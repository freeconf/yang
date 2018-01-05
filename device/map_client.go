package device

import (
	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/node"
)

type MapClient struct {
	proto       ProtocolHandler
	browser     *node.Browser
	baseAddress string
}

func NewMapClient(d Device, baseAddress string, proto ProtocolHandler) *MapClient {
	b, err := d.Browser("map")
	if err != nil {
		panic(err)
	}
	return &MapClient{
		proto:       proto,
		browser:     b,
		baseAddress: baseAddress,
	}
}

type NotifySubscription node.NotifyCloser

func (self NotifySubscription) Close() error {
	return node.NotifyCloser(self)()
}

func (self *MapClient) Device(id string) (Device, error) {
	sel := self.browser.Root().Find("device=" + id)
	if sel.LastErr != nil {
		return nil, sel.LastErr
	}
	return self.device(id)
}

type DeviceHnd struct {
	DeviceId string
	Address  string
}

func (self *MapClient) device(deviceId string) (Device, error) {
	address := self.baseAddress + "/device=" + deviceId
	c2.Debug.Printf("map client address %s", address)
	return self.proto(address)
}

func (self *MapClient) OnUpdate(l ChangeListener) c2.Subscription {
	return self.onUpdate("update", l)
}

func (self *MapClient) OnModuleUpdate(module string, l ChangeListener) c2.Subscription {
	return self.onUpdate("update?filter=module/name%3d'"+module+"'", l)
}

func (self *MapClient) onUpdate(path string, l ChangeListener) c2.Subscription {
	closer, err := self.browser.Root().Find(path).Notifications(func(msg node.Selection) {
		id, err := msg.GetValue("deviceId")
		if err != nil {
			c2.Err.Print(err)
			return
		}
		d, err := self.device(id.String())
		if err != nil {
			c2.Err.Print(err)
			return
		}
		l(d, id.String(), Added)
	})
	if err != nil {
		panic(err)
	}
	return NotifySubscription(closer)
}
