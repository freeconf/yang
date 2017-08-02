package nodes

import (
	"bytes"
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
)

func TestDecoupledMetaCopy(t *testing.T) {
	m, _ := yang.LoadModuleCustomImport(yang.TestDataSimpleYang, nil)
	tape, _ := meta.FindByPath(m, "turing-machine/tape")
	yangPath := meta.PathStreamSource("../yang")
	tapeCopy, _ := DecoupledMetaCopy(yangPath, tape.(meta.MetaList))
	if tapeCopy == nil {
		t.Error("null")
	}
	if tapeCopy.GetIdent() != "tape" {
		t.Error(tapeCopy.GetIdent())
	}
	// with meta decoupled, we should be able to navigate tape meta w/o "tape-cells" group
	test := &meta.Module{Ident: "test"}
	test.AddMeta(tapeCopy)
	var actualBytes bytes.Buffer
	err := SelectModule(test, true).Root().InsertInto(NewJsonWriter(&actualBytes).Node()).LastErr
	if err != nil {
		t.Error(err)
	}
	t.Log(actualBytes.String())
}
