package meta

func HasChildren(parent MetaList) bool {
	i := Children(parent)
	return !i.HasNext()
}

func ListLen(parent MetaList) (len int) {
	i := Children(parent)
	for i.HasNext() {
		len++
		i.Next()
	}
	return
}

func ListLenNoExpand(parent MetaList) (len int) {
	i := Children(parent)
	for i.HasNext() {
		len++
		i.Next()
	}
	return
}

// GetPath as determined in information model (not data model)
func GetPath(m Meta) string {
	s := m.GetIdent()
	if p := m.GetParent(); p != nil {
		return GetPath(p) + "/" + s
	}
	return s
}

// GetModule finds root meta definition, which is the Module
func GetModule(m Meta) *Module {
	candidate := m
	for candidate.GetParent() != nil {
		candidate = candidate.GetParent()
	}
	return candidate.(*Module)
}

func moveModuleMeta(dest *Module, src *Module) error {
	iters := []Iterator{
		ChildrenNoResolve(src.GetGroupings()),
		ChildrenNoResolve(src.GetTypedefs()),
		ChildrenNoResolve(src.DataDefs()),
	}
	for _, iter := range iters {
		for iter.HasNext() {
			if m, err := iter.Next(); err != nil {
				return err
			} else {
				dest.AddMeta(m)
			}
		}
	}
	return nil
}
