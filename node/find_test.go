package node

import (
	"testing"
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
				if test.key != found.Key()[0].Str {
					t.Errorf("\nExpected:%s\n  Actual:%s", test.key, found.Key()[0].Str)
				}
			}
		}
	}
}
