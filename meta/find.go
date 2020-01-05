package meta

import "strings"

// TODO: Support namespaces
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
		// TODO: qualify by namespace
		// prefix := path[:colon]
		path = path[colon+1:]
	}
	if hd, ok := p.(HasDataDefinitions); ok {
		return hd.Definition(path)
	}
	panic(SchemaPath(p) + " does not have definitions")
}
