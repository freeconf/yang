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
		t.Log(trigger)
		c2.AssertEqual(t, 1, beginCount)
		c2.AssertEqual(t, 0, endCount)
		table.handle("bbb", r, false)
		c2.AssertEqual(t, 1, endCount)
		table.handle("ccc", r, true)
		c2.AssertEqual(t, 1, beginCount)
		table.Remove(trigger)
		table.handle("bbb", r, true)
		c2.AssertEqual(t, 1, beginCount)
		table.Install(trigger)
		table.handle("bbb", r, true)
		c2.AssertEqual(t, 2, beginCount)
	}
}
