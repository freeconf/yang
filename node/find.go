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
func (sel Selection) Find(path string) Selection {
	p := path
	s := sel
	for strings.HasPrefix(p, "../") {
		if s.Parent == nil {
			s.LastErr = fmt.Errorf("%w. no parent path to resolve %s", fc.NotFoundError, p)
			return s
		}
		s = *s.Parent
		p = p[3:]
	}
	var u *url.URL
	u, s.LastErr = url.Parse(p)
	if s.LastErr != nil {
		return s
	}
	return s.FindUrl(u)
}

// FindUrl navigates to another selection with possible constraints as url parameters.  Constraints
// are added to any existing contraints.  Original selector and constraints will remain unaltered
func (sel Selection) FindUrl(url *url.URL) Selection {
	if sel.LastErr != nil || url == nil {
		return sel
	}
	var targetSlice PathSlice
	targetSlice, sel.LastErr = ParseUrlPath(url, sel.Meta())
	if sel.LastErr != nil {
		return Selection{LastErr: sel.LastErr}
	}
	if len(url.Query()) > 0 {
		buildConstraints(&sel, url.Query())
		if sel.LastErr != nil {
			return Selection{LastErr: sel.LastErr}
		}
	}
	return sel.FindSlice(targetSlice)
}

func (sel Selection) FindSlice(xslice PathSlice) Selection {
	segs := xslice.Segments()
	copy := sel
	for i := 0; i < len(segs); i++ {
		isLast := i == len(segs)-1
		if meta.IsAction(segs[i].Meta) || meta.IsNotification(segs[i].Meta) {
			if !isLast {
				err := fmt.Errorf("%w. Cannot select inside action or notification", fc.BadRequestError)
				return Selection{LastErr: err}
			}
			childSel := copy
			childSel.Parent = &copy
			childSel.Path = segs[i]
			return childSel
		} else if meta.IsList(segs[i].Meta) || meta.IsContainer(segs[i].Meta) {
			r := &ChildRequest{
				Request: Request{
					Selection: copy,
					Target:    xslice.Tail,
				},
				Meta: segs[i].Meta.(meta.HasDataDefinitions),
			}
			if copy = copy.selekt(r); copy.IsNil() || copy.LastErr != nil {
				return copy
			}
			if meta.IsList(segs[i].Meta) {
				if segs[i].Key == nil {
					if !isLast {
						err := fmt.Errorf("%w. Cannot select inside list with key", fc.BadRequestError)
						return Selection{LastErr: err}
					}
					break
				}
				r := &ListRequest{
					Request: Request{
						Selection: copy,
						Target:    xslice.Tail,
					},
					First: true,
					Meta:  segs[i].Meta.(*meta.List),
					Key:   segs[i].Key,
				}
				// not interested in key, should match seg[i].key in theory
				copy, _ = copy.selectListItem(r)
			}
		} else if meta.IsLeaf(segs[i].Meta) {
			return Selection{
				LastErr: fmt.Errorf("%w. Cannot select leaves", fc.BadRequestError),
				Context: sel.Context,
			}
		}
		if copy.LastErr != nil || copy.IsNil() {
			return copy
		}
	}
	return copy
}
