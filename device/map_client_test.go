package device

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func Test_MapClient(t *testing.T) {
	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "."},
		&meta.FileStreamSource{Root: "../yang"},
	)
	d := New(ypath)
	d.Add("test", &nodes.Basic{})
	dm := NewMap()
	dm.Add("dev0", d)
	dmMod := yang.RequireModule(ypath, "map")
	noProto := dm.Device
	deviceIdAsAddress := func(id string, d Device) string {
		return id
	}
	dmNode := MapNode(dm, deviceIdAsAddress, noProto)
	dmClient := &MapClient{
		proto:   noProto,
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
