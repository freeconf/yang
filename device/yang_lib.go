package device

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/node"
)

type ResolveModule interface {
	ResolveModuleHnd(hnd ModuleHnd) (*meta.Module, error)
}

func LoadModules(ietfYangLib *node.Browser, resolver ResolveModule) (map[string]*meta.Module, error) {
	mods := make(map[string]*meta.Module)
	n := yangLibClientNode(mods, resolver)
	if err := ietfYangLib.Root().Find("modules-state/module").InsertInto(n).LastErr; err != nil {
		return nil, err
	}
	return mods, nil
}

func yangLibClientNode(mods map[string]*meta.Module, resolver ResolveModule) node.Node {
	var hnd ModuleHnd
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			if r.New {
				hnd = ModuleHnd{Name: r.Key[0].Str}
				return node.ReflectNode(&hnd), key, nil
			}
			return nil, nil, nil
		},
		OnEndEdit: func(r node.NodeRequest) error {
			mod, err := resolver.ResolveModuleHnd(hnd)
			if err != nil {
				return err
			}
			mods[mod.GetIdent()] = mod
			return nil
		},
	}
}

type ModuleHnd struct {
	Name      string
	Schema    string
	Revision  string
	Namespace string
}
