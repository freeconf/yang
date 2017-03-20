package conf

import (
	"time"

	"context"

	"github.com/c2stack/c2g/node"
)

// Implements RFC Draft in spirit-only
//   https://tools.ietf.org/html/draft-ietf-netconf-call-home-17
//
// Draft calls for server-initiated registration and this implementation is client-initiated
// which may or may-not be part of the final draft.  Client-initiated registration at first
// glance appears to be more useful, but this may prove to be a wrong assumption on my part.
//
type CallHome struct {
	DeviceId      string
	Address       string
	DeviceAddress string
	Proxy         *node.Browser
	RateMs        int
	registerTimer *time.Ticker
	Registered    bool
	LastErr       string
}

type registration struct {
	DeviceId      string
	DeviceAddress string
}

func (self *CallHome) Register(c context.Context) {
	r := &registration{DeviceId: self.DeviceId, DeviceAddress: self.DeviceAddress}
	self.Proxy.Root().Find("register").Action(c, node.ReflectNode(&r))
}

func (self *CallHome) Unregister(c context.Context) {
	r := &registration{DeviceId: self.DeviceId, DeviceAddress: self.DeviceAddress}
	self.Proxy.Root().Find("unregister").Action(c, node.ReflectNode(&r))
}
