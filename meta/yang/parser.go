//line parser.y:2
package yang

import (
	"fmt"
	__yyfmt__ "fmt"
	"strings"

	"github.com/dhubler/c2g/meta"
)

//line parser.y:2
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
const yyInitialStackSize = 16

//line parser.y:623

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

const yyLast = 393

var yyAct = [...]int{

	130, 226, 129, 134, 162, 153, 148, 137, 117, 142,
	131, 169, 111, 7, 103, 7, 150, 26, 135, 3,
	96, 94, 133, 138, 223, 227, 227, 26, 10, 10,
	24, 10, 181, 132, 180, 127, 125, 47, 123, 245,
	6, 10, 54, 53, 9, 60, 51, 52, 55, 126,
	225, 56, 119, 58, 66, 150, 64, 59, 57, 6,
	10, 16, 10, 9, 109, 176, 108, 244, 11, 10,
	63, 109, 62, 108, 91, 114, 124, 243, 95, 106,
	105, 233, 112, 118, 164, 139, 139, 11, 232, 146,
	154, 163, 113, 104, 143, 121, 231, 141, 141, 230,
	155, 171, 171, 229, 95, 228, 120, 10, 140, 140,
	172, 221, 106, 105, 114, 167, 216, 177, 158, 159,
	124, 112, 209, 185, 203, 188, 104, 118, 200, 99,
	196, 113, 47, 197, 189, 186, 178, 54, 53, 121,
	60, 51, 52, 55, 199, 168, 56, 139, 58, 165,
	120, 144, 59, 57, 204, 76, 61, 17, 208, 141,
	171, 171, 237, 154, 236, 164, 215, 242, 210, 211,
	140, 10, 163, 155, 219, 235, 220, 218, 217, 125,
	47, 214, 207, 202, 201, 54, 53, 198, 60, 51,
	52, 55, 126, 195, 56, 10, 58, 187, 166, 213,
	59, 57, 212, 125, 47, 205, 175, 174, 234, 54,
	53, 80, 60, 51, 52, 55, 126, 10, 56, 184,
	58, 79, 78, 10, 59, 57, 239, 75, 100, 101,
	74, 73, 47, 72, 71, 196, 69, 54, 53, 99,
	60, 51, 52, 55, 67, 99, 56, 10, 58, 22,
	190, 240, 59, 57, 238, 194, 47, 45, 193, 192,
	65, 54, 53, 191, 60, 51, 52, 55, 182, 99,
	56, 173, 58, 47, 20, 19, 59, 57, 54, 53,
	39, 60, 51, 52, 55, 18, 10, 56, 241, 58,
	206, 5, 183, 59, 57, 47, 14, 90, 89, 88,
	54, 53, 87, 60, 51, 52, 55, 86, 85, 56,
	84, 58, 47, 83, 82, 59, 57, 54, 53, 39,
	60, 51, 52, 55, 81, 77, 56, 10, 58, 109,
	68, 108, 59, 57, 21, 125, 12, 136, 44, 46,
	122, 116, 115, 42, 110, 70, 41, 161, 126, 160,
	50, 157, 156, 152, 151, 49, 98, 97, 93, 92,
	27, 128, 43, 224, 222, 179, 107, 102, 40, 149,
	147, 145, 48, 38, 37, 36, 35, 34, 33, 32,
	31, 30, 29, 28, 170, 25, 8, 15, 23, 13,
	4, 2, 1,
}
var yyPact = [...]int{

	-6, -1000, 29, 332, 48, 148, 280, -1000, -1000, 270,
	269, 330, 242, 291, 147, 63, 46, -1000, -1000, -1000,
	-1000, -1000, -1000, 252, -1000, -1000, -1000, 237, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 326,
	229, 227, 226, 224, 223, 220, 146, 321, 215, 214,
	204, 320, 310, 309, 306, 304, 303, 298, 295, 294,
	293, -1000, -1000, 19, -1000, -1000, -1000, 205, -1000, 50,
	-1000, 235, 16, 183, 315, 315, -1000, 142, 17, 95,
	274, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 140, 190, 205, -1000, 136, -1000, 111, 111, 266,
	200, 199, 57, -1000, -1000, 127, -1000, 25, 263, 288,
	211, -1000, 126, -1000, -1000, 189, 16, -1000, 125, 244,
	-1000, -1000, -1000, 258, -1000, 254, 253, 250, 185, 183,
	-1000, 124, -1000, -1000, -1000, 179, 315, -1000, -1000, 119,
	-1000, -1000, -1000, 176, -1000, 175, 115, -22, -1000, 198,
	286, 174, 95, -1000, 113, -1000, 111, 111, 195, 192,
	173, 274, -1000, 107, -1000, 170, -1000, -1000, -1000, 169,
	111, -1000, 168, 102, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 7, 96, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	94, 90, 87, 79, 72, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 183, -1000, -1000, -1000, -1000,
	167, 156, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 154, 249, 8, 246, -1000, 284, -1000, -1000,
	-1000, -1000, -1000, -1000, 159, -1000, -1000, -1000, 68, -1000,
	58, 30, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 392, 391, 390, 389, 388, 387, 10, 291, 386,
	30, 385, 3, 11, 384, 383, 382, 381, 380, 379,
	378, 377, 376, 375, 374, 373, 372, 371, 370, 6,
	369, 2, 368, 367, 14, 23, 9, 366, 365, 364,
	363, 362, 361, 0, 33, 22, 360, 359, 358, 21,
	20, 357, 356, 355, 354, 353, 5, 352, 351, 350,
	349, 347, 4, 346, 345, 344, 12, 343, 342, 341,
	8, 340, 339, 338, 18, 337, 7, 257, 1,
}
var yyR1 = [...]int{

	0, 1, 2, 6, 4, 4, 7, 3, 3, 8,
	8, 8, 8, 9, 10, 10, 5, 5, 13, 13,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 14, 14, 23, 27, 27, 27, 26, 28, 28,
	29, 30, 15, 32, 33, 33, 34, 34, 34, 36,
	35, 37, 38, 38, 39, 39, 39, 18, 41, 42,
	42, 31, 31, 43, 43, 43, 43, 22, 11, 46,
	47, 47, 48, 48, 49, 49, 49, 49, 51, 52,
	24, 53, 54, 54, 55, 55, 56, 56, 56, 56,
	57, 58, 25, 59, 60, 60, 61, 61, 62, 62,
	16, 64, 63, 65, 65, 66, 66, 66, 17, 67,
	68, 69, 69, 70, 70, 70, 70, 70, 70, 70,
	71, 21, 72, 19, 73, 74, 75, 75, 76, 76,
	76, 76, 76, 45, 44, 20, 77, 40, 40, 78,
	50,
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
	3,
}
var yyChk = [...]int{

	-1000, -1, -2, 25, -3, -8, 11, -7, -9, 15,
	12, 39, 4, -4, -8, -6, 13, 9, 5, 5,
	5, 4, 7, -5, -10, -11, -12, -46, -15, -16,
	-17, -18, -19, -20, -21, -22, -23, -24, -25, 28,
	-32, -63, -67, -41, -73, -77, -72, 21, -26, -53,
	-59, 30, 31, 27, 26, 32, 35, 42, 37, 41,
	29, 9, 9, 7, 10, 8, -10, 7, 4, 7,
	-64, 7, 7, 7, 7, 7, 9, 4, 7, 7,
	7, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, -7, -47, -48, -49, -7, -50, -51, -52, 34,
	23, 24, -33, -34, -35, -7, -36, -37, 16, 14,
	-65, -66, -7, -50, -12, -68, -69, -70, -7, 36,
	-44, -45, -71, 22, -12, 20, 33, 19, -42, -31,
	-43, -7, -44, -45, -12, -74, -75, -76, -35, -7,
	-44, -45, -36, -74, 9, -27, -7, -28, -29, -30,
	38, -54, -55, -56, -7, -50, -57, -58, 23, 24,
	-60, -61, -62, -7, -12, 9, 8, -49, 9, -13,
	-14, -12, -13, 5, 7, 7, 8, -34, 9, -38,
	9, 7, 5, 4, 8, -66, 9, 8, -70, 9,
	6, 5, 5, 5, 5, 8, -43, 9, 8, -76,
	9, 8, 8, 9, -29, 7, 4, 8, -56, 9,
	-13, -13, 7, 7, 8, -62, 9, 8, 8, -12,
	8, 9, -39, 17, -40, 43, -78, 18, 9, 9,
	9, 9, 9, 9, -31, 8, 8, 8, 5, -78,
	5, 4, 8, 9, 9, 9,
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
	77, 140, 0, 0, 55, 0, 137, 0, 49, 114,
	118, 134, 133, 120, 0, 88, 89, 53, 0, 138,
	0, 0, 40, 54, 56, 139,
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
		//line parser.y:123
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:129
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:137
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:140
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:144
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:154
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:160
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:165
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
		//line parser.y:211
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:223
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:233
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:240
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:248
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:255
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:269
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
		//line parser.y:281
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:291
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
		//line parser.y:301
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:311
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:318
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:337
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:349
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:356
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:370
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:375
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:382
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:387
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:395
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:402
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:416
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:421
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:428
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:433
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:441
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:449
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:468
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:480
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:496
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:503
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:524
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
		//line parser.y:534
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:541
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:549
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:556
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:576
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
		//line parser.y:585
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
		//line parser.y:599
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:606
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:615
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
