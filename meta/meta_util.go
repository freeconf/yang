package meta

import (
	"strings"
	"unicode"
)

func FindByIdent(i MetaIterator, ident string) (Meta, error) {
	child, err := i.NextMeta()
	if err != nil {
		return nil, err
	}
	for child != nil {
		if child.GetIdent() == ident {
			return child, nil
		}
		child, err = i.NextMeta()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func FindByIdent2(parent Meta, ident string) (Meta, error) {
	i := NewMetaListIterator(parent, true)
	return FindByIdent(i, ident)
}

func FindByIdentExpandChoices(m Meta, ident string) (Meta, error) {
	parent, isMetaList := m.(MetaList)
	if !isMetaList {
		return nil, nil
	}
	i := NewMetaListIterator(parent, true)
	var choice *Choice
	var isChoice bool
	for i.HasNextMeta() {
		child, err := i.NextMeta()
		if err != nil {
			return nil, err
		}
		choice, isChoice = child.(*Choice)
		if isChoice {
			cases := NewMetaListIterator(choice, false)
			for cases.HasNextMeta() {
				ccase, err := cases.NextMeta()
				if err != nil {
					return nil, err
				}
				found, err := FindByIdentExpandChoices(ccase.(*ChoiceCase), ident)
				if found != nil || err != nil {
					return found, err
				}
			}
		} else {
			if child.GetIdent() == ident {
				return child, nil
			}
		}
	}
	return nil, nil
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

func moveModuleMeta(dest *Module, src *Module) error {
	iters := []MetaIterator{
		NewMetaListIterator(src.GetGroupings(), false),
		NewMetaListIterator(src.GetTypedefs(), false),
		NewMetaListIterator(src.DataDefs(), false),
	}
	for _, iter := range iters {
		for iter.HasNextMeta() {
			if m, err := iter.NextMeta(); err != nil {
				return err
			} else {
				dest.AddMeta(m)
			}
		}
	}
	return nil
}

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

func IsAction(m Meta) bool {
	_, isAction := m.(*Rpc)
	return isAction
}

func IsNotification(m Meta) bool {
	_, isNotification := m.(*Notification)
	return isNotification
}

func IsLeaf(m Meta) bool {
	switch m.(type) {
	case *Leaf, *LeafList, *Any:
		return true
	}
	return false
}

func IsKeyLeaf(parent MetaList, leaf Meta) bool {
	if !IsList(parent) || !IsLeaf(leaf) {
		return false
	}
	for _, keyIdent := range parent.(*List).Key {
		if keyIdent == leaf.GetIdent() {
			return true
		}
	}
	return false
}

func ListEmpty(parent MetaList) (empty bool) {
	i := NewMetaListIterator(parent, true)
	return !i.HasNextMeta()
}

func ListLen(parent MetaList) (len int) {
	i := NewMetaListIterator(parent, true)
	for i.HasNextMeta() {
		len++
		i.NextMeta()
	}
	return
}

func ListLenNoExpand(parent MetaList) (len int) {
	i := NewMetaListIterator(parent, false)
	for i.HasNextMeta() {
		len++
		i.NextMeta()
	}
	return
}

func IsList(m Meta) bool {
	_, isList := m.(*List)
	return isList
}

func IsContainer(m Meta) bool {
	return !IsList(m) && !IsLeaf(m)
}

func MetaNameToFieldName(in string) string {
	// assumes fix is always shorter because char can be dropped and not added
	fixed := make([]rune, len(in))
	cap := true
	j := 0
	for _, r := range in {
		if r == '-' {
			cap = true
		} else {
			if cap {
				fixed[j] = unicode.ToUpper(r)
			} else {
				fixed[j] = r
			}
			j += 1
			cap = false
		}
	}
	return string(fixed[:j])
}

func ListToArray(l MetaList) ([]Meta, error) {
	// PERFORMANCE: is it better to iterate twice, pass 1 to find length?
	meta := make([]Meta, 0)
	i := NewMetaListIterator(l, true)
	for i.HasNextMeta() {
		m, err := i.NextMeta()
		if err != nil {
			return nil, err
		}
		meta = append(meta, m)
	}
	return meta, nil
}

func GetPath(m Meta) string {
	s := m.GetIdent()
	if p := m.GetParent(); p != nil {
		return GetPath(p) + "/" + s
	}
	return s
}

func GetModule(m Meta) *Module {
	candidate := m
	for candidate.GetParent() != nil {
		candidate = candidate.GetParent()
	}
	return candidate.(*Module)
}

func FindByPathWithoutResolvingProxies(root Meta, path string) (Meta, error) {
	return find(root, path, false)
}

func FindByPath(root Meta, path string) (Meta, error) {
	return find(root, path, true)
}

func find(root Meta, path string, resolveProxies bool) (def Meta, err error) {
	if strings.HasPrefix(path, "../") {
		return find(root.GetParent(), path[3:], resolveProxies)
	} else if strings.HasPrefix(path, "/") {
		p := root
		for p.GetParent() != nil {
			p = p.GetParent()
		}
		return find(p, path[1:], resolveProxies)
	}
	elems := strings.SplitN(path, "/", -1)
	lastLevel := len(elems) - 1
	var ok bool
	list := root
	i := NewMetaListIterator(list, resolveProxies)
	for level, elem := range elems {
		def, err = FindByIdent(i, elem)
		if def == nil || err != nil {
			return nil, err
		}
		if level < lastLevel {
			if list, ok = def.(MetaList); ok {
				i = NewMetaListIterator(list, resolveProxies)
			} else {
				return nil, nil
			}
		}
	}
	return
}
