package browse

import (
	"github.com/c2g/c2"
	"github.com/c2g/meta"
	"github.com/c2g/meta/yang"
	"github.com/c2g/node"
	"io"
)

// Takes a given selection anywhere in a given meta set and stores it into a given node
// that can then be restored back into the original selection at any time. The meta first
// has to be decoupled from it's ancestors so restoration process does not need access
// or original schema.  This will not work on recursive metasets.
type SelectionSnapshot struct {
	DataMeta *meta.Container
	Resolver MetaResolver
}

func RestoreSelection(c node.Context, n node.Node, resolver MetaResolver) (*node.Selection, error) {
	snap := &SelectionSnapshot{
		Resolver: resolver,
	}
	if snap.Resolver == nil {
		snap.Resolver = DownloadMeta
	}
	return snap.Restore(c, n)
}

func (self *SelectionSnapshot) Restore(c node.Context, n node.Node) (*node.Selection, error) {
	m := yang.InternalModule("snapshot")
	self.DataMeta = meta.FindByIdent2(m, "data").(*meta.Container)
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	onSchemaLoad := make(chan error)
	go func() {
		err := c.Select(m, n).UpsertInto(self.node(push, onSchemaLoad)).LastErr
		// errors can come in meta region or data region. If they come in the meta region
		// then onSchemaLoad is still valid and we send error there to make error handling
		// synchronous.   otherwise we're into the data section and error will be asynchronous
		// to this function but synchronous to the caller of returned selection
		if onSchemaLoad != nil && err != nil {
			onSchemaLoad <- err
		}
		pipe.Close(err)
	}()
	// wait until self.DataMeta is valid...
	err, valid := <-onSchemaLoad
	if self.DataMeta == nil && err == io.EOF {
		return nil, c2.NewErr("No meta found in restore data")
	}
	if valid && err != nil {
		return nil, err
	}
	return node.Select(self.DataMeta, pull), nil
}

func (self *SelectionSnapshot) selectImports() node.Node {
	n := &node.MyNode{}
	n.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if r.New {
			return n, nil, nil
		}
		return nil, nil, nil
	}
	n.OnWrite = func(r node.FieldRequest, v *node.Value) error {
		switch r.Meta.GetIdent() {
		case "container":
			c := &meta.Container{Ident: "unknown"}
			if err := self.Resolver(v.Str, c); err != nil {
				return err
			}
			self.DataMeta.AddMeta(c)
		case "list":
			l := &meta.List{Ident: "unknown"}
			if err := self.Resolver(v.Str, l); err != nil {
				return err
			}
			self.DataMeta.AddMeta(l)
		}
		return nil
	}
	return n
}

func (self *SelectionSnapshot) selectMetaDefinition() node.Node {
	return &node.Extend{
		Node: node.SchemaData{Resolve: false}.MetaList(self.DataMeta),
		OnSelect: func(p node.Node, r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "import":
				if r.New {
					return self.selectImports(), nil
				}
				return nil, nil
			}
			return p.Select(r)
		},
	}
}

func SaveSelection(to *node.Selection) *node.Selection {
	snap := &SelectionSnapshot{}
	return snap.Save(to)
}

func (self *SelectionSnapshot) node(dataNode node.Node, onSchemaLoad chan error) node.Node {
	n := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "meta":
			return self.selectMetaDefinition(), nil
		case "data":
			if onSchemaLoad != nil {
				onSchemaLoad <- nil
			}
			return dataNode, nil
		}
		return nil, nil
	}
	return n
}

func (self *SelectionSnapshot) Save(to *node.Selection) *node.Selection {
	m := yang.InternalModule("snapshot")
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = meta.FindByIdent2(m, "data").(*meta.Container)
	copy := node.DecoupledMetaCopy(to.Meta().(meta.MetaList))
	isList := meta.IsList(to.Meta())
	var toNode node.Node
	if isList && !to.InsideList() {
		self.DataMeta.AddMeta(copy)
		toNode = &node.MyNode{
			OnSelect: func(r node.ContainerRequest) (node.Node, error) {
				return to.Node(), nil
			},
		}
	} else {
		i := meta.NewMetaListIterator(copy, false)
		for i.HasNextMeta() {
			self.DataMeta.AddMeta(i.NextMeta())
		}
		toNode = to.Node()
	}

	return node.Select(m, self.node(toNode, nil))
}

func init() {
	yang.InternalYang()["snapshot"] = `
module snapshot {
	namespace "";
	prefix "";
	import yanglib;
	revision 0;

	container meta {
	        list import {
	           choice importer {
	             case import-container {
			   leaf container {
			      type string;
			   }
	             }
	             case import-list {
			   leaf list {
			      type string;
			   }
	             }
	           }
	        }
		uses containers-lists-leafs-uses-choice;
	}

	container data {
		/* empty placeholder */
	}
}
`
}
