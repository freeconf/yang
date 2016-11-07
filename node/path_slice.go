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

func NewPathSlice(path string, m meta.MetaList) (p PathSlice) {
	var err error
	if p, err = ParsePath(path, m); err != nil {
		if err != nil {
			panic(err.Error())
		}
	}
	return p
}

func ParsePath(path string, m meta.MetaList) (PathSlice, error) {
	u, err := url.Parse(path)
	if err != nil {
		return PathSlice{}, err
	}
	return ParseUrlPath(u, m)
}

func ParseUrlPath(u *url.URL, m meta.Meta) (PathSlice, error) {
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
		seg.meta = meta.FindByIdentExpandChoices(p.meta, ident)
		if seg.meta == nil {
			return PathSlice{}, c2.NewErrC(ident+" not found in "+p.meta.GetIdent(), 404)
		}

		// now we know meta, convert keys to proper data type
		if len(keyStrs) > 0 {
			if seg.key, err = CoerseKeys(seg.meta.(*meta.List), keyStrs); err != nil {
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

// func (a PathSlice) Equal(b PathSlice) bool {
// 	if self.Len() != bPath.Len() {
// 		return false
// 	}
// 	a := self.Tail
// 	b := bPath.Tail
// 	for a != nil {
// 		if a.meta != b.meta {
// 			return false
// 		}
// 		if len(a.key) != len(b.key) {
// 			return false
// 		}
// 		for i, k := range a.key {
// 			if !k.Equal(b.key[i]) {
// 				return false
// 			}
// 		}
// 		a = a.parent
// 		b = b.parent
// 	}
// 	return true
// }

// func (self PathSlice) PopHead() PathSlice {
// 	if self.Head == self.Tail {
// 		return self
// 	}
// 	// only singly-linked list so we have to start at tail
// 	candidate := self.Tail
// 	for candidate != nil {
// 		if candidate.parent.Equal(self.Head) {
// 			self.Head = candidate
// 			return self
// 		}
// 		candidate = candidate.parent
// 	}
// 	panic("slice has discontinuous parts")
// }

func (self PathSlice) Empty() bool {
	return self.Tail == self.Head
}

// func (self PathSlice) SplitAfter(point *Path) PathSlice {
// 	self.Head = point
// 	return self
// }

// func (self PathSlice) Append(child *Path) PathSlice {
// 	child.parent = self.Tail
// 	self.Tail = child
// 	return self
// }

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
