package conf

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

type DeviceManagerClient struct {
	client  Client
	browser *node.Browser
}

func NewDeviceManagerClient(d Device, client Client) *DeviceManagerClient {
	b, err := d.Browser("device-manager")
	if err != nil {
		panic(err)
	}
	return &DeviceManagerClient{
		client:  client,
		browser: b,
	}
}

type NotifySubscription node.NotifyCloser

func (self NotifySubscription) Close() error {
	return node.NotifyCloser(self)()
}

func (self *DeviceManagerClient) Device(id string) (Device, error) {
	sel := self.browser.Root().Find("device=" + id)
	if sel.LastErr != nil {
		return nil, sel.LastErr
	}
	return self.device(sel)
}

func (self *DeviceManagerClient) device(sel node.Selection) (Device, error) {
	address, err := sel.GetValue("address")
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, c2.NewErr("no address found")
	}
	return self.client.NewDevice(address.Str)
}

func (self *DeviceManagerClient) OnUpdate(l DeviceChangeListener) c2.Subscription {
	return self.onUpdate("update", l)
}

func (self *DeviceManagerClient) OnModuleUpdate(module string, l DeviceChangeListener) c2.Subscription {
	return self.onUpdate("update?filter=module/name%3d'"+module+"'", l)
}

func (self *DeviceManagerClient) onUpdate(path string, l DeviceChangeListener) c2.Subscription {
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
