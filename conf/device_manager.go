package conf

import (
	"container/list"

	"github.com/c2stack/c2g/c2"
)

type DeviceManager struct {
	devices   map[string]Device
	listeners *list.List
}

func NewDeviceManager() *DeviceManager {
	dm := &DeviceManager{
		devices:   make(map[string]Device),
		listeners: list.New(),
	}
	return dm
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

type DeviceChangeListener func(d Device, id string, c Change)

func (self *DeviceManager) OnUpdate(l DeviceChangeListener) c2.Subscription {
	self.updateListener(l, Added)
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *DeviceManager) updateListener(l DeviceChangeListener, c Change) {
	for id, d := range self.devices {
		l(d, id, c)
	}
}

func (self *DeviceManager) updateListeners(d Device, id string, c Change) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(DeviceChangeListener)(d, id, c)
		p = p.Next()
	}
}

func (self *DeviceManager) OnModuleUpdate(module string, l DeviceChangeListener) c2.Subscription {
	return self.OnUpdate(func(d Device, id string, c Change) {
		if hnd := d.Modules()[module]; hnd != nil {
			l(d, id, c)
		}
	})
}

func (self *DeviceManager) Device(deviceId string) (Device, error) {
	return self.devices[deviceId], nil
}

func (self *DeviceManager) Add(id string, d Device) {
	self.devices[id] = d
	self.updateListeners(d, id, Added)
}
