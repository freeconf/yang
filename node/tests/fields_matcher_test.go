package tests

import (
	"bytes"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func TestFieldsMatcherOnList(t *testing.T) {
	m, err := yang.LoadModuleFromString(nil, `
module x { namespace ""; prefix ""; revision 0;
  list a {
    key "id";
    leaf id {
      type string;
    }
    leaf b {
      type string;
    }
  }
}
	`)
	if err != nil {
		t.Fatal(err)
	}
	n := nodes.ReadJson(`
{
	"a" : [{
	  "id" : "1",
	  "b" : "B1"
	},{
	  "id" : "2",
	  "b" : "B2"
	}]
}`)

	b := node.NewBrowser(m, n)
	var buf bytes.Buffer
	out := nodes.NewJsonWriter(&buf)
	if err := b.Root().Find("a?fields=id").InsertInto(out.Node()).LastErr; err != nil {
		t.Error(err)
	}
	if notEq := c2.CheckEqual(`{"a":[{"id":"1"},{"id":"2"}]}`, buf.String()); notEq != nil {
		t.Error(notEq)
	}
}
