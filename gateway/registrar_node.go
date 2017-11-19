package gateway

import (
	"strings"

	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/val"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func RegistrarNode(registrar Registrar) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "registrations":
				return registrationsNode(registrar), nil
			}
			return nil, nil
		},
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "register":
				var reg Registration
				if err := r.Input.InsertInto(regNode(&reg)).LastErr; err != nil {
					return nil, err
				}
				ctx := r.Selection.Context
				if regAddr, hasRegAddr := ctx.Value(device.RemoteIpAddressKey).(string); hasRegAddr {
					reg.Address = strings.Replace(reg.Address, "{REQUEST_ADDRESS}", regAddr, 1)
				}
				registrar.RegisterDevice(reg.DeviceId, reg.Address)
				return nil, nil
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			switch r.Meta.Ident() {
			case "update":
				sub := registrar.OnRegister(func(reg Registration) {
					r.Send(regNode(&reg))
				})
				return sub.Close, nil
			}
			return nil, nil
		},
	}
}

func registrationsNode(registrar Registrar) node.Node {

	// assume local registrar, need better way to iterate
	index := node.NewIndex(registrar.(*LocalRegistrar).regs)

	return &nodes.Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var reg Registration
			found := false
			var id string
			key := r.Key
			if key != nil {
				id = key[0].String()
				reg, found = registrar.LookupRegistration(id)
			} else if r.Row < registrar.RegistrationCount() {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					id = v.String()
					reg, found = registrar.LookupRegistration(id)
					key = []val.Value{val.String(reg.DeviceId)}
				}
			}
			if found {
				return nodes.ReflectChild(&reg), key, nil
			}
			return nil, nil, nil
		},
	}
}

func regNode(reg *Registration) node.Node {
	return nodes.ReflectChild(reg)
}
