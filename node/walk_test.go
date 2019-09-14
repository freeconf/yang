package node_test

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/parser"
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
	ypath := &meta.FileStreamSource{Root: "../parser/testdata"}
	m := parser.RequireModule(ypath, "rtstone")
	rdr := nodes.ReadJSON(config)
	sel := node.NewBrowser(m, rdr).Root()
	if actual, err := nodes.WriteJSON(sel); err != nil {
		t.Error(err)
	} else {
		t.Log(actual)
	}
}
