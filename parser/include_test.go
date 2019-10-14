package parser

import (
	"strings"
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/source"
)

func TestInclude(t *testing.T) {
	subYang := `
submodule sub {
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
	source := source.Any(
		source.Named("main", strings.NewReader(mainYang)),
		source.Named("sub", strings.NewReader(subYang)))
	m, err := LoadModule(source, "main", "")
	if err != nil {
		t.Error(err)
	} else {
		if m := meta.Find(m, "x"); m == nil {
			t.Error("Could not find x container")
		}
		if m := meta.Find(m, "sub-x"); m == nil {
			t.Error("Could not find sub-x container")
		}
	}
}
