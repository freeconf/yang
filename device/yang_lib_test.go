package device

import (
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func Test_YangLibNode(t *testing.T) {
	d := bird(`{"bird":[{
		"name" : "robin"
	},{
		"name" : "blue jay"
	}]}`)
	moduleNameAsAddress := func(m *meta.Module) string {
		return m.GetIdent()
	}
	if err := d.Add("ietf-yang-library", LocalDeviceYangLibNode(moduleNameAsAddress, d)); err != nil {
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
	actual, err := node.WriteJson(b.Root())
	if err != nil {
		t.Error(err)
	}
	expected := `{"modules-state":{"module":[{"name":"ietf-yang-library","revision":"2016-06-21","schema":"ietf-yang-library","namespace":"urn:ietf:params:xml:ns:yang:ietf-yang-library"},{"name":"testdata-bird","revision":"0","schema":"testdata-bird","namespace":""}]}}`
	if err := c2.CheckEqual(expected, actual); err != nil {
		t.Error(err)
	}
}
