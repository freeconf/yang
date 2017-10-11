package device_test

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/gateway"
	"github.com/c2stack/c2g/meta"
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
