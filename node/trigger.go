package node

import (
	"container/list"
)

// TriggerTable is registry of all trigger functions for a browser
type TriggerTable struct {
	table *list.List
}

// NewTriggerTable creates new registry of all trigger functions for a browser
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
	i := self.table.Front()
	for i != nil {
		t := i.Value.(*Trigger)
		if begin && t.OnBegin != nil {
			if err := t.OnBegin(t, r); err != nil {
				return err
			}
		} else if !begin && t.OnEnd != nil {
			if err := t.OnEnd(t, r); err != nil {
				return err
			}
		}
		i = i.Next()
	}
	return nil
}

// Install will register new trigger functions
func (self *TriggerTable) Install(t *Trigger) {
	t.hnd = self.table.PushBack(t)
}

// Remove will no longer call trigger funcs
func (self *TriggerTable) Remove(t *Trigger) {
	if t.hnd != nil {
		self.table.Remove(t.hnd)
	}
}

// TriggerFunc callback for triggers
type TriggerFunc func(t *Trigger, r NodeRequest) error

// Trigger for registering global listeners on a data model.  Useful for triggering on save or
// push notifications
type Trigger struct {
	hnd *list.Element

	// OnBegin will be called on the start of any edit. Return error to stop edit
	OnBegin TriggerFunc

	// OnEnd will be called at the end of any edit. Return error to mark edit as failed
	OnEnd TriggerFunc
}
