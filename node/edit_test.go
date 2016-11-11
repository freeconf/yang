package node

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
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
		find     string
		data     map[string]interface{}
		expected string
	}{
		{
			data: map[string]interface{}{
				"x": "hello",
			},
			expected: `{"x":"hello"}`,
		},
		{
			data: map[string]interface{}{
				"c": map[string]interface{}{
					"x": "hello",
				},
			},
			expected: `{"c":{"x":"hello"}}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string]interface{}{
					{
						"x": "hello",
					},
				},
			},
			expected: `{"l":[{"x":"hello"}]}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string]interface{}{
					{
						"x": "hello",
						"c": map[string]interface{}{
							"x": "goodbye",
						},
					},
				},
			},
			find:     "l=hello",
			expected: `{"x":"hello","c":{"x":"goodbye"}}`,
		},
		{
			data: map[string]interface{}{
				"l": []map[string]interface{}{
					{
						"x": "hello",
						"c": map[string]interface{}{
							"x": "goodbye",
						},
					},
				},
			},
			find:     "l",
			expected: `{"l":[{"x":"hello","c":{"x":"goodbye"}}]}`,
		},
	}

	for _, test := range tests {
		bd := MapNode(test.data)
		var buff bytes.Buffer
		sel := NewBrowser(m, bd).Root()
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

const editTestModule = `
module food {
	prefix "x";
	namespace "y";
	revision 0000-00-00 {
		description "";
	}
	list fruits  {
		key "name";
		leaf name {
			type string;
		}
		container origin {
			leaf country {
				type string;
			}
		}
	}
}
`

func TestEditListItem(t *testing.T) {
	m := YangFromString(editTestModule)
	root := testDataRoot()
	bd := MapNode(root)
	json := NewJsonReader(strings.NewReader(`{"origin":{"country":"Canada"}}`)).Node()

	// UPDATE
	// Here we're testing editing a specific list item. With FindTarget walk controller
	// needs to leave walkstate in a position for WalkTarget controller to make the edit
	// on the right item.
	log.Println("Testing edit\n")
	sel := NewBrowser(m, bd).Root()
	if err := sel.Find("fruits=apple").UpdateFrom(json).LastErr; err != nil {
		t.Fatal(err)
	}
	actual := MapValue(root, "fruits.1.origin.country")
	if actual != "Canada" {
		t.Error("Edit failed", actual)
	}

	// INSERT
	log.Println("Testing insert\n")
	insertData := `{
  "fruits": [
    {
      "name": "pear",
      "origin": {
        "country": "Columbia"
      }
    },
    {
      "name": "guava",
      "origin": {
        "country": "Boliva"
      }
    }
  ]
}`
	json = NewJsonReader(strings.NewReader(insertData)).Node()
	if err := sel.Find("fruits").InsertFrom(json).LastErr; err != nil {
		t.Fatal(err)
	}
	actual, found := root["fruits"]
	if !found {
		t.Error("fruits not found")
	} else {
		fruits := actual.([]map[string]interface{})
		if len(fruits) != 4 {
			t.Error("Expected 4 fruits but got ", len(fruits))
		}
	}
}

func testDataRoot() map[string]interface{} {
	return map[string]interface{}{
		"fruits": []map[string]interface{}{
			map[string]interface{}{
				"name": "banana",
				"origin": map[string]interface{}{
					"country": "Brazil",
				},
			},
			map[string]interface{}{
				"name": "apple",
				"origin": map[string]interface{}{
					"country": "US",
				},
			},
		},
	}
}

func YangFromString(s string) *meta.Module {
	m, err := yang.LoadModuleCustomImport(s, nil)
	if err != nil {
		panic(err)
	}
	return m
}
