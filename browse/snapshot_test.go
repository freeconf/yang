package browse

import (
	"testing"
	"github.com/c2g/meta/yang"
	"github.com/c2g/node"
	"strings"
	"bytes"
	"fmt"
	"os"
)

func TestSnapshotRestore(t *testing.T) {

	tests := []struct {
		snapshot string
		expected string
	}{
		{
			snapshot : `
{
  "meta": {
    "definitions": [{
      "ident" : "hobbies",
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
                    "type": {
                      "ident": "string"
                    }
                  }
                }
              ]
            }
          }
        ]
      }
    }]
  },
  "data": {
    "hobbies": {
      "birding": {
        "favorite-species": "towhee"
      }
    }
  }
}`,
			expected : `{"hobbies":{"birding":{"favorite-species":"towhee"}}}`,
		},
		{
			snapshot: `
{
  "meta": {
    "definitions":[{
      "ident":"hobbies",
      "list": {
        "ident": "hobbies",
        "key":["name"],
        "definitions": [
          {
            "ident": "name",
            "leaf": {
              "ident": "name",
              "type": {
                "ident": "string"
              }
            }
          },
          {
            "ident": "favorite",
            "container": {
              "ident": "favorite",
              "definitions": [
                {
                  "ident": "label",
                  "leaf": {
                    "ident": "label",
                    "type": {
                      "ident": "string"
                    }
                  }
                }
              ]
            }
          }
        ]
      }
    }]
  },
  "data": {
    "hobbies": [
      {
        "name": "birding",
        "favorite": {
          "label": "towhee"
        }
      }
    ]
  }
}`,
			expected: `{"hobbies":[{"name":"birding","favorite":{"label":"towhee"}}]}`,
		},
	}

	c := node.NewContext()
	for i, test := range tests {
		in := node.NewJsonReader(strings.NewReader(test.snapshot)).Node()
		snap, err := RestoreSelection(c, in, nil)
		if err != nil {
			t.Errorf("#%d - %s", i, err.Error())
			continue
		}
		var actualBytes bytes.Buffer
		if err = c.Selector(snap).InsertInto(node.NewJsonWriter(&actualBytes).Node()).LastErr; err != nil {
			t.Errorf("#%d - %s", i, err.Error())
			continue
		}
		actual := actualBytes.String()
		if actual != test.expected {
			t.Errorf("#%d - %s", i, actual)
		}
	}
}

func TestSnapshotSave(t *testing.T) {
	moduleStr := `
module test {
	prefix "t";
	namespace "t";
	revision 0;

        %s

	container hockey {
		leaf favorite-team {
			type string;
		}
	}
}`
	tests := []struct {
		yang string
		data string
		url string
		expected string
		roundtrip string // from the perspective of the test.url
	}{
		{
			yang :
				`
					container hobbies {
						container birding {
							leaf favorite-species {
								type string;
							}
						}
					}
				`,
			data :
				`{
					"hobbies" : {
						"birding" : {
							"favorite-species" : "towhee"
						}
					}
				}`,
			url : "hobbies",
			expected :
				`"data":{"birding":{"favorite-species":"towhee"}}`,
			roundtrip :
				`{"birding":{"favorite-species":"towhee"}}`,

		},
		{
			yang :
			`
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
			`,
			data :
			`{
				"hobbies" : [{
					"name" : "birding",
					"favorite" : {
						"label" : "towhee"
					}
				}]
			}`,
			url : "hobbies",
			expected :
				`"data":{"hobbies":[{"name":"birding","favorite":{"label":"towhee"}}]}}`,
			roundtrip:
				`{"hobbies":[{"name":"birding","favorite":{"label":"towhee"}}]}`,
		},
		{
			yang :
			`
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
			`,
			data :
			`{
				"hobbies" : [{
					"name" : "birding",
					"favorite" : {
						"label" : "towhee"
					}
				}]
			}`,
			url : "hobbies=birding",
			expected :
				`"data":{"name":"birding","favorite":{"label":"towhee"}}`,
			roundtrip:
				`{"name":"birding","favorite":{"label":"towhee"}}`,
		},
	}
	for i, test := range tests {
		mstr := fmt.Sprintf(moduleStr, test.yang)
		mod, err := yang.LoadModuleCustomImport(mstr, nil)
		if err != nil {
			panic(err)
		}
		n := node.NewJsonReader(strings.NewReader(test.data)).Node()
		c := node.NewContext()
		sel := c.Select(mod, n).Find(test.url)
		if sel.LastErr != nil {
			t.Error("#%d - %s", i, sel.LastErr.Error())
			continue
		}
		snap := SaveSelection(sel.Selection)
		var actualBytes bytes.Buffer
		if err = c.Selector(snap).InsertInto(node.NewJsonWriter(&actualBytes).Node()).LastErr; err != nil {
			t.Errorf("#%d - %s", i, err.Error())
			continue
		}
		actual := actualBytes.String()
		if !strings.Contains(actual, test.expected) {
			t.Errorf("#%d - %s", i, actual)
			continue
		}

		roundtrip, rtErr := RestoreSelection(c, node.NewJsonReader(&actualBytes).Node(), nil)
		if rtErr != nil {
			t.Errorf("#%d roundtrip - %s", i, rtErr.Error())
			continue
		}
		var roundtripBytes bytes.Buffer
		if restoreErr := c.Selector(roundtrip).InsertInto(node.NewJsonWriter(&roundtripBytes).Node()).LastErr; restoreErr != nil {
			t.Errorf("#%d roundtrip restore - %s", i, restoreErr.Error())
			continue
		}
		roundtripActual := roundtripBytes.String()
		if roundtripActual != test.roundtrip {
			t.Errorf("#%d roundtrip wrong expectation. actual:%s", i, roundtripActual)
			continue
		}
	}
}

// Disabled becuase requires running server
func _TestSnapshotMetaDownload(t *testing.T) {
	data := `
{
  "meta": {
    "import" : [
      {"list": "http://localhost:8009/meta/module/definitions=records/list?userToken=api:5"}
    ]
  },
  "data": {
    "records": [
      {
        "_id": "2101242312321",
        "firstName": "Charles",
        "lastName": "Abany",
        "userToken": "1024608925",
        "address": {
          "full": "250 Baldwin AV",
          "city": "Framingham",
          "state": "MA",
          "zip": "01701"
        }
      },
      {
        "_id": "2101242312321",
        "firstName": "Isabelle",
        "lastName": "Abany",
        "userToken": "1072730327",
        "address": {
          "full": "250 Baldwin AV",
          "city": "Framingham",
          "state": "MA",
          "zip": "01701"
        }
      }
    ]
  }
}`
	c := node.NewContext()
	s, err := RestoreSelection(c, node.NewJsonReader(strings.NewReader(data)).Node(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if err = c.Selector(s).InsertInto(node.NewJsonWriter(os.Stdout).Node()).LastErr; err != nil {
		t.Error(err)
	}
}

