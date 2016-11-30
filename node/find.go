package node

import (
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

func (self Selection) FindSlice(xslice PathSlice) Selection {
	segs := xslice.Segments()
	sel := self
	for i := 0; i < len(segs); i++ {
		isLast := i == len(segs)-1
		if meta.IsAction(segs[i].meta) || meta.IsNotification(segs[i].meta) {
			if !isLast {
				return Selection{LastErr: c2.NewErrC("Cannot select inside action or notification", 400)}
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
				Meta: segs[i].meta.(meta.MetaList),
			}
			if sel = sel.Select(r); sel.IsNil() || sel.LastErr != nil {
				return sel
			}
			if meta.IsList(segs[i].meta) {
				if segs[i].key == nil {
					if !isLast {
						return Selection{LastErr: c2.NewErrC("Cannot select inside list with key", 400)}
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
			return Selection{LastErr: c2.NewErrC("Cannot select leaves", 400)}
		}
		if sel.LastErr != nil || sel.IsNil() {
			return sel
		}
	}
	return sel
}
