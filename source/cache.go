package source

import (
	"fmt"
	"io"
	"os"
)

// Backup keeps a copy of every reasource that
func Cached(src Opener, backup Cacher) Opener {
	return func(resourceId string, ext string) (io.Reader, error) {
		s, err := backup.ReadFromCache(resourceId, ext)
		if s == nil {
			found, err := src(resourceId, ext)
			if err != nil || found == nil {
				return nil, err
			}
			if err := backup.WriteToCache(resourceId, ext, found); err != nil {
				return nil, err
			}
			// now we should be able to get reader from backup
			s, err = backup.ReadFromCache(resourceId, ext)
		}
		return s, err
	}
}

// Cacher implementations manage cached resources
type Cacher interface {
	WriteToCache(resourceId string, ext string, copy io.Reader) error
	ReadFromCache(resourceId string, ext string) (io.Reader, error)
}

// DirCache will cache yang files into local file directory
func DirCache(root string) Cacher {
	return dirCache{
		root:   root,
		opener: Dir(root),
	}
}

type dirCache struct {
	root   string
	opener Opener
}

func (d dirCache) ReadFromCache(resourceId string, ext string) (io.Reader, error) {
	return d.opener(resourceId, ext)
}

func (d dirCache) WriteToCache(resourceId string, ext string, copy io.Reader) error {
	path := fmt.Sprint(d.root, "/", resourceId, ext)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, copy)
	return err
}
