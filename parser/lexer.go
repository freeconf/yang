package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/freeconf/yang/meta"
)

// This uses the go feature call go tools in the build process. To ensure this gets
//  called before compilation, make this call before building
//
//    go generate
//
//  To build the goyacc binary, run
//
//    go get golang.org/x/tools/cmd/goyacc
//
//go:generate goyacc -o parser.go parser.y

//
// When making changes to this code, be aware that order of statements
// can be important because tokens can mean multiple things in different
// contexts

type token struct {
	typ int
	val string
}

type stateFunc func(*lexer) stateFunc
type tokenSink func(*token)

const (
	parseEof = iota
	parseErr
)

const (
	char_doublequote         = '"'
	char_singlequote         = '\''
	char_backslash           = '\\'
	str_comment_start        = "/*"
	str_comment_end          = "*/"
	str_comment_inline_start = "//"
)

// needs to be in-sync w/ %token list in parser.y
var keywords = [...]string{
	"[ident]",
	"[string]",
	"[number]",
	"[extension]",
	"{",
	"}",
	";",

	// KEEP LIST IN SYNC WITH parser.y
	"namespace",
	"description",
	"revision",
	"type",
	"prefix",
	"default",
	"length",
	"enum",
	"key",
	"config",
	"uses",
	"unique",
	"input",
	"output",
	"module",
	"container",
	"list",
	"rpc",
	"notification",
	"typedef",
	"grouping",
	"leaf",
	"mandatory",
	"reference",
	"leaf-list",
	"max-elements",
	"min-elements",
	"choice",
	"case",
	"import",
	"include",
	"action",
	"anyxml",
	"anydata",
	"path",
	"value",
	"true",
	"false",
	"contact",
	"organization",
	"refine",
	"unbounded",
	"augment",
	"submodule",
	"+",
	"identity",
	"base",
	"feature",
	"if-feature",
	"when",
	"must",
	"yang-version",
	"range",
	"extension",
	"argument",
	"yin-element",
	"pattern",
	"units",
	"fraction-digits",
	"status",
	"current",
	"obsolete",
	"deprecated",
	"presence",
	"deviation",
	"deviate",
	"not-supported",
	"add",
	"replace",
	"delete",
	"ordered-by",
	"system",
	"user",
	"require-instance",
	"error-app-tag",
	"error-message",
	"bit",
	"position",
}

const eof rune = 0

func (l *lexer) keyword(ttype int) string {
	if ttype < token_ident {
		panic("Not a keyword")
	}
	return keywords[ttype-token_ident]
}

