package source

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Opener will find yang files (resources) abstracting away where or how
type Opener func(src string, ext string) (io.Reader, error)

// Named - useful when you have a single yang file by name
func Named(name string, r io.Reader) Opener {
	return func(src string, extCandidate string) (io.Reader, error) {
		if src == name {
			return r, nil
		}
		return nil, nil
	}
}

// Path looks in a series of directories separated by ':'.  Example ./foo:./bar
func Path(path string) Opener {
	dirs := strings.Split(path, ":")
	sources := make([]Opener, len(dirs))
	for i, dir := range dirs {
		sources[i] = Dir(dir)
		i++
	}
	return Any(sources...)
}

// Any will sequentially look in any of these openers for the particular resource
func Any(s ...Opener) Opener {
	return func(resourceId string, ext string) (io.Reader, error) {
		for _, source := range s {
			found, err := source(resourceId, ext)
			if found != nil {
				return found, err
			}
		}
		return nil, os.ErrNotExist
	}
}

// Dir will simply find yang files in a directory
func Dir(root string) Opener {
	return func(resourceId string, ext string) (io.Reader, error) {
		path := fmt.Sprint(root, "/", resourceId, ext)
		stream, err := os.Open(path)
		if os.IsNotExist(err) {
			return nil, nil
		}
		return stream, err
	}
}
