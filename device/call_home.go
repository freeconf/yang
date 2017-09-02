package device

import (
	"container/list"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/nodes"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
type CallHome struct {
	options       CallHomeOptions
	proto         ProtocolHandler
	registerTimer *time.Ticker
	Registered    bool
	mapDevice     Device
	LastErr       string
	listeners     *list.List
}

type CallHomeOptions struct {
	DeviceId     string
	Address      string
	Endpoint     string
	LocalAddress string
	RetryRateMs  int
}

func NewCallHome(proto ProtocolHandler) *CallHome {
	return &CallHome{
		proto:     proto,
		listeners: list.New(),
	}
}

type RegisterUpdate int

const (
	Register RegisterUpdate = iota
	Unregister
)

type RegisterListener func(d Device, update RegisterUpdate)

func (self *CallHome) OnRegister(l RegisterListener) c2.Subscription {
	if self.Registered {
		l(self.mapDevice, Register)
	}
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *CallHome) Options() CallHomeOptions {
	return self.options
}

func (self *CallHome) ApplyOptions(options CallHomeOptions) error {
	if self.options == options {
		return nil
	}
	self.options = options
	self.Registered = false
	c2.Debug.Print("connecting to ", self.options.Address)
	self.Register()
	return nil
}

func (self *CallHome) updateListeners(update RegisterUpdate) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(RegisterListener)(self.mapDevice, update)
		p = p.Next()
	}
}

func (self *CallHome) Register() {
retry:
	regUrl := self.options.Address + self.options.Endpoint
	d, err := self.proto(regUrl)
	if err != nil {
		c2.Err.Printf("failed to build device with address %s. %s", regUrl, err)
	} else {
		if err = self.register(d); err != nil {
			c2.Err.Printf("failed to register %s", err)
		} else {
			return
		}
	}
	if self.options.RetryRateMs == 0 {
		panic("failed to register and no retry rate configured")
	}
	<-time.After(time.Duration(self.options.RetryRateMs) * time.Millisecond)
	goto retry
}

func (self *CallHome) register(d Device) error {
	dm, err := d.Browser("map")
	if err != nil {
		return err
	}
	r := map[string]interface{}{
		"deviceId": self.options.DeviceId,
		"address":  self.options.LocalAddress,
	}
	err = dm.Root().Find("register").Action(nodes.ReflectChild(r)).LastErr
	if err == nil {
		self.mapDevice = d
		self.updateListeners(Register)
		self.Registered = true
	}
	return err
}
