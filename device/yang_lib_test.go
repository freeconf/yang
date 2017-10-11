package device_test

import (
	"flag"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/testdata"
)

var update = flag.Bool("update", false, "update golden test files")

func TestYangLibNode(t *testing.T) {
	d, _ := testdata.BirdDevice(`{"bird":[{
		"name" : "robin"
	},{
		"name" : "blue jay"
	}]}`)
	moduleNameAsAddress := func(m *meta.Module) string {
		return m.GetIdent()
	}
	if err := d.Add("ietf-yang-library", device.LocalDeviceYangLibNode(moduleNameAsAddress, d)); err != nil {
		t.Error(err)
	}
	b, err := d.Browser("ietf-yang-library")
	if err != nil {
		t.Error(err)
		return
	}
	if b == nil {
		t.Error("no browser")
		return
	}
	actual, err := nodes.WritePrettyJSON(b.Root())
	if err != nil {
		t.Error(err)
	}
	c2.Gold(t, *update, []byte(actual), "yang_lib.json")
}
