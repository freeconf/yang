package meta

func HasChildren(parent MetaList) bool {
	i := Children(parent)
	return !i.HasNext()
}

// Len counts the items in an iterator
func Len(i Iterator) (len int) {
	for i.HasNext() {
		len++
		i.Next()
	}
	return
}

// GetPath as determined in the information model (e.g. YANG), not data model (e.g. RESTCONF)
func GetPath(m Meta) string {
	s := m.GetIdent()
	if p := m.GetParent(); p != nil {
		return GetPath(p) + "/" + s
	}
	return s
}

// Root finds root meta definition, which is the Module
func Root(m Meta) *Module {
	candidate := m
	for candidate.GetParent() != nil {
		candidate = candidate.GetParent()
	}
	return candidate.(*Module)
}
