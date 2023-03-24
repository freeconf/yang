package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
)

func TestFieldInContainerWithSameName(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `
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
	if err != nil {
		t.Fatal(err)
	}
	n := nodeutil.ReadJSON(`
{
	"a" : {
		"a": "A",
		"b": "B"
 	}
}`)
	b := node.NewBrowser(m, n)
	actual, err := nodeutil.WriteJSON(b.Root().Constrain("fields=a"))
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"a":{"a":"A","b":"B"}}`
	fc.AssertEqual(t, expected, actual)
}

func TestFieldsMatcherOnList(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `
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
	if err != nil {
		t.Fatal(err)
	}
	n := nodeutil.ReadJSON(`
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
	actual, err := nodeutil.WriteJSON(b.Root().Find("a?fields=id"))
	if err != nil {
		t.Error(err)
	} else {
		fc.AssertEqual(t, `{"a":[{"id":"1"},{"id":"2"}]}`, actual)
	}
}
