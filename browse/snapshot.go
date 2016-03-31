package browse

import (
	"github.com/blitter/blit"
	"github.com/blitter/meta"
	"github.com/blitter/meta/yang"
	"github.com/blitter/node"
)

// Takes a given selection anywhere in a given meta set and stores it into a given node
// that can then be restored back into the original selection at any time. The meta first
// has to be decoupled from it's ancestors so restoration process does not need access
// or original schema.  This will not work on recursive metasets.
type SelectionSnapshot struct {
	bodyMeta    meta.MetaList
	DataMeta    meta.MetaList
	InsideList  bool
	Key         []string
	restoreMode bool
}

func RestoreSelection(n node.Node) (*node.Selection, error) {
	snap := &SelectionSnapshot{}
	return snap.Restore(n)
}

func (self *SelectionSnapshot) Restore(n node.Node) (*node.Selection, error) {
	m := yang.InternalModule("snapshot")
	self.restoreMode = true
	self.bodyMeta = meta.FindByIdent2(m, "data").(meta.MetaList)
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	c := node.NewContext()
	onSchemaLoad := make(chan error)
	defer func() {
		close(onSchemaLoad)
		onSchemaLoad = nil
	}()
	go func() {
		err := c.Select(m, n).UpsertInto(self.node(push, onSchemaLoad)).LastErr
		// errors can come in meta region or data region. If they come in the meta region
		// then onSchemaLoad is still valid and we send error there to make error handling
		// synchronous.   otherwise we're into the data section and error will be asynchronous
		// to this function but synchronous to the caller of returned selection
		if onSchemaLoad != nil && err != nil {
blit.Debug.Printf("error %s", err.Error())
			//onSchemaLoad <- err
		}
		pipe.Close(err)
	}()
	// wait until self.DataMeta is valid...
	err := <-onSchemaLoad
	if err != nil {
		return nil, err
	}
	if self.DataMeta == nil {
		return nil, blit.NewErr("No meta found in restore data")
	}
	return node.Select(self.DataMeta, pull), nil
}

func (self *SelectionSnapshot) selectMetaDefinition() node.Node {
	return &node.Extend{
		Node: node.SchemaData{Resolve: false}.MetaList(self.DataMeta),
		OnWrite: func(p node.Node, r node.FieldRequest, v *node.Value) error {
			switch r.Meta.GetIdent() {
			case "url":
				if err := DownloadMeta(v.Str, self.DataMeta); err != nil {
					return err
				}
				return nil
			}
			return p.Write(r, v)
		},
		OnRead: func(p node.Node, r node.FieldRequest) (*node.Value, error) {
			switch r.Meta.GetIdent() {
			case "url":
				return nil, nil
			}
			return p.Read(r)
		},
	}
}

func (self *SelectionSnapshot) selectMeta() node.Node {
	return &node.Extend{
		Node: node.MarshalContainer(self),
		OnChoose: func(p node.Node, sel *node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			switch choice.GetIdent() {
			case "handle":
				if meta.IsList(self.DataMeta) {
					if self.InsideList {
						return choice.GetCase("inline-list-item"), nil
					}
					return choice.GetCase("inline-list"), nil
				}
				return choice.GetCase("inline-container"), nil
			}
			return nil, nil
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
					return self.selectMetaDefinition(), nil
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

func (self *SelectionSnapshot) node(to node.Node, onSchemaLoad chan error) node.Node {
	n := &node.MyNode{}
	data := &node.MyNode{}
	n.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "meta":
			return self.selectMeta(), nil
		case "data":
			if onSchemaLoad != nil {
				onSchemaLoad <- nil
			}
			return data, nil
		}
		return nil, nil
	}
	data.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		if r.Meta.GetIdent() == self.DataMeta.GetIdent() {
			if self.InsideList {
				return data, nil
			} else {
				return to, nil
			}
		}
		return nil, nil
	}
	data.OnNext = func(r node.ListRequest) (node.Node, []*node.Value, error) {
		key := r.Key
		if r.First {
			if len(r.Meta.Key) > 0 {
				if key == nil {
					var keyErr error
					key, keyErr = node.CoerseKeys(r.Meta, self.Key)
					if keyErr != nil {
						return nil, nil, keyErr
					}
				} else if self.Key == nil {
					self.Key = node.SerializeKey(key)
				}
			}
			// tricky - if restoring, source node is not positions to read list item yet
			// but when saving, it has already called n.Next() and selection has that returned
			// node and that was what was expected/given to Save() so we do something a little
			// different here on save v.s. restore
			if self.restoreMode {
				toNext, _, toNextErr := to.Next(r)
				return toNext, key, toNextErr
			}
			return to, key, nil
		}
		return nil, nil, nil
	}
	return n
}

func (self *SelectionSnapshot) Save(to *node.Selection) *node.Selection {
	m := yang.InternalModule("snapshot")
	self.bodyMeta = meta.FindByIdent2(m, "data").(meta.MetaList)
	// we resolve meta because consumer will need all meta self-contained
	// to validate and/or persist w/o original meta, parent heirarchy
	self.DataMeta = node.DecoupledMetaCopy(to.Meta().(meta.MetaList))
	self.InsideList = to.InsideList()
	if len(to.Key()) > 0 {
		self.Key = node.SerializeKey(to.Key())
	}
	self.bodyMeta.Clear()
	self.bodyMeta.AddMeta(self.DataMeta)
	return node.Select(m, self.node(to.Node(), nil))
}

func init() {
	yang.InternalYang()["snapshot"] = `
module snapshot {
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
					leaf url {
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
					leaf-list key {
					    type string;
					}
					leaf url {
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
					leaf-list key {
					    type string;
					}
					leaf url {
						type string;
					}
    					uses containers-lists-leafs-uses-choice;
				}
				leaf-list key {
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
