package xpath

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// This uses the go feature call go tools in the build process. To ensure this gets
//  called before compilation, make this call before building
//	  go get golang.org/x/tools/cmd/goyacc
//    cd src
//    go generate github.com/freeconf/c2g/xpath
//
//go:generate goyacc -o parser.go parser.y

type Token struct {
	typ int
	val string
}

const eof rune = 0

const (
	ParseEnd = 0
	ParseErr
)

type stateFunc func(*lexer) stateFunc

type lexer struct {
	pos       int
	start     int
	width     int
	state     stateFunc
	input     string
	tokens    []Token
	stack     *stack
	head      int
	tail      int
	lastError error
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		l.pos = len(l.input)
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) isEnd() bool {
	return l.pos >= len(l.input)
}

func (l *lexer) acceptWS() {
	for unicode.IsSpace(l.next()) {
	}
	l.backup()
	l.ignore()
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) error(msg string) stateFunc {
	l.tokens = append(l.tokens, Token{
		ParseErr,
		msg,
	})
	l.Error(msg)
	return nil
}

func (l *lexer) acceptLiteral(ttype int) bool {
	first := true
	for {
		switch l.next() {
		case eof:
			return false
		case '\'':
			if !first {
				l.emit(ttype)
				return true
			}
		default:
			if first {
				l.backup()
				return false
			}
		}
		first = false
	}
}

func (l *lexer) acceptNumeric(ttype int) bool {
	first := true
	for {
		r := l.next()
		if unicode.IsDigit(r) || (!first && r == '.') {
			first = false
		} else {
			if r != eof {
				l.backup()
			}
			if first {
				return false
			}
			l.emit(ttype)
			return true
		}
	}
}

func (l *lexer) acceptAlphaNumeric(ttype int) bool {
	accepted := false
	for {
		r := l.next()
		// TODO: review spec on legal chars
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && !(r == '-') && !(r == '_') && !(r == '.') {
			l.backup()
			if accepted {
				l.emit(ttype)
			}
			return accepted
		}
		accepted = true
	}
}

func (l *lexer) emit(t int) {
	l.pushToken(Token{t, l.input[l.start:l.pos]})
	l.start = l.pos
	l.acceptWS()
}

func (l *lexer) acceptOperator() bool {
	switch l.next() {
	case '=':
		l.emit(token_operator)
	case '!':
		if l.next() == '=' {
			l.emit(token_operator)
		} else {
			l.backup()
			return false
		}
	case '<':
		if l.next() == '=' {
			l.emit(token_operator)
		} else {
			l.backup()
			l.emit(token_operator)
		}
	case '>':
		if l.next() == '=' {
			l.emit(token_operator)
		} else {
			l.backup()
			l.emit(token_operator)
		}
	default:
		l.backup()
		return false
	}
	return true
}

func lexBegin(l *lexer) stateFunc {
	if l.isEnd() {
		return nil
	}

	if l.acceptToken(kywd_slash) {
		return lexBegin
	}

	if l.acceptOperator() {
		if l.acceptNumeric(token_number) {
			return lexBegin
		}
		if l.acceptLiteral(token_literal) {
			return lexBegin
		}
	}

	if l.acceptToken(token_name) {
		return lexBegin
	}

	if l.acceptToken(token_number) {
		return lexBegin
	}
	return l.error("unknown statement")
}

func (l *lexer) popToken() Token {
	token := l.tokens[l.tail]
	l.tail = (l.tail + 1) % len(l.tokens)
	return token
}

func (l *lexer) pushToken(t Token) {
	l.tokens[l.head] = t
	l.head = (l.head + 1) % len(l.tokens)
}

func (l *lexer) acceptToken(ttype int) bool {
	var keyword string
	switch ttype {
	case token_name:
		return l.acceptAlphaNumeric(ttype)
	case token_number:
		return l.acceptNumeric(ttype)
	case kywd_slash:
		keyword = "/"
	}
	if !strings.HasPrefix(l.input[l.pos:], keyword) {
		return false
	}
	l.pos += len(keyword)
	l.emit(ttype)
	return true
}

func (l *lexer) nextToken() (Token, error) {
	for {
		if l.head != l.tail {
			token := l.popToken()
			return token, nil
		} else {
			if l.state == nil {
				return Token{typ: 0}, nil
			}
			l.state = l.state(l)
		}
	}
}

type stack struct {
	steps []Path
	count int
}

func (self *stack) push(p Path) {
	self.steps[self.count] = p
	self.count++
}

func (self *stack) pop() Path {
	self.count--
	return self.steps[self.count]
}

func (self *stack) peek() Path {
	return self.steps[self.count-1]
}

const (
	lexRingBufferSize = 64
	nestedPathStack   = 256
)

func lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make([]Token, lexRingBufferSize),
		head:   0,
		tail:   0,
		stack: &stack{
			steps: make([]Path, nestedPathStack),
		},
		state: lexBegin,
	}
	l.acceptWS()
	return l
}
