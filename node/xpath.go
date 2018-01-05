package node

import "github.com/freeconf/gconf/xpath"

func (self Selection) XFind(path xpath.Path) Selection {
	sel := self
	p := path
	r := xpathResolver{impl: xpathImpl{}}
	for p != nil {
		found := r.resolvePath(p, sel)
		if found.IsNil() || found.LastErr != nil {
			return found
		}
		p = p.Next()
		sel = found
	}
	return sel
}

func (self Selection) XPredicate(p xpath.Path) (bool, error) {
	found := self.XFind(p)
	if found.LastErr != nil {
		return false, found.LastErr
	}
	return !found.IsNil(), nil
}

type xpathResolver struct {
	impl xpathInterpretter
}

func (self xpathResolver) resolvePath(p xpath.Path, sel Selection) Selection {
	switch x := p.(type) {
	case *xpath.Segment:
		return self.impl.resolveSegment(self, x, sel)
	case *xpath.AbsolutePath:
		return self.impl.resolveAbsolutePath(self, sel)
	}
	panic("unknown xpath segment type")
}

func (self xpathResolver) resolveExpression(name string, e xpath.Expression, sel Selection) (bool, error) {
	switch x := e.(type) {
	case *xpath.Operator:
		return self.impl.resolveOperator(self, x, name, sel)
	}
	panic("unknown xpath expression")
}

type xpathInterpretter interface {
	resolveAbsolutePath(r xpathResolver, s Selection) Selection
	resolveSegment(r xpathResolver, seg *xpath.Segment, s Selection) Selection
	resolveOperator(r xpathResolver, oper *xpath.Operator, ident string, s Selection) (bool, error)
}
