package tests

import (
	"bytes"
	"strings"
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
	rdr := nodes.NewJsonReader(strings.NewReader(config)).Node()
	var actualBuff bytes.Buffer
	wtr := nodes.NewJsonWriter(&actualBuff).Node()
	if err := node.NewBrowser(m, rdr).Root().UpsertInto(wtr).LastErr; err != nil {
		t.Error(err)
	}
	t.Log(string(actualBuff.Bytes()))
}

func TestWalkYang(t *testing.T) {
	var err error
	module := LoadSampleModule(t)
	var actualBuff bytes.Buffer
	wtr := nodes.NewJsonWriter(&actualBuff).Node()
	if err = nodes.SelectModule(module, true).Root().UpsertInto(wtr).LastErr; err != nil {
		t.Error(err)
	} else {
		t.Log(string(actualBuff.Bytes()))
	}
}
