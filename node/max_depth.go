package node

type MaxDepth struct {
	InitialDepth int
	MaxDepth int
}

func (self MaxDepth) CheckContainerPreConstraints(r *ContainerRequest, navigating bool) (bool, error) {
	if navigating {
		return true, nil
	}
	depth := r.Selection.Path.Len() + 1
	if depth - self.InitialDepth >= self.MaxDepth {
		r.ConstraintsHandler.IncompleteResponse(r.Selection.Path)
		// NON-FATAL
		return false, nil
	}
	return true, nil
}
