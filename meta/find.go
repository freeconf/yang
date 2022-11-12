package meta

import "strings"

func Find(p Meta, path string) Definition {
	if strings.HasPrefix(path, "../") {
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
		prefix := path[:colon]
		mod, err := RootModule(p).ModuleByPrefix(prefix)
		if err != nil {
			return nil
		}
		return Find(mod, path[colon+1:])
	}
	if hd, ok := p.(HasDataDefinitions); ok {
		return hd.Definition(path)
	}
	panic(SchemaPath(p) + " does not have definitions")
}
