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

var yangDir = "../testdata:../yang"

func BirdDevice(json string) (*device.Local, map[string]*Bird) {
	ypath := meta.PathStreamSource(yangDir)
	yang.RequireModule(ypath, "bird")
	d := device.New(ypath)
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
	ypath := meta.PathStreamSource(yangDir)
	data := make(map[string]*Bird)
	m := yang.RequireModule(ypath, "bird")
	b := node.NewBrowser(m, BirdModule(data))
	if json != "" {
		if err := b.Root().UpsertFrom(nodes.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return b, data
}

func BirdModule(birds map[string]*Bird) node.Node {
	return &nodes.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "bird":
				return nodes.ReflectList(birds), nil
			}
			return nil, nil
		},
	}
}