func (t token) String() string {
	if t.typ == parseErr {
		return fmt.Sprintf("ERROR: %q", t.val)
	}
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

func (l *lexer) error(msg string) stateFunc {
	l.tokens = append(l.tokens, token{
		parseErr,
		msg,
	})
	l.Error(msg)
	return nil
}

func (l *lexer) importModule(into *meta.Module, moduleName string) error {
	return nil
}

type yangMetaStack struct {
	defs  []interface{}
	count int
}

func (s *yangMetaStack) push(def interface{}) interface{} {
	s.defs[s.count] = def
	s.count++
	return def
}

func (s *yangMetaStack) pop() interface{} {
	s.count--
	return s.defs[s.count]
}

func (s *yangMetaStack) peek() interface{} {
	return s.defs[s.count-1]
}

func (s *yangMetaStack) peekModule() *meta.Module {
	for i := len(s.defs) - 1; i >= 0; i-- {
		if m, valid := s.defs[i].(*meta.Module); valid {
			return m
		}
	}
	return nil
}

func newDefStack(size int) *yangMetaStack {
	return &yangMetaStack{
		defs:  make([]interface{}, size),
		count: 0,
	}
}

type lexer struct {
	pos        int
	start      int
	width      int
	state      stateFunc
	input      string
	tokens     []token
	head       int
	tail       int
	stack      *yangMetaStack
	loader     meta.Loader
	featureSet meta.FeatureSet
	parent     *meta.Module
	lastError  error
	builder    *meta.Builder
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) Position() (line, col int) {
	for p := 0; p < l.pos; p++ {
		if l.input[p] == '\n' {
			line += 1
			col = 0
		} else {
			col += 1
		}
	}
	return
}

func (l *lexer) isEof() bool {
	return l.pos >= len(l.input)
}

func (l *lexer) acceptWS() {
	for {
		for unicode.IsSpace(l.next()) {
		}
		l.backup()

		if strings.HasPrefix(l.input[l.pos:], str_comment_start) {
			for {
				l.next()
				if strings.HasPrefix(l.input[l.pos:], str_comment_end) {
					l.pos += len(str_comment_end)
					break
				}
			}
		} else if strings.HasPrefix(l.input[l.pos:], str_comment_inline_start) {
			for {
				l.next()
				if l.input[l.pos] == '\n' {
					l.pos++
					break
				}
			}
		} else {
			break
		}
	}
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

func (l *lexer) acceptToken(ttype int) bool {
	var keyword string
	switch ttype {
	case token_extension:
		return l.acceptToks(ttype, isIdent, isPrefixedIdent)
	case token_ident:
		return l.acceptToks(ttype, isIdent, nil)
	case token_string:
		return l.acceptString()
	case token_number:
		return l.acceptNumber(token_number)
	case token_curly_open:
		keyword = "{"
		break
	case token_curly_close:
		keyword = "}"
		break
	case token_semi:
		keyword = ";"
		break
	default:
		keyword = l.keyword(ttype)
	}
	if !strings.HasPrefix(l.input[l.pos:], keyword) {
		return false
	}
	l.pos += len(keyword)
	l.emit(ttype)
	return true
}

func (l *lexer) acceptRun(ttype int, valid string) bool {
	found := false
	for strings.IndexRune(valid, l.next()) >= 0 {
		found = true
	}
	l.backup()
	if found {
		l.emit(ttype)
	}
	return found
}

func (l *lexer) acceptString() bool {
	begin := l.next()
	isDblQuote := begin == char_doublequote
	isSglQuote := begin == char_singlequote
	isSpaceDelim := !isSglQuote && !isDblQuote
	if !isSpaceDelim && !isDblQuote && !isSglQuote {
		l.backup()
		return false
	}
	for {
		term := false
		r := l.next()
		if isSpaceDelim {
			if r == eof {
				term = true
			} else if isStringDelim(r) {
				term = true
				l.backup()
			}
		} else {
			if r == eof {
				// bad format
				return false
			}
			if r == char_backslash && isDblQuote {
				l.next()
				continue
			}
			if r == begin && !isSpaceDelim {
				term = true
			}
		}
		if term {
			l.emit(token_string)
			if !isSpaceDelim {
				if l.acceptToken(kywd_str_plus) {
					return l.acceptString()
				}
			}
			return true
		}
	}
}

// strings that are not surrounded by quotes (single or double) are allowed
func isStringDelim(r rune) bool {
	return unicode.IsSpace(r) || r == ';' || r == '{'
}

func (l *lexer) acceptNumber(ttype int) bool {
	accepted := false
	for i := 0; ; i++ {
		r := l.next()
		sign := ((r == '-' || r == '+') && i == 0)
		decimal := (r == '.' && i != 0)
		if !unicode.IsDigit(r) && !sign && !decimal {
			l.backup()
			if accepted {
				l.emit(ttype)
			}
			return accepted
		}
		accepted = true
	}
}

func (l *lexer) acceptInteger(ttype int) bool {
	accepted := false
	for i := 0; ; i++ {
		r := l.next()
		sign := ((r == '-' || r == '+') && i == 0)
		if !unicode.IsDigit(r) && !sign {
			l.backup()
			if accepted {
				l.emit(ttype)
			}
			return accepted
		}
		accepted = true
	}
}

func isIdent(r rune) bool {
	return isAlphaNumeric(r) || r == ':' || r == '.'
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r) || r == '-' || r == '_'
}

func isPrefixedIdent(s string) bool {
	count := 0
	for _, r := range s {
		if r == ':' {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return count == 1
}

type runeTest func(r rune) bool
type strTest func(s string) bool

func (l *lexer) acceptToks(ttype int, rfunc runeTest, sfunc strTest) bool {
	accepted := false
	for {
		r := l.next()
		// TODO: review spec on legal chars
		if !rfunc(r) {
			s := l.input[l.start:l.pos]
			l.backup()
			if sfunc != nil {
				if !sfunc(s) {
					return false
				}
			}
			if accepted {
				l.emit(ttype)
			}
			return accepted
		}
		accepted = true
	}
}

func lexBegin(l *lexer) stateFunc {
	if l.isEof() {
		return nil
	}

	// FORMAT: xxx zzz { ...
	// order from longest keyword to shorted to ensure "foobar" doesn't get picked
	// up as "foo"
	types := []int{
		kywd_notification,
		kywd_container,
		kywd_leaf_list,
		kywd_submodule,
		kywd_grouping,
		kywd_typedef,
		kywd_action,
		kywd_module,
		kywd_choice,
		kywd_leaf,
		kywd_list,
		kywd_case,
		kywd_rpc,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptToken(token_ident) {
				return l.error("expected ident")
			}
			if !l.acceptToken(token_curly_open) {
				return l.error("expected {")
			}
			return lexBegin
		}
	}

	// FORMAT : xxx "path" { ...
	types = []int{
		kywd_deviation,
		kywd_augment,
		kywd_refine,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptString() {
				return l.error("expected ident or string of a path")
			}
			if !l.acceptToken(token_curly_open) {
				return l.error("expected {")
			}
			return lexBegin
		}
	}

	// FORMAT : aaa { ...
	types = []int{
		kywd_input,
		kywd_output,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptToken(token_curly_open) {
				return l.error("expected {")
			}
			return lexBegin
		}
	}

	// FORMAT:
	//  deviate (fixed_set) ;
	if l.acceptToken(kywd_deviate) {
		if l.acceptToken(kywd_not_supported) {
			if l.acceptToken(token_curly_open) {
				return lexBegin
			}
			return l.acceptEndOfStatement()
		}
		deviateTypes := []int{
			kywd_replace,
			kywd_add,
			kywd_delete,
		}
		for _, ttype := range deviateTypes {
			if l.acceptToken(ttype) {
				if !l.acceptToken(token_curly_open) {
					return l.error("expected {")
				}
				return lexBegin
			}
		}
		return l.error("expected deviate type")
	}

	// FORMAT: Either
	//  xxx zzz;
	// or
	//  xxx zzz { ...
	types = []int{
		kywd_not_supported,
		kywd_extension,
		kywd_identity,
		kywd_include,
		kywd_anydata,
		kywd_feature,
		kywd_anyxml,
		kywd_import,
		kywd_type,
		kywd_bit,
		kywd_uses,
		kywd_base,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptToken(token_ident) {
				return l.error("expecting string")
			}
			return l.acceptEndOfStatement()
		}
	}

	if l.acceptToken(kywd_status) {
		allowed := []int{
			kywd_current,
			kywd_obsolete,
			kywd_deprecated,
		}
		for _, t := range allowed {
			if l.acceptToken(t) {
				return l.acceptEndOfStatement()
			}
		}
		return l.error("unexpected token after status")
	}

	// FORMAT:
	// xxx (number || string);
	types = []int{
		kywd_default,
		kywd_value,
		kywd_position,
		kywd_fraction_digits,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptNumber(token_number) && !l.acceptString() {
				return l.error("expecting number or string")
			}
			return l.acceptEndOfStatement()
		}
	}

	// FORMAT: xxx [true|false]
	types = []int{
		kywd_mandatory,
		kywd_config,
		kywd_yin_element,
		kywd_require_instance,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptToken(kywd_true) && !l.acceptToken(kywd_false) {
				return l.error("expecting true or false")
			}
			return l.acceptEndOfStatement()
		}
	}

	// FORMAT: xxx "zzz";
	// longest first just ensures most specific will match over
	// least specific
	types = []int{
		kywd_error_message,
		kywd_error_app_tag,
		kywd_organization,
		kywd_yang_version,
		kywd_description,
		kywd_if_feature,
		kywd_namespace,
		kywd_reference,
		kywd_revision,
		kywd_argument,
		kywd_presence,
		kywd_contact,
		kywd_pattern,
		kywd_prefix,
		kywd_length,
		kywd_unique,
		kywd_range,
		kywd_units,
		kywd_enum,
		kywd_type,
		kywd_path,
		kywd_when,
		kywd_must,
		kywd_key,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptString() {
				return l.error("expecting string")
			}
			return l.acceptEndOfStatement()
		}
	}

	if l.acceptToken(kywd_ordered_by) {
		types = []int{
			kywd_system,
			kywd_user,
		}
		for _, ttype := range types {
			if l.acceptToken(ttype) {
				return l.acceptEndOfStatement()
			}
		}
		return l.error("unexpected order-by type")
	}

	// FORMAT: xxx number;
	types = []int{
		kywd_max_elements,
		kywd_min_elements,
	}
	for _, ttype := range types {
		if l.acceptToken(ttype) {
			if !l.acceptToken(kywd_unbounded) {
				if !l.acceptInteger(token_number) {
					return l.error("expecting integer")
				}
			}
			return l.acceptEndOfStatement()
		}
	}

	if l.acceptToken(token_curly_close) {
		return lexBegin
	}

	if l.acceptToken(token_extension) {
		for {
			if l.acceptToken(token_semi) {
				return lexBegin
			}
			if l.acceptToken(token_curly_open) {
				return lexBegin
			}
			if l.acceptToken(token_number) {
				continue
			}
			if l.acceptToken(token_string) {
				continue
			}
			return l.error("expecting token")
		}
	}

	return l.error("unknown statement")
}

