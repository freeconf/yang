package device

import "github.com/c2stack/c2g/node"
import "github.com/c2stack/c2g/meta"
import "github.com/c2stack/c2g/meta/yang"

// Implementation of RFC7895

func LocalDeviceYangLibNode(ld *Local) node.Node {
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

func localYangLibModuleState(ld *Local) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "module":
				mods := ld.Modules()
				if len(mods) > 0 {
					return YangLibModuleList(mods, ld.SchemaSource()), nil
				}
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			return nil
		},
	}
}

func YangLibModuleList(mods map[string]*meta.Module, yangPath meta.StreamSource) node.Node {
	index := node.NewIndex(mods)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			var m *meta.Module
			if r.New {
				m, err := yang.LoadModule(yangPath, key[0].Str)
				if err != nil {
					return nil, nil, err
				}
				mods[m.GetIdent()] = m
			} else if r.Key != nil {
				m = mods[r.Key[0].Str]
			} else {
				if v := index.NextKey(r.Row); v != node.NO_VALUE {
					module := v.String()
					if m = mods[module]; m != nil {
						key = node.SetValues(r.Meta.KeyMeta(), m.GetIdent())
					}
				}
			}
			if m != nil {
				return yangLibModuleHandleNode(m), key, nil
			}
			return nil, nil, nil
		},
	}
}

func yangLibModuleHandleNode(m *meta.Module) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			// deviation
			// submodule
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "name":
				hnd.Val = &node.Value{Str: m.GetIdent()}
			case "revision":
				hnd.Val = &node.Value{Str: m.Revision.GetIdent()}
			case "schema":
			case "namespace":
			case "feature":
			case "conformance-type":
			}
			return nil
		},
	}
}
