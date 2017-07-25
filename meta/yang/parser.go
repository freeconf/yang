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

const yyLast = 402

var yyAct = [...]int{

	132, 228, 131, 136, 164, 155, 150, 139, 20, 119,
	133, 171, 113, 8, 105, 8, 137, 96, 250, 12,
	20, 152, 98, 12, 3, 135, 144, 229, 12, 140,
	129, 127, 41, 125, 134, 63, 65, 48, 47, 253,
	54, 45, 46, 49, 128, 152, 50, 121, 52, 41,
	18, 246, 53, 51, 48, 47, 33, 54, 45, 46,
	49, 225, 229, 50, 12, 52, 111, 66, 110, 53,
	51, 93, 245, 178, 251, 116, 126, 12, 97, 111,
	107, 110, 114, 120, 166, 141, 141, 227, 235, 148,
	156, 165, 145, 248, 115, 247, 108, 234, 123, 106,
	143, 143, 157, 173, 173, 233, 97, 122, 232, 142,
	142, 231, 174, 169, 230, 107, 116, 223, 183, 179,
	182, 218, 126, 114, 60, 187, 59, 211, 190, 120,
	205, 108, 198, 239, 106, 115, 12, 202, 6, 12,
	14, 199, 10, 191, 123, 188, 201, 160, 161, 141,
	180, 170, 167, 122, 146, 92, 206, 91, 101, 76,
	210, 57, 173, 173, 143, 156, 13, 166, 217, 56,
	212, 213, 238, 142, 165, 215, 221, 157, 6, 12,
	14, 237, 10, 222, 220, 219, 216, 209, 41, 204,
	203, 200, 197, 48, 47, 33, 54, 45, 46, 49,
	189, 214, 50, 168, 52, 12, 13, 207, 53, 51,
	236, 244, 177, 176, 80, 12, 102, 103, 79, 78,
	75, 74, 73, 127, 41, 72, 71, 101, 241, 48,
	47, 69, 54, 45, 46, 49, 128, 198, 50, 12,
	52, 67, 64, 192, 53, 51, 175, 127, 41, 242,
	240, 196, 195, 48, 47, 194, 54, 45, 46, 49,
	128, 186, 50, 193, 52, 12, 184, 61, 53, 51,
	58, 55, 5, 252, 41, 243, 208, 17, 185, 48,
	47, 39, 54, 45, 46, 49, 90, 101, 50, 12,
	52, 89, 88, 87, 53, 51, 86, 85, 41, 84,
	83, 82, 81, 48, 47, 138, 54, 45, 46, 49,
	77, 101, 50, 12, 52, 68, 62, 15, 53, 51,
	38, 40, 41, 124, 118, 117, 36, 48, 47, 112,
	54, 45, 46, 49, 70, 35, 50, 163, 52, 41,
	162, 44, 53, 51, 48, 47, 159, 54, 45, 46,
	49, 158, 154, 50, 12, 52, 111, 153, 110, 53,
	51, 43, 127, 100, 99, 95, 94, 21, 130, 37,
	226, 224, 181, 109, 104, 128, 34, 151, 149, 147,
	42, 32, 31, 30, 29, 28, 27, 26, 25, 24,
	23, 22, 172, 19, 9, 7, 11, 16, 4, 2,
	1, 249,
}
var yyPact = [...]int{

	-1, -1000, 127, 313, 167, -1000, 266, -1000, 160, 152,
	265, 117, 262, 312, 25, 235, 28, -1000, -1000, -1000,
	-1000, 234, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 311, 224, 219, 218, 215, 214, 213,
	150, 306, 212, 211, 207, 298, 297, 296, 295, 293,
	292, 289, 288, 287, 282, 148, -1000, -1000, 146, -1000,
	16, -1000, -1000, -1000, -1000, -1000, -1000, 193, -1000, 52,
	-1000, 277, 11, 227, 342, 342, -1000, 145, 7, 124,
	301, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 143, 195, 193, -1000, 142, -1000, 318,
	318, 240, 206, 205, 65, -1000, -1000, 141, -1000, 111,
	261, 274, 253, -1000, 136, -1000, -1000, 192, 11, -1000,
	134, 237, -1000, -1000, -1000, 258, -1000, 250, 247, 246,
	184, 227, -1000, 132, -1000, -1000, -1000, 183, 342, -1000,
	-1000, 128, -1000, -1000, -1000, 182, -1000, 181, 121, -17,
	-1000, 200, 272, 179, 124, -1000, 118, -1000, 318, 318,
	194, 168, 178, 301, -1000, 112, -1000, 177, -1000, -1000,
	-1000, 176, 318, -1000, 175, 108, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 44, 105, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 102, 99, 96, 88, 79, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 227, -1000, -1000,
	-1000, -1000, 173, 164, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 125, 245, 9, 244, -1000, 271,
	-1000, -1000, -1000, -1000, -1000, -1000, 203, -1000, -1000, -1000,
	63, -1000, 42, 86, -1000, -1000, -1000, -1000, -26, 66,
	269, -1000, 30, -1000,
}
var yyPgo = [...]int{

	0, 401, 400, 399, 398, 397, 396, 395, 10, 272,
	394, 50, 393, 3, 11, 392, 391, 390, 389, 388,
	387, 386, 385, 384, 383, 382, 381, 380, 379, 378,
	6, 377, 2, 376, 374, 14, 29, 26, 373, 372,
	371, 370, 369, 368, 0, 34, 25, 367, 366, 365,
	17, 22, 364, 363, 361, 357, 352, 5, 351, 346,
	341, 340, 337, 4, 335, 334, 329, 12, 326, 325,
	324, 9, 323, 321, 320, 16, 305, 7, 281, 1,
}
var yyR1 = [...]int{

	0, 2, 3, 6, 7, 7, 8, 4, 4, 9,
	9, 9, 9, 9, 10, 11, 11, 5, 5, 14,
	14, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 15, 15, 24, 28, 28, 28, 27, 29,
	29, 30, 31, 16, 33, 34, 34, 35, 35, 35,
	37, 36, 38, 39, 39, 40, 40, 40, 19, 42,
	43, 43, 32, 32, 44, 44, 44, 44, 23, 12,
	47, 48, 48, 49, 49, 50, 50, 50, 50, 52,
	53, 25, 54, 55, 55, 56, 56, 57, 57, 57,
	57, 58, 59, 26, 60, 61, 61, 62, 62, 63,
	63, 17, 65, 64, 66, 66, 67, 67, 67, 18,
	68, 69, 70, 70, 71, 71, 71, 71, 71, 71,
	71, 72, 22, 73, 20, 74, 75, 76, 76, 77,
	77, 77, 77, 77, 46, 45, 21, 78, 41, 41,
	79, 79, 1, 51,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 5, 2, 1, 2, 3,
	1, 2, 2, 3, 2, 1, 1, 1, 2, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 4, 0, 2, 1, 2, 1,
	2, 4, 2, 4, 2, 1, 2, 1, 2, 1,
	3, 2, 2, 1, 3, 3, 1, 3, 4, 2,
	0, 1, 1, 2, 2, 1, 1, 1, 3, 4,
	2, 0, 1, 1, 2, 2, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 2, 1, 3,
	3, 2, 2, 4, 2, 0, 1, 1, 2, 2,
	1, 2, 3, 2, 1, 2, 2, 1, 1, 4,
	2, 1, 1, 2, 2, 3, 1, 1, 1, 3,
	1, 3, 2, 2, 4, 2, 1, 1, 2, 1,
	2, 1, 1, 1, 3, 3, 4, 2, 1, 2,
	3, 5, 3, 3,
}
var yyChk = [...]int{

	-1000, -2, -3, 25, -4, -9, 11, -7, -8, -10,
	15, -6, 12, 39, 13, 4, -5, -9, -11, -12,
	-13, -47, -16, -17, -18, -19, -20, -21, -22, -23,
	-24, -25, -26, 28, -33, -64, -68, -42, -74, -78,
	-73, 21, -27, -54, -60, 30, 31, 27, 26, 32,
	35, 42, 37, 41, 29, 5, 9, 9, 5, 9,
	7, 5, 4, 10, 7, 8, -11, 7, 4, 7,
	-65, 7, 7, 7, 7, 7, 9, 4, 7, 7,
	7, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 9, 9, -8, -48, -49, -50, -8, -51, -52,
	-53, 34, 23, 24, -34, -35, -36, -8, -37, -38,
	16, 14, -66, -67, -8, -51, -13, -69, -70, -71,
	-8, 36, -45, -46, -72, 22, -13, 20, 33, 19,
	-43, -32, -44, -8, -45, -46, -13, -75, -76, -77,
	-36, -8, -45, -46, -37, -75, 9, -28, -8, -29,
	-30, -31, 38, -55, -56, -57, -8, -51, -58, -59,
	23, 24, -61, -62, -63, -8, -13, 9, 8, -50,
	9, -14, -15, -13, -14, 6, 7, 7, 8, -35,
	9, -39, 9, 7, 5, 4, 8, -67, 9, 8,
	-71, 9, 6, 5, 5, 5, 5, 8, -44, 9,
	8, -77, 9, 8, 8, 9, -30, 7, 4, 8,
	-57, 9, -14, -14, 7, 7, 8, -63, 9, 8,
	8, -13, 8, 9, -40, 17, -41, 43, -79, 18,
	9, 9, 9, 9, 9, 9, -32, 8, 8, 8,
	5, -79, 5, 4, 8, 9, 9, 9, 7, -1,
	44, 8, 4, 9,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 17, 15,
	16, 0, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 11, 12, 0, 4,
	0, 6, 14, 3, 2, 1, 18, 71, 70, 0,
	101, 0, 0, 60, 0, 0, 122, 0, 35, 83,
	95, 44, 103, 110, 59, 125, 137, 123, 38, 82,
	94, 9, 13, 0, 0, 72, 73, 0, 76, 19,
	19, 0, 0, 0, 0, 45, 47, 0, 49, 0,
	0, 0, 0, 104, 0, 107, 108, 0, 111, 112,
	0, 0, 116, 117, 118, 0, 120, 0, 0, 0,
	0, 61, 62, 0, 65, 66, 67, 0, 126, 127,
	129, 0, 131, 132, 133, 0, 68, 0, 0, 37,
	39, 0, 0, 0, 84, 85, 0, 88, 19, 19,
	0, 0, 0, 96, 97, 0, 100, 0, 69, 74,
	75, 0, 20, 32, 0, 0, 79, 80, 43, 46,
	48, 51, 53, 0, 0, 52, 102, 105, 106, 109,
	113, 114, 0, 0, 0, 0, 0, 58, 63, 64,
	124, 128, 130, 136, 34, 36, 40, 0, 42, 81,
	86, 87, 0, 0, 91, 92, 93, 98, 99, 5,
	77, 33, 78, 143, 0, 0, 56, 0, 138, 0,
	50, 115, 119, 135, 134, 121, 0, 89, 90, 54,
	0, 139, 0, 0, 41, 55, 57, 140, 0, 0,
	0, 141, 0, 142,
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
		//line parser.y:125
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:131
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:139
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:142
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:146
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:156
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:163
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 14:
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
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:214
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:226
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:236
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:243
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:251
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:258
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 50:
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
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 55:
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
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:304
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:314
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:321
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:340
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:352
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:373
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:378
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:385
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:390
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:398
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:405
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:419
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:424
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:431
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:436
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:444
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:452
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:483
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:499
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:506
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 121:
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
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:537
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:544
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:552
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:559
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 134:
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
	case 135:
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
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:602
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:609
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:618
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 141:
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
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:633
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
