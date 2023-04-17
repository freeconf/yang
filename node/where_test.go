package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/nodeutil"
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
