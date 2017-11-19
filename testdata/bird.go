package testdata

import (
	"github.com/c2stack/c2g/device"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
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

var YangPath = meta.PathStreamSource("../testdata:../yang")

func BirdDevice(json string) (*device.Local, map[string]*Bird) {
	d := device.New(YangPath)
	b, birds := BirdBrowser(json)
	d.AddBrowser(b)
	if json != "" {
		if err := b.Root().UpsertFrom(nodes.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return d, birds
}

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
	return yang.RequireModule(YangPath, "bird")
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
