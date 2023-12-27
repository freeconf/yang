package node_test

import (
	"testing"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"

	"github.com/freeconf/yang/val"
)

func TestWalkPathTest(t *testing.T) {
	mstr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	container a {
        container b {
            leaf c {   
                type string;
            }
        }
	}
    list d {
        key "e";
        leaf e {
            type string;
        }
    }
}`

	expected := []string{
		"a",
		"a/b",
		"a/b/c",
		"d",
		"d=x/e",
	}
	ndx := 0
	checkPath := func(r *node.Request) {
		actual := r.Path.StringNoModule()
		if ndx == len(expected) {
			t.Errorf("Extra path %s", actual)
			return
		}
		if actual != expected[ndx] {
			t.Errorf("Expected: %s\n  Actual: %s", expected[ndx], actual)
		}
		ndx++
	}
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	b := node.NewBrowser(m, &nodeutil.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			checkPath(&r.Request)
			return r.Selection.Node, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			checkPath(&r.Request)
			return nil
		},
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var key []val.Value
			var n node.Node
			if r.First {
				var err error
				if key, err = node.NewValues(r.Meta.KeyMeta(), "x"); err != nil {
					return nil, nil, err
				}
				n = r.Selection.Node
			}
			return n, key, nil
		},
	})
	if err := b.Root().InsertInto(nodeutil.Null()); err != nil {
		t.Error(err)
	}
	if ndx != len(expected) {
		t.Errorf("Not all children selected Actual, %d, expected, %d", ndx, len(expected))
	}
}
