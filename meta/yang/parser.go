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
	yys    int
	ident  string
	token  string
	stack  *yangMetaStack
	loader ModuleLoader
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

//line parser.y:703

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 449

var yyAct = [...]int{

	151, 259, 150, 155, 183, 174, 169, 158, 24, 138,
	152, 200, 132, 8, 124, 8, 103, 116, 110, 281,
	107, 171, 154, 3, 24, 156, 163, 256, 260, 159,
	6, 13, 16, 153, 11, 6, 13, 16, 260, 11,
	45, 13, 13, 195, 190, 52, 51, 37, 58, 49,
	50, 53, 22, 258, 54, 279, 56, 278, 17, 18,
	57, 55, 69, 17, 18, 13, 105, 171, 104, 211,
	68, 210, 67, 74, 101, 284, 194, 106, 277, 112,
	13, 111, 276, 135, 145, 266, 117, 108, 126, 113,
	133, 139, 185, 160, 160, 265, 118, 167, 175, 184,
	134, 264, 108, 142, 127, 162, 162, 125, 176, 164,
	66, 263, 65, 106, 141, 262, 161, 161, 282, 187,
	112, 13, 111, 202, 202, 13, 117, 130, 193, 129,
	113, 261, 203, 198, 126, 135, 118, 63, 207, 62,
	251, 145, 133, 108, 215, 250, 249, 218, 139, 248,
	127, 226, 134, 125, 246, 13, 239, 188, 243, 233,
	142, 13, 105, 230, 104, 229, 179, 180, 160, 227,
	206, 141, 219, 216, 13, 234, 130, 108, 129, 238,
	162, 202, 202, 108, 175, 208, 185, 245, 199, 240,
	241, 161, 13, 184, 176, 196, 191, 186, 165, 148,
	146, 45, 144, 100, 99, 253, 52, 51, 13, 58,
	49, 50, 53, 147, 84, 54, 140, 56, 60, 121,
	122, 57, 55, 275, 270, 269, 268, 13, 254, 252,
	108, 247, 244, 237, 232, 146, 45, 231, 267, 228,
	225, 52, 51, 217, 58, 49, 50, 53, 147, 197,
	54, 242, 56, 273, 235, 13, 57, 55, 13, 272,
	130, 205, 129, 146, 45, 204, 146, 88, 226, 52,
	51, 87, 58, 49, 50, 53, 147, 214, 54, 147,
	56, 13, 86, 83, 57, 55, 82, 81, 80, 79,
	45, 77, 75, 72, 220, 52, 51, 271, 58, 49,
	50, 53, 224, 108, 54, 13, 56, 223, 222, 221,
	57, 55, 212, 192, 45, 43, 189, 64, 73, 52,
	51, 61, 58, 49, 50, 53, 59, 108, 54, 283,
	56, 45, 274, 236, 57, 55, 52, 51, 37, 58,
	49, 50, 53, 5, 13, 54, 213, 56, 21, 98,
	97, 57, 55, 45, 96, 95, 94, 93, 52, 51,
	92, 58, 49, 50, 53, 91, 90, 54, 89, 56,
	45, 85, 76, 57, 55, 52, 51, 71, 58, 49,
	50, 53, 70, 19, 54, 157, 56, 42, 44, 143,
	57, 55, 137, 136, 40, 131, 78, 39, 182, 181,
	48, 178, 177, 173, 172, 47, 120, 119, 115, 114,
	25, 149, 41, 257, 255, 209, 128, 123, 38, 170,
	168, 166, 46, 36, 35, 34, 33, 32, 31, 30,
	29, 28, 27, 26, 201, 23, 109, 15, 102, 14,
	10, 9, 7, 12, 20, 4, 2, 1, 280,
}
var yyPact = [...]int{

	-2, -1000, 24, 379, 19, -1000, 321, -1000, 209, -1000,
	-1000, 316, 130, 312, 103, 63, 52, 378, 373, 286,
	310, -1000, -1000, -1000, -1000, 285, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 368, 284, 282,
	281, 280, 279, 276, 205, 367, 275, 264, 260, 364,
	362, 361, 356, 353, 352, 351, 350, 346, 345, 195,
	-1000, 194, -1000, 30, -1000, -1000, 53, -1000, 109, -1000,
	-1000, -1000, -1000, -1000, -1000, 196, -1000, 113, -1000, 293,
	180, 243, 246, 246, -1000, 189, 29, 143, 332, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 188, 149, -1000, 311, 34, 187, -1000, 308, 68,
	-1000, 33, 186, -1000, 241, 196, -1000, 179, -1000, 349,
	349, 258, 254, 162, -1000, -1000, 176, -1000, 62, 307,
	342, 269, -1000, 164, -1000, -1000, 235, 180, -1000, 163,
	288, -1000, -1000, -1000, 304, -1000, 303, 302, 297, 232,
	243, -1000, 160, -1000, -1000, -1000, 231, 246, -1000, -1000,
	154, -1000, -1000, -1000, 229, -1000, 226, 150, -17, -1000,
	247, 329, 225, 143, -1000, 147, -1000, 349, 349, 244,
	151, 224, 332, -1000, 145, -1000, 223, -1000, -1000, 140,
	137, -1000, 136, -1000, -1000, 131, -1000, -1000, -1000, -1000,
	221, 349, -1000, 220, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 10, 122, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	106, 102, 92, 86, 76, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 243, -1000, -1000, -1000, -1000,
	218, 217, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 216, 292, 20, 248, -1000,
	328, -1000, -1000, -1000, -1000, -1000, -1000, 215, -1000, -1000,
	-1000, 73, -1000, 69, 48, -1000, -1000, -1000, -1000, -25,
	110, 325, -1000, 66, -1000,
}
var yyPgo = [...]int{

	0, 448, 447, 446, 445, 444, 443, 442, 10, 343,
	441, 440, 439, 438, 16, 20, 437, 436, 18, 52,
	435, 3, 11, 434, 433, 432, 431, 430, 429, 428,
	427, 426, 425, 424, 423, 422, 421, 420, 6, 419,
	2, 418, 417, 14, 29, 26, 416, 415, 414, 413,
	412, 411, 0, 33, 22, 410, 409, 408, 17, 407,
	406, 405, 404, 403, 5, 402, 401, 400, 399, 398,
	4, 397, 396, 395, 12, 394, 393, 392, 9, 389,
	388, 387, 25, 385, 7, 315, 1,
}
var yyR1 = [...]int{

	0, 2, 3, 6, 7, 7, 8, 4, 4, 9,
	9, 9, 9, 9, 9, 12, 13, 13, 14, 14,
	14, 14, 10, 10, 16, 17, 17, 18, 18, 18,
	11, 11, 19, 19, 5, 5, 22, 22, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 23,
	23, 32, 36, 36, 36, 35, 37, 37, 38, 39,
	24, 41, 42, 42, 43, 43, 43, 45, 44, 46,
	47, 47, 48, 48, 48, 27, 50, 51, 51, 40,
	40, 52, 52, 52, 52, 31, 20, 55, 56, 56,
	57, 57, 58, 58, 58, 58, 59, 60, 33, 61,
	62, 62, 63, 63, 64, 64, 64, 64, 65, 66,
	34, 67, 68, 68, 69, 69, 70, 70, 25, 72,
	71, 73, 73, 74, 74, 74, 26, 75, 76, 77,
	77, 78, 78, 78, 78, 78, 78, 78, 79, 30,
	80, 28, 81, 82, 83, 83, 84, 84, 84, 84,
	84, 54, 53, 29, 85, 49, 49, 86, 86, 1,
	15,
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
	1, 3, 3, 4, 2, 1, 2, 3, 5, 3,
	3,
}
var yyChk = [...]int{

	-1000, -2, -3, 25, -4, -9, 11, -7, -8, -10,
	-11, 15, -6, 12, -12, -16, 13, 39, 40, 4,
	-5, -9, -19, -20, -21, -55, -24, -25, -26, -27,
	-28, -29, -30, -31, -32, -33, -34, 28, -41, -71,
	-75, -50, -81, -85, -80, 21, -35, -61, -67, 30,
	31, 27, 26, 32, 35, 42, 37, 41, 29, 5,
	9, 5, 9, 7, 5, 9, 7, 9, 7, 10,
	4, 4, 7, 8, -19, 7, 4, 7, -72, 7,
	7, 7, 7, 7, 9, 4, 7, 7, 7, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 9,
	9, -8, -13, -14, 15, 13, -8, -15, 34, -17,
	-18, 13, -8, -15, -56, -57, -58, -8, -15, -59,
	-60, 23, 24, -42, -43, -44, -8, -45, -46, 16,
	14, -73, -74, -8, -15, -21, -76, -77, -78, -8,
	36, -53, -54, -79, 22, -21, 20, 33, 19, -51,
	-40, -52, -8, -53, -54, -21, -82, -83, -84, -44,
	-8, -53, -54, -45, -82, 9, -36, -8, -37, -38,
	-39, 38, -62, -63, -64, -8, -15, -65, -66, 23,
	24, -68, -69, -70, -8, -21, 9, -14, 8, 5,
	10, 9, 5, -18, 8, 10, 9, 8, -58, 9,
	-22, -23, -21, -22, 7, 7, 8, -43, 9, -47,
	9, 7, 5, 4, 8, -74, 9, 8, -78, 9,
	6, 5, 5, 5, 5, 8, -52, 9, 8, -84,
	9, 8, 8, 9, -38, 7, 4, 8, -64, 9,
	-22, -22, 7, 7, 8, -70, 9, 8, 9, 9,
	9, 9, 8, -21, 8, -48, 17, -49, 43, -86,
	18, 9, 9, 9, 9, 9, 9, -40, 8, 8,
	8, 5, -86, 5, 4, 8, 9, 9, 9, 7,
	-1, 44, 8, 4, 9,
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
	120, 127, 76, 142, 154, 140, 55, 99, 111, 9,
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
	0, 0, 0, 0, 0, 75, 80, 81, 141, 145,
	147, 153, 51, 53, 57, 0, 59, 98, 103, 104,
	0, 0, 108, 109, 110, 115, 116, 5, 18, 19,
	160, 27, 94, 50, 95, 0, 0, 73, 0, 155,
	0, 67, 132, 136, 152, 151, 138, 0, 106, 107,
	71, 0, 156, 0, 0, 58, 72, 74, 157, 0,
	0, 0, 158, 0, 159,
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
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:164
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:169
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
		//line parser.y:189
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = yyDollar[2].token
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:198
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:203
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:209
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
		//line parser.y:234
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:239
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:277
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:289
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:299
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:306
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:314
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:321
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:335
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
		//line parser.y:347
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:357
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
		//line parser.y:367
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:377
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:384
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:403
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:415
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:422
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:436
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:441
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:448
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:453
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:461
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:468
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:482
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:487
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:499
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:507
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:515
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:534
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:546
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:562
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:590
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
		//line parser.y:600
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:607
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:615
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:622
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:642
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:651
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig("true" == yyDollar[2].token)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:665
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:672
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:681
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 158:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:685
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:696
		{
			yyVAL.token = yyDollar[2].token
		}
	}
	goto yystack /* stack new state and value */
}
