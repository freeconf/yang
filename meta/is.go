package meta

// IsList returns true if meta is *Leaf or *LeafList
func IsLeaf(m Meta) bool {
	switch m.(type) {
	case *Leaf, *LeafList, *Any:
		return true
	}
	return false
}

// IsNotification returns true if meta is *Notification
func IsNotification(m Meta) bool {
	_, match := m.(*Notification)
	return match
}

// IsList returns true if meta is *List
func IsList(m Meta) bool {
	_, match := m.(*List)
	return match
}

// IsContainer return true if meta is *Module or *Container
func IsContainer(m Meta) bool {
	switch m.(type) {
	case *Container, *Module:
		return true
	}
	return false
}

// IsAction returns true is meta is *Rpc (YANG rpm or action)
func IsAction(m Meta) bool {
	_, match := m.(*Rpc)
	return match
}

func IsChoice(m Meta) bool {
	_, match := m.(*Choice)
	return match
}

func IsChoiceCase(m Meta) bool {
	_, match := m.(*ChoiceCase)
	return match
}

// IsDataDef is *Container, *List or Leaf
func IsDataDef(m Meta) bool {
	return IsList(m) || IsContainer(m) || IsLeaf(m)
}
