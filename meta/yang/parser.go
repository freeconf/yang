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

func tokenPath(s string) string {
	if len(s) > 0 && s[0] == '"' {
		return tokenString(s)
	}
	return s
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

//line parser.y:71
type yySymType struct {
	yys     int
	token   string
	boolean bool
	stack   *yangMetaStack
	loader  ModuleLoader
}

const token_ident = 57346
const token_string = 57347
const token_path = 57348
const token_int = 57349
const token_number = 57350
const token_custom = 57351
const token_curly_open = 57352
const token_curly_close = 57353
const token_semi = 57354
const token_rev_ident = 57355
const kywd_namespace = 57356
const kywd_description = 57357
const kywd_revision = 57358
const kywd_type = 57359
const kywd_prefix = 57360
const kywd_default = 57361
const kywd_length = 57362
const kywd_enum = 57363
const kywd_key = 57364
const kywd_config = 57365
const kywd_uses = 57366
const kywd_unique = 57367
const kywd_input = 57368
const kywd_output = 57369
const kywd_module = 57370
const kywd_container = 57371
const kywd_list = 57372
const kywd_rpc = 57373
const kywd_notification = 57374
const kywd_typedef = 57375
const kywd_grouping = 57376
const kywd_leaf = 57377
const kywd_mandatory = 57378
const kywd_reference = 57379
const kywd_leaf_list = 57380
const kywd_max_elements = 57381
const kywd_choice = 57382
const kywd_case = 57383
const kywd_import = 57384
const kywd_include = 57385
const kywd_action = 57386
const kywd_anyxml = 57387
const kywd_anydata = 57388
const kywd_path = 57389
const kywd_value = 57390
const kywd_true = 57391
const kywd_false = 57392
const kywd_contact = 57393
const kywd_organization = 57394
const kywd_refine = 57395

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_path",
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
	"kywd_refine",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:825

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 513

var yyAct = [...]int{

	175, 173, 300, 11, 290, 11, 119, 172, 242, 215,
	206, 178, 201, 191, 252, 181, 30, 231, 159, 187,
	153, 182, 124, 144, 136, 130, 253, 254, 114, 177,
	321, 203, 18, 176, 3, 179, 291, 18, 30, 319,
	228, 318, 174, 18, 121, 10, 120, 10, 211, 212,
	170, 168, 63, 166, 19, 287, 291, 58, 57, 19,
	66, 55, 56, 59, 169, 19, 60, 162, 64, 18,
	196, 116, 65, 61, 62, 225, 128, 219, 133, 223,
	18, 18, 289, 18, 126, 138, 125, 147, 78, 155,
	161, 19, 184, 184, 324, 203, 193, 199, 208, 217,
	156, 167, 19, 19, 317, 19, 148, 227, 145, 316,
	218, 18, 131, 115, 116, 18, 131, 297, 127, 164,
	132, 186, 186, 163, 128, 185, 185, 137, 188, 146,
	133, 154, 160, 19, 183, 183, 138, 19, 192, 198,
	207, 216, 220, 265, 147, 264, 222, 18, 126, 28,
	125, 233, 233, 155, 296, 226, 115, 295, 234, 161,
	230, 294, 293, 148, 156, 145, 127, 238, 292, 19,
	167, 282, 132, 247, 258, 281, 83, 249, 137, 237,
	241, 184, 240, 18, 255, 151, 146, 150, 164, 280,
	122, 193, 163, 118, 322, 154, 260, 18, 95, 117,
	94, 160, 112, 111, 263, 19, 208, 93, 141, 142,
	186, 315, 312, 268, 185, 217, 272, 307, 305, 19,
	18, 233, 233, 183, 278, 304, 218, 273, 274, 63,
	279, 285, 283, 192, 58, 57, 277, 66, 55, 56,
	59, 271, 19, 60, 284, 64, 267, 262, 207, 65,
	61, 62, 6, 18, 22, 261, 14, 216, 77, 70,
	76, 69, 63, 259, 257, 248, 302, 58, 57, 43,
	66, 55, 56, 59, 229, 19, 60, 303, 64, 276,
	23, 24, 65, 61, 62, 275, 306, 269, 236, 16,
	17, 314, 235, 309, 98, 18, 97, 96, 92, 91,
	302, 90, 313, 168, 63, 258, 89, 88, 301, 58,
	57, 86, 66, 55, 56, 59, 169, 19, 60, 18,
	64, 151, 84, 150, 65, 61, 62, 168, 81, 18,
	75, 243, 221, 266, 244, 250, 310, 168, 63, 308,
	169, 19, 301, 58, 57, 256, 66, 55, 56, 59,
	169, 19, 60, 246, 64, 251, 224, 18, 65, 61,
	62, 6, 18, 22, 74, 14, 63, 73, 72, 71,
	68, 58, 57, 67, 66, 55, 56, 59, 323, 19,
	60, 311, 64, 82, 19, 270, 65, 61, 62, 23,
	24, 18, 5, 151, 49, 150, 63, 27, 16, 17,
	245, 58, 57, 43, 66, 55, 56, 59, 110, 109,
	60, 108, 64, 19, 63, 107, 65, 61, 62, 58,
	57, 106, 66, 55, 56, 59, 105, 104, 60, 103,
	64, 102, 101, 100, 65, 61, 62, 99, 85, 80,
	79, 25, 180, 48, 50, 165, 158, 157, 46, 152,
	87, 45, 214, 213, 54, 210, 209, 205, 204, 53,
	140, 139, 135, 134, 31, 299, 298, 195, 194, 190,
	189, 51, 171, 47, 288, 286, 239, 149, 143, 44,
	202, 200, 197, 52, 42, 41, 40, 39, 38, 37,
	36, 35, 34, 33, 32, 232, 29, 129, 21, 123,
	20, 13, 12, 9, 8, 113, 7, 15, 26, 4,
	2, 1, 320,
}
var yyPact = [...]int{

	6, -1000, 347, 437, 238, -1000, 368, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 365, 249, 364, 363, 362, 359,
	320, 248, 75, 436, 435, 318, 372, -1000, -1000, -1000,
	-1000, 312, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 434, 301, 297, 296, 291, 289, 288,
	195, 188, 287, 286, 284, 433, 429, 428, 427, 425,
	423, 422, 417, 411, 407, 405, 404, 191, 190, -1000,
	65, 187, 181, 34, 178, 132, -1000, 100, -1000, -1000,
	-1000, -1000, -1000, -1000, 182, -1000, 376, -1000, 205, 28,
	314, 304, 304, -1000, -1000, 17, 54, 22, 205, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 66, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 323, -1000, 68, -1000, 351, 62, -1000, -1000, 96,
	-1000, 27, -1000, -1000, 263, 182, -1000, -1000, -1000, 390,
	390, 282, 278, 168, -1000, -1000, -1000, -1000, -1000, 170,
	326, 396, 342, -1000, -1000, -1000, -1000, 254, 28, -1000,
	-1000, -1000, 328, -1000, -1000, -1000, 350, -1000, -23, -23,
	340, 253, 314, -1000, -1000, -1000, -1000, -1000, -1000, 252,
	304, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 244, 236,
	17, -1000, -1000, -1000, -1000, 133, 327, 235, -1000, -1000,
	-10, -1000, 277, 381, 230, 22, -1000, -1000, -1000, 390,
	390, 275, 269, 225, 205, -1000, -1000, -1000, -1000, -1000,
	-1000, 326, -1000, -1000, 177, 163, -1000, -1000, 159, -1000,
	-1000, 221, 390, -1000, 220, -1000, -1000, -1000, -1000, -1000,
	-1000, 35, 156, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	150, 149, 145, -1000, -1000, 142, 105, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 65, -1000, -1000, -1000, 314,
	-1000, -1000, -1000, 214, 207, -1000, -1000, -1000, -1000, 34,
	-1000, -1000, -1000, -1000, -1000, -1000, 206, 334, 15, 331,
	-1000, 377, -1000, -1000, -1000, -1000, -1000, -1000, 201, 65,
	-1000, -1000, -1000, 280, -1000, -1000, 200, -1000, 97, -1000,
	92, 29, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -18,
	183, 374, -1000, 82, -1000,
}
var yyPgo = [...]int{

	0, 512, 14, 8, 511, 510, 509, 508, 507, 506,
	505, 28, 42, 0, 6, 392, 504, 503, 502, 501,
	500, 499, 22, 498, 497, 25, 149, 496, 11, 17,
	495, 494, 493, 492, 491, 490, 489, 488, 487, 486,
	485, 484, 483, 482, 481, 12, 480, 7, 479, 478,
	23, 21, 19, 477, 476, 475, 474, 473, 472, 1,
	33, 29, 471, 470, 469, 13, 468, 467, 2, 466,
	465, 464, 463, 462, 24, 461, 460, 459, 458, 457,
	10, 456, 455, 454, 453, 452, 9, 451, 450, 449,
	20, 448, 447, 446, 18, 445, 444, 443, 35, 442,
	15, 394, 4,
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
	58, 47, 47, 59, 59, 59, 59, 59, 62, 38,
	38, 63, 63, 64, 64, 65, 65, 65, 67, 68,
	68, 66, 66, 69, 70, 70, 27, 71, 72, 72,
	73, 73, 74, 74, 74, 74, 75, 76, 40, 77,
	78, 78, 79, 79, 80, 80, 80, 80, 81, 82,
	41, 83, 84, 84, 85, 85, 86, 86, 86, 32,
	88, 87, 89, 89, 90, 90, 90, 33, 91, 92,
	93, 93, 94, 94, 94, 94, 94, 94, 94, 94,
	95, 37, 96, 96, 35, 97, 98, 99, 99, 100,
	100, 100, 100, 100, 100, 61, 2, 2, 60, 36,
	101, 56, 56, 102, 102, 1, 13, 16, 17,
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
	1, 1, 2, 1, 1, 1, 1, 1, 2, 2,
	4, 0, 1, 1, 2, 1, 1, 1, 2, 1,
	1, 2, 4, 1, 1, 2, 4, 2, 0, 1,
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

	-1000, -4, -5, 28, -6, -15, 14, -9, -16, -17,
	-12, -13, -18, -19, 18, -8, 51, 52, 15, 37,
	-20, -23, 16, 42, 43, 4, -7, -15, -26, -27,
	-28, -71, -31, -32, -33, -34, -35, -36, -37, -38,
	-39, -40, -41, 31, -48, -87, -91, -57, -97, -101,
	-96, -62, -42, -77, -83, 33, 34, 30, 29, 35,
	38, 45, 46, 24, 40, 44, 32, 5, 5, 12,
	10, 5, 5, 5, 5, 10, 12, 10, 13, 4,
	4, 10, 11, -26, 10, 4, 10, -88, 10, 10,
	10, 10, 10, 12, 12, 10, 10, 10, 10, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 12, 12, -10, -11, -12, -13, 12, 12, -14,
	12, 10, 12, -21, -22, 18, 16, -12, -13, -24,
	-25, 16, -12, -13, -72, -73, -74, -12, -13, -75,
	-76, 26, 27, -49, -50, -51, -12, -13, -52, -53,
	19, 17, -89, -90, -12, -13, -28, -92, -93, -94,
	-12, -13, 39, -60, -61, -95, 25, -28, 23, 36,
	22, -58, -47, -59, -12, -13, -60, -61, -28, -98,
	-99, -100, -51, -12, -13, -60, -61, -52, -98, -63,
	-64, -65, -12, -13, -66, -67, 53, -43, -12, -13,
	-44, -45, -46, 41, -78, -79, -80, -12, -13, -81,
	-82, 26, 27, -84, -85, -86, -12, -13, -28, 11,
	-11, 9, -22, 11, 5, 13, -25, 11, 13, 11,
	-74, -29, -30, -28, -29, 10, 10, 11, -50, -54,
	12, 10, -3, 5, 8, 4, 11, -90, 11, -94,
	7, 5, -2, 49, 50, -2, 5, 11, -59, 11,
	-100, 11, 11, -65, 12, 10, 6, 11, -45, 10,
	4, 11, -80, -29, -29, 10, 10, 11, -86, -3,
	12, 12, 12, 11, -28, 11, -55, 20, -56, 47,
	-102, 21, 12, 12, 12, 12, 12, 12, -69, -70,
	-68, -12, -13, -47, 11, 11, -14, 11, 5, -102,
	5, 4, 11, -68, 11, 11, 12, 12, 12, 10,
	-1, 48, 11, 4, 12,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 11, 0, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 42, 40,
	41, 0, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 4,
	0, 0, 0, 0, 0, 0, 38, 0, 3, 22,
	32, 2, 1, 43, 118, 117, 0, 149, 0, 0,
	89, 0, 0, 171, 99, 101, 60, 130, 142, 70,
	151, 158, 88, 175, 190, 172, 173, 98, 64, 129,
	141, 13, 21, 0, 6, 8, 9, 197, 198, 10,
	23, 0, 196, 0, 25, 0, 0, 29, 30, 0,
	33, 0, 36, 37, 0, 119, 120, 122, 123, 44,
	44, 0, 0, 0, 71, 73, 74, 75, 76, 0,
	0, 0, 0, 152, 154, 155, 156, 0, 159, 160,
	162, 163, 0, 165, 166, 167, 0, 169, 0, 0,
	0, 0, 90, 91, 93, 94, 95, 96, 97, 0,
	176, 177, 179, 180, 181, 182, 183, 184, 0, 0,
	102, 103, 105, 106, 107, 0, 0, 0, 61, 62,
	63, 65, 0, 0, 0, 131, 132, 134, 135, 44,
	44, 0, 0, 0, 143, 144, 146, 147, 148, 5,
	7, 0, 26, 31, 0, 0, 34, 39, 0, 116,
	121, 0, 45, 57, 0, 126, 127, 69, 72, 80,
	82, 0, 0, 77, 78, 81, 150, 153, 157, 161,
	0, 0, 0, 186, 187, 0, 0, 87, 92, 174,
	178, 189, 100, 104, 111, 0, 108, 59, 66, 0,
	68, 128, 133, 0, 0, 138, 139, 140, 145, 0,
	27, 28, 35, 124, 58, 125, 0, 0, 85, 0,
	191, 0, 79, 164, 168, 188, 185, 170, 0, 113,
	114, 109, 110, 0, 136, 137, 0, 83, 0, 192,
	0, 0, 112, 115, 67, 24, 84, 86, 193, 0,
	0, 0, 194, 0, 195,
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
	52, 53,
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
		//line parser.y:144
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:150
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:158
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:161
		{
			yyVAL.stack.Pop()
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:173
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:182
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:193
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:198
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
		//line parser.y:222
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:231
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:237
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
		//line parser.y:262
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:267
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:305
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:318
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:328
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:335
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:343
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:350
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:365
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:366
		{
			yyVAL.token = yyDollar[1].token
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:369
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
		//line parser.y:382
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:392
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
		//line parser.y:402
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:412
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:419
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:439
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:444
		{
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
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:474
		{
			yyVAL.stack.Push(&meta.Refine{Ident: tokenPath(yyDollar[2].token)})
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:499
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:513
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:520
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:534
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:539
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:546
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:551
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:559
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:566
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:580
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:585
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:592
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:597
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:605
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:613
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:633
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:645
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:661
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:668
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:690
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:700
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:707
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:710
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 174:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:718
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:725
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:746
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[2].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:756
		{
			yyVAL.boolean = true
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:757
		{
			yyVAL.boolean = false
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:759
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[2].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:773
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:780
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:789
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:793
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:804
		{
			yyVAL.token = yyDollar[2].token
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:809
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:815
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:821
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
