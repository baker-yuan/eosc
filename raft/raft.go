package raft

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/eosc/log"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"go.etcd.io/etcd/client/pkg/v3/types"
	"go.etcd.io/etcd/pkg/wait"
	"go.etcd.io/etcd/raft/v3"

	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/rafthttp"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	stats "go.etcd.io/etcd/server/v3/etcdserver/api/v2stats"
	"go.etcd.io/etcd/server/v3/wal"
	"go.etcd.io/etcd/server/v3/wal/walpb"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

var defaultSnapshotCount uint64 = 10000
var snapshotCatchUpEntriesN uint64 = 10000

var retryFrequency time.Duration = 2000

// raft节点结构
type raftNode struct {
	// 节点ID
	nodeID uint64

	// eosc 服务相关
	Service IService

	// 节点相关
	node raft.Node

	// 当前leader
	lead uint64
	// 节点状态
	confState raftpb.ConfState

	// 集群相关
	join bool
	// 集群模式
	isCluster bool
	// peers列表
	peers *Peers

	// 快照相关
	snapShotter *snap.Snapshotter
	// 达到snapCount时本地保存快照
	snapCount uint64
	// 快照记录的索引
	snapshotIndex uint64
	snapDir       string // path to snapshot directory

	// 日志与内存
	raftStorage  *MemoryStorage
	wal          *wal.WAL
	waldir       string // path to WAL directory
	appliedIndex uint64

	// 与其他节点通信
	transport *rafthttp.Transport
	stopC     chan struct{} // signals proposal channel closed
	httpStopC chan struct{} // signals http server to shutdown
	httpDoneC chan struct{} // signals http server shutdown complete

	// 日志相关，后续改为eosc_log
	logger *zap.Logger
	waiter wait.Wait
	active bool
}

