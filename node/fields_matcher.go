package node

type FieldsMatcher struct {
	expression string
	selector   PathMatcher
}

func NewFieldsMatcher(initialPath *Path, expression string) (fm *FieldsMatcher, err error) {
	fm = &FieldsMatcher{
		expression : expression,
	}
	return fm, nil
}

func (self *FieldsMatcher) CheckContainerPreConstraints(r *ContainerRequest, navigating bool) (bool, error) {
	if navigating {
		return true, nil
	} else  if self.selector == nil {
		if err := self.init(r.Selection.Path); err != nil {
			return false, err
		}
	}
	candidate := NewContainerPath(r.Selection.Path, r.Meta)
	return self.selector.PathMatches(candidate), nil
}

func (self *FieldsMatcher) init(root *Path) error {
	var err error
	if self.selector, err = ParsePathExpression(root, self.expression); err != nil {
		return err
	}
	return nil
}

func (self *FieldsMatcher) CheckFieldPreConstraints(r *FieldRequest, navigating bool) (bool, error) {
	if navigating {
		return true, nil
	} else if self.selector == nil {
		if err := self.init(r.Selection.Path); err != nil {
			return false, err
		}
	}
	return self.selector.FieldMatches(r.Selection.Path, r.Meta), nil
}
