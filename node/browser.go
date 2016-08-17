package node

import "github.com/dhubler/c2g/meta"

type Browser struct {
	Meta     meta.MetaList
	Triggers *TriggerTable
	getNode  func() Node
}

func (self *Browser) Root() *Selection {
	return &Selection{
		browser: self,
		path:    &Path{meta: self.Meta},
		node:    self.getNode(),
	}
}

func NewBrowser(m meta.MetaList, src func() Node) *Browser {
	return &Browser{
		Meta:     m,
		Triggers: NewTriggerTable(),
		getNode:  src,
	}
}

// NewBrowser2 obviously does not resolve the source node for each new selection
// so the state of at least the root node is shared for all subsequent operations.
// In short, either do not keep a copy of this very browser for very long or know
// what you're doing
func NewBrowser2(m meta.MetaList, src Node) *Browser {
	return &Browser{
		Meta:     m,
		Triggers: NewTriggerTable(),
		getNode: func() Node {
			return src
		},
	}
}
