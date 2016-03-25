package yang

import (
	"meta"
	"testing"
)

func TestParseModuleStatement(t *testing.T) {
	yang := `
module ff {
	namespace "ns";

	description "mod";

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
	l := lex(yang, nil)
	err := yyParse(l)
	if err != 0 {
		t.Errorf("Error parsing %d", err)
	}
	d := l.stack.Peek()
	m := d.(*meta.Module)
	if m.Ident != "ff" {
		t.Errorf("module name expected ff, got %s", m.Ident)
	}
	if m.Revision.Ident != "99-99-9999" {
		t.Errorf("revision is %s", m.Revision.Ident)
	}
	if m.GetFirstMeta() == nil {
		t.Errorf("Container x is missing")
	}
	if m.GetFirstMeta().GetIdent() != "x" {
		t.Errorf("Container x not identified")
	}
}
