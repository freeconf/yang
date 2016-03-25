package node

type FieldsMatcher struct {
	Selector   PathMatcher
}

func NewFieldsMatcher(initialPath *Path, expression string) (fm *FieldsMatcher, err error) {
	fm = &FieldsMatcher{}
	if fm.Selector, err = ParsePathExpression(initialPath, expression); err != nil {
		return nil, err
	}
	return fm, nil
}

func (self *FieldsMatcher) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	return self.Selector.PathMatches(r.Selection.path), nil
}

func (self *FieldsMatcher) CheckFieldPreConstraints(r *FieldRequest) (bool, error) {
	return self.Selector.FieldMatches(r.Selection.path, r.Meta), nil
}
