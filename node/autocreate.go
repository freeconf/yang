package node

type AutoCreate struct {
}

func (self AutoCreate) CheckListPostConstraints(r ListRequest, child *Selection, key []*Value) (bool, error) {
	panic("not implemented")
	//if child == nil {
	//	r.New = true
	//	nextNode, r.Selection.path.key, _, err = r.Selection.node.Next(r.Selection, r)
	//	if err != nil {
	//		return nil, err
	//	} else if nextNode == nil {
	//		return nil, blit.NewErr("Could not autocreate list item for " + selection.path.String())
	//	}
	//
	//}
	//return true, nil
}

func (self AutoCreate) CheckContainerPostConstraints(r ContainerRequest, child *Selection) (bool, error) {
	panic("not implemented")
	//return true, nil
}