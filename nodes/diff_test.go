package nodes

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
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
	a := ReadJSON(`{
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
	`)

	// old
	b := ReadJSON(`{
		"movie" : {
			"mame" : "StarWars",
			"character" : {
				"name" : "Princess Laya"
			}
		},
		"videoGame" : {
			"name" : "GTA V"
		}
	}`)

	sel := node.NewBrowser(m, Diff(b, a)).Root()
	expected := `{"movie":{"character":{"name":"Princess Laya"}},"videoGame":{"name":"GTA V"}}`
	if actual, err := WriteJSON(sel); err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
