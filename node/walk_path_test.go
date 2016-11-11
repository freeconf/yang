package node

import (
	"testing"
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
	checkPath := func(r *Request) {
		slice := &PathSlice{Head: r.Base, Tail: r.Path}
		actual := slice.String()
		if ndx == len(expected) {
			t.Errorf("Extra path %s", actual)
			return
		}
		if actual != expected[ndx] {
			t.Errorf("Expected: %s\n  Actual: %s", expected[ndx], actual)
		}
		ndx++
	}
	b := NewBrowser(YangFromString(mstr), &MyNode{
		OnSelect: func(r ContainerRequest) (Node, error) {
			checkPath(&r.Request)
			return r.Selection.Node, nil
		},
		OnField: func(r FieldRequest, hnd *ValueHandle) error {
			checkPath(&r.Request)
			return nil
		},
		OnNext: func(r ListRequest) (Node, []*Value, error) {
			var key []*Value
			var n Node
			if r.First {
				key = SetValues(r.Meta.KeyMeta(), "x")
				n = r.Selection.Node
			}
			return n, key, nil
		},
	})
	if err := b.Root().InsertInto(DevNull()).LastErr; err != nil {
		t.Error(err)
	}
	if ndx != len(expected) {
		t.Errorf("Not all children selected Actual, %d, expected, %d", ndx, len(expected))
	}
}
