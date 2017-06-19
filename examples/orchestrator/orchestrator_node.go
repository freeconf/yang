package orchestrator

import "github.com/c2stack/c2g/node"

func Node(o *Orchestrator) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "app":
				return appListNode(o), nil
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
				if err := o.Factory.NewApp(a); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
