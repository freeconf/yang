package nodeutil

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

func TestSliceAsList(t *testing.T) {
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
	app := &sliceAsListApp{
		P: []*sliceAsListP{
			{G: 100},
			{G: 55},
			{G: 999},
		},
	}
	b := node.NewBrowser(m, &Node{Object: app})
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"p":[{"g":100},{"g":55},{"g":999}]}`, actual)

	byKey, err := b.Root().Find("p=55")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, byKey != nil)
	actual, err = WriteJSON(byKey)
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"g":55}`, actual)

	fc.RequireEqual(t, nil, byKey.Delete())
	fc.AssertEqual(t, 2, len(app.P))

	err = b.Root().UpsertFrom(ReadJSON(`{"p":[{"g":1000}]}`))
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 3, len(app.P))

	app = &sliceAsListApp{}
	b = node.NewBrowser(m, &Node{Object: app})
	err = b.Root().UpsertFrom(ReadJSON(`{"p":[{"g":100}]}`))
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 1, len(app.P))
}

type sliceAsListP struct {
	G int
}

type sliceAsListApp struct {
	P []*sliceAsListP
}
