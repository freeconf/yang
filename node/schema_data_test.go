package node

import (
	"bytes"
	"fmt"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"testing"
	"flag"
	"github.com/c2stack/c2g/c2"
	"io/ioutil"
)

var updateFlag = flag.Bool("update", false, "Update the golden files.")

func printMeta(m meta.Meta, level string) {
	fmt.Printf("%s%s\n", level, m.GetIdent())
	if nest, isNest := m.(meta.MetaList); isNest {
		if len(level) >= 16 {
			panic("Max level reached")
		}
		i2 := meta.NewMetaListIterator(nest, false)
		for i2.HasNextMeta() {
			printMeta(i2.NextMeta(), level+"  ")
		}
	}
}

func TestYangBrowse(t *testing.T) {
	moduleStr := `
module json-test {
	prefix "t";
	namespace "t";
	revision 0;
	list hobbies {
		container birding {
			leaf favorite-species {
				type string;
			}
		}
		container hockey {
			leaf favorite-team {
				type string;
			}
		}
		container metric {
			config "false";
			leaf v {
				type string;
			}
		}
	}
	action foo {
	  input {
	  	leaf a {
	  	   type string;
	  	}
	  }
	  output {
	  	leaf b {
	  	   type string;
	  	}
	  }
	}
	notification n {
	  leaf-list ll {
	    type int32;
	  }
	}
}`
	m, err := yang.LoadModuleCustomImport(moduleStr, nil)
	if err != nil {
		t.Fatal("bad module", err)
	}
	var actual bytes.Buffer
	if err = SelectModule(yang.YangPath(), m, false).Root().Selector().InsertInto(NewJsonPretty(&actual).Node()).LastErr; err != nil {
		t.Error(err)
	}
	goldenFile := "testdata/schema_data_test-TestYangBrowse.json"
	if *updateFlag {
		if err := ioutil.WriteFile(goldenFile, actual.Bytes(), 0644); err != nil {
			panic(err.Error())
		}
	}
	if err := c2.Diff2(goldenFile, actual.Bytes()); err != nil {
		t.Error(err)
	}
}
//	read := &meta.Module{Ident:"read"}
//	if err = SelectModule(read, false).Root().Selector().UpsertFrom(NewJsonReader(&actual).Node()).LastErr; err != nil {
//		t.Error(err)
//	}
//	var roundtrip bytes.Buffer
//	if err = SelectModule(read, false).Root().Selector().InsertInto(NewJsonWriter(&roundtrip).Node()).LastErr; err != nil {
//		t.Error(err)
//	} else {
//		t.Log("Round Trip:\n", string(roundtrip.Bytes()))
//	}
//}

// TODO: support typedefs - simpleyang datatypes that use typedefs return format=0
func TestYangWrite(t *testing.T) {
	simple, err := yang.LoadModuleCustomImport(yang.TestDataSimpleYang, nil)
	if err != nil {
		t.Fatal(err)
	}
	copy := &meta.Module{}
	yangPath := yang.YangPath()
	from := SelectModule(yangPath, simple, false).Root()
	to := SelectModule(yangPath, copy, false).Root()
	err = from.Selector().UpsertInto(to.Node()).LastErr
	if err != nil {
		t.Fatal(err)
	}

	// dump original and clone to see if anything is missing
	diff := Diff(from.Node(), to.Node())
	var out bytes.Buffer
	from.Fork(diff).Selector().InsertInto(NewJsonWriter(&out).Node())
	t.Log(out.String())
}
