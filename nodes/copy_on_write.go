package nodes

import (
	"github.com/freeconf/gconf/meta"
	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/val"
)

type CopyOnWrite struct {
}

func (self CopyOnWrite) Node(s node.Selection, read node.Node, write node.Node) node.Node {
	if meta.IsList(s.Meta()) && !s.InsideList {
		return self.list(read, write)
	}
	return self.container(read, write)
}

func (self CopyOnWrite) list(read node.Node, write node.Node) node.Node {
	if read == nil {
		panic("nil read")
	}
	return &Extend{
		Base: read,
		OnNext: func(p node.Node, r node.ListRequest) (node.Node, []val.Value, error) {
			if r.New {
				return write.Next(r)
			}
			rChild, key, err := read.Next(r)
			if err != nil || rChild == nil {
				return nil, key, err
			}
			rNew := r
			rNew.New = true
			rNew.Key = key
			wChild, _, err := write.Next(rNew)
			if err != nil {
				return nil, key, err
			}
			return self.container(rChild, wChild), key, err
		},
	}
}

func (self CopyOnWrite) container(read node.Node, write node.Node) node.Node {
	return &Extend{
		Base: read,
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			if r.New {
				return write.Child(r)
			}
			rChild, err := read.Child(r)
			if err != nil || rChild == nil {
				return nil, err
			}
			rNew := r
			rNew.New = true
			wChild, err := write.Child(rNew)
			if err != nil {
				return nil, err
			}
			if meta.IsList(r.Meta) {
				return self.list(rChild, wChild), nil
			}
			return self.container(rChild, wChild), nil
		},
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			if r.Write {
				return write.Field(r, hnd)
			}
			return p.Field(r, hnd)
		},
	}
}
