package node

type MaxDepth struct {
	MaxDepth int
}

func (self MaxDepth) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	depth := r.Selection.Path.Len() - r.Base.Len()
	if depth >= self.MaxDepth {
		r.ConstraintsHandler.IncompleteResponse(r.Selection.Path)
		// NON-FATAL
		return false, nil
	}
	return true, nil
}
