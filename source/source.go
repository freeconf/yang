package source

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"path"
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

// EmbedDir will find a yang file in an embedded file system
// because Embed doesn't support symlinks, EmbedDir also looks for version specific files of yang
// with a pattern of <resourceId>@<version><ext>, e.g. ietf-inet-types@2013-07-15.yang
func EmbedDir(fs embed.FS, root string) Opener {
	return func(resourceId string, ext string) (io.Reader, error) {
		name := fmt.Sprint(resourceId, ext)
		bin, err := fs.ReadFile(path.Join(root, name))
		if err == nil {
			return bytes.NewReader(bin), nil
		}

		entries, err := fs.ReadDir(root)
		if err != nil {
			return nil, nil
		}

		for _, e := range entries {
			if strings.HasPrefix(e.Name(), resourceId+"@") && strings.HasSuffix(e.Name(), ext) {
				bin, err = fs.ReadFile(path.Join(root, e.Name()))
				if err != nil {
					return nil, err
				}
				return bytes.NewReader(bin), nil
			}
		}
		return nil, nil
	}
}
