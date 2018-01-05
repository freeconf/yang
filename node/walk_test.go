package node_test

import (
	"testing"

	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/meta/yang"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
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
