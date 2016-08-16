package node

import (
	"testing"
	"github.com/dhubler/c2g/c2"
	"regexp"
)

func TestTriggers(t *testing.T) {
	var count int
	tests := []*Trigger {
		&Trigger{
			Origin:"aaa/string",
			Target:"bbb",
			EventType: NEW,
			OnFire: func(t *Trigger, e Event) error {
				count++
				return nil
			},
		},
		&Trigger{
			Origin:"aaa/regexp",
			TargetRegx:regexp.MustCompile("b.*"),
			EventType: NEW,
			OnFire: func(t *Trigger, e Event) error {
				count++
				return nil
			},
		},
	}
	for _, test := range tests {
		count = 0
		table := NewTriggerTable()
		trigger := test
		table.Install(trigger)
		table.Fire("bbb", NEW.New(nil))
		if err := c2.CheckEqual(1, count); err != nil {
			t.Error(trigger, err)
		}
		table.Fire("ccc", NEW.New(nil))
		if err := c2.CheckEqual(1, count); err != nil {
			t.Error(trigger, err)
		}
		table.Remove(trigger)
		table.Fire("bbb", NEW.New(nil))
		if err := c2.CheckEqual(1, count); err != nil {
			t.Error(trigger, err)
		}
		table.Install(trigger)
		table.Fire("bbb", NEW.New(nil))
		if err := c2.CheckEqual(2, count); err != nil {
			t.Error(trigger, err)
		}
		table.RemoveByOrigin("aaa")
		table.Fire("bbb", NEW.New(nil))
		if err := c2.CheckEqual(2, count); err != nil {
			t.Error(trigger, err)
		}
	}
}
