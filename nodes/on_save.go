package nodes

import (
	"github.com/freeconf/gconf/node"
)

func OnSave(orig node.Selection, onSave func(node.Selection) error) (node.Node, error) {
	temp := ReflectChild(make(map[string]interface{}))
	if err := orig.InsertInto(temp).LastErr; err != nil {
		return nil, err
	}
	return &Extend{
		Base: orig.Node,
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			if r.EditRoot {
				return onSave(r.Selection.Split(temp))
			}
			return nil
		},
	}, nil
}
