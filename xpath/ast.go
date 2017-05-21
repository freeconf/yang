package xpath

import (
	"fmt"
	"strings"

	"strconv"
)

// examples
//  /event/event-class='fault'
//  /event/severity<=4
//  /linkUp|/linkDown
//  /*/reporting-entity/card!='Ethernet0'
//  /*/email-addr[contains(.,'company.com')]
//  (/example-mod:event1/name='joe' and
//           /example-mod:event1/status='online')
//  (/m1:* or /m2:*)
//
//  /moduleName='car'

func Parse(pstr string) (Path, error) {
	l := lex(pstr)
	if err := yyParse(l); err != 0 {
		return nil, l.lastError
	}
	return l.stack.pop(), nil
}

type Expression interface {
	String() string
}

type Operator struct {
	Oper string
	Lhs  interface{}
}

func (self *Operator) String() string {
	s := self.Oper
	if str, isStr := self.Lhs.(string); isStr {
		s = s + "'" + str + "'"
	} else {
		s = fmt.Sprintf("%s%v", self.Oper, self.Lhs)
	}
	return s
}

func num(s string) (interface{}, error) {
	if strings.ContainsRune(s, '.') {
		return strconv.ParseFloat(s, 64)
	}
	return strconv.ParseInt(s, 10, 64)
}

func literal(s string) interface{} {
	cutset := "'"
	return strings.TrimRight(strings.TrimLeft(s, cutset), cutset)
}

type Path interface {
	SetParent(Path)
	Parent() Path
	String() string
	Append(Path)
	Next() Path
}

type Segment struct {
	parent Path
	next   Path
	Ident  string
	Expr   Expression
}

func (self *Segment) String() string {
	s := self.Ident
	if self.next != nil {
		s = fmt.Sprintf("%s/%s", s, self.next.String())
	}
	if self.Expr != nil {
		s = s + self.Expr.String()
	}
	return s
}

func (self *Segment) Parent() Path {
	return self.parent
}

func (self *Segment) SetParent(parent Path) {
	self.parent = parent
}

func (self *Segment) Next() Path {
	return self.next
}

func (self *Segment) Append(next Path) {
	next.SetParent(self)
	self.next = next
}

type AbsolutePath struct {
	next Path
}

func (self *AbsolutePath) Parent() Path {
	return nil
}

func (self *AbsolutePath) SetParent(parent Path) {
	panic("Cannot set parent of absolute path")
}

func (self *AbsolutePath) Next() Path {
	return self.next
}

func (self *AbsolutePath) String() string {
	return "/" + self.next.String()
}

func (self *AbsolutePath) Append(p Path) {
	self.next = p
}
