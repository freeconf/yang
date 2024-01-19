package nodeutil

import (
	"errors"
	"strings"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
)

var testActionMstr = `module x {
	rpc noArgs {}

	rpc withArgs {
		input {
			leaf bar {
				type string;
			}
		}
		output {
			leaf foo {
				type string;
			}
		}
	}

	rpc withArgsStruct {
		input {
			leaf bar {
				type string;
			}
		}
		output {
			leaf foo {
				type string;
			}
		}
	}

	rpc withArgsAndErr {
		input {
			leaf bar {
				type string;
			}
		}
		output {
			leaf foo {
				type string;
			}
		}
	}

	rpc exploded {
		input {
			leaf bar {
				type string;
			}
		}
		output {
			leaf foo {
				type string;
			}
		}
	}

	rpc explodedWithErr {
		input {
			leaf bar {
				type string;
			}
		}
		output {
			leaf foo {
				type string;
			}
		}
	}
}`

func TestAction(t *testing.T) {
	app := &testActionApp{}
	n := &Node{
		Object: app,
		OnOptions: func(n *Node, m meta.Definition, o NodeOptions) NodeOptions {
			if strings.HasPrefix(m.Ident(), "exploded") {
				o.ActionInputExploded = true
				o.ActionOutputExploded = true
			}
			return o
		},
	}
	m, err := parser.LoadModuleFromString(nil, testActionMstr)
	b := node.NewBrowser(m, n)
	fc.RequireEqual(t, nil, err)
	t.Run("noArgs", func(t *testing.T) {
		_, err := sel(b.Root().Find("noArgs")).Action(nil)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, 1, app.noArgsCalled)
	})

	t.Run("withArgs", func(t *testing.T) {
		in, _ := ReadJSON(`{"bar":"x"}`)
		resp, err := sel(b.Root().Find("withArgs")).Action(in)
		fc.AssertEqual(t, nil, err)
		actual, err := WriteJSON(resp)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, `{"foo":"x"}`, actual)
	})

	t.Run("withArgsStruct", func(t *testing.T) {
		in, _ := ReadJSON(`{"bar":"x"}`)
		resp := sel(sel(b.Root().Find("withArgsStruct")).Action(in))
		actual, err := WriteJSON(resp)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, `{"foo":"x"}`, actual)
	})

	t.Run("withArgsAndErr", func(t *testing.T) {
		in, _ := ReadJSON(`{"bar":"x"}`)
		resp := sel(sel(b.Root().Find("withArgsAndErr")).Action(in))
		actual, err := WriteJSON(resp)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, `{"foo":"x"}`, actual)

		in, _ = ReadJSON(`{"bar":"bad-input"}`)
		_, errExpected := sel(b.Root().Find("withArgsAndErr")).Action(in)
		fc.AssertEqual(t, "bad-input", errExpected.Error())
	})

	t.Run("exploded", func(t *testing.T) {
		in, _ := ReadJSON(`{"bar":"x"}`)
		resp := sel(sel(b.Root().Find("exploded")).Action(in))
		actual, err := WriteJSON(resp)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, `{"foo":"x"}`, actual)
	})

	t.Run("explodedWithErr", func(t *testing.T) {
		in, _ := ReadJSON(`{"bar":"x"}`)
		resp := sel(sel(b.Root().Find("explodedWithErr")).Action(in))
		actual, err := WriteJSON(resp)
		fc.AssertEqual(t, nil, err)
		fc.AssertEqual(t, `{"foo":"x"}`, actual)

		in, _ = ReadJSON(`{"bar":"bad-input"}`)
		_, errExpected := sel(b.Root().Find("explodedWithErr")).Action(in)
		fc.AssertEqual(t, "bad-input", errExpected.Error())
	})
}

func sel(s *node.Selection, err error) *node.Selection {
	if err != nil {
		panic(err.Error())
	}
	if s == nil {
		panic("no selection")
	}
	return s
}

type testActionApp struct {
	noArgsCalled int
}

func (app *testActionApp) NoArgs() {
	app.noArgsCalled += 1
}

type BasicOutput struct {
	Foo string
}

type BasicInput struct {
	Bar string
}

func (app *testActionApp) WithArgs(in *BasicInput) *BasicOutput {
	return &BasicOutput{Foo: in.Bar}
}

func (app *testActionApp) WithArgsStruct(in BasicInput) BasicOutput {
	return BasicOutput{Foo: in.Bar}
}

func (app *testActionApp) WithArgsAndErr(in *BasicInput) (*BasicOutput, error) {
	if in.Bar == "bad-input" {
		return nil, errors.New(in.Bar)
	}
	return &BasicOutput{Foo: in.Bar}, nil
}

func (app *testActionApp) Exploded(foo string) string {
	return foo
}

func (app *testActionApp) ExplodedWithErr(foo string) (string, error) {
	if foo == "bad-input" {
		return "", errors.New(foo)
	}
	return foo, nil
}
