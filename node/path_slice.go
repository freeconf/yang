package node

import (
	"bytes"
	"net/url"
	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
)

type PathSlice struct {
	Head *Path
	Tail *Path
}

func NewPathSlice(path string, m meta.HasDefinitions) (p PathSlice) {
	var err error
	if p, err = ParsePath(path, m); err != nil {
		if err != nil {
			panic(err.Error())
		}
	}
	return p
}

func ParsePath(path string, m meta.HasDefinitions) (PathSlice, error) {
	u, err := url.Parse(path)
	if err != nil {
		return PathSlice{}, err
	}
	return ParseUrlPath(u, m)
}

func ParseUrlPath(u *url.URL, m meta.Definition) (PathSlice, error) {
	var err error
	p := NewRootPath(m)
	slice := PathSlice{
		Head: p,
		Tail: p,
	}
	segments := strings.Split(u.EscapedPath(), "/")
	for _, segment := range segments {

		// a/b/c same as a/b/c/
		if segment == "" {
			break
		}

		var ident string
		var keyStrs []string

		// next path segment
		seg := &Path{parent: p}
		equalsMark := strings.Index(segment, "=")

		// has key
		if equalsMark >= 0 {
			if ident, err = url.QueryUnescape(segment[:equalsMark]); err != nil {
				return PathSlice{}, err
			}
			keyStrs = strings.Split(segment[equalsMark+1:], ",")
			for i, escapedKeystr := range keyStrs {
				if keyStrs[i], err = url.QueryUnescape(escapedKeystr); err != nil {
					return PathSlice{}, err
				}
			}
			// no key
		} else {
			if ident, err = url.QueryUnescape(segment); err != nil {
				return PathSlice{}, err
			}
		}

		// find meta associated with path ident
		seg.meta = meta.Find(p.meta.(meta.HasDefinitions), ident)
		if seg.meta == nil {
			return PathSlice{}, c2.NewErrC(ident+" not found in "+p.meta.Ident(), 404)
		}

		if len(keyStrs) > 0 {
			if seg.key, err = NewValuesByString(seg.meta.(*meta.List).KeyMeta(), keyStrs...); err != nil {
				return PathSlice{}, err
			}
		}

		// append to tail
		seg.parent = slice.Tail
		slice.Tail = seg

		p = seg
	}
	return slice, nil
}

func (self PathSlice) Empty() bool {
	return self.Tail == self.Head
}

func (self PathSlice) Len() (len int) {
	p := self.Tail
	for p != self.Head {
		len++
		p = p.parent
		if p == nil {
			panic("bad path slice.  head was never found in tail's parent list")
		}
	}
	return
}

func (self PathSlice) String() string {
	var b bytes.Buffer
	for _, segment := range self.Segments() {
		segment.toBuffer(&b)
	}
	return b.String()
}

func (self PathSlice) Segments() []*Path {
	segments := make([]*Path, self.Len())
	p := self.Tail
	for i := len(segments) - 1; i >= 0; i-- {
		segments[i] = p
		p = p.parent
	}
	return segments
}
