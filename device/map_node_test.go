package device

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_MapNode(t *testing.T) {
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
	b := node.NewBrowser(dmMod, dmNode)
	actual, err := node.WriteJson(b.Root().Find("device=dev0"))
	if err != nil {
		t.Error(err)
	}
	expected := `{"id":"dev0","address":"dev0","module":[{"name":"test","revision":"0"}]}`
	if err := c2.CheckEqual(expected, actual); err != nil {
		t.Error(err)
	}
}