func (l *lexer) acceptEndOfStatement() stateFunc {
	if !l.acceptToken(token_semi) && !l.acceptToken(token_curly_open) {
		return l.error("expecting semicolon or '{'")
	}
	return lexBegin
}

func (l *lexer) emit(t int) {
	l.pushToken(token{t, l.input[l.start:l.pos]})
	l.start = l.pos
	l.acceptWS()
}

func (l *lexer) popToken() token {
	token := l.tokens[l.tail]
	l.tail = (l.tail + 1) % len(l.tokens)
	return token
}

func (l *lexer) pushToken(t token) {
	l.tokens[l.head] = t
	l.head = (l.head + 1) % len(l.tokens)
}

func (l *lexer) nextToken() (token, error) {
	for {
		if l.head != l.tail {
			token := l.popToken()
			if token.typ == parseEof {
				return token, errors.New(token.val)
			}
			return token, nil
		}
		if l.state == nil {
			return token{parseEof, "EOF"}, nil
		}
		l.state = l.state(l)
	}
}

const (
	lexRingBufferSize = 64
	nestedYangDefMax  = 256
)

func lex(input string, loader meta.Loader) *lexer {
	l := &lexer{
		input:  input,
		tokens: make([]token, lexRingBufferSize),
		head:   0,
		tail:   0,
		state:  lexBegin,
		stack:  newDefStack(256),
		loader: loader,
	}
	l.acceptWS()
	return l
}

// useful only in test cases
func lexDump(y string, w io.Writer) error {
	l := lex(string(y), nil)
	for {
		token, err := l.nextToken()
		if err != nil {
			return err
		} else if l.lastError != nil {
			return l.lastError
		} else if token.typ == parseEof {
			return nil
		}
		l := fmt.Sprintf("%s %s\n", l.keyword(token.typ), token.String())
		if _, err := w.Write([]byte(l)); err != nil {
			return err
		}
	}
}
