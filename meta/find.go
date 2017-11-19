package meta

import "strings"
import "github.com/c2stack/c2g/c2"

// TODO: Support namespaces
func Find(p HasDefinitions, path string) Definition {
	c2.Debug.Printf("looking for %s on %s", path, p.Ident())
	if strings.HasPrefix(path, "../") {
		return Find(p.Parent().(HasDefinitions), path[3:])
	}
	if strings.HasPrefix(path, "/") {
		return Find(Root(p), path[1:])
	}
	if seg := strings.IndexRune(path, '/'); seg > 0 {
		c2.Debug.Printf("seg=%s, tail=%s", path[:seg], path[seg+1:])
		if child := p.Definition(path[:seg]); child != nil {
			return Find(child.(HasDefinitions), path[seg+1:])
		}
		return nil
	}
	return p.Definition(path)
}
