package node

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

// Find navigates to another selector automatically applying constraints to returned selector.
// This supports paths that start with any number of "../" where FindUrl does not.
func (sel *Selection) Find(path string) (*Selection, error) {
	p := path
	s := sel
	for strings.HasPrefix(p, "../") {
		if s.parent == nil {
			return nil, fmt.Errorf("%w. no parent path to resolve %s", fc.NotFoundError, p)
		}
		p = p[3:]
		s = s.Parent()
	}
	s, err := s.makeCopy()
	if err != nil {
		return nil, err
	}
	if qmark := strings.IndexRune(path, '?'); qmark >= 0 {
		u, err := url.Parse(p)
		if err != nil {
			return nil, err
		}
		if err = buildConstraints(s, u.Query()); err != nil {
			return nil, err
		}
		path = path[:qmark]
	}

	targetSlice, err := parseUrlPath(path, sel.Meta())
	if err != nil {
		return nil, err
	}
	return s.findSlice(targetSlice)
}

func (sel *Selection) findSlice(segs []*Path) (*Selection, error) {
	if len(segs) == 0 {
		return sel, nil
	}
	p := sel
	tail := segs[len(segs)-1]
	var err error
	for i := 0; i < len(segs); i++ {
		isLast := i == len(segs)-1
		if meta.IsAction(segs[i].Meta) || meta.IsNotification(segs[i].Meta) || meta.IsLeaf(segs[i].Meta) {
			if !isLast {
				return nil, fmt.Errorf("%w. Cannot select inside action, leaf or notification", fc.BadRequestError)
			}
			copy := *p
			copy.parent = p
			copy.Path = segs[i]
			return &copy, nil
		} else if meta.IsList(segs[i].Meta) || meta.IsContainer(segs[i].Meta) {
			r := &ChildRequest{
				Request: Request{
					Selection: p,
					Target:    tail,
				},
				Meta: segs[i].Meta.(meta.HasDataDefinitions),
			}
			if p, err = p.selekt(r); p == nil || err != nil {
				return nil, err
			}
			if meta.IsList(segs[i].Meta) {
				if segs[i].Key == nil {
					if !isLast {
						return nil, fmt.Errorf("%w. Cannot select inside list with key", fc.BadRequestError)
					}
					break
				}
				r := &ListRequest{
					Request: Request{
						Selection: p,
						Target:    tail,
					},
					First: true,
					Meta:  segs[i].Meta.(*meta.List),
					Key:   segs[i].Key,
				}
				// not interested in key, should match seg[i].key in theory
				var visible bool
				p, visible, _, err = p.selectListItem(r)
				if !visible || err != nil {
					return nil, err
				}
			}
		}
	}
	return p, nil
}
