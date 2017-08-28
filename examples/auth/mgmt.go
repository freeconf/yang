package auth

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func Manage(rbac *Rbac) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "role":
				return roleMgmt(rbac.Roles), nil
			}
			return nil, nil
		},
	}
}

func roleMgmt(role map[string]*Role) node.Node {
	return &nodes.Extend{
		Base: nodes.Reflect(role),
	}
}
