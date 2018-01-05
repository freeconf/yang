package node_test

import (
	"testing"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
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
	n := nodes.ReadJSON(`
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
	actual, err := nodes.WriteJSON(b.Root().Find("a?fields=id"))
	if err != nil {
		t.Error(err)
	} else {
		c2.AssertEqual(t, `{"a":[{"id":"1"},{"id":"2"}]}`, actual)
	}
}
