package device

import (
	"github.com/freeconf/c2g/node"
	"github.com/freeconf/c2g/nodes"
	"github.com/freeconf/c2g/val"
)

func CallHomeNode(ch *CallHome) node.Node {
	options := ch.Options()
	return &nodes.Extend{
		Base: nodes.ReflectChild(&options),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "registered":
				hnd.Val = val.Bool(ch.Registered)
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			return ch.ApplyOptions(options)
		},
	}
}
