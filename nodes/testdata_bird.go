package nodes

import (
	"reflect"
	"strings"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

type Bird struct {
	Name     string
	Wingspan int
}

func BirdBrowser(nodeDir string, json string) (*node.Browser, map[string]*Bird) {
	yangPath := &meta.FileStreamSource{Root: nodeDir}
	data := make(map[string]*Bird)
	m := yang.RequireModule(yangPath, "testdata-bird")
	b := node.NewBrowser(m, BirdModule(data))
	if json != "" {
		if err := b.Root().UpsertFrom(ReadJson(json)).LastErr; err != nil {
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
				return BirdList(birds), nil
			}
			return nil, nil
		},
	}
}

func BirdList(birds map[string]*Bird) node.Node {
	index := node.NewIndex(birds)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &Basic{
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			var b *Bird
			key := r.Key
			if r.New {
				b = &Bird{Name: key[0].String()}
				birds[b.Name] = b
			} else if key != nil {
				if r.Delete {
					delete(birds, key[0].String())
				} else {
					b = birds[key[0].String()]
				}
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					name := v.String()
					if b = birds[name]; b != nil {
						var err error
						if key, err = node.NewValues(r.Meta.KeyMeta(), name); err != nil {
							return nil, nil, err
						}
					}
				}
			}
			if b != nil {
				return BirdNode(b), key, nil
			}
			return nil, nil, nil
		},
	}
}

func BirdNode(b *Bird) node.Node {
	return ReflectNode(b)
}
