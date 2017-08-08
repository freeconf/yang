package restconf

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/stock"
	"github.com/c2stack/c2g/val"
)

func Node(mgmt *Server, ypath meta.StreamSource) node.Node {
	if mgmt.DeviceHandler == nil {
		mgmt.DeviceHandler = NewDeviceHandler()
	}
	return &nodes.Extend{
		Node: nodes.Reflect(mgmt.DeviceHandler),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "web":
				if r.New {
					mgmt.Web = stock.NewHttpServer(mgmt.DeviceHandler)
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
			switch r.Meta.GetIdent() {
			case "debug":
				if r.Write {
					c2.DebugLog(hnd.Val.Value().(bool))
				} else {
					hnd.Val = val.Bool(c2.DebugLogEnabled())
				}
			case "streamCount":
				hnd.Val = val.Int32(mgmt.DeviceHandler.notifiers.Len())
			case "subscriptionCount":
				hnd.Val = val.Int32(mgmt.DeviceHandler.SubscriptionCount())
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
