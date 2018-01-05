package device

import (
	"container/list"
	"time"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/nodes"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
type CallHome struct {
	options        CallHomeOptions
	d              Device
	registrarProto ProtocolHandler
	registerTimer  *time.Ticker
	Registered     bool
	registrar      Device
	LastErr        string
	listeners      *list.List
}

type CallHomeOptions struct {
	DeviceId     string
	Address      string
	Endpoint     string
	LocalAddress string
	RetryRateMs  int
}

func NewCallHome(registrarProto ProtocolHandler) *CallHome {
	return &CallHome{
		registrarProto: registrarProto,
		listeners:      list.New(),
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
		l(self.registrar, Register)
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

func (self *CallHome) updateListeners(registrar Device, update RegisterUpdate) {
	self.registrar = registrar
	p := self.listeners.Front()
	for p != nil {
		p.Value.(RegisterListener)(self.registrar, update)
		p = p.Next()
	}
}

func (self *CallHome) Register() {
retry:
	regUrl := self.options.Address + self.options.Endpoint
	registrar, err := self.registrarProto(regUrl)
	if err != nil {
		c2.Err.Printf("failed to build device with address %s. %s", regUrl, err)
	} else {
		if err = self.register(registrar); err != nil {
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

func (self *CallHome) register(registrar Device) error {
	reg, err := registrar.Browser("registrar")
	if err != nil {
		return err
	}
	r := map[string]interface{}{
		"deviceId": self.options.DeviceId,
		"address":  self.options.LocalAddress,
	}
	err = reg.Root().Find("register").Action(nodes.ReflectChild(r)).LastErr
	if err == nil {
		self.updateListeners(registrar, Register)
		self.Registered = true
	}
	return err
}
