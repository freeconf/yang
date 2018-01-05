package node_test

import (
	"testing"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/testdata"
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
	c2.DebugLog(true)
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
	v, _ := i.Selection.GetValue("name")
	c2.AssertEqual(t, "blue jay", v.String())
	i = i.Next()
	v, _ = i.Selection.GetValue("name")
	c2.AssertEqual(t, "robin", v.String())
	i = i.Next()
	if !i.Selection.IsNil() {
		t.Error("expected no value")
	}
}
