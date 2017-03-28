package conf

import (
	"container/list"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

// handles:
//   register and unregister
//   download yang and cache it
//   delegates request to remote nodes as if they were local
//   works w/store so store can persist data
type Proxy struct {
	mounts    map[string]*Mount
	factory   ProtocolHandler
	serve     DeviceServer
	yangPath  meta.StreamSource
	listeners *list.List
}

type Mount struct {
	Device        Device
	DeviceId      string
	RemoteAddress string
	RemotePort    string
	Mount         string
	Data          string
	Schema        string
	Stream        string
}

type ProtocolHandler func(yangPath meta.StreamSource, address string, port string) (Device, error)

type DeviceServer func(d Device, path string) error

func NewProxy(yangPath meta.StreamSource, proto ProtocolHandler, server DeviceServer) *Proxy {
	return &Proxy{
		factory:   proto,
		serve:     server,
		yangPath:  yangPath,
		mounts:    make(map[string]*Mount),
		listeners: list.New(),
	}
}

type ProxyNewMount func(id string, m *Mount)

func (self *Proxy) OnUpdate(l ProxyNewMount) c2.Subscription {
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *Proxy) updateListeners(id string, m *Mount) {
	p := self.listeners.Front()
	for p != nil {
		l := p.Value.(ProxyNewMount)
		l(id, m)
		p = p.Next()
	}
}

func (self *Proxy) GetMount(deviceId string) *Mount {
	return self.mounts[deviceId]
}

func (self *Proxy) Mount(id string, address string, port string) error {
	d, err := self.factory(self.yangPath, address, port)
	if err != nil {
		return err
	}
	path := "/dev=" + id + "/"
	mount := &Mount{
		Device:        d,
		DeviceId:      id,
		Data:          path + "data/",
		Stream:        path + "stream/",
		Schema:        path + "schema/",
		RemoteAddress: address,
		RemotePort:    port,
	}
	if err := self.serve(d, path); err != nil {
		return err
	}
	self.mounts[id] = mount
	self.updateListeners(id, mount)
	return nil
}
