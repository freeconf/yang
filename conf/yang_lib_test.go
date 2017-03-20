package conf

import (
	"testing"

	"bytes"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Test_YangLibNode(t *testing.T) {
	m := yang.RequireModule(yang.YangPath(), "ietf-yang-library")
	reg := NewModuleRegistry("x")
	b := node.NewBrowser(m, IetfLibraryNode(reg))
	var actual bytes.Buffer
	if err := b.Root().InsertInto(node.NewJsonWriter(&actual).Node()).LastErr; err != nil {
		t.Error(err)
	}

}
