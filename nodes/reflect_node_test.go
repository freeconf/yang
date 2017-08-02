package nodes

import (
	"strings"
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

type TestMessage struct {
	Message struct {
		Hello string
	}
}

func TestMarshal(t *testing.T) {
	mstr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	container message {
		leaf hello {
			type string;
		}
	}
}
`
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	var obj TestMessage
	c := ReflectNode(&obj)
	sel := node.NewBrowser(m, c).Root()
	r := NewJsonReader(strings.NewReader(`{"message":{"hello":"bob"}}`)).Node()
	if err = sel.UpsertFrom(r).LastErr; err != nil {
		t.Fatal(err)
	}
	if obj.Message.Hello != "bob" {
		t.Fatal("Not selected")
	}
}

type TestMessageItem struct {
	Id string
}

func TestMarshalIndex(t *testing.T) {
	mstr := `
module m {
	prefix "";
	namespace "";
	revision 0;
	list messages {
		key "id";
		leaf id {
			type string;
		}
	}
}
`
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	objs := make(map[string]*TestMessageItem)
	marshaller := &MarshalMap{
		Map: objs,
		OnNewItem: func(node.ListRequest) interface{} {
			return &TestMessageItem{}
		},
		OnSelectItem: func(item interface{}) node.Node {
			return ReflectNode(item)
		},
	}
	d := NewJsonReader(strings.NewReader(`{"messages":[{"id":"bob"},{"id":"barb"}]}`))
	sel := node.NewBrowser(m, d.Node()).Root().Find("messages")
	if err = sel.UpsertInto(marshaller.Node()).LastErr; err != nil {
		t.Fatal(err)
	}
	if objs["bob"].Id != "bob" {
		t.Fatal("Not inserted")
	}
	n := marshaller.Node()
	r := node.ListRequest{
		Meta: m.DataDefs().GetFirstMeta().(*meta.List),
		Request: node.Request{
			Selection: sel,
		},
		First: true,
	}
	r.Key, _ = node.NewValues(r.Meta.KeyMeta(), "bob")
	foundByKeyNode, _, nextByKeyErr := n.Next(r)
	if nextByKeyErr != nil {
		t.Fatal(nextByKeyErr)
	}
	if foundByKeyNode == nil {
		t.Error("lookup by key failed")
	}
	r.Key = []val.Value{}
	foundFirstNode, _, nextFirstErr := n.Next(r)
	if nextFirstErr != nil {
		t.Fatal(nextFirstErr)
	}
	if foundFirstNode == nil {
		t.Error("lookup by next failed")
	}
}
