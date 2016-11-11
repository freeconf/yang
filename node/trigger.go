package node

import (
	"container/list"
	"fmt"
	"regexp"
)

type TriggerTable struct {
	table *list.List
}

func NewTriggerTable() *TriggerTable {
	return &TriggerTable{
		table: list.New(),
	}
}

func (self *TriggerTable) beginEdit(r NodeRequest) error {
	return self.handle(r.Selection.Path.String(), r, true)
}

func (self *TriggerTable) endEdit(r NodeRequest) error {
	return self.handle(r.Selection.Path.String(), r, false)
}

func (self *TriggerTable) handle(path string, r NodeRequest, begin bool) error {
	// TODO: Threadsafe
	var err error
	i := self.table.Front()
	for i != nil {
		t := i.Value.(*Trigger)
		if t.TargetRegx != nil {
			if t.TargetRegx.MatchString(path) {
				err = self.fire(t, r, begin, err)
			}
		} else if t.Target == path {
			err = self.fire(t, r, begin, err)
		}
		i = i.Next()
	}
	return err
}

func (self *TriggerTable) fire(t *Trigger, r NodeRequest, begin bool, lastErr error) error {
	if begin && t.OnBegin != nil {
		return t.OnBegin(t, r)
	} else if !begin && t.OnEnd != nil {
		return t.OnEnd(t, r)
	}
	return lastErr
}

func (self *TriggerTable) Install(t *Trigger) {
	// TODO: Threadsafe
	t.hnd = self.table.PushBack(t)
}

func (self *TriggerTable) Remove(t *Trigger) {
	// TODO: Threadsafe
	if t.hnd != nil {
		self.table.Remove(t.hnd)
	}
}

type TriggerFunc func(t *Trigger, r NodeRequest) error

type Trigger struct {
	Target     string
	TargetRegx *regexp.Regexp
	hnd        *list.Element
	OnBegin    TriggerFunc
	OnEnd      TriggerFunc
}

func (self *Trigger) String() string {
	target := self.Target
	if self.TargetRegx != nil {
		target = self.TargetRegx.String()
	}
	return fmt.Sprintf("%s, onBegin=%p, onEnd=%p", target, self.OnBegin, self.OnEnd)
}
