package device

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
)

func CallHomeNode(ch *CallHome) node.Node {
	options := ch.Options()
	return &nodes.Extend{
		Node: nodes.ReflectNode(&options),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
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