// JoinCluster 新建一个加入已知集群的请求
// join和isCluster一定为true
// 1、应用于新建一个想要加入已知集群的节点，会向已知节点发送请求获取id等新建节点信息
// 已知节点如果还处于非集群模式，会先切换成集群模式
// 2、也可以用于节点crash后的重启处理
func JoinCluster(local string, target string, service IService) (*raftNode, error) {
	msg := &JoinMsg{
		Host: local,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	// 向集群中的某个节点发送要加入的请求
	resp, err := http.Post(fmt.Sprintf("%s/join", target), "application/json;charset=utf-8", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &Response{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return nil, err
	}
	if res.Code == "000000" {
		err = json.Unmarshal(res.Data, msg)
		if err != nil {
			return nil, err
		}
		log.Infof("receive join message id(%d), peers number(%d)", msg.Id, len(msg.Peers))
		return joinAndCreateRaft(msg.Id, local, service, msg.Peers), nil
	} else {
		return nil, fmt.Errorf(res.Msg)
	}
}

// joinAndCreateRaft 收到id，peer等信息后，新建并加入集群，新建日志文件等处理
func joinAndCreateRaft(id int, host string, service IService, peerList map[uint64]string) *raftNode {
	rc := &raftNode{
		nodeID:    uint64(id),
		Service:   service,
		join:      true,
		isCluster: true,
		// 日志文件目前直接以id命名初始化
		waldir:    fmt.Sprintf("eosc-%d", id),
		snapDir:   fmt.Sprintf("eosc-%d-snap", id),
		snapCount: defaultSnapshotCount,
		stopC:     make(chan struct{}),
		httpStopC: make(chan struct{}),
		httpDoneC: make(chan struct{}),
		logger:    zap.NewExample(),
		waiter:    wait.New(),
		lead:      0,
		active:    true,
	}

	// 确保peer一定有节点本身
	peerList[rc.nodeID] = host
	rc.peers = NewPeers(peerList, len(peerList))
	// 创建并启动 transport 实例，该负责节点之间的网络通信，
	// 非集群模式下主要是为了listener的Handler处理，监听join请求，此时transport尚未start
	rc.transport = &rafthttp.Transport{
		Logger:             rc.logger,
		ID:                 types.ID(rc.nodeID),
		ClusterID:          0x1000,
		Raft:               rc,
		ServerStats:        stats.NewServerStats("", ""),
		LeaderStats:        stats.NewLeaderStats(zap.NewExample(), strconv.Itoa(int(rc.nodeID))),
		ErrorC:             make(chan error),
		DialRetryFrequency: rate.Every(retryFrequency * time.Millisecond),
	}

	// raft启动
	go rc.serveRaft()
	go rc.startRaft()
	return rc
}

// CreateRaftNode 初始化节点
// peers至少会包括节点本身，如果join为true，isCluster也默认为true
// 已经切换到集群模式下的节点不会回到非集群模式，除非改变节点ID或删除相关日志文件
// 1、创建非集群节点,isCluster为false，此时peers可为空
// 2、创建集群节点，isCluster为true，若此时peers为空(或仅有节点本身)，表示该集群仅有一个节点
// peers也可以是其余集群中的其他节点，表示这是一个多节点集群，此时其他节点也需通过同样的配置和方式启动，
// 推荐使用JoinCluster来新建多节点集群节点
// 3、创建加入已知集群的节点，join为true，isCluster为true，此时peers需包括其他节点地址，推荐使用JoinCluster来新建非单点集群节点
func CreateRaftNode(id int, host string, service IService, peers string, keys string, join bool, isCluster bool) (*raftNode, error) {
	rc := &raftNode{
		nodeID:    uint64(id),
		Service:   service,
		join:      join,
		waldir:    fmt.Sprintf("eosc-%d", id), // 日志文件路径
		snapDir:   fmt.Sprintf("eosc-%d-snap", id),
		snapCount: defaultSnapshotCount,
		stopC:     make(chan struct{}),
		httpStopC: make(chan struct{}),
		httpDoneC: make(chan struct{}),
		logger:    zap.NewExample(),
		waiter:    wait.New(),
		lead:      0,
		active:    true,
	}
	// 建过集群的节点不能再换回去（暂时采用该方案）
	if rc.join || wal.Exist(rc.waldir) {
		rc.isCluster = true
	}

	log.Infof("current mode is cluster %v.", rc.isCluster)

	peerList, err := Adjust(rc.nodeID, host, peers, keys, rc.isCluster)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rc.peers = NewPeers(peerList, len(peerList))

	// 创建并启动 transport 实例，负责节点之间的网络通信，
	// 非集群模式下主要是为了listener的Handler处理，监听join请求，此时transport尚未start
	rc.transport = &rafthttp.Transport{
		Logger:             rc.logger,
		ID:                 types.ID(rc.nodeID),
		ClusterID:          0x1000,
		Raft:               rc,
		ServerStats:        stats.NewServerStats("", ""),
		LeaderStats:        stats.NewLeaderStats(zap.NewExample(), strconv.Itoa(int(rc.nodeID))),
		ErrorC:             make(chan error),
		DialRetryFrequency: rate.Every(2000 * time.Millisecond),
	}

	// 监听节点端口，用transport处理节点通信，此时这种情况下只是监听join
	go rc.serveRaft()

	// 集群模式下启动节点，已经是认为是集群的节点（有日志文件存在）也会启动集群模式
	if isCluster || wal.Exist(rc.waldir) {
		go rc.startRaft()
	}
	return rc, nil
}

// startRaft 启动raft服务，在集群模式下启动或join模式下启动
// 非集群模式下启动的节点不会调用该start函数
func (rc *raftNode) startRaft() {
	log.Info("start raft Service")

	// 判断快照文件夹是否存在，不存在则创建
	if !fileutil.Exist(rc.snapDir) {
		if err := os.Mkdir(rc.snapDir, 0750); err != nil {
			log.Fatalf("eosc: cannot create dir for snapshot (%v)", err)
		}
	}
	// 新建快照管理
	rc.snapShotter = snap.New(zap.NewExample(), rc.snapDir)

	// 判断是否有日志文件目录
	oldWal := wal.Exist(rc.waldir)

	// 将日志wal重写入raftNode实例中，读取快照和日志，并初始化raftStorage
	rc.wal = rc.replayWAL()

	// 集群模式下启动节点的时候，重新reload快照到service中
	// TODO 非集群想要切换成集群的时候，要么这里做进一步校验，要么切换前先存好快照和日志
	err := rc.ReadSnap(rc.snapShotter)
	if err != nil {
		log.Info("reload snap to Service error:", err)
	}

	// 节点配置
	c := &raft.Config{
		ID:                        rc.nodeID,
		ElectionTick:              10,
		HeartbeatTick:             1,
		Storage:                   rc.raftStorage,
		MaxSizePerMsg:             1024 * 1024,
		MaxInflightMsgs:           256,
		MaxUncommittedEntriesSize: 1 << 30,
	}
	peersList := rc.peers.GetAllPeers()
	// 启动node节点
	if rc.join || oldWal {
		// 选择加入集群或已有日志消息（曾经切换过集群模式）
		rc.node = raft.RestartNode(c)
	} else {

		// 启动节点时添加peers
		rpeers := make([]raft.Peer, 0, rc.peers.GetPeerNum())
		for id := range peersList {
			rpeers = append(rpeers, raft.Peer{ID: uint64(id)})
		}
		// 新开一个集群
		rc.node = raft.StartNode(c, rpeers)
	}

	// 开启节点间通信
	// 通信实例开始运行
	err = rc.transport.Start()
	if err != nil {
		log.Info("transport start error:", err)
	}
	// 与集群中的其他节点建立通信

	for k, v := range peersList {
		// transport加入peer列表，节点本身不添加
		if k != rc.nodeID {
			rc.transport.AddPeer(types.ID(k), []string{v})
		}
	}

	// 开始读ready
	go rc.serveChannels()
}

// 监听ready通道，集群模式下会开始监听
func (rc *raftNode) serveChannels() {
	sn, err := rc.raftStorage.Snapshot()
	if err != nil {
		log.Panic(err)
	}
	// 获取当前的日志最新的状态信息，index和term等
	rc.confState = sn.Metadata.ConfState
	rc.snapshotIndex = sn.Metadata.Index
	rc.appliedIndex = sn.Metadata.Index

	defer rc.wal.Close()
	// 心跳定时器
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			rc.node.Tick()
		case rd := <-rc.node.Ready():
			if rd.SoftState != nil {
				rc.lead = rd.SoftState.Lead
			}
			//islead = rd.RaftState == raft.StateLeader
			rc.saveToStorage(rd)

			// 将信息发给其他节点
			rc.transport.Send(rd.Messages)

			// 处理需要commit的日志
			ok := rc.publishEntries(rc.entriesToApply(rd.CommittedEntries))
			if !ok {
				// 此时节点停止
				rc.stop()
				return
			}
			// 必要时生成快照
			rc.maybeTriggerSnapshot()
			// 告知底层raft该Ready已处理完
			rc.node.Advance()
		// transport出错
		case err = <-rc.transport.ErrorC:
			log.Info(err.Error())
			rc.stop()
			return
		case <-rc.stopC:
			return
		}
	}
}

