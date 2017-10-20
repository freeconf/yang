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
	lval.loader = l.loader
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
	yys     int
	ident   string
	token   string
	boolean bool
	stack   *yangMetaStack
	loader  ModuleLoader
}

const token_ident = 57346
const token_string = 57347
const token_int = 57348
const token_number = 57349
const token_curly_open = 57350
const token_curly_close = 57351
const token_semi = 57352
const token_rev_ident = 57353
const kywd_namespace = 57354
const kywd_description = 57355
const kywd_revision = 57356
const kywd_type = 57357
const kywd_prefix = 57358
const kywd_default = 57359
const kywd_length = 57360
const kywd_enum = 57361
const kywd_key = 57362
const kywd_config = 57363
const kywd_uses = 57364
const kywd_unique = 57365
const kywd_input = 57366
const kywd_output = 57367
const kywd_module = 57368
const kywd_container = 57369
const kywd_list = 57370
const kywd_rpc = 57371
const kywd_notification = 57372
const kywd_typedef = 57373
const kywd_grouping = 57374
const kywd_leaf = 57375
const kywd_mandatory = 57376
const kywd_reference = 57377
const kywd_leaf_list = 57378
const kywd_max_elements = 57379
const kywd_choice = 57380
const kywd_case = 57381
const kywd_import = 57382
const kywd_include = 57383
const kywd_action = 57384
const kywd_anyxml = 57385
const kywd_path = 57386
const kywd_value = 57387
const kywd_true = 57388
const kywd_false = 57389

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_int",
	"token_number",
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
	"kywd_true",
	"kywd_false",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:719

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 453

