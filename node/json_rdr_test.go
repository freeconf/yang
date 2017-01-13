package node

import (
	"github.com/c2stack/c2g/meta/yang"
	"strings"
	"testing"
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
		sel := NewBrowserSource(module, rdr.Node).Root()
		found := sel.Find(test)
		if found.LastErr != nil {
			t.Error("failed to transmit json", found.LastErr)
		} else if found.IsNil() {
			t.Error(test, "- Target not found, state nil")
		} else {
			actual := found.Path.String()
			if actual != "json-test/"+test {
				t.Error("json-test/"+test, "!=", actual)
			}
		}
	}
}

func TestNumberParse(t *testing.T) {
	moduleStr := `
module json-test {
	prefix "t";
	namespace "t";
	revision 0;
	container data {
		leaf id {
			type int32;
		}
		leaf idstr {
			type int32;
		}
		leaf idstrwrong {
			type int32;
		}
		leaf-list readings{
			type decimal64;
		}
	}
}
	`
	module, err := yang.LoadModuleCustomImport(moduleStr, nil)
	if err != nil {
		t.Fatal(err)
	}
	json := `{ "data": {
			"id": 4,
			"idstr": "4",
			"idstrwrong": "4s",
			"readings": [
				"3.555454",
				"45.04545",
				324545.04
			]
		}
	}`

	//test get id
	rdr := NewJsonReader(strings.NewReader(json))
	sel := NewBrowserSource(module, rdr.Node).Root().Find("data")
	found, err := sel.Get("id")
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/id - Target not found, state nil")
	} else {
		if 4 != found.(int) {
			t.Error(found.(int), "!=", 4)
		}
	}

	//test get idstr
	rdr = NewJsonReader(strings.NewReader(json))
	sel = NewBrowserSource(module, rdr.Node).Root().Find("data")
	found, err = sel.Get("idstr")
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/idstr - Target not found, state nil")
	} else {
		if 4 != found.(int) {
			t.Error(found.(int), "!=", 4)
		}
	}

	//test idstrwrong fail
	rdr = NewJsonReader(strings.NewReader(json))
	sel = NewBrowserSource(module, rdr.Node).Root().Find("data")
	found, err = sel.Get("idstrwrong")
	if err == nil {
		t.Error("Failed to throw error on invalid input")
	}

	rdr = NewJsonReader(strings.NewReader(json))
	sel = NewBrowserSource(module, rdr.Node).Root().Find("data")
	found, err = sel.Get("readings")
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/readings - Target not found, state nil")
	} else {
		expected := []float64{3.555454, 45.04545, 324545.04}
		readings := found.([]float64)

		if expected[0] != readings[0] || expected[1] != readings[1] || expected[2] != readings[2] {
			t.Error(found.([]int), "!=", expected)
		}
	}
}
