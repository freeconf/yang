package nodeutil

import (
	"reflect"
	"strings"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
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
	app := &reflectTestApp{
		C: &reflectTestC{
			L: "hi",
		},
	}
	b := node.NewBrowser(m, &Node{Object: app})
	actual, err := WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"c":{"l":"hi"}}`, actual)

	app = &reflectTestApp{}
	b = node.NewBrowser(m, &Node{Object: app})
	err = b.Root().UpsertFrom(ReadJSON(`{"c":{"l":"hi"}}`))
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, true, app.C != nil)
	fc.AssertEqual(t, "hi", app.C.L)

	app = &reflectTestApp{
		C: &reflectTestC{
			L: "hi",
		},
	}
	n := &Node{
		Object: app,
		OnRead: func(ref *Node, m meta.Definition, t reflect.Type, v reflect.Value) (reflect.Value, error) {
			if t.Kind() == reflect.String {
				return reflect.ValueOf(strings.ToUpper(v.String())), nil
			}
			return v, nil
		},
		OnWrite: func(ref *Node, m meta.Definition, t reflect.Type, v reflect.Value) (reflect.Value, error) {
			if t.Kind() == reflect.String {
				return reflect.ValueOf(strings.ToLower(v.String())), nil
			}
			return v, nil
		},
	}
	b = node.NewBrowser(m, n)
	actual, err = WriteJSON(b.Root())
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, `{"c":{"l":"HI"}}`, actual)

	err = b.Root().UpsertFrom(ReadJSON(`{"c":{"l":"BYE"}}`))
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, true, app.C != nil)
	fc.AssertEqual(t, "bye", app.C.L)
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
