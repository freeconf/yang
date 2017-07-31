package node

import (
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
	"github.com/c2stack/c2g/xpath"
)

type xpathImpl struct {
}

func (self xpathImpl) resolveSegment(r xpathResolver, seg *xpath.Segment, s Selection) Selection {
	m, err := meta.FindByIdent2(s.Meta(), seg.Ident)
	if err != nil {
		return Selection{Node: ErrorNode{Err: err}}
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
	panic("type not supported " + m.GetIdent())
}

func (self xpathImpl) resolveOperator(r xpathResolver, oper *xpath.Operator, ident string, s Selection) (bool, error) {
	m, err := meta.FindByIdent2(s.Meta(), ident)
	if err != nil {
		return false, err
	}
	b, err := NewValue(m.(meta.HasDataType).GetDataType(), oper.Lhs)
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
