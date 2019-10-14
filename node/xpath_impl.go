package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
	"github.com/freeconf/yang/xpath"
)

type xpathImpl struct {
}

func (self xpathImpl) resolveSegment(r xpathResolver, seg *xpath.Segment, s Selection) Selection {
	m := meta.Find(s.Meta().(meta.HasDefinitions), seg.Ident)
	if m == nil {
		return Selection{LastErr: fmt.Errorf("'%s' not found in xpath", seg.Ident)}
	}
	if meta.IsContainer(m) {
		return r.resolvePath(seg.Next(), s.Find(seg.Ident))
	}
	if meta.IsList(m) {
		s := s.Find(seg.Ident)
		li := s.First()
		nextSeg := seg.Next()
		for !li.Selection.IsNil() {
			if s = r.resolvePath(nextSeg, li.Selection); !s.IsNil() || s.LastErr != nil {
				return s
			}
			li = li.Next()
		}
		return Selection{}
	}
	if meta.IsLeaf(m) {
		match, err := r.resolveExpression(seg.Ident, seg.Expr, s)
		if err != nil {
			return s.Split(ErrorNode{Err: err})
		}
		if !match {
			return Selection{}
		}
		return s
	}
	panic("type not supported " + m.Ident())
}

func (self xpathImpl) resolveOperator(r xpathResolver, oper *xpath.Operator, ident string, s Selection) (bool, error) {
	m := meta.Find(s.Meta().(meta.HasDefinitions), ident)
	if m == nil {
		return false, fmt.Errorf("'%s' not found in xpath", ident)
	}
	b, err := NewValue(m.(meta.HasType).Type(), oper.Lhs)
	if err != nil {
		return false, err
	}
	a, err := s.GetValue(ident)
	if err != nil {
		return false, err
	}
	switch oper.Oper {
	case "=":
		return val.Equal(a, b), nil
	case "!=":
		return !val.Equal(a, b), nil
	default:
		c := a.(val.Comparable).Compare(b.(val.Comparable))
		switch oper.Oper {
		case "<":
			return c < 0, nil
		case ">":
			return c > 0, nil
		case ">=":
			return c >= 0, nil
		case "<=":
			return c <= 0, nil
		}
	}
	panic("unrecognized operator: " + oper.Oper)
}

func (self xpathImpl) resolveAbsolutePath(r xpathResolver, s Selection) Selection {
	found := s
	if found.Parent != nil {
		found = *found.Parent
	}
	return found
}
