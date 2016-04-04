
package node

import "github.com/c2g/c2"

type MaxNode struct {
	Count int
	Max int
}

func (self MaxNode) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	self.Count++
	if self.Count > self.Max  {
		r.Context.Handler().IncompleteResponse(r.Selection.path)
		// FATAL
		return false, c2.NewErrC("Too many nodes", 413)
	}
	return true, nil
}
