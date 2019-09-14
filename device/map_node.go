package device

import (
	"github.com/freeconf/yang/val"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
)

type ProxyContextKey int

const RemoteIpAddressKey ProxyContextKey = 0

type LocalLocationService map[string]string

func MapNode(mgr Map) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "device":
				return deviceRecordListNode(mgr), nil
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.Ident() {
			case "update":
				sub := mgr.OnUpdate(func(d Device, id string, c Change) {
					n := deviceChangeNode(id, d, c)
					r.Send(n)
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

func deviceChangeNode(id string, d Device, c Change) node.Node {
	return &nodes.Extend{
		Base: deviceNode(id, d),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "change":
				var err error
				hnd.Val, err = node.NewValue(r.Meta.Type(), int(c))
				if err != nil {
					return err
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func deviceRecordListNode(devices Map) node.Node {
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var d Device
			var id string
			key := r.Key
			var err error
			if key != nil {
				id = key[0].String()
				d, err = devices.Device(id)
				if err != nil {
					return nil, nil, err
				}
			} else if r.Row < devices.Len() {
				id = devices.NthDeviceId(r.Row)
				d, err = devices.Device(id)
				if err != nil {
					return nil, nil, err
				} else if d != nil {
					key = []val.Value{val.String(id)}
				}
			}
			if d != nil {
				return deviceNode(id, d), key, nil
			}
			return nil, nil, nil
		},
	}
}

func deviceHndNode(hnd *DeviceHnd) node.Node {
	return nodes.ReflectChild(hnd)
}

func deviceNode(id string, d Device) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "module":
				return deviceModuleList(d.Modules()), nil
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "deviceId":
				hnd.Val = val.String(id)
			}
			return nil
		},
	}
}

func deviceModuleList(mods map[string]*meta.Module) node.Node {
	index := node.NewIndex(mods)
	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			key := r.Key
			var m *meta.Module
			if r.Key != nil {
				m = mods[r.Key[0].String()]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					module := v.String()
					if m = mods[module]; m != nil {
						key = []val.Value{val.String(m.Ident())}
					}
				}
			}
			if m != nil {
				return deviceModuleNode(m), key, nil
			}
			return nil, nil, nil
		},
	}
}

func deviceModuleNode(m *meta.Module) node.Node {
	return &nodes.Basic{
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "name":
				hnd.Val = val.String(m.Ident())
			case "revision":
				hnd.Val = val.String(m.Revision().Ident())
			}
			return nil
		},
	}
}
