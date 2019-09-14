package secure

import (
	"reflect"

	"github.com/freeconf/yang/val"

	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
)

func Manage(rbac *Rbac) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "authentication":

			case "authorization":
				return authorizeMgmt(rbac), nil
			}
			return nil, nil
		},
	}
}

func authorizeMgmt(rbac *Rbac) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "role":
				return rolesMgmt(rbac.Roles), nil
			}
			return nil, nil
		},
	}
}

func rolesMgmt(role map[string]*Role) node.Node {
	return nodes.Reflect{
		OnChild: func(p nodes.Reflect, v reflect.Value) node.Node {
			switch x := v.Interface().(type) {
			case *AccessControl:
				return accessControlMgmt(x)
			}
			return p.Child(v)
		},
	}.List(role)
}

func accessControlMgmt(ac *AccessControl) node.Node {
	return &nodes.Extend{
		Base: nodes.ReflectChild(ac),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "perm":
				if r.Write {
					ac.Permissions = Permission(hnd.Val.Value().(val.Enum).Id)
				} else {
					var err error
					hnd.Val, err = node.NewValue(r.Meta.Type(), ac.Permissions)
					if err != nil {
						return err
					}
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
	}
}
