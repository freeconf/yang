package node

func Diff(a Node, b Node) Node {
	n := &MyNode{}
	n.OnSelect = func(r ContainerRequest) (n Node, err error) {
		var aNode, bNode Node
		r.New = false
		if aNode, err = a.Select(r); err != nil {
			return nil, err
		}
		if bNode, err = b.Select(r); err != nil {
			return nil, err
		}
		if aNode == nil {
			return nil, nil
		}
		if bNode == nil {
			return aNode, nil
		}
		return Diff(aNode, bNode), nil
	}
	n.OnNext = func(r ListRequest) (Node, []*Value, error) {
		var err error
		var aNode, bNode Node
		var aKey []*Value
		r.New = false
		if aNode, aKey, err = a.Next(r); err != nil {
			return nil, nil, err
		}
		if bNode, _, err = b.Next(r); err != nil {
			return nil, nil, err
		}
		if aNode == nil {
			return nil, nil, nil
		}
		if bNode == nil {
			return aNode, aKey, nil
		}

		// TODO: compare keys?

		return Diff(aNode, bNode), aKey, nil
	}
	n.OnRead = func(r FieldRequest) (changedValue *Value, err error) {
		var aVal, bVal *Value
		if aVal, err = a.Read(r); err != nil {
			return nil, err
		}
		if bVal, err = b.Read(r); err != nil {
			return nil, err
		}
		if aVal == nil {
			if bVal == nil {
				return nil, nil
			}
			return bVal, nil
		}
		if aVal.Equal(bVal) {
			return nil, nil
		}
		return aVal, nil
	}
	return n
}
