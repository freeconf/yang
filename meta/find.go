package meta

import "strings"

func Find(p Meta, path string) Definition {
	if strings.HasPrefix(path, "../") {
		if p.Parent() == nil {
			return nil
		}
		return Find(p.Parent(), path[3:])
	}
	if strings.HasPrefix(path, "/") {
		return Find(RootModule(p), path[1:])
	}
	if seg := strings.IndexRune(path, '/'); seg > 0 {
		if child := Find(p, path[:seg]); child != nil {
			return Find(child, path[seg+1:])
		}
		return nil
	}
	if colon := strings.IndexRune(path, ':'); colon > 0 {
		ident := path[colon+1:]
		if _, isModule := p.(*Module); isModule {
			prefix := path[:colon]
			mod, err := RootModule(p).ModuleByPrefix(prefix)
			if err != nil {
				return nil
			}
			p = mod
		}
		return Find(p, ident)
	}
	if hd, ok := p.(HasDataDefinitions); ok {
		return hd.Definition(path)
	} else {
		if _, ok := p.(HasCases); ok {
			if choice, ok := p.(*Choice); ok {
				if c, found := choice.Cases()[path]; found {
					return c
				}
				return nil
			}
		}
	}
	return nil
}
