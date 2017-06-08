package conf

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_DeviceManagerClient(t *testing.T) {
	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "."},
		&meta.FileStreamSource{Root: "../yang"},
	)
	d := NewDevice(ypath)
	d.Add("test", &node.MyNode{})
	dm := NewDeviceManager()
	dm.Add("dev0", d)
	dmMod := yang.RequireModule(ypath, "device-manager")
	local := localDm{dm: dm}
	dmNode := DeviceManagerNode(dm, local, local)
	dmClient := &DeviceManagerClient{
		client:  local,
		browser: node.NewBrowser(dmMod, dmNode),
	}
	dmClient.OnModuleUpdate("test", func(d Device, id string, c Change) {
		t.Log("HERE")
	})
}

type localDm struct {
	dm *DeviceManager
}

func (self localDm) DeviceAddress(id string, d Device) string {
	return id
}

func (self localDm) NewDevice(address string) (Device, error) {
	return self.dm.Device(address)
}
