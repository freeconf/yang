package meta

import (
	"container/list"

	"github.com/c2stack/c2g/c2"
)

type defs struct {
	actions       map[string]*Rpc
	notifications map[string]*Notification
	dataDefs      []DataDef
	dataDefsIndex map[string]DataDef
	unresolved    *list.List
}

func newDefs() *defs {
	return &defs{
		actions:       make(map[string]*Rpc),
		notifications: make(map[string]*Notification),
		unresolved:    list.New(),
	}
}

func (self *defs) clone(parent Meta, deep bool) *defs {
	copy := &defs{}
	copy.notifications = make(map[string]*Notification, len(self.notifications))
	for _, x := range self.notifications {
		y := x.clone(deep)
		y.setParent(parent)
		copy.notifications[y.Ident()] = y
	}

	copy.actions = make(map[string]*Rpc, len(self.actions))
	for _, x := range self.actions {
		y := x.clone(deep)
		y.setParent(parent)
		copy.actions[y.Ident()] = y
	}
	copy.dataDefs = make([]DataDef, len(self.dataDefs))
	copy.dataDefsIndex = make(map[string]DataDef, len(self.dataDefs))
	for i, x := range self.dataDefs {
		y := x.clone(deep)
		y.setParent(parent)
		copy.dataDefs[i] = y
		copy.dataDefsIndex[y.Ident()] = y
	}
	return copy
}

func (self *defs) definition(ident string) Definition {
	if x, found := self.actions[ident]; found {
		return x
	}
	if x, found := self.notifications[ident]; found {
		return x
	}
	return self.dataDefsIndex[ident]
}

func (self *defs) compile(parent Meta) error {
	if self.unresolved == nil {
		return c2.NewErr(parent.(Identifiable).Ident() + " already compiled")
	}
	for _, x := range self.actions {
		if err := x.compile(); err != nil {
			return err
		}
	}
	for _, x := range self.notifications {
		if err := x.compile(); err != nil {
			return err
		}
	}
	approxLen := self.unresolved.Len() // sans expanding uses/groups
	self.dataDefs = make([]DataDef, 0, approxLen)
	self.dataDefsIndex = make(map[string]DataDef, approxLen)
	unresolved := self.unresolved
	self.unresolved = nil
	for p := unresolved.Front(); p != nil; p = p.Next() {
		c := p.Value.(compilable)
		if err := c.compile(); err != nil {
			return err
		}
		self.add(parent, c)
	}
	return nil
}

func (self *defs) add(parent Meta, o interface{}) {
	switch x := o.(type) {
	case *Rpc:
		x.parent = parent
		self.actions[x.Ident()] = x
		return
	case *Notification:
		x.parent = parent
		self.notifications[x.Ident()] = x
		return
	}
	if self.unresolved != nil {
		incoming := o.(movable)
		incoming.setParent(parent)
		self.unresolved.PushBack(incoming)
	} else {
		switch x := o.(type) {
		case *Choice:
			x.parent = parent
			self.indexChoiceCases(x)
		case *Uses:
			return
		}
		// when we're adding resolved data-defs back in, we need to
		// replace duplicates because this could be from an augment
		ddef := o.(DataDef)
		if _, found := self.dataDefsIndex[ddef.Ident()]; found {
			for i, x := range self.dataDefs {
				if ddef.Ident() == x.Ident() {
					self.dataDefsIndex[ddef.Ident()] = ddef
					self.dataDefs[i] = ddef
					return
				}
			}
		}
		self.dataDefsIndex[ddef.Ident()] = ddef
		self.dataDefs = append(self.dataDefs, ddef)
	}
}

// dive into choice cases because their contents are at same level in the information model.
func (self *defs) indexChoiceCases(c *Choice) {
	for _, cse := range c.DataDefs() {
		for _, d := range cse.(HasDataDefs).DataDefs() {
			self.dataDefsIndex[d.Ident()] = d
			if nested, valid := d.(*Choice); valid {
				self.indexChoiceCases(nested)
			}
		}
	}
}
