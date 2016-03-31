package node

import (
	"encoding/json"
	"testing"
	"github.com/c2g/meta/yang"
	"github.com/c2g/meta"
	"bytes"
)

func TestMapValue(t *testing.T) {
	var err error
	dataJson := `{"a":{"b":{"x":"waldo"}},"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`
	var data map[string]interface{}
	if err = json.Unmarshal([]byte(dataJson), &data); err != nil {
		t.Error(err)
	}
	if MapValue(data, "a.b.x") != "waldo" {
		t.Error("can't find waldo")
	}
	if MapValue(data, "p.1.k") != "waldo" {
		t.Error("can't find waldo")
	}
}

func TestDecoupledMetaCopy(t *testing.T) {
	m, _ := yang.LoadModuleCustomImport(yang.TestDataSimpleYang, nil)
	tape := meta.FindByPath(m, "turing-machine/tape").(meta.MetaList)
	tapeCopy := DecoupledMetaCopy(tape)
	if tapeCopy == nil {
		t.Error("null")
	}
	if tapeCopy.GetIdent() != "tape" {
		t.Error(tapeCopy.GetIdent())
	}
	// with meta decoupled, we should be able to navigate tape meta w/o "tape-cells" group
	test := &meta.Module{Ident:"test"}
	test.AddMeta(tapeCopy)
	var actualBytes bytes.Buffer
	err := NewContext().Selector(SelectModule(test, true)).InsertInto(NewJsonWriter(&actualBytes).Node()).LastErr
	if err != nil {
		t.Error(err)
	}
	t.Log(actualBytes.String())
}