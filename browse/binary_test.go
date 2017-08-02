package browse

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func _TestBinaryBrowser(t *testing.T) {
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
	tests := []string{
		`{"c":"hello"}`,
		`{"a":{"b":{"s":"waldo","b":true,"i":99,"l":100,"d":1.5,"e":"one"}}}`,
		`{"a":{"b":{"sl":["waldo"],"bl":[true],"il":[99,100],"ll":[100,101],"dl":[1.5,2.5],"el":["one","two"]}}}`,
		`{"p":[{"k":"walter"}]}`,
		`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo","r":[{"z":99}]}]}`,
	}
	for _, test := range tests {
		var buff bytes.Buffer
		w := NewBinaryWriter(&buff)
		if err = node.NewBrowser(m, w.Node()).Root().InsertFrom(nodes.ReadJSON(test)).LastErr; err != nil {
			t.Error(err)
		}
		original := "\n" + hex.Dump(buff.Bytes())
		t.Log(original)
		r := NewBinaryReader(&buff)
		sel := node.NewBrowser(m, r.Node()).Root()
		if actual, err := nodes.WriteJSON(sel); err != nil {
			t.Log("\n" + hex.Dump(buff.Bytes()))
			t.Error(err)
		} else if test != actual {
			t.Log("\noriginal:\n%s", original)
			t.Errorf("\nExpected:%s\n  Actual:%s", test, actual)
		}
	}
}
