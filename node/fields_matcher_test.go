package node

import (
	"testing"
	"github.com/c2stack/c2g/meta/yang"
	"bytes"
	"github.com/c2stack/c2g/c2"
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
	n := ReadJson(`
{
	"a" : [{
	  "id" : "1",
	  "b" : "B1"
	},{
	  "id" : "2",
	  "b" : "B2"
	}]
}`)

	b := NewBrowser2(m, n)
	var buf bytes.Buffer
	out := NewJsonWriter(&buf)
	if err := b.Root().Find("a?fields=id").InsertInto(out.Node()).LastErr; err != nil {
		t.Error(err)
	}
	if notEq := c2.CheckEqual(`{"a":[{"id":"1"},{"id":"2"}]}`, buf.String()); notEq != nil {
		t.Error(notEq)
	}
}
