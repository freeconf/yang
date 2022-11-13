package nodeutil_test

import (
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/val"
)

func ExampleExtend() {
	model := `
		leaf bar {
			type string;
		}
		leaf bleep {
			type string;
		}
	`
	f := foo{
		Bar: "x",
	}
	bleep := "y"
	data := &nodeutil.Extend{
		Base: nodeutil.ReflectChild(&f),
		OnField: func(parent node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "bleep":
				if r.Write {
					bleep = hnd.Val.String()
				} else {
					hnd.Val = val.String(bleep)
				}
			default:
				return parent.Field(r, hnd)
			}
			return nil
		},
	}

	sel := exampleSelection(model, data)

	examplePrint(sel)
	// Output:
	// {"x:bar":"x","x:bleep":"y"}
}
