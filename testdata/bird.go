package testdata

import (
	"net/netip"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
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

type IPBird struct {
	Name     netip.Addr
	Wingspan int
	Species  *IPSpecies
}

type IPSpecies struct {
	Name  netip.Addr
	Class string
}

var YangPath = source.Path("../testdata:../yang")

func BirdBrowser(json string) (*node.Browser, map[string]*Bird) {
	data := make(map[string]*Bird)
	b := node.NewBrowser(BirdModule(), BirdNode(data))
	if json != "" {
		if err := b.Root().UpsertFrom(nodeutil.ReadJSON(json)); err != nil {
			panic(err)
		}
	}
	return b, data
}

func BirdModule() *meta.Module {
	return parser.RequireModule(YangPath, "bird")
}

func BirdNode(birds map[string]*Bird) node.Node {
	return &nodeutil.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "bird":
				return nodeutil.ReflectList(birds), nil
			}
			return nil, nil
		},
	}
}
