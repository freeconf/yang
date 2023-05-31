package parser

import (
	"strings"
	"testing"

	"github.com/freeconf/yang/fc"
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
	m, err := LoadModule(source, "main")
	fc.RequireEqual(t, nil, err)
	x := meta.Find(m, "x")
	fc.AssertEqual(t, true, x != nil, "Could not find x container")
	subx := meta.Find(m, "sub-x")
	fc.AssertEqual(t, true, subx != nil, "Could not find sub-x container")
	fc.AssertEqual(t, "mod", m.Description())
}
