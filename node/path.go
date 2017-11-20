package node

import (
	"bytes"
	"strings"

	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/val"
)

// Immutable otherwise children paths become illegal if parent state changes
type Path struct {
	meta   meta.Definition
	key    []val.Value
	parent *Path
}

func NewRootPath(m meta.Definition) *Path {
	return &Path{meta: m}
}

func NewListItemPath(parent *Path, m *meta.List, key []val.Value) *Path {
	return &Path{parent: parent, meta: m, key: key}
}

func (path *Path) SetKey(key []val.Value) *Path {
	return &Path{parent: path.parent, meta: path.meta, key: key}
}

func NewContainerPath(parent *Path, m meta.HasDefinitions) *Path {
	return &Path{parent: parent, meta: m}
}

func (path *Path) Parent() *Path {
	return path.parent
}

func (path *Path) MetaParent() meta.Path {
	if path.parent == nil {
		// subtle difference returning nil and interface reference to nil struct.
		// See http://stackoverflow.com/questions/13476349/check-for-nil-and-nil-interface-in-go
		// by rights in go, all callers should check for interface check for nil and nil interface
		// so this hack some-what contributes to the bad practice of not doing so.
		return nil
	}
	return path.parent
}

func (path *Path) Meta() meta.Definition {
	return path.meta
}

func (path *Path) Key() []val.Value {
	return path.key
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
		p = p.parent
	}
	return strings.Join(strs, "/")
}

func (seg *Path) toBuffer(b *bytes.Buffer) {
	if seg.meta == nil {
		return
	}
	if b.Len() > 0 {
		b.WriteRune('/')
	}
	b.WriteString(seg.meta.Ident())
	if len(seg.key) > 0 {
		b.WriteRune('=')
		for i, k := range seg.key {
			if i != 0 {
				b.WriteRune(',')
			}
			b.WriteString(k.String())
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
		sa = sa.parent
		sb = sb.parent
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
		sa = sa.parent
		sb = sb.parent
	}
	return true
}

func (path *Path) Len() (len int) {
	p := path
	for p != nil {
		len++
		p = p.parent
	}
	return
}

func (a *Path) equalSegment(b *Path, compareKey bool) bool {
	if a.meta == nil {
		if b.meta != nil {
			return false
		}
		if a.meta.Ident() != b.meta.Ident() {
			return false
		}
	}
	if compareKey {
		if len(a.key) != len(b.key) {
			return false
		}
		for i, k := range a.key {
			if !val.Equal(k, b.key[i]) {
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
		p = p.parent
	}
	return segs
}
