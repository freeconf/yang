package node

import (
	"testing"

	"github.com/c2stack/c2g/meta/yang"
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

//func TestSelectionEvents(t *testing.T) {
//	m, err := yang.LoadModuleCustomImport(selectionTestModule, nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	data := make(map[string]interface{})
//	b := NewBrowser2(m, MapNode(data))
//	sel := b.Root()
//	var relPathFired bool
//	b.Triggers.Install(&Trigger{
//		Origin:    "x",
//		Target:    "m/message",
//		EventType: END_EDIT,
//		OnFire: func(t *Trigger, e Event) error {
//			relPathFired = true
//			return nil
//		},
//	})
//	var regexFired bool
//	b.Triggers.Install(&Trigger{
//		Origin: "y",
//		TargetRegx: regexp.MustCompile(".*"),
//		EventType: DELETE,
//		OnFire: func(*Trigger, Event) error {
//			regexFired = true
//			return nil
//		},
//	})
//	json := NewJsonReader(strings.NewReader(`{"message":{"hello":"bob"}}`)).Node()
//	if err = sel.UpsertFrom(json).LastErr; err != nil {
//		t.Fatal(err)
//	}
//	if !relPathFired {
//		t.Fatal("Event not fired")
//	}
//	sel.Find("deep").Delete()
//	if !regexFired {
//		t.Fatal("regex not fired")
//	}
//}

func TestSelectionPeek(t *testing.T) {
	m, err := yang.LoadModuleCustomImport(selectionTestModule, nil)
	if err != nil {
		t.Fatal(err)
	}
	var expected = "Justin Bieber Fan Club Member"
	n := &MyNode{
		Peekable: expected,
	}
	sel := NewBrowser(m, n).Root()
	actual := sel.Peek(t)
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
