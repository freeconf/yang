package device

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

type MapClient struct {
	client  Client
	browser *node.Browser
}

func NewMapClient(d Device, client Client) *MapClient {
	b, err := d.Browser("device-manager")
	if err != nil {
		panic(err)
	}
	return &MapClient{
		client:  client,
		browser: b,
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
	return self.device(sel)
}

func (self *MapClient) device(sel node.Selection) (Device, error) {
	address, err := sel.GetValue("address")
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, c2.NewErr("no address found")
	}
	return self.client.NewDevice(address.Str)
}

func (self *MapClient) OnUpdate(l ChangeListener) c2.Subscription {
	return self.onUpdate("update", l)
}

func (self *MapClient) OnModuleUpdate(module string, l ChangeListener) c2.Subscription {
	return self.onUpdate("update?filter=module/name%3d'"+module+"'", l)
}

func (self *MapClient) onUpdate(path string, l ChangeListener) c2.Subscription {
	closer, err := self.browser.Root().Find(path).Notifications(func(msg node.Selection) {
		id, err := msg.GetValue("id")
		if err != nil {
			c2.Err.Print(err)
			return
		}
		d, err := self.device(msg)
		if err != nil {
			c2.Err.Print(err)
			return
		}
		l(d, id.Str, Added)
	})
	if err != nil {
		panic(err)
	}
	return NotifySubscription(closer)
}
