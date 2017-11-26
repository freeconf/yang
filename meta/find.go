package meta

import "strings"

// TODO: Support namespaces
func Find(p HasDefinitions, path string) Definition {
	if strings.HasPrefix(path, "../") {
		return Find(p.Parent().(HasDefinitions), path[3:])
	}
	if strings.HasPrefix(path, "/") {
		return Find(Root(p), path[1:])
	}
	if seg := strings.IndexRune(path, '/'); seg > 0 {
		if child := p.Definition(path[:seg]); child != nil {
			return Find(child.(HasDefinitions), path[seg+1:])
		}
		return nil
	}
	return p.Definition(path)
}
