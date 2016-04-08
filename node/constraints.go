package node

import (
	"reflect"
	"sort"
)

type entry struct {
	id         string
	weight     int
	priority   int
	constraint interface{}
	prelist    ListPreConstraint
	postlist   ListPostConstraint
	precont    ContainerPreConstraint
	postcont   ContainerPostConstraint
	prefield   FieldPreConstraint
	postfield  FieldPostConstraint
}

type Constraints struct {
	entries  map[string]*entry
	compiled entrySlice
}

func NewConstraints(parent *Constraints) *Constraints {
	c := &Constraints{}
	if parent != nil && parent.entries != nil{
		c.entries = make(map[string]*entry, len(parent.entries))
		for k, e := range parent.entries {
			c.entries[k] = e
		}
	}
	return c
}

type entrySlice []*entry

func (self entrySlice) Len() int {
	return len(self)
}

func (self entrySlice) Less(a int, b int) bool {
	if self[a].priority != self[b].priority {
		return self[a].priority < self[b].priority
	}
	return self[a].weight < self[b].weight
}

func (self entrySlice) Swap(a int, b int) {
	self[a], self[b] = self[b], self[a]
}

/*
 priority - the priority of the target constraint, lower value means more preferred.
 weight - A relative weight for constraint with the same priority, higher value means more preferred.
*/

func (self *Constraints) AddConstraint(id string, weight int, priority int, constraint interface{}) {
	oneMatch := false
	e := &entry{
		id:         id,
		weight:     weight,
		priority:   priority,
		constraint: constraint,
	}
	if v, ok := constraint.(ListPreConstraint); ok {
		oneMatch = true
		e.prelist = v
	}
	if v, ok := constraint.(ListPostConstraint); ok {
		oneMatch = true
		e.postlist = v
	}
	if v, ok := constraint.(ContainerPreConstraint); ok {
		oneMatch = true
		e.precont = v
	}
	if v, ok := constraint.(ContainerPostConstraint); ok {
		oneMatch = true
		e.postcont = v
	}
	if v, ok := constraint.(FieldPreConstraint); ok {
		oneMatch = true
		e.prefield = v
	}
	if v, ok := constraint.(FieldPostConstraint); ok {
		oneMatch = true
		e.postfield = v
	}
	if !oneMatch {
		panic(reflect.TypeOf(constraint).Name() + " does not implement any of the known constraint types.")
	}
	if self.entries == nil {
		self.entries = make(map[string]*entry, 1)
	}
	self.entries[id] = e
	self.compiled = nil
}

func (self *Constraints) Constraint(id string) interface{} {
	if e, found := self.entries[id]; found {
		return e.constraint
	}
	return nil
}

func (self *Constraints) compile() entrySlice {
	if self.compiled == nil {
		compiled := make(entrySlice, len(self.entries))
		i := 0
		for _, v := range self.entries {
			compiled[i] = v
			i++
		}

		sort.Sort(compiled)
		self.compiled = compiled
	}
	return self.compiled
}

func (self *Constraints) CheckListPreConstraints(r *ListRequest, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.prelist != nil {
			if more, err := v.prelist.CheckListPreConstraints(r, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckListPostConstraints(r ListRequest, child *Selection, key []*Value, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.postlist != nil {
			if more, err := v.postlist.CheckListPostConstraints(r, child, key, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckContainerPreConstraints(r *ContainerRequest, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.precont != nil {
			if more, err := v.precont.CheckContainerPreConstraints(r, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckContainerPostConstraints(r ContainerRequest, child *Selection, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.postcont != nil {
			if more, err := v.postcont.CheckContainerPostConstraints(r, child, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckFieldPreConstraints(r *FieldRequest, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.prefield != nil {
			if more, err := v.prefield.CheckFieldPreConstraints(r, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckFieldPostConstraints(r FieldRequest, val *Value, navigating bool) (bool, error) {
	for _, v := range self.compile() {
		if v.postfield != nil {
			if more, err := v.postfield.CheckFieldPostConstraints(r, val, navigating); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

type ListPreConstraint interface {
	CheckListPreConstraints(r *ListRequest, navigating bool) (bool, error)
}

type ListPostConstraint interface {
	CheckListPostConstraints(r ListRequest, child *Selection, key []*Value, navigating bool) (bool, error)
}

type ContainerPreConstraint interface {
	CheckContainerPreConstraints(r *ContainerRequest, navigating bool) (bool, error)
}

type ContainerPostConstraint interface {
	CheckContainerPostConstraints(r ContainerRequest, child *Selection, navigating bool) (bool, error)
}

type FieldPreConstraint interface {
	CheckFieldPreConstraints(r *FieldRequest, navigating bool) (bool, error)
}

type FieldPostConstraint interface {
	CheckFieldPostConstraints(r FieldRequest, v *Value, navigating bool) (bool, error)
}
