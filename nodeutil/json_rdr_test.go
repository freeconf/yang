package nodeutil

import (
	"strings"
	"testing"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
	"github.com/freeconf/yang/val"
)

func TestJsonWalk(t *testing.T) {
	moduleStr := `
module json-test {
	prefix "t";
	namespace "t";
	revision 0;
	list hobbies {
		key "name";
	    leaf name {
	      type string;
	    }
		container favorite {
		    leaf common-name {
		      type string;
		    }
		    leaf location {
				type string;
		    }
		}
	}
}
	`
	module, err := parser.LoadModuleFromString(nil, moduleStr)
	if err != nil {
		t.Fatal(err)
	}
	json := `{"hobbies":[
{"name":"birding", "favorite": {"common-name" : "towhee", "extra":"double-mint", "location":"out back"}},
{"name":"hockey", "favorite": {"common-name" : "bruins", "location" : "Boston"}}
]}`
	tests := []string{
		"hobbies",
		"hobbies=birding",
		"hobbies=birding/favorite",
	}
	for _, test := range tests {
		n, err := ReadJSON(json)
		fc.AssertEqual(t, nil, err)
		sel := node.NewBrowser(module, n).Root()
		found, err := sel.Find(test)
		fc.RequireEqual(t, nil, err, "failed to transmit json")
		fc.RequireEqual(t, true, found != nil, "target not found")
		fc.AssertEqual(t, "json-test/"+test, found.Path.String())
	}
}

func TestJsonRdrUnion(t *testing.T) {
	mstr := `
	module x {
		revision 0;
		leaf y {
			type union {
				type int32;
				type string;
			}
		}
	}
		`
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{in: `{"y":24}`, out: `{"y":24}`},
		{in: `{"y":"hi"}`, out: `{"y":"hi"}`},
	}
	for _, json := range tests {
		t.Log(json.in)
		n, err := ReadJSON(json.in)
		fc.AssertEqual(t, nil, err)
		actual, err := WriteJSON(node.NewBrowser(m, n).Root())
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, json.out, actual)
	}
}

func TestJsonRdrTypedefUnionList(t *testing.T) {
	mstr := `
    module x {
        revision 0;
        typedef ip-prefix {
            type union {
                type string;
            }
        }
        leaf-list ip {
            type ip-prefix;
        }
    }
        `
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  `{"ip":["10.0.0.1","10.0.0.2"]}`,
			out: `{"ip":["10.0.0.1","10.0.0.2"]}`,
		},
	}
	for _, json := range tests {
		t.Log(json.in)
		n, err := ReadJSON(json.in)
		fc.AssertEqual(t, nil, err)
		actual, err := WriteJSON(node.NewBrowser(m, n).Root())
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, json.out, actual)
	}
}

func TestJsonRdrTypedefMixedUnion(t *testing.T) {
	mstr := `
    module x {
        revision 0;
        typedef ip-prefix {
            type union {
                type string;
            }
        }
        leaf-list ips {
            type ip-prefix;
        }
		leaf ip {
			type ip-prefix;
		}
    }
        `
	m, err := parser.LoadModuleFromString(nil, mstr)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  `{"ips":["10.0.0.1","10.0.0.2"],"ip":"10.0.0.3"}`,
			out: `{"ips":["10.0.0.1","10.0.0.2"],"ip":"10.0.0.3"}`,
		},
	}
	for _, json := range tests {
		t.Log(json.in)
		n, err := ReadJSON(json.in)
		fc.AssertEqual(t, nil, err)
		actual, err := WriteJSON(node.NewBrowser(m, n).Root())
		if err != nil {
			t.Error(err)
		}
		fc.AssertEqual(t, json.out, actual)
	}
}

func TestNumberParse(t *testing.T) {
	moduleStr := `
module json-test {
	prefix "t";
	namespace "t";
	revision 0;
	container data {
		leaf id {
			type int32;
		}
		leaf idstr {
			type int32;
		}
		leaf idstrwrong {
			type int32;
		}
		leaf-list readings{
			type decimal64;
		}
	}
}
	`
	module, err := parser.LoadModuleFromString(nil, moduleStr)
	if err != nil {
		t.Fatal(err)
	}
	n, err := ReadJSON(`{ "data": {
			"id": 4,
			"idstr": "4",
			"idstrwrong": "4s",
			"readings": [
				"3.555454",
				"45.04545",
				324545.04
			]
		}
	}`)
	fc.AssertEqual(t, nil, err)

	root := node.NewBrowser(module, n).Root()

	//test get id
	sel, err := root.Find("data/id")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, sel != nil)
	v, err := sel.Get()
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 4, v.Value())

	//test get idstr
	sel, err = root.Find("data/idstr")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, sel != nil)
	v, err = sel.Get()
	fc.RequireEqual(t, nil, err)
	fc.AssertEqual(t, 4, v.Value())

	//test idstrwrong fail
	sel, err = root.Find("data/idstrwrong")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, sel != nil)
	_, err = sel.Get()
	fc.RequireEqual(t, true, err != nil, "Failed to throw error on invalid input")

	sel, err = root.Find("data/readings")
	fc.RequireEqual(t, nil, err)
	fc.RequireEqual(t, true, sel != nil)
	v, err = sel.Get()
	fc.RequireEqual(t, nil, err)
	expected := []float64{3.555454, 45.04545, 324545.04}
	fc.AssertEqual(t, expected, v.Value())
}

