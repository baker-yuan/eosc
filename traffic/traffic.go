/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package traffic

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/eolinker/eosc/log"
)

var (
	ErrorInvalidListener = errors.New("invalid listener")
)
var _ ITraffic = (*Traffic)(nil)

type Traffic struct {
	locker sync.Mutex
	data   *tTrafficData
}

func (t *Traffic) Expire(ports []int) {
	t.locker.Lock()
	defer t.locker.Unlock()

	newData := newTTrafficData()

	old := t.data.clone()

	for _, p := range ports {
		addr := ResolveTCPAddr("", p)
		name := addrToName(addr)
		if o, has := old.Del(name); has {
			log.Debug("move traffic:", name)
			newData.add(o)
		}
	}
	for n, o := range old.All() {

		log.Debug("close old : ", n)
		o.shutdown()
		//if err := o.shutdown(); err != nil {
		//	log.Warn("close listener:", err, " ", o.Addr())
		//}
		log.Debug("close old done:", n)
	}

}

func NewTraffic() *Traffic {
	return &Traffic{
		data:   newTTrafficData(),
		locker: sync.Mutex{},
	}
}
func (t *Traffic) Read(in io.Reader) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	data := newTTrafficData()
	data.Read(in)
	t.data = data
	return nil
}

func (t *Traffic) ListenTcp(ip string, port int) (net.Listener, error) {
	log.Debug("traffic try ListenTcp:", ip, ":", port)
	tcpAddr := ResolveTCPAddr(ip, port)
	name := addrToName(tcpAddr)
	t.locker.Lock()
	defer t.locker.Unlock()

	log.Debug("traffic listen:", name)
	if o, has := t.data.get(name); has {

		//if !ok {
		//	log.Debug("traffic ListenTcp:", ip, ":", port, ", not listener")
		//
		//	return nil, ErrorInvalidListener
		//}
		log.Debug("traffic ListenTcp:", ip, ":", port, ", ok")

		return o, nil
	}
	log.Debug("traffic ListenTcp:", ip, ":", port, ", not has")

	return nil, nil
}

type ITraffic interface {
	ListenTcp(ip string, port int) (net.Listener, error)
	Close()
	Expire([]int)
}

func (t *Traffic) Close() {
	t.locker.Lock()
	list := t.data.list()
	t.data = newTTrafficData()
	t.locker.Unlock()
	for _, it := range list {
		err := it.Close()
		if err != nil {
			log.Info("close traffic listener:", err)
		}
	}
}

func resolve(value string) net.IP {
	ip := net.ParseIP(value)
	if ip == nil {
		return net.IPv6zero
	}
	if ip.Equal(net.IPv4zero) {
		return net.IPv6zero
	}
	return ip
}

func ResolveTCPAddr(ip string, port int) *net.TCPAddr {

	return &net.TCPAddr{
		IP:   resolve(ip),
		Port: port,
		Zone: "",
	}
}
func toName(ln net.Listener) string {
	addr := ln.Addr()
	return addrToName(addr)
}

func addrToName(addr net.Addr) string {
	return fmt.Sprintf("%s://%s", addr.Network(), addr.String())

}
