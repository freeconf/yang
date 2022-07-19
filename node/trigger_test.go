package node

import (
	"testing"

	"github.com/freeconf/yang/fc"
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
	table.handle(r, true)
	table.handle(r, false)
	fc.AssertEqual(t, 1, beginCount)
	fc.AssertEqual(t, 1, endCount)
	table.Remove(tgr)

	table.handle(r, true)
	table.handle(r, false)
	fc.AssertEqual(t, 1, beginCount)
	fc.AssertEqual(t, 1, endCount)
}
