package node

import (
	"bytes"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Path in data tree (not meta) resembling a RESTCONF URL when printed to a string
type Path struct {
	Meta   meta.Definition
	Key    []val.Value
	Parent *Path
}

func (seg *Path) StringNoModule() string {
	return seg.str(false)
}

func (seg *Path) String() string {
	return seg.str(true)
}

func (seg *Path) str(showModule bool) string {
	l := seg.Len()
	if !showModule {
		l--
	}
	strs := make([]string, l)
	p := seg
	var b bytes.Buffer
	for i := l - 1; i >= 0; i-- {
		b.Reset()
		p.toBuffer(&b)
		strs[i] = b.String()
		p = p.Parent
	}
	return strings.Join(strs, "/")
}

func (seg *Path) toBuffer(b *bytes.Buffer) {
	if seg.Meta == nil {
		return
	}
	if b.Len() > 0 {
		b.WriteRune('/')
	}
	b.WriteString(seg.Meta.Ident())
	if len(seg.Key) > 0 {
		b.WriteRune('=')
		for i, k := range seg.Key {
			if i != 0 {
				b.WriteRune(',')
			}
			if k != nil {
				b.WriteString(k.String())
			} else {
				b.WriteString("<nil>")
			}
		}
	}
}

func (a *Path) EqualNoKey(b *Path) bool {
	if a.Len() != b.Len() {
		return false
	}
	sa := a
	sb := b
	// work up as comparing children are most likely to lead to differences faster
	for sa != nil {
		if !sa.equalSegment(sb, false) {
			return false
		}
		sa = sa.Parent
		sb = sb.Parent
	}
	return true
}

func (a *Path) Equal(b *Path) bool {
	if a.Len() != b.Len() {
		return false
	}
	sa := a
	sb := b
	// work up as comparing children are most likely to lead to differences faster
	for sa != nil {
		if !sa.equalSegment(sb, true) {
			return false
		}
		sa = sa.Parent
		sb = sb.Parent
	}
	return true
}

func (path *Path) Len() (len int) {
	p := path
	for p != nil {
		len++
		p = p.Parent
	}
	return
}

func (a *Path) equalSegment(b *Path, compareKey bool) bool {
	if a.Meta == nil {
		if b.Meta != nil {
			return false
		}
		if a.Meta.Ident() != b.Meta.Ident() {
			return false
		}
	}
	if compareKey {
		if len(a.Key) != len(b.Key) {
			return false
		}
		for i, k := range a.Key {
			if !val.Equal(k, b.Key[i]) {
				return false
			}
		}
	}
	return true
}

func (path *Path) Segments() []*Path {
	segs := make([]*Path, path.Len())
	p := path
	for i := len(segs) - 1; i >= 0; i-- {
		segs[i] = p
		p = p.Parent
	}
	return segs
}