// publishEntries writes committed log entries to commit channel and returns
// whether all entries could be published.
// 日志提交处理
func (rc *raftNode) publishEntries(ents []raftpb.Entry) bool {
	if len(ents) == 0 {
		return true
	}
	for i := range ents {
		switch ents[i].Type {
		case raftpb.EntryNormal:
			if len(ents[i].Data) == 0 {
				// ignore empty messages
				continue
			}
			m := &Message{}
			var err error
			err = m.Decode(ents[i].Data)
			if err != nil {
				log.Info(err)
				continue
			}
			err = rc.Service.CommitHandler(m.Cmd, m.Data)
			if m.Type == INIT && m.From == rc.nodeID {
				// 释放InitSend方法的等待，仅针对切换集群的对应节点
				if err != nil {
					log.Info(err)
					rc.waiter.Trigger(uint64(m.From), err.Error())
					continue
				} else {
					rc.waiter.Trigger(uint64(m.From), "")
				}
			}
		case raftpb.EntryConfChange:
			var cc raftpb.ConfChange
			cc.Unmarshal(ents[i].Data)

			rc.confState = *rc.node.ApplyConfChange(cc)
			// eosc: 修改快照里的confState
			rc.raftStorage.UpdateConState(&rc.confState)
			switch cc.Type {
			// 新增节点
			case raftpb.ConfChangeAddNode:
				if len(cc.Context) > 0 {
					// transport不需要加自己
					if cc.NodeID != rc.nodeID {
						p := rc.transport.Get(types.ID(cc.NodeID))
						// 不存在才加进去
						if p == nil {
							rc.transport.AddPeer(types.ID(cc.NodeID), []string{string(cc.Context)})
						}
					}
					_, ok := rc.peers.GetPeerByID(cc.NodeID)
					if !ok {
						// 已存在，不再新增
						rc.peers.SetPeer(cc.NodeID, string(cc.Context))
					}
				}
			case raftpb.ConfChangeRemoveNode:
				if cc.NodeID == rc.nodeID {
					log.Info("current node has been removed from the cluster!")
					return false
				}
				p := rc.transport.Get(types.ID(cc.NodeID))
				if p != nil {
					// 存在才移除
					rc.transport.RemovePeer(types.ID(cc.NodeID))
				}
				_, ok := rc.peers.GetPeerByID(cc.NodeID)
				if ok {
					// 存在，减去
					rc.peers.DeletePeerByID(cc.NodeID)
				}
			}
		}
	}
	// after commit, update appliedIndex
	rc.appliedIndex = ents[len(ents)-1].Index
	return true
}

