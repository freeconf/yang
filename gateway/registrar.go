package gateway

import (
	"container/list"

	"github.com/freeconf/c2g/c2"
)

type Registration struct {
	DeviceId string
	Address  string
}

type Registrar interface {
	RegistrationCount() int
	LookupRegistration(deviceId string) (Registration, bool)
	RegisterDevice(deviceId string, address string)
	OnRegister(l RegisterListener) c2.Subscription
}

type RegisterListener func(Registration)

type LocalRegistrar struct {
	regs      map[string]Registration
	listeners *list.List
}

func NewLocalRegistrar() *LocalRegistrar {
	return &LocalRegistrar{
		regs:      make(map[string]Registration),
		listeners: list.New(),
	}
}

func (self *LocalRegistrar) LookupRegistration(deviceId string) (Registration, bool) {
	found, reg := self.regs[deviceId]
	return found, reg
}

func (self *LocalRegistrar) RegisterDevice(deviceId string, address string) {
	self.regs[deviceId] = Registration{Address: address, DeviceId: deviceId}
}

func (self *LocalRegistrar) updateListeners(reg Registration) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(RegisterListener)(reg)
		p.Next()
	}
}

func (self *LocalRegistrar) RegistrationCount() int {
	return len(self.regs)
}

func (self *LocalRegistrar) OnRegister(l RegisterListener) c2.Subscription {
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}
