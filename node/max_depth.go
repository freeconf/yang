package node

import (
	"github.com/c2stack/c2g/c2"
)

type MaxDepth struct {
	MaxDepth int
}

func (self MaxDepth) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	depth := r.Selection.Path.Len() - r.Base.Len()
	if depth >= self.MaxDepth {
		return false, c2.NewErrC("response for request too large", 413)
	}
	return true, nil
}
