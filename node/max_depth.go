package node

import "github.com/freeconf/yang/meta"

type MaxDepth struct {
	MaxDepth int
}

func (md MaxDepth) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	return md.checkPathLen(r.Selection.Path, r.Base), nil
}

func (md MaxDepth) checkPathLen(current *Path, base *Path) bool {
	depth := 0
	p := current
	for p.Meta != base.Meta {
		isListItem := meta.IsList(p.Meta) && p.Parent.Meta == p.Meta
		if !isListItem {
			// lists have 2 entries in a path, list node and list item node
			depth++
		}
		if depth >= md.MaxDepth {
			return false
		}
		p = p.Parent
	}
	return true
}

func (md MaxDepth) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	return md.checkPathLen(r.Selection.Path, r.Base), nil
}
