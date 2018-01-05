package device

import (
	"testing"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
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
