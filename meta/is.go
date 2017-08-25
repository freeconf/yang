package meta

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

func IsList(m Meta) bool {
	_, isList := m.(*List)
	return isList
}

// IsContainer Module, Container
func IsContainer(m Meta) bool {
	return !IsList(m) && !IsLeaf(m)
}

// IsKeyLeaf tests if this leaf makes up the or one of the keys in a list
// meta
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
