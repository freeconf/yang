package node

import (
	"net/url"
	"strings"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/meta"
)

// Find navigates to another selector automatically applying constraints to returned selector.
// This supports paths that start with any number of "../" where FindUrl does not.
func (self Selection) Find(path string) Selection {
	p := path
	s := self
	for strings.HasPrefix(p, "../") {
		if s.Parent == nil {
			s.LastErr = fc.NotFoundError("No parent path to resolve " + p)
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
func (self Selection) FindUrl(url *url.URL) Selection {
	if self.LastErr != nil || url == nil {
		return self
	}
	var targetSlice PathSlice
	targetSlice, self.LastErr = ParseUrlPath(url, self.Meta())
	if self.LastErr != nil {
		return self
	}
	if len(url.Query()) > 0 {
		buildConstraints(&self, url.Query())
		if self.LastErr != nil {
			return self
		}
	}
	return self.FindSlice(targetSlice)
}

func (self Selection) FindSlice(xslice PathSlice) Selection {
	segs := xslice.Segments()
	sel := self
	for i := 0; i < len(segs); i++ {
		isLast := i == len(segs)-1
		if meta.IsAction(segs[i].meta) || meta.IsNotification(segs[i].meta) {
			if !isLast {
				return Selection{LastErr: fc.BadRequestError("Cannot select inside action or notification")}
			}
			childSel := sel
			childSel.Parent = &sel
			childSel.Path = segs[i]
			return childSel
		} else if meta.IsList(segs[i].meta) || meta.IsContainer(segs[i].meta) {
			r := &ChildRequest{
				Request: Request{
					Selection: sel,
					Target:    xslice.Tail,
				},
				Meta: segs[i].meta.(meta.HasDataDefs),
			}
			if sel = sel.Select(r); sel.IsNil() || sel.LastErr != nil {
				return sel
			}
			if meta.IsList(segs[i].meta) {
				if segs[i].key == nil {
					if !isLast {
						return Selection{LastErr: fc.BadRequestError("Cannot select inside list with key")}
					}
					break
				}
				r := &ListRequest{
					Request: Request{
						Selection: sel,
						Target:    xslice.Tail,
					},
					First: true,
					Meta:  segs[i].meta.(*meta.List),
					Key:   segs[i].key,
				}
				// not interested in key, should match seg[i].key in theory
				sel, _ = sel.SelectListItem(r)
			}
		} else if meta.IsLeaf(segs[i].meta) {
			return Selection{
				LastErr: fc.BadRequestError("Cannot select leaves"),
				Context: self.Context,
			}
		}
		if sel.LastErr != nil || sel.IsNil() {
			return sel
		}
	}
	return sel
}
