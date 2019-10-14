package node_test

import (
	"testing"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

func TestWalkJson(t *testing.T) {
	config := `{
	"game" : {
		"base-radius" : 14,
		"teams" : [{
  		  "color" : "red",
		  "team" : {
		    "members" : ["joe","mary"]
		  }
		}]
	}
}`
	ypath := source.Dir("../parser/testdata")
	m := parser.RequireModule(ypath, "rtstone")
	rdr := nodeutil.ReadJSON(config)
	sel := node.NewBrowser(m, rdr).Root()
	if actual, err := nodeutil.WriteJSON(sel); err != nil {
		t.Error(err)
	} else {
		t.Log(actual)
	}
}
