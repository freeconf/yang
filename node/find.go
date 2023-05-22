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
		if s.Parent == nil {
			return nil, fmt.Errorf("%w. no parent path to resolve %s", fc.NotFoundError, p)
		}
		s = s.Parent
		p = p[3:]
	}
	u, err := url.Parse(p)
	if err != nil {
		return nil, err
	}
	return s.FindUrl(u)
}

// FindUrl navigates to another selection with possible constraints as url parameters.  Constraints
// are added to any existing contraints.  Original selector and constraints will remain unaltered
func (sel *Selection) FindUrl(url *url.URL) (*Selection, error) {
	if url == nil {
		return nil, nil
	}
	targetSlice, err := ParseUrlPath(url, sel.Meta())
	if err != nil {
		return nil, err
	}
	copy := *sel
	if len(url.Query()) > 0 {
		if err = buildConstraints(&copy, url.Query()); err != nil {
			return nil, err
		}
	}
	return copy.FindSlice(targetSlice)
}

func (sel *Selection) FindSlice(xslice PathSlice) (*Selection, error) {
	segs := xslice.Segments()
	p := sel
	var err error
	for i := 0; i < len(segs); i++ {
		isLast := i == len(segs)-1
		if meta.IsAction(segs[i].Meta) || meta.IsNotification(segs[i].Meta) || meta.IsLeaf(segs[i].Meta) {
			if !isLast {
				return nil, fmt.Errorf("%w. Cannot select inside action, leaf or notification", fc.BadRequestError)
			}
			copy := *p
			copy.Parent = p
			copy.Path = segs[i]
			return &copy, nil
		} else if meta.IsList(segs[i].Meta) || meta.IsContainer(segs[i].Meta) {
			r := &ChildRequest{
				Request: Request{
					Selection: p,
					Target:    xslice.Tail,
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
						Target:    xslice.Tail,
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
