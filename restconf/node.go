package restconf

import (
	"github.com/freeconf/yang/c2"
	"github.com/freeconf/yang/device"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
	"github.com/freeconf/yang/stock"
	"github.com/freeconf/yang/val"
)

func Node(mgmt *Server, ypath meta.StreamSource) node.Node {
	return &nodes.Extend{
		Base: nodes.ReflectChild(mgmt),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "web":
				if r.New {
					mgmt.Web = stock.NewHttpServer(mgmt)
				}
				if mgmt.Web != nil {
					return stock.WebServerNode(mgmt.Web), nil
				}
			case "callHome":
				if r.New {
					rc := ProtocolHandler(ypath)
					mgmt.CallHome = device.NewCallHome(rc)
				}
				if mgmt.CallHome != nil {
					return device.CallHomeNode(mgmt.CallHome), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "debug":
				if r.Write {
					c2.DebugLog(hnd.Val.Value().(bool))
				} else {
					hnd.Val = val.Bool(c2.DebugLogEnabled())
				}
			case "streamCount":
				hnd.Val = val.Int32(mgmt.notifiers.Len())
			case "subscriptionCount":
				hnd.Val = val.Int32(mgmt.SubscriptionCount())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
