package node_test

import (
	"testing"

	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
)

func TestFieldInContainerWithSameName(t *testing.T) {
	m := parser.RequireModuleFromString(nil, `
module x {
	container a {
		leaf a {
			type string;
		}
		leaf b {
			type string;
		}
	}
}
	`)
	n := nodes.ReadJSON(`
{
	"a" : {
		"a": "A",
		"b": "B"
 	}
}`)
	b := node.NewBrowser(m, n)
	actual, err := nodes.WriteJSON(b.Root().Constrain("fields=a"))
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"a":{"a":"A","b":"B"}}`
	c2.AssertEqual(t, expected, actual)
}

func TestFieldsMatcherOnList(t *testing.T) {
	m := parser.RequireModuleFromString(nil, `
module x {
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
