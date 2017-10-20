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
const token_custom = 57350
const token_curly_open = 57351
const token_curly_close = 57352
const token_semi = 57353
const token_rev_ident = 57354
const kywd_namespace = 57355
const kywd_description = 57356
const kywd_revision = 57357
const kywd_type = 57358
const kywd_prefix = 57359
const kywd_default = 57360
const kywd_length = 57361
const kywd_enum = 57362
const kywd_key = 57363
const kywd_config = 57364
const kywd_uses = 57365
const kywd_unique = 57366
const kywd_input = 57367
const kywd_output = 57368
const kywd_module = 57369
const kywd_container = 57370
const kywd_list = 57371
const kywd_rpc = 57372
const kywd_notification = 57373
const kywd_typedef = 57374
const kywd_grouping = 57375
const kywd_leaf = 57376
const kywd_mandatory = 57377
const kywd_reference = 57378
const kywd_leaf_list = 57379
const kywd_max_elements = 57380
const kywd_choice = 57381
const kywd_case = 57382
const kywd_import = 57383
const kywd_include = 57384
const kywd_action = 57385
const kywd_anyxml = 57386
const kywd_anydata = 57387
const kywd_path = 57388
const kywd_value = 57389
const kywd_true = 57390
const kywd_false = 57391

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_int",
	"token_number",
	"token_custom",
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
	"kywd_anydata",
	"kywd_path",
	"kywd_value",
	"kywd_true",
	"kywd_false",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:728

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 457

