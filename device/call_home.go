package device

import (
	"time"

	"github.com/c2stack/c2g/node"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
type CallHome struct {
	options       CallHomeOptions
	client        Client
	device        Device
	remote        *node.Browser
	registerTimer *time.Ticker
	Registered    bool
	LastErr       string
}

type CallHomeOptions struct {
	DeviceId     string
	Address      string
	LocalAddress string
	LocalPort    string
	RateMs       int
}

func NewCallHome(client Client) *CallHome {
	return &CallHome{
		client: client,
	}
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
	d, err := self.client.NewDevice(self.options.Address)
	if err != nil {
		return err
	}
	self.device = d
	self.remote, err = d.Browser("device-manager")
	if err != nil {
		return err
	}
	self.Register()
	return nil
}

func (self *CallHome) Device() Device {
	return self.device
}

func (self *CallHome) Register() error {
	r := &struct {
		Id      string
		Port    string
		Address string
	}{
		Id:      self.options.DeviceId,
		Port:    self.options.LocalPort,
		Address: self.options.LocalAddress,
	}
	err := self.remote.Root().Find("register").Action(node.ReflectNode(r)).LastErr
	if err != nil {
		self.Registered = true
	}
	return err
}
