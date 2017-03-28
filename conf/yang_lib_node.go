package conf

import "github.com/c2stack/c2g/node"

// Implementation of RFC7895

func LocalDeviceYangLibNode(ld *LocalDevice) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "modules-state":
				return localYangLibModuleState(ld), nil
			}
			return nil, nil
		},
	}
}

func localYangLibModuleState(ld *LocalDevice) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "module":
				mods, err := ld.ModuleHandles()
				if err != nil {
					return nil, err
				} else if len(mods) > 0 {
					return YangLibModuleList(mods), nil
				}
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			return nil
		},
	}
}

func YangLibModuleList(mods map[string]*ModuleHandle) node.Node {
	index := node.NewIndex(mods)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			var e *ModuleHandle
			if r.New {
				e = &ModuleHandle{
					Name: r.Key[0].Str,
				}
				mods[e.Name] = e
			} else if r.Key != nil {
				e = mods[r.Key[0].Str]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					module := v.String()
					if e = mods[module]; e != nil {
						key = node.SetValues(r.Meta.KeyMeta(), e.Name)
					}
				}
			}
			if e != nil {
				return yangLibModuleHandleNode(e), key, nil
			}
			return nil, nil, nil
		},
	}
}

func yangLibModuleHandleNode(e *ModuleHandle) node.Node {
	return &node.Extend{
		Node: node.ReflectNode(e),
		OnChild: func(p node.Node, r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "submodule":
				if r.New {
					e.Submodule = make(map[string]*ModuleHandle)
				}
				if e.Submodule != nil {
					return YangLibModuleList(e.Submodule), nil
				}
			default:
				return p.Child(r)
			}
			return nil, nil
		},
	}
}
