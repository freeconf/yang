package node

import (
	"regexp"
	"strings"
	"testing"

	"github.com/dhubler/c2g/meta/yang"
)

var selectionTestModule = `
module m {
	prefix "";
	namespace "";
	revision 0;
	container message {
		leaf hello {
			type string;
		}
		container deep {
			leaf goodbye {
				type string;
			}
		}
	}
}
`

func TestSelectionEvents(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(selectionTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	store := NewBufferStore()
	b := NewStoreData(m, store).Browser()
	sel := b.Root()
	var relPathFired bool
	b.Triggers.Install(&Trigger{
		Origin:    "x",
		Target:    "m/message",
		EventType: NEW,
		OnFire: func(t *Trigger, e Event) error {
			relPathFired = true
			return nil
		},
	})
	var regexFired bool
	b.Triggers.Install(&Trigger{
		Origin:     "y",
		TargetRegx: regexp.MustCompile(".*"),
		EventType:  LEAVE_EDIT,
		OnFire: func(*Trigger, Event) error {
			regexFired = true
			return nil
		},
	})
	json := NewJsonReader(strings.NewReader(`{"message":{"hello":"bob"}}`)).Node()
	if err = sel.Selector().UpsertFrom(json).LastErr; err != nil {
		t.Fatal(err)
	}
	if !relPathFired {
		t.Fatal("Event not fired")
	}
	if !regexFired {
		t.Fatal("regex not fired")
	}
}

func TestSelectionPeek(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(selectionTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	var expected = "Justin Bieber Fan Club Member"
	n := &MyNode{
		Peekable: expected,
	}
	sel := NewBrowser2(m, n).Root()
	actual := sel.Peek()
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
