package node

import (
	"bytes"
	"strings"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

// Immutable otherwise children paths become illegal if parent state changes
type Path struct {
	Meta   meta.Definition
	Key    []val.Value
	Parent *Path
}

// func NewRootPath(m meta.Definition) *Path {
// 	return &Path{Meta: m}
// }

// func NewListItemPath(parent *Path, m *meta.List, key []val.Value) *Path {
// 	return &Path{Parent: parent, Meta: m, Key: key}
// }

// func (path *Path) SetKey(key []val.Value) *Path {
// 	return &Path{Parent: path.Parent, Meta: path.Meta, Key: key}
// }

// func NewContainerPath(parent *Path, m meta.HasDefinitions) *Path {
// 	return &Path{Parent: parent, Meta: m}
// }

// func (path *Path) Parent() *Path {
// 	return path.parent
// }

// func (path *Path) MetaParent() meta.Path {
// 	if path.Parent == nil {
// 		// subtle difference returning nil and interface reference to nil struct.
// 		// See http://stackoverflow.com/questions/13476349/check-for-nil-and-nil-interface-in-go
// 		// by rights in go, all callers should check for interface check for nil and nil interface
// 		// so this hack some-what contributes to the bad practice of not doing so.
// 		return nil
// 	}
// 	return path.Parent.Meta
// }

// func (path *Path) Meta() meta.Definition {
// 	return path.meta
// }

// func (path *Path) Key() []val.Value {
// 	return path.key
// }

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
