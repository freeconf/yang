package nodeutil_test

import (
	"fmt"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/val"
)

// Example defines a model with a single string called "message" as it's only allowed
// field (or 'leaf').
//
// Models can be matched to a data for lot's of things including reading data, so we create a simple
// data source that always returns hello.
//
// Models and Data come together as a browser.  A browser is all you need to do anything with the data
// that confirms to the model.
func ExampleBasic_onField() {
	model := `
	  leaf foo {
		  type string;
	  }`

	data := &nodeutil.Basic{

		// Custom implementations of reading and writing fields called "leafs" or
		// "leaf-lists" in YANG.
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "foo":
				if r.Write {
					fmt.Println(hnd.Val.String())
				} else {
					hnd.Val = val.String("READ")
				}
			}
			return nil
		},
	}
	sel := exampleSelection(model, data)

	examplePrint(sel)

	n, _ := nodeutil.ReadJSON(`{"foo":"WRITE"}`)
	sel.UpsertFrom(n)
	// Output:
	// {"foo":"READ"}
	// WRITE
}

type foo struct {
	Bar string
}

// TestReadingStruct expands on TestSimplestExample by wrapping a 'container' around the
// message.  Containers are like a Golang struct.
func ExampleBasic_onChild() {
	model := `
		container foo {
			leaf bar {
				type string;
			}
		}
		`

	f := &foo{
		Bar: "x",
	}
	data := &nodeutil.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "foo":
				if r.New {
					f = &foo{}
				} else if r.Delete {
					f = nil
				}
				if f != nil {
					return nodeutil.ReflectChild(f), nil
				}
			}
			return nil, nil
		},
	}

	sel := exampleSelection(model, data)

	fmt.Println("Reading")
	examplePrint(sel)

	fmt.Println("Deleting")
	selFoo, err := sel.Find("foo")
	if err != nil {
		panic(err)
	}
	if err = selFoo.Delete(); err != nil {
		panic(err)
	}
	examplePrint(sel)

	fmt.Println("Creating")
	n, _ := nodeutil.ReadJSON(`{"foo":{"bar":"y"}}`)
	sel.InsertFrom(n)
	examplePrint(sel)

	// Output:
	// Reading
	// {"foo":{"bar":"x"}}
	// Deleting
	// {}
	// Creating
	// {"foo":{"bar":"y"}}
}

/*
ExampleBasic_onNext We need to handle adding, removing and naviagtion thru a list both by key
and sequentially.  Because most lists have a key, Go's map is often most
useful structure to store lists. node.Index helps navigating thru a map sequentally but
you can use your own method.
*/
func ExampleBasic_onNext() {
	model := `
			list foo {
			   key "bar";
			   leaf bar {
				   type string;
			   }
			}`

	// Data
	m := map[string]*foo{
		"a": {Bar: "a"},
	}
	dataList := func() node.Node {
		// helps navigate a map sequentially
		i := node.NewIndex(m)
		return &nodeutil.Basic{
			OnNextItem: func(r node.ListRequest) nodeutil.BasicNextItem {
				var f *foo
				return nodeutil.BasicNextItem{
					New: func() error {
						f = &foo{}
						m[r.Key[0].String()] = f
						return nil
					},

					GetByKey: func() error {
						f = m[r.Key[0].String()]
						return nil
					},

					GetByRow: func() ([]val.Value, error) {
						v := i.NextKey(r.Row)
						if v == node.NO_VALUE {
							return nil, nil
						}
						id := v.String()
						f = m[id]
						return []val.Value{val.String(id)}, nil
					},

					DeleteByKey: func() error {
						delete(m, r.Key[0].String())
						return nil
					},

					Node: func() (node.Node, error) {
						if f != nil {
							return nodeutil.ReflectChild(f), nil
						}
						return nil, nil
					},
				}
			},
		}
	}
	sel := exampleSelection(model, &nodeutil.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			return dataList(), nil
		},
	})

	fmt.Println("Reading")
	examplePrint(sel)

	fmt.Println("Deleting")
	aSel, err := sel.Find("foo=a")
	if err != nil || aSel == nil {
		panic(err)
	}
	aSel.Delete()
	examplePrint(sel)

	fmt.Println("Creating")
	n, _ := nodeutil.ReadJSON(`{"foo":[{"bar":"b"}]}`)
	err = sel.UpsertFrom(n)
	if err != nil {
		panic(err)
	}
	examplePrint(sel)

	// Output:
	// Reading
	// {"foo":[{"bar":"a"}]}
	// Deleting
	// {"foo":[]}
	// Creating
	// {"foo":[{"bar":"b"}]}
}

/*
OnActions you just decode the input and encode the output and return it as response.
*/
func ExampleBasic_onAction() {

	// YANG 1.1 - 'rpc' is same as 'action' but only used at the top-level of a
	// module for backward compatibility with YANG 1.0
	model := `
			rpc sum {
			   input {
			     leaf a {
			       type int32;
			     }
			     leaf b {
			       type int32;
			     }
			   }
			   output {
			     leaf result {
			       type int32;
			     }
			   }
			}`

	// Data
	data := &nodeutil.Basic{
		OnAction: func(r node.ActionRequest) (out node.Node, err error) {
			switch r.Meta.Ident() {
			case "sum":
				var n nums
				if err := r.Input.InsertInto(nodeutil.ReflectChild(&n)); err != nil {
					return nil, err
				}
				result := map[string]interface{}{
					"result": n.Sum(),
				}
				return nodeutil.ReflectChild(result), nil
			}
			return
		},
	}

	sel, err := exampleSelection(model, data).Find("sum")
	if err != nil {
		panic(err)
	}

	// JSON is a useful format to use as input, but this can come from any node
	// that would return "a" and "b" leaves.
	n, _ := nodeutil.ReadJSON(`{"a":10,"b":32}`)
	result, err := sel.Action(n)
	if err != nil {
		panic(err)
	}
	examplePrint(result)

	// Output:
	// {"result":42}
}

type nums struct {
	A int
	B int
}

func (n nums) Sum() int {
	return n.A + n.B
}

type exampleBox struct {
	message string
}

func examplePrint(sel *node.Selection) {
	s, err := nodeutil.WriteJSON(sel)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func exampleSelection(yangFragment string, n node.Node) *node.Selection {
	mstr := fmt.Sprintf(`module x {
		namespace "";
		prefix "";
		revision 0;

		%s
	}`, yangFragment)
	model, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		panic(err.Error())
	}
	brwsr := node.NewBrowser(model, n)
	return brwsr.Root()
}
