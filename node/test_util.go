package node
import (
	"github.com/c2g/meta"
	"github.com/c2g/meta/yang"
	"bytes"
)

type Testing interface {
	Fatal(args ...interface{})
	Errorf(format string, args ...interface{})
}

type ModuleTestSetup struct {
	Module *meta.Module
	Store *BufferStore
	Data *StoreData
}

func ModuleSetup(mstr string, t Testing) (setup *ModuleTestSetup) {
	setup = &ModuleTestSetup{}
	var err error
	setup.Module, err = yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	setup.Store = NewBufferStore()
	setup.Data = NewStoreData(setup.Module, setup.Store)
	return
}

func (setup *ModuleTestSetup) ToString(t Testing) string {
	var actualBuff bytes.Buffer
	err := setup.Data.Browser().Root().Selector().InsertInto(NewJsonWriter(&actualBuff).Node()).LastErr
	if err != nil {
		t.Fatal(err)
	}
	return actualBuff.String()
}

func AssertStrEqual(t Testing, expected string, actual string) bool {
	if expected != actual {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
		return false
	}
	return true
}


