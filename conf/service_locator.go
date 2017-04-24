package conf

import (
	"context"
	"fmt"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

type InterfaceLocator interface {
	FindDevice(deviceId string) Device
	FindBrowser(module string) *node.Browser
}

func FindBrowsers(sl InterfaceLocator, module string, onNew ModuleChangeListener, onRemove ModuleChangeListener) (node.NotifyCloser, error) {
	proxy := sl.FindBrowser("proxy")
	if proxy == nil {
		return nil, c2.NewErrC("No proxy module found", 404)
	}
	s := proxy.Root().Find(fmt.Sprintf("module=%s", module))
	if s.LastErr != nil {
		return nil, s.LastErr
	}
	// Find all existing and call onNew for each one
	deviceHandles := make([]struct{ Id string }, 0)
	if err := s.Find("device?fields=id").InsertInto(node.ReflectNode(deviceHandles)).LastErr; err != nil {
		return nil, err
	}
	for _, hnd := range deviceHandles {
		device := sl.FindDevice(hnd.Id)
		if device == nil {
			panic("No device found " + hnd.Id)
		}
		onNew(hnd.Id, device, module)
	}
	return s.Find("update").Notifications(func(c context.Context, msg node.Selection) {
		deviceIdVal, err := msg.GetValue("deviceId")
		if err != nil {
			panic(err)
		}
		device := sl.FindDevice(deviceIdVal.Str)
		if device == nil {
			panic("No device found " + deviceIdVal.Str)
		}
		changeVal, err := msg.GetValue("change")
		if err != nil {
			panic(err)
		}
		if changeVal.Str == "new" {
			onNew(deviceIdVal.Str, device, module)
		} else {
			onRemove(deviceIdVal.Str, device, module)
		}
	})
}

type DeviceChangeListener func(device Device)

type DeviceRegistry interface {
	OnDeviceNew(l DeviceChangeListener)
	OnDeviceRemove(l DeviceChangeListener)
}

type ModuleChangeListener func(deviceId string, device Device, module string)

type ModuleRegistry interface {
	OnModuleNew(l ModuleChangeListener)
	OnModuleRemove(l ModuleChangeListener)
}

func AllNotifications(reg DeviceRegistry, module string, path string, stream node.NotifyStream, errs chan<- error) node.NotifyCloser {
	subs := make(map[Device]node.NotifyCloser)
	reg.OnDeviceNew(func(d Device) {
		b, err := d.Browser(module)
		if err != nil {
			errs <- err
		}
		if b == nil {
			return
		}
		s := b.Root().Find(path)
		if s.LastErr != nil {
			errs <- s.LastErr
		} else {
			sub, err := s.Notifications(stream)
			if err != nil {
				errs <- err
			} else {
				subs[d] = sub
			}
		}
	})

	reg.OnDeviceRemove(func(d Device) {
		if sub, found := subs[d]; found {
			sub()
		}
		delete(subs, d)
	})

	return func() error {
		for _, sub := range subs {
			sub()
		}
		return nil
	}
}
