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

	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/log"
)

var (
	ErrorInvalidListener = errors.New("invalid listener")
)

type Traffic struct {
	locker sync.Mutex
	data   eosc.IUntyped
}

func NewTraffic() *Traffic {
	return &Traffic{
		data:   eosc.NewUntyped(),
		locker: sync.Mutex{},
	}
}

func (t *Traffic) ListenTcp(ip string, port int) (*net.TCPListener, error) {
	//if strings.HasPrefix(addr, "0.0.0.0:") {
	//	addr = strings.TrimPrefix(addr, "0.0.0.0")
	//}

	tcpAddr := ResolveTCPAddr(ip, port)
	t.locker.Lock()
	defer t.locker.Unlock()

	name := fmt.Sprintf("%s://%s", tcpAddr.Network(), tcpAddr.String())
	log.Debug("traffic listen:", name)
	if o, has := t.data.Get(name); has {
		listener, ok := o.(*net.TCPListener)
		if !ok {
			return nil, ErrorInvalidListener
		}

		return listener, nil
	}

	return nil, nil
}

type ITraffic interface {
	ListenTcp(ip string, port int) (*net.TCPListener, error)
}

func (t *Traffic) Read(r io.Reader) {
	t.locker.Lock()
	defer t.locker.Unlock()

	listeners, err := Reader(r)
	if err != nil {
		log.Warn("read listeners:", err)
		return
	}
	for _, ln := range listeners {
		t.add(ln)
	}

}

func (t *Traffic) add(ln *net.TCPListener) {
	tcpAddr := ln.Addr()
	name := fmt.Sprintf("%s://%s", tcpAddr.Network(), tcpAddr.String())
	log.Info("traffic add:", name)
	t.data.Set(name, ln)
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
