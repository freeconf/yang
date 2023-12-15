package nodeutil

import (
	"fmt"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/patch/xml"
	"github.com/freeconf/yang/source"
)

func TestXmlWtr2(t *testing.T) {
	tests := []struct {
		Yang     string
		Val      any
		expected string
		enumAsId bool
	}{
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      &struct{ X string }{X: "a"},
			expected: `<m xmlns="mm"><x>a</x></m>`,
		},
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      &struct{ X int }{X: 99},
			expected: `<m xmlns="mm"><x>99</x></m>`,
		},
		{
			Yang:     `leaf x { type enumeration { enum zero; enum one; }}`,
			Val:      &struct{ X string }{X: "zero"},
			expected: `<m xmlns="mm"><x>zero</x></m>`,
		},
		{
			Yang:     `leaf x { type enumeration { enum five {value 5;} enum six; }}`,
			Val:      &struct{ X string }{X: "six"},
			expected: `<m xmlns="mm"><x>6</x></m>`,
			enumAsId: true,
		},
	}
	for _, test := range tests {
		m, err := parser.LoadModuleFromString(nil, fmt.Sprintf(`module m { namespace "mm"; %s }`, test.Yang))
		if err != nil {
			t.Fatal(err)
		}
		b := node.NewBrowser(m, &Node{Object: test.Val})
		w := &XMLWtr2{
			XMLName:   XmlName(m),
			EnumAsIds: test.enumAsId,
			ns:        "mm",
		}
		b.Root().UpdateInto(w)
		actual, err := xml.Marshal(w)
		fc.RequireEqual(t, nil, err)
		fc.AssertEqual(t, test.expected, string(actual))
	}
}

func TestXmlWriterListInList2(t *testing.T) {
	moduleStr := `
module m {
	prefix "t";
	namespace "t";
	revision 0000-00-00 {
		description "x";
	}
	typedef td {
		type string;
	}
	container c1 {
		list l1 {
			list l2 {
				key "a";
				leaf a {
					type td;
				}
				leaf b {
					type string;
				}
			}
		}
	}
}
	`
	m, _ := parser.LoadModuleFromString(nil, moduleStr)
	root := map[string]interface{}{
		"c1": map[string]interface{}{
			"l1": []map[string]interface{}{
				{
					"l2": []map[string]interface{}{
						{
							"a": "hi",
							"b": "bye",
						},
					},
				},
			},
		},
	}
	b := &Node{Object: root}
	c1 := sel(node.NewBrowser(m, b).Root().Find("c1"))
	actual, err := WriteXMLFrag(c1, false)
	if err != nil {
		t.Fatal(err)
	}
	expected := `<l1 xmlns="t"><l2><a>hi</a><b>bye</b></l2></l1>`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestXmlAnyData2(t *testing.T) {
	moduleStr := `
module m {
	prefix "t";
	namespace "t";
	revision 0000-00-00 {
		description "x";
	}
	container c1 {
		list l1 {
			list l2 {
				key "a";
				leaf a {
					type any;
				}
				leaf b {
					type any;
				}
			}
		}
	}
}
	`
	m, _ := parser.LoadModuleFromString(nil, moduleStr)
	root := map[string]interface{}{
		"c1": map[string]interface{}{
			"l1": []map[string]interface{}{
				{
					"l2": []map[string]interface{}{
						{
							"a": "hi",
							"b": 99,
						},
					},
				},
			},
		},
	}
	b := &Node{Object: root}
	c1 := sel(node.NewBrowser(m, b).Root().Find("c1"))
	actual, err := WriteXMLDoc(c1, false)
	if err != nil {
		t.Fatal(err)
	}
	expected := `<c1 xmlns="t"><l1><l2><a>hi</a><b>99</b></l2></l1></c1>`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestQualifiedXmlIdentityRef2(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	d := map[string]interface{}{
		"type": "derived-type",
	}
	b := node.NewBrowser(m, &Node{Object: d})
	actual, err := WriteXMLFrag(sel(b.Root().Find("type")), false)
	if err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, `<type xmlns="http://test.org/ns/yang/module/test">module-types:derived-type</type>`, actual)
}

func TestXmlLeafList2(t *testing.T) {
	moduleStr := `
module m {
	prefix "t";
	namespace "t";
	revision 0000-00-00 {
		description "x";
	}
	container c {
		leaf-list l {
			type string;
		}
	}
}
	`
	m, _ := parser.LoadModuleFromString(nil, moduleStr)
	root := map[string]interface{}{
		"c": map[string]interface{}{
			"l": []interface{}{
				"hi",
				"bye",
			},
		},
	}

	b := &Node{Object: root}
	c := sel(node.NewBrowser(m, b).Root().Find("c"))
	actual, err := WriteXMLDoc(c, false)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `<c xmlns="t"><l>hi</l><l>bye</l></c>`, actual)
}
