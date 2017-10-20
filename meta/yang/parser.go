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
const kywd_contact = 57392
const kywd_organization = 57393

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
	"kywd_contact",
	"kywd_organization",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:747

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 471

var yyAct = [...]int{

	166, 266, 115, 165, 170, 223, 198, 189, 184, 30,
	173, 167, 212, 233, 11, 153, 11, 139, 131, 125,
	147, 169, 119, 171, 234, 235, 290, 186, 3, 267,
	178, 30, 19, 174, 10, 6, 19, 22, 293, 14,
	19, 126, 168, 263, 267, 51, 286, 288, 19, 287,
	58, 57, 43, 65, 55, 56, 59, 19, 18, 60,
	209, 63, 18, 23, 24, 64, 61, 62, 194, 195,
	265, 204, 16, 17, 186, 19, 121, 206, 120, 18,
	117, 111, 116, 222, 19, 221, 145, 122, 144, 127,
	78, 28, 161, 150, 160, 285, 132, 18, 141, 273,
	148, 154, 200, 175, 175, 162, 272, 182, 190, 199,
	123, 157, 128, 177, 177, 19, 179, 142, 83, 133,
	140, 19, 121, 149, 120, 271, 136, 137, 270, 269,
	122, 191, 156, 268, 176, 176, 127, 18, 258, 214,
	214, 203, 132, 18, 207, 77, 257, 76, 215, 211,
	141, 150, 75, 123, 74, 256, 219, 160, 148, 128,
	69, 180, 68, 114, 154, 133, 239, 228, 230, 142,
	113, 112, 140, 208, 157, 110, 236, 19, 126, 109,
	93, 149, 218, 241, 175, 19, 19, 145, 145, 144,
	144, 291, 244, 284, 177, 156, 248, 214, 214, 18,
	190, 278, 200, 276, 254, 249, 250, 275, 255, 199,
	261, 259, 19, 253, 247, 176, 243, 242, 260, 163,
	161, 51, 159, 191, 240, 238, 58, 57, 229, 65,
	55, 56, 59, 162, 210, 60, 155, 63, 201, 252,
	251, 64, 61, 62, 245, 217, 216, 97, 96, 274,
	95, 92, 91, 90, 202, 283, 89, 88, 277, 19,
	86, 84, 81, 224, 281, 225, 280, 161, 51, 231,
	279, 237, 232, 58, 57, 239, 65, 55, 56, 59,
	162, 205, 60, 19, 63, 73, 72, 71, 64, 61,
	62, 161, 51, 70, 67, 66, 292, 58, 57, 49,
	65, 55, 56, 59, 162, 227, 60, 282, 63, 19,
	5, 246, 64, 61, 62, 27, 226, 108, 51, 107,
	106, 105, 104, 58, 57, 103, 65, 55, 56, 59,
	102, 18, 60, 101, 63, 100, 99, 19, 64, 61,
	62, 6, 19, 22, 98, 14, 51, 94, 85, 80,
	79, 58, 57, 25, 65, 55, 56, 59, 172, 18,
	60, 48, 63, 82, 18, 50, 64, 61, 62, 23,
	24, 158, 152, 151, 46, 146, 51, 87, 16, 17,
	45, 58, 57, 43, 65, 55, 56, 59, 197, 196,
	60, 19, 63, 54, 193, 192, 64, 61, 62, 188,
	51, 187, 53, 135, 134, 58, 57, 130, 65, 55,
	56, 59, 129, 31, 60, 164, 63, 47, 51, 264,
	64, 61, 62, 58, 57, 262, 65, 55, 56, 59,
	220, 143, 60, 138, 63, 44, 185, 183, 64, 61,
	62, 181, 52, 42, 41, 40, 39, 38, 37, 36,
	35, 34, 33, 32, 213, 29, 124, 21, 118, 20,
	13, 12, 9, 8, 7, 15, 26, 4, 2, 1,
	289,
}
var yyPact = [...]int{

	1, -1000, 328, 349, 22, -1000, 290, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 289, 151, 288, 282, 281, 280,
	143, 136, 78, 346, 345, 253, 353, -1000, -1000, -1000,
	-1000, 252, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 344, 251, 248, 247, 244, 243, 242,
	169, 343, 241, 239, 238, 340, 332, 331, 329, 326,
	321, 318, 317, 316, 315, 313, 168, 164, -1000, 18,
	160, 159, 152, 71, -1000, 107, -1000, 26, -1000, -1000,
	-1000, -1000, -1000, -1000, 101, -1000, 171, -1000, 323, 198,
	269, 70, 70, -1000, 150, 34, 43, 377, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 228, -1000, -1000, -1000, -1000, -1000, 246, 61, -1000,
	276, 65, -1000, -1000, 163, -1000, 48, -1000, -1000, 224,
	101, -1000, -1000, -1000, 395, 395, 237, 236, 172, -1000,
	-1000, -1000, -1000, 74, 258, 312, 295, -1000, -1000, -1000,
	-1000, 218, 198, -1000, -1000, 263, -1000, -1000, -1000, 267,
	-1000, -24, -24, 266, 215, 269, -1000, -1000, -1000, -1000,
	-1000, 214, 70, -1000, -1000, -1000, -1000, -1000, -1000, 207,
	-1000, 206, -1000, -13, -1000, 235, 307, 204, 43, -1000,
	-1000, -1000, 395, 395, 231, 230, 203, 377, -1000, -1000,
	-1000, -1000, 258, -1000, -1000, 144, 135, -1000, -1000, 127,
	-1000, -1000, 201, 395, -1000, 200, -1000, -1000, -1000, -1000,
	-1000, -1000, 24, 122, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 118, 117, 114, -1000, -1000, 95, 88, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 269, -1000, -1000, -1000, 197,
	193, -1000, -1000, -1000, -1000, 71, -1000, -1000, -1000, -1000,
	-1000, -1000, 191, 265, 9, 259, -1000, 303, -1000, -1000,
	-1000, -1000, -1000, -1000, 245, -1000, -1000, 183, -1000, 84,
	-1000, 35, 38, -1000, -1000, -1000, -1000, -1000, -21, 181,
	292, -1000, 27, -1000,
}
var yyPgo = [...]int{

	0, 470, 13, 5, 469, 468, 467, 466, 465, 464,
	11, 2, 310, 463, 462, 34, 461, 460, 459, 458,
	22, 457, 456, 19, 91, 455, 4, 12, 454, 453,
	452, 451, 450, 449, 448, 447, 446, 445, 444, 443,
	442, 441, 437, 8, 436, 3, 435, 433, 17, 33,
	30, 431, 430, 425, 419, 417, 415, 0, 42, 21,
	413, 412, 407, 18, 404, 403, 402, 401, 399, 7,
	395, 394, 393, 389, 388, 6, 380, 377, 375, 20,
	374, 373, 372, 15, 371, 365, 361, 23, 358, 10,
	299, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 6, 6, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 18, 11,
	11, 19, 19, 20, 20, 20, 20, 16, 16, 21,
	22, 22, 23, 23, 23, 17, 17, 24, 24, 7,
	7, 27, 27, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 28, 28, 37, 41, 41, 41,
	40, 42, 42, 43, 44, 29, 46, 47, 47, 48,
	48, 48, 3, 3, 50, 49, 51, 52, 52, 53,
	53, 53, 32, 55, 56, 56, 45, 45, 57, 57,
	57, 57, 36, 25, 60, 61, 61, 62, 62, 63,
	63, 63, 63, 64, 65, 38, 66, 67, 67, 68,
	68, 69, 69, 69, 69, 70, 71, 39, 72, 73,
	73, 74, 74, 75, 75, 30, 77, 76, 78, 78,
	79, 79, 79, 31, 80, 81, 82, 82, 83, 83,
	83, 83, 83, 83, 83, 84, 35, 85, 85, 33,
	86, 87, 88, 88, 89, 89, 89, 89, 89, 59,
	2, 2, 58, 34, 90, 54, 54, 91, 91, 1,
	15, 13, 14,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 3, 1, 2, 3,
	1, 1, 1, 1, 1, 1, 1, 3, 2, 1,
	5, 1, 2, 3, 3, 1, 1, 2, 4, 2,
	1, 2, 3, 1, 1, 2, 4, 1, 1, 1,
	2, 0, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 0, 1, 1,
	2, 1, 2, 4, 2, 4, 2, 1, 2, 1,
	1, 1, 1, 1, 3, 2, 2, 1, 3, 3,
	1, 3, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 3, 4, 2, 0, 1, 1, 2, 1,
	1, 3, 3, 2, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 3, 3, 2, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 2, 3, 2, 1, 2,
	1, 1, 1, 4, 2, 1, 1, 2, 1, 3,
	1, 1, 1, 3, 1, 3, 2, 2, 2, 4,
	2, 1, 1, 2, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 4, 2, 1, 2, 3, 5, 3,
	3, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -12, 13, -9, -13, -14,
	-15, -10, -16, -17, 17, -8, 50, 51, 36, 14,
	-18, -21, 15, 41, 42, 4, -7, -12, -24, -25,
	-26, -60, -29, -30, -31, -32, -33, -34, -35, -36,
	-37, -38, -39, 30, -46, -76, -80, -55, -86, -90,
	-85, 23, -40, -66, -72, 32, 33, 29, 28, 34,
	37, 44, 45, 39, 43, 31, 5, 5, 11, 9,
	5, 5, 5, 5, 11, 9, 11, 9, 12, 4,
	4, 9, 10, -24, 9, 4, 9, -77, 9, 9,
	9, 9, 9, 11, 4, 9, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 11,
	11, -10, 11, 11, 11, -11, 11, 9, -19, -20,
	17, 15, -10, -15, -22, -23, 15, -10, -15, -61,
	-62, -63, -10, -15, -64, -65, 25, 26, -47, -48,
	-49, -10, -50, -51, 18, 16, -78, -79, -10, -15,
	-26, -81, -82, -83, -10, 38, -58, -59, -84, 24,
	-26, 22, 35, 21, -56, -45, -57, -10, -58, -59,
	-26, -87, -88, -89, -49, -10, -58, -59, -50, -87,
	11, -41, -10, -42, -43, -44, 40, -67, -68, -69,
	-10, -15, -70, -71, 25, 26, -73, -74, -75, -10,
	-26, 10, 8, -20, 10, 5, 12, -23, 10, 12,
	10, -63, -27, -28, -26, -27, 9, 9, 10, -48,
	-52, 11, 9, -3, 5, 7, 4, 10, -79, 10,
	-83, 6, 5, -2, 48, 49, -2, 5, 10, -57,
	10, -89, 10, 10, -43, 9, 4, 10, -69, -27,
	-27, 9, 9, 10, -75, -3, 11, 11, 11, 10,
	-26, 10, -53, 19, -54, 46, -91, 20, 11, 11,
	11, 11, 11, 11, -45, 10, 10, -11, 10, 5,
	-91, 5, 4, 10, 10, 11, 11, 11, 9, -1,
	47, 10, 4, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 7, 0, 10, 11, 12,
	13, 14, 15, 16, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 39, 37,
	38, 0, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 4, 0,
	0, 0, 0, 0, 27, 0, 35, 0, 3, 18,
	29, 2, 1, 40, 95, 94, 0, 125, 0, 0,
	84, 0, 0, 146, 0, 57, 107, 119, 66, 127,
	134, 83, 150, 164, 147, 148, 60, 106, 118, 9,
	17, 0, 171, 172, 170, 6, 19, 0, 0, 21,
	0, 0, 25, 26, 0, 30, 0, 33, 34, 0,
	96, 97, 99, 100, 41, 41, 0, 0, 0, 67,
	69, 70, 71, 0, 0, 0, 0, 128, 130, 131,
	132, 0, 135, 136, 138, 0, 140, 141, 142, 0,
	144, 0, 0, 0, 0, 85, 86, 88, 89, 90,
	91, 0, 151, 152, 154, 155, 156, 157, 158, 0,
	92, 0, 58, 59, 61, 0, 0, 0, 108, 109,
	111, 112, 41, 41, 0, 0, 0, 120, 121, 123,
	124, 5, 0, 22, 28, 0, 0, 31, 36, 0,
	93, 98, 0, 42, 54, 0, 103, 104, 65, 68,
	75, 77, 0, 0, 72, 73, 76, 126, 129, 133,
	137, 0, 0, 0, 160, 161, 0, 0, 82, 87,
	149, 153, 163, 56, 62, 0, 64, 105, 110, 0,
	0, 115, 116, 117, 122, 0, 23, 24, 32, 101,
	55, 102, 0, 0, 80, 0, 165, 0, 74, 139,
	143, 162, 159, 145, 0, 113, 114, 0, 78, 0,
	166, 0, 0, 63, 20, 79, 81, 167, 0, 0,
	0, 168, 0, 169,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
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
		//line parser.y:136
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:142
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:150
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:153
		{
			yyVAL.stack.Pop()
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:157
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:167
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:178
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:183
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
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:207
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:216
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:221
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:227
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
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:252
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:257
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:295
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:307
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:317
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:324
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:332
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:339
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:355
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:356
		{
			yyVAL.token = yyDollar[1].token
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:358
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:370
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:380
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:390
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:400
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:407
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:426
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:438
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:445
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:459
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:464
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:471
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:476
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:484
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:491
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:505
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:510
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:517
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:522
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:530
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:538
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:557
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:585
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:592
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:613
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:623
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:630
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:633
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:641
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:648
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:668
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:678
		{
			yyVAL.boolean = true
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:679
		{
			yyVAL.boolean = false
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:681
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:695
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:702
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:711
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 168:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:715
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:726
		{
			yyVAL.token = yyDollar[2].token
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:731
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Reference = tokenString(yyDollar[2].token)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:737
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:743
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
