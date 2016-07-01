package node

import (
	"container/list"
	"regexp"
	"strings"
	"fmt"
)

type TriggerTable struct {
	table *list.List
}

func NewTriggerTable() *TriggerTable {
	return &TriggerTable{
		table: list.New(),
	}
}

func (self *TriggerTable) Fire(target string, e Event) error {
	// TODO: Threadsafe
	var lastErr error
	i := self.table.Front()
	for i != nil {
		var err error
		t := i.Value.(*Trigger)
		if t.EventType == e.Type {
			if t.TargetRegx != nil {
				if t.TargetRegx.MatchString(target) {
					err = t.OnFire(t, e)
				}
			} else if t.Target == target {
				err = t.OnFire(t, e)
			}
		}
		if err != nil && lastErr == nil {
			lastErr = err
		}
		i = i.Next()
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

// RemoveByOrigin removes all children of this origin as well
func (self *TriggerTable) RemoveByOrigin(origin string) {
	i := self.table.Front()
	for i != nil {
		t := i.Value.(*Trigger)
		if strings.HasPrefix(t.Origin, origin) {
			self.table.Remove(i)
		}
		i = i.Next()
	}
}

type TriggerFunc func(t *Trigger, e Event) error

type Trigger struct {
	Origin     string
	Target     string
	TargetRegx *regexp.Regexp
	hnd        *list.Element
	EventType  EventType
	OnFire     TriggerFunc
}

func (self *Trigger) String() string {
	target := self.Target
	if self.TargetRegx != nil {
		target = self.TargetRegx.String()
	}
	return fmt.Sprintf("%v:%s, origin=%s, listener=%p", self.EventType, target, self.Origin, self.OnFire)
}

