package node

import (
	"meta"
)

type Context struct {
	NoFail      bool
	NewLocation string
	Violations  []error
	Constraints *Constraints
}

func NewContext() *Context {
	return &Context{}
}

func (self *Context) IncompleteResponse(location *Path) {
	//blit.Err.Println("Incomplete response served at " + location.String())
}

func (self *Context) LocatableNode(location *Path) {
	self.NewLocation = location.String()
}

func (self *Context) ConstraintViolation(violation error) error {
	if !self.NoFail {
		return violation
	}
	if self.Violations == nil {
		self.Violations = []error{violation}
	} else {
		self.Violations = append(self.Violations, violation)
	}
	return nil
}

func (self *Context) Select(m meta.MetaList, node Node) Selector {
	return Selector{
		Context:   self,
		Selection: Select(m, node),
		Constraints: self.Constraints,
	}
}

func (self *Context) Selector(s *Selection) Selector {
	return Selector{
		Context:   self,
		Selection: s,
		Constraints: self.Constraints,
	}
}
