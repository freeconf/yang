package node

import (
	"testing"
	"strings"
	"github.com/c2g/meta/yang"
	"regexp"
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
	sel := NewStoreData(m, store).Select()
	var relPathFired bool
	sel.OnPath(NEW, "m/message", func() error {
		relPathFired = true
		return nil
	})
	var regexFired bool
	sel.OnRegex(LEAVE_EDIT, regexp.MustCompile(".*"), func() error {
		regexFired = true
		return nil
	})
	json := NewJsonReader(strings.NewReader(`{"message":{"hello":"bob"}}`)).Node()
	c := NewContext()
	if err = c.Selector(sel).UpsertFrom(json).LastErr; err != nil {
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
		Peekables:map[string]interface{} {"a":expected},
	}
	sel := Select(m, n)
	actual :=  sel.Peek("a")
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}
