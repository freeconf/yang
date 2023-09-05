package node

import "github.com/freeconf/yang/xpath"

func (sel *Selection) XFind(path xpath.Path) (*Selection, error) {
	p := sel
	var err error
	xp := path
	r := xpathResolver{impl: xpathImpl{}}
	for xp != nil {
		p, err = r.resolvePath(xp, p)
		if p == nil || err != nil {
			return nil, err
		}
		xp = xp.Next()
	}
	return p, nil
}

func (sel *Selection) XPredicate(p xpath.Path) (bool, error) {
	found, err := sel.XFind(p)
	return found != nil, err
}

type xpathResolver struct {
	impl xpathInterpretter
}

func (resolver xpathResolver) resolvePath(p xpath.Path, sel *Selection) (*Selection, error) {
	switch x := p.(type) {
	case *xpath.Segment:
		return resolver.impl.resolveSegment(resolver, x, sel)
	case *xpath.AbsolutePath:
		return resolver.impl.resolveAbsolutePath(resolver, sel)
	}
	panic("unknown xpath segment type")
}

func (resolver xpathResolver) resolveExpression(name string, e xpath.Expression, sel *Selection) (bool, error) {
	switch x := e.(type) {
	case *xpath.Operator:
		return resolver.impl.resolveOperator(resolver, x, name, sel)
	}
	panic("unknown xpath expression")
}

type xpathInterpretter interface {
	resolveAbsolutePath(r xpathResolver, s *Selection) (*Selection, error)
	resolveSegment(r xpathResolver, seg *xpath.Segment, s *Selection) (*Selection, error)
	resolveOperator(r xpathResolver, oper *xpath.Operator, ident string, s *Selection) (bool, error)
}
