package node
import (
	"testing"
	"github.com/dhubler/c2g/meta/yang"
	"strings"
	"bytes"
)

var mstr = `
module m {
	namespace "";
	prefix "";
	revision 0;
	container a {
		container b {
			leaf x {
				type string;
			}
		}
	}
	list p {
		key "k";
		leaf k {
			type string;
		}
		container q {
			leaf s {
				type string;
			}
		}
		list r {
			leaf z {
				type int32;
			}
		}
	}
}
`

func TestCollectionWrite(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		data string
		path string
	} {
		{
			`{"a":{"b":{"x":"waldo"}}}`,
			"a.b.x",
		},
		{
			`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
			"p.1.k",
		},
	}
	for _, test := range tests {
		root := make(map[string]interface{})
		bd := MapNode(root)
		sel := NewBrowser2(m, bd).Root().Selector()
		if err = sel.InsertFrom(NewJsonReader(strings.NewReader(test.data)).Node()).LastErr; err != nil {
			t.Error(err)
		}
		actual := MapValue(root, test.path)
		if actual != "waldo" {
			t.Error(actual)
		}
	}
}

func TestCollectionRead(t *testing.T) {
	m := YangFromString(mstr)
	tests := []struct {
		root map[string]interface{}
		expected string
	} {
		{
			map[string]interface{}{
				"a" : map[string]interface{}{
					"b" : map[string]interface{}{
						"x" : "waldo",
					},
				},
			},
			`{"a":{"b":{"x":"waldo"}}}`,
		},
		{
			map[string]interface{}{
				"p" : []interface{}{
					map[string]interface{}{"k" :"walter"},
					map[string]interface{}{"k" :"waldo"},
					map[string]interface{}{"k" :"weirdo"},
				},
			},
			`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
		},
	}
	for _, test := range tests {
		bd := MapNode(test.root)
		var buff bytes.Buffer
		sel := NewBrowser2(m, bd).Root().Selector()
		if err := sel.InsertInto(NewJsonWriter(&buff).Node()).LastErr; err != nil {
			t.Error(err)
		}
		actual := buff.String()
		if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}

