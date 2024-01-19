package nodeutil

import (
	"errors"
	"reflect"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

var testStructMstr = `module x {
	container c {
		leaf z {
			type string;
		}
	}
	leaf l {
		type string;
	}
	list p {
		key g;
		leaf g {
			type int32;
		}
	}
	leaf q {
		type int32;
	}
	leaf g {
		type string;
	}
	leaf f {
		type string;
	}
	container z {
		leaf zz {
			type string;
		}
	}
}`

func TestStructAsContainer(t *testing.T) {
	app := &reflectStructTestApp{
		C: &reflectStructC{
			Z: "hi",
		},
		L: "whatever",
	}
	m, err := parser.LoadModuleFromString(nil, testStructMstr)
	fc.RequireEqual(t, nil, err)
	ref := &Node{}
	x := newStructAsContainer(ref, reflect.ValueOf(app))
	t.Run("child container", func(t *testing.T) {
		c := meta.Find(m, "c").(meta.HasDataDefinitions)
		n, err := x.newChild(c)
		fc.RequireEqual(t, nil, err)
		fc.AssertEqual(t, false, n.IsNil())
		newC, valid := n.Interface().(*reflectStructC)
		fc.AssertEqual(t, true, valid)
		fc.AssertEqual(t, "", newC.Z)

		h, err := x.getHandler(c)
		fc.RequireEqual(t, nil, err)
		cv, err := h.get()
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, app.C, cv.Interface())
	})

	t.Run("leaf", func(t *testing.T) {
		l := meta.Find(m, "l")
		h, err := x.getHandler(l)
		fc.RequireEqual(t, nil, err)
		lv, err := h.get()
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, "whatever", lv.String())
		h.set(reflect.ValueOf("zzz"))
		fc.AssertEqual(t, "zzz", app.L)
		h.clear()
		fc.AssertEqual(t, "", app.L)
	})

	t.Run("child list", func(t *testing.T) {
		p := meta.Find(m, "p").(meta.HasDataDefinitions)
		n, err := x.newChild(p)
		fc.RequireEqual(t, nil, err)
		_, valid := n.Interface().(map[int]*reflectStructP)
		fc.RequireEqual(t, true, valid)
	})
}

func TestStructAsContainer2(t *testing.T) {
	app := &reflectStructTestApp{}
	m, err := parser.LoadModuleFromString(nil, testStructMstr)
	fc.RequireEqual(t, nil, err)
	b := node.NewBrowser(m, &Node{Object: app})

	t.Run("replace container", func(t *testing.T) {
		n, _ := ReadJSON(`{"zz":"bye"}`)
		sel(b.Root().Find("z")).UpsertFrom(n)
		fc.AssertEqual(t, "bye", app.Z.Zz)
	})

}

func TestFindReflectByMethod(t *testing.T) {
	m, err := parser.LoadModuleFromString(nil, testStructMstr)
	fc.RequireEqual(t, nil, err)

	app := &reflectStructTestApp{Q: 100}
	appVal := reflect.ValueOf(app)
	var opts NodeOptions

	qMeta := meta.Find(m, "q")
	q := findReflectByField(appVal, qMeta, opts)
	fc.AssertEqual(t, true, q != nil)
	fc.AssertEqual(t, "Q", q.f.Name)
	fc.AssertEqual(t, "GetQ", q.getter.Name)
	fc.AssertEqual(t, "", q.setter.Name)
	qv, err := q.get()
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, int64(100), qv.Int()) // choose getter

	gMeta := meta.Find(m, "g")
	g := findReflectByField(appVal, gMeta, opts)
	fc.AssertEqual(t, true, g != nil)
	fc.AssertEqual(t, "", g.f.Name)
	fc.AssertEqual(t, "G", g.getter.Name)
	fc.AssertEqual(t, "", g.setter.Name)
	gv, err := g.get()
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, "Gee", gv.String())

	fMeta := meta.Find(m, "f")
	f := findReflectByField(appVal, fMeta, opts)
	fc.AssertEqual(t, true, f != nil)
	fc.AssertEqual(t, "F", f.getter.Name)
	fv, err := f.get()
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, "Eff", fv.String())
	fc.AssertEqual(t, "SetF", f.setter.Name)
	fc.AssertEqual(t, nil, f.set(reflect.ValueOf("Heffe")))
	fc.AssertEqual(t, "Heffe", app.f)

	fAsG := findReflectByField(appVal, fMeta, NodeOptions{Ident: "G"})
	fc.AssertEqual(t, true, fAsG != nil)
	fc.AssertEqual(t, "G", fAsG.getter.Name)

	gWithErr := findReflectByField(appVal, gMeta, NodeOptions{
		GetterPrefix: "GetAlwaysError",
		SetterPrefix: "SetAlwaysError",
	})
	fc.AssertEqual(t, true, gWithErr != nil)
	fc.AssertEqual(t, "GetAlwaysErrorG", gWithErr.getter.Name)
	fc.AssertEqual(t, "SetAlwaysErrorG", gWithErr.setter.Name)
	_, err = gWithErr.get()
	fc.RequireEqual(t, false, err == nil)
	fc.AssertEqual(t, "expected get error", err.Error())
	err = gWithErr.set(reflect.ValueOf("doesn't matter"))
	fc.RequireEqual(t, false, err == nil)
	fc.AssertEqual(t, "expected set error", err.Error())

	fc.AssertEqual(t, reflect.String, gWithErr.fieldType().Kind())
}

type reflectStructC struct {
	Z string
}

type reflectStructP struct {
	G int
}

type reflectStructTestApp struct {
	C *reflectStructC
	L string
	P map[int]*reflectStructP
	Q int
	f string
	Z relectStructZ
}

type relectStructZ struct {
	Zz string
}

func (a *reflectStructTestApp) GetQ() int {
	return 10
}

func (a *reflectStructTestApp) G() string {
	return "Gee"
}

func (a *reflectStructTestApp) F() (string, error) {
	if a.f != "" {
		return a.f, nil
	}
	return "Eff", nil
}

func (a *reflectStructTestApp) SetF(v string) error {
	a.f = v
	return nil
}

func (a *reflectStructTestApp) GetAlwaysErrorG() (string, error) {
	return "doesn't matter", errors.New("expected get error")
}

func (a *reflectStructTestApp) SetAlwaysErrorG(string) error {
	return errors.New("expected set error")
}
