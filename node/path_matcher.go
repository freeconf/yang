package node

import (
	"bytes"
	"strings"
)

type PathMatcher interface {
	PathMatches(base *Path, tail *Path) bool
}

type PathMatchExpression struct {
	paths []segments
}

// a single, denormalized list of idents after parsing expression
//
// simple expression:
//
//	/aaa/bbb/ccc
//
// would be
//
//	['aaa', 'bbb', 'ccc']
type segments []string

func ParsePathExpression(selector string) (*PathMatchExpression, error) {
	pe := &PathMatchExpression{}
	pe.parsex(&lex{selector: selector})
	return pe, nil
}

type lex struct {
	pos      int
	selector string
}

func (l *lex) next() (s string) {
	tokenlen := strings.IndexAny(l.selector[l.pos:], "(;)/")
	if tokenlen < 0 {
		s = l.selector[l.pos:]
		l.pos = len(l.selector)
	} else if tokenlen == 0 {
		s = l.selector[l.pos : l.pos+1]
		l.pos++
	} else {
		end := l.pos + tokenlen
		s = l.selector[l.pos:end]
		l.pos = end
	}
	return s
}

func (l *lex) done() bool {
	return l.pos >= len(l.selector)
}

func (e *PathMatchExpression) parsex(l *lex) {
	s := e
	var split *PathMatchExpression
	for !l.done() {
		t := l.next()
		switch t {
		case "(":
			nested := &PathMatchExpression{}
			nested.parsex(l)
			s.expandPaths(nested)
		case ";":
			if split != nil {
				e.appendPaths(s)
			}
			split = &PathMatchExpression{}
			s = split
		case ")":
			if split != nil {
				e.appendPaths(s)
			}
			return
		case "/":
			// ignore natural delimitor already used in lexer
		default:
			s.addSegment(t)
		}
	}
	if split != nil {
		e.appendPaths(s)
	}
}

// expandPaths incoming paths into current paths
// Example:
//
//	If this expression has paths
//	   [a,b]
//	   [c,d]
//	and you add following paths
//	   [e, f]
//	   [g, h]
//	this expressions new paths becomes
//	   [a, b, e, f]
//	   [a, b, g, h]
//	   [c, d, e, f]
//	   [c, d, g, h]
func (e *PathMatchExpression) expandPaths(sub *PathMatchExpression) {
	expanded := make([]segments, len(e.paths)*len(sub.paths))
	for i, dest := range e.paths {
		for j, src := range sub.paths {
			k := (i * len(sub.paths)) + j
			expanded[k] = append(dest, src...)
		}
	}
	e.paths = expanded
}

// appendPaths to incoming paths into current paths
// Example:
//
//	If this expression has paths
//	   [a,b]
//	   [c,d]
//	and you add following paths
//	   [e, f]
//	   [g, h]
//	this expressions new paths becomes
//	   [a, b]
//	   [c, d]
//	   [e, f]
//	   [g, h]
func (e *PathMatchExpression) appendPaths(next *PathMatchExpression) {
	e.paths = append(e.paths, next.paths...)
}

// addSegment adds to incoming paths a new segement
// Example:
//
//	If this expression has paths
//	   [a,b]
//	   [c,d]
//	and you add 'g'
//	this expressions new paths becomes
//	   [a, b, g]
//	   [c, d, g]
func (e *PathMatchExpression) addSegment(ident string) {
	if len(e.paths) == 0 {
		e.paths = []segments{[]string{ident}}
	} else {
		for i, path := range e.paths {
			e.paths[i] = append(path, ident)
		}
	}
}

// PathMatches returns true if a path selector is root of candidate when you subtract the base.
// While this might be easier to do if we just converted everything to strings, we
// do not do that because it would be less efficient then walking the three arguments
// and we need this to be fast as selectors can be excersize for every node.
//
// Example Match:
//
//	base      : some/path
//	candidate : some/path=key/more/path/here/and/here
//	slice     :               more/path
//
// Example Mismatch (even though this starts to match first few times thru loop):
//
//	base      : some/path
//	candidate : some/path=key/more/path/here/and/here
//	slice     :               and/here
func (e *PathMatchExpression) PathMatches(base *Path, candidate *Path) bool {
	// NOTE: empty selector means select everything
	if len(e.paths) == 0 {
		return true
	}
	for _, path := range e.paths {

		// NOTE: empty selector means select everything
		if len(path) == 0 {
			return true
		}

		if e.match(path, base, candidate) {
			return true
		}
	}
	return false
}

func (e *PathMatchExpression) match(segs segments, base *Path, candidate *Path) bool {
	p := candidate
	j := (candidate.Len() - base.Len()) - 1

	// start navigation at the end of the tail as it would likely be more efficient the longer
	// the path
	for i := len(segs) - 1; i >= 0; {

		// we keep peeling back slice as long as it continues to match candidate as we
		// peel that back as well.
		if j == i {
			if p.Meta.Ident() != segs[i] {
				return false
			}
			i--
		}
		p = p.Parent
		if p == nil {
			panic("illegal call : base was not found to be any parent of candidate")
		}
		j--
	}

	// the subpath AFTER base path matches, now we have to see if we have same
	// base paths
	return p.EqualNoKey(base)
}

func (e *PathMatchExpression) String() string {
	var buff bytes.Buffer
	for i, slice := range e.paths {
		if i != 0 {
			buff.WriteRune(',')
		}
		buff.WriteRune('[')
		for j, ident := range slice {
			if j != 0 {
				buff.WriteRune(',')
			}
			buff.WriteString(ident)
		}
		buff.WriteRune(']')
	}
	return buff.String()
}
