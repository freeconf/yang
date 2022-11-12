package nodeutil

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/val"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

func TestJsonWriterLeafs(t *testing.T) {
	fc.DebugLog(true)
	tests := []struct {
		Yang     string
		Val      val.Value
		expected string
		enumAsId bool
	}{
		{
			Yang:     `leaf-list x { type string;}`,
			Val:      val.StringList([]string{"a", "b"}),
			expected: `"x":["a","b"]`,
		},
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      val.String("a"),
			expected: `"x":"a"`,
		},
		{
			Yang:     `leaf x { type union { type int32; type string;}}`,
			Val:      val.Int32(99),
			expected: `"x":99`,
		},
		{
			Yang:     `leaf x { type enumeration { enum zero; enum one; }}`,
			Val:      val.Enum{Id: 0, Label: "zero"},
			expected: `"x":"zero"`,
		},
		{
			Yang:     `leaf x { type enumeration { enum five {value 5;} enum six; }}`,
			Val:      val.Enum{Id: 6, Label: "six"},
			expected: `"x":6`,
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
		w := &JSONWtr{
			_out:                     buf,
			EnumAsIds:                test.enumAsId,
			QualifyNamespaceDisabled: true,
		}
		w.writeValue(node.NewRootPath(m), m.DataDefinitions()[0], test.Val)
		buf.Flush()
		fc.AssertEqual(t, test.expected, actual.String())
	}
}

func TestJsonWriterListInList(t *testing.T) {
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
	`
	m, _ := parser.LoadModuleFromString(nil, moduleStr)
	root := map[string]interface{}{
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
	}
	b := ReflectChild(root)
	sel := node.NewBrowser(m, b).Root()
	actual, err := WriteJSON(sel)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"l1":[{"l2":[{"a":"hi","b":"bye"}]}]}`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestJsonAnyData(t *testing.T) {
	tests := []struct {
		anything interface{}
		expected string
	}{
		{
			anything: map[string]interface{}{
				"a": "A",
				"b": "B",
			},
			expected: `"x":{"a":"A","b":"B"}`,
		},
		{
			anything: []interface{}{
				map[string]interface{}{
					"a": "A",
				},
				map[string]interface{}{
					"b": "B",
				},
			},
			expected: `"x":[{"a":"A"},{"b":"B"}]`,
		},
	}
	for _, test := range tests {
		b := &meta.Builder{}
		m := b.Module("m", nil)
		var actual bytes.Buffer
		buf := bufio.NewWriter(&actual)
		w := &JSONWtr{
			_out:                     buf,
			QualifyNamespaceDisabled: true,
		}
		l := b.Leaf(m, "x")
		w.writeValue(node.NewRootPath(m), l, val.Any{Thing: test.anything})
		buf.Flush()
		fc.AssertEqual(t, test.expected, actual.String())
	}
}

func TestQualifiedJson(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "example-barmod")
	d := map[string]interface{}{
		"top": map[string]interface{}{
			"foo": 10,
			"bar": true,
		},
	}
	b := node.NewBrowser(m, ReflectChild(d))
	actual, err := WriteJSON(b.Root())
	if err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, `{"top":{"foo":10,"bar":true}}`, actual)
}
