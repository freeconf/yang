package conf

import (
	"time"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
type CallHome struct {
	options       CallHomeOptions
	proto         ProtocolHandler
	ypath         meta.StreamSource
	remote        *node.Browser
	registerTimer *time.Ticker
	Registered    bool
	LastErr       string
}

type CallHomeOptions struct {
	DeviceId            string
	LocalAddress        string
	RegistrationAddress string
	RegistrationPort    string
	LocalPort           string
	RateMs              int
}

func NewCallHome(ypath meta.StreamSource, proto ProtocolHandler) *CallHome {
	return &CallHome{
		ypath: ypath,
		proto: proto,
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
	d, err := self.proto(self.ypath, self.options.RegistrationAddress, self.options.RegistrationPort)
	if err != nil {
		return err
	}
	self.remote, err = d.Browser("call-home-register")
	if err != nil {
		return err
	}
	self.Register()
	return nil
}

type registration struct {
	Id      string
	Port    string
	Address string
}

func (self *CallHome) Register() error {
	r := &registration{Id: self.options.DeviceId, Port: self.options.LocalPort, Address: self.options.LocalAddress}
	err := self.remote.Root().Find("register").Action(node.ReflectNode(r)).LastErr
	if err != nil {
		self.Registered = true
	}
	return err
}
