package node

import (
	"fmt"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
	"github.com/freeconf/yang/xpath"
)

type xpathImpl struct {
}

func (xp xpathImpl) resolvePath(seg *xpath.Path, s *Selection) (*Selection, error) {
	m := meta.Find(s.Meta().(meta.HasDefinitions), seg.Ident)
	if m == nil {
		return nil, fmt.Errorf("'%s' not found in xpath", seg.Ident)
	}
	if meta.IsContainer(m) {
		sel, err := s.Find(seg.Ident)
		if err != nil || sel == nil {
			return nil, err
		}
		return xp.resolvePath(seg.Next, sel)
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
		nextSeg := seg.Next
		for li.Selection != nil {
			if s, err = xp.resolvePath(nextSeg, li.Selection); s != nil || err != nil {
				return s, err
			}
			if li, err = li.Next(); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
	if meta.IsLeaf(m) {
		match, err := xp.resolveExpression(seg.Ident, seg.Expr, s)
		if err != nil || !match {
			return nil, err
		}
		return s, nil
	}
	panic("type not supported " + m.Ident())
}

func (xp xpathImpl) resolveExpression(name string, e xpath.Expression, sel *Selection) (bool, error) {
	switch x := e.(type) {
	case *xpath.Operator:
		return xp.resolveOperator(x, name, sel)
	}
	panic("unknown xpath expression")
}

func (xp xpathImpl) resolveOperator(oper *xpath.Operator, ident string, s *Selection) (bool, error) {
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

func (xp xpathImpl) resolveAbsolutePath(s *Selection) (*Selection, error) {
	found := s
	for found.parent != nil {
		var err error
		if found, err = found.parent.makeCopy(); err != nil {
			return nil, err
		}
	}
	return found, nil
}
