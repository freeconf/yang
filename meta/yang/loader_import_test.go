package yang

import (
	"fmt"
	"testing"

	"github.com/dhubler/c2g/meta"
)

func TestLoaderImport(t *testing.T) {
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
	import sub;
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
		if meta.FindByIdent2(m, "x") == nil {
			t.Error("Could not find x container")
		}
		if meta.FindByIdent2(m, "sub-x") == nil {
			t.Error("Could not find sub-x container")
		}
	}
}
