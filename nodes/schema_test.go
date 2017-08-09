package nodes

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
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
			m, _ := i2.NextMeta()
			printMeta(m, level+"  ")
		}
	}
}

func Test_SchemaRead(t *testing.T) {
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
	sel := Schema(m, false).Root()
	actual, err := WritePrettyJSON(sel)
	if err != nil {
		t.Error(err)
	}
	goldenFile := "testdata/schema_data_test-TestYangBrowse.json"
	if *updateFlag {
		if err := ioutil.WriteFile(goldenFile, []byte(actual), 0644); err != nil {
			panic(err.Error())
		}
	}
	if err := c2.Diff2(goldenFile, []byte(actual)); err != nil {
		t.Error(err)
	}
}

// TODO: support typedefs - simpleyang datatypes that use typedefs return format=0
func Test_SchemaWrite(t *testing.T) {
	simple, err := yang.LoadModuleCustomImport(yang.TestDataSimpleYang, nil)
	if err != nil {
		t.Fatal(err)
	}
	copy := &meta.Module{}
	from := Schema(simple, false).Root()
	to := Schema(copy, false).Root()
	err = from.UpsertInto(to.Node).LastErr
	if err != nil {
		t.Fatal(err)
	}

	// dump original and clone to see if anything is missing
	diff := Diff(from.Node, to.Node)
	if out, err := WriteJSON(from.Split(diff)); err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}
