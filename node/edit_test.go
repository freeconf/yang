package node_test

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"testing"

	"github.com/freeconf/yang/fc"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
)

var update = flag.Bool("update", false, "update gold files instead of testing against them")

func TestChoiceInAction(t *testing.T) {
	mstr := `module m { prefix ""; namespace ""; revision 0;
		rpc r {
			input {
				choice c {
					leaf x {
						type string;
					}
					leaf y {
						type string;
					}
				}	
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	n := &nodeutil.Basic{}
	var actual string
	n.OnAction = func(r node.ActionRequest) (node.Node, error) {
		actual, _ = nodeutil.WriteJSON(r.Input)
		t.Log(actual)
		return n, nil
	}
	root := node.NewBrowser(m, n).Root()
	expected := `{"x":"hello"}`
	in := nodeutil.ReadJSON(expected)
	if err := root.Find("r").Action(in).LastErr; err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, expected, actual)
}

// https://github.com/freeconf/yang/issues/19
func TestNodeEditDelete(t *testing.T) {
	mstr := `module m { prefix ""; namespace ""; revision 0;
		container c {
			leaf x {
				type string;
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	n := &nodeutil.Basic{}
	n.OnChild = func(r node.ChildRequest) (node.Node, error) {
		return n, nil
	}
	var actual bytes.Buffer
	n.OnBeginEdit = func(r node.NodeRequest) error {
		fmt.Fprintf(&actual, "begin %s(%v)\n", r.Selection.Meta().Ident(), r.Delete)
		return nil
	}
	n.OnEndEdit = func(r node.NodeRequest) error {
		fmt.Fprintf(&actual, "end %s(%v)\n", r.Selection.Meta().Ident(), r.Delete)
		return nil
	}
	root := node.NewBrowser(m, n).Root()
	if err := root.Find("c").Delete(); err != nil {
		t.Fatal(err)
	}
	fc.AssertEqual(t, `begin c(true)
begin m(true)
end c(true)
end m(true)
`, actual.String())
}

func TestEditListNoKey(t *testing.T) {
	mstr := `module m { prefix ""; namespace ""; revision 0;
		list l {
			leaf x {
				type string;
			}
		}
	}`
	data := map[string]interface{}{
		"l": []map[string]interface{}{
			{
				"x": "hi",
			},
			{
				"y": "bye",
			},
		},
	}
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	sel := node.NewBrowser(m, nodeutil.ReflectChild(data)).Root()
	var actual bytes.Buffer
	if err := sel.InsertInto(nodeutil.Dump(nodeutil.Null(), &actual)).LastErr; err != nil {
		t.Error(err)
	}
	fc.Gold(t, *update, actual.Bytes(), "gold/TestEditListNoKey.dmp")
}

func TestChoiceLeafUpsert(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `
		module x {
			revision 0;

			container a {
				choice x {
					case aa {
						leaf aa {
							type string;
						}
						container c {
							leaf cc {
								type string;
							}
						}
					}
					case bb {
						leaf bb {
							type string;
						}
					}
				}
			}
		}
	`)
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"aa": "x",
			"c": map[string]interface{}{
				"cc": "x",
			},
		},
	}
	if err != nil {
		t.Fatal(err)
	}
	b := node.NewBrowser(m, nodeutil.ReflectChild(data))
	sel := b.Root().Find("a")
	err = sel.UpsertFrom(nodeutil.ReadJSON(`
		{
			"bb" : "y"
		}
	`)).LastErr
	if err != nil {
		t.Error(err)
	}
	actual, err := nodeutil.WriteJSON(sel)
	if err != nil {
		t.Fatal(err)
	}
	if actual != `{"bb":"y"}` {
		t.Error(actual)
	}
	_, foundA := data["aa"]
	if foundA {
		t.Error("reflect implemtation should have removed case leaf")
	}
	_, foundC := data["c"]
	if foundC {
		t.Error("reflect implemtation should have removed case container")
	}
}

