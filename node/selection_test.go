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
	sel, err := b.Root().Find("bird=robin")
	fc.RequireEqual(t, nil, err)
	actual := sel.Peek(nil)
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
	sel, err := b.Root().Find("bird")
	fc.RequireEqual(t, nil, err)
	i, err := sel.First()
	fc.RequireEqual(t, nil, err)
	name, err := i.Selection.Find("name")
	fc.RequireEqual(t, nil, err)
	v, err := name.Get()
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, "blue jay", v.String())
	i, err = i.Next()
	fc.RequireEqual(t, nil, err)
	name, err = i.Selection.Find("name")
	fc.RequireEqual(t, nil, err)
	v, err = name.Get()
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, "robin", v.String())
	i, err = i.Next()
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, true, i.Selection == nil, "expected no value")
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
	sel, err := root.Find("bird=robin/species")
	fc.RequireEqual(t, nil, err)
	n, _ := nodeutil.ReadJSON(`
		{"class":"dragon"}
	`)
	err = sel.ReplaceFrom(n)
	fc.AssertEqual(t, nil, err)
	sel, err = root.Find("bird=robin")
	fc.RequireEqual(t, nil, err)
	actual, err := js.JSON(sel)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"name":"robin","wingspan":10,"species":{"class":"dragon"}}`, actual)

	// list item
	sel, err = root.Find("bird=robin")
	fc.RequireEqual(t, nil, err)
	n, _ = nodeutil.ReadJSON(`
		{"name": "robin", "wingspan":11}
	`)
	sel.ReplaceFrom(n)
	fc.AssertEqual(t, nil, err)
	actual, err = js.JSON(root)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[{"name":"blue jay","wingspan":0},{"name":"robin","wingspan":11}]}`, actual)

	// whole list
	n, _ = nodeutil.ReadJSON(`
		{"bird":[{"name":"blue jay"},{"name":"malak","species":{"class":"jedi"}}]}
	`)
	root.ReplaceFrom(n)
	fc.AssertEqual(t, nil, err)
	actual, err = js.JSON(root)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"bird":[{"name":"blue jay","wingspan":0},{"name":"malak","wingspan":0,"species":{"class":"jedi"}}]}`, actual)
}
