package meta

import "strings"

func FindInIterator(i Iterator, ident string) (Meta, error) {
	child, err := i.Next()
	if err != nil {
		return nil, err
	}
	for child != nil {
		if child.GetIdent() == ident {
			return child, nil
		}
		child, err = i.Next()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func Find(parent MetaList, ident string) (Meta, error) {
	i := Children(parent, true)
	return FindInIterator(i, ident)
}

func FindByIdentExpandChoices(m MetaList, ident string) (Meta, error) {
	i := Children(m, true)
	var choice *Choice
	var isChoice bool
	for i.HasNext() {
		child, err := i.Next()
		if err != nil {
			return nil, err
		}
		choice, isChoice = child.(*Choice)
		if isChoice {
			cases := Children(choice, false)
			for cases.HasNext() {
				ccase, err := cases.Next()
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

func FindByPathWithoutResolvingProxies(root MetaList, path string) (Meta, error) {
	return find(root, path, false)
}

func FindByPath(root MetaList, path string) (Meta, error) {
	return find(root, path, true)
}

func find(root MetaList, path string, resolveProxies bool) (def Meta, err error) {
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
	i := Children(list, resolveProxies)
	for level, elem := range elems {
		def, err = FindInIterator(i, elem)
		if def == nil || err != nil {
			return nil, err
		}
		if level < lastLevel {
			if list, ok = def.(MetaList); ok {
				i = Children(list, resolveProxies)
			} else {
				return nil, nil
			}
		}
	}
	return
}
