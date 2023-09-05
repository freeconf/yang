package nodeutil

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
	"github.com/freeconf/yang/val"
)

func TestXmlWriterLeafs(t *testing.T) {
	fc.DebugLog(true)
	tests := []struct {
		Yang     string
		Val      val.Value
		expected string
		enumAsId bool
	}{
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      val.String("a"),
			expected: `<x xmlns="m">a</x>`,
		},
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      val.Int32(99),
			expected: `<x xmlns="m">99</x>`,
		},
		{
			Yang:     `leaf x { type enumeration { enum zero; enum one; }}`,
			Val:      val.Enum{Id: 0, Label: "zero"},
			expected: `<x xmlns="m">zero</x>`,
		},
		{
			Yang:     `leaf x { type enumeration { enum five {value 5;} enum six; }}`,
			Val:      val.Enum{Id: 6, Label: "six"},
			expected: `<x xmlns="m">6</x>`,
			enumAsId: true,
		},
	}
	for _, test := range tests {
		m, err := parser.LoadModuleFromString(nil, fmt.Sprintf(`module m { namespace ""; %s }`, test.Yang))
		if err != nil {
			t.Fatal(err)
		}
		var actual bytes.Buffer
		buf := bufio.NewWriter(&actual)
		w := &XMLWtr{
			_out:      buf,
			EnumAsIds: test.enumAsId,
		}
		w.writeLeafElement("m", &node.Path{Parent: &node.Path{Meta: m}, Meta: m.DataDefinitions()[0]}, test.Val)
		buf.Flush()
		fc.AssertEqual(t, test.expected, actual.String())
	}
}

func TestXmlWriterListInList(t *testing.T) {
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
	b := ReflectChild(root)
	sel, _ := node.NewBrowser(m, b).Root().Find("c1")
	actual, err := WriteXML(*sel)
	if err != nil {
		t.Fatal(err)
	}
	expected := `<c1 xmlns="t"><l1><l2><a>hi</a><b>bye</b></l2></l1></c1>`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestXmlAnyData(t *testing.T) {
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
	b := ReflectChild(root)
	sel, _ := node.NewBrowser(m, b).Root().Find("c1")
	actual, err := WriteXML(*sel)
	if err != nil {
		t.Fatal(err)
	}
	expected := `<c1 xmlns="t"><l1><l2><a>hi</a><b>99</b></l2></l1></c1>`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
func TestQualifiedXmlIdentityRef(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	d := map[string]interface{}{
		"type": "derived-type",
	}
	b := node.NewBrowser(m, ReflectChild(d))
	wtr := &XMLWtr{}
	sel, _ := b.Root().Find("type")
	actual, err := wtr.XML(*sel)
	if err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, `<type xmlns="http://test.org/ns/yang/module/test">module-types:derived-type</type>`, actual)
}

func TestXmlLeafList(t *testing.T) {
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

	b := ReflectChild(root)
	sel, _ := node.NewBrowser(m, b).Root().Find("c")
	actual, err := WriteXML(*sel)
	if err != nil {
		t.Fatal(err)
	}
	expected := `<c xmlns="t"><l>hi</l><l>bye</l></c>`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
