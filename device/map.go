package device

import (
	"container/list"

	"github.com/c2stack/c2g/c2"
)

type Map struct {
	devices   map[string]Device
	listeners *list.List
}

func NewMap() *Map {
	dm := &Map{
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

type ChangeListener func(d Device, id string, c Change)

func (self *Map) OnUpdate(l ChangeListener) c2.Subscription {
	self.updateListener(l, Added)
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *Map) updateListener(l ChangeListener, c Change) {
	for id, d := range self.devices {
		l(d, id, c)
	}
}

func (self *Map) updateListeners(d Device, id string, c Change) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(ChangeListener)(d, id, c)
		p = p.Next()
	}
}

func (self *Map) OnModuleUpdate(module string, l ChangeListener) c2.Subscription {
	return self.OnUpdate(func(d Device, id string, c Change) {
		if hnd := d.Modules()[module]; hnd != nil {
			l(d, id, c)
		}
	})
}

func (self *Map) Device(deviceId string) (Device, error) {
	return self.devices[deviceId], nil
}

func (self *Map) Add(id string, d Device) {
	self.devices[id] = d
	self.updateListeners(d, id, Added)
}
