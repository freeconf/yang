package device

import (
	"github.com/c2stack/c2g/node"
)

type ProxyContextKey int

const RemoteIpAddressKey ProxyContextKey = 0

func MapNode(mgr *Map, server Server, client Client) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "device":
				return deviceRecordListNode(mgr.devices, server), nil
			}
			return nil, nil
		},
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "register":
				reg, err := registrationRequest(r.Input)
				if err != nil {
					return nil, err
				}
				if reg.Address == "" {
					ctx := r.Selection.Context
					reg.Address = ctx.Value(RemoteIpAddressKey).(string)
				}
				d, err := client.NewDevice(reg.Address)
				if err != nil {
					return nil, err
				}
				mgr.Add(reg.Id, d)
				return nil, nil
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "update":
				sub := mgr.OnUpdate(func(d Device, id string, c Change) {
					n := deviceChangeNode(id, d, server, c)
					r.Send(n)
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

func deviceChangeNode(id string, d Device, server Server, c Change) node.Node {
	return &node.Extend{
		Node: deviceNode(id, d, server),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "change":
				hnd.Val = node.NewEnumValue(r.Meta.GetDataType(), int(c))
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func deviceRecordListNode(devices map[string]Device, server Server) node.Node {
	index := node.NewIndex(devices)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var d Device
			var id string
			key := r.Key
			if key != nil {
				id = key[0].Str
				d = devices[id]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					if id = v.String(); id != "" {
						if d = devices[id]; d != nil {
							key = node.SetValues(r.Meta.KeyMeta(), id)
						}
					}
				}
			}
			if d != nil {
				return deviceNode(id, d, server), key, nil
			}
			return nil, nil, nil
		},
	}
}

func deviceNode(id string, d Device, server Server) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "module":
				return YangLibModuleList(d.Modules(), d.SchemaSource()), nil
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "id":
				hnd.Val = &node.Value{Str: id}
			case "address":
				hnd.Val = &node.Value{Str: server.DeviceAddress(id, d)}
			}
			return nil
		},
	}
}

type RegistrationRequest struct {
	Address string
	Port    string
	Id      string
}

func registrationRequest(s node.Selection) (RegistrationRequest, error) {
	var reg RegistrationRequest
	if err := s.InsertInto(node.ReflectNode(&reg)).LastErr; err != nil {
		return reg, err
	}
	return reg, nil
}
