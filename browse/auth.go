package browse

import (
	"github.com/c2g/c2"
	"github.com/c2g/node"
	"regexp"
	"container/list"
)

type Acl struct {
	List *list.List
}

func NewAcl() *Acl {
	return &Acl{
		List: list.New(),
	}
}

type AccessControl struct {
	Id          string
	Roles       []string
	Path        string
	Selector    string
	Permissions Permission
	pathRegx    *regexp.Regexp
}

func (self *AccessControl) SetPath(path string) (err error) {
	self.Path = path
	self.pathRegx, err = regexp.Compile(self.Path)
	return
}

type Permission int

const (
	Read Permission = 1 << iota
	Write
	Execute
	Subscribe
)

func decodePermission(s []string) Permission {
	var p Permission
	for _, pstr := range s {
		switch pstr {
		case "r":
			p = p | Read
		case "w":
			p = p | Write
		case "x":
			p = p | Execute
		case "s":
			p = p | Subscribe
		}

	}
	return p
}

var UnauthorizedError = c2.NewErrC("Unauthorized", 401)

func (self *Acl) CheckListPreConstraints(r *node.ListRequest, navigating bool) (bool, error) {
	p := Read
	if r.New {
		p = Write
	}
	if r.Key != nil {
		return self.check(r.Selection.Path().StringNoModule() + "=" + node.EncodeKey(r.Key), p)
	}
	return true, nil
}

func (self *Acl) CheckContainerPreConstraints(r *node.ContainerRequest, navigating bool) (bool, error) {
	p := Read
	if r.New {
		p = Write
	}
	return self.check(r.Selection.Path().StringNoModule() + "/" + r.Meta.GetIdent(), p)
}

func (self *Acl) CheckFieldPreConstraints(r *node.FieldRequest, write bool, navigating bool) (bool, error) {
	p := Read
	if r.Write {
		p = Write
	}
	return self.check(r.Selection.Path().StringNoModule() + "/" + r.Meta.GetIdent(), p)
}

func (self *Acl) check(targetPath string, p Permission) (bool, error) {
	i := self.List.Front()
	for i != nil {
		ac := i.Value.(*AccessControl)
		if ac.pathRegx.MatchString(targetPath) {
			if (ac.Permissions & p) == p {
				return true, nil
			}
		}
		i = i.Next()
	}
	return false, UnauthorizedError
}
