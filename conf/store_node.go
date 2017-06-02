package conf

import "github.com/c2stack/c2g/node"
import "github.com/c2stack/c2g/c2"
import "github.com/c2stack/c2g/meta"

type StoreNode struct {
}

func (self StoreNode) Node(a node.Node, b node.Node) node.Node {
	var edit node.Node
	return &node.MyNode{
		OnPeek: func(s node.Selection, consumer interface{}) interface{} {
			if b != nil {
				if o := b.Peek(s, consumer); o != nil {
					return o
				}
			}
			if a != nil {
				return a.Peek(s, consumer)
			}
			return nil
		},
		OnChoose: func(sel node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			if choice.GetParent().(meta.HasDetails).Details().Config(sel.Path) {
				if a != nil {
					return a.Choose(sel, choice)
				}
			} else {
				if b != nil {
					return b.Choose(sel, choice)
				}
			}
			return nil, nil
		},
		OnBeginEdit: func(r node.NodeRequest) error {
			if r.EditRoot {
				var err error
				if edit, err = self.createEdit(r, a, b); err != nil {
					return err
				}
			}
			return nil
		},
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			if b != nil {
				return b.Action(r)
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			if b != nil {
				return b.Notify(r)
			}
			return nil, nil
		},
		OnEndEdit: func(r node.NodeRequest) error {
			if edit == nil {
				panic("Begin edit never called")
			}
			edit.EndEdit(r)
			if r.EditRoot {
				err := r.Selection.Split(edit).InsertInto(b).LastErr
				if err != nil {
					return err
				}
				if err = r.Selection.Split(edit).InsertInto(a).LastErr; err != nil {
					// MESSY STATE : Need to support undo on node "b"
					c2.Err.Printf("Device is configured but store could not save.  Device needs rebooting")
					return err
				}
			}
			return nil
		},
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			if edit != nil {
				return edit.Child(r)
			}
			var err error
			var aChild, bChild node.Node
			if a != nil {
				if aChild, err = a.Child(r); err != nil {
					return nil, err
				}
			}
			if b != nil {
				if bChild, err = b.Child(r); err != nil {
					return nil, err
				}
			}
			if a != nil || b != nil {
				return self.Node(aChild, bChild), nil
			}
			return nil, nil
		},
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			if edit != nil {
				return edit.Next(r)
			}
			var err error
			var aChild, bChild node.Node
			var aKey, bKey []*node.Value
			if a != nil {
				if aChild, aKey, err = a.Next(r); err != nil {
					return nil, aKey, err
				}
			}
			if b != nil {
				if bChild, bKey, err = b.Next(r); err != nil {
					return nil, bKey, err
				}
			}
			if a != nil || b != nil {
				key := aKey
				if key == nil {
					key = bKey
				}
				return self.Node(aChild, bChild), key, nil
			}
			return nil, nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			if edit != nil {
				return edit.Field(r, hnd)
			}
			if r.Meta.(meta.HasDetails).Details().Config(r.Path) {
				if a != nil {
					return a.Field(r, hnd)
				}
			} else {
				if b != nil {
					return b.Field(r, hnd)
				}
			}

			return nil
		},
	}
}

func (self StoreNode) createEdit(r node.NodeRequest, a node.Node, b node.Node) (node.Node, error) {
	params := "depth=1&content=config&with-defaults=trim"
	data := make(map[string]interface{})
	edit := node.MapNode(data)

	// shallow copy in existing config so inserts know what already exists
	err := r.Selection.Split(a).Constrain(params).InsertInto(edit).LastErr
	if err != nil {
		return nil, err
	}
	return edit, nil
}
