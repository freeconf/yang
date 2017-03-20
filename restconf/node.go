package restconf

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/stock"
)

func Node(mgmt *Management) node.Node {
	if mgmt.DeviceHandler == nil {
		mgmt.DeviceHandler = NewDeviceHandler()
	}
	return &node.Extend{
		Node: node.ReflectNode(mgmt.DeviceHandler),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "web":
				if r.New {
					mgmt.Web = stock.NewHttpServer(mgmt.DeviceHandler)
				}
				if mgmt.Web != nil {
					return stock.WebServerNode(mgmt.Web), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
	}
}