var yyAct = [...]int{

	151, 263, 150, 155, 183, 174, 169, 158, 24, 224,
	152, 200, 138, 8, 132, 8, 124, 116, 110, 103,
	107, 285, 154, 163, 24, 156, 6, 13, 16, 159,
	11, 225, 226, 153, 260, 264, 45, 13, 171, 3,
	264, 52, 51, 37, 58, 49, 50, 53, 13, 195,
	54, 22, 56, 190, 17, 18, 57, 55, 13, 69,
	262, 188, 283, 171, 282, 13, 105, 206, 104, 179,
	180, 13, 74, 130, 101, 129, 211, 106, 210, 112,
	108, 288, 281, 135, 145, 280, 117, 108, 126, 113,
	133, 139, 185, 160, 160, 270, 118, 167, 175, 184,
	134, 127, 68, 142, 67, 162, 162, 125, 176, 164,
	66, 269, 65, 106, 141, 268, 161, 161, 267, 63,
	112, 62, 187, 202, 202, 266, 117, 13, 193, 265,
	113, 255, 203, 198, 126, 135, 118, 254, 121, 122,
	207, 145, 133, 253, 252, 250, 217, 127, 139, 108,
	220, 230, 134, 125, 243, 13, 105, 227, 104, 13,
	142, 130, 13, 129, 130, 233, 129, 237, 160, 194,
	146, 141, 234, 13, 111, 238, 231, 108, 221, 242,
	162, 202, 202, 147, 175, 218, 185, 249, 208, 244,
	245, 161, 13, 184, 176, 108, 199, 13, 111, 148,
	146, 45, 144, 196, 191, 257, 52, 51, 186, 58,
	49, 50, 53, 147, 279, 54, 140, 56, 13, 108,
	165, 57, 55, 100, 99, 84, 146, 45, 60, 286,
	274, 273, 52, 51, 272, 58, 49, 50, 53, 147,
	258, 54, 271, 56, 256, 251, 247, 57, 55, 248,
	13, 241, 236, 6, 13, 16, 235, 11, 146, 45,
	232, 229, 219, 276, 52, 51, 197, 58, 49, 50,
	53, 147, 230, 54, 216, 56, 246, 239, 13, 57,
	55, 17, 18, 213, 205, 214, 43, 45, 204, 88,
	87, 86, 52, 51, 83, 58, 49, 50, 53, 82,
	108, 54, 13, 56, 81, 80, 79, 57, 55, 77,
	75, 45, 72, 222, 277, 73, 52, 51, 275, 58,
	49, 50, 53, 228, 108, 54, 223, 56, 45, 192,
	189, 57, 55, 52, 51, 37, 58, 49, 50, 53,
	64, 13, 54, 61, 56, 59, 5, 287, 57, 55,
	45, 21, 278, 240, 215, 52, 51, 98, 58, 49,
	50, 53, 97, 96, 54, 95, 56, 45, 94, 93,
	57, 55, 52, 51, 92, 58, 49, 50, 53, 91,
	90, 54, 89, 56, 85, 76, 71, 57, 55, 70,
	19, 157, 42, 44, 143, 137, 136, 40, 131, 78,
	39, 182, 181, 48, 178, 177, 173, 172, 47, 120,
	119, 115, 114, 25, 149, 41, 261, 259, 209, 128,
	123, 38, 170, 168, 166, 46, 36, 35, 34, 33,
	32, 31, 30, 29, 28, 27, 26, 201, 23, 109,
	15, 102, 14, 10, 9, 7, 12, 20, 4, 2,
	1, 212, 284,
}
var yyPact = [...]int{

	13, -1000, 241, 386, 14, -1000, 340, -1000, 218, -1000,
	-1000, 338, 111, 335, 102, 94, 48, 385, 382, 304,
	306, -1000, -1000, -1000, -1000, 302, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 381, 301, 298,
	297, 296, 291, 286, 215, 380, 283, 282, 281, 378,
	376, 375, 370, 365, 364, 361, 359, 358, 353, 214,
	-1000, 213, -1000, 35, -1000, -1000, 142, -1000, 184, -1000,
	-1000, -1000, -1000, -1000, -1000, 114, -1000, 146, -1000, 289,
	179, 237, 149, 149, -1000, 210, 24, 45, 328, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 198, 52, -1000, 325, 42, 194, -1000, 324, 160,
	-1000, 38, 193, -1000, 257, 114, -1000, 186, -1000, 345,
	345, 280, 276, 58, -1000, -1000, 178, -1000, 68, 278,
	350, 265, -1000, 175, -1000, -1000, 253, 179, -1000, 168,
	307, -1000, -1000, -1000, 321, -1000, -15, -15, 318, 252,
	237, -1000, 166, -1000, -1000, -1000, 251, 149, -1000, -1000,
	162, -1000, -1000, -1000, 247, -1000, 243, 157, -1, -1000,
	269, 349, 242, 45, -1000, 144, -1000, 345, 345, 268,
	238, 240, 328, -1000, 135, -1000, 236, -1000, -1000, 134,
	133, -1000, 127, -1000, -1000, 121, -1000, -1000, -1000, -1000,
	235, 345, -1000, 231, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 16, 119, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 115, 108, 105, -1000, -1000, 101, 85, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 237,
	-1000, -1000, -1000, -1000, 225, 222, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 221,
	313, 21, 309, -1000, 348, -1000, -1000, -1000, -1000, -1000,
	-1000, 205, -1000, -1000, -1000, 75, -1000, 72, 54, -1000,
	-1000, -1000, -1000, -24, 220, 343, -1000, 71, -1000,
}
var yyPgo = [...]int{

	0, 452, 9, 451, 450, 449, 448, 447, 446, 445,
	10, 346, 444, 443, 442, 441, 19, 20, 440, 439,
	18, 51, 438, 3, 11, 437, 436, 435, 434, 433,
	432, 431, 430, 429, 428, 427, 426, 425, 424, 423,
	6, 422, 2, 421, 420, 16, 29, 23, 419, 418,
	417, 416, 415, 414, 0, 33, 22, 413, 412, 411,
	17, 410, 409, 408, 407, 406, 5, 405, 404, 403,
	402, 401, 4, 400, 399, 398, 14, 397, 396, 395,
	12, 394, 393, 392, 25, 391, 7, 286, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 6, 6, 11,
	11, 11, 11, 11, 11, 14, 15, 15, 16, 16,
	16, 16, 12, 12, 18, 19, 19, 20, 20, 20,
	13, 13, 21, 21, 7, 7, 24, 24, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 25,
	25, 34, 38, 38, 38, 37, 39, 39, 40, 41,
	26, 43, 44, 44, 45, 45, 45, 3, 3, 47,
	46, 48, 49, 49, 50, 50, 50, 29, 52, 53,
	53, 42, 42, 54, 54, 54, 54, 33, 22, 57,
	58, 58, 59, 59, 60, 60, 60, 60, 61, 62,
	35, 63, 64, 64, 65, 65, 66, 66, 66, 66,
	67, 68, 36, 69, 70, 70, 71, 71, 72, 72,
	27, 74, 73, 75, 75, 76, 76, 76, 28, 77,
	78, 79, 79, 80, 80, 80, 80, 80, 80, 80,
	81, 32, 82, 30, 83, 84, 85, 85, 86, 86,
	86, 86, 86, 56, 2, 2, 55, 31, 87, 51,
	51, 88, 88, 1, 17,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 5, 2, 1, 2, 3,
	1, 2, 1, 1, 3, 2, 1, 2, 3, 3,
	2, 1, 2, 4, 2, 1, 2, 3, 2, 1,
	2, 4, 1, 1, 1, 2, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 4, 0, 2, 1, 2, 1, 2, 4, 2,
	4, 2, 1, 2, 1, 2, 1, 1, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 2, 1, 1, 1, 3, 4, 2,
	0, 1, 1, 2, 2, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 2, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 2, 1,
	2, 3, 2, 1, 2, 2, 1, 1, 4, 2,
	1, 1, 2, 2, 3, 1, 1, 1, 3, 1,
	3, 2, 2, 4, 2, 1, 1, 2, 1, 2,
	1, 1, 1, 3, 1, 1, 3, 4, 2, 1,
	2, 3, 5, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 26, -6, -11, 12, -9, -10, -12,
	-13, 16, -8, 13, -14, -18, 14, 40, 41, 4,
	-7, -11, -21, -22, -23, -57, -26, -27, -28, -29,
	-30, -31, -32, -33, -34, -35, -36, 29, -43, -73,
	-77, -52, -83, -87, -82, 22, -37, -63, -69, 31,
	32, 28, 27, 33, 36, 43, 38, 42, 30, 5,
	10, 5, 10, 8, 5, 10, 8, 10, 8, 11,
	4, 4, 8, 9, -21, 8, 4, 8, -74, 8,
	8, 8, 8, 8, 10, 4, 8, 8, 8, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 10,
	10, -10, -15, -16, 16, 14, -10, -17, 35, -19,
	-20, 14, -10, -17, -58, -59, -60, -10, -17, -61,
	-62, 24, 25, -44, -45, -46, -10, -47, -48, 17,
	15, -75, -76, -10, -17, -23, -78, -79, -80, -10,
	37, -55, -56, -81, 23, -23, 21, 34, 20, -53,
	-42, -54, -10, -55, -56, -23, -84, -85, -86, -46,
	-10, -55, -56, -47, -84, 10, -38, -10, -39, -40,
	-41, 39, -64, -65, -66, -10, -17, -67, -68, 24,
	25, -70, -71, -72, -10, -23, 10, -16, 9, 5,
	11, 10, 5, -20, 9, 11, 10, 9, -60, 10,
	-24, -25, -23, -24, 8, 8, 9, -45, 10, -49,
	10, 8, -3, 5, 7, 4, 9, -76, 10, 9,
	-80, 10, 6, 5, -2, 46, 47, -2, 5, 9,
	-54, 10, 9, -86, 10, 9, 9, 10, -40, 8,
	4, 9, -66, 10, -24, -24, 8, 8, 9, -72,
	10, 9, 10, 10, 10, 10, 9, -23, 9, -50,
	18, -51, 44, -88, 19, 10, 10, 10, 10, 10,
	10, -42, 9, 9, 9, 5, -88, 5, 4, 9,
	10, 10, 10, 8, -1, 45, 9, 4, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 0, 12,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 34, 32, 33, 0, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 0, 4, 0, 6, 22, 0, 30, 0, 3,
	15, 24, 2, 1, 35, 90, 89, 0, 120, 0,
	0, 79, 0, 0, 141, 0, 52, 102, 114, 61,
	122, 129, 78, 144, 158, 142, 55, 101, 113, 9,
	14, 0, 0, 16, 0, 0, 0, 21, 0, 0,
	25, 0, 0, 29, 0, 91, 92, 0, 95, 36,
	36, 0, 0, 0, 62, 64, 0, 66, 0, 0,
	0, 0, 123, 0, 126, 127, 0, 130, 131, 0,
	0, 135, 136, 137, 0, 139, 0, 0, 0, 0,
	80, 81, 0, 84, 85, 86, 0, 145, 146, 148,
	0, 150, 151, 152, 0, 87, 0, 0, 54, 56,
	0, 0, 0, 103, 104, 0, 107, 36, 36, 0,
	0, 0, 115, 116, 0, 119, 0, 17, 23, 0,
	0, 20, 0, 26, 31, 0, 28, 88, 93, 94,
	0, 37, 49, 0, 98, 99, 60, 63, 65, 70,
	72, 0, 0, 67, 68, 71, 121, 124, 125, 128,
	132, 133, 0, 0, 0, 154, 155, 0, 0, 77,
	82, 83, 143, 147, 149, 157, 51, 53, 57, 0,
	59, 100, 105, 106, 0, 0, 110, 111, 112, 117,
	118, 5, 18, 19, 164, 27, 96, 50, 97, 0,
	0, 75, 0, 159, 0, 69, 134, 138, 156, 153,
	140, 0, 108, 109, 73, 0, 160, 0, 0, 58,
	74, 76, 161, 0, 0, 0, 162, 0, 163,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47,
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
		//line parser.y:132
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:138
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:146
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:149
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:153
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:163
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:171
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:176
		{
			if yyVAL.loader == nil {
				yylex.Error("No loader defined")
				goto ret1
			} else {
				if sub, err := yyVAL.loader(yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				} else {
					i := &meta.Import{Module: sub}
					yyVAL.stack.Push(i)
				}
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:196
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:205
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:210
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:216
		{
			if yyVAL.loader == nil {
				yylex.Error("No loader defined")
				goto ret1
			} else {
				if sub, err := yyVAL.loader(yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				} else {
					i := &meta.Include{Module: sub}
					yyVAL.stack.Push(i)
				}
			}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:246
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:284
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:296
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:306
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:313
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:328
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:344
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:345
		{
			yyVAL.token = yyDollar[1].token
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:347
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:369
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:379
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:389
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:396
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:415
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:427
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:434
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:448
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:460
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:465
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:473
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:480
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:494
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
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
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:511
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:519
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:527
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:546
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:558
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:574
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:581
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:602
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:612
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:619
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:627
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:634
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:654
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:664
		{
			yyVAL.boolean = true
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:665
		{
			yyVAL.boolean = false
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:667
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:681
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:688
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:697
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 162:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:701
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:712
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
