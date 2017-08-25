package meta

func DeepCopy(m Meta) Meta {
	var c Meta
	switch t := m.(type) {
	case *Leaf:
		x := *t
		x.DataType = cloneDataType(&x, x.DataType)
		c = &x
	case *LeafList:
		x := *t
		x.DataType = cloneDataType(&x, x.DataType)
		c = &x
	case *Any:
		x := *t
		c = &x
	case *Container:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *List:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *Uses:
		x := *t
		// TODO: Uses will eventually have children, when that happens, uncomment this
		//deepCloneList(&x, &x)
		c = &x
	case *Grouping:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *Rpc:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *RpcInput:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *RpcOutput:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *Notification:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *Module:
		x := *t
		deepCloneList(&x, &x.Defs)
		deepCloneList(&x, &x.Groupings)
		deepCloneList(&x, &x.Typedefs)
		c = &x
	case *Choice:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	case *ChoiceCase:
		x := *t
		deepCloneList(&x, &x)
		c = &x
	}
	return c
}

func deepCloneList(p MetaList, src MetaList) {
	i := src.GetFirstMeta()
	p.Clear()
	for i != nil {
		copy := DeepCopy(i)
		p.AddMeta(copy)
		i = i.GetSibling()
	}
}

func cloneDataType(parent HasDataType, dt *DataType) *DataType {
	if dt == nil {
		return nil
	}
	copy := *dt
	copy.Parent = parent
	copy.resolvedPtr = nil
	return &copy
}
