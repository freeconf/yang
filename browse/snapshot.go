package browse

import (
	"io"

	"github.com/dhubler/c2g/c2"
	"github.com/dhubler/c2g/meta"
	"github.com/dhubler/c2g/meta/yang"
	"github.com/dhubler/c2g/node"
)

// Takes a given selection anywhere in a given meta set and stores it into a given node
// that can then be restored back into the original selection at any time. The meta first
// has to be decoupled from it's ancestors so restoration process does not need access
// or original schema.  This will not work on recursive metasets.
type SelectionSnapshot struct {
	YangPath meta.StreamSource
	DataMeta *meta.Container
	Resolver MetaResolver
}

func RestoreSelection(yangPath meta.StreamSource, n node.Node, resolver MetaResolver) (*node.Selection, error) {
	snap := &SelectionSnapshot{
		YangPath: yangPath,
		Resolver: resolver,
	}
	if snap.Resolver == nil {
		snap.Resolver = DownloadMeta
	}
	return snap.Restore(n)
}

func (self *SelectionSnapshot) Restore(n node.Node) (*node.Selection, error) {
	m := yang.RequireModule(self.YangPath, "snapshot")
	self.DataMeta = meta.FindByIdent2(m, "data").(*meta.Container)
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	onSchemaLoad := make(chan error)
	go func() {
		s := node.NewBrowser2(m, n).Root().Selector()
		err := s.UpsertInto(self.node(push, onSchemaLoad)).LastErr
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
	return node.NewBrowser2(self.DataMeta, pull).Root(), nil
}

func (self *SelectionSnapshot) selectImports() node.Node {
	n := &node.MyNode{}
	n.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		if r.New {
			return n, nil, nil
		}
		return nil, nil, nil
	}
	n.OnField = func(r node.FieldRequest, hnd *node.ValueHandle) error {
		switch r.Meta.GetIdent() {
		case "container":
			c := &meta.Container{Ident: "unknown"}
			if err := self.Resolver(self.YangPath, hnd.Val.Str, c); err != nil {
				return err
			}
			self.DataMeta.AddMeta(c)
		case "list":
			l := &meta.List{Ident: "unknown"}
			if err := self.Resolver(self.YangPath, hnd.Val.Str, l); err != nil {
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

func SaveSelection(yangPath meta.StreamSource, to *node.Selection) *node.Selection {
	snap := &SelectionSnapshot{
		YangPath: yangPath,
	}
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
	m := yang.RequireModule(self.YangPath, "snapshot")
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = meta.FindByIdent2(m, "data").(*meta.Container)
	copy := node.DecoupledMetaCopy(self.YangPath, to.Meta().(meta.MetaList))
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

	return node.NewBrowser2(m, self.node(toNode, nil)).Root()
}
