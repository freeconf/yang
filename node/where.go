package node

import (
	"fmt"

	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/val"
	"github.com/freeconf/yang/xpath"
)

// RFC Draft for list pagination
// see https://www.ietf.org/id/draft-ietf-netconf-list-pagination-rc-01.html#name-the-where-query-parameter

type Where struct {
	Filter      string // XPath filter
	xpathFilter xpath.Path
}

func NewWhere(filter string) (*Where, error) {
	p, err := xpath.Parse(filter)
	if err != nil {
		return nil, fmt.Errorf("%w. invalid xpath expression: %s", fc.BadRequestError, err)
	}
	return &Where{xpathFilter: p, Filter: filter}, nil
}

func (w *Where) CheckListPostConstraints(r ListRequest, child Selection, key []val.Value) (bool, bool, error) {
	if child.InsideList {
		match, err := child.XPredicate(w.xpathFilter)
		fmt.Printf("path=%s, match=%v, err=%s\n", child.Path.StringNoModule(), match, err)
		return true, match, err
	}
	return true, true, nil
}

// func (w *Where) CheckContainerPostConstraints(r ChildRequest, child Selection) (bool, error) {
// 	if r.Path == nil {
// 		fmt.Printf("nil\n")
// 		return true, nil
// 	}
// 	//match, err := child.XPredicate(w.xpathFilter)
// 	//fmt.Printf("path=%s, match=%v, err=%s\n", r.Path.StringNoModule(), match, err)
// 	return true, nil
// }
