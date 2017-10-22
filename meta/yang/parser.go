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

//line parser.y:761

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 486

var yyAct = [...]int{

	171, 276, 117, 170, 176, 197, 179, 206, 233, 30,
	243, 173, 192, 157, 11, 151, 11, 142, 134, 128,
	175, 222, 122, 112, 174, 177, 6, 18, 22, 300,
	14, 30, 185, 180, 244, 245, 51, 194, 273, 277,
	277, 58, 57, 43, 65, 55, 56, 59, 3, 19,
	60, 303, 63, 18, 23, 24, 64, 61, 62, 28,
	219, 166, 51, 16, 17, 275, 18, 58, 57, 18,
	65, 55, 56, 59, 167, 19, 60, 298, 63, 297,
	296, 114, 64, 61, 62, 216, 83, 126, 19, 131,
	78, 19, 194, 154, 165, 210, 136, 295, 145, 18,
	153, 159, 209, 182, 182, 283, 282, 190, 199, 208,
	162, 281, 184, 184, 161, 280, 183, 183, 186, 146,
	143, 19, 172, 114, 279, 10, 18, 10, 18, 18,
	149, 278, 148, 126, 268, 211, 166, 202, 203, 131,
	139, 140, 224, 224, 213, 136, 267, 217, 19, 167,
	19, 19, 221, 145, 18, 154, 149, 266, 148, 229,
	225, 165, 153, 187, 18, 124, 238, 123, 159, 301,
	240, 249, 18, 129, 146, 143, 19, 162, 246, 120,
	119, 161, 118, 116, 115, 251, 19, 232, 110, 231,
	182, 77, 113, 76, 19, 75, 109, 74, 125, 184,
	130, 93, 258, 183, 254, 224, 224, 135, 199, 144,
	209, 152, 158, 264, 181, 181, 294, 208, 189, 198,
	207, 265, 259, 260, 288, 6, 18, 22, 270, 14,
	286, 228, 285, 271, 113, 18, 69, 149, 68, 148,
	269, 218, 263, 257, 125, 18, 129, 253, 19, 252,
	130, 250, 248, 23, 24, 239, 135, 19, 220, 284,
	262, 261, 16, 17, 144, 255, 227, 19, 287, 291,
	226, 97, 96, 152, 18, 95, 290, 92, 91, 158,
	90, 168, 166, 51, 164, 249, 89, 88, 58, 57,
	86, 65, 55, 56, 59, 167, 19, 60, 160, 63,
	84, 181, 81, 64, 61, 62, 241, 234, 293, 235,
	289, 212, 18, 49, 247, 242, 215, 73, 214, 198,
	166, 51, 18, 124, 72, 123, 58, 57, 207, 65,
	55, 56, 59, 167, 19, 60, 237, 63, 71, 70,
	18, 64, 61, 62, 19, 67, 66, 5, 302, 51,
	292, 256, 27, 236, 58, 57, 108, 65, 55, 56,
	59, 107, 19, 60, 18, 63, 106, 105, 104, 64,
	61, 62, 103, 51, 102, 101, 100, 99, 58, 57,
	98, 65, 55, 56, 59, 82, 19, 60, 94, 63,
	85, 80, 79, 64, 61, 62, 25, 178, 51, 48,
	50, 163, 156, 58, 57, 43, 65, 55, 56, 59,
	155, 46, 60, 150, 63, 87, 51, 45, 64, 61,
	62, 58, 57, 205, 65, 55, 56, 59, 204, 54,
	60, 201, 63, 200, 196, 195, 64, 61, 62, 53,
	138, 137, 133, 132, 31, 169, 47, 274, 272, 230,
	147, 141, 44, 193, 191, 188, 52, 42, 41, 40,
	39, 38, 37, 36, 35, 34, 33, 32, 223, 29,
	127, 21, 121, 20, 13, 12, 9, 8, 111, 7,
	15, 26, 4, 2, 1, 299,
}
var yyPact = [...]int{

	21, -1000, 212, 392, 13, -1000, 341, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 340, 227, 334, 333, 319, 312,
	186, 182, 78, 388, 387, 293, 375, -1000, -1000, -1000,
	-1000, 291, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 386, 281, 278, 277, 271, 269, 268,
	190, 384, 266, 263, 262, 376, 373, 372, 371, 370,
	368, 364, 363, 362, 357, 352, 185, 177, -1000, 55,
	173, 172, 171, 168, -1000, 150, -1000, 158, -1000, -1000,
	-1000, -1000, -1000, -1000, 115, -1000, 140, -1000, 350, 260,
	39, 114, 114, -1000, 152, 52, 112, 350, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 85, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 303,
	-1000, 308, -1000, 311, 73, -1000, -1000, 231, -1000, 48,
	-1000, -1000, 248, 115, -1000, -1000, -1000, 393, 393, 261,
	257, 221, -1000, -1000, -1000, -1000, -1000, 178, 302, 349,
	326, -1000, -1000, -1000, -1000, 245, 260, -1000, -1000, -1000,
	300, -1000, -1000, -1000, 310, -1000, -14, -14, 309, 242,
	39, -1000, -1000, -1000, -1000, -1000, -1000, 241, 114, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 239, -1000, 237, -1000,
	-1000, -3, -1000, 256, 347, 233, 112, -1000, -1000, -1000,
	393, 393, 252, 251, 232, 350, -1000, -1000, -1000, -1000,
	-1000, -1000, 302, -1000, -1000, 146, 135, -1000, -1000, 123,
	-1000, -1000, 230, 393, -1000, 223, -1000, -1000, -1000, -1000,
	-1000, -1000, 19, 120, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 113, 104, 100, -1000, -1000, 95, 94, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 39, -1000, -1000, -1000, 222,
	220, -1000, -1000, -1000, -1000, 171, -1000, -1000, -1000, -1000,
	-1000, -1000, 214, 305, 20, 264, -1000, 346, -1000, -1000,
	-1000, -1000, -1000, -1000, 298, -1000, -1000, 206, -1000, 86,
	-1000, 69, 68, -1000, -1000, -1000, -1000, -1000, -18, 159,
	344, -1000, 40, -1000,
}
var yyPgo = [...]int{

	0, 485, 10, 8, 484, 483, 482, 481, 480, 479,
	478, 23, 122, 11, 2, 347, 477, 476, 475, 474,
	473, 472, 22, 471, 470, 19, 59, 469, 4, 21,
	468, 467, 466, 465, 464, 463, 462, 461, 460, 459,
	458, 457, 456, 455, 454, 12, 453, 3, 452, 451,
	17, 33, 32, 450, 449, 448, 447, 446, 445, 0,
	24, 20, 444, 443, 442, 18, 441, 440, 439, 435,
	434, 5, 433, 431, 429, 428, 423, 7, 417, 415,
	413, 15, 411, 410, 402, 13, 401, 400, 399, 25,
	397, 6, 313, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 8, 9, 9, 10, 10, 11, 11,
	12, 6, 6, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 20, 14, 14, 21, 21, 22, 22, 22,
	22, 18, 18, 23, 24, 24, 25, 25, 25, 19,
	19, 26, 26, 7, 7, 29, 29, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 30, 30,
	39, 43, 43, 43, 43, 42, 44, 44, 45, 46,
	31, 48, 49, 49, 50, 50, 50, 50, 3, 3,
	52, 51, 53, 54, 54, 55, 55, 55, 34, 57,
	58, 58, 47, 47, 59, 59, 59, 59, 59, 38,
	27, 62, 63, 63, 64, 64, 65, 65, 65, 65,
	66, 67, 40, 68, 69, 69, 70, 70, 71, 71,
	71, 71, 72, 73, 41, 74, 75, 75, 76, 76,
	77, 77, 77, 32, 79, 78, 80, 80, 81, 81,
	81, 33, 82, 83, 84, 84, 85, 85, 85, 85,
	85, 85, 85, 85, 86, 37, 87, 87, 35, 88,
	89, 90, 90, 91, 91, 91, 91, 91, 91, 61,
	2, 2, 60, 36, 92, 56, 56, 93, 93, 1,
	13, 16, 17,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 1, 2, 1, 1,
	3, 1, 2, 3, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 5, 1, 2, 3, 3, 1,
	1, 2, 4, 2, 1, 2, 3, 1, 1, 2,
	4, 1, 1, 1, 2, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	4, 0, 1, 1, 1, 2, 1, 2, 4, 2,
	4, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	3, 2, 2, 1, 3, 3, 1, 3, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 3,
	4, 2, 0, 1, 1, 2, 1, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	3, 3, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 2, 3, 2, 1, 2, 1, 1,
	1, 4, 2, 1, 1, 2, 1, 1, 3, 1,
	1, 1, 3, 1, 3, 2, 2, 2, 4, 2,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 4, 2, 1, 2, 3, 5, 3,
	3, 3, 3,
}
var yyChk = [...]int{

	-1000, -4, -5, 27, -6, -15, 13, -9, -16, -17,
	-12, -13, -18, -19, 17, -8, 50, 51, 14, 36,
	-20, -23, 15, 41, 42, 4, -7, -15, -26, -27,
	-28, -62, -31, -32, -33, -34, -35, -36, -37, -38,
	-39, -40, -41, 30, -48, -78, -82, -57, -88, -92,
	-87, 23, -42, -68, -74, 32, 33, 29, 28, 34,
	37, 44, 45, 39, 43, 31, 5, 5, 11, 9,
	5, 5, 5, 5, 11, 9, 11, 9, 12, 4,
	4, 9, 10, -26, 9, 4, 9, -79, 9, 9,
	9, 9, 9, 11, 4, 9, 9, 9, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 11,
	11, -10, -11, -12, -13, 11, 11, -14, 11, 9,
	11, -21, -22, 17, 15, -12, -13, -24, -25, 15,
	-12, -13, -63, -64, -65, -12, -13, -66, -67, 25,
	26, -49, -50, -51, -12, -13, -52, -53, 18, 16,
	-80, -81, -12, -13, -28, -83, -84, -85, -12, -13,
	38, -60, -61, -86, 24, -28, 22, 35, 21, -58,
	-47, -59, -12, -13, -60, -61, -28, -89, -90, -91,
	-51, -12, -13, -60, -61, -52, -89, 11, -43, -12,
	-13, -44, -45, -46, 40, -69, -70, -71, -12, -13,
	-72, -73, 25, 26, -75, -76, -77, -12, -13, -28,
	10, -11, 8, -22, 10, 5, 12, -25, 10, 12,
	10, -65, -29, -30, -28, -29, 9, 9, 10, -50,
	-54, 11, 9, -3, 5, 7, 4, 10, -81, 10,
	-85, 6, 5, -2, 48, 49, -2, 5, 10, -59,
	10, -91, 10, 10, -45, 9, 4, 10, -71, -29,
	-29, 9, 9, 10, -77, -3, 11, 11, 11, 10,
	-28, 10, -55, 19, -56, 46, -93, 20, 11, 11,
	11, 11, 11, 11, -47, 10, 10, -14, 10, 5,
	-93, 5, 4, 10, 10, 11, 11, 11, 9, -1,
	47, 10, 4, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 11, 0, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 43, 41,
	42, 0, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 4, 0,
	0, 0, 0, 0, 31, 0, 39, 0, 3, 22,
	33, 2, 1, 44, 102, 101, 0, 133, 0, 0,
	90, 0, 0, 155, 0, 61, 114, 126, 71, 135,
	142, 89, 159, 174, 156, 157, 65, 113, 125, 13,
	21, 0, 6, 8, 9, 181, 182, 10, 23, 0,
	180, 0, 25, 0, 0, 29, 30, 0, 34, 0,
	37, 38, 0, 103, 104, 106, 107, 45, 45, 0,
	0, 0, 72, 74, 75, 76, 77, 0, 0, 0,
	0, 136, 138, 139, 140, 0, 143, 144, 146, 147,
	0, 149, 150, 151, 0, 153, 0, 0, 0, 0,
	91, 92, 94, 95, 96, 97, 98, 0, 160, 161,
	163, 164, 165, 166, 167, 168, 0, 99, 0, 62,
	63, 64, 66, 0, 0, 0, 115, 116, 118, 119,
	45, 45, 0, 0, 0, 127, 128, 130, 131, 132,
	5, 7, 0, 26, 32, 0, 0, 35, 40, 0,
	100, 105, 0, 46, 58, 0, 110, 111, 70, 73,
	81, 83, 0, 0, 78, 79, 82, 134, 137, 141,
	145, 0, 0, 0, 170, 171, 0, 0, 88, 93,
	158, 162, 173, 60, 67, 0, 69, 112, 117, 0,
	0, 122, 123, 124, 129, 0, 27, 28, 36, 108,
	59, 109, 0, 0, 86, 0, 175, 0, 80, 148,
	152, 172, 169, 154, 0, 120, 121, 0, 84, 0,
	176, 0, 0, 68, 24, 85, 87, 177, 0, 0,
	0, 178, 0, 179,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:224
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:229
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:235
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
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:260
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:265
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:303
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:316
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:326
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:333
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:341
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:348
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:365
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:366
		{
			yyVAL.token = yyDollar[1].token
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:368
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[2].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:380
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:390
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:400
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:410
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:417
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:437
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:449
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:456
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:475
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:482
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:487
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:495
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:502
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:521
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:528
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:533
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:541
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:549
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:581
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:597
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:604
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:626
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:636
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:643
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:646
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:654
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:661
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:682
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:692
		{
			yyVAL.boolean = true
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:693
		{
			yyVAL.boolean = false
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:695
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:709
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:716
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:725
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 178:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:729
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:740
		{
			yyVAL.token = yyDollar[2].token
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:745
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:751
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:757
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
