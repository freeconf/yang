package browse

import (
	"testing"
	"github.com/c2g/meta/yang"
	"strings"
	"bytes"
	"encoding/hex"
	"github.com/c2g/node"
)

func TestBinaryBrowser(t *testing.T) {
	mstr := `
module m {
	namespace "";
	prefix "";
	revision 0;
	leaf c {
		type string;
	}
	container a {
		container b {
			leaf s {
				type string;
			}
			leaf b {
				type boolean;
			}
			leaf i {
				type int32;
			}
			leaf l {
				type int64;
			}
			leaf d {
				type decimal64;
			}
			leaf e {
				type enumeration {
					enum one;
					enum two;
				}
			}
			leaf-list sl {
				type string;
			}
			leaf-list bl {
				type boolean;
			}
			leaf-list il {
				type int32;
			}
			leaf-list ll {
				type int64;
			}
			leaf-list dl {
				type decimal64;
			}
			leaf-list el {
				type enumeration {
					enum one;
					enum two;
				}
			}
		}
	}
	list p {
		key "k";
		leaf k {
			type string;
		}
		container q {
			leaf s {
				type string;
			}
		}
		list r {
			leaf z {
				type int32;
			}
		}
	}
}
`
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []string {
		`{"c":"hello"}`,
		`{"a":{"b":{"s":"waldo","b":true,"i":99,"l":100,"d":1.5,"e":"one"}}}`,
		`{"a":{"b":{"sl":["waldo"],"bl":[true],"il":[99,100],"ll":[100,101],"dl":[1.5,2.5],"el":["one","two"]}}}`,
		`{"p":[{"k":"walter"}]}`,
		`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo","r":[{"z":99}]}]}`,
	}
	c := node.NewContext()
	for _, test := range tests {
		var buff bytes.Buffer
		w := NewBinaryWriter(&buff)
		if err = c.Select(m, w.Node()).InsertFrom(node.NewJsonReader(strings.NewReader(test)).Node()).LastErr; err != nil {
			t.Error(err)
		}
		original := "\n" + hex.Dump(buff.Bytes())
		r := NewBinaryReader(&buff)
		var actualBuff bytes.Buffer
		if err = c.Select(m, r.Node()).InsertInto(node.NewJsonWriter(&actualBuff).Node()).LastErr; err != nil {
			t.Log("\n" + hex.Dump(buff.Bytes()))
			t.Error(err)
		}
		actual := actualBuff.String()
		if test != actual {
			t.Log("\noriginal:\n%s", original)
			t.Errorf("\nExpected:%s\n  Actual:%s", test, actual)
		}
	}
}
