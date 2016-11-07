package node

type FieldsMatcher struct {
	expression string
	selector   PathMatcher
}

func NewFieldsMatcher(expression string) (fm *FieldsMatcher, err error) {
	fm = &FieldsMatcher{
		expression: expression,
	}
	fm.selector, err = ParsePathExpression(expression)
	return
}

func (self *FieldsMatcher) CheckContainerPreConstraints(r *ContainerRequest) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	return self.selector.PathMatches(r.Base, r.Path), nil
}

// func (self *FieldsMatcher) CheckListPreConstraints(r *ListRequest, navigating bool) (bool, error) {
// 	// "fields" constraint doesn't control items in list, but we take this opportunity to initialize the root
// 	// path if it's a list
// 	if !navigating && self.selector == nil {
// 		if err := self.init(r.Selection.Path); err != nil {
// 			return false, err
// 		}
// 	}
// 	return true, nil
// }

func (self *FieldsMatcher) CheckFieldPreConstraints(r *FieldRequest, hnd *ValueHandle) (bool, error) {
	if r.IsNavigation() {
		return true, nil
	}
	return self.selector.PathMatches(r.Base, r.Path), nil
}