func TestJsonEmpty(t *testing.T) {
	moduleStr := `
module json-test {
	leaf x {
		type empty;
	}
}
	`
	m, err := parser.LoadModuleFromString(nil, moduleStr)
	fc.AssertEqual(t, nil, err)
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	n, err := ReadJSON(`{"x":{}}`)
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, nil, b.Root().InsertFrom(n))
	fc.AssertEqual(t, val.NotEmpty, actual["x"])
}

func TestReadQualifiedJsonIdentRef(t *testing.T) {
	ypath := source.Dir("./testdata")
	m := parser.RequireModule(ypath, "module-test")
	n, err := ReadJSON(`{
		"module-test:type":"module-types:derived-type",
		"module-test:type2":"local-type"
	}`)
	fc.AssertEqual(t, nil, err)
	actual := make(map[string]interface{})
	b := node.NewBrowser(m, ReflectChild(actual))
	fc.AssertEqual(t, nil, b.Root().InsertFrom(n))
	fc.AssertEqual(t, "derived-type", actual["type"].(val.IdentRef).Label)
	fc.AssertEqual(t, "local-type", actual["type2"].(val.IdentRef).Label)
}

func TestValidateHappyCase(t *testing.T) {
	mstring := `
	module x {
		revision 0;
		container c {
			leaf l1 {
				type int32;
				mandatory true;
			}
			leaf l2 {
				type int32;
			}
		}
		list l {
			leaf l1 {
				type int32;
				mandatory true;
			}
			leaf l2 {
				type int32;
			}
		}
	}`
	payload := `
	{
		"c": {
			"l1": 1
		},
		"l": [
			{"l1": 1, "l2": 2},
			{"l1": 1}
		]
	}`
	module, err := parser.LoadModuleFromString(nil, mstring)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(payload)

	n, err := ReadJSON(payload)
	fc.AssertEqual(t, nil, err)
	selection := node.NewBrowser(module, n).Root()

	reader := JSONRdr{In: strings.NewReader(payload)}
	if err := reader.Validate(selection); err != nil {
		t.Errorf("validation should pass, but got error: %s", err)
	}
}

func TestValidateForInvalidPayloads(t *testing.T) {
	tests := []struct{
		mstring     string
		payload     string
		msg         string
		expectedErr string
	}{
		{
			mstring: `
			module x {
				revision 0;
				leaf l {
					type int32;
					mandatory true;
				}
			}`,
			payload: `{}`,
			msg: "should fail when mandatory container child is missing",
			expectedErr: "missing mandatory node: x/l",
		},
		{
			mstring: `
			module x {
				revision 0;
				container c {
					leaf l1 {
						type string;
					}
				}
			}`,
			payload: `{"c": {"l1": 1, "extra": 3}}`,
			msg: "should fail on unexpected container child",
			expectedErr: "unexpected node: x/c/extra",
		},
		{
			mstring: `
			module x {
				revision 0;
				list l {
					leaf l1 {
						type string;
						mandatory true;
					}
				}
			}`,
			payload: `{"l": [{}]}`,
			msg: "should fail when mandatory list child is missing",
			expectedErr: "missing mandatory node: x/l/l1",
		},
		{
			mstring: `
			module x {
				revision 0;
				list l {
					leaf l1 {
						type string;
						mandatory true;
					}
				}
			}`,
			payload: `{"l": [{"l1": "foo", "extra": 1}]}`,
			msg: "should fail on unexpected list child",
			expectedErr: "unexpected node: x/l/extra",
		},
	}

	for _, test := range tests {
		module, err := parser.LoadModuleFromString(nil, test.mstring)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(test.payload)

		n, err := ReadJSON(test.payload)
		fc.AssertEqual(t, nil, err)
		selection := node.NewBrowser(module, n).Root()

		reader := JSONRdr{In: strings.NewReader(test.payload)}
		if err := reader.Validate(selection); err != nil {
			fc.AssertEqual(t, strings.Contains(err.Error(), test.expectedErr), true, "unexpected error")
		} else {
			t.Errorf(test.msg)
		}
	}
}
