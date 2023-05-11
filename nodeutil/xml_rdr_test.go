package nodeutil

import (
	"testing"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

func TestXmlWalk(t *testing.T) {
	moduleStr := `
module xml-test {
	prefix "t";
	namespace "t";
	revision 0;
	container hobbies {
		list hobbie {
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
}
	`
	module, err := parser.LoadModuleFromString(nil, moduleStr)
	if err != nil {
		t.Fatal(err)
	}
	xml := `<hobbies><hobbie><name>birding</name><favorite><common-name>towhee</common-name><extra>double-mint</extra><location>out back</location></favorite></hobbie><hobbie><name>hockey</name><favorite><common-name>bruins</common-name><location>Boston</location></favorite></hobbie></hobbies></xml-test>`

	tests := []string{
		"hobbies",
		"hobbies/hobbie=birding",
		"hobbies/hobbie=birding/favorite",
	}
	for _, test := range tests {
		sel := node.NewBrowser(module, ReadXML(xml)).Root()
		found := sel.Find(test)
		if found.LastErr != nil {
			t.Error("failed to transmit json", found.LastErr)
		} else if found.IsNil() {
			t.Error(test, "- Target not found, state nil")
		} else {
			actual := found.Path.String()
			if actual != "xml-test/"+test {
				t.Error("xml-test/"+test, "!=", actual)
			}
		}
	}
}

/*
func TestJsonRdrUnion(t *testing.T) {
	mstr := `
	module x {
		revision 0;
		leaf y {
			type union {
				type int32;
				type string;
			}
		}
	}
		`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{in: `{"y":24}`, out: `{"y":24}`},
		{in: `{"y":"hi"}`, out: `{"y":"hi"}`},
	}
	for _, json := range tests {
		t.Log(json.in)
		actual, err := WriteJSON(node.NewBrowser(m, ReadJSON(json.in)).Root())
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, json.out, actual)
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
	module, err := parser.LoadModuleFromString(nil, moduleStr)
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
	sel := node.NewBrowser(module, ReadJSON(json)).Root().Find("data")
	found, err := sel.Find("id").Get()
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/id - Target not found, state nil")
	} else {
		if 4 != found.Value().(int) {
			t.Error(found.Value().(int), "!=", 4)
		}
	}

	//test get idstr
	sel = node.NewBrowser(module, ReadJSON(json)).Root().Find("data")
	found, err = sel.Find("idstr").Get()
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/idstr - Target not found, state nil")
	} else {
		if 4 != found.Value().(int) {
			t.Error(found.Value().(int), "!=", 4)
		}
	}

	//test idstrwrong fail
	sel = node.NewBrowser(module, ReadJSON(json)).Root().Find("data")
	found, err = sel.Find("idstrwrong").Get()
	if err == nil {
		t.Error("Failed to throw error on invalid input")
	}

	sel = node.NewBrowser(module, ReadJSON(json)).Root().Find("data")
	found, err = sel.Find("readings").Get()
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/readings - Target not found, state nil")
	} else {
		expected := []float64{3.555454, 45.04545, 324545.04}
		readings := found.Value().([]float64)

		if expected[0] != readings[0] || expected[1] != readings[1] || expected[2] != readings[2] {
			t.Error(found.Value().([]int), "!=", expected)
		}
	}
}

func TestJsonEmpty(t *testing.T) {
	moduleStr := `
module json-test {
	leaf x {
		type empty;
	}
}
	`
	m, err := parser.LoadModuleFromString(nil, moduleStr)
	fc.AssertEqual(t, nil, err)
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	in := `{"x":{}}`
	fc.AssertEqual(t, nil, b.Root().InsertFrom(ReadJSON(in)).LastErr)
	fc.AssertEqual(t, val.NotEmpty, actual["x"])
}

func TestReadQualifiedJsonIdentRef(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	in := `{
		"module-test:type":"module-types:derived-type",
		"module-test:type2":"local-type"
	}`
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	fc.AssertEqual(t, nil, b.Root().InsertFrom(ReadJSON(in)).LastErr)
	fc.AssertEqual(t, "derived-type", actual["type"].(val.IdentRef).Label)
	fc.AssertEqual(t, "local-type", actual["type2"].(val.IdentRef).Label)
}
*/
