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
const kywd_path = 57387
const kywd_value = 57388
const kywd_true = 57389
const kywd_false = 57390

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
	"kywd_path",
	"kywd_value",
	"kywd_true",
	"kywd_false",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:724

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 447

var yyAct = [...]int{

	153, 255, 101, 152, 157, 211, 185, 176, 171, 24,
	221, 154, 200, 160, 8, 134, 8, 140, 118, 112,
	126, 109, 156, 105, 165, 24, 161, 279, 158, 3,
	155, 6, 13, 16, 173, 11, 222, 223, 252, 256,
	13, 45, 22, 13, 113, 256, 52, 51, 37, 58,
	49, 50, 53, 13, 282, 54, 13, 56, 13, 17,
	18, 57, 55, 73, 254, 110, 173, 181, 182, 123,
	124, 13, 197, 132, 100, 131, 193, 108, 110, 114,
	110, 68, 275, 137, 147, 277, 119, 276, 128, 115,
	135, 141, 187, 162, 162, 274, 120, 169, 177, 186,
	136, 129, 144, 127, 164, 164, 262, 103, 178, 102,
	143, 166, 163, 163, 191, 210, 108, 209, 13, 107,
	67, 106, 66, 114, 261, 260, 202, 202, 190, 119,
	65, 195, 64, 115, 259, 203, 199, 128, 137, 120,
	110, 258, 257, 196, 147, 135, 207, 13, 113, 216,
	129, 141, 127, 227, 247, 136, 62, 218, 61, 13,
	224, 132, 144, 131, 246, 245, 244, 148, 206, 110,
	143, 162, 13, 229, 132, 167, 131, 99, 98, 232,
	149, 83, 164, 236, 202, 202, 280, 177, 273, 187,
	163, 242, 237, 238, 13, 243, 186, 178, 267, 265,
	264, 150, 148, 45, 146, 250, 249, 248, 52, 51,
	241, 58, 49, 50, 53, 149, 235, 54, 142, 56,
	231, 230, 228, 57, 55, 226, 217, 13, 107, 272,
	106, 198, 188, 13, 240, 239, 233, 263, 205, 204,
	87, 148, 45, 86, 85, 82, 266, 52, 51, 110,
	58, 49, 50, 53, 149, 269, 54, 13, 56, 81,
	80, 79, 57, 55, 227, 148, 45, 78, 76, 74,
	71, 52, 51, 189, 58, 49, 50, 53, 149, 215,
	54, 219, 56, 13, 270, 268, 57, 55, 225, 6,
	13, 16, 45, 11, 212, 220, 213, 52, 51, 194,
	58, 49, 50, 53, 192, 110, 54, 13, 56, 63,
	60, 59, 57, 55, 5, 281, 45, 17, 18, 21,
	72, 52, 51, 271, 58, 49, 50, 53, 234, 110,
	54, 214, 56, 45, 97, 96, 57, 55, 52, 51,
	37, 58, 49, 50, 53, 95, 13, 54, 94, 56,
	93, 92, 91, 57, 55, 45, 90, 89, 88, 84,
	52, 51, 75, 58, 49, 50, 53, 70, 69, 54,
	19, 56, 45, 43, 159, 57, 55, 52, 51, 42,
	58, 49, 50, 53, 44, 145, 54, 139, 56, 138,
	40, 133, 57, 55, 77, 39, 184, 183, 48, 180,
	179, 175, 174, 47, 122, 121, 117, 116, 25, 151,
	41, 253, 251, 208, 130, 125, 38, 172, 170, 168,
	46, 36, 35, 34, 33, 32, 31, 30, 29, 28,
	27, 26, 201, 23, 111, 15, 104, 14, 10, 9,
	7, 12, 20, 4, 2, 1, 278,
}
var yyPact = [...]int{

	2, -1000, 276, 366, 18, -1000, 306, -1000, -1000, -1000,
	-1000, 305, 147, 304, 121, 111, 69, 364, 363, 261,
	310, -1000, -1000, -1000, -1000, 260, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 358, 259, 258,
	252, 251, 250, 236, 170, 355, 235, 234, 231, 354,
	353, 352, 348, 347, 346, 344, 341, 331, 330, 167,
	166, -1000, 39, 98, -1000, 213, -1000, 29, -1000, -1000,
	-1000, -1000, -1000, -1000, 44, -1000, 57, -1000, 293, 180,
	243, 145, 145, -1000, 164, 26, 42, 332, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	222, -1000, -1000, 265, 104, -1000, 299, 64, -1000, -1000,
	294, 133, -1000, 60, -1000, -1000, 221, 44, -1000, -1000,
	-1000, 349, 349, 230, 229, 158, -1000, -1000, -1000, -1000,
	106, 289, 327, 269, -1000, -1000, -1000, -1000, 216, 180,
	-1000, -1000, 275, -1000, -1000, -1000, 290, -1000, -11, -11,
	283, 215, 243, -1000, -1000, -1000, -1000, -1000, 212, 145,
	-1000, -1000, -1000, -1000, -1000, -1000, 211, -1000, 210, -1000,
	-6, -1000, 227, 324, 206, 42, -1000, -1000, -1000, 349,
	349, 226, 225, 200, 332, -1000, -1000, -1000, -1000, 289,
	-1000, -1000, 155, 154, 153, -1000, -1000, 143, -1000, -1000,
	197, 349, -1000, 195, -1000, -1000, -1000, -1000, -1000, -1000,
	19, 131, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 130,
	123, 114, -1000, -1000, 113, 95, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 243, -1000, -1000, -1000, 190, 189, -1000,
	-1000, -1000, -1000, 98, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 188, 280, 25, 279, -1000, 319, -1000, -1000, -1000,
	-1000, -1000, -1000, 219, -1000, -1000, 178, -1000, 84, -1000,
	71, 76, -1000, -1000, -1000, -1000, -1000, -19, 176, 311,
	-1000, 43, -1000,
}
var yyPgo = [...]int{

	0, 446, 10, 5, 445, 444, 443, 442, 441, 440,
	11, 2, 314, 439, 438, 437, 436, 23, 21, 435,
	434, 19, 42, 433, 4, 12, 432, 431, 430, 429,
	428, 427, 426, 425, 424, 423, 422, 421, 420, 419,
	418, 8, 417, 3, 416, 415, 20, 26, 24, 414,
	413, 412, 411, 410, 409, 0, 30, 22, 408, 407,
	406, 18, 405, 404, 403, 402, 401, 7, 400, 399,
	398, 397, 396, 6, 395, 394, 391, 15, 390, 389,
	387, 17, 385, 384, 379, 28, 374, 13, 373, 1,
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
	81, 81, 82, 33, 83, 31, 84, 85, 86, 86,
	87, 87, 87, 87, 87, 57, 2, 2, 56, 32,
	88, 52, 52, 89, 89, 1, 18,
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
	3, 1, 3, 2, 2, 4, 2, 1, 1, 2,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 4,
	2, 1, 2, 3, 5, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -12, 13, -9, -10, -13,
	-14, 17, -8, 14, -15, -19, 15, 41, 42, 4,
	-7, -12, -22, -23, -24, -58, -27, -28, -29, -30,
	-31, -32, -33, -34, -35, -36, -37, 30, -44, -74,
	-78, -53, -84, -88, -83, 23, -38, -64, -70, 32,
	33, 29, 28, 34, 37, 44, 39, 43, 31, 5,
	5, 11, 9, 5, 11, 9, 11, 9, 12, 4,
	4, 9, 10, -22, 9, 4, 9, -75, 9, 9,
	9, 9, 9, 11, 4, 9, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 11, 11,
	-10, -11, 11, 9, -16, -17, 17, 15, -10, -18,
	36, -20, -21, 15, -10, -18, -59, -60, -61, -10,
	-18, -62, -63, 25, 26, -45, -46, -47, -10, -48,
	-49, 18, 16, -76, -77, -10, -18, -24, -79, -80,
	-81, -10, 38, -56, -57, -82, 24, -24, 22, 35,
	21, -54, -43, -55, -10, -56, -57, -24, -85, -86,
	-87, -47, -10, -56, -57, -48, -85, 11, -39, -10,
	-40, -41, -42, 40, -65, -66, -67, -10, -18, -68,
	-69, 25, 26, -71, -72, -73, -10, -24, 10, 8,
	-17, 10, 5, 12, 5, -21, 10, 12, 10, -61,
	-25, -26, -24, -25, 9, 9, 10, -46, -50, 11,
	9, -3, 5, 7, 4, 10, -77, 10, -81, 6,
	5, -2, 47, 48, -2, 5, 10, -55, 10, -87,
	10, 10, -41, 9, 4, 10, -67, -25, -25, 9,
	9, 10, -73, -3, 11, 11, 11, 11, 10, -24,
	10, -51, 19, -52, 45, -89, 20, 11, 11, 11,
	11, 11, 11, -43, 10, 10, -11, 10, 5, -89,
	5, 4, 10, 10, 11, 11, 11, 9, -1, 46,
	10, 4, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 11, 12,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 36, 34, 35, 0, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 0, 24, 0, 32, 0, 3, 15,
	26, 2, 1, 37, 92, 91, 0, 122, 0, 0,
	81, 0, 0, 143, 0, 54, 104, 116, 63, 124,
	131, 80, 146, 160, 144, 57, 103, 115, 9, 14,
	0, 6, 16, 0, 0, 18, 0, 0, 22, 23,
	0, 0, 27, 0, 30, 31, 0, 93, 94, 96,
	97, 38, 38, 0, 0, 0, 64, 66, 67, 68,
	0, 0, 0, 0, 125, 127, 128, 129, 0, 132,
	133, 135, 0, 137, 138, 139, 0, 141, 0, 0,
	0, 0, 82, 83, 85, 86, 87, 88, 0, 147,
	148, 150, 151, 152, 153, 154, 0, 89, 0, 55,
	56, 58, 0, 0, 0, 105, 106, 108, 109, 38,
	38, 0, 0, 0, 117, 118, 120, 121, 5, 0,
	19, 25, 0, 0, 0, 28, 33, 0, 90, 95,
	0, 39, 51, 0, 100, 101, 62, 65, 72, 74,
	0, 0, 69, 70, 73, 123, 126, 130, 134, 0,
	0, 0, 156, 157, 0, 0, 79, 84, 145, 149,
	159, 53, 59, 0, 61, 102, 107, 0, 0, 112,
	113, 114, 119, 0, 20, 21, 166, 29, 98, 52,
	99, 0, 0, 77, 0, 161, 0, 71, 136, 140,
	158, 155, 142, 0, 110, 111, 0, 75, 0, 162,
	0, 0, 60, 17, 76, 78, 163, 0, 0, 0,
	164, 0, 165,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48,
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
		//line parser.y:133
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:139
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:147
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:150
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:154
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:164
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:172
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:177
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
		//line parser.y:201
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:210
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:215
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:221
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
		//line parser.y:246
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:251
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:289
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:301
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:311
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:318
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:326
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:333
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:349
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:350
		{
			yyVAL.token = yyDollar[1].token
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:352
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
		//line parser.y:364
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:374
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
		//line parser.y:384
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:394
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:401
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:420
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:432
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:439
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:465
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:470
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:478
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:485
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:499
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:504
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:511
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:516
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:524
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:532
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:551
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:563
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:579
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:586
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:607
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
		//line parser.y:617
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:624
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:632
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:639
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:659
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:669
		{
			yyVAL.boolean = true
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:670
		{
			yyVAL.boolean = false
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:672
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:686
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:693
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:702
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 164:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:706
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:717
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
