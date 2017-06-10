package device

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_MapClient(t *testing.T) {
	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "."},
		&meta.FileStreamSource{Root: "../yang"},
	)
	d := New(ypath)
	d.Add("test", &node.MyNode{})
	dm := NewMap()
	dm.Add("dev0", d)
	dmMod := yang.RequireModule(ypath, "device-manager")
	local := localDm{dm: dm}
	dmNode := MapNode(dm, local, local)
	dmClient := &MapClient{
		client:  local,
		browser: node.NewBrowser(dmMod, dmNode),
	}
	var gotUpdate bool
	dmClient.OnModuleUpdate("test", func(d Device, id string, c Change) {
		gotUpdate = true
	})
	if !gotUpdate {
		t.Error("never got test message")
	}
}

type localDm struct {
	dm *Map
}

func (self localDm) DeviceAddress(id string, d Device) string {
	return id
}

func (self localDm) NewDevice(address string) (Device, error) {
	return self.dm.Device(address)
}
