package node

import (
	"testing"

	"github.com/freeconf/gconf/c2"
)

type TriggerEvent struct {
}

func TestTriggers(t *testing.T) {
	beginCount := 0
	endCount := 0
	table := NewTriggerTable()
	tgr := &Trigger{
		OnBegin: func(*Trigger, NodeRequest) error {
			beginCount++
			return nil
		},
		OnEnd: func(*Trigger, NodeRequest) error {
			endCount++
			return nil
		},
	}
	table.Install(tgr)
	var r NodeRequest
	table.handle("bbb", r, true)
	table.handle("bbb", r, false)
	c2.AssertEqual(t, 1, beginCount)
	c2.AssertEqual(t, 1, endCount)
	table.Remove(tgr)

	table.handle("bbb", r, true)
	table.handle("bbb", r, false)
	c2.AssertEqual(t, 1, beginCount)
	c2.AssertEqual(t, 1, endCount)
}
