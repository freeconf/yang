// +build ignore

package store

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func RegistrarNode(registrar *Registrar) node.Node {
	return &node.Extend{
		Node: node.Reflect(registrar),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "endpoint":
				if registrar.Endpoints != nil {
					return EndpointsNode(registrar), nil
				}
				return nil, nil
			}
			return p.Child(r)
		},
		OnAction: func(p node.Node, r node.ActionRequest) (output node.Node, err error) {
			switch r.Meta.GetIdent() {
			case "register":
				return RegisterEndpointNode(registrar, r.Selection, r.Meta, r.Input)
			}
			return p.Action(r)
		},
	}
}

func EndpointNode(endpoint *Endpoint) node.Node {
	return &node.Extend{
		Node: node.Reflect(endpoint),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			if r.Meta.GetIdent() == endpoint.Meta.GetIdent() {
				return endpoint.handleRequest(node.PathSlice{Head: r.Path, Tail: r.Target})
			}
			return p.Child(r)
		},
		OnChoose: func(p node.Node, sel node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			return choice.GetCase(endpoint.Meta.GetIdent()), nil
		},
		OnAction: func(p node.Node, r node.ActionRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "syncConfig":
				if err := endpoint.pushConfig(); err != nil {
					return nil, err
				}
				return nil, endpoint.pullConfig()
			case "pushConfig":
				return nil, endpoint.pushConfig()
			case "pullConfig":
				return nil, endpoint.pullConfig()
			}
			return p.Action(r)
		},
	}
}

func EndpointsNode(registrar *Registrar) node.Node {
	n := &node.MarshalMap{
		Map: registrar.Endpoints,
		OnSelectItem: func(item interface{}) node.Node {
			return EndpointNode(item.(*Endpoint))
		},
	}
	return n.Node()
}

func RegisterEndpointNode(registrar *Registrar, sel node.Selection, rpc *meta.Rpc, input node.Selection) (output node.Node, err error) {
	reg := &Endpoint{
		YangPath:     registrar.YangPath,
		ClientSource: registrar.ClientSource,
	}
	regNode := node.Reflect(reg)
	if err = input.UpsertInto(regNode).LastErr; err != nil {
		return nil, err
	}
	if err = registrar.RegisterEndpoint(reg); err != nil {
		return nil, err
	}
	return regNode, nil
}
