package yang

import (
	"fmt"
	"testing"

	"github.com/c2stack/c2g/meta"
)

func TestInclude(t *testing.T) {
	subYang := `
module sub {
	namespace "sub-ns";
	description "sub mod";
	revision 99-99-9999 {
	  description "bingo";
	}

	container sub-x {
	  description "sub-z";
	  leaf sub-y {
	    type int32;
	  }
	}
}
`
	mainYang := `
module main {
	namespace "ns";
	description "mod";
	include sub;
	revision 99-99-9999 {
	  description "bingo";
	}

	container x {
	  description "z";
	  leaf y {
	    type int32;
	  }
	}
}
	`
	resources := func(resource string) (string, error) {
		switch resource {
		case "main":
			return mainYang, nil
		case "sub":
			return subYang, nil
		default:
			return "", &yangError{fmt.Sprint("Unexpected resource ", resource)}
		}
	}
	source := &meta.StringSource{Streamer: resources}
	m, err := LoadModule(source, "main")
	if err != nil {
		t.Error(err)
	} else {
		if m, err := meta.Find(m, "x"); m == nil || err != nil {
			t.Error("Could not find x container, err=%s", err)
		}
		if m, err := meta.Find(m, "sub-x"); m == nil || err != nil {
			t.Error("Could not find sub-x container, err=%s", err)
		}
	}
}
