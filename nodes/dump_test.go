package nodes

import (
	"bytes"
	"testing"

	"github.com/freeconf/gconf/meta"

	"github.com/freeconf/gconf/meta/yang"
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
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}

	var dump bytes.Buffer
	out := Dump(Null(), &dump)
	ypath := &meta.FileStreamSource{Root: "../yang"}
	ymod := yang.RequireModule(ypath, "yang")
	for _, d := range ymod.DataDefs()[0].(meta.HasDataDefs).DataDefs() {
		t.Logf("def %s", d.Ident())
	}
	if err = Schema(ymod, m).Root().InsertInto(out).LastErr; err != nil {
		t.Fatal(err)
	}
	t.Log(dump.String())
}
