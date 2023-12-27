package xpath

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/freeconf/yang/meta"
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

func Parse2(lookup ShortcodeToModule, pstr string) (*Path, error) {
	l := lex(pstr)
	l.lookup = lookup
	if err := yyParse(l); err != 0 {
		return nil, l.lastError
	}
	p := l.stack.pop()
	for p.Parent != nil {
		p = p.Parent
	}
	return p, nil
}

type ShortcodeToModule func(shortcode string) (*meta.Module, error)

func Parse(pstr string) (*Path, error) {
	nolookup := func(string) (*meta.Module, error) {
		return nil, fmt.Errorf("lookup function for XML ns shortcodes not specified")
	}
	return Parse2(nolookup, pstr)
}

type Expression interface {
	String() string
}

type Operator struct {
	Oper string
	Lhs  interface{}
}

func (o *Operator) String() string {
	s := o.Oper
	if str, isStr := o.Lhs.(string); isStr {
		s = s + "'" + str + "'"
	} else {
		s = fmt.Sprintf("%s%v", o.Oper, o.Lhs)
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

type Path struct {
	Parent *Path
	Ident  string
	Expr   Expression
	Next   *Path
}

func (p *Path) String() string {
	s := p.Ident
	if p.Next != nil {
		s = fmt.Sprintf("%s/%s", s, p.Next.String())
	}
	if p.Expr != nil {
		s = s + p.Expr.String()
	}
	return s
}