// 处理本节点需要committed的日志
func (rc *raftNode) entriesToApply(ents []raftpb.Entry) (nents []raftpb.Entry) {
	if len(ents) == 0 {
		return ents
	}
	firstIdx := ents[0].Index
	if firstIdx > rc.appliedIndex+1 {
		log.Fatalf("first index of committed entry[%d] should <= progress.appliedIndex[%d]+1", firstIdx, rc.appliedIndex)
	}
	if rc.appliedIndex-firstIdx+1 < uint64(len(ents)) {
		nents = ents[rc.appliedIndex-firstIdx+1:]
	}
	return nents
}

// 日志/快照文件相关

// ReadSnap 读取快照内容到service
func (rc *raftNode) ReadSnap(snapshotter *snap.Snapshotter) error {
	// 读取快照的所有内容
	snapshot, err := snapshotter.Load()
	// 快照不存在
	if err == snap.ErrNoSnapshot {
		return nil
	}
	if err != nil {
		return err
	}
	// 快照不为空的话写进service
	if snapshot != nil {
		// 将快照内容缓存到service中
		log.Infof("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
		err = rc.Service.ResetSnap(snapshot.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

// saveToStorage 内存信息存储
func (rc *raftNode) saveToStorage(rd raft.Ready) {
	var err error
	err = rc.wal.Save(rd.HardState, rd.Entries)
	if err != nil {
		log.Fatal("save wal error: ", err)
	}
	// 安装快照信息
	if !raft.IsEmptySnap(rd.Snapshot) {
		log.Info("begin save snap")
		err = rc.saveSnap(rd.Snapshot)
		if err != nil {
			log.Fatal("save snap error: ", err)
		}
		// 有快照信息则根据快照修改内存
		err = rc.raftStorage.ApplySnapshot(rd.Snapshot)
		if err != nil {
			log.Fatal("raftStorage apply snapshot error: ", err)
		}
		rc.publishSnapshot(rd.Snapshot)
	}
	// 日志信息写入节点缓存
	err = rc.raftStorage.Append(rd.Entries)
	if err != nil {
		log.Fatal("raftStorage append entries error: ", err)
	}
}

// 保存快照文件
func (rc *raftNode) saveSnap(snap raftpb.Snapshot) error {
	walSnap := walpb.Snapshot{
		Index:     snap.Metadata.Index,
		Term:      snap.Metadata.Term,
		ConfState: &snap.Metadata.ConfState,
	}
	// save the snapshot file before writing the snapshot to the wal.
	// This makes it possible for the snapshot file to become orphaned, but prevents
	// a WAL snapshot entry from having no corresponding snapshot file.
	if err := rc.snapShotter.SaveSnap(snap); err != nil {
		return err
	}
	if err := rc.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}
	// 压缩日志
	return rc.wal.ReleaseLockTo(snap.Metadata.Index)
}

// service 从快照中重取数据
func (rc *raftNode) publishSnapshot(snapshotToSave raftpb.Snapshot) {
	if raft.IsEmptySnap(snapshotToSave) {
		return
	}
	log.Infof("publishing snapshot at index %d", rc.snapshotIndex)
	defer log.Infof("finished publishing snapshot at index %d", rc.snapshotIndex)

	if snapshotToSave.Metadata.Index <= rc.appliedIndex {
		log.Fatalf("snapshot index [%d] should > progress.appliedIndex [%d]", snapshotToSave.Metadata.Index, rc.appliedIndex)
	}
	err := rc.ReadSnap(rc.snapShotter)
	if err != nil {
		log.Info("read snap from snap shotter error:", err)
	}
	rc.confState = snapshotToSave.Metadata.ConfState
	rc.snapshotIndex = snapshotToSave.Metadata.Index
	rc.appliedIndex = snapshotToSave.Metadata.Index
}

// 保存现有快照
func (rc *raftNode) maybeTriggerSnapshot() {
	// 还不到保存快照的数里
	if rc.appliedIndex-rc.snapshotIndex <= rc.snapCount {
		return
	}
	log.Infof("start snapshot [applied index: %d | last snapshot index: %d]", rc.appliedIndex, rc.snapshotIndex)

	// 获取service中的信息
	data, err := rc.Service.GetSnapshot()
	if err != nil {
		log.Panic(err)
	}
	// 利用现有信息生成要保存的快照信息
	snapContent, err := rc.raftStorage.CreateSnapshot(rc.appliedIndex, &rc.confState, data)
	if err != nil {
		log.Panic(err)
	}
	if err = rc.saveSnap(snapContent); err != nil {
		log.Panic(err)
	}
	// 暂时先不做日志的压缩处理，有bug如下：
	// 已有集群中的节点生成快照后，后续新节点加入集群时无法成功同步
	// 修复处理：集群节点配置更新的时候也更新raftStorage中已有快照的confState
	compactIndex := uint64(1)
	if rc.appliedIndex > snapshotCatchUpEntriesN {
		compactIndex = rc.appliedIndex - snapshotCatchUpEntriesN
	}
	if err = rc.raftStorage.Compact(compactIndex); err != nil {
		log.Panic(err)
	}
	log.Infof("compacted log at index %d", compactIndex)
	rc.snapshotIndex = rc.appliedIndex
}

// 从现有文件中读取日志
func (rc *raftNode) replayWAL() *wal.WAL {
	// 先获取现有快照
	snapshot := rc.loadSnapshot()
	// 再获取现有日志
	w := rc.openWAL(snapshot)
	_, st, ents, err := w.ReadAll()
	if err != nil {
		log.Fatalf("eosc: failed to read WAL (%v)", err)
	}

	// 节点日志缓存初始化
	rc.raftStorage = NewMemoryStorage()
	if snapshot != nil {

		err = rc.raftStorage.ApplySnapshot(*snapshot)
		if err != nil {
			log.Infof("eosc: failed to apply snapshot for raftStorage (%v)", err)
		}
	}
	err = rc.raftStorage.SetHardState(st)
	if err != nil {
		log.Infof("eosc: failed to set hardState for raftStorage (%v)", err)
	}
	// append to storage so raft starts at the right place in log
	err = rc.raftStorage.Append(ents)
	if err != nil {
		log.Infof("eosc: failed to append ents for raftStorage (%v)", err)
	}
	return w
}

// 读取快照文件
func (rc *raftNode) loadSnapshot() *raftpb.Snapshot {
	if wal.Exist(rc.waldir) {
		walSnaps, err := wal.ValidSnapshotEntries(rc.logger, rc.waldir)
		if err != nil {
			log.Fatalf("eosc: error listing snapshots (%v)", err)
		}
		// 获取最新的快照
		snapshot, err := rc.snapShotter.LoadNewestAvailable(walSnaps)
		if err != nil && err != snap.ErrNoSnapshot {
			log.Fatalf("eosc: error loading snapshot (%v)", err)
		}
		return snapshot
	}
	return &raftpb.Snapshot{}
}

// 读取(创建)wal日志文件
func (rc *raftNode) openWAL(snapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(rc.waldir) {
		// 创建本地文件
		if err := os.Mkdir(rc.waldir, 0750); err != nil {
			log.Fatalf("eosc: cannot create dir for wal (%v)", err)
		}
		// 创建wal日志对象
		w, err := wal.Create(zap.NewExample(), rc.waldir, nil)
		if err != nil {
			log.Fatalf("eosc: create wal error (%v)", err)
		}
		w.Close()
	}
	// 该结构用于日志对象定位快照索引，方便日志读取
	walSnap := walpb.Snapshot{}
	// 获取目前快照的最新记录
	if snapshot != nil {
		walSnap.Index, walSnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	log.Infof("loading WAL at term %d and index %d", walSnap.Term, walSnap.Index)
	// 开启日志
	w, err := wal.Open(zap.NewExample(), rc.waldir, walSnap)
	if err != nil {
		log.Fatalf("eosc: error loading wal (%v)", err)
	}
	return w
}

// 停止服务相关(暂时不直接关闭程序)
// stop closes http and stops raft.
func (rc *raftNode) stop() {
	rc.stopHTTP()
	close(rc.stopC)
	rc.node.Stop()
	rc.active = false
	//os.Exit(0)
}

// 停止http服务
func (rc *raftNode) stopHTTP() {
	rc.transport.Stop()
	close(rc.httpStopC)
	<-rc.httpDoneC
}

// Send 客户端发送propose请求的处理
// 由客户端API调用或Leader收到转发后调用
// 如果是非集群模式(isCluster为false)，直接处理(即service.ProcessHandler后直接service.CommitHandler)
// 如果是集群模式，分两种情况
// 1、当前节点是leader，经service.ProcessHandler后由node.Propose处理后返回，
// 后续会由各个节点的node.Ready读取后进行Commit时由service.CommitHandler处理
// 2、当前节点不是leader，获取当前leader节点地址，转发至leader进行处理(rc.proposeHandler)，
// leader收到请求后经service.ProcessHandler后由node.Propose处理后返回，
// 后续会由各个节点的node.Ready读取后进行Commit时由service.CommitHandler处理
func (rc *raftNode) Send(command string, send []byte) error {
	// 移除节点后，因为有外部api，故不会停止程序，以此做隔离
	if !rc.active {
		return fmt.Errorf("current node is stop")
	}
	// 非集群模式下直接处理
	if !rc.isCluster {
		cmd, data, err := rc.Service.ProcessHandler(command, send)
		if err != nil {
			return err
		}
		return rc.Service.CommitHandler(cmd, data)
	}
	// 集群模式下的处理
	addr, isLeader, err := rc.getLeader()
	if err != nil {
		return err
	}
	log.Infof("send:leader is node(%d)", rc.lead)
	// 如果自己本身就是leader，直接处理，否则转发由leader处理
	if isLeader {
		// Service.ProcessHandler要么leader执行，要么非集群模式下自己执行
		cmd, data, err := rc.Service.ProcessHandler(command, send)
		if err != nil {
			return err
		}
		m := &Message{
			From: rc.nodeID,
			Type: PROPOSE,
			Cmd:  cmd,
			Data: data,
		}
		b, err := m.Encode()
		if err != nil {
			return err
		}
		return rc.node.Propose(context.TODO(), b)
	} else {
		return rc.postMessage(addr, command, send)
	}
}

// GetPeers 获取集群的peer列表，供API调用
func (rc *raftNode) GetPeers() (map[uint64]string, int, error) {
	if !rc.active {
		return nil, 0, fmt.Errorf("current node is stop")
	}
	peerList := rc.peers.GetAllPeers()
	peerCount := rc.peers.GetConfigCount()
	return peerList, peerCount, nil
}

// AddConfigChange 客户端发送增加/删除节点的发送处理
func (rc *raftNode) AddConfigChange(nodeID uint64, host string) error {
	if !rc.active {
		return fmt.Errorf("current node is stop")
	}
	if !rc.isCluster {
		return fmt.Errorf("current node is not cluster mode")
	}
	p := rc.transport.Get(types.ID(nodeID))
	if p != nil {
		return fmt.Errorf("added node is existed")
	}
	cc := raftpb.ConfChange{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  nodeID,
		Context: []byte(host),
		ID:      uint64(rc.peers.GetConfigCount() + 1),
	}
	return rc.node.ProposeConfChange(context.TODO(), cc)
}

// DeleteConfigChange 客户端发送删除节点的发送处理
func (rc *raftNode) DeleteConfigChange(nodeID uint64) error {
	if !rc.active {
		return fmt.Errorf("current node is stop")
	}
	if !rc.isCluster {
		return fmt.Errorf("current node is not cluster mode")
	}
	p := rc.transport.Get(types.ID(nodeID))
	if p == nil && nodeID != uint64(rc.nodeID) {
		return fmt.Errorf("deleted node is not existed")
	}
	cc := raftpb.ConfChange{
		Type:   raftpb.ConfChangeRemoveNode,
		NodeID: nodeID,
		ID:     uint64(rc.peers.GetConfigCount() + 1),
	}
	return rc.node.ProposeConfChange(context.TODO(), cc)
}

// InitSend 切换集群时调用，一般一个集群仅调用一次
// 将service现有的缓存信息(基于service.GetInit获取)加载到日志中，便于其他节点同步
// 此时节点刚切换到集群状态，一般会是日志中的第一条信息
// 并通过rc.waiter等待service.ProcessInit处理完后进行后续操作(同步等待)
func (rc *raftNode) InitSend() error {
	// 集群模式初始化的时候才会调
	if !rc.isCluster {
		return fmt.Errorf("need to change cluster mode")
	}
	cmd, data, err := rc.Service.GetInit()
	if err != nil {
		return err
	}
	m := &Message{
		From: rc.nodeID,
		Type: INIT,
		Cmd:  cmd,
		Data: data,
	}
	b, err := m.Encode()
	if err != nil {
		return err
	}
	err = rc.node.Propose(context.TODO(), b)
	if err != nil {
		return err
	}
	// 等待处理完
	c := rc.waiter.Register(rc.nodeID)
	res := <-c
	str, ok := res.(string)
	if !ok {
		return fmt.Errorf("init send wait channel interface assert error")
	}
	if len(str) > 0 {
		return fmt.Errorf(str)
	}
	return nil
}

// changeCluster 切换集群时调用，一般是非集群节点收到其他节点的join请求时触发(rc.joinHandler)
// 开始运行集群节点,新建日志文件，启动transport和node，
// 并开始监听node.ready,将现有缓存加入日志中rc.InitSend
func (rc *raftNode) changeCluster() error {
	log.Info("change cluster mode")
	rc.isCluster = true
	// 判断快照文件夹是否存在，不存在则创建
	if !fileutil.Exist(rc.snapDir) {
		if err := os.Mkdir(rc.snapDir, 0750); err != nil {
			return fmt.Errorf("eosc: node(%d) cannot create dir for snapshot (%v)", rc.nodeID, err)
		}
	}
	// 新建快照管理
	rc.snapShotter = snap.New(zap.NewExample(), rc.snapDir)

	// 判断是否有日志文件目录，此时应该是没有的
	oldwal := wal.Exist(rc.waldir)
	if oldwal {
		return fmt.Errorf("node(%d) has been cluster mode, wal is existed", rc.nodeID)
	}
	// 将日志wal重写入raftNode实例中，读取快照和日志，并初始化raftStorage,此处主要是新建日志文件
	rc.wal = rc.replayWAL()
	// 节点配置
	c := &raft.Config{
		ID:                        uint64(rc.nodeID),
		ElectionTick:              10,
		HeartbeatTick:             1,
		Storage:                   rc.raftStorage,
		MaxSizePerMsg:             1024 * 1024,
		MaxInflightMsgs:           256,
		MaxUncommittedEntriesSize: 1 << 30,
	}
	// 启动节点时添加peers，一般情况下此时只有自己
	rpeers := make([]raft.Peer, 0, rc.peers.GetPeerNum())
	peerList := rc.peers.GetAllPeers()
	for id := range peerList {
		rpeers = append(rpeers, raft.Peer{ID: uint64(id)})
	}
	// 启动node节点
	// 新开一个集群
	rc.node = raft.StartNode(c, rpeers)

	// 通信实例开始运行
	err := rc.transport.Start()
	if err != nil {
		return err
	}
	// 与集群中的其他节点建立通信
	for k, v := range peerList {
		// transport加入peer列表，节点本身不添加
		if k != rc.nodeID {
			rc.transport.AddPeer(types.ID(k), []string{v})
		}
	}
	// 读ready
	go rc.serveChannels()
	log.Info("change cluster mode successfully")
	// 开始打包处理初始化信息
	err = rc.InitSend()
	if err != nil {
		return err
	}
	return nil
}

// 通信相关
// serveRaft 用于监听当前节点的指定端口，处理与其他节点的网络连接，需更改
func (rc *raftNode) serveRaft() {
	log.Info("eosc: start raft serve listener")
	v, ok := rc.peers.GetPeerByID(rc.nodeID)
	if !ok {
		log.Fatalf("eosc: Failed read current node(%d) url ", rc.nodeID)
	}
	addr, err := url.Parse(v)
	if err != nil {
		log.Fatalf("eosc: Failed parsing URL (%v)", err)
	}
	ln, err := newStoppableListener(addr.Host, rc.httpStopC)
	if err != nil {
		log.Fatalf("eosc: Failed to listen rafthttp (%v)", err)
	}
	// 调用rc.transport.Handler()对连接进行处理
	err = (&http.Server{Handler: rc.Handler()}).Serve(ln)
	select {
	case <-rc.httpStopC:
	default:
		log.Fatalf("eosc: Failed to serve rafthttp (%v)", err)
	}
	close(rc.httpDoneC)
}

// Handler http请求处理
func (rc *raftNode) Handler() http.Handler {
	sm := http.NewServeMux()
	// 其他节点加入集群的处理
	sm.HandleFunc("/join", rc.joinHandler)
	// 其他节点转发到leader的处理
	sm.HandleFunc("/propose", rc.proposeHandler)
	sm.Handle("/", rc.transport.Handler())
	return sm
}

// joinHandler 收到其他节点加入集群的处理
// 1、如果已经是集群模式，直接返回相关id，peer等信息方便处理
// 2、如果不是集群模式，先切换集群rc.changeCluster,再返回相关信息
// 3、该处理也可应用于集群节点crash后的重启
func (rc *raftNode) joinHandler(w http.ResponseWriter, r *http.Request) {

	joinMsg := &JoinMsg{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(body, joinMsg)
	if err != nil {
		writeResult(w, &Response{
			Code: "111111",
			Msg:  err.Error(),
		})
		return
	}
	log.Infof("host(%s) apply join the cluster", joinMsg.Host)
	// 先判断是不是集群模式
	// 是的话返回要加入的相关信息
	// 不是的话先切换集群模式，再初始化startRaft()，再返回加入的相关信息
	if !rc.isCluster {
		// 非集群模式，先本节点切换成集群模式
		err = rc.changeCluster()
		if err != nil {
			// 切换错误
			writeResult(w, &Response{
				Code: "111111",
				Msg:  err.Error(),
			})
			return
		}
	}
	// 切换完了，开始新增对应节点并返回新增条件信息
	joinMsg.Peers = rc.peers.GetAllPeers()
	if id, exist := rc.peers.CheckExist(joinMsg.Host); exist {
		// 已经在集群中的了，直接返回信息
		joinMsg.Id = int(id)
		b, _ := json.Marshal(joinMsg)

		writeResult(w, &Response{
			Code: "000000",
			Msg:  "success",
			Data: b,
		})
		return
	}
	// 现有的变更id+1
	joinMsg.Id = rc.peers.GetConfigCount() + 1
	// 已经是集群了，发送新增节点的消息后返回
	err = rc.AddConfigChange(uint64(joinMsg.Id), joinMsg.Host)
	if err != nil {
		writeResult(w, &Response{
			Code: "111111",
			Msg:  err.Error(),
		})
		return
	}
	b, _ := json.Marshal(joinMsg)
	writeResult(w, &Response{
		Code: "000000",
		Msg:  "success",
		Data: b,
	})
	return
}

// proposeHandler 其他节点转发到leader的propose处理，由rc.Send触发
func (rc *raftNode) proposeHandler(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Code: "111111",
	}
	// 只有leader才会收到该消息
	_, isLeader, err := rc.getLeader()
	if !isLeader {
		// no current leader
		res.Msg = "can not find leader"
		writeResult(w, res)
		return
	}
	if err != nil {
		res.Msg = err.Error()
		writeResult(w, res)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		res.Msg = err.Error()
		writeResult(w, res)
		return
	}

	msg := &ProposeMsg{}
	err = json.Unmarshal(b, msg)
	if err != nil {
		res.Msg = err.Error()
		writeResult(w, res)
		return
	}
	log.Infof("receive propose request from node(%d)", msg.From)
	err = rc.Send(msg.Cmd, msg.Data)
	if err != nil {
		res.Msg = err.Error()
	} else {
		res.Code = "000000"
		res.Msg = "success"
	}
	writeResult(w, res)
}

// 工具方法
// postMessage 转发消息，基于json
func (rc *raftNode) postMessage(addr string, command string, data []byte) error {
	// 转给leader
	m := &ProposeMsg{
		From: rc.nodeID,
		To:   int(rc.lead),
		Cmd:  command,
		Data: data,
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("%s/propose", addr), "application/json;charset=utf-8", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	res := &Response{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return err
	}
	if res.Code == "000000" {
		return nil
	} else {
		msg := res.Msg
		return fmt.Errorf(msg)
	}
}

// getLeader 获取leader地址以及判断当前节点是不是leader
func (rc *raftNode) getLeader() (string, bool, error) {
	if rc.lead == raft.None {
		if rc.node.Status().Lead == raft.None {
			return "", false, fmt.Errorf("current node(%d) has no leader", rc.lead)
		} else {
			rc.lead = rc.node.Status().Lead
		}
	}
	flag := false
	if rc.nodeID == rc.lead {
		flag = true
	}
	v, ok := rc.peers.GetPeerByID(rc.lead)
	if !ok {
		return "", flag, fmt.Errorf("current node has no leader(%d) host", rc.lead)
	}
	return v, flag, nil
}

// Adjust 参数校验与调整
func Adjust(id uint64, host string, peers string, keys string, isCluster bool) (map[uint64]string, error) {
	peerList := make(map[uint64]string)
	peerList[id] = host
	// 非集群模式不需要peer列表，此时peer列表仅有节点本身
	if !isCluster {
		return peerList, nil
	}
	clusterList := strings.Split(peers, ",")
	keyList := strings.Split(keys, ",")
	if len(keyList) != len(clusterList) {
		return nil, fmt.Errorf("the length of keys is %d while the length of peers is %d", len(keyList), len(clusterList))
	} else {
		for i, key := range keyList {
			k, err := strconv.Atoi(key)
			if err != nil {
				return nil, err
			}
			if _, ok := peerList[uint64(k)]; !ok {
				peerList[uint64(k)] = clusterList[i]
			}
		}
	}
	return peerList, nil
}

// writeResult 响应http
func writeResult(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}

func (rc *raftNode) Process(ctx context.Context, m raftpb.Message) error { return rc.node.Step(ctx, m) }
func (rc *raftNode) IsIDRemoved(id uint64) bool                          { return false }
func (rc *raftNode) ReportUnreachable(id uint64)                         { rc.node.ReportUnreachable(id) }
func (rc *raftNode) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	rc.node.ReportSnapshot(id, status)
}
