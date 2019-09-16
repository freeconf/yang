package device

import (
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
)

func TestMapNode(t *testing.T) {
	ypath := meta.MultipleSources(
		&meta.FileStreamSource{Root: "./testdata"},
		&meta.FileStreamSource{Root: "../yang"},
	)
	d := New(ypath)
	d.Add("test", &nodes.Basic{})
	dm := NewMap()
	dm.Add("dev0", d)
	dmMod := parser.RequireModule(ypath, "fc-map")
	dmNode := MapNode(dm)
	b := node.NewBrowser(dmMod, dmNode)
	actual, err := nodes.WriteJSON(b.Root().Find("device=dev0"))
	if err != nil {
		t.Error(err)
	}
	expected := `{"deviceId":"dev0","module":[{"name":"test","revision":"0"}]}`
	c2.AssertEqual(t, expected, actual)
}
