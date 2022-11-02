package node

import (
	"context"

	"github.com/freeconf/yang/meta"
)

// Browser is a handle to a data source and starting point for interfacing with any freeconf enabled interface.
// It's the starting point to the top-most selection, or the Root().
type Browser struct {

	// Information model of browser
	Meta *meta.Module

	// Regsitry of listeners when data model under browser is modified
	Triggers *TriggerTable

	// True if you want no leaf data checks like pattern, length, range, etc
	// you would only want to do this if you had a good reason
	DisableConstraints bool

	// Function to get data model behind browser
	src func() Node
}

// Root is top-most selection.  From here you can use Find to navigate to other parts of
// application or any of the Insert command to start getting data in or out.
func (self *Browser) Root() Selection {
	return Selection{
		Browser:     self,
		Path:        &Path{meta: self.Meta},
		Node:        self.src(),
		Constraints: self.baseConstraints(),
		Context:     context.Background(),
	}
}

// Root is top-most selection.  From here you can use Find to navigate to other parts of
// application or any of the Insert command to start getting data in or out.
func (self *Browser) RootWithContext(ctx context.Context) Selection {
	return Selection{
		Browser:     self,
		Path:        &Path{meta: self.Meta},
		Node:        self.src(),
		Constraints: self.baseConstraints(),
		Context:     ctx,
	}
}

func (self *Browser) baseConstraints() *Constraints {
	c := &Constraints{}
	c.AddConstraint("~when", 100, 0, CheckWhen{})
	if !self.DisableConstraints {
		c.AddConstraint("field", 100, 0, fieldConstraints{})
	}
	return c
}

// NewBrowserSource unites a model (MetaList) with a data source (Node).  Here the node instance
// is requested for each browse operation allowing the node state to be fresh for each request.
func NewBrowserSource(m *meta.Module, src func() Node) *Browser {
	return &Browser{
		Meta:     m,
		Triggers: NewTriggerTable(),
		src:      src,
	}
}

// NewBrowser  obviously does not resolve the source node for each new selection
// so the state of at least the root node is shared for all subsequent operations.
// In short, either do not keep a copy of this very browser for very long or know
// what you're doing
func NewBrowser(m *meta.Module, n Node) *Browser {
	return &Browser{
		Meta:     m,
		Triggers: NewTriggerTable(),
		src: func() Node {
			return n
		},
	}
}
