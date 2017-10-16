package nodes

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/val"
)

/*
   Proxy all but config prremoteties to a delegate node.  For the config read prremoteties
   simply return local copy, for config writes send a copy to far end and if returns ok
   then trigger storage to save.
*/
type ConfigProxy struct {
}

func (self ConfigProxy) Node(local node.Node, remote node.Node) node.Node {
	if local == nil {
		panic("nil local")
	}
	var edit, changes node.Node
	return &Basic{
		OnPeek: func(s node.Selection, consumer interface{}) interface{} {
			if remote != nil {
				if o := remote.Peek(s, consumer); o != nil {
					return o
				}
			}
			if local != nil {
				return local.Peek(s, consumer)
			}
			return nil
		},
		OnChoose: func(sel node.Selection, choice *meta.Choice) (*meta.ChoiceCase, error) {
			if choice.GetParent().(meta.HasDetails).Details().Config(sel.Path) {
				if local != nil {
					return local.Choose(sel, choice)
				}
			} else {
				if remote != nil {
					return remote.Choose(sel, choice)
				}
			}
			return nil, nil
		},
		OnBeginEdit: func(r node.NodeRequest) error {
			if r.EditRoot {
				var err error
				if edit, changes, err = self.createEdit(r.Selection, local); err != nil {
					return err
				}
			}
			return nil
		},
		OnAction: func(r node.ActionRequest) (node.Node, error) {
			if remote != nil {
				return remote.Action(r)
			}
			return nil, nil
		},
		OnNotify: func(r node.NotifyRequest) (node.NotifyCloser, error) {
			if remote != nil {
				return remote.Notify(r)
			}
			return nil, nil
		},
		OnEndEdit: func(r node.NodeRequest) error {
			if edit != nil {
				if err := edit.EndEdit(r); err != nil {
					return err
				}
			}
			if r.EditRoot {
				// First we see if remote side likes the edit

				err := r.Selection.Split(changes).UpsertInto(remote).LastErr
				if err != nil {
					return err
				}

				// Second we save edit to local node
				if err = r.Selection.Split(changes).UpsertInto(local).LastErr; err != nil {
					// MESSY STATE : Need to support undo on remote node? However local
					// node should rarely fail on a syntactically correct edit.
					c2.Err.Printf("Device is configured but store could not save.  Device might need rebooting")
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
			var localChild, remoteChild node.Node
			if local != nil {
				if localChild, err = local.Child(r); err != nil {
					return nil, err
				}
			}
			if remote != nil {
				if remoteChild, err = remote.Child(r); err != nil {
					return nil, err
				}
			}
			if localChild != nil || remoteChild != nil {
				return self.Node(localChild, remoteChild), nil
			}
			return nil, nil
		},
		OnNext: func(r node.ListRequest) (node.Node, []val.Value, error) {
			if edit != nil {
				return edit.Next(r)
			}
			var err error
			var localChild, remoteChild node.Node
			var localKey, remoteKey []val.Value
			if local != nil {
				if localChild, localKey, err = local.Next(r); err != nil {
					return nil, localKey, err
				}
			}
			if remote != nil {
				if remoteChild, remoteKey, err = remote.Next(r); err != nil {
					return nil, remoteKey, err
				}
			}
			if localChild != nil || remoteChild != nil {
				key := localKey
				if key == nil {
					key = remoteKey
				}
				return self.Node(localChild, remoteChild), key, nil
			}
			return nil, nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			if edit != nil {
				return edit.Field(r, hnd)
			}
			if r.Meta.(meta.HasDetails).Details().Config(r.Path) {
				if local != nil {
					return local.Field(r, hnd)
				}
			} else {
				if remote != nil {
					return remote.Field(r, hnd)
				}
			}

			return nil
		},
	}
}

func (self ConfigProxy) createEdit(s node.Selection, local node.Node) (node.Node, node.Node, error) {
	var changes node.Node
	if meta.IsList(s.Meta()) && !s.InsideList {
		// by making capacity = len + 1, we know appending a single item
		// will not change data reference.  Ideally, reflect API should
		// allow me to do this more cleanly
		data := make([]map[string]interface{}, 0, 1)
		changes = ReflectList(data)
	} else {
		data := make(map[string]interface{})
		changes = ReflectChild(data)
	}
	return CopyOnWrite{}.Node(s, local, changes), changes, nil
}
