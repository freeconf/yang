package testdata

import (
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodes"
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

var YangPath = source.Path("../testdata:../yang")

func BirdBrowser(json string) (*node.Browser, map[string]*Bird) {
	data := make(map[string]*Bird)
	b := node.NewBrowser(BirdModule(), BirdNode(data))
	if json != "" {
		if err := b.Root().UpsertFrom(nodes.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return b, data
}

func BirdModule() *meta.Module {
	return parser.RequireModule(YangPath, "bird")
}

func BirdNode(birds map[string]*Bird) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "bird":
				return nodes.ReflectList(birds), nil
			}
			return nil, nil
		},
	}
}
