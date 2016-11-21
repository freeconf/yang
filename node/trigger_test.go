package node

import (
	"regexp"
	"testing"

	"github.com/c2stack/c2g/c2"
)

type TriggerEvent struct {
}

func TestTriggers(t *testing.T) {
	var beginCount int
	var endCount int
	tests := []*Trigger{
		&Trigger{
			Target: "bbb",
			OnBegin: func(*Trigger, NodeRequest) error {
				beginCount++
				return nil
			},
			OnEnd: func(*Trigger, NodeRequest) error {
				endCount++
				return nil
			},
		},
		&Trigger{
			TargetRegx: regexp.MustCompile("b.*"),
			OnBegin: func(*Trigger, NodeRequest) error {
				beginCount++
				return nil
			},
			OnEnd: func(*Trigger, NodeRequest) error {
				endCount++
				return nil
			},
		},
	}
	for _, test := range tests {
		beginCount = 0
		endCount = 0
		table := NewTriggerTable()
		trigger := test
		table.Install(trigger)
		var r NodeRequest
		table.handle("bbb", r, true)
		if err := c2.CheckEqual(1, beginCount); err != nil {
			t.Error(trigger, err)
		}
		if err := c2.CheckEqual(0, endCount); err != nil {
			t.Error(trigger, err)
		}
		table.handle("bbb", r, false)
		if err := c2.CheckEqual(1, endCount); err != nil {
			t.Error(trigger, err)
		}
		table.handle("ccc", r, true)
		if err := c2.CheckEqual(1, beginCount); err != nil {
			t.Error(trigger, err)
		}
		table.Remove(trigger)
		table.handle("bbb", r, true)
		if err := c2.CheckEqual(1, beginCount); err != nil {
			t.Error(trigger, err)
		}
		table.Install(trigger)
		table.handle("bbb", r, true)
		if err := c2.CheckEqual(2, beginCount); err != nil {
			t.Error(trigger, err)
		}
	}
}
