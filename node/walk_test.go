package node_test

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func LoadSampleModule(t *testing.T) *meta.Module {
	m, err := yang.LoadModuleCustomImport(yang.TestDataRomancingTheStone, nil)
	if err != nil {
		t.Error(err.Error())
	}
	return m
}

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
	m := LoadSampleModule(t)
	rdr := nodes.ReadJSON(config)
	sel := node.NewBrowser(m, rdr).Root()
	if actual, err := nodes.WriteJSON(sel); err != nil {
		t.Error(err)
	} else {
		t.Log(actual)
	}
}

func TestWalkYang(t *testing.T) {
	module := LoadSampleModule(t)
	sel := nodes.Schema(module, true).Root()
	if actual, err := nodes.WriteJSON(sel); err != nil {
		t.Error(err)
	} else {
		t.Log(actual)
	}
}
