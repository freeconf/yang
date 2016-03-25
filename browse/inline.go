package browse

import (
	"node"
	"fmt"
	"meta"
	"meta/yang"
)

// All docs have form
// {
//   meta:{}
//   node:{}
//}
//
// we force custom node to be in fixed container called "node" so we could never collide with other container
// called "meta".  Otherwise, it would be rather useful conceptually to have docs like
// {
//   meta : { },
//   foo: {}
//}
// however, what if someone's custom node was actually called "meta", then this wouldn't work.
//
type Inline struct {
	Module   *meta.Module
	DataMeta meta.MetaList
	Url string
}

func NewInline() *Inline {
	self := &Inline{}
	// because we modifiy the goober, we need to load a fresh one each time
	self.Module = FreshYang("inline")
	return self
}

func (self *Inline) Save(c *node.Context, goober meta.MetaList, onWrite node.Node) node.Node {
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = node.DecoupledMetaCopy(goober)
	if self.DataMeta.GetIdent() != "node" {
		node.RenameMeta(self.DataMeta, "node")
	}
	self.Module.AddMeta(self.DataMeta)
	return self.node(onWrite, nil)
}

// Useful with pipe node as well if you don't have an output node but are looking for
// and output node
func (self *Inline) Load(c *node.Context, input node.Node, output node.Node, waitForSchemaLoad chan error) (error) {
	// probably the coolest single line i've ever written, this loads the node and the meta
	// then redirects the node output to the given node w/o ever keeping a copy.  This
	// inspired the name bitblit which then lead to nodeblit!
	self.DataMeta = &meta.Container{Ident:"node"}
	self.Module.AddMeta(self.DataMeta)

	err := c.Select(self.Module, self.node(output, waitForSchemaLoad)).UpsertFrom(input).LastErr

	// in unusual case there is no meta, unblock anything waiting on meta to load
	if err == nil && waitForSchemaLoad != nil && meta.ListLen(self.DataMeta) == 0 {
		waitForSchemaLoad <- nil
	}

	return err
}

func FreshYang(name string) *meta.Module {
	// TODO: performance - return deep copy of cached copy
	inlineYang, err := yang.LoadModule(yang.InternalYang(), name)
	if err != nil {
		msg := fmt.Sprintf("Error parsing %s yang, %s", name, err.Error())
		panic(msg)
	}
	return inlineYang
}

func (self *Inline) selectSchema() node.Node {
	return &node.Extend{
		Node: node.SchemaData{Resolve: false}.MetaList(self.DataMeta),
		OnChoose: func(p node.Node, sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			switch choice.GetIdent() {
			case "handle":
				if len(self.Url) > 0 {
					return choice.GetCase("remote"), nil
				}
				return choice.GetCase("inline"), nil
			}
			return p.Choose(sel, choice)
		},
		OnRead: func(p node.Node, r node.FieldRequest) (*node.Value, error) {
			switch r.Meta.GetIdent() {
			case "url":
				return &node.Value{Str:self.Url}, nil
			}
			return p.Read(r)
		},
		OnWrite: func(p node.Node, r node.FieldRequest, v *node.Value) error {
			switch r.Meta.GetIdent() {
			case "url":
				if err := DownloadMeta(v.Str, self.DataMeta); err != nil {
					return err
				}
				node.RenameMeta(self.DataMeta, "node")
				self.Url = v.Str
				return nil
			}
			return p.Write(r, v)
		},
	}
}


func (self *Inline) node(nodeNode node.Node, waitForSchemaLoad chan error) node.Node {
	n := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "node":
			return nodeNode, nil
		case "meta":
			if waitForSchemaLoad != nil {
				r.Selection.OnChild(node.LEAVE_EDIT, r.Meta, func () error {
					waitForSchemaLoad <- nil
					return nil
				})
			}
			return self.selectSchema(), nil
		}
		return nil, nil
	}
	return n
}


func init() {
	yang.InternalYang()["inline"] = `
module inline {
	namespace "";
	prefix "";
	import yang;
	revision 0;

	container schema {
		choice handle {
			case inline {
				uses containers-lists-leafs-uses-choice;
			}
			case remote {
				leaf url {
					type string;
				}
			}
		}
	}
}
`
}