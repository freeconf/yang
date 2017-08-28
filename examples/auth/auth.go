package auth

import (
	"regexp"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/node"
)

// Role-based Access Control
//   see https://en.wikipedia.org/wiki/Role-based_access_control
type Rbac struct {
	Roles map[string]*Role
}

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
	pathRegx    *regexp.Regexp
}

func (self *AccessControl) Matches(targetPath string) (bool, error) {
	if self.pathRegx == nil {
		var err error
		if self.pathRegx, err = regexp.Compile(self.Path); err != nil {
			return false, err
		}
	}
	return self.pathRegx.MatchString(targetPath), nil
}

type Permission string

const (
	Read = Permission("read")
	Full = Permission("full")
	None = Permission("none")
)

func (allowed Permission) Allows(requested Permission) bool {
	switch allowed {
	case Read:
		return requested != Full
	case Full:
		return true
	}
	return false
}

var UnauthorizedError = c2.NewErrC("Unauthorized", 401)

func (self *Role) CheckListPreConstraints(r *node.ListRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	requested := Read
	if r.New {
		requested = Full
	}
	if r.Key != nil {
		return self.check(r.Selection.Path.StringNoModule()+"="+node.EncodeKey(r.Key), requested)
	}
	return true, nil
}

func (self *Role) CheckContainerPreConstraints(r *node.ChildRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	requested := Read
	if r.New {
		requested = Full
	}
	return self.check(r.Selection.Path.StringNoModule()+"/"+r.Meta.GetIdent(), requested)
}

func (self *Role) CheckFieldPreConstraints(r *node.FieldRequest, hnd *node.ValueHandle) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	requested := Read
	if r.Write {
		requested = Full
	}
	return self.check(r.Selection.Path.StringNoModule()+"/"+r.Meta.GetIdent(), requested)
}

func (self *Role) CheckActionPreConstraints(r *node.ActionRequest) (bool, error) {
	return self.check(r.Selection.Path.StringNoModule(), Full)
}

func (self *Role) check(targetPath string, requested Permission) (bool, error) {
	// HACK: occasional leading path messing things up, find out why this is inconsistent
	targetPath = strings.TrimLeft(targetPath, "/")

	allowed := None
	for _, ac := range self.Access {
		if found, err := ac.Matches(targetPath); found {
			allowed = ac.Permissions
		} else if err != nil {
			return false, err
		}
	}
	return allowed.Allows(requested), UnauthorizedError
}

type NoAccess struct {
}

func (self NoAccess) CheckListPreConstraints(r *node.ListRequest) (bool, error) {
	return false, UnauthorizedError
}

func (self NoAccess) CheckContainerPreConstraints(r *node.ChildRequest) (bool, error) {
	return false, UnauthorizedError
}

func (self NoAccess) CheckFieldPreConstraints(r *node.FieldRequest, hnd *node.ValueHandle) (bool, error) {
	return false, UnauthorizedError
}
