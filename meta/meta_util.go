package meta

import (
	"strings"
	"unicode"
)

func FindByIdent(i MetaIterator, ident string) Meta {
	child := i.NextMeta()
	for child != nil {
		if child.GetIdent() == ident {
			return child
		}
		child = i.NextMeta()
	}
	return nil
}

func FindByIdent2(parent Meta, ident string) Meta {
	i := NewMetaListIterator(parent, true)
	return FindByIdent(i, ident)
}

func FindByIdentExpandChoices(m Meta, ident string) Meta {
	parent, isMetaList := m.(MetaList)
	if !isMetaList {
		return nil
	}
	i := NewMetaListIterator(parent, true)
	var choice *Choice
	var isChoice bool
	for i.HasNextMeta() {
		child := i.NextMeta()
		choice, isChoice = child.(*Choice)
		if isChoice {
			cases := NewMetaListIterator(choice, false)
			for cases.HasNextMeta() {
				ccase := cases.NextMeta().(*ChoiceCase)
				found := FindByIdentExpandChoices(ccase, ident)
				if found != nil {
					return found
				}
			}
		} else {
			if child.GetIdent() == ident {
				return child
			}
		}
		//child = i.NextMeta()
	}
	return nil
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

func ListToArray(l MetaList) []Meta {
	// PERFORMANCE: is it better to iterate twice, pass 1 to find length?
	meta := make([]Meta, 0)
	i := NewMetaListIterator(l, true)
	for i.HasNextMeta() {
		m := i.NextMeta()
		meta = append(meta, m)
	}
	return meta
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

func FindByPathWithoutResolvingProxies(root Meta, path string) Meta {
	c := find(root, path, false)
	return c
}

func FindByPath(root Meta, path string) Meta {
	return find(root, path, true)
}

func find(root Meta, path string, resolveProxies bool) (def Meta) {
	if strings.HasPrefix(path, "../") {
		return find(root.GetParent(), path[3:], resolveProxies)
	}
	elems := strings.SplitN(path, "/", -1)
	lastLevel := len(elems) - 1
	var ok bool
	list := root
	i := NewMetaListIterator(list, resolveProxies)
	for level, elem := range elems {
		def = FindByIdent(i, elem)
		if def == nil {
			return nil
		}
		if level < lastLevel {
			if list, ok = def.(MetaList); ok {
				i = NewMetaListIterator(list, resolveProxies)
			} else {
				return nil
			}
		}
	}
	return
}
