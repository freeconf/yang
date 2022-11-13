package node

import (
	"bytes"
	"strings"
)

type PathMatcher interface {
	PathMatches(base *Path, tail *Path) bool
}

type PathMatchExpression struct {
	slices []*segSlice
}

type seg struct {
	parent *seg
	ident  string
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
		n := &seg{
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

func ParsePathExpression(selector string) (*PathMatchExpression, error) {
	pe := &PathMatchExpression{}
	pe.parsex(&lex{selector: selector})
	return pe, nil
}

type lex struct {
	pos      int
	selector string
}

func (self *lex) next() (s string) {
	tokenlen := strings.IndexAny(self.selector[self.pos:], "(;)/")
	if tokenlen < 0 {
		s = self.selector[self.pos:]
		self.pos = len(self.selector)
	} else if tokenlen == 0 {
		s = self.selector[self.pos : self.pos+1]
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
			// ignore natural delimitor already used in lexer
		default:
			s.addSegment(t)
		}
	}
	if split != nil {
		self.addNextExpression(s)
	}
}

func (self *PathMatchExpression) addSubExpression(sub *PathMatchExpression) {
	expanded := make([]*segSlice, len(self.slices)*len(sub.slices))
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
				ident:  ident,
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

func (self *PathMatchExpression) PathMatches(base *Path, candidate *Path) bool {
	// NOTE: empty selector means select everything
	if len(self.slices) == 0 {
		return true
	}
	for _, slice := range self.slices {

		// NOTE: empty selector means select everything
		if slice.tail == nil {
			return true
		}

		if self.sliceMatches(slice, base, candidate) {
			return true
		}
	}
	return false
}

// Find out if a path selector is root of candidate when you subtract the base.
// While this might be easier to do if we just converted everything to strings, we
// do not do that because it would be less efficient then walking the three arguments
// and we need this to be fast as selectors can be excersize for every node.
//
// Example Match:
//  base      : some/path
//  candidate : some/path=key/more/path/here/and/here
//  slice     :               more/path
//
// Example Mismatch (even though this starts to match first few times thru loop):
//  base      : some/path
//  candidate : some/path=key/more/path/here/and/here
//  slice     :               and/here
//
func (self *PathMatchExpression) sliceMatches(slice *segSlice, base *Path, candidate *Path) bool {
	s := slice.tail
	p := candidate
	subpathIndex := candidate.Len() - base.Len()
	filterPathIndex := slice.Len()

	// start navigation at the end of the tail
	for {
		if filterPathIndex == 0 {
			// the subpath AFTER base path matches, now we have to see if we have same
			// base paths
			if p.EqualNoKey(base) {
				return true
			} else if p == nil {
				panic("illegal call : base was not found to be any parent of candidate")
			}
		}
		// we keep peeling back slice as long as it continues to match candidate as we
		// peel that back as well.
		if subpathIndex == filterPathIndex {
			if p.Meta.Ident() != s.ident {
				return false
			}
			s = s.parent
			filterPathIndex--
		}
		p = p.Parent
		subpathIndex--
	}
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
