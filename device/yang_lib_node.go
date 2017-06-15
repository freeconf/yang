package device

import (
	"reflect"
	"strings"

	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

// Implementation of RFC7895

// Export device by it's address so protocol server can serve a device
// often referred to northbound
type ModuleAddresser func(m *meta.Module) string

func LocalDeviceYangLibNode(addresser ModuleAddresser, d Device) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "modules-state":
				return localYangLibModuleState(addresser, d), nil
			}
			return nil, nil
		},
	}
}

func localYangLibModuleState(addresser ModuleAddresser, d Device) node.Node {
	return &node.MyNode{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.GetIdent() {
			case "module":
				mods := d.Modules()
				if len(mods) > 0 {
					return YangLibModuleList(addresser, mods), nil
				}
			}
			return nil, nil
		},
		OnField: func(r node.FieldRequest, hnd *node.ValueHandle) error {
			return nil
		},
	}
}

func YangLibModuleList(addresser ModuleAddresser, mods map[string]*meta.Module) node.Node {
	index := node.NewIndex(mods)
	index.Sort(func(a, b reflect.Value) bool {
		return strings.Compare(a.String(), b.String()) < 0
	})
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			var m *meta.Module
			if r.Key != nil {
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
				return yangLibModuleHandleNode(addresser, m), key, nil
			}
			return nil, nil, nil
		},
	}
}

func yangLibModuleHandleNode(addresser ModuleAddresser, m *meta.Module) node.Node {
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
				hnd.Val = &node.Value{Str: addresser(m)}
			case "namespace":
				hnd.Val = &node.Value{Str: m.Namespace}
			case "feature":
			case "conformance-type":
			}
			return nil
		},
	}
}
