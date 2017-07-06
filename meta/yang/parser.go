//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/c2stack/c2g/meta"
	"strconv"
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

//line parser.y:64
type yySymType struct {
	yys      int
	ident    string
	token    string
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
const kywd_value = 57386

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
	"kywd_value",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:640

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 399

var yyAct = [...]int{

	130, 226, 129, 134, 162, 153, 148, 137, 117, 142,
	131, 169, 111, 7, 103, 7, 248, 26, 135, 150,
	96, 94, 133, 138, 223, 227, 3, 26, 10, 10,
	24, 227, 10, 132, 64, 127, 125, 47, 123, 251,
	6, 10, 54, 53, 9, 60, 51, 52, 55, 126,
	225, 56, 119, 58, 66, 150, 244, 59, 57, 6,
	10, 16, 10, 9, 109, 176, 108, 243, 11, 10,
	246, 109, 245, 108, 91, 114, 124, 233, 95, 106,
	105, 232, 112, 118, 164, 139, 139, 11, 231, 146,
	154, 163, 113, 104, 143, 121, 230, 141, 141, 229,
	155, 171, 171, 181, 95, 180, 120, 10, 140, 140,
	172, 228, 106, 105, 114, 167, 221, 177, 158, 159,
	124, 112, 63, 185, 62, 188, 104, 118, 216, 99,
	196, 113, 47, 209, 203, 200, 197, 54, 53, 121,
	60, 51, 52, 55, 199, 189, 56, 139, 58, 186,
	120, 178, 59, 57, 204, 168, 165, 144, 208, 141,
	171, 171, 76, 154, 61, 164, 215, 242, 210, 211,
	140, 10, 163, 155, 219, 17, 249, 237, 236, 125,
	47, 235, 220, 218, 217, 54, 53, 214, 60, 51,
	52, 55, 126, 207, 56, 10, 58, 202, 201, 198,
	59, 57, 195, 125, 47, 187, 166, 213, 234, 54,
	53, 212, 60, 51, 52, 55, 126, 10, 56, 184,
	58, 205, 175, 10, 59, 57, 239, 174, 100, 101,
	80, 79, 47, 78, 75, 196, 74, 54, 53, 99,
	60, 51, 52, 55, 73, 99, 56, 10, 58, 72,
	71, 69, 59, 57, 67, 22, 47, 190, 250, 173,
	65, 54, 53, 240, 60, 51, 52, 55, 238, 99,
	56, 194, 58, 47, 193, 192, 59, 57, 54, 53,
	39, 60, 51, 52, 55, 191, 10, 56, 182, 58,
	20, 19, 18, 59, 57, 47, 241, 5, 206, 183,
	54, 53, 14, 60, 51, 52, 55, 90, 89, 56,
	88, 58, 47, 87, 86, 59, 57, 54, 53, 39,
	60, 51, 52, 55, 85, 84, 56, 10, 58, 109,
	83, 108, 59, 57, 82, 125, 81, 77, 68, 21,
	12, 45, 136, 44, 46, 122, 116, 115, 126, 42,
	110, 70, 41, 161, 160, 50, 157, 156, 152, 151,
	49, 98, 97, 93, 92, 27, 128, 43, 224, 222,
	179, 107, 102, 40, 149, 147, 145, 48, 38, 37,
	36, 35, 34, 33, 32, 31, 30, 29, 28, 170,
	25, 8, 15, 23, 13, 4, 2, 1, 247,
}
var yyPact = [...]int{

	1, -1000, 29, 336, 48, 166, 287, -1000, -1000, 286,
	285, 335, 248, 291, 155, 115, 24, -1000, -1000, -1000,
	-1000, -1000, -1000, 252, -1000, -1000, -1000, 247, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 334,
	244, 243, 242, 237, 229, 227, 153, 333, 226, 224,
	223, 332, 330, 326, 321, 320, 310, 309, 306, 304,
	303, -1000, -1000, 20, -1000, -1000, -1000, 205, -1000, 50,
	-1000, 235, 16, 183, 315, 315, -1000, 148, 17, 95,
	274, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 147, 198, 205, -1000, 146, -1000, 111, 111, 253,
	220, 215, 57, -1000, -1000, 142, -1000, 96, 283, 295,
	211, -1000, 140, -1000, -1000, 197, 16, -1000, 136, 251,
	-1000, -1000, -1000, 280, -1000, 270, 269, 266, 194, 183,
	-1000, 127, -1000, -1000, -1000, 191, 315, -1000, -1000, 126,
	-1000, -1000, -1000, 190, -1000, 189, 125, -19, -1000, 214,
	294, 185, 95, -1000, 124, -1000, 111, 111, 204, 200,
	179, 274, -1000, 119, -1000, 176, -1000, -1000, -1000, 175,
	111, -1000, 174, 107, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 7, 102, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	90, 87, 79, 72, 68, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 183, -1000, -1000, -1000, -1000,
	173, 170, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 169, 263, 13, 258, -1000, 292, -1000, -1000,
	-1000, -1000, -1000, -1000, 159, -1000, -1000, -1000, 58, -1000,
	47, 63, -1000, -1000, -1000, -1000, -28, 168, 254, -1000,
	30, -1000,
}
var yyPgo = [...]int{

	0, 398, 397, 396, 395, 394, 393, 392, 10, 297,
	391, 30, 390, 3, 11, 389, 388, 387, 386, 385,
	384, 383, 382, 381, 380, 379, 378, 377, 376, 375,
	6, 374, 2, 373, 372, 14, 23, 9, 371, 370,
	369, 368, 367, 366, 0, 33, 22, 365, 364, 363,
	21, 20, 362, 361, 360, 359, 358, 5, 357, 356,
	355, 354, 353, 4, 352, 351, 350, 12, 349, 347,
	346, 8, 345, 344, 343, 18, 342, 7, 341, 1,
}
var yyR1 = [...]int{

	0, 2, 3, 7, 5, 5, 8, 4, 4, 9,
	9, 9, 9, 10, 11, 11, 6, 6, 14, 14,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 15, 15, 24, 28, 28, 28, 27, 29, 29,
	30, 31, 16, 33, 34, 34, 35, 35, 35, 37,
	36, 38, 39, 39, 40, 40, 40, 19, 42, 43,
	43, 32, 32, 44, 44, 44, 44, 23, 12, 47,
	48, 48, 49, 49, 50, 50, 50, 50, 52, 53,
	25, 54, 55, 55, 56, 56, 57, 57, 57, 57,
	58, 59, 26, 60, 61, 61, 62, 62, 63, 63,
	17, 65, 64, 66, 66, 67, 67, 67, 18, 68,
	69, 70, 70, 71, 71, 71, 71, 71, 71, 71,
	72, 22, 73, 20, 74, 75, 76, 76, 77, 77,
	77, 77, 77, 46, 45, 21, 78, 41, 41, 79,
	79, 1, 51,
}
var yyR2 = [...]int{

	0, 5, 3, 2, 2, 5, 2, 2, 3, 2,
	1, 1, 2, 2, 1, 1, 1, 2, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 4, 0, 2, 1, 2, 1, 2,
	4, 2, 4, 2, 1, 2, 1, 2, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 2, 1, 1, 1, 3, 4, 2,
	0, 1, 1, 2, 2, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 2, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 2, 1,
	2, 3, 2, 1, 2, 2, 1, 1, 4, 2,
	1, 1, 2, 2, 3, 1, 1, 1, 3, 1,
	3, 2, 2, 4, 2, 1, 1, 2, 1, 2,
	1, 1, 1, 3, 3, 4, 2, 1, 2, 3,
	5, 3, 3,
}
var yyChk = [...]int{

	-1000, -2, -3, 25, -4, -9, 11, -8, -10, 15,
	12, 39, 4, -5, -9, -7, 13, 9, 5, 5,
	5, 4, 7, -6, -11, -12, -13, -47, -16, -17,
	-18, -19, -20, -21, -22, -23, -24, -25, -26, 28,
	-33, -64, -68, -42, -74, -78, -73, 21, -27, -54,
	-60, 30, 31, 27, 26, 32, 35, 42, 37, 41,
	29, 9, 9, 7, 10, 8, -11, 7, 4, 7,
	-65, 7, 7, 7, 7, 7, 9, 4, 7, 7,
	7, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, -8, -48, -49, -50, -8, -51, -52, -53, 34,
	23, 24, -34, -35, -36, -8, -37, -38, 16, 14,
	-66, -67, -8, -51, -13, -69, -70, -71, -8, 36,
	-45, -46, -72, 22, -13, 20, 33, 19, -43, -32,
	-44, -8, -45, -46, -13, -75, -76, -77, -36, -8,
	-45, -46, -37, -75, 9, -28, -8, -29, -30, -31,
	38, -55, -56, -57, -8, -51, -58, -59, 23, 24,
	-61, -62, -63, -8, -13, 9, 8, -50, 9, -14,
	-15, -13, -14, 6, 7, 7, 8, -35, 9, -39,
	9, 7, 5, 4, 8, -67, 9, 8, -71, 9,
	6, 5, 5, 5, 5, 8, -44, 9, 8, -77,
	9, 8, 8, 9, -30, 7, 4, 8, -57, 9,
	-14, -14, 7, 7, 8, -63, 9, 8, 8, -13,
	8, 9, -40, 17, -41, 43, -79, 18, 9, 9,
	9, 9, 9, 9, -32, 8, 8, 8, 5, -79,
	5, 4, 8, 9, 9, 9, 7, -1, 44, 8,
	4, 9,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 10, 11, 0,
	0, 0, 0, 0, 0, 0, 0, 7, 9, 12,
	6, 13, 2, 0, 16, 14, 15, 0, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 4, 0, 3, 1, 17, 70, 69, 0,
	100, 0, 0, 59, 0, 0, 121, 0, 34, 82,
	94, 43, 102, 109, 58, 124, 136, 122, 37, 81,
	93, 0, 0, 71, 72, 0, 75, 18, 18, 0,
	0, 0, 0, 44, 46, 0, 48, 0, 0, 0,
	0, 103, 0, 106, 107, 0, 110, 111, 0, 0,
	115, 116, 117, 0, 119, 0, 0, 0, 0, 60,
	61, 0, 64, 65, 66, 0, 125, 126, 128, 0,
	130, 131, 132, 0, 67, 0, 0, 36, 38, 0,
	0, 0, 83, 84, 0, 87, 18, 18, 0, 0,
	0, 95, 96, 0, 99, 0, 68, 73, 74, 0,
	19, 31, 0, 0, 78, 79, 42, 45, 47, 50,
	52, 0, 0, 51, 101, 104, 105, 108, 112, 113,
	0, 0, 0, 0, 0, 57, 62, 63, 123, 127,
	129, 135, 33, 35, 39, 0, 41, 80, 85, 86,
	0, 0, 90, 91, 92, 97, 98, 5, 76, 32,
	77, 142, 0, 0, 55, 0, 137, 0, 49, 114,
	118, 134, 133, 120, 0, 88, 89, 53, 0, 138,
	0, 0, 40, 54, 56, 139, 0, 0, 0, 140,
	0, 141,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44,
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
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
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
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			yyrcvr.char = -1
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
		//line parser.y:126
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:132
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:140
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:143
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:147
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:157
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:163
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:168
		{
			var err error
			if yyVAL.importer == nil {
				yylex.Error("No importer defined")
				goto ret1
			} else {
				m := yyVAL.stack.Peek().(*meta.Module)
				if err = yyVAL.importer(m, yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				}
			}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:214
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:226
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:236
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:243
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:251
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:258
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:272
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(tokenString(yyDollar[2].token))
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:294
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:304
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:314
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:321
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:340
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:352
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:373
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:378
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:385
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:390
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:398
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:405
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:419
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:424
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:431
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:436
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:444
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:452
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:483
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:499
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:506
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:527
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:537
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:544
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:552
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:559
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:579
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:588
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:602
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:609
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:618
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:622
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:633
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
