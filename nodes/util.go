package nodes

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

func Singleton(name string, f ChildFunc) node.Node {
	return &Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			if r.Meta.GetIdent() == name {
				return f(r)
			}
			return nil, nil
		},
	}
}

// Copys meta while expanding all groups and typedefs.  This has the potentional
// to dramatically increase the size of your meta and more dangerously, go into infinite
// recursion on recursive metas
func DecoupledMetaCopy(yangPath meta.StreamSource, src meta.MetaList) (meta.MetaList, error) {
	yangModule := yang.RequireModule(yangPath, "yang")
	var copy meta.MetaList
	m, err := meta.FindByPath(yangModule, "module/definitions")
	if err != nil {
		return nil, err
	}
	if meta.IsList(src) {
		m, err = meta.FindByIdentExpandChoices(m, "list")
		if err != nil {
			return nil, err
		}
		copy = &meta.List{}
	} else {
		m, err = meta.FindByIdentExpandChoices(m, "container")
		if err != nil {
			return nil, err
		}
		copy = &meta.Container{}
	}
	srcNode := SchemaData{true}.MetaList(src)
	destNode := SchemaData{true}.MetaList(copy)
	err = node.NewBrowser(m.(meta.MetaList), srcNode).Root().InsertInto(destNode).LastErr
	return copy, err
}
