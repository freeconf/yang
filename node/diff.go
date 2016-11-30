package node

func Diff(a Node, b Node) Node {
	n := &MyNode{}
	n.OnChild = func(r ChildRequest) (n Node, err error) {
		var aNode, bNode Node
		r.New = false
		if aNode, err = a.Child(r); err != nil {
			return nil, err
		}
		if bNode, err = b.Child(r); err != nil {
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
	n.OnField = func(r FieldRequest, hnd *ValueHandle) (err error) {
		if err = a.Field(r, hnd); err != nil {
			return err
		}
		aVal := hnd.Val
		if err = b.Field(r, hnd); err != nil {
			return err
		}
		bVal := hnd.Val
		if aVal == nil {
			if bVal == nil {
				return nil
			}
			hnd.Val = bVal
			return nil
		}
		if aVal.Equal(bVal) {
			hnd.Val = nil
			return nil
		}
		hnd.Val = aVal
		return nil
	}
	return n
}
