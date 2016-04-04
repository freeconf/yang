package node

import "github.com/c2g/meta"

type AnyData interface {

	Writer() (Node)

	Reader() (Node)

	// Can be nil
	Meta() meta.MetaList
}

type AnyNode struct {
	Read   Node
	Write  Node
	Schema meta.MetaList
}

func (self AnyNode) Meta() (meta.MetaList) {
	if self.Schema == nil {
		panic("Format doesn't have schema")
	}
	return self.Schema
}

func (self AnyNode) Reader() (Node) {
	if self.Read == nil {
		panic("Not a reader")
	}
	return self.Read
}

func (self AnyNode) Writer() Node {
	if self.Write == nil {
		panic("Not a writer")
	}
	return self.Write
}