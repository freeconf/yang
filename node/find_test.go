package node

import (
	"testing"

	"github.com/c2stack/c2g/meta"
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

	m := YangFromString(mstr)
	root := NewBrowser(m, MapNode(data)).Root()
	tests := []struct {
		path string
		key  string
	}{
		{
			path: "fruit/apple",
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
	}
	for _, test := range tests {
		t.Log("Testing path", test.path)
		target := NewPathSlice(test.path, m)
		found := root.FindSlice(target)
		if found.LastErr != nil {
			t.Error(found.LastErr)
		} else if found.IsNil() {
			t.Error(test.path, " not found")
		} else {
			actual := found.Meta().GetIdent()
			expected := target.Tail.meta.GetIdent()
			if expected != actual {
				t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
			}
			if test.key != "" {
				if test.key != found.Key()[0].String() {
					t.Errorf("\nExpected:%s\n  Actual:%s", test.key, found.Key()[0].String())
				}
			}
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
		target := NewBrowser(m, MapNode(root)).Root().Find(test)
		if target.LastErr != nil {
			t.Fatal(target.LastErr)
		} else if target.IsNil() {
			t.Fatal("Could not find target " + test)
		}
	}
}

func LoadPathTestData() (*meta.Module, map[string]interface{}) {
	// avoid using json to load because that needs edit/INSERT and
	// we don't want to use code to load seed data that we're trying to test
	data := map[string]interface{}{
		"fruits": []map[string]interface{}{
			map[string]interface{}{
				"name": "banana",
				"origin": map[string]interface{}{
					"country": "Brazil",
				},
				"plane": map[string]interface{}{
					"name": "747c",
				},
			},
			map[string]interface{}{
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
	return YangFromString(walkTestModule), data
}
