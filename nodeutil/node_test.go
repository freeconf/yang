package nodeutil

import (
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

func TestReflectBasics(t *testing.T) {
	mstr := `module x {
		container c {
			leaf l {
				type string;
			}	
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	fc.RequireEqual(t, nil, err)
	app := reflectTestApp{
		C: &reflectTestC{
			L: "hi",
		},
	}
	b := node.NewBrowser(m, &Node{Object: &app})
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"c":{"l":"hi"}}`, actual)

	app = reflectTestApp{}
	b = node.NewBrowser(m, &Node{Object: &app})
	err = b.Root().UpsertFrom(ReadJSON(`{"c":{"l":"hi"}}`))
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, true, app.C != nil)
	fc.AssertEqual(t, "hi", app.C.L)
}

func TestReflect(t *testing.T) {
	mstr := `module x {
		container c {
			leaf l {
				type string;
			}	
		}
		container z {
			leaf l {
				type string;
			}
		}
	}`
	m, err := parser.LoadModuleFromString(nil, mstr)
	fc.RequireEqual(t, nil, err, mstr)

	readFrom := &reflectTestApp{
		C: &reflectTestC{
			L: "hi",
		},
		Z: map[string]any{
			"l": "bye",
		},
	}

	expected := `{"c":{"l":"hi"},"z":{"l":"bye"}}`

	// read
	b := node.NewBrowser(m, &Node{Object: readFrom})
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, expected, actual)

	// write
	wtrInto := &reflectTestApp{}
	bwtr := node.NewBrowser(m, &Node{Object: wtrInto})
	err = bwtr.Root().UpsertFrom(ReadJSON(expected))
	fc.RequireEqual(t, nil, err)
}

type reflectTestC struct {
	L string
}

type reflectTestP struct {
	G int
}

type reflectTestApp struct {
	C *reflectTestC
	Z map[string]any
	P map[int]*reflectTestP
}
