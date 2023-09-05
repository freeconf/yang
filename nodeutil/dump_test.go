package nodeutil

import (
	"bytes"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/parser"
)

func TestDump(t *testing.T) {
	mstr := `
module food {
	prefix "x";
	namespace "y";
	revision 0000-00-00 {
		description "";
	}
	list fruits  {
		key "name";
		leaf name {
			type string;
		}
		container origin {
			leaf country {
				type string;
			}
		}
	}
}
`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}

	var dump bytes.Buffer
	out := Dump(Null(), &dump)
	ypath := source.Dir("../yang")
	ymod := parser.RequireModule(ypath, "fc-yang")
	for _, d := range ymod.DataDefinitions()[0].(meta.HasDataDefinitions).DataDefinitions() {
		t.Logf("def %s", d.Ident())
	}
	fc.AssertEqual(t, err, Schema(ymod, m).Root().InsertInto(out))
	t.Log(dump.String())
}
