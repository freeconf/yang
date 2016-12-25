package restconf

import (
	"time"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func ServiceNode(mgmt *Management) node.Node {
	options := mgmt.Options()
	return &node.Extend{
		Node: node.ReflectNode(&options),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "callHome":
				if r.New {
					mgmt.CallHome = &CallHome{
						Module:       mgmt.Handler.Root.Meta.(*meta.Module),
						ClientSource: mgmt.Web,
					}
				}
				if mgmt.CallHome != nil {
					return CallHomeNode(mgmt.CallHome), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		// OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
		// 	switch r.Meta.GetIdent() {
		// 	case "notifyKeepaliveTimeoutMs", "path":
		// 		if r.Write {
		// 			return node.WriteField(r.Meta, &options, hnd.Val)
		// 		}
		// 		var err error
		// 		if hnd.Val, err = node.ReadField(r.Meta, &options); err != nil {
		// 			return err
		// 		}
		// 	default:
		// 		return p.Field(r, hnd)
		// 	}
		// 	return nil
		// },
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			mgmt.ApplyOptions(options)
			return nil
		},
	}
}

func CallHomeNode(ch *CallHome) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(ch),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "registration":
				if ch.Registration != nil {
					return node.ReflectNode(ch.Registration), nil
				}
			}
			return nil, nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			// We wait for 1 second because on initial configuration load the
			// callback url isn't valid until the web server is also configured.
			time.AfterFunc(1*time.Second, func() {
				if err := ch.StartRegistration(); err != nil {
					c2.Err.Printf("Initial registration failed %s", err)
				}
			})
			return p.EndEdit(r)
		},
	}
}
