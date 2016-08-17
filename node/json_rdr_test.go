package node

import (
	"strings"
	"testing"

	"github.com/dhubler/c2g/meta/yang"
)

func TestJsonWalk(t *testing.T) {
	moduleStr := `
module json-test {
	prefix "t";
	namespace "t";
	revision 0;
	list hobbies {
		key "name";
	    leaf name {
	      type string;
	    }
		container favorite {
		    leaf common-name {
		      type string;
		    }
		    leaf location {
				type string;
		    }
		}
	}
}
	`
	module, err := yang.LoadModuleCustomImport(moduleStr, nil)
	if err != nil {
		t.Fatal(err)
	}
	json := `{"hobbies":[
{"name":"birding", "favorite": {"common-name" : "towhee", "extra":"double-mint", "location":"out back"}},
{"name":"hockey", "favorite": {"common-name" : "bruins", "location" : "Boston"}}
]}`
	tests := []string{
		"hobbies",
		"hobbies=birding",
		"hobbies=birding/favorite",
	}
	for _, test := range tests {
		rdr := NewJsonReader(strings.NewReader(json))
		sel := NewBrowser(module, rdr.Node).Root().Selector()
		found := sel.Find(test)
		if found.LastErr != nil {
			t.Error("failed to transmit json", err)
		} else if found.Selection == nil {
			t.Error(test, "- Target not found, state nil")
		} else {
			actual := found.Selection.Path().String()
			if actual != "json-test/"+test {
				t.Error("json-test/"+test, "!=", actual)
			}
		}
	}
}
