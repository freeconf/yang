package node

import (
	"meta"
	"bytes"
	"strings"
)

type PathMatcher interface {
	PathMatches(*Path) bool
	FieldMatches(*Path, meta.HasDataType) bool
}

type pathMatchEntry struct {
	segments []string
}

type PathMatchExpression struct {
	root *Path
	slices []*segSlice
}

type seg struct {
	parent *seg
	ident string
}

type segSlice struct {
	head *seg
	tail *seg
}

func (self *segSlice) Len() int {
	len := 1
	s := self.tail
	for s != self.head {
		s = s.parent
		len++
	}
	return len
}

func (self *segSlice) copy() *segSlice {
	orig := self.tail
	var copy segSlice
	for orig != nil {
		n := &seg {
			ident: orig.ident,
		}
		if copy.head == nil {
			copy.head = n
			copy.tail = n
		} else {
			copy.head.parent = n
			copy.head = n
		}
		orig = orig.parent
	}
	return &copy
}

func ParsePathExpression(root *Path, selector string) (*PathMatchExpression, error) {
	pe := &PathMatchExpression{root:root}
	pe.parsex(&lex{selector:selector})
	return pe, nil
}

type lex struct {
	pos int
	selector string
}

func (self *lex) next() (s string) {
	tokenlen := strings.IndexAny(self.selector[self.pos:], "(;)/")
	if tokenlen < 0 {
		s = self.selector[self.pos:]
		self.pos = len(self.selector)
	} else if (tokenlen == 0) {
		s = self.selector[self.pos:self.pos + 1]
		self.pos++
	} else {
		end := self.pos + tokenlen
		s = self.selector[self.pos:end]
		self.pos = end
	}
	return s
}

func (self *lex) done() bool {
	return self.pos >= len(self.selector)
}

func (self *PathMatchExpression) parsex(l *lex) {
	s := self
	var split *PathMatchExpression
	for !l.done() {
		t := l.next()
		switch t {
		case "(":
			nested := &PathMatchExpression{}
			nested.parsex(l)
			s.addSubExpression(nested)
		case ";":
			if split != nil {
				self.addNextExpression(s)
			}
			split = &PathMatchExpression{}
			s = split
		case ")":
			if split != nil {
				self.addNextExpression(s)
			}
			return
		case "/":
			// ignore natural delimitor
		default:
			s.addSegment(t)
		}
	}
	if split != nil {
		self.addNextExpression(s)
	}
}

func (self *PathMatchExpression) addSubExpression(sub *PathMatchExpression) {
	expanded := make([]*segSlice, len(self.slices) * len(sub.slices))
	for i, slice := range self.slices {
		for j, subSlicesOrig := range sub.slices {
			subSlices := subSlicesOrig.copy()
			n := &segSlice{
				head: slice.head,
				tail: subSlices.tail,
			}
			subSlices.head.parent = slice.tail
			k := (i * len(sub.slices)) + j
			expanded[k] = n
		}
	}
	self.slices = expanded
}

func (self *PathMatchExpression) addNextExpression(next *PathMatchExpression) {
	self.slices = append(self.slices, next.slices...)
}

func (self *PathMatchExpression) addSegment(ident string) {
	if len(self.slices) == 0 {
		seg := &seg{
			ident: ident,
		}
		self.slices = []*segSlice{
			&segSlice{
				head: seg,
				tail: seg,
			},
		}
	} else {
		for _, slice := range self.slices {
			seg := &seg{
				ident: ident,
				parent: slice.tail,
			}
			slice.tail = seg
		}
	}
}

func (self *PathMatchExpression) copy() {
	for _, slice := range self.slices {
		seg := seg{
			parent: slice.tail,
		}
		slice.tail = &seg
	}
}

func (self *PathMatchExpression) PathMatches(candidate *Path) bool {
	for _, slice := range self.slices {
		if self.sliceMatches(slice, candidate) {
			return true
		}
	}
	return false
}

func (self *PathMatchExpression) sliceMatches(slice *segSlice, candidate *Path) bool {
	aRootLen := self.root.Len()
	aSegLen := slice.Len()
	aLen := aRootLen + aSegLen
	bLen := candidate.Len()
	if bLen < aLen {
		return false
	}

	// PERF: We start at end because in practice we eliminate candidates faster
	// we may also want to consider caching results
	p := candidate
	aSegTail := slice.tail
	aRootTail := self.root
	for i := bLen; i > 0; i-- {
		if i <= aLen {
			if i > aRootLen {
				if aSegTail.ident != p.goober.GetIdent() {
					return false
				}
				aSegTail = aSegTail.parent
			} else {
				return aRootTail.Equal(p)
			}
		}
		p = p.parent
	}
	return true
}

func (self *PathMatchExpression) FieldMatches(candidate *Path, goober meta.HasDataType) bool {
	c2 := &Path{
		parent: candidate,
		goober: goober,
	}
	return self.PathMatches(c2)
}

func (self *PathMatchExpression) String() string {
	var buff bytes.Buffer
	for i, slice := range self.slices {
		if i != 0 {
			buff.WriteRune(',')
		}
		buff.WriteRune('[')
		self.writeSeg(&buff, slice.tail)
		buff.WriteRune(']')
	}
	return buff.String()
}

func (self *PathMatchExpression) writeSeg(buff *bytes.Buffer, seg *seg) {
	if seg.parent != nil {
		self.writeSeg(buff, seg.parent)
		buff.WriteRune(',')
	}
	buff.WriteString(seg.ident)
}


