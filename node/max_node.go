
package node

import "github.com/blitter/blit"

type MaxNode struct {
	Count int
	Max int
}

func (self MaxNode) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	self.Count++
	if self.Count > self.Max  {
		r.Context.IncompleteResponse(r.Selection.path)
		// FATAL
		return false, blit.NewErrC("Too many nodes", 413)
	}
	return true, nil
}
