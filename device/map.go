package device

import (
	"container/list"

	"github.com/c2stack/c2g/c2"
)

type LocalMap struct {
	devices   map[string]Device
	ids       []string
	listeners *list.List
}

type Map interface {
	OnUpdate(l ChangeListener) c2.Subscription
	OnModuleUpdate(module string, l ChangeListener) c2.Subscription
	Device(deviceId string) (Device, error)
	Add(id string, d Device)
	NthDeviceId(int) string
	Len() int
}

func NewMap() *LocalMap {
	return &LocalMap{
		devices:   make(map[string]Device),
		listeners: list.New(),
	}
}

type Change int

const (
	Added = iota
	Removed
)

func (self Change) String() string {
	labels := []string{"added", "removed"}
	return labels[int(self)]
}

type ChangeListener func(d Device, id string, c Change)

func (self *LocalMap) OnUpdate(l ChangeListener) c2.Subscription {
	self.updateListener(l, Added)
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *LocalMap) updateListener(l ChangeListener, c Change) {
	for id, d := range self.devices {
		l(d, id, c)
	}
}

func (self *LocalMap) updateListeners(d Device, id string, c Change) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(ChangeListener)(d, id, c)
		p = p.Next()
	}
}

func (self *LocalMap) OnModuleUpdate(module string, l ChangeListener) c2.Subscription {
	return self.OnUpdate(func(d Device, id string, c Change) {
		if hnd := d.Modules()[module]; hnd != nil {
			l(d, id, c)
		}
	})
}

func (self *LocalMap) NthDeviceId(i int) string {
	return self.ids[i]
}

func (self *LocalMap) Device(deviceId string) (Device, error) {
	return self.devices[deviceId], nil
}

func (self *LocalMap) Add(id string, d Device) {
	self.devices[id] = d
	self.ids = append(self.ids, id)
	self.updateListeners(d, id, Added)
}

func (self *LocalMap) Len() int {
	return len(self.devices)
}

// implementation of DeviceMapServer
func (self *LocalMap) DeviceAddressFromMap(id string, d Device) string {
	return id
}
