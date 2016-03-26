package browse

import (
	"testing"
	"meta/yang"
	"meta"
	"node"
	"strings"
	"bytes"
	"os"
)



func TestInlineCreate(t *testing.T) {
	moduleStr := `
module my-module {
	prefix "t";
	namespace "t";
	revision 0;
	container hobbies {
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
	}
}`
	m, err := yang.LoadModuleCustomImport(moduleStr, nil)
	if err != nil {
		panic(err)
	}
	nodeSlice := `{
	"hobbies" : {
		"birding" : {
			"favorite-species" : "towhee"
		}
	}
}`
	n := node.NewJsonReader(strings.NewReader(nodeSlice)).Node()
	c := node.NewContext()
	inline := NewInline()
	onWriteData := inline.Save(c, meta.FindByIdent2(m, "hobbies").(meta.MetaList), n)
	var actualBytes bytes.Buffer
	if toErr := c.Select(inline.Module, onWriteData).InsertInto(node.NewJsonWriter(&actualBytes).Node()).LastErr; toErr != nil {
		t.Fatal(toErr)
	}
	c.Selector(node.SelectModule(inline.Module, true)).InsertInto(node.NewJsonWriter(os.Stdout).Node())
	t.Log(actualBytes.String())
}


func TestInlineReincarnate(t *testing.T) {
	inlineData := `
{
  "meta": {
    "definitions": [
      {
        "ident": "hobbies",
        "container": {
          "ident": "hobbies",
          "definitions": [
            {
              "ident": "birding",
              "container": {
                "ident": "birding",
                "definitions": [
                  {
                    "ident": "favorite-species",
                    "leaf": {
                      "ident": "favorite-species",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  }
                ]
              }
            },
            {
              "ident": "hockey",
              "container": {
                "ident": "hockey",
                "definitions": [
                  {
                    "ident": "favorite-team",
                    "leaf": {
                      "ident": "favorite-team",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  }
                ]
              }
            }
          ]
        }
      }
    ]
  },
  "data" : {
    "hobbies": {
      "birding": {
        "favorite-species": "towhee"
      }
    }
  }
}
`
	inlineNode := node.NewJsonReader(strings.NewReader(inlineData)).Node()
	var actualBytes bytes.Buffer
	c := node.NewContext()
	inline := NewInline()
	waitForSchemaLoad := make(chan error)
	go func() {
		defer close(waitForSchemaLoad)
		err := inline.Load(c, inlineNode, node.NewJsonWriter(&actualBytes).Node(), waitForSchemaLoad)
		if err != nil {
			waitForSchemaLoad <- err
		}
	}()
	if loadErr := <- waitForSchemaLoad; loadErr != nil {
		t.Fatal(loadErr)
	}
	expected := `{"hobbies":{"birding":{"favorite-species":"towhee"}}}`
	actual := actualBytes.String()
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
	//c.Selector(node.SelectModule(inline.Module, true)).InsertInto(node.NewJsonWriter(os.Stdout).Node())
}