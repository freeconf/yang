package auth

import (
	"context"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type Auth interface {
	ConstrainRoot(role string, c *node.Constraints)
}

// This does not implement NETMOD ACLs, but rather a simplistic implementation
// to be both useful and example of more complex implementations
type Rbac struct {
	Roles map[string]*Role
}

func NewRbac() *Rbac {
	return &Rbac{
		Roles: make(map[string]*Role),
	}
}

func (self *Rbac) ConstrainRoot(role string, c *node.Constraints) {
	r, found := self.Roles[role]
	if !found {
		r = noAccess
	}
	c.AddConstraint("auth", 0, 0, r)
}

var noAccess = &Role{}

type Role struct {
	Id     string
	Access map[string]*AccessControl
}

func NewRole() *Role {
	return &Role{
		Access: make(map[string]*AccessControl),
	}
}

type AccessControl struct {
	Path        string
	Permissions Permission
}

type Permission int

const (
	None Permission = iota
	Read
	Full
)

var UnauthorizedError = c2.NewErrC("unauthorized", 401)

func (self *Role) CheckListPreConstraints(r *node.ListRequest) (bool, error) {
	requested := Read
	if r.New {
		requested = Full
	}
	return self.check(r.Meta, r.Selection.Context, requested)
}

func (self *Role) CheckContainerPreConstraints(r *node.ChildRequest) (bool, error) {
	requested := Read
	if r.New {
		requested = Full
	}
	return self.check(r.Meta, r.Selection.Context, requested)
}

func (self *Role) CheckFieldPreConstraints(r *node.FieldRequest, hnd *node.ValueHandle) (bool, error) {
	requested := Read
	if r.Write {
		requested = Full
	}
	return self.check(r.Meta, r.Selection.Context, requested)
}

func (self *Role) CheckNotifyFilterConstraints(msg node.Selection) (bool, error) {
	return self.check(msg.Meta(), msg.Context, Full)
}

type contextKey int

var permKey contextKey = 0

func (self *Role) CheckActionPreConstraints(r *node.ActionRequest) (bool, error) {
	return self.check(r.Meta, r.Selection.Context, Full)
}

func (self *Role) ContextConstraint(s node.Selection) context.Context {
	if acl, found := self.Access[meta.GetPath(s.Meta())]; found {
		return context.WithValue(s.Context, permKey, acl.Permissions)
	}
	return s.Context
}

func (self *Role) check(m meta.Meta, c context.Context, requested Permission) (bool, error) {
	allowed := None
	path := meta.GetPath(m)
	if acl, found := self.Access[path]; found {
		allowed = acl.Permissions
	} else if x := c.Value(permKey); x != nil {
		allowed = x.(Permission)
	}
	if requested == Read {
		return allowed >= Read, nil
	}
	if allowed >= requested {
		return true, nil
	}
	return false, UnauthorizedError
}
