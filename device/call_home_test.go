package device

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

func Test_CallHome(t *testing.T) {
	c2.DebugLog(true)
	south := bird("")
	north := New(&meta.FileStreamSource{Root: "../yang"})
	moduleNameAsAddress := func(m *meta.Module) string {
		return m.GetIdent()
	}
	deviceIdAsAddress := func(id string, d Device) string {
		return id
	}
	if err := south.Add("ietf-yang-library", LocalDeviceYangLibNode(moduleNameAsAddress, south)); err != nil {
		t.Error(err)
	}
	noProto := func(addr string) (Device, error) {
		switch addr {
		case "south":
			return south, nil
		case "north":
			return north, nil
		}
		panic(addr)
	}

	dm := NewMap()
	if err := north.Add("map", MapNode(dm, deviceIdAsAddress, noProto)); err != nil {
		t.Error(err)
	}
	ch := NewCallHome(noProto)
	options := ch.Options()
	options.DeviceId = "x"
	options.Address = "north"
	options.LocalAddress = "south"
	var gotUpdate bool
	ch.OnRegister(func(d Device, update RegisterUpdate) {
		gotUpdate = true
	})
	ch.ApplyOptions(options)
	if !gotUpdate {
		t.Error("no update recieved")
	}
	if err := c2.CheckEqual(1, len(dm.devices)); err != nil {
		t.Error(err)
	}
}
