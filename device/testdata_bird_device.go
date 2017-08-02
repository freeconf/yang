package device

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/nodes"
)

func bird(json string) *Local {
	ypath := meta.PathStreamSource("../nodes:../yang")
	d := New(ypath)
	birds := make(map[string]*nodes.Bird)
	d.Add("testdata-bird", nodes.BirdModule(birds))
	b, err := d.Browser("testdata-bird")
	if err != nil {
		panic(err)
	}
	if json != "" {
		if err := b.Root().UpsertFrom(nodes.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return d
}
