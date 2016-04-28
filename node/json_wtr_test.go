package node

import (
	"bytes"
	"testing"
)

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
	m := YangFromString(moduleStr)
	root := map[string]interface{}{
		"l1": []map[string]interface{}{
			map[string]interface{}{"l2" : []map[string]interface{}{
				map[string]interface{}{
						"a" : "hi",
						"b" : "bye",
					},
				},
			},
		},
	}
	b := MapNode(root)
	var json bytes.Buffer
	c := NewContext()
	if err := c.Select(m, b).UpsertInto(NewJsonWriter(&json).Node()).LastErr; err != nil {
		t.Fatal(err)
	}
	actual := json.String()
	expected := `{"l1":[{"l2":[{"a":"hi","b":"bye"}]}]}`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
