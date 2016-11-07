package node

import (
	"bytes"
	"testing"
)

func TestEditor(t *testing.T) {
	mstr := `module m { prefix ""; namespace ""; revision 0;

		leaf x {
			type string;
		}

		container c {
			leaf x {
				type string;
			}
		}
		list l {
			key "x";
			leaf x {
				type string;
			}
			container c {
				leaf x {
					type string;
				}
			}
		}
	}`
	m := YangFromString(mstr)

	tests := []struct {
		find string
		data map[string]interface{}
		expected string
	}{
		{
			data: map[string]interface{}{
				"x": "hello",
			},
			expected : `{"x":"hello"}`,

		},
		{
			data: map[string]interface{}{
				"c": map[string] interface{} {
					"x" : "hello",
				},
			},
			expected : `{"c":{"x":"hello"}}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string] interface{} {
					{
						"x" : "hello",
					},
				},
			},
			expected : `{"l":[{"x":"hello"}]}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string] interface{} {
					{
						"x" : "hello",
						"c" : map[string]interface{} {
							"x" : "goodbye",
						},
					},
				},
			},
			find : "l=hello",
			expected : `{"x":"hello","c":{"x":"goodbye"}}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string] interface{} {
					{
						"x" : "hello",
						"c" : map[string]interface{} {
							"x" : "goodbye",
						},
					},
				},
			},
			find : "l",
			expected : `{"l":[{"x":"hello","c":{"x":"goodbye"}}]}`,
		},
	}

	for _, test := range tests {
		bd := MapNode(test.data)
		var buff bytes.Buffer
		sel := NewBrowser2(m, bd).Root()
		if test.find != "" {
			sel = sel.Find(test.find)
		}
		if err := sel.InsertInto(NewJsonWriter(&buff).Node()).LastErr; err != nil {
			t.Error(err)
		}
		actual := buff.String()
		if actual != test.expected {
			t.Errorf("\nExpected:%s\n  Actual:%s", test.expected, actual)
		}
	}
}