package node

import (
	"bytes"
	"testing"
	"github.com/c2stack/c2g/meta"
	"bufio"
	"github.com/c2stack/c2g/c2"
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
	sel := NewBrowser2(m, b).Root().Selector()
	if err := sel.UpsertInto(NewJsonWriter(&json).Node()).LastErr; err != nil {
		t.Fatal(err)
	}
	actual := json.String()
	expected := `{"l1":[{"l2":[{"a":"hi","b":"bye"}]}]}`
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestJsonAnyData(t *testing.T) {
	var actual bytes.Buffer
	buf := bufio.NewWriter(&actual)
	w := &JsonWriter{
		out: buf,
	}
	m := meta.NewLeaf("x", "na")
	anything := map[string]interface{} {
		"a" : "A",
		"b" : "B",
	}
	v := &Value{Type:meta.NewDataType(nil, "any"), AnyData: anything}
	w.writeValue(m, v)
	buf.Flush()
	if err := c2.CheckEqual(`"x":{"a":"A","b":"B"}`, actual.String()); err != nil {
		t.Error(err)
	}
}
