package meta

import (
	"container/list"
)

type defs struct {
	actions       map[string]*Rpc
	notifications map[string]*Notification
	dataDefs      []Definition
	dataDefsIndex map[string]Definition
	unresolved    *list.List
	compiled      bool
}

func newDefs() *defs {
	return &defs{
		actions:       make(map[string]*Rpc),
		notifications: make(map[string]*Notification),
		unresolved:    list.New(),
	}
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

type resolvedListener func(Definition) error

func (self *defs) resolve(parent Meta, pool schemaPool) error {
	if self.unresolved == nil {
		return nil
	}

	// size is rough estimate and could be larger as 'uses' get resolved
	self.dataDefs = make([]Definition, 0, self.unresolved.Len())
	self.dataDefsIndex = make(map[string]Definition, self.unresolved.Len())

	// we remove all datadefs, then add them back in except Uses, which we
	// resolve.  We go horizontal (across children) first in tree not down
	// to support circular groupings.
	unresolved := self.unresolved
	self.unresolved = nil
	resolved := func(ddef Definition) error {
		self.add(parent, ddef)
		return nil
	}
	if err := self.resolveDefs(parent, pool, unresolved, resolved); err != nil {
		return err
	}

	// move down tree
	for _, x := range self.notifications {
		if err := x.resolve(pool); err != nil {
			return err
		}
	}
	for _, x := range self.actions {
		if err := x.resolve(pool); err != nil {
			return err
		}
	}
	for _, x := range self.dataDefs {
		if r, ok := x.(resolver); ok {
			if err := r.resolve(pool); err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *defs) resolveDefs(parent Meta, pool schemaPool, unresolved *list.List, resolved resolvedListener) error {
	for p := unresolved.Front(); p != nil; p = p.Next() {
		switch x := p.Value.(type) {
		case *Uses:
			if err := x.resolve(parent, pool, resolved); err != nil {
				return err
			}
		case *Choice:
			if err := resolved(x); err != nil {
				return err
			}

		default:
			if err := resolved(p.Value.(Definition)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *defs) compile(parent Meta) error {
	if self.compiled {
		return nil
	}
	self.compiled = true
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
	for _, x := range self.dataDefs {
		if c, ok := x.(*Choice); ok {
			self.buildCases(c)
		}
		if err := x.(compilable).compile(); err != nil {
			return err
		}
	}
	return nil
}

func (y *defs) buildCases(c *Choice) {
	for _, kase := range c.Cases() {
		for _, ddef := range kase.DataDefs() {
			y.dataDefsIndex[ddef.Ident()] = ddef
			if nested, ok := ddef.(*Choice); ok {
				y.buildCases(nested)
			}
		}
	}
}

// clones
func (self *defs) clone(parent Meta) *defs {
	copy := &defs{
		actions:       make(map[string]*Rpc, len(self.actions)),
		notifications: make(map[string]*Notification, len(self.notifications)),
		unresolved:    list.New(),
	}
	for _, x := range self.actions {
		y := x.clone(parent).(*Rpc)
		copy.actions[y.Ident()] = y
	}
	for _, x := range self.notifications {
		y := x.clone(parent).(*Notification)
		copy.notifications[y.Ident()] = y
	}

	for p := self.unresolved.Front(); p != nil; p = p.Next() {
		y := p.Value.(cloneable).clone(parent)
		copy.unresolved.PushBack(y)
	}
	return copy
}

func (self *defs) add(parent Meta, d Definition) {
	if self.unresolved != nil {
		self.unresolved.PushBack(d)
	} else {
		if hasIf, ok := d.(HasIfFeatures); ok {
			if on, err := checkFeature(hasIf); err != nil {
				panic(err.Error())
			} else if !on {
				return
			}
		}

		switch x := d.(type) {
		case *Rpc:
			self.actions[x.Ident()] = x
			return
		case *Notification:
			self.notifications[x.Ident()] = x
			return
		}
		// when we're adding resolved data-defs back in, we need to
		// replace duplicates because this could be from an augment
		if _, found := self.dataDefsIndex[d.Ident()]; found {
			for i, x := range self.dataDefs {
				if d.Ident() == x.Ident() {
					self.dataDefsIndex[d.Ident()] = d
					self.dataDefs[i] = d
					return
				}
			}
		}
		self.dataDefsIndex[d.Ident()] = d
		self.dataDefs = append(self.dataDefs, d)
	}
}
