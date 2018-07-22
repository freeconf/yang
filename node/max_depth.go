package node

type MaxDepth struct {
	MaxDepth int
}

func (self MaxDepth) CheckContainerPreConstraints(r *ChildRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	depth := r.Selection.Path.Len() - r.Base.Len()
	if depth >= self.MaxDepth {
		return false, nil
	}
	return true, nil
}
