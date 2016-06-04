package node

import "github.com/c2g/meta"

type Browser func() *Selection

func NewBrowser(m meta.MetaList, src func() Node) Browser {
	return func() *Selection {
		return Select(m, src())
	}
}
