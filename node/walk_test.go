package node

import (
	"bytes"
	"github.com/c2g/meta"
	"github.com/c2g/meta/yang"
	"strings"
	"testing"
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
	rdr := NewJsonReader(strings.NewReader(config)).Node()
	var actualBuff bytes.Buffer
	wtr := NewJsonWriter(&actualBuff).Node()
	if err := NewBrowser2(m, rdr).Root().Selector().UpsertInto(wtr).LastErr; err != nil {
		t.Error(err)
	}
	t.Log(string(actualBuff.Bytes()))
}

func TestWalkYang(t *testing.T) {
	var err error
	module := LoadSampleModule(t)
	var actualBuff bytes.Buffer
	wtr := NewJsonWriter(&actualBuff).Node()
	if err = SelectModule(module, true).Root().Selector().UpsertInto(wtr).LastErr; err != nil {
		t.Error(err)
	} else {
		t.Log(string(actualBuff.Bytes()))
	}
}
