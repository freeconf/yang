package browse

import (
	"bytes"
	"github.com/blitter/meta"
	"github.com/blitter/meta/yang"
	"github.com/blitter/node"
	"os"
	"strings"
	"testing"
)

func TestInlineSave(t *testing.T) {
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
	actual := actualBytes.String()
	expectedFragement := `"data":{"hobbies":{"birding":{"favorite-species":"towhee"}}}`
	if !strings.Contains(actual, expectedFragement) {
		t.Error(actual)
	}
}

func TestInlineLoadContainer(t *testing.T) {
	inlineData := `
{
  "meta": {
    "container" : {
    	    "ident" : "hobbies",
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
  },
  "data":{
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
	if loadErr := <-waitForSchemaLoad; loadErr != nil {
		t.Fatal(loadErr)
	}
	expected := `{"hobbies":{"birding":{"favorite-species":"towhee"}}}`
	actual := actualBytes.String()
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
	//c.Selector(node.SelectModule(inline.Module, true)).InsertInto(node.NewJsonWriter(os.Stdout).Node())
}

func TestInlineLoadListWhole(t *testing.T) {
	inlineData := `
{
  "meta": {
    "list" : {
      "ident" : "hello",
      "definitions": [
        {
          "ident": "hobby",
          "leaf": {
            "ident": "hobby",
            "type": {
              "ident": "string"
            }
          }
        }
      ]
    }
  },
  "data" : {
     "hello" : [{
        "hobby":"birding"
      }]
  }
}
`
	inlineNode := node.NewJsonReader(strings.NewReader(inlineData)).Node()
	node.Dump(inlineNode, os.Stdout)
	var actualBytes bytes.Buffer
	c := node.NewContext()
	inline := NewInline()
	if err := inline.Load(c, inlineNode, node.NewJsonWriter(&actualBytes).Node(), nil); err != nil {
		t.Fatal(err)
	}
	expected := `{"hello":[{"hobby":"birding"}]}`
	actual := actualBytes.String()
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestInlineLoadListItem(t *testing.T) {
	inlineData := `
{
  "meta": {
    "list-item" : {
      "ident" : "hello",
      "definitions": [
        {
          "ident": "hobby",
          "leaf": {
            "ident": "hobby",
            "type": {
              "ident": "string"
            }
          }
        }
      ]
    }
  },
  "data" : {
      "hello": [{
        "hobby":"birding"
      }]
  }
}
`
	inlineNode := node.NewJsonReader(strings.NewReader(inlineData)).Node()
	node.Dump(inlineNode, os.Stdout)
	var actualBytes bytes.Buffer
	c := node.NewContext()
	inline := NewInline()
	if err := inline.Load(c, inlineNode, node.NewJsonWriter(&actualBytes).Node(), nil); err != nil {
		t.Fatal(err)
	}
	expected := `{"hobby":"birding"}`
	actual := actualBytes.String()
	if actual != expected {
		t.Errorf("\nExpected:%s\n  Actual:%s", expected, actual)
	}
}

func TestInlineSaveListItem(t *testing.T) {
	moduleStr := `
module my-module {
	prefix "t";
	namespace "t";
	revision 0;
	list hobbies {
		key "name";
		leaf name {
			type string;
		}
		container favorite {
			leaf label {
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
	"hobbies" : [{
		"name" : "birding",
		 "favorite" : {
			"label" : "towhee"
		}
	}]
}`
	n := node.NewJsonReader(strings.NewReader(nodeSlice)).Node()
	c := node.NewContext()
	inline := NewInline()
	sel := c.Select(m, n).Find("hobbies=birding")
	if sel.LastErr != nil {
		t.Fatal(sel.LastErr)
	}
	var fragment bytes.Buffer
	fragmentNode := inline.SaveSelection(c, sel.Selection)
	err = c.Select(inline.Module, fragmentNode).InsertInto(node.NewJsonWriter(&fragment).Node()).LastErr
	if err != nil {
		t.Fatal(err)
	}
	restore := NewInline()
	var actualBytes bytes.Buffer
	restoreNode := node.NewJsonReader(&fragment).Node()
	if err := restore.Load(c, restoreNode, node.NewJsonWriter(&actualBytes).Node(), nil); err != nil {
		t.Fatal(err)
	}
	expected := `{"name":"birding","favorite":{"label":"towhee"}}`
	actual := actualBytes.String()
	if actual != expected {
		t.Errorf("Expected:%s\n  Actual:%s", expected, actual)
	}
}
