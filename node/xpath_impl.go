package node

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/xpath"
)

type xpathImpl struct {
}

func (self xpathImpl) resolveSegment(r xpathResolver, seg *xpath.Segment, s Selection) Selection {
	c2.Debug.Print(seg.Ident)
	m := meta.FindByIdent2(s.Meta(), seg.Ident)
	if meta.IsContainer(m) || meta.IsList(m) {
		if s.InsideList {
			panic("here")
		}
		return s.Find(seg.Ident)
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
	m := meta.FindByIdent2(s.Meta(), ident).(meta.HasDataType)
	b, err := SetValue(m.GetDataType(), oper.Lhs)
	if err != nil {
		return false, err
	}
	a, err := s.GetValue(ident)
	if err != nil {
		return false, err
	}
	switch oper.Oper {
	case "=":
		return a.Value() == b.Value(), nil
	case "!=":
		return a.Value() != b.Value(), nil
	case "<":
		return a.Compare(b) < 0, nil
	case ">":
		return a.Compare(b) > 0, nil
	case ">=":
		return a.Compare(b) >= 0, nil
	case "<=":
		return a.Compare(b) <= 0, nil
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