func TestChoiceContainerUpsert(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, `
		module x {
			revision 0;

			choice x {
				case a {
					container a {
						leaf aa {
							type string;
						}
					}
				}
				case b {
					container b {
						leaf bb {
							type string;
						}
					}
				}
			}
		}
	`)
	if err != nil {
		t.Fatal(err)
	}
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"aa": "x",
		},
	}
	b := node.NewBrowser(m, nodeutil.ReflectChild(data))
	sel := b.Root()
	err = sel.UpsertFrom(nodeutil.ReadJSON(`
		{
			"b" : {"bb" : "y"}
		}
	`)).LastErr
	if err != nil {
		t.Error(err)
	}
	actual, err := nodeutil.WriteJSON(sel)
	if err != nil {
		t.Fatal(err)
	}
	if actual != `{"b":{"bb":"y"}}` {
		t.Error(actual)
	}
	_, foundA := data["aa"]
	if foundA {
		t.Error("reflect implemtation should have removed case leaf")
	}
}

func TestEditor(t *testing.T) {
	mstr := `module m { prefix ""; namespace ""; revision 0;

		leaf x {
			type string;
		}

		container c {
			leaf x {
				type string;
			}
			leaf y {
				type int32;
				default 10;
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
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
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
			expected: `{"c":{"x":"hello","y":10}}`,
		},
		{
			data: map[string]interface{}{
				"c": map[string]interface{}{
					"x": "hello",
					"y": 5,
				},
			},
			expected: `{"c":{"x":"hello","y":5}}`,
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
		bd := nodeutil.ReflectChild(test.data)
		sel := node.NewBrowser(m, bd).Root()
		if test.find != "" {
			sel = sel.Find(test.find)
		}
		if actual, err := nodeutil.WriteJSON(sel); err != nil {
			t.Error(err)
		} else if actual != test.expected {
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
	m, err := parser.LoadModuleFromString(nil, editTestModule)
	if err != nil {
		t.Fatal(err)
	}
	root := testDataRoot()
	bd := nodeutil.ReflectChild(root)
	json := nodeutil.ReadJSON(`{"origin":{"country":"Canada"}}`)

	// UPDATE
	// Here we're testing editing a specific list item. With FindTarget walk controller
	// needs to leave walkstate in a position for WalkTarget controller to make the edit
	// on the right item.
	log.Println("Testing edit")
	sel := node.NewBrowser(m, bd).Root()
	if err := sel.Find("fruits=apple").UpdateFrom(json).LastErr; err != nil {
		t.Fatal(err)
	}
	actual := fc.MapValue(root, "fruits", 1, "origin", "country")
	if actual != "Canada" {
		t.Error("Edit failed", actual)
	}

	// INSERT
	log.Println("Testing insert")
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
	json = nodeutil.ReadJSON(insertData)
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

func TestEditChoiceInGroup(t *testing.T) {
	tests := []struct {
		schema   string
		data     string
		expected string
	}{
		{
			schema: `
				grouping g {
					choice c {
						case a {
							leaf a {
								type string;
							}
						}
					}
				}
		
				uses g;
			`,
			data:     `{"a":"hi"}`,
			expected: `{"a":"hi"}`,
		},
		{
			schema: `
				grouping g {
					leaf e {
						type string;
					}
					choice c {						
						case a {
							container a {
								leaf aa {
									type string;
								}	
							}
						}
						case b {
							container b {
								leaf bb {
									type string;
								}	
							}
						}
					}
				}
				container z {
					leaf c {
						type string;
					}
					uses g;
				}
			`,
			data:     `{"z":{"b":{"bb":"hi"}}}`,
			expected: `{"z":{"b":{"bb":"hi"}}}`,
		},
	}

	for _, test := range tests {
		mstr := fmt.Sprintf(`
			module x {
				revision 0;
				%s
			}`, test.schema)
		m, err := parser.LoadModuleFromString(nil, mstr)
		if err != nil {
			t.Fatal(err)
		}
		data := make(map[string]interface{})
		n := nodeutil.ReflectChild(data)
		b := node.NewBrowser(m, n)
		err = b.Root().UpsertFromSetDefaults(nodeutil.ReadJSON(test.data)).LastErr
		if err != nil {
			t.Fatal(err)
		}
		actual, err := nodeutil.WriteJSON(b.Root())
		if err != nil {
			t.Fatal(err)
		}
		if actual != test.expected {
			t.Error(actual)
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
