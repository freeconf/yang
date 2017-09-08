package secure

import "github.com/c2stack/c2g/node"

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
