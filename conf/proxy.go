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
	mounts          map[string]*Mount
	factory         ProtocolHandler
	serve           DeviceServer
	yangPath        meta.StreamSource
	listeners       *list.List
	moduleListeners *list.List
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

type DeviceSource func(deviceId string) Device

type ProtocolHandler func(yangPath meta.StreamSource, address string, port string, deviceId string) (Device, error)

type DeviceServer func(id string, d Device) error

func NewProxy(yangPath meta.StreamSource, proto ProtocolHandler, server DeviceServer) *Proxy {
	return &Proxy{
		factory:         proto,
		serve:           server,
		yangPath:        yangPath,
		mounts:          make(map[string]*Mount),
		listeners:       list.New(),
		moduleListeners: list.New(),
	}
}

type ProxyNewMount func(m *Mount)

func (self *Proxy) OnUpdate(l ProxyNewMount) c2.Subscription {
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

type ProxyNewModule func(module string, m *Mount)

func (self *Proxy) OnModuleUpdate(replay bool, l ProxyNewModule) c2.Subscription {
	if replay {
		go func() {
			for _, m := range self.mounts {
				hnds, err := m.Device.ModuleHandles()
				if err != nil {
					panic(err)
				}
				self.updateModuleListener(l, m, hnds)
			}
		}()
	}
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *Proxy) updateListeners(m *Mount) error {
	self.updateDeviceListeners(m)
	return self.updateModuleListeners(m)
}

func (self *Proxy) updateDeviceListeners(m *Mount) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(ProxyNewMount)(m)
		p = p.Next()
	}
}

func (self *Proxy) updateModuleListeners(m *Mount) error {
	if self.moduleListeners.Len() == 0 {
		return nil
	}
	hnds, err := m.Device.ModuleHandles()
	if err != nil {
		return err
	}
	p := self.listeners.Front()
	for p != nil {
		self.updateModuleListener(p.Value.(ProxyNewModule), m, hnds)
		p = p.Next()
	}
	return nil
}

func (self *Proxy) updateModuleListener(l ProxyNewModule, m *Mount, hnds map[string]*ModuleHandle) {
	for _, hnd := range hnds {
		l(hnd.Name, m)
	}
}

func (self *Proxy) GetMount(deviceId string) *Mount {
	return self.mounts[deviceId]
}

func (self *Proxy) Mount(id string, address string, port string) error {
	d, err := self.factory(self.yangPath, address, port, id)
	if err != nil {
		return err
	}
	if d == nil {
		return c2.NewErrC("no device found "+id, 404)
	}
	mount := &Mount{
		Device:   d,
		DeviceId: id,

		// northbound restconf addresses
		Data:   "data=" + id + "/",
		Stream: "stream=" + id + "/",
		Schema: "schema=" + id + "/",

		// southbound address
		RemoteAddress: address,
		RemotePort:    port,
	}
	if err := self.serve(id, d); err != nil {
		return err
	}
	self.mounts[id] = mount
	self.updateListeners(mount)
	return nil
}
