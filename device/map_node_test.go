package device

import (
	"testing"

	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
)

func TestMapNode(t *testing.T) {
	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "./testdata"},
		&meta.FileStreamSource{Root: "."},
	)
	d := New(ypath)
	d.Add("test", &nodes.Basic{})
	dm := NewMap()
	dm.Add("dev0", d)
	dmMod := yang.RequireModule(ypath, "map")
	dmNode := MapNode(dm)
	b := node.NewBrowser(dmMod, dmNode)
	actual, err := nodes.WriteJSON(b.Root().Find("device=dev0"))
	if err != nil {
		t.Error(err)
	}
	expected := `{"deviceId":"dev0","module":[{"name":"test","revision":"0"}]}`
	c2.AssertEqual(t, expected, actual)
}
