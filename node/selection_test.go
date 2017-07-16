package node

import "testing"
import "github.com/c2stack/c2g/c2"

func Test_Peek(t *testing.T) {
	b, _ := BirdBrowser(".", `
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
	} else if b, ok := actual.(*Bird); !ok {
		t.Errorf("not a bird %v", actual)
	} else if b.Name != "robin" {
		t.Error(b.Name)
	}
}

func Test_Next(t *testing.T) {
	c2.DebugLog(true)
	b, _ := BirdBrowser(".", `
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
	if err := c2.CheckEqual("blue jay", v.Str); err != nil {
		t.Error(err)
	}
	i = i.Next()
	v, _ = i.Selection.GetValue("name")
	if err := c2.CheckEqual("robin", v.Str); err != nil {
		t.Error(err)
	}
	i = i.Next()
	if !i.Selection.IsNil() {
		t.Error("expected no value")
	}
}
