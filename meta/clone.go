package meta

import (
	"fmt"
)

func Copy(m Meta, deep bool) Meta {
	switch t := m.(type) {
	case *Leaf:
		x := *t
		x.DataType = cloneDataType(&x, x.DataType)
		return &x
	case *LeafList:
		x := *t
		x.DataType = cloneDataType(&x, x.DataType)
		return &x
	case *Any:
		x := *t
		return &x
	case *Container:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *List:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *Uses:
		x := *t
		// TODO: Uses will eventually have children, when that happens, uncomment this
		//deepCloneList(&x, &x)
		return &x
	case *Grouping:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *Rpc:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *RpcInput:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *RpcOutput:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *Notification:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *Module:
		x := *t
		if deep {
			deepCloneList(&x, &x.Defs)
			deepCloneList(&x, &x.Groupings)
			deepCloneList(&x, &x.Typedefs)
		}
		return &x
	case *Choice:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *ChoiceCase:
		x := *t
		if deep {
			deepCloneList(&x, &x)
		}
		return &x
	case *Refine:
		x := *t
		return &x
	}
	panic(fmt.Sprintf("clone not implemented for type %T", m))
}

func deepCloneList(p MetaList, src MetaList) {
	i := src.GetFirstMeta()
	p.Clear()
	for i != nil {
		copy := Copy(i, true)
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
