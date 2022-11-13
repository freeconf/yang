package nodeutil

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
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
	if m, err = parser.LoadModuleFromString(nil, moduleStr); err != nil {
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
	expected := `{"m:movie":{"character":{"name":"Princess Laya"}},"m:videoGame":{"name":"GTA V"}}`
	if actual, err := WriteJSON(sel); err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
