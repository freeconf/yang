package device

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

func bird(json string) *Local {
	ypath := meta.PathStreamSource("../node:../yang")
	d := New(ypath)
	birds := make(map[string]*node.Bird)
	d.Add("testdata-bird", node.BirdModule(birds))
	b, err := d.Browser("testdata-bird")
	if err != nil {
		panic(err)
	}
	if json != "" {
		if err := b.Root().UpsertFrom(node.ReadJson(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return d
}
