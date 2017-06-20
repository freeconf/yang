package app

import "github.com/c2stack/c2g/node"
import "github.com/c2stack/c2g/meta"

func Node(o *Orchestrator, ypath meta.StreamSource) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "app":
				return appListNode(o), nil
			case "builder":
				if r.New || o.Builder != nil {
					return builderNode(o, ypath), nil
				}
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "count":
				hnd.Val = &node.Value{Int: len(o.Apps)}
			}
			return nil
		},
	}
}

func builderNode(o *Orchestrator, ypath meta.StreamSource) node.Node {
	return &node.MyNode{
		OnChoose: func(s node.Selection, c *meta.Choice) (*meta.ChoiceCase, error) {
			switch o.Builder.(type) {
			case *InMemory:
				return c.GetCase("inMemory"), nil
			case *Exec:
				return c.GetCase("exec"), nil
			}
			return nil, nil
		},
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "inMemory":
				if r.New {
					o.Builder = NewInMemory(ypath)
				}
				if o.Builder != nil {
					return &node.MyNode{}, nil
				}
			case "exec":
				if r.New {
					o.Builder = &Exec{}
				}
				if o.Builder != nil {
					return &node.MyNode{}, nil
				}
			}
			return nil, nil
		},
	}
}

func appListNode(o *Orchestrator) node.Node {
	index := node.NewIndex(o.Apps)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			var a *App
			key := r.Key
			if r.New {
				a = &App{Id: key[0].Str}
				o.Apps[a.Id] = a
			} else if key != nil {
				a = o.Apps[key[0].Str]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					id := v.String()
					if a = o.Apps[id]; a != nil {
						key = node.SetValues(r.Meta.KeyMeta(), id)
					}
				}
			}
			if a != nil {
				return appNode(o, a), key, nil
			}
			return nil, nil, nil
		},
	}
}

func appNode(o *Orchestrator, a *App) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(a),
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			if r.New {
				if err := o.Builder.NewApp(a); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
