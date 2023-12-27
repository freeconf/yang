package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"

	"github.com/freeconf/yang/meta"
)

func TestFindPathSlice(t *testing.T) {
	mstr := `
		module food {
			namespace ""; prefix ""; revision 0;
			container fruit {
				container apple {
					leaf kind {
						type string;
					}
				}
				action peel {
					input {}
				}
				notification spoil {}
			}
			list country {
				key "name";
				leaf name {
					type string;
				}
				container detail {
					leaf ally {
						type string;
					}
				}
				action vote {
					input {}
				}
				notification anarchy {}
			}
		}
	`
	data := map[string]interface{}{
		"fruit": map[string]interface{}{
			"apple": map[string]interface{}{
				"kind": "macintosh",
			},
		},
		"country": []map[string]interface{}{
			{
				"name": "US",
				"detail": map[string]interface{}{
					"ally": "Canada",
				},
			},
			{
				"name": "Canada",
			},
		},
	}

	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	root := node.NewBrowser(m, nodeutil.ReflectChild(data)).Root()
	tests := []struct {
		path           string
		customExpected string
		key            string
	}{
		{
			path: "fruit/apple",
		},
		{
			path: "fruit/apple/kind",
		},
		{
			path: "country",
		},
		{
			path: "country=US",
			key:  "US",
		},
		{
			path: "country=US/detail",
		},
		{
			path: "fruit/peel",
		},
		{
			path: "fruit/spoil",
		},
		{
			path: "country=US/vote",
		},
		{
			path: "country=US/anarchy",
		},
		{
			path:           "food:country",
			customExpected: "country",
		},
	}
	for _, test := range tests {
		t.Log("Testing path", test.path)
		found, err := root.Find(test.path)
		fc.RequireEqual(t, nil, err)
		fc.RequireEqual(t, true, found != nil, test.path+" not found")
		actual := found.Path.StringNoModule()
		if test.customExpected != "" {
			fc.AssertEqual(t, test.customExpected, actual)
		} else {
			fc.AssertEqual(t, test.path, actual)
		}
		if test.key != "" {
			fc.AssertEqual(t, test.key, found.Key()[0].String())
		}
	}
}

const walkTestModule = `
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
		choice shipment {
			case water {
				container boat {
					leaf name {
						type string;
					}
				}
			}
			case air {
				container plane {
					leaf name {
						type string;
					}
				}
			}
		}
	}
}
`

func TestFindPathIntoListItemContainer(t *testing.T) {
	m, root := LoadPathTestData()
	tests := []string{
		"fruits=apple/origin",
		"fruits=apple/boat",
	}
	for _, test := range tests {
		root := node.NewBrowser(m, nodeutil.ReflectChild(root)).Root()
		target, err := root.Find(test)
		fc.RequireEqual(t, nil, err)
		fc.AssertEqual(t, true, target != nil, "Could not find target "+test)
	}
}

func LoadPathTestData() (*meta.Module, map[string]interface{}) {
	// avoid using json to load because that needs edit/INSERT and
	// we don't want to use code to load seed data that we're trying to test
	data := map[string]interface{}{
		"fruits": []map[string]interface{}{
			{
				"name": "banana",
				"origin": map[string]interface{}{
					"country": "Brazil",
				},
				"plane": map[string]interface{}{
					"name": "747c",
				},
			},
			{
				"name": "apple",
				"origin": map[string]interface{}{
					"country": "US",
				},
				"boat": map[string]interface{}{
					"name": "SS Hudson",
				},
			},
		},
	}
	m, err := parser.LoadModuleFromString(nil, walkTestModule)
	if err != nil {
		panic(err)
	}
	return m, data
}
