package node

import "github.com/freeconf/yang/c2"

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
		return false, c2.HttpError(413)
	}
	return true, nil
}
