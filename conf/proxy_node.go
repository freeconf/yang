package conf

import (
	"github.com/c2stack/c2g/node"
)

type ProxyContextKey int

const RemoteIpAddressKey ProxyContextKey = 0

func ProxyNode(proxy *Proxy) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "device":
				return proxyMountListNode(proxy.mounts), nil
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
				return nil, proxy.Mount(reg.Id, reg.Address, reg.Port)
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.GetIdent() {
			case "deviceUpdate":
				sub := proxy.OnUpdate(func(m *Mount) {
					payload := map[string]interface{}{
						"device": m.DeviceId,
						"change": "added",
					}
					r.Send(node.MapNode(payload))
				})
				return sub.Close, nil
			case "moduleUpdate":
				sub := proxy.OnModuleUpdate(true, func(module string, m *Mount) {
					payload := map[string]interface{}{
						"device": m.DeviceId,
						"module": module,
						"change": "added",
					}
					r.Send(node.MapNode(payload))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

func proxyMountListNode(mounts map[string]*Mount) node.Node {
	index := node.NewIndex(mounts)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var m *Mount
			var id string
			key := r.Key
			if key != nil {
				id = key[0].Str
				m = mounts[id]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					if id = v.String(); id != "" {
						if m = mounts[id]; m != nil {
							key = node.SetValues(r.Meta.KeyMeta(), id)
						}
					}
				}
			}
			if m != nil {
				return proxyMountNode(m), key, nil
			}
			return nil, nil, nil
		},
	}
}

func proxyMountNode(m *Mount) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(m),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "module":
				hnds, err := m.Device.ModuleHandles()
				if err != nil {
					return nil, err
				} else if hnds != nil {
					return proxyModuleHandles(hnds), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "id":
				hnd.Val = &node.Value{Str: m.DeviceId}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}

func proxyModuleHandles(hnds map[string]*ModuleHandle) node.Node {
	index := node.NewIndex(hnds)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var hnd *ModuleHandle
			var name string
			key := r.Key
			if key != nil {
				name = key[0].Str
				hnd = hnds[name]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					if name = v.String(); name != "" {
						if hnd = hnds[name]; hnd != nil {
							key = node.SetValues(r.Meta.KeyMeta(), name)
						}
					}
				}
			}
			if hnd != nil {
				return node.ReflectNode(hnd), key, nil
			}
			return nil, nil, nil
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
