package conf

import (
	"testing"

	"bytes"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func Test_YangLibNode(t *testing.T) {
	ypath := &meta.FileStreamSource{Root: "../yang"}
	d := NewDevice(ypath)
	if err := d.Add("ietf-yang-library", LocalDeviceYangLibNode(d)); err != nil {
		t.Error(err)
	}
	var actual bytes.Buffer
	b, err := d.Browser("ietf-yang-library")
	if err != nil {
		t.Error(err)
	} else if b == nil {
		t.Error("no browser")
	} else if err = b.Root().InsertInto(node.NewJsonWriter(&actual).Node()).LastErr; err != nil {
		t.Error(err)
	}
	t.Log(actual.String())
}
