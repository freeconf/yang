package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/testdata"
)

func TestWhere(t *testing.T) {
	b, _ := testdata.BirdBrowser(`
{
	"bird" : [{
		"name" : "blue jay",
		"wingspan": 99
	},{
		"name" : "sparrow"
	},{
		"name" : "robin",
		"wingspan": 80
	},{
		"name" : "heron"
	},{
		"name" : "pee wee",
		"species" : {
			"name" : "fly catcher"
		}
	}]
}
`)
	actual, err := nodeutil.WriteJSON(b.Root().Find("bird?where=name%3d'robin'"))
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[{"name":"robin","wingspan":80}]}`, actual)
}

func TestWhereEnum(t *testing.T) {
	mstr := `module m {
		list bird {
			key name;
			leaf name {
				type enumeration {
					enum robin;
					enum sparrow;
					enum heron;
				}
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	fc.RequireEqual(t, nil, err)
	birds := []map[string]any{
		{"name": "robin"},
	}
	b := node.NewBrowser(m, nodeutil.ReflectChild(map[string]any{"bird": birds}))
	s := b.Root().Find("bird?where=name%3d'robin'")
	fc.AssertEqual(t, nil, s.LastErr)
	actual, err := nodeutil.WriteJSON(s)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[{"name":"robin"}]}`, actual)

	s = b.Root().Find("bird?where=name%3d'sparrow'")
	fc.AssertEqual(t, nil, s.LastErr)
	actual, err = nodeutil.WriteJSON(s)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[]}`, actual)
}
