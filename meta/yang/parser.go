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

//line parser.y:753

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 477

var yyAct = [...]int{

	168, 272, 114, 167, 173, 203, 229, 194, 189, 30,
	176, 169, 239, 154, 10, 148, 10, 139, 182, 218,
	131, 125, 172, 177, 119, 174, 240, 241, 296, 191,
	3, 30, 6, 18, 22, 170, 14, 18, 11, 146,
	11, 145, 51, 18, 273, 171, 18, 58, 57, 43,
	65, 55, 56, 59, 215, 19, 60, 212, 63, 19,
	23, 24, 64, 61, 62, 269, 273, 78, 19, 16,
	17, 224, 191, 18, 294, 18, 293, 146, 116, 145,
	115, 111, 299, 292, 199, 200, 228, 122, 227, 127,
	28, 291, 271, 151, 162, 19, 132, 19, 141, 279,
	149, 155, 206, 178, 178, 143, 278, 186, 195, 204,
	140, 123, 159, 128, 181, 181, 277, 83, 183, 77,
	133, 76, 142, 276, 150, 156, 275, 179, 179, 274,
	122, 187, 196, 205, 264, 158, 127, 180, 180, 220,
	220, 263, 132, 209, 214, 75, 213, 74, 18, 126,
	141, 217, 151, 262, 123, 221, 225, 143, 162, 149,
	128, 184, 140, 234, 117, 155, 133, 236, 245, 18,
	19, 18, 113, 146, 142, 145, 159, 242, 112, 163,
	136, 137, 110, 150, 18, 126, 247, 178, 69, 156,
	68, 19, 164, 19, 109, 93, 297, 250, 181, 158,
	290, 254, 220, 220, 284, 195, 19, 206, 260, 18,
	121, 179, 120, 282, 204, 261, 281, 255, 256, 267,
	265, 180, 259, 258, 266, 253, 249, 18, 248, 196,
	246, 19, 244, 235, 165, 163, 51, 161, 205, 216,
	207, 58, 57, 257, 65, 55, 56, 59, 164, 19,
	60, 157, 63, 251, 223, 280, 64, 61, 62, 222,
	208, 289, 97, 96, 283, 18, 95, 92, 91, 90,
	89, 88, 286, 163, 51, 86, 84, 81, 237, 58,
	57, 245, 65, 55, 56, 59, 164, 19, 60, 18,
	63, 230, 287, 231, 64, 61, 62, 163, 51, 285,
	243, 238, 211, 58, 57, 73, 65, 55, 56, 59,
	164, 19, 60, 233, 63, 72, 71, 18, 64, 61,
	62, 6, 18, 22, 70, 14, 51, 67, 66, 298,
	288, 58, 57, 5, 65, 55, 56, 59, 27, 19,
	60, 49, 63, 252, 19, 232, 64, 61, 62, 23,
	24, 18, 108, 107, 106, 105, 104, 210, 16, 17,
	51, 18, 121, 103, 120, 58, 57, 102, 65, 55,
	56, 59, 82, 19, 60, 101, 63, 100, 99, 98,
	64, 61, 62, 19, 175, 51, 94, 85, 80, 79,
	58, 57, 43, 65, 55, 56, 59, 25, 48, 60,
	50, 63, 160, 51, 153, 64, 61, 62, 58, 57,
	152, 65, 55, 56, 59, 46, 147, 60, 87, 63,
	45, 202, 201, 64, 61, 62, 54, 198, 197, 193,
	192, 53, 135, 134, 130, 129, 31, 166, 47, 270,
	268, 226, 144, 138, 44, 190, 188, 185, 52, 42,
	41, 40, 39, 38, 37, 36, 35, 34, 33, 32,
	219, 29, 124, 21, 118, 20, 13, 12, 9, 8,
	7, 15, 26, 4, 2, 1, 295,
}
var yyPact = [...]int{

	3, -1000, 308, 393, 19, -1000, 323, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 322, 179, 319, 311, 310, 300,
	136, 110, 55, 385, 384, 268, 362, -1000, -1000, -1000,
	-1000, 267, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 383, 266, 262, 261, 260, 259, 258,
	184, 382, 257, 254, 253, 375, 374, 373, 371, 363,
	359, 352, 351, 350, 349, 348, 183, 171, -1000, 29,
	167, 161, 69, 153, -1000, 195, -1000, 170, -1000, -1000,
	-1000, -1000, -1000, -1000, 155, -1000, 23, -1000, 337, 213,
	275, 157, 157, -1000, 150, 32, 59, 337, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 230, -1000, -1000, -1000, -1000, 252, -1000, 347, -1000,
	297, 45, -1000, -1000, 134, -1000, 42, -1000, -1000, 229,
	155, -1000, -1000, -1000, 380, 380, 250, 245, 61, -1000,
	-1000, -1000, -1000, -1000, 77, 286, 341, 303, -1000, -1000,
	-1000, -1000, 223, 213, -1000, -1000, -1000, 272, -1000, -1000,
	-1000, 296, -1000, -22, -22, 295, 222, 275, -1000, -1000,
	-1000, -1000, -1000, -1000, 220, 157, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 218, -1000, 216, -1000, -1000, -11, -1000,
	244, 339, 215, 59, -1000, -1000, -1000, 380, 380, 234,
	214, 212, 337, -1000, -1000, -1000, -1000, -1000, 286, -1000,
	-1000, 142, 130, -1000, -1000, 123, -1000, -1000, 210, 380,
	-1000, 209, -1000, -1000, -1000, -1000, -1000, -1000, 46, 118,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 115, 112, 105,
	-1000, -1000, 95, 88, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 275, -1000, -1000, -1000, 206, 203, -1000, -1000, -1000,
	-1000, 69, -1000, -1000, -1000, -1000, -1000, -1000, 194, 294,
	24, 287, -1000, 326, -1000, -1000, -1000, -1000, -1000, -1000,
	251, -1000, -1000, 190, -1000, 80, -1000, 72, 65, -1000,
	-1000, -1000, -1000, -1000, -19, 186, 325, -1000, 71, -1000,
}
var yyPgo = [...]int{

	0, 476, 12, 6, 475, 474, 473, 472, 471, 470,
	11, 2, 333, 469, 468, 35, 467, 466, 465, 464,
	24, 463, 462, 21, 90, 461, 4, 19, 460, 459,
	458, 457, 456, 455, 454, 453, 452, 451, 450, 449,
	448, 447, 446, 8, 445, 3, 444, 443, 17, 23,
	18, 442, 441, 440, 439, 438, 437, 0, 45, 22,
	436, 435, 434, 20, 433, 432, 431, 430, 429, 7,
	428, 427, 426, 422, 421, 5, 420, 418, 416, 15,
	415, 410, 404, 13, 402, 400, 398, 25, 384, 10,
	341, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 6, 6, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 18, 11,
	11, 19, 19, 20, 20, 20, 20, 16, 16, 21,
	22, 22, 23, 23, 23, 17, 17, 24, 24, 7,
	7, 27, 27, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 28, 28, 37, 41, 41, 41,
	41, 40, 42, 42, 43, 44, 29, 46, 47, 47,
	48, 48, 48, 48, 3, 3, 50, 49, 51, 52,
	52, 53, 53, 53, 32, 55, 56, 56, 45, 45,
	57, 57, 57, 57, 57, 36, 25, 60, 61, 61,
	62, 62, 63, 63, 63, 63, 64, 65, 38, 66,
	67, 67, 68, 68, 69, 69, 69, 69, 70, 71,
	39, 72, 73, 73, 74, 74, 75, 75, 75, 30,
	77, 76, 78, 78, 79, 79, 79, 31, 80, 81,
	82, 82, 83, 83, 83, 83, 83, 83, 83, 83,
	84, 35, 85, 85, 33, 86, 87, 88, 88, 89,
	89, 89, 89, 89, 89, 59, 2, 2, 58, 34,
	90, 54, 54, 91, 91, 1, 15, 13, 14,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 3, 1, 2, 3,
	1, 1, 1, 1, 1, 1, 1, 3, 2, 1,
	5, 1, 2, 3, 3, 1, 1, 2, 4, 2,
	1, 2, 3, 1, 1, 2, 4, 1, 1, 1,
	2, 0, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 0, 1, 1,
	1, 2, 1, 2, 4, 2, 4, 2, 1, 2,
	1, 1, 1, 1, 1, 1, 3, 2, 2, 1,
	3, 3, 1, 3, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 3, 4, 2, 0, 1,
	1, 2, 1, 1, 3, 3, 2, 2, 4, 2,
	0, 1, 1, 2, 1, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 1, 1, 1, 2,
	3, 2, 1, 2, 1, 1, 1, 4, 2, 1,
	1, 2, 1, 1, 3, 1, 1, 1, 3, 1,
	3, 2, 2, 2, 4, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 4,
	2, 1, 2, 3, 5, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -12, 13, -9, -13, -14,
	-10, -15, -16, -17, 17, -8, 50, 51, 14, 36,
	-18, -21, 15, 41, 42, 4, -7, -12, -24, -25,
	-26, -60, -29, -30, -31, -32, -33, -34, -35, -36,
	-37, -38, -39, 30, -46, -76, -80, -55, -86, -90,
	-85, 23, -40, -66, -72, 32, 33, 29, 28, 34,
	37, 44, 45, 39, 43, 31, 5, 5, 11, 9,
	5, 5, 5, 5, 11, 9, 11, 9, 12, 4,
	4, 9, 10, -24, 9, 4, 9, -77, 9, 9,
	9, 9, 9, 11, 4, 9, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 11,
	11, -10, 11, 11, -11, 11, 9, 11, -19, -20,
	17, 15, -10, -15, -22, -23, 15, -10, -15, -61,
	-62, -63, -10, -15, -64, -65, 25, 26, -47, -48,
	-49, -10, -15, -50, -51, 18, 16, -78, -79, -10,
	-15, -26, -81, -82, -83, -10, -15, 38, -58, -59,
	-84, 24, -26, 22, 35, 21, -56, -45, -57, -10,
	-15, -58, -59, -26, -87, -88, -89, -49, -10, -15,
	-58, -59, -50, -87, 11, -41, -10, -15, -42, -43,
	-44, 40, -67, -68, -69, -10, -15, -70, -71, 25,
	26, -73, -74, -75, -10, -15, -26, 10, 8, -20,
	10, 5, 12, -23, 10, 12, 10, -63, -27, -28,
	-26, -27, 9, 9, 10, -48, -52, 11, 9, -3,
	5, 7, 4, 10, -79, 10, -83, 6, 5, -2,
	48, 49, -2, 5, 10, -57, 10, -89, 10, 10,
	-43, 9, 4, 10, -69, -27, -27, 9, 9, 10,
	-75, -3, 11, 11, 11, 10, -26, 10, -53, 19,
	-54, 46, -91, 20, 11, 11, 11, 11, 11, 11,
	-45, 10, 10, -11, 10, 5, -91, 5, 4, 10,
	10, 11, 11, 11, 9, -1, 47, 10, 4, 11,
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
	29, 2, 1, 40, 98, 97, 0, 129, 0, 0,
	86, 0, 0, 151, 0, 57, 110, 122, 67, 131,
	138, 85, 155, 170, 152, 153, 61, 109, 121, 9,
	17, 0, 177, 178, 6, 19, 0, 176, 0, 21,
	0, 0, 25, 26, 0, 30, 0, 33, 34, 0,
	99, 100, 102, 103, 41, 41, 0, 0, 0, 68,
	70, 71, 72, 73, 0, 0, 0, 0, 132, 134,
	135, 136, 0, 139, 140, 142, 143, 0, 145, 146,
	147, 0, 149, 0, 0, 0, 0, 87, 88, 90,
	91, 92, 93, 94, 0, 156, 157, 159, 160, 161,
	162, 163, 164, 0, 95, 0, 58, 59, 60, 62,
	0, 0, 0, 111, 112, 114, 115, 41, 41, 0,
	0, 0, 123, 124, 126, 127, 128, 5, 0, 22,
	28, 0, 0, 31, 36, 0, 96, 101, 0, 42,
	54, 0, 106, 107, 66, 69, 77, 79, 0, 0,
	74, 75, 78, 130, 133, 137, 141, 0, 0, 0,
	166, 167, 0, 0, 84, 89, 154, 158, 169, 56,
	63, 0, 65, 108, 113, 0, 0, 118, 119, 120,
	125, 0, 23, 24, 32, 104, 55, 105, 0, 0,
	82, 0, 171, 0, 76, 144, 148, 168, 165, 150,
	0, 116, 117, 0, 80, 0, 172, 0, 0, 64,
	20, 81, 83, 173, 0, 0, 0, 174, 0, 175,
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
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:308
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:318
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:325
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:333
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:340
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:357
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:358
		{
			yyVAL.token = yyDollar[1].token
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:360
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:372
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:382
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:392
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:402
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:409
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:429
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:441
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:448
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:462
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:467
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:474
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:479
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:487
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:508
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:513
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:520
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:525
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:533
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:541
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:561
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:573
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:589
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:596
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:618
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:628
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:635
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:638
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:646
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:653
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:674
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:684
		{
			yyVAL.boolean = true
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:685
		{
			yyVAL.boolean = false
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:687
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:701
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:708
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:717
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 174:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:721
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:732
		{
			yyVAL.token = yyDollar[2].token
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:737
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:743
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:749
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
