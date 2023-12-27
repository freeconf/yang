package node

import (
	"github.com/freeconf/yang/xpath"
)

func (sel *Selection) XFind(path *xpath.Path) (*Selection, error) {
	p := sel
	var err error
	xp := path
	r := xpathImpl{}
	for xp != nil {
		p, err = r.resolvePath(xp, p)
		if p == nil || err != nil {
			return nil, err
		}
		xp = xp.Next
	}
	return p, nil
}

func (sel *Selection) XPredicate(p *xpath.Path) (bool, error) {
	found, err := sel.XFind(p)
	return found != nil, err
}
