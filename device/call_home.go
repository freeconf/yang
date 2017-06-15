package device

import (
	"container/list"
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
type CallHome struct {
	options       CallHomeOptions
	proto         ProtocolHandler
	registerTimer *time.Ticker
	Registered    bool
	LastErr       string
	listeners     *list.List
}

type CallHomeOptions struct {
	DeviceId     string
	Address      string
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

func (self *CallHome) updateListeners(d Device, update RegisterUpdate) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(RegisterListener)(d, update)
		p = p.Next()
	}
}

func (self *CallHome) Register() {
retry:
	d, err := self.proto(self.options.Address)
	if err == nil {
		if err = self.register(d); err == nil {
			return
		}
	}
	if self.options.RetryRateMs == 0 {
		panic("failed to register and no retry rate configured")
	}
	c2.Err.Print("registration failed, retrying....  Err:", err)
	<-time.After(time.Duration(self.options.RetryRateMs) * time.Millisecond)
	goto retry
}

func (self *CallHome) register(d Device) error {
	dm, err := d.Browser("device-manager")
	if err != nil {
		return err
	}
	r := map[string]interface{}{
		"deviceId": self.options.DeviceId,
		"address":  self.options.LocalAddress,
	}
	err = dm.Root().Find("register").Action(node.MapNode(r)).LastErr
	if err == nil {
		self.updateListeners(d, Register)
		self.Registered = true
	}
	return err
}
