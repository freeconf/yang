package node

import "github.com/c2stack/c2g/meta"

// Browser is constructor of root-most selection together with managing triggers
// for data operations.
type Browser struct {
	Meta     meta.MetaList
	//Triggers *TriggerTable
	getNode  func() Node
}

// Single root selector capable of leading to all other selectors in
// the data tree.
func (self *Browser) Root() Selection {
	return Selection{
		Browser: self,
		Path:    &Path{meta: self.Meta},
		Node:    self.getNode(),
		Constraints: &Constraints{},
	}
}

// NewBrowser unites a model to a data source, and the data source can create
// a new node for each request to ensure new state is used starting with
// the root data node.
func NewBrowser(m meta.MetaList, src func() Node) *Browser {
	return &Browser{
		Meta:     m,
		//Triggers: NewTriggerTable(),
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
		//Triggers: NewTriggerTable(),
		getNode:  func() Node {
			return src
		},
	}
}
