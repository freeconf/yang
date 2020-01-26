package node

import "github.com/freeconf/yang/fc"

type MaxNode struct {
	Count int
	Max   int
}

func (self MaxNode) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	self.Count++
	if self.Count > self.Max {
		return false, fc.ConflictError
	}
	return true, nil
}
