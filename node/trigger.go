package node

import (
	"container/list"
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
	// this is called A LOT so we avoid doing anything if we do not have to
	if self.table.Len() == 0 {
		return nil
	}
	return self.handle(r.Selection.Path.String(), r, true)
}

func (self *TriggerTable) endEdit(r NodeRequest) error {
	// this is called A LOT so we avoid doing anything if we do not have to
	if self.table.Len() == 0 {
		return nil
	}
	return self.handle(r.Selection.Path.String(), r, false)
}

func (self *TriggerTable) handle(path string, r NodeRequest, begin bool) error {
	var err error
	i := self.table.Front()
	for i != nil {
		t := i.Value.(*Trigger)
		if begin && t.OnBegin != nil {
			return t.OnBegin(t, r)
		} else if !begin && t.OnEnd != nil {
			return t.OnEnd(t, r)
		}
		i = i.Next()
	}
	return err
}

func (self *TriggerTable) Install(t *Trigger) {
	t.hnd = self.table.PushBack(t)
}

func (self *TriggerTable) Remove(t *Trigger) {
	if t.hnd != nil {
		self.table.Remove(t.hnd)
	}
}

type TriggerFunc func(t *Trigger, r NodeRequest) error

type Trigger struct {
	hnd     *list.Element
	OnBegin TriggerFunc
	OnEnd   TriggerFunc
}
