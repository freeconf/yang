package node

import (
	"bytes"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"strings"
	"testing"
)

func TestDiff(t *testing.T) {
	moduleStr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	container movie {
	    leaf name {
	      type string;
	    }
		container character {
		    leaf name {
		      type string;
		    }
		}
	}
	container car {
		leaf name {
			type string;
		}
		leaf year {
			type int32;
		}
	}
	container videoGame {
		leaf name {
			type string;
		}
	}
}
	`
	var err error
	var m *meta.Module
	if m, err = yang.LoadModuleCustomImport(moduleStr, nil); err != nil {
		t.Fatal(err)
	}

	// new
	a := `{
		"movie" : {
			"mame" : "StarWars",
			"character" : {
				"name" : "Hans Solo"
			}
		},
		"car" : {
			"name" : "Malibu"
		}
	}
	`
	aData := NewJsonReader(strings.NewReader(a)).Node()

	// old
	b := `{
		"movie" : {
			"mame" : "StarWars",
			"character" : {
				"name" : "Princess Laya"
			}
		},
		"videoGame" : {
			"name" : "GTA V"
		}
	}`
	bData := NewJsonReader(strings.NewReader(b)).Node()
	var out bytes.Buffer
	if err = NewBrowser(m, NewJsonWriter(&out).Node()).Root().InsertFrom(Diff(bData, aData)).LastErr; err != nil {
		t.Error(err)
	}
	actual := out.String()
	expected := `{"movie":{"character":{"name":"Princess Laya"}},"videoGame":{"name":"GTA V"}}`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
