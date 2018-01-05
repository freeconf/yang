package device_test

import (
	"testing"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/gateway"
	"github.com/freeconf/gconf/meta"
)

func TestCallHome(t *testing.T) {
	c2.DebugLog(true)

	registrar := gateway.NewLocalRegistrar()
	regDevice := device.New(&meta.FileStreamSource{Root: "../yang"})
	if err := regDevice.Add("registrar", gateway.RegistrarNode(registrar)); err != nil {
		t.Error(err)
	}
	caller := device.NewCallHome(func(string) (device.Device, error) {
		return regDevice, nil
	})
	options := caller.Options()
	options.DeviceId = "x"
	options.Address = "north"
	options.LocalAddress = "south"
	var gotUpdate bool
	caller.OnRegister(func(d device.Device, update device.RegisterUpdate) {
		gotUpdate = true
	})
	caller.ApplyOptions(options)
	if !gotUpdate {
		t.Error("no update recieved")
	}
	c2.AssertEqual(t, 1, registrar.RegistrationCount())
}
