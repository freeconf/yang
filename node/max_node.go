
package node

import "github.com/dhubler/c2g/c2"

type MaxNode struct {
	Count int
	Max int
}

func (self MaxNode) CheckContainerPreConstraints(r *ContainerRequest, navigating bool) (bool, error) {
	if navigating {
		return true, nil
	}
	self.Count++
	if self.Count > self.Max  {
		r.ConstraintsHandler.IncompleteResponse(r.Selection.path)
		// FATAL
		return false, c2.NewErrC("Too many nodes", 413)
	}
	return true, nil
}
