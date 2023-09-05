package nodeutil

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
	"github.com/freeconf/yang/val"
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
	xml := `<hobbies><hobbie><name>birding</name><favorite><common-name>towhee</common-name><extra>double-mint</extra><location>out back</location></favorite></hobbie><hobbie><name>hockey</name><favorite><common-name>bruins</common-name><location>Boston</location></favorite></hobbie></hobbies>`
	tests := []string{
		"hobbies",
		"hobbies/hobbie=birding",
		"hobbies/hobbie=birding/favorite",
	}
	for _, test := range tests {
		sel := node.NewBrowser(module, ReadXML(xml)).Root()
		found, err := sel.Find(test)
		fc.RequireEqual(t, nil, err, "failed to transmit xml")
		fc.RequireEqual(t, true, found != nil, "target not found")
		fc.AssertEqual(t, "xml-test/"+test, found.Path.String())
	}
}

func TestXmlRdrUnion(t *testing.T) {
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
		{in: `{"y":24}`, out: `<y xmlns="x">24</y>`},
		{in: `{"y":"hi"}`, out: `<y xmlns="x">hi</y>`},
	}
	for _, xml := range tests {
		t.Log(xml.in)
		var sel *node.Selection
		var err error
		sel, err = node.NewBrowser(m, ReadJSON(xml.in)).Root().Find("y")
		actual, err := WriteXML(*sel)
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, xml.out, actual)
	}
}

func TestXMLNumberParse(t *testing.T) {
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

	xml := "<data><id>4</id><idstr>4</idstr><readings>3.555454</readings><readings>45.04545</readings><readings>324545.04</readings></data>"
	//test get id
	var sel *node.Selection
	sel, _ = node.NewBrowser(module, ReadXML(xml)).Root().Find("data/id")

	found, err := sel.Get()
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
	sel, _ = node.NewBrowser(module, ReadXML(xml)).Root().Find("data/idstr")
	found, err = sel.Get()
	if err != nil {
		t.Error("failed to transmit json", err)
	} else if found == nil {
		t.Error("data/idstr - Target not found, state nil")
	} else {
		if 4 != found.Value().(int) {
			t.Error(found.Value().(int), "!=", 4)
		}
	}

	sel, _ = node.NewBrowser(module, ReadXML(xml)).Root().Find("data/readings")
	found, err = sel.Get()
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

func TestXmlEmpty(t *testing.T) {
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
	in := `<x/>`
	fc.AssertEqual(t, nil, b.Root().InsertFrom(ReadXML(in)))
	fc.AssertEqual(t, val.NotEmpty, actual["x"])
}

func TestReadQualifiedXmlIdentRef(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	in := `{
		"module-test:type":"module-types:derived-type",
		"module-test:type2":"local-type"
	}`
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	fc.AssertEqual(t, nil, b.Root().InsertFrom(ReadJSON(in)))
	fc.AssertEqual(t, "derived-type", actual["type"].(val.IdentRef).Label)
	fc.AssertEqual(t, "local-type", actual["type2"].(val.IdentRef).Label)
}
