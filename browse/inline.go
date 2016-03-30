package browse

import (
	"fmt"
	"github.com/blitter/blit"
	"github.com/blitter/meta"
	"github.com/blitter/meta/yang"
	"github.com/blitter/node"
)

// All docs have form
// {
//   meta:{}
//   data:{}
//}
//
// we force this structure so nodes in "data" would never collide with "meta"
//
type Inline struct {
	Module     *meta.Module
	bodyMeta   meta.MetaList
	DataMeta   meta.MetaList
	Url        string
	InsideList bool
	Key        []string
}

func NewInline() *Inline {
	self := &Inline{}
	// because we modifiy the meta, we need to load a fresh one each time
	self.Module = FreshYang("inline")
	self.bodyMeta = meta.FindByIdent2(self.Module, "data").(meta.MetaList)
	return self
}

func (self *Inline) Save(c *node.Context, m meta.MetaList, onWrite node.Node) node.Node {
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = node.DecoupledMetaCopy(m)

	self.bodyMeta.Clear()
	self.bodyMeta.AddMeta(self.DataMeta)
	return self.node(onWrite, nil)
}

func (self *Inline) SaveSelection(c *node.Context, sel *node.Selection, n node.Node) node.Node {
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = node.DecoupledMetaCopy(sel.Meta().(meta.MetaList))
	self.InsideList = sel.InsideList()
	if len(sel.Key()) > 0 {
		self.Key = node.SerializeKey(sel.Key())
	}

	self.bodyMeta.Clear()
	self.bodyMeta.AddMeta(self.DataMeta)
	return self.node(n, nil)
}

// Useful with pipe node as well if you don't have an output node but are looking for
// and output node
func (self *Inline) Load(c *node.Context, input node.Node, output node.Node, waitForSchemaLoad chan error) error {
	// probably the coolest single line i've ever written, this loads the node and the meta
	// then redirects the node output to the given node w/o ever keeping a copy.  This
	// inspired the name bitblit which then lead to datablit!

	err := c.Select(self.Module, self.node(output, waitForSchemaLoad)).UpsertFrom(input).LastErr

	// in unusual case there is no meta, unblock anything waiting on meta to load
	if err == nil && waitForSchemaLoad != nil && self.DataMeta == nil {
		waitForSchemaLoad <- blit.NewErr("No meta found")
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
		Node: node.MarshalContainer(self),
		OnChoose: func(p node.Node, sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			switch choice.GetIdent() {
			case "handle":
				if len(self.Url) > 0 {
					return choice.GetCase("remote"), nil
				} else if meta.IsList(self.DataMeta) {
					if self.InsideList {
						return choice.GetCase("inline-list-item"), nil
					}
					return choice.GetCase("inline-list"), nil
				}
				return choice.GetCase("inline-container"), nil
			}
			return p.Choose(sel, choice)
		},
		OnSelect: func(p node.Node, r node.ContainerRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "container", "list", "list-item":
				if r.New {
					if r.Meta.GetIdent() == "container" {
						self.DataMeta = &meta.Container{Ident: "data"}
					} else {
						self.DataMeta = &meta.List{Ident: "data"}
						if r.Meta.GetIdent() == "list-item" {
							self.InsideList = true
						}
					}
					self.bodyMeta.Clear()
					self.bodyMeta.AddMeta(self.DataMeta)
				}
				if self.DataMeta != nil {
					return node.SchemaData{Resolve: false}.MetaList(self.DataMeta), nil
				}
				return nil, nil
			}
			return p.Select(r)
		},
		OnWrite: func(p node.Node, r node.FieldRequest, v *node.Value) error {
			switch r.Meta.GetIdent() {
			case "url":
				if err := DownloadMeta(v.Str, self.DataMeta); err != nil {
					return err
				}
				//node.RenameMeta(self.DataMeta, "data")
				self.Url = v.Str
				return nil
			}
			return p.Write(r, v)
		},
	}
}

func (self *Inline) node(nodeNode node.Node, waitForSchemaLoad chan error) node.Node {
	n := &node.MyNode{}
	li := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "meta":
			if waitForSchemaLoad != nil {
				r.Selection.OnChild(node.LEAVE_EDIT, r.Meta, func() error {
					waitForSchemaLoad <- nil
					return nil
				})
			}
			return self.selectSchema(), nil
		case "data":
			if self.InsideList {
				return li, nil
			}
			return nodeNode, nil
		}
		return nil, nil
	}
	li.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if r.Meta.GetIdent() == self.DataMeta.GetIdent() {
			return li, nil
		}
		return nil, nil
	}
	li.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		key := r.Key
		if r.First {
			if len(r.Meta.Key) > 0 {
				if key == nil {
					var keyErr error
					key, keyErr = node.CoerseKeys(r.Meta, self.Key)
					if keyErr != nil {
						return nil, nil, keyErr
					}
				} else {
					self.Key = node.SerializeKey(key)
				}
			}
			return nodeNode, key, nil
		}
		return nil, nil, nil
	}
	return n
}

func init() {
	yang.InternalYang()["inline"] = `
module inline {
	namespace "";
	prefix "";
	import yanglib;
	revision 0;

	container meta {
		choice handle {
			case inline-container {
				container container {
					leaf ident {
						type string;
					}
    					uses containers-lists-leafs-uses-choice;
				}
			}
			case inline-list {
				container list {
					leaf ident {
						type string;
					}
    					uses containers-lists-leafs-uses-choice;
				}
			}
			case inline-list-item {
				container list-item {
					leaf ident {
						type string;
					}
    					uses containers-lists-leafs-uses-choice;
				}
				leaf-list key {
					type string;
				}
			}
			case remote {
				leaf url {
					type string;
				}
			}
		}
	}
	container data {
		/* empty placeholder */
	}
}
`
}
