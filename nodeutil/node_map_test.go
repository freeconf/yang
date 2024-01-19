package nodeutil

import (
	"reflect"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

func TestMapAsContainer(t *testing.T) {
	app := map[string]any{
		"c": map[string]string{
			"z": "hi",
		},
		"l": "before",
	}
	mstr := `module x {
		container c {
			leaf z {
				type string;
			}
		}
		leaf l {
			type string;
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	fc.RequireEqual(t, nil, err)
	ref := &Node{}
	x := newMapAsContainer(ref, reflect.ValueOf(app))

	cMeta := meta.Find(m, "c")

	cObj, err := x.get(cMeta)
	fc.RequireEqual(t, nil, err)
	c, valid := cObj.Interface().(map[string]string)
	fc.RequireEqual(t, true, valid)
	fc.AssertEqual(t, app["c"], c)

	lMeta := meta.Find(m, "l")
	lObj, err := x.get(lMeta)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, "before", lObj.Interface().(string))

	x.set(lMeta, reflect.ValueOf("after"))
	fc.AssertEqual(t, "after", app["l"])
}

func TestMapAsList(t *testing.T) {
	mstr := `module x {
		list p {
			key g;
			leaf g {
				type int32;
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	fc.RequireEqual(t, nil, err)
	app := mapAsListApp{
		P: map[int]*MapAsListP{
			100: {G: 100},
			55:  {G: 55},
			999: {G: 999},
		},
	}
	b := node.NewBrowser(m, &Node{Object: &app})
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"p":[{"g":55},{"g":100},{"g":999}]}`, actual)

	byKey, err := b.Root().Find("p=55")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, byKey != nil)
	actual, err = WriteJSON(byKey)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"g":55}`, actual)

	fc.RequireEqual(t, nil, byKey.Delete())
	fc.AssertEqual(t, 2, len(app.P))

	n, _ := ReadJSON(`{"p":[{"g":1000}]}`)
	err = b.Root().UpsertFrom(n)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 3, len(app.P))

	app = mapAsListApp{}
	b = node.NewBrowser(m, &Node{Object: &app})
	n, _ = ReadJSON(`{"p":[{"g":100}]}`)
	err = b.Root().UpsertFrom(n)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 1, len(app.P))
}

type MapAsListP struct {
	G int
}

type mapAsListApp struct {
	P map[int]*MapAsListP
}
