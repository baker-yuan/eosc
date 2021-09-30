package process_master

import (
	"bytes"
	"context"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/eolinker/eosc/traffic"

	"github.com/eolinker/eosc/log"

	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/service"
)

var _ service.WorkerServiceClient = (*WorkerController)(nil)

type WorkerProcessController interface {
	Stop()
	NewWorker() error
	Start()
}
type WorkerController struct {
	locker            sync.Mutex
	dms               []eosc.IDataMarshaller
	current           *WorkerProcess
	expireWorkers     []*WorkerProcess
	trafficController traffic.IController
	isStop            bool
	checkClose        chan int
}

func NewWorkerController(trafficController traffic.IController, dms ...eosc.IDataMarshaller) *WorkerController {
	dmsAll := make([]eosc.IDataMarshaller, 0, len(dms)+1)
	dmsAll = append(dmsAll, trafficController)
	for _, v := range dms {
		dmsAll = append(dmsAll, v)
	}

	return &WorkerController{
		trafficController: trafficController,
		dms:               dmsAll,
		checkClose:        make(chan int, 0),
	}
}
func (wc *WorkerController) Stop() {
	wc.locker.Lock()
	defer wc.locker.Unlock()

	if wc.isStop {
		return
	}
	close(wc.checkClose)
	wc.isStop = true
	if wc.current != nil {
		wc.current.Close()
		wc.expireWorkers = append(wc.expireWorkers, wc.current)
		wc.current = nil
	}

}
func (wc *WorkerController) check(w *WorkerProcess) {
	err := w.cmd.Wait()
	if err != nil {
		log.Warn("worker exit:", err)
	}

	if wc.getClient() == w {
		err := wc.NewWorker()
		if err != nil {
			log.Error("worker create:", err)
		}
	} else {

		for i, v := range wc.expireWorkers {
			if v == w {
				wc.expireWorkers = append(wc.expireWorkers[:i], wc.expireWorkers[i+1:]...)
			}
		}
	}
}
func (wc *WorkerController) Start() {

	wc.NewWorker()

	go func() {
		t := time.NewTicker(time.Second)
		in := &service.WorkerHelloRequest{
			Hello: "hello",
		}
		next := time.NewTimer(time.Second * 5)
		next.Stop()
		var last []int = nil
		defer next.Stop()
		defer t.Stop()
		for {
			select {
			case <-t.C:
				client := wc.getClient()
				if client != nil {
					response, err := client.Ping(context.TODO(), in)
					if err != nil {
						continue
					}
					ports := sortAndSet(response.Resource.Port)
					if !equal(last, ports) {
						last = ports
						next.Reset(time.Second * 5)
					}
				}
			case <-next.C:
				{
					isCreate, err := wc.trafficController.Reset(last)
					if err != nil {
						continue
					}
					if isCreate {
						wc.NewWorker()
					}
				}
			case <-wc.checkClose:
				return
			}
		}

	}()
}
func (wc *WorkerController) NewWorker() error {
	wc.locker.Lock()
	defer wc.locker.Unlock()
	err := wc.new()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			_, err := wc.current.Ping(context.TODO(), &service.WorkerHelloRequest{Hello: "hello"})
			if err != nil {
				//log.Debug("work controller ping: ", err)
				continue
			}

			return nil
		}
	}
	return nil
}
func (wc *WorkerController) new() error {
	log.Debug("create worker process start")
	buf := bytes.NewBuffer(nil)
	var fileAll []*os.File
	index := 3
	for _, dm := range wc.dms {
		data, files, err := dm.Encode(index)
		log.Debugf("encode:data[%d] file[%d]", len(data), len(files))
		if err != nil {
			log.Warn("create worker process fail:", err)
			return err
		}
		index += len(files)
		fileAll = append(fileAll, files...)
		buf.Write(data)

	}

	workerProcess, err := wc.newWorkerProcess(buf, fileAll)
	if err != nil {
		return err
	}

	old := wc.current
	wc.current = workerProcess
	go wc.check(wc.current)

	if old != nil {
		old.Close()
	}

	return nil
}

func (wc *WorkerController) getClient() *WorkerProcess {
	wc.locker.Lock()
	defer wc.locker.Unlock()
	if wc.current == nil {
		return nil
	}
	return wc.current
}

func equal(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}
func sortAndSet(vs []int32) []int {
	if len(vs) == 0 {
		return nil
	}

	m := make(map[int]int)
	for _, v := range vs {
		m[int(v)] = 1
	}
	rs := make([]int, 0, len(m))
	for v := range m {
		rs = append(rs, v)
	}
	sort.Ints(rs)
	return rs
}
