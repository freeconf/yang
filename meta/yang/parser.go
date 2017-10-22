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

//line parser.y:756

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 485

var yyAct = [...]int{

	170, 275, 116, 169, 175, 232, 205, 196, 191, 30,
	178, 172, 242, 221, 11, 156, 11, 141, 150, 127,
	174, 299, 133, 111, 173, 176, 121, 193, 184, 243,
	244, 30, 179, 6, 18, 22, 18, 14, 3, 276,
	272, 276, 302, 51, 18, 123, 18, 122, 58, 57,
	43, 65, 55, 56, 59, 218, 19, 60, 19, 63,
	18, 23, 24, 64, 61, 62, 19, 274, 19, 51,
	16, 17, 193, 215, 58, 57, 77, 65, 55, 56,
	59, 113, 19, 60, 295, 63, 125, 294, 130, 64,
	61, 62, 153, 164, 209, 135, 282, 144, 18, 152,
	158, 208, 181, 181, 18, 128, 189, 198, 207, 161,
	281, 183, 183, 160, 145, 182, 182, 185, 142, 297,
	19, 296, 113, 227, 261, 280, 19, 18, 118, 148,
	117, 147, 125, 279, 210, 18, 278, 148, 130, 147,
	277, 223, 223, 165, 135, 171, 216, 212, 10, 19,
	10, 224, 144, 267, 153, 220, 166, 19, 228, 213,
	164, 152, 266, 18, 123, 28, 122, 158, 237, 145,
	248, 239, 18, 142, 265, 18, 161, 186, 119, 245,
	160, 115, 114, 201, 202, 19, 138, 139, 250, 181,
	109, 231, 82, 230, 19, 108, 92, 19, 183, 253,
	300, 293, 182, 257, 223, 223, 18, 198, 148, 208,
	147, 263, 287, 258, 259, 112, 207, 264, 285, 76,
	124, 75, 129, 69, 260, 68, 292, 269, 19, 134,
	18, 143, 284, 151, 157, 270, 180, 180, 165, 51,
	188, 197, 206, 268, 58, 57, 262, 65, 55, 56,
	59, 166, 19, 60, 256, 63, 112, 252, 283, 64,
	61, 62, 251, 249, 217, 247, 124, 286, 18, 128,
	238, 51, 129, 219, 254, 289, 58, 57, 134, 65,
	55, 56, 59, 226, 248, 60, 143, 63, 225, 18,
	19, 64, 61, 62, 96, 151, 167, 165, 51, 163,
	95, 157, 94, 58, 57, 91, 65, 55, 56, 59,
	166, 19, 60, 159, 63, 90, 89, 18, 64, 61,
	62, 88, 87, 180, 85, 165, 51, 83, 80, 74,
	211, 58, 57, 240, 65, 55, 56, 59, 166, 19,
	60, 197, 63, 290, 288, 236, 64, 61, 62, 18,
	206, 246, 241, 6, 18, 22, 214, 14, 51, 233,
	73, 234, 72, 58, 57, 71, 65, 55, 56, 59,
	301, 19, 60, 70, 63, 81, 19, 67, 64, 61,
	62, 23, 24, 66, 5, 291, 255, 235, 51, 27,
	16, 17, 107, 58, 57, 43, 65, 55, 56, 59,
	106, 105, 60, 104, 63, 103, 102, 101, 64, 61,
	62, 100, 99, 98, 97, 93, 84, 79, 78, 25,
	49, 177, 48, 50, 162, 155, 154, 46, 149, 86,
	45, 204, 203, 54, 200, 199, 195, 194, 53, 137,
	136, 132, 131, 31, 168, 47, 273, 271, 229, 146,
	140, 44, 192, 190, 187, 52, 42, 41, 40, 39,
	38, 37, 36, 35, 34, 33, 32, 222, 29, 126,
	21, 120, 20, 13, 12, 9, 8, 110, 7, 15,
	26, 4, 2, 1, 298,
}
var yyPact = [...]int{

	11, -1000, 340, 415, 20, -1000, 378, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 372, 214, 368, 360, 357, 355,
	320, 210, 64, 414, 413, 319, 365, -1000, -1000, -1000,
	-1000, 318, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 412, 315, 313, 312, 307, 306, 296,
	185, 411, 293, 291, 285, 410, 409, 408, 407, 403,
	402, 401, 399, 397, 396, 388, 184, 179, -1000, 22,
	171, 170, 119, 167, 30, -1000, 90, -1000, -1000, -1000,
	-1000, -1000, -1000, 161, -1000, 192, -1000, 46, 275, 303,
	121, 121, -1000, 166, 32, 158, 46, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	84, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 322, -1000,
	149, -1000, 351, 61, -1000, -1000, 254, -1000, 43, -1000,
	-1000, 263, 161, -1000, -1000, -1000, 248, 248, 279, 274,
	113, -1000, -1000, -1000, -1000, -1000, 182, 354, 383, 335,
	-1000, -1000, -1000, -1000, 260, 275, -1000, -1000, -1000, 327,
	-1000, -1000, -1000, 347, -1000, -19, -19, 346, 255, 303,
	-1000, -1000, -1000, -1000, -1000, -1000, 253, 121, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 252, -1000, 247, -1000, -1000,
	-13, -1000, 265, 382, 244, 158, -1000, -1000, -1000, 248,
	248, 215, 115, 236, 46, -1000, -1000, -1000, -1000, -1000,
	-1000, 354, -1000, -1000, 163, 151, -1000, -1000, 142, -1000,
	-1000, 233, 248, -1000, 225, -1000, -1000, -1000, -1000, -1000,
	-1000, 21, 129, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	125, 122, 114, -1000, -1000, 99, 85, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 303, -1000, -1000, -1000, 222, 208,
	-1000, -1000, -1000, -1000, 119, -1000, -1000, -1000, -1000, -1000,
	-1000, 202, 339, 19, 338, -1000, 381, -1000, -1000, -1000,
	-1000, -1000, -1000, 216, -1000, -1000, 191, -1000, 76, -1000,
	73, 110, -1000, -1000, -1000, -1000, -1000, -26, 190, 366,
	-1000, 31, -1000,
}
var yyPgo = [...]int{

	0, 484, 12, 5, 483, 482, 481, 480, 479, 478,
	477, 23, 145, 11, 2, 384, 476, 475, 474, 473,
	472, 471, 26, 470, 469, 19, 165, 468, 4, 13,
	467, 466, 465, 464, 463, 462, 461, 460, 459, 458,
	457, 456, 455, 454, 453, 8, 452, 3, 451, 450,
	17, 32, 28, 449, 448, 447, 446, 445, 444, 0,
	24, 20, 443, 442, 441, 22, 440, 439, 438, 437,
	436, 7, 435, 434, 433, 432, 431, 6, 430, 429,
	428, 18, 427, 426, 425, 15, 424, 423, 422, 25,
	421, 10, 420, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 10, 11, 11,
	12, 6, 6, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 20, 14, 14, 21, 21, 22, 22, 22,
	22, 18, 23, 24, 24, 25, 25, 25, 19, 19,
	26, 26, 7, 7, 29, 29, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 30, 30, 39,
	43, 43, 43, 43, 42, 44, 44, 45, 46, 31,
	48, 49, 49, 50, 50, 50, 50, 3, 3, 52,
	51, 53, 54, 54, 55, 55, 55, 34, 57, 58,
	58, 47, 47, 59, 59, 59, 59, 59, 38, 27,
	62, 63, 63, 64, 64, 65, 65, 65, 65, 66,
	67, 40, 68, 69, 69, 70, 70, 71, 71, 71,
	71, 72, 73, 41, 74, 75, 75, 76, 76, 77,
	77, 77, 32, 79, 78, 80, 80, 81, 81, 81,
	33, 82, 83, 84, 84, 85, 85, 85, 85, 85,
	85, 85, 85, 86, 37, 87, 87, 35, 88, 89,
	90, 90, 91, 91, 91, 91, 91, 91, 61, 2,
	2, 60, 36, 92, 56, 56, 93, 93, 1, 13,
	16, 17,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 1, 2, 1, 1,
	3, 1, 2, 3, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 5, 1, 2, 3, 3, 1,
	1, 4, 2, 1, 2, 3, 1, 1, 2, 4,
	1, 1, 1, 2, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	0, 1, 1, 1, 2, 1, 2, 4, 2, 4,
	2, 1, 2, 1, 1, 1, 1, 1, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 3, 4,
	2, 0, 1, 1, 2, 1, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 3,
	3, 2, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 2, 3, 2, 1, 2, 1, 1, 1,
	4, 2, 1, 1, 2, 1, 1, 3, 1, 1,
	1, 3, 1, 3, 2, 2, 2, 4, 2, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 3, 4, 2, 1, 2, 3, 5, 3, 3,
	3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -15, 13, -9, -16, -17,
	-12, -13, -18, -19, 17, -8, 50, 51, 14, 36,
	-20, -23, 15, 41, 42, 4, -7, -15, -26, -27,
	-28, -62, -31, -32, -33, -34, -35, -36, -37, -38,
	-39, -40, -41, 30, -48, -78, -82, -57, -88, -92,
	-87, 23, -42, -68, -74, 32, 33, 29, 28, 34,
	37, 44, 45, 39, 43, 31, 5, 5, 11, 9,
	5, 5, 5, 5, 9, 11, 9, 12, 4, 4,
	9, 10, -26, 9, 4, 9, -79, 9, 9, 9,
	9, 9, 11, 4, 9, 9, 9, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 11, 11,
	-10, -11, -12, -13, 11, 11, -14, 11, 9, 11,
	-21, -22, 17, 15, -12, -13, -24, -25, 15, -12,
	-13, -63, -64, -65, -12, -13, -66, -67, 25, 26,
	-49, -50, -51, -12, -13, -52, -53, 18, 16, -80,
	-81, -12, -13, -28, -83, -84, -85, -12, -13, 38,
	-60, -61, -86, 24, -28, 22, 35, 21, -58, -47,
	-59, -12, -13, -60, -61, -28, -89, -90, -91, -51,
	-12, -13, -60, -61, -52, -89, 11, -43, -12, -13,
	-44, -45, -46, 40, -69, -70, -71, -12, -13, -72,
	-73, 25, 26, -75, -76, -77, -12, -13, -28, 10,
	-11, 8, -22, 10, 5, 12, -25, 10, 12, 10,
	-65, -29, -30, -28, -29, 9, 9, 10, -50, -54,
	11, 9, -3, 5, 7, 4, 10, -81, 10, -85,
	6, 5, -2, 48, 49, -2, 5, 10, -59, 10,
	-91, 10, 10, -45, 9, 4, 10, -71, -29, -29,
	9, 9, 10, -77, -3, 11, 11, 11, 10, -28,
	10, -55, 19, -56, 46, -93, 20, 11, 11, 11,
	11, 11, 11, -47, 10, 10, -14, 10, 5, -93,
	5, 4, 10, 10, 11, 11, 11, 9, -1, 47,
	10, 4, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 11, 0, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 42, 40,
	41, 0, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 4, 0,
	0, 0, 0, 0, 0, 38, 0, 3, 22, 32,
	2, 1, 43, 101, 100, 0, 132, 0, 0, 89,
	0, 0, 154, 0, 60, 113, 125, 70, 134, 141,
	88, 158, 173, 155, 156, 64, 112, 124, 13, 21,
	0, 6, 8, 9, 180, 181, 10, 23, 0, 179,
	0, 25, 0, 0, 29, 30, 0, 33, 0, 36,
	37, 0, 102, 103, 105, 106, 44, 44, 0, 0,
	0, 71, 73, 74, 75, 76, 0, 0, 0, 0,
	135, 137, 138, 139, 0, 142, 143, 145, 146, 0,
	148, 149, 150, 0, 152, 0, 0, 0, 0, 90,
	91, 93, 94, 95, 96, 97, 0, 159, 160, 162,
	163, 164, 165, 166, 167, 0, 98, 0, 61, 62,
	63, 65, 0, 0, 0, 114, 115, 117, 118, 44,
	44, 0, 0, 0, 126, 127, 129, 130, 131, 5,
	7, 0, 26, 31, 0, 0, 34, 39, 0, 99,
	104, 0, 45, 57, 0, 109, 110, 69, 72, 80,
	82, 0, 0, 77, 78, 81, 133, 136, 140, 144,
	0, 0, 0, 169, 170, 0, 0, 87, 92, 157,
	161, 172, 59, 66, 0, 68, 111, 116, 0, 0,
	121, 122, 123, 128, 0, 27, 28, 35, 107, 58,
	108, 0, 0, 85, 0, 174, 0, 79, 147, 151,
	171, 168, 153, 0, 119, 120, 0, 83, 0, 175,
	0, 0, 67, 24, 84, 86, 176, 0, 0, 0,
	177, 0, 178,
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
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:165
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:175
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:186
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:191
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
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:215
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:224
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:230
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
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:255
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:260
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:298
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:311
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:328
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:336
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:343
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:360
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:361
		{
			yyVAL.token = yyDollar[1].token
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:363
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:375
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:395
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:405
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:412
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:432
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:444
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:451
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:465
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:470
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:477
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:482
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:490
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:497
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:511
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:516
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:523
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:528
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:536
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:544
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:564
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:576
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:592
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:599
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:621
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:631
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:638
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:641
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:649
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:656
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:677
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:687
		{
			yyVAL.boolean = true
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:688
		{
			yyVAL.boolean = false
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:690
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:704
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:711
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:720
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 177:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:724
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:735
		{
			yyVAL.token = yyDollar[2].token
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:740
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:746
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:752
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
