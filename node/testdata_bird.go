package node

import (
	"reflect"
	"strings"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
)

type Bird struct {
	Name     string
	Wingspan int
}

func birds(data map[string]*Bird, json string) *Browser {
	m := yang.RequireModule(&meta.FileStreamSource{Root: "."}, "testdata-bird")
	b := NewBrowser(m, BirdModule(data))
	if json != "" {
		if err := b.Root().UpsertFrom(ReadJson(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return b
}

func BirdModule(birds map[string]*Bird) Node {
	return &MyNode{
		OnChild: func(r ChildRequest) (Node, error) {
			switch r.Meta.GetIdent() {
			case "bird":
				return BirdList(birds), nil
			}
			return nil, nil
		},
	}
}

func BirdList(birds map[string]*Bird) Node {
	index := NewIndex(birds)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &MyNode{
		OnNext: func(r ListRequest) (Node, []*Value, error) {
			var b *Bird
			key := r.Key
			if r.New {
				b = &Bird{Name: key[0].Str}
				birds[b.Name] = b
			} else if key != nil {
				if r.Delete {
					delete(birds, key[0].Str)
				} else {
					b = birds[key[0].Str]
				}
			} else {
				if v := index.NextKey(r.Row); v != NO_VALUE {
					name := v.String()
					if b = birds[name]; b != nil {
						key = SetValues(r.Meta.KeyMeta(), name)
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

func BirdNode(b *Bird) Node {
	return ReflectNode(b)
}
