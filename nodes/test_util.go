package nodes

import (
	"testing"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

type Bird struct {
	Name     string
	Wingspan int
	Species  *Species
}

type Species struct {
	Name  string
	Class string
}

var birdYangStr = `
module testdata-bird {
    prefix "";
    namespace "";
    revision 0;

    list bird {
        key "name";
        leaf name {
            type string;
        }
        leaf wingspan {
            type int32;
        }
    }
}
`
var TestYangPath = &meta.StringSource{
	Streamer: func(resource string) (string, error) {
		switch resource {
		case "testdata-bird":
			return birdYangStr, nil
		}
		return "", nil
	},
}

func BirdBrowser(json string) (*node.Browser, map[string]*Bird) {
	data := make(map[string]*Bird)
	m := yang.RequireModule(TestYangPath, "testdata-bird")
	b := node.NewBrowser(m, BirdModule(data))
	if json != "" {
		if err := b.Root().UpsertFrom(ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return b, data
}

func BirdModule(birds map[string]*Bird) node.Node {
	return &Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "bird":
				return ReflectList(birds), nil
			}
			return nil, nil
		},
	}
}

func assertEq(t *testing.T, a interface{}, b interface{}) bool {
	eq := a == b
	if !eq {
		t.Errorf("%v != %v", a, b)
	}
	return eq
}