var yyAct = [...]int{

	155, 257, 103, 154, 159, 213, 187, 178, 173, 24,
	223, 156, 202, 162, 8, 136, 8, 142, 120, 114,
	128, 111, 158, 107, 167, 24, 163, 281, 160, 175,
	157, 6, 13, 16, 3, 11, 224, 225, 254, 258,
	13, 45, 13, 258, 13, 199, 52, 51, 37, 59,
	49, 50, 53, 183, 184, 54, 22, 57, 195, 17,
	18, 58, 55, 56, 112, 256, 175, 13, 208, 134,
	284, 133, 13, 69, 134, 102, 133, 74, 110, 279,
	116, 278, 277, 276, 139, 149, 105, 121, 104, 130,
	117, 137, 143, 189, 164, 164, 264, 122, 171, 179,
	188, 138, 131, 146, 129, 166, 166, 242, 263, 180,
	13, 145, 168, 165, 165, 6, 13, 16, 110, 11,
	262, 125, 126, 261, 212, 116, 211, 260, 204, 204,
	192, 121, 112, 197, 68, 117, 67, 205, 201, 130,
	139, 122, 259, 17, 18, 198, 149, 137, 209, 13,
	115, 218, 131, 143, 129, 229, 193, 138, 249, 220,
	13, 109, 226, 108, 146, 248, 13, 109, 66, 108,
	65, 112, 145, 164, 63, 231, 62, 247, 13, 115,
	246, 234, 112, 169, 166, 238, 204, 204, 112, 179,
	101, 189, 165, 244, 239, 240, 13, 245, 188, 180,
	112, 100, 84, 152, 150, 45, 148, 282, 251, 275,
	52, 51, 269, 59, 49, 50, 53, 151, 267, 54,
	144, 57, 266, 252, 241, 58, 55, 56, 13, 250,
	134, 243, 133, 237, 233, 232, 150, 230, 228, 265,
	274, 219, 200, 190, 13, 235, 207, 206, 268, 151,
	88, 87, 150, 45, 86, 83, 82, 271, 52, 51,
	81, 59, 49, 50, 53, 151, 229, 54, 13, 57,
	80, 79, 77, 58, 55, 56, 150, 45, 75, 72,
	191, 221, 52, 51, 272, 59, 49, 50, 53, 151,
	217, 54, 270, 57, 13, 227, 222, 58, 55, 56,
	214, 196, 215, 45, 194, 64, 61, 60, 52, 51,
	5, 59, 49, 50, 53, 21, 112, 54, 13, 57,
	283, 273, 236, 58, 55, 56, 216, 45, 99, 98,
	97, 96, 52, 51, 95, 59, 49, 50, 53, 73,
	112, 54, 94, 57, 93, 92, 91, 58, 55, 56,
	90, 89, 45, 85, 76, 71, 70, 52, 51, 37,
	59, 49, 50, 53, 19, 43, 54, 13, 57, 161,
	42, 44, 58, 55, 56, 147, 45, 141, 140, 40,
	135, 52, 51, 78, 59, 49, 50, 53, 39, 186,
	54, 185, 57, 48, 45, 182, 58, 55, 56, 52,
	51, 181, 59, 49, 50, 53, 177, 176, 54, 47,
	57, 124, 123, 119, 58, 55, 56, 118, 25, 153,
	41, 255, 253, 210, 132, 127, 38, 174, 172, 170,
	46, 36, 35, 34, 33, 32, 31, 30, 29, 28,
	27, 26, 203, 23, 113, 15, 106, 14, 10, 9,
	7, 12, 20, 4, 2, 1, 280,
}
var yyPact = [...]int{

	7, -1000, 102, 360, 18, -1000, 302, -1000, -1000, -1000,
	-1000, 301, 165, 300, 159, 125, 61, 352, 351, 270,
	329, -1000, -1000, -1000, -1000, 269, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 350, 263, 262,
	261, 251, 247, 246, 191, 349, 245, 242, 241, 347,
	346, 342, 341, 340, 338, 330, 327, 326, 325, 324,
	190, 179, -1000, 30, 77, -1000, 152, -1000, 164, -1000,
	-1000, -1000, -1000, -1000, -1000, 96, -1000, 53, -1000, 304,
	182, 254, 214, 214, -1000, 172, 26, 28, 353, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 233, -1000, -1000, 272, 146, -1000, 299, 46,
	-1000, -1000, 296, 135, -1000, 33, -1000, -1000, 232, 96,
	-1000, -1000, -1000, 371, 371, 238, 237, 58, -1000, -1000,
	-1000, -1000, 115, 295, 322, 280, -1000, -1000, -1000, -1000,
	231, 182, -1000, -1000, 275, -1000, -1000, -1000, 291, -1000,
	-12, -12, 290, 228, 254, -1000, -1000, -1000, -1000, -1000,
	227, 214, -1000, -1000, -1000, -1000, -1000, -1000, 225, -1000,
	224, -1000, -11, -1000, 236, 318, 223, 28, -1000, -1000,
	-1000, 371, 371, 215, 98, 221, 353, -1000, -1000, -1000,
	-1000, 295, -1000, -1000, 169, 166, 154, -1000, -1000, 147,
	-1000, -1000, 219, 371, -1000, 213, -1000, -1000, -1000, -1000,
	-1000, -1000, 19, 131, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 116, 112, 109, -1000, -1000, 97, 85, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 254, -1000, -1000, -1000, 212,
	208, -1000, -1000, -1000, -1000, 77, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 202, 287, 23, 279, -1000, 317, -1000,
	-1000, -1000, -1000, -1000, -1000, 230, -1000, -1000, 199, -1000,
	72, -1000, 71, 70, -1000, -1000, -1000, -1000, -1000, -20,
	197, 316, -1000, 59, -1000,
}
var yyPgo = [...]int{

	0, 456, 10, 5, 455, 454, 453, 452, 451, 450,
	11, 2, 310, 449, 448, 447, 446, 23, 21, 445,
	444, 19, 56, 443, 4, 12, 442, 441, 440, 439,
	438, 437, 436, 435, 434, 433, 432, 431, 430, 429,
	428, 8, 427, 3, 426, 425, 20, 26, 24, 424,
	423, 422, 421, 420, 419, 0, 30, 22, 418, 417,
	413, 18, 412, 411, 409, 407, 406, 7, 401, 395,
	393, 391, 389, 6, 388, 383, 380, 15, 379, 378,
	377, 17, 375, 371, 370, 28, 369, 13, 365, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 6, 6, 12,
	12, 12, 12, 12, 12, 15, 11, 11, 16, 16,
	17, 17, 17, 17, 13, 13, 19, 20, 20, 21,
	21, 21, 14, 14, 22, 22, 7, 7, 25, 25,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 26, 26, 35, 39, 39, 39, 38, 40, 40,
	41, 42, 27, 44, 45, 45, 46, 46, 46, 3,
	3, 48, 47, 49, 50, 50, 51, 51, 51, 30,
	53, 54, 54, 43, 43, 55, 55, 55, 55, 34,
	23, 58, 59, 59, 60, 60, 61, 61, 61, 61,
	62, 63, 36, 64, 65, 65, 66, 66, 67, 67,
	67, 67, 68, 69, 37, 70, 71, 71, 72, 72,
	73, 73, 28, 75, 74, 76, 76, 77, 77, 77,
	29, 78, 79, 80, 80, 81, 81, 81, 81, 81,
	81, 81, 82, 33, 83, 83, 31, 84, 85, 86,
	86, 87, 87, 87, 87, 87, 57, 2, 2, 56,
	32, 88, 52, 52, 89, 89, 1, 18,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 3, 1, 2, 3,
	1, 1, 1, 1, 3, 2, 1, 5, 1, 2,
	3, 3, 1, 1, 2, 4, 2, 1, 2, 3,
	1, 1, 2, 4, 1, 1, 1, 2, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 4, 0, 1, 1, 2, 1, 2,
	4, 2, 4, 2, 1, 2, 1, 1, 1, 1,
	1, 3, 2, 2, 1, 3, 3, 1, 3, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 1, 3,
	4, 2, 0, 1, 1, 2, 1, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	3, 3, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 2, 3, 2, 1, 2, 1, 1, 1,
	4, 2, 1, 1, 2, 1, 3, 1, 1, 1,
	3, 1, 3, 2, 2, 2, 4, 2, 1, 1,
	2, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	4, 2, 1, 2, 3, 5, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -12, 13, -9, -10, -13,
	-14, 17, -8, 14, -15, -19, 15, 41, 42, 4,
	-7, -12, -22, -23, -24, -58, -27, -28, -29, -30,
	-31, -32, -33, -34, -35, -36, -37, 30, -44, -74,
	-78, -53, -84, -88, -83, 23, -38, -64, -70, 32,
	33, 29, 28, 34, 37, 44, 45, 39, 43, 31,
	5, 5, 11, 9, 5, 11, 9, 11, 9, 12,
	4, 4, 9, 10, -22, 9, 4, 9, -75, 9,
	9, 9, 9, 9, 11, 4, 9, 9, 9, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	11, 11, -10, -11, 11, 9, -16, -17, 17, 15,
	-10, -18, 36, -20, -21, 15, -10, -18, -59, -60,
	-61, -10, -18, -62, -63, 25, 26, -45, -46, -47,
	-10, -48, -49, 18, 16, -76, -77, -10, -18, -24,
	-79, -80, -81, -10, 38, -56, -57, -82, 24, -24,
	22, 35, 21, -54, -43, -55, -10, -56, -57, -24,
	-85, -86, -87, -47, -10, -56, -57, -48, -85, 11,
	-39, -10, -40, -41, -42, 40, -65, -66, -67, -10,
	-18, -68, -69, 25, 26, -71, -72, -73, -10, -24,
	10, 8, -17, 10, 5, 12, 5, -21, 10, 12,
	10, -61, -25, -26, -24, -25, 9, 9, 10, -46,
	-50, 11, 9, -3, 5, 7, 4, 10, -77, 10,
	-81, 6, 5, -2, 48, 49, -2, 5, 10, -55,
	10, -87, 10, 10, -41, 9, 4, 10, -67, -25,
	-25, 9, 9, 10, -73, -3, 11, 11, 11, 11,
	10, -24, 10, -51, 19, -52, 46, -89, 20, 11,
	11, 11, 11, 11, 11, -43, 10, 10, -11, 10,
	5, -89, 5, 4, 10, 10, 11, 11, 11, 9,
	-1, 47, 10, 4, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 11, 12,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 36, 34, 35, 0, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 4, 0, 0, 24, 0, 32, 0, 3,
	15, 26, 2, 1, 37, 92, 91, 0, 122, 0,
	0, 81, 0, 0, 143, 0, 54, 104, 116, 63,
	124, 131, 80, 147, 161, 144, 145, 57, 103, 115,
	9, 14, 0, 6, 16, 0, 0, 18, 0, 0,
	22, 23, 0, 0, 27, 0, 30, 31, 0, 93,
	94, 96, 97, 38, 38, 0, 0, 0, 64, 66,
	67, 68, 0, 0, 0, 0, 125, 127, 128, 129,
	0, 132, 133, 135, 0, 137, 138, 139, 0, 141,
	0, 0, 0, 0, 82, 83, 85, 86, 87, 88,
	0, 148, 149, 151, 152, 153, 154, 155, 0, 89,
	0, 55, 56, 58, 0, 0, 0, 105, 106, 108,
	109, 38, 38, 0, 0, 0, 117, 118, 120, 121,
	5, 0, 19, 25, 0, 0, 0, 28, 33, 0,
	90, 95, 0, 39, 51, 0, 100, 101, 62, 65,
	72, 74, 0, 0, 69, 70, 73, 123, 126, 130,
	134, 0, 0, 0, 157, 158, 0, 0, 79, 84,
	146, 150, 160, 53, 59, 0, 61, 102, 107, 0,
	0, 112, 113, 114, 119, 0, 20, 21, 167, 29,
	98, 52, 99, 0, 0, 77, 0, 162, 0, 71,
	136, 140, 159, 156, 142, 0, 110, 111, 0, 75,
	0, 163, 0, 0, 60, 17, 76, 78, 164, 0,
	0, 0, 165, 0, 166,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49,
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
		//line parser.y:134
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:140
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:148
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:151
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:155
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:165
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:173
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:178
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
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:202
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:211
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:216
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:222
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
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:247
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:252
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:290
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:302
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:312
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:319
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:327
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:334
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:350
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:351
		{
			yyVAL.token = yyDollar[1].token
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:353
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:365
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:375
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:395
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:402
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:421
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:433
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:440
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:454
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:459
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:466
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:479
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:486
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:500
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:505
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:512
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:517
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:525
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:533
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:552
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:564
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:580
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:587
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:608
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:618
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:625
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:628
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:636
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:643
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:663
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:673
		{
			yyVAL.boolean = true
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:674
		{
			yyVAL.boolean = false
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:676
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:690
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:697
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:706
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:710
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:721
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
