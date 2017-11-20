package yang

import (
	"testing"

	"github.com/freeconf/c2g/meta"
)

func TestImport(t *testing.T) {
	subYang := `
module sub {
	namespace "";
	revision 0;

	grouping x {
	  leaf sub {
	    type int32;
	  }
	}
}`

	mainYang := `
module main {
	namespace "";
	import sub {
		prefix "s";
	}
	revision 0;

	grouping x {
	  leaf main {
	    type int32;
	  }
	}
	container x {
	  leaf y {
	    type int32;
	  }
	}
	uses s:x;
}`
	source := &meta.StringSource{Streamer: func(m string) (string, error) {
		switch m {
		case "main":
			return mainYang, nil
		case "sub":
			return subYang, nil
		}
		panic(m)
	}}
	m, err := LoadModule(source, "main")
	if err != nil {
		t.Error(err)
	} else {
		if m := meta.Find(m, "sub"); m == nil {
			t.Error("Could not find s:sub container")
		}
	}
}
