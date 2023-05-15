package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
	"github.com/freeconf/yang/xpath"
)

type xpathImpl struct {
}

func (self xpathImpl) resolveSegment(r xpathResolver, seg *xpath.Segment, s *Selection) (*Selection, error) {
	m := meta.Find(s.Meta().(meta.HasDefinitions), seg.Ident)
	if m == nil {
		return nil, fmt.Errorf("'%s' not found in xpath", seg.Ident)
	}
	if meta.IsContainer(m) {
		sel, err := s.Find(seg.Ident)
		if err != nil || sel == nil {
			return nil, err
		}
		return r.resolvePath(seg.Next(), sel)
	}
	if meta.IsList(m) {
		sel, err := s.Find(seg.Ident)
		if sel == nil || err != nil {
			return nil, err
		}
		li, err := sel.First()
		if err != nil {
			return nil, err
		}
		nextSeg := seg.Next()
		for li.Selection != nil {
			if s, err = r.resolvePath(nextSeg, li.Selection); s != nil || err != nil {
				return s, err
			}
			if li, err = li.Next(); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
	if meta.IsLeaf(m) {
		match, err := r.resolveExpression(seg.Ident, seg.Expr, s)
		if err != nil || !match {
			return nil, err
		}
		return s, nil
	}
	panic("type not supported " + m.Ident())
}

func (self xpathImpl) resolveOperator(r xpathResolver, oper *xpath.Operator, ident string, s *Selection) (bool, error) {
	m := meta.Find(s.Meta().(meta.HasDefinitions), ident)
	if m == nil {
		return false, fmt.Errorf("'%s' not found in xpath", ident)
	}
	b, err := NewValue(m.(meta.HasType).Type(), oper.Lhs)
	if err != nil {
		return false, err
	}
	s, err = s.Find(ident)
	if err != nil {
		return false, err
	}
	a, err := s.Get()
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

func (self xpathImpl) resolveAbsolutePath(r xpathResolver, s *Selection) (*Selection, error) {
	found := s
	for found.Parent != nil {
		found = found.Parent
	}
	return found, nil
}
