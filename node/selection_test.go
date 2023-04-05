package node_test

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/testdata"
)

func TestPeek(t *testing.T) {
	b, _ := testdata.BirdBrowser(`
{
	"bird" : [{
		"name" : "blue jay"
	},{
		"name" : "robin"
	}]
}
`)
	actual := b.Root().Find("bird=robin").Peek(nil)
	if actual == nil {
		t.Error("no value from peek")
	} else if b, ok := actual.(*testdata.Bird); !ok {
		t.Errorf("not a bird %v", actual)
	} else if b.Name != "robin" {
		t.Error(b.Name)
	}
}

func TestNext(t *testing.T) {
	fc.DebugLog(true)
	b, _ := testdata.BirdBrowser(`
{
	"bird" : [{
		"name" : "blue jay"
	},{
		"name" : "robin"
	}]
}
`)
	i := b.Root().Find("bird").First()
	v, _ := i.Selection.Find("name").Get()
	fc.AssertEqual(t, "blue jay", v.String())
	i = i.Next()
	v, _ = i.Selection.Find("name").Get()
	fc.AssertEqual(t, "robin", v.String())
	i = i.Next()
	if !i.Selection.IsNil() {
		t.Error("expected no value")
	}
}

func TestReplaceFrom(t *testing.T) {
	fc.DebugLog(true)
	b, _ := testdata.BirdBrowser(`
{
	"bird" : [{
		"name" : "blue jay"
	},{
		"name" : "robin",
		"wingspan": 10,
		"species":{
			"name" : "thrush"
		}
	}]
}
`)
	root := b.Root()
	js := nodeutil.JSONWtr{}

	// container
	err := root.Find("bird=robin/species").ReplaceFrom(nodeutil.ReadJSON(`
		{"species":{"name":"dragon"}}
	`))
	fc.AssertEqual(t, nil, err)
	actual, err := js.JSON(root.Find("bird=robin"))
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"name":"robin","wingspan":10,"species":{"name":"dragon"}}`, actual)

	// list item
	err = root.Find("bird=robin").ReplaceFrom(nodeutil.ReadJSON(`
		{"bird":[{"name": "robin", "wingspan":11}]}
	`))
	fc.AssertEqual(t, nil, err)
	actual, err = js.JSON(root)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[{"name":"blue jay","wingspan":0},{"name":"robin","wingspan":11}]}`, actual)
}
