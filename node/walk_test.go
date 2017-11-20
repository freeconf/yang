package node_test

import (
	"testing"

	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/meta/yang"
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
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
	ypath := &meta.FileStreamSource{Root: "../meta/yang/testdata"}
	m := yang.RequireModule(ypath, "rtstone")
	rdr := nodes.ReadJSON(config)
	sel := node.NewBrowser(m, rdr).Root()
	if actual, err := nodes.WriteJSON(sel); err != nil {
		t.Error(err)
	} else {
		t.Log(actual)
	}
}
