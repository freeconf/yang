package node

import (
	"context"
	"reflect"
	"sort"

	"github.com/freeconf/yang/val"
)

type ActionPreConstraint interface {
	CheckActionPreConstraints(r *ActionRequest) (bool, error)
}

type ActionPostConstraint interface {
	CheckActionPostConstraints(r ActionRequest) (bool, error)
}

type ListPreConstraint interface {
	CheckListPreConstraints(r *ListRequest) (bool, error)
}

type ListPostConstraint interface {

	// returns
	//   proceed - stop processing any possible, subsequentremaining items in list
	//   visible - don't process this list item
	//   err - if something happened
	CheckListPostConstraints(r ListRequest, child *Selection, key []val.Value) (bool, bool, error)
}

type ContainerPreConstraint interface {
	CheckContainerPreConstraints(r *ChildRequest) (bool, error)
}

type ContainerPostConstraint interface {
	CheckContainerPostConstraints(r ChildRequest, child *Selection) (bool, error)
}

type FieldPreConstraint interface {
	CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error)
}

type FieldPostConstraint interface {
	CheckFieldPostConstraints(r FieldRequest, hnd *ValueHandle) (bool, error)
}

type NotifyFilterConstraint interface {
	CheckNotifyFilterConstraints(msg *Selection) (bool, error)
}

type ContextConstraint interface {
	ContextConstraint(*Selection) context.Context
}

type entry struct {
	id           string
	weight       int
	priority     int
	constraint   interface{}
	prelist      ListPreConstraint
	postlist     ListPostConstraint
	precont      ContainerPreConstraint
	postcont     ContainerPostConstraint
	prefield     FieldPreConstraint
	postfield    FieldPostConstraint
	preaction    ActionPreConstraint
	postaction   ActionPostConstraint
	notifyfilter NotifyFilterConstraint
	ctx          ContextConstraint
}

type Constraints struct {
	entries  []*entry
	compiled entrySlice
}

func NewConstraints(parent *Constraints) *Constraints {
	c := &Constraints{}
	c.entries = make([]*entry, len(parent.entries))
	for i, e := range parent.entries {
		c.entries[i] = e
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
	atLeastOneMatch := false
	e := &entry{
		id:         id,
		weight:     weight,
		priority:   priority,
		constraint: constraint,
	}
	if v, ok := constraint.(ListPreConstraint); ok {
		atLeastOneMatch = true
		e.prelist = v
	}
	if v, ok := constraint.(ListPostConstraint); ok {
		atLeastOneMatch = true
		e.postlist = v
	}
	if v, ok := constraint.(ContainerPreConstraint); ok {
		atLeastOneMatch = true
		e.precont = v
	}
	if v, ok := constraint.(ContainerPostConstraint); ok {
		atLeastOneMatch = true
		e.postcont = v
	}
	if v, ok := constraint.(FieldPreConstraint); ok {
		atLeastOneMatch = true
		e.prefield = v
	}
	if v, ok := constraint.(FieldPostConstraint); ok {
		atLeastOneMatch = true
		e.postfield = v
	}
	if v, ok := constraint.(ActionPreConstraint); ok {
		atLeastOneMatch = true
		e.preaction = v
	}
	if v, ok := constraint.(ActionPostConstraint); ok {
		atLeastOneMatch = true
		e.postaction = v
	}
	if v, ok := constraint.(NotifyFilterConstraint); ok {
		atLeastOneMatch = true
		e.notifyfilter = v
	}
	if v, ok := constraint.(ContextConstraint); ok {
		atLeastOneMatch = true
		e.ctx = v
	}
	if !atLeastOneMatch {
		panic(reflect.TypeOf(constraint).Name() + " does not implement any of the known constraint types.")
	}
	self.entries = append(self.entries, e)
	self.compiled = nil
}

func (self *Constraints) Constraint(id string) interface{} {
	for _, e := range self.entries {
		if e.id == id {
			return e
		}
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

func (self *Constraints) CheckListPreConstraints(r *ListRequest) (bool, error) {
	for _, v := range self.compile() {
		if v.prelist != nil {
			if more, err := v.prelist.CheckListPreConstraints(r); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckListPostConstraints(r ListRequest, child *Selection, key []val.Value) (bool, bool, error) {
	for _, v := range self.compile() {
		if v.postlist != nil {
			if more, visible, err := v.postlist.CheckListPostConstraints(r, child, key); !more || !visible || err != nil {
				return more, visible, err
			}
		}
	}
	return true, true, nil
}

func (self *Constraints) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	for _, v := range self.compile() {
		if v.precont != nil {
			if more, err := v.precont.CheckContainerPreConstraints(r); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckContainerPostConstraints(r ChildRequest, child *Selection) (bool, error) {
	for _, v := range self.compile() {
		if v.postcont != nil {
			if more, err := v.postcont.CheckContainerPostConstraints(r, child); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	for _, v := range self.compile() {
		if v.prefield != nil {
			if more, err := v.prefield.CheckFieldPreConstraints(r, hnd); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckFieldPostConstraints(r FieldRequest, hnd *ValueHandle) (bool, error) {
	for _, v := range self.compile() {
		if v.postfield != nil {
			if more, err := v.postfield.CheckFieldPostConstraints(r, hnd); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckActionPreConstraints(r *ActionRequest) (bool, error) {
	for _, v := range self.compile() {
		if v.preaction != nil {
			if more, err := v.preaction.CheckActionPreConstraints(r); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckActionPostConstraints(r ActionRequest) (bool, error) {
	for _, v := range self.compile() {
		if v.postaction != nil {
			if more, err := v.postaction.CheckActionPostConstraints(r); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) CheckNotifyFilterConstraints(msg *Selection) (bool, error) {
	for _, v := range self.compile() {
		if v.notifyfilter != nil {
			if more, err := v.notifyfilter.CheckNotifyFilterConstraints(msg); !more || err != nil {
				return more, err
			}
		}
	}
	return true, nil
}

func (self *Constraints) ContextConstraint(s *Selection) context.Context {
	c := s.Context
	for _, v := range self.compile() {
		if v.ctx != nil {
			c = v.ctx.ContextConstraint(s)
		}
	}
	return c
}
