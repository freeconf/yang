//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/c2g/meta"
	"strings"
)

type yangError struct {
	s string
}

func (err *yangError) Error() string {
	return err.s
}

func tokenString(s string) string {
	return s[1 : len(s)-1]
}

func (l *lexer) Lex(lval *yySymType) int {
	t, _ := l.nextToken()
	if t.typ == ParseEof {
		return 0
	}
	lval.token = t.val
	lval.stack = l.stack
	lval.importer = l.importer
	return int(t.typ)
}

func (l *lexer) Error(e string) {
	line, col := l.Position()
	msg := fmt.Sprintf("%s - line %d, col %d", e, line, col)
	l.lastError = &yangError{msg}
}

func HasError(l yyLexer, e error) bool {
	if e == nil {
		return false
	}
	l.Error(e.Error())
	return true
}

func popAndAddMeta(yylval *yySymType) error {
	i := yylval.stack.Pop()
	if def, ok := i.(meta.Meta); ok {
		parent := yylval.stack.Peek()
		if parentList, ok := parent.(meta.MetaList); ok {
			return parentList.AddMeta(def)
		} else {
			return &yangError{fmt.Sprintf("Cannot add \"%s\" to \"%s\"; not collection type.", i.GetIdent(), parent.GetIdent())}
		}
	} else {
		return &yangError{fmt.Sprintf("\"%s\" cannot be stored in a collection type.", i.GetIdent())}
	}
}

//line parser.y:63
type yySymType struct {
	yys      int
	ident    string
	token    string
	dataType *meta.DataType
	stack    *yangMetaStack
	importer ImportModule
}

const token_ident = 57346
const token_string = 57347
const token_int = 57348
const token_curly_open = 57349
const token_curly_close = 57350
const token_semi = 57351
const token_rev_ident = 57352
const kywd_namespace = 57353
const kywd_description = 57354
const kywd_revision = 57355
const kywd_type = 57356
const kywd_prefix = 57357
const kywd_default = 57358
const kywd_length = 57359
const kywd_enum = 57360
const kywd_key = 57361
const kywd_config = 57362
const kywd_uses = 57363
const kywd_unique = 57364
const kywd_input = 57365
const kywd_output = 57366
const kywd_module = 57367
const kywd_container = 57368
const kywd_list = 57369
const kywd_rpc = 57370
const kywd_notification = 57371
const kywd_typedef = 57372
const kywd_grouping = 57373
const kywd_leaf = 57374
const kywd_mandatory = 57375
const kywd_reference = 57376
const kywd_leaf_list = 57377
const kywd_max_elements = 57378
const kywd_choice = 57379
const kywd_case = 57380
const kywd_import = 57381
const kywd_include = 57382
const kywd_action = 57383
const kywd_anyxml = 57384
const kywd_path = 57385

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_int",
	"token_curly_open",
	"token_curly_close",
	"token_semi",
	"token_rev_ident",
	"kywd_namespace",
	"kywd_description",
	"kywd_revision",
	"kywd_type",
	"kywd_prefix",
	"kywd_default",
	"kywd_length",
	"kywd_enum",
	"kywd_key",
	"kywd_config",
	"kywd_uses",
	"kywd_unique",
	"kywd_input",
	"kywd_output",
	"kywd_module",
	"kywd_container",
	"kywd_list",
	"kywd_rpc",
	"kywd_notification",
	"kywd_typedef",
	"kywd_grouping",
	"kywd_leaf",
	"kywd_mandatory",
	"kywd_reference",
	"kywd_leaf_list",
	"kywd_max_elements",
	"kywd_choice",
	"kywd_case",
	"kywd_import",
	"kywd_include",
	"kywd_action",
	"kywd_anyxml",
	"kywd_path",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:618

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 141
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 391

