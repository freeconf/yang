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
const kywd_true = 57387
const kywd_false = 57388

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
	"kywd_true",
	"kywd_false",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:711

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 451

var yyAct = [...]int{

	151, 261, 150, 155, 183, 174, 169, 158, 24, 222,
	152, 200, 138, 8, 132, 8, 124, 116, 110, 103,
	107, 283, 154, 163, 24, 156, 6, 13, 16, 159,
	11, 223, 224, 153, 258, 262, 45, 13, 171, 3,
	262, 52, 51, 37, 58, 49, 50, 53, 13, 195,
	54, 22, 56, 190, 17, 18, 57, 55, 13, 69,
	260, 188, 281, 171, 280, 13, 105, 206, 104, 179,
	180, 13, 74, 130, 101, 129, 211, 106, 210, 112,
	108, 286, 279, 135, 145, 278, 117, 108, 126, 113,
	133, 139, 185, 160, 160, 268, 118, 167, 175, 184,
	134, 127, 68, 142, 67, 162, 162, 125, 176, 164,
	66, 267, 65, 106, 141, 266, 161, 161, 265, 63,
	112, 62, 187, 202, 202, 264, 117, 13, 193, 263,
	113, 253, 203, 198, 126, 135, 118, 252, 121, 122,
	207, 145, 133, 251, 250, 248, 215, 127, 139, 108,
	218, 228, 134, 125, 241, 13, 105, 225, 104, 13,
	142, 130, 13, 129, 130, 231, 129, 235, 160, 194,
	146, 141, 232, 13, 111, 236, 229, 108, 219, 240,
	162, 202, 202, 147, 175, 216, 185, 247, 208, 242,
	243, 161, 13, 184, 176, 108, 199, 13, 111, 148,
	146, 45, 144, 196, 191, 255, 52, 51, 186, 58,
	49, 50, 53, 147, 277, 54, 140, 56, 13, 108,
	165, 57, 55, 100, 99, 84, 146, 45, 60, 284,
	272, 271, 52, 51, 270, 58, 49, 50, 53, 147,
	269, 54, 256, 56, 254, 249, 246, 57, 55, 239,
	285, 234, 233, 230, 13, 227, 217, 197, 245, 244,
	237, 274, 146, 45, 205, 204, 88, 87, 52, 51,
	228, 58, 49, 50, 53, 147, 214, 54, 86, 56,
	13, 83, 220, 57, 55, 82, 6, 13, 16, 45,
	11, 81, 80, 79, 52, 51, 77, 58, 49, 50,
	53, 75, 108, 54, 13, 56, 72, 275, 273, 57,
	55, 226, 221, 45, 17, 18, 212, 73, 52, 51,
	192, 58, 49, 50, 53, 189, 108, 54, 64, 56,
	45, 61, 59, 57, 55, 52, 51, 37, 58, 49,
	50, 53, 5, 13, 54, 276, 56, 21, 238, 213,
	57, 55, 45, 98, 97, 96, 95, 52, 51, 94,
	58, 49, 50, 53, 93, 92, 54, 91, 56, 45,
	90, 89, 57, 55, 52, 51, 85, 58, 49, 50,
	53, 76, 71, 54, 70, 56, 19, 43, 157, 57,
	55, 42, 44, 143, 137, 136, 40, 131, 78, 39,
	182, 181, 48, 178, 177, 173, 172, 47, 120, 119,
	115, 114, 25, 149, 41, 259, 257, 209, 128, 123,
	38, 170, 168, 166, 46, 36, 35, 34, 33, 32,
	31, 30, 29, 28, 27, 26, 201, 23, 109, 15,
	102, 14, 10, 9, 7, 12, 20, 4, 2, 1,
	282,
}
var yyPact = [...]int{

	14, -1000, 275, 382, 15, -1000, 327, -1000, 219, -1000,
	-1000, 326, 112, 323, 103, 95, 49, 380, 378, 299,
	309, -1000, -1000, -1000, -1000, 294, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 377, 289, 286,
	285, 284, 278, 274, 216, 372, 271, 260, 259, 367,
	366, 363, 361, 360, 355, 352, 351, 350, 349, 215,
	-1000, 214, -1000, 36, -1000, -1000, 143, -1000, 185, -1000,
	-1000, -1000, -1000, -1000, -1000, 115, -1000, 147, -1000, 292,
	180, 242, 150, 150, -1000, 211, 25, 46, 331, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 199, 53, -1000, 320, 43, 195, -1000, 315, 161,
	-1000, 39, 194, -1000, 249, 115, -1000, 187, -1000, 348,
	348, 258, 257, 59, -1000, -1000, 179, -1000, 69, 311,
	345, 268, -1000, 176, -1000, -1000, 248, 180, -1000, 169,
	276, -1000, -1000, -1000, 307, -1000, -14, -14, 306, 247,
	242, -1000, 167, -1000, -1000, -1000, 245, 150, -1000, -1000,
	163, -1000, -1000, -1000, 244, -1000, 243, 158, 0, -1000,
	253, 344, 241, 46, -1000, 145, -1000, 348, 348, 252,
	251, 238, 331, -1000, 136, -1000, 237, -1000, -1000, 135,
	134, -1000, 128, -1000, -1000, 122, -1000, -1000, -1000, -1000,
	236, 348, -1000, 234, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 17, 120, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	116, 109, 106, -1000, -1000, 102, 86, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 242, -1000, -1000,
	-1000, -1000, 226, 223, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 222, 303, 22,
	302, -1000, 341, -1000, -1000, -1000, -1000, -1000, -1000, 206,
	-1000, -1000, -1000, 76, -1000, 73, 55, -1000, -1000, -1000,
	-1000, -23, 221, 246, -1000, 72, -1000,
}
var yyPgo = [...]int{

	0, 450, 9, 449, 448, 447, 446, 445, 444, 10,
	342, 443, 442, 441, 440, 19, 20, 439, 438, 18,
	51, 437, 3, 11, 436, 435, 434, 433, 432, 431,
	430, 429, 428, 427, 426, 425, 424, 423, 422, 6,
	421, 2, 420, 419, 16, 29, 23, 418, 417, 416,
	415, 414, 413, 0, 33, 22, 412, 411, 410, 17,
	409, 408, 407, 406, 405, 5, 404, 403, 402, 401,
	400, 4, 399, 398, 397, 14, 396, 395, 394, 12,
	393, 392, 391, 25, 388, 7, 387, 1,
}
var yyR1 = [...]int{

	0, 3, 4, 7, 8, 8, 9, 5, 5, 10,
	10, 10, 10, 10, 10, 13, 14, 14, 15, 15,
	15, 15, 11, 11, 17, 18, 18, 19, 19, 19,
	12, 12, 20, 20, 6, 6, 23, 23, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 24,
	24, 33, 37, 37, 37, 36, 38, 38, 39, 40,
	25, 42, 43, 43, 44, 44, 44, 46, 45, 47,
	48, 48, 49, 49, 49, 28, 51, 52, 52, 41,
	41, 53, 53, 53, 53, 32, 21, 56, 57, 57,
	58, 58, 59, 59, 59, 59, 60, 61, 34, 62,
	63, 63, 64, 64, 65, 65, 65, 65, 66, 67,
	35, 68, 69, 69, 70, 70, 71, 71, 26, 73,
	72, 74, 74, 75, 75, 75, 27, 76, 77, 78,
	78, 79, 79, 79, 79, 79, 79, 79, 80, 31,
	81, 29, 82, 83, 84, 84, 85, 85, 85, 85,
	85, 55, 2, 2, 54, 30, 86, 50, 50, 87,
	87, 1, 16,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 5, 2, 1, 2, 3,
	1, 2, 1, 1, 3, 2, 1, 2, 3, 3,
	2, 1, 2, 4, 2, 1, 2, 3, 2, 1,
	2, 4, 1, 1, 1, 2, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 4, 0, 2, 1, 2, 1, 2, 4, 2,
	4, 2, 1, 2, 1, 2, 1, 3, 2, 2,
	1, 3, 3, 1, 3, 4, 2, 0, 1, 1,
	2, 2, 1, 1, 1, 3, 4, 2, 0, 1,
	1, 2, 2, 1, 3, 3, 2, 2, 4, 2,
	0, 1, 1, 2, 2, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 2, 1, 2, 3,
	2, 1, 2, 2, 1, 1, 4, 2, 1, 1,
	2, 2, 3, 1, 1, 1, 3, 1, 3, 2,
	2, 4, 2, 1, 1, 2, 1, 2, 1, 1,
	1, 3, 1, 1, 3, 4, 2, 1, 2, 3,
	5, 3, 3,
}
var yyChk = [...]int{

	-1000, -3, -4, 25, -5, -10, 11, -8, -9, -11,
	-12, 15, -7, 12, -13, -17, 13, 39, 40, 4,
	-6, -10, -20, -21, -22, -56, -25, -26, -27, -28,
	-29, -30, -31, -32, -33, -34, -35, 28, -42, -72,
	-76, -51, -82, -86, -81, 21, -36, -62, -68, 30,
	31, 27, 26, 32, 35, 42, 37, 41, 29, 5,
	9, 5, 9, 7, 5, 9, 7, 9, 7, 10,
	4, 4, 7, 8, -20, 7, 4, 7, -73, 7,
	7, 7, 7, 7, 9, 4, 7, 7, 7, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 9,
	9, -9, -14, -15, 15, 13, -9, -16, 34, -18,
	-19, 13, -9, -16, -57, -58, -59, -9, -16, -60,
	-61, 23, 24, -43, -44, -45, -9, -46, -47, 16,
	14, -74, -75, -9, -16, -22, -77, -78, -79, -9,
	36, -54, -55, -80, 22, -22, 20, 33, 19, -52,
	-41, -53, -9, -54, -55, -22, -83, -84, -85, -45,
	-9, -54, -55, -46, -83, 9, -37, -9, -38, -39,
	-40, 38, -63, -64, -65, -9, -16, -66, -67, 23,
	24, -69, -70, -71, -9, -22, 9, -15, 8, 5,
	10, 9, 5, -19, 8, 10, 9, 8, -59, 9,
	-23, -24, -22, -23, 7, 7, 8, -44, 9, -48,
	9, 7, 5, 4, 8, -75, 9, 8, -79, 9,
	6, 5, -2, 45, 46, -2, 5, 8, -53, 9,
	8, -85, 9, 8, 8, 9, -39, 7, 4, 8,
	-65, 9, -23, -23, 7, 7, 8, -71, 9, 8,
	9, 9, 9, 9, 8, -22, 8, -49, 17, -50,
	43, -87, 18, 9, 9, 9, 9, 9, 9, -41,
	8, 8, 8, 5, -87, 5, 4, 8, 9, 9,
	9, 7, -1, 44, 8, 4, 9,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 0, 12,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 34, 32, 33, 0, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 0, 4, 0, 6, 22, 0, 30, 0, 3,
	15, 24, 2, 1, 35, 88, 87, 0, 118, 0,
	0, 77, 0, 0, 139, 0, 52, 100, 112, 61,
	120, 127, 76, 142, 156, 140, 55, 99, 111, 9,
	14, 0, 0, 16, 0, 0, 0, 21, 0, 0,
	25, 0, 0, 29, 0, 89, 90, 0, 93, 36,
	36, 0, 0, 0, 62, 64, 0, 66, 0, 0,
	0, 0, 121, 0, 124, 125, 0, 128, 129, 0,
	0, 133, 134, 135, 0, 137, 0, 0, 0, 0,
	78, 79, 0, 82, 83, 84, 0, 143, 144, 146,
	0, 148, 149, 150, 0, 85, 0, 0, 54, 56,
	0, 0, 0, 101, 102, 0, 105, 36, 36, 0,
	0, 0, 113, 114, 0, 117, 0, 17, 23, 0,
	0, 20, 0, 26, 31, 0, 28, 86, 91, 92,
	0, 37, 49, 0, 96, 97, 60, 63, 65, 68,
	70, 0, 0, 69, 119, 122, 123, 126, 130, 131,
	0, 0, 0, 152, 153, 0, 0, 75, 80, 81,
	141, 145, 147, 155, 51, 53, 57, 0, 59, 98,
	103, 104, 0, 0, 108, 109, 110, 115, 116, 5,
	18, 19, 162, 27, 94, 50, 95, 0, 0, 73,
	0, 157, 0, 67, 132, 136, 154, 151, 138, 0,
	106, 107, 71, 0, 158, 0, 0, 58, 72, 74,
	159, 0, 0, 0, 160, 0, 161,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
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
		//line parser.y:129
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:135
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:143
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:146
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:150
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:160
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:168
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:173
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
		//line parser.y:193
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:202
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:207
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:213
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
		//line parser.y:238
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:243
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:281
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:293
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:303
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:310
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:318
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:325
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:339
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(tokenString(yyDollar[2].token))
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:351
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:361
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:371
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:381
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:388
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:407
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:419
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:426
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:440
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:445
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:452
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:457
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:465
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:472
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:486
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:491
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:498
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:503
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:511
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:519
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:538
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:550
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:566
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:573
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:594
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:604
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:611
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:619
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:626
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:646
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:656
		{
			yyVAL.boolean = true
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:657
		{
			yyVAL.boolean = false
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:659
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:673
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:680
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:689
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 160:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:693
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:704
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
