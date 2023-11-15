package nodeutil

import (
	"bytes"
	"strings"
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
	module, err := parser.LoadModuleFromString(nil, moduleStr)
	fc.RequireEqual(t, nil, err)
	xml := `
	<xml-test xmlns="t">
		<hobbies>
			<name>birding</name>
			<favorite>
				<common-name>towhee</common-name>
				<extra>double-mint</extra>
				<location>out back</location>
			</favorite>
		</hobbies>
		<hobbies>
			<name>hockey</name>
			<favorite>
				<common-name>bruins</common-name>
				<location>Boston</location>
			</favorite>
		</hobbies>
	</xml-test>
	`
	tests := []string{
		"hobbies",
		"hobbies=birding",
		"hobbies=birding/favorite",
	}
	for _, test := range tests {
		sel := node.NewBrowser(module, readXml(xml)).Root()
		found, err := sel.Find(test)
		fc.RequireEqual(t, nil, err, test)
		fc.RequireEqual(t, true, found != nil, test)
		fc.AssertEqual(t, "xml-test/"+test, found.Path.String())
	}
}

func TestXMLNumberParse(t *testing.T) {
	moduleStr := `
module xml-test {
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
	fc.RequireEqual(t, nil, err)

	xml := `
	<xml-test xmlns="t">
		<data>
			<id>4</id>
			<idstr>4</idstr>
			<readings>3.555454</readings>
			<readings>45.04545</readings>
			<readings>324545.04</readings>
		</data>
	</xml-test>`

	//test get id
	root := node.NewBrowser(module, readXml(xml)).Root()

	data := sel(root.Find("data/id"))
	found, err := data.Get()
	fc.RequireEqual(t, nil, err, "failed to transmit json")
	fc.RequireEqual(t, true, found != nil, "data/id - Target not found, state nil")
	fc.AssertEqual(t, 4, found.Value().(int))

	//test get idstr
	data = sel(root.Find("data/idstr"))
	found, err = data.Get()
	fc.RequireEqual(t, nil, err, "failed to transmit json")
	fc.RequireEqual(t, true, found != nil, "data/idstr - Target not found, state nil")
	fc.AssertEqual(t, 4, found.Value().(int))

	data = sel(root.Find("data/readings"))
	found, err = data.Get()
	fc.RequireEqual(t, nil, err, "failed to transmit json")
	fc.RequireEqual(t, true, found != nil, "data/readings - Target not found, state nil")
	expected := []float64{3.555454, 45.04545, 324545.04}
	readings := found.Value().([]float64)
	fc.AssertEqual(t, expected, readings)
}

func readXml(data string) *XmlNode {
	n, err := ReadXMLDoc(strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	return n
}

func TestXmlEmpty(t *testing.T) {
	moduleStr := `
module xml-test {
	leaf x {
		type empty;
	}
}
	`
	m, err := parser.LoadModuleFromString(nil, moduleStr)
	fc.AssertEqual(t, nil, err)
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, &Node{Object: actual})
	in := `<xml-test><x/></xml-test>`
	fc.AssertEqual(t, nil, b.Root().InsertFrom(readXml(in)))
	fc.AssertEqual(t, val.NotEmpty, actual["x"])
}

func TestReadQualifiedXmlIdentRef(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	in := `
	<module-test xmlns="http://test.org/ns/yang/module/test">
	        <type>module-types:derived-type</type>
		    <type2>local-type</type2>
	</module-test>`
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	fc.AssertEqual(t, nil, b.Root().InsertFrom(readXml(in)))
	fc.AssertEqual(t, "derived-type", actual["type"].(val.IdentRef).Label)
	fc.AssertEqual(t, "local-type", actual["type2"].(val.IdentRef).Label)
}

func TestXmlChoice(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "choice")
	in := `<choice><z>here</z></choice>`
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	fc.AssertEqual(t, nil, b.Root().InsertFrom(readXml(in)))
	fc.AssertEqual(t, "here", actual["z"])
}

func TestXmlRdrListByRow(t *testing.T) {
	moduleStr := `
module xml-test {
	leaf x {
		type string;
	}
	list y {
		leaf z {
			type string;
		}
	}
}
	`
	m, err := parser.LoadModuleFromString(nil, moduleStr)
	fc.RequireEqual(t, nil, err)
	in := `
	<xml-test>
	    <x>Exs</x>
		<y><z>row 0</z></y>
		<y><z>row 1</z></y>
	</xml-test>
	`
	b := node.NewBrowser(m, readXml(in))
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"x":"Exs","y":[{"z":"row 0"},{"z":"row 1"}]}`, actual)
}

func TestXmlConflict(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "conflict")
	in := `
	<conflict xmlns="zero">
		<x>zero</x>
		<x xmlns="one">one</x>
		<x xmlns="two">two</x>
	</conflict>`
	b := node.NewBrowser(m, readXml(in))
	var actual bytes.Buffer
	w := NewJSONWtr(&actual)
	w.QualifyNamespace = true
	err := b.Root().UpdateInto(w.Node())
	fc.RequireEqual(t, nil, err)
	// this is wrong, should have all three, but there is bug in
	// underlying edit.go that doesn't request all three "x" fields
	fc.AssertEqual(t, `{"conflict:x":"zero"}`, actual.String())
}