var yyAct = [...]int{

	136, 229, 135, 140, 159, 170, 96, 154, 143, 148,
	125, 119, 156, 103, 10, 137, 111, 27, 7, 94,
	7, 3, 141, 230, 10, 10, 144, 27, 24, 64,
	139, 133, 108, 49, 131, 246, 245, 190, 55, 54,
	156, 10, 52, 53, 56, 109, 138, 57, 127, 59,
	49, 244, 66, 60, 58, 55, 54, 234, 233, 52,
	53, 56, 232, 99, 57, 187, 59, 186, 226, 230,
	60, 58, 107, 63, 231, 62, 224, 122, 132, 91,
	121, 114, 223, 95, 104, 222, 213, 113, 161, 120,
	126, 207, 145, 145, 228, 204, 152, 160, 112, 106,
	149, 172, 172, 201, 173, 129, 107, 147, 147, 95,
	6, 10, 16, 168, 9, 105, 178, 195, 104, 192,
	114, 128, 122, 146, 146, 121, 113, 183, 132, 184,
	191, 179, 169, 106, 120, 194, 200, 112, 11, 166,
	126, 150, 6, 10, 182, 238, 9, 78, 10, 105,
	117, 203, 116, 61, 17, 129, 196, 237, 145, 236,
	10, 208, 117, 212, 116, 161, 172, 172, 214, 215,
	11, 128, 243, 147, 160, 220, 10, 242, 221, 219,
	218, 10, 211, 206, 108, 49, 205, 202, 199, 146,
	55, 54, 164, 165, 52, 53, 56, 109, 177, 57,
	193, 59, 10, 99, 167, 60, 58, 217, 216, 209,
	108, 49, 235, 176, 175, 81, 55, 54, 80, 77,
	52, 53, 56, 109, 76, 57, 10, 59, 75, 240,
	74, 60, 58, 73, 108, 49, 200, 241, 239, 65,
	55, 54, 10, 71, 52, 53, 56, 109, 68, 57,
	67, 59, 49, 100, 101, 60, 58, 55, 54, 40,
	41, 52, 53, 56, 99, 10, 57, 22, 59, 198,
	197, 188, 60, 58, 49, 47, 181, 180, 174, 55,
	54, 20, 19, 52, 53, 56, 18, 99, 57, 210,
	59, 49, 189, 90, 60, 58, 55, 54, 40, 41,
	52, 53, 56, 5, 49, 57, 89, 59, 14, 55,
	54, 60, 58, 52, 53, 56, 88, 87, 57, 10,
	59, 117, 86, 116, 60, 58, 85, 108, 84, 83,
	82, 79, 70, 69, 21, 12, 142, 46, 48, 130,
	109, 124, 123, 44, 118, 72, 43, 102, 29, 163,
	162, 158, 157, 51, 98, 97, 93, 92, 28, 134,
	45, 227, 225, 185, 115, 110, 42, 155, 153, 151,
	50, 39, 38, 37, 36, 35, 34, 33, 32, 31,
	30, 171, 26, 25, 8, 15, 23, 13, 4, 2,
	1,
}
var yyPact = [...]int{

	-4, -1000, 131, 331, 99, 145, 281, -1000, -1000, 277,
	276, 330, 260, 270, 144, 66, 19, -1000, -1000, -1000,
	-1000, -1000, -1000, 231, -1000, -1000, -1000, -1000, 243, 241,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	329, 328, 236, 226, 223, 221, 217, 212, 138, 327,
	211, 208, 326, 325, 324, 322, 318, 313, 312, 302,
	289, -1000, -1000, 13, -1000, -1000, -1000, 230, 214, -1000,
	-1000, 148, -1000, 253, 12, 214, 307, 307, -1000, 132,
	2, 169, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 130, 196, 230, -1000, 123, -1000, 283, 283, 273,
	207, 206, 190, -1000, 122, -1000, -1000, -1000, 272, 271,
	136, -1000, -1000, 120, -1000, 58, 266, 288, 29, -1000,
	110, -1000, -1000, 192, 12, -1000, 108, 150, -1000, -1000,
	-1000, 265, -1000, 264, 180, 214, -1000, 94, -1000, -1000,
	-1000, 179, 307, -1000, -1000, 86, -1000, -1000, -1000, 178,
	-1000, 175, 82, -26, -1000, 202, 285, 174, 169, -1000,
	77, -1000, 283, 283, 201, 200, 172, -1000, -1000, -1000,
	171, 283, -1000, 170, 76, -1000, -1000, -1000, -1000, -1000,
	73, 67, -1000, -1000, -1000, -1000, -1000, 51, 65, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 53, 49, 48, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 214,
	-1000, -1000, -1000, -1000, 151, 149, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 137, 233, 5, 232, -1000,
	173, -1000, -1000, -1000, -1000, 164, -1000, -1000, -1000, 42,
	-1000, 27, 26, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 390, 389, 388, 387, 386, 385, 15, 303, 384,
	28, 383, 382, 3, 5, 381, 380, 379, 378, 377,
	376, 375, 374, 373, 372, 371, 370, 369, 368, 7,
	367, 2, 366, 365, 16, 26, 9, 364, 363, 362,
	361, 360, 359, 0, 46, 30, 358, 357, 356, 19,
	6, 355, 354, 353, 352, 351, 4, 350, 349, 348,
	347, 13, 346, 345, 344, 11, 343, 342, 341, 10,
	339, 338, 337, 22, 336, 8, 275, 1,
}
var yyR1 = [...]int{

	0, 1, 2, 6, 4, 4, 7, 3, 3, 8,
	8, 8, 8, 9, 10, 10, 10, 5, 5, 14,
	14, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 15, 15, 24, 27, 27, 27, 26, 28, 28,
	29, 30, 16, 32, 33, 33, 34, 34, 34, 36,
	35, 37, 38, 38, 39, 39, 39, 19, 41, 42,
	42, 31, 31, 43, 43, 43, 43, 23, 11, 46,
	47, 47, 48, 48, 49, 49, 49, 49, 51, 52,
	25, 53, 54, 54, 55, 55, 56, 56, 56, 56,
	57, 58, 12, 59, 60, 60, 61, 61, 61, 61,
	17, 63, 62, 64, 64, 65, 65, 65, 18, 66,
	67, 68, 68, 69, 69, 69, 69, 69, 69, 69,
	70, 22, 71, 20, 72, 73, 74, 74, 75, 75,
	75, 75, 75, 45, 44, 21, 76, 40, 40, 77,
	50,
}
var yyR2 = [...]int{

	0, 5, 3, 2, 2, 5, 2, 2, 3, 2,
	1, 1, 2, 2, 1, 1, 1, 1, 2, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 4, 0, 2, 1, 2, 1, 2,
	4, 2, 4, 2, 1, 2, 1, 2, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 2, 1, 1, 1, 3, 4, 2,
	0, 1, 1, 2, 2, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 2, 1, 3, 3,
	2, 2, 4, 2, 1, 2, 2, 1, 1, 1,
	2, 3, 2, 1, 2, 2, 1, 1, 4, 2,
	1, 1, 2, 2, 3, 1, 1, 1, 3, 1,
	3, 2, 2, 4, 2, 1, 1, 2, 1, 2,
	1, 1, 1, 3, 3, 4, 2, 1, 2, 3,
	3,
}
var yyChk = [...]int{

	-1000, -1, -2, 25, -3, -8, 11, -7, -9, 15,
	12, 39, 4, -4, -8, -6, 13, 9, 5, 5,
	5, 4, 7, -5, -10, -11, -12, -13, -46, -59,
	-16, -17, -18, -19, -20, -21, -22, -23, -24, -25,
	28, 29, -32, -62, -66, -41, -72, -76, -71, 21,
	-26, -53, 30, 31, 27, 26, 32, 35, 42, 37,
	41, 9, 9, 7, 10, 8, -10, 7, 7, 4,
	4, 7, -63, 7, 7, 7, 7, 7, 9, 4,
	7, 7, 4, 4, 4, 4, 4, 4, 4, 4,
	4, -7, -47, -48, -49, -7, -50, -51, -52, 34,
	23, 24, -60, -61, -7, -44, -45, -13, 20, 33,
	-33, -34, -35, -7, -36, -37, 16, 14, -64, -65,
	-7, -50, -13, -67, -68, -69, -7, 36, -44, -45,
	-70, 22, -13, 19, -42, -31, -43, -7, -44, -45,
	-13, -73, -74, -75, -35, -7, -44, -45, -36, -73,
	9, -27, -7, -28, -29, -30, 38, -54, -55, -56,
	-7, -50, -57, -58, 23, 24, 9, 8, -49, 9,
	-14, -15, -13, -14, 5, 7, 7, 8, -61, 9,
	5, 5, 8, -34, 9, -38, 9, 7, 5, 4,
	8, -65, 9, 8, -69, 9, 6, 5, 5, 8,
	-43, 9, 8, -75, 9, 8, 8, 9, -29, 7,
	4, 8, -56, 9, -14, -14, 7, 7, 8, 8,
	-13, 8, 9, 9, 9, -39, 17, -40, 43, -77,
	18, 9, 9, 9, 9, -31, 8, 8, 8, 5,
	-77, 5, 4, 8, 9, 9, 9,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 10, 11, 0,
	0, 0, 0, 0, 0, 0, 0, 7, 9, 12,
	6, 13, 2, 0, 17, 14, 15, 16, 0, 0,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 4, 0, 3, 1, 18, 70, 0, 69,
	93, 0, 100, 0, 0, 59, 0, 0, 121, 0,
	34, 82, 43, 102, 109, 58, 124, 136, 122, 37,
	81, 0, 0, 71, 72, 0, 75, 19, 19, 0,
	0, 0, 0, 94, 0, 97, 98, 99, 0, 0,
	0, 44, 46, 0, 48, 0, 0, 0, 0, 103,
	0, 106, 107, 0, 110, 111, 0, 0, 115, 116,
	117, 0, 119, 0, 0, 60, 61, 0, 64, 65,
	66, 0, 125, 126, 128, 0, 130, 131, 132, 0,
	67, 0, 0, 36, 38, 0, 0, 0, 83, 84,
	0, 87, 19, 19, 0, 0, 0, 68, 73, 74,
	0, 20, 31, 0, 0, 78, 79, 92, 95, 96,
	0, 0, 42, 45, 47, 50, 52, 0, 0, 51,
	101, 104, 105, 108, 112, 113, 0, 0, 0, 57,
	62, 63, 123, 127, 129, 135, 33, 35, 39, 0,
	41, 80, 85, 86, 0, 0, 90, 91, 5, 76,
	32, 77, 140, 134, 133, 0, 0, 55, 0, 137,
	0, 49, 114, 118, 120, 0, 88, 89, 53, 0,
	138, 0, 0, 40, 54, 56, 139,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:124
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yylval.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:130
		{
			d := yylval.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yylval.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:138
		{
			yylval.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:141
		{
			yylval.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:145
		{
			yylval.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:155
		{
			d := yylval.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:161
		{
			m := yylval.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:166
		{
			var err error
			if yylval.importer == nil {
				yylex.Error("No importer defined")
				goto ret1
			} else {
				m := yylval.stack.Peek().(*meta.Module)
				if err = yylval.importer(m, yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				}
			}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:212
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:224
		{
			yylval.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:234
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			yylval.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:249
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:256
		{
			yylval.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:270
		{
			if hasType, valid := yylval.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(tokenString(yyDollar[2].token))
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:279
		{
			y := yylval.stack.Peek().(meta.HasDataType)
			y.SetDataType(yylval.dataType)
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
			y := yylval.stack.Peek().(meta.HasDataType)
			yylval.dataType = meta.NewDataType(y, yyDollar[2].token)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:294
		{
			var err error
			if err = yylval.dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:302
		{
			yylval.dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:310
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:317
		{
			yylval.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:336
		{
			yylval.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:348
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:355
		{
			yylval.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:369
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:374
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:381
		{
			yylval.stack.Push(&meta.RpcInput{})
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:386
		{
			yylval.stack.Push(&meta.RpcOutput{})
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:394
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:401
		{
			yylval.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:415
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:420
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:427
		{
			yylval.stack.Push(&meta.RpcInput{})
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:432
		{
			yylval.stack.Push(&meta.RpcOutput{})
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:440
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:447
		{
			yylval.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:464
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:476
		{
			yylval.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:492
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:499
		{
			yylval.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:520
		{
			if list, valid := yylval.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:530
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:537
		{
			yylval.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:545
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:552
		{
			yylval.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:572
		{
			if hasDetails, valid := yylval.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:581
		{
			if hasDetails, valid := yylval.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:595
		{
			if HasError(yylex, popAndAddMeta(&yylval)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:602
		{
			yylval.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:611
		{
			yylval.dataType.AddEnumeration(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
