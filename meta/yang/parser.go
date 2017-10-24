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
	num     int64
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
const kywd_min_elements = 57382
const kywd_choice = 57383
const kywd_case = 57384
const kywd_import = 57385
const kywd_include = 57386
const kywd_action = 57387
const kywd_anyxml = 57388
const kywd_anydata = 57389
const kywd_path = 57390
const kywd_value = 57391
const kywd_true = 57392
const kywd_false = 57393
const kywd_contact = 57394
const kywd_organization = 57395
const kywd_refine = 57396
const kywd_unbounded = 57397

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
	"kywd_min_elements",
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
	"kywd_unbounded",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:887

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 546

var yyAct = [...]int{

	181, 179, 150, 11, 193, 11, 300, 312, 119, 178,
	251, 183, 223, 214, 239, 192, 182, 209, 199, 187,
	263, 160, 136, 154, 195, 188, 144, 130, 124, 114,
	185, 264, 265, 259, 6, 18, 22, 338, 14, 211,
	18, 3, 180, 301, 63, 10, 18, 10, 18, 58,
	57, 43, 66, 55, 56, 59, 28, 19, 60, 219,
	220, 64, 19, 23, 24, 65, 61, 62, 19, 236,
	19, 116, 16, 17, 297, 301, 128, 233, 133, 204,
	336, 260, 335, 83, 121, 138, 120, 147, 18, 156,
	162, 227, 190, 190, 164, 18, 201, 207, 216, 225,
	78, 166, 299, 194, 194, 163, 165, 341, 191, 191,
	19, 148, 145, 115, 116, 211, 334, 19, 127, 333,
	132, 184, 309, 196, 128, 308, 30, 137, 307, 146,
	133, 155, 161, 306, 189, 189, 138, 305, 200, 206,
	215, 224, 304, 228, 147, 275, 303, 274, 30, 249,
	302, 248, 230, 292, 156, 242, 115, 234, 238, 245,
	162, 291, 18, 18, 164, 151, 127, 152, 148, 145,
	246, 166, 132, 141, 142, 163, 165, 255, 137, 339,
	268, 257, 18, 131, 19, 19, 146, 190, 290, 122,
	235, 95, 118, 94, 18, 131, 155, 266, 194, 201,
	117, 112, 161, 191, 19, 77, 270, 76, 231, 332,
	157, 169, 18, 126, 216, 125, 19, 273, 111, 70,
	226, 69, 93, 225, 331, 329, 278, 282, 18, 189,
	324, 322, 283, 284, 19, 288, 175, 63, 321, 295,
	289, 200, 58, 57, 293, 66, 55, 56, 59, 176,
	19, 60, 287, 281, 64, 277, 215, 272, 65, 61,
	62, 241, 241, 18, 126, 224, 125, 261, 271, 269,
	6, 18, 22, 267, 14, 157, 314, 256, 315, 237,
	319, 169, 18, 286, 151, 19, 152, 317, 285, 320,
	175, 318, 316, 19, 279, 244, 243, 98, 323, 23,
	24, 97, 96, 176, 19, 326, 170, 171, 16, 17,
	229, 254, 314, 92, 315, 18, 319, 91, 313, 330,
	90, 89, 268, 317, 63, 88, 86, 318, 316, 58,
	57, 84, 66, 55, 56, 59, 81, 19, 60, 241,
	241, 64, 75, 276, 226, 65, 61, 62, 252, 18,
	327, 253, 325, 262, 313, 258, 174, 175, 63, 168,
	232, 74, 294, 58, 57, 73, 66, 55, 56, 59,
	176, 19, 60, 170, 171, 64, 18, 72, 71, 65,
	61, 62, 68, 67, 175, 63, 340, 5, 328, 49,
	58, 57, 27, 66, 55, 56, 59, 176, 19, 60,
	18, 280, 64, 250, 110, 109, 65, 61, 62, 63,
	108, 107, 106, 105, 58, 57, 82, 66, 55, 56,
	59, 104, 19, 60, 103, 102, 64, 101, 186, 63,
	65, 61, 62, 100, 58, 57, 43, 66, 55, 56,
	59, 99, 85, 60, 80, 18, 64, 151, 63, 152,
	65, 61, 62, 58, 57, 79, 66, 55, 56, 59,
	18, 25, 60, 48, 152, 64, 50, 19, 175, 65,
	61, 62, 167, 159, 158, 46, 153, 87, 45, 222,
	221, 176, 19, 54, 170, 171, 218, 217, 213, 212,
	53, 140, 139, 135, 134, 31, 311, 310, 203, 202,
	198, 197, 51, 177, 47, 298, 296, 247, 149, 143,
	44, 210, 208, 205, 52, 42, 41, 40, 39, 38,
	37, 36, 35, 34, 33, 32, 240, 29, 129, 21,
	123, 20, 13, 12, 9, 8, 113, 7, 15, 26,
	4, 2, 1, 173, 172, 337,
}
var yyPact = [...]int{

	13, -1000, 256, 457, 20, -1000, 378, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 377, 209, 373, 372, 360, 356,
	332, 195, 87, 451, 440, 326, 405, -1000, -1000, -1000,
	-1000, 321, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 438, 316, 315, 311, 310, 307, 303,
	210, 181, 292, 291, 287, 437, 429, 423, 421, 420,
	417, 409, 408, 407, 406, 401, 400, 206, 189, -1000,
	31, 188, 180, 74, 177, 248, -1000, 167, -1000, -1000,
	-1000, -1000, -1000, -1000, 147, -1000, 430, -1000, 385, 334,
	361, 267, 267, -1000, -1000, 25, 73, 33, 385, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 80, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 301, -1000, 197, -1000, 355, 64, -1000, -1000, 179,
	-1000, 56, -1000, -1000, 268, 147, -1000, -1000, -1000, 424,
	424, 286, 285, 148, -1000, -1000, -1000, -1000, -1000, 139,
	-1000, 399, 343, 300, -1000, -1000, -1000, -1000, 266, 334,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 350, -1000,
	26, 260, -1000, -1000, 348, -19, -19, 262, 361, -1000,
	-1000, -1000, -1000, -1000, -1000, 258, 267, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 257, 246, 25, -1000,
	-1000, -1000, -1000, 135, 337, 244, -1000, -1000, -3, -1000,
	284, 397, 242, 33, -1000, -1000, -1000, 424, 424, 278,
	273, 241, 385, -1000, -1000, -1000, -1000, -1000, -1000, 343,
	-1000, -1000, 176, 149, -1000, -1000, 141, -1000, -1000, 233,
	424, -1000, 228, -1000, -1000, -1000, -1000, -1000, -1000, 54,
	-1000, 138, -1000, -1000, -1000, -1000, -1000, -1000, 134, 130,
	125, 121, 116, 113, -1000, -1000, 110, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 445, -1000, -1000, -1000, 361,
	-1000, -1000, -1000, 227, 220, -1000, -1000, -1000, -1000, 74,
	-1000, -1000, -1000, -1000, -1000, -1000, 219, 347, 22, 345,
	-1000, 384, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	214, 445, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	213, -1000, -1000, 198, -1000, 107, -1000, 104, 70, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -12, 168, 382, -1000,
	95, -1000,
}
var yyPgo = [...]int{

	0, 545, 20, 10, 2, 544, 543, 542, 541, 540,
	539, 538, 537, 536, 29, 42, 0, 8, 387, 535,
	534, 533, 532, 531, 530, 28, 529, 528, 27, 56,
	527, 121, 14, 526, 525, 524, 523, 522, 521, 520,
	519, 518, 517, 516, 515, 514, 513, 512, 17, 511,
	9, 510, 509, 26, 25, 24, 508, 507, 506, 505,
	504, 503, 1, 16, 11, 502, 501, 500, 18, 499,
	498, 7, 15, 4, 497, 496, 495, 494, 493, 22,
	492, 491, 490, 489, 488, 13, 487, 486, 483, 480,
	479, 12, 478, 477, 476, 23, 475, 474, 473, 21,
	472, 466, 463, 30, 428, 19, 389, 6,
}
var yyR1 = [...]int{

	0, 7, 8, 11, 12, 12, 13, 13, 14, 14,
	15, 9, 9, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 23, 17, 17, 24, 24, 25, 25, 25,
	25, 21, 26, 27, 27, 28, 28, 28, 22, 22,
	29, 29, 10, 10, 32, 32, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 31, 33, 33, 42,
	46, 46, 46, 46, 45, 47, 47, 48, 49, 34,
	51, 52, 52, 53, 53, 53, 53, 3, 3, 4,
	55, 54, 56, 57, 57, 58, 58, 58, 37, 60,
	61, 61, 50, 50, 62, 62, 62, 62, 62, 65,
	41, 41, 66, 66, 67, 67, 68, 68, 68, 70,
	71, 71, 71, 71, 71, 71, 71, 69, 69, 74,
	75, 75, 30, 76, 77, 77, 78, 78, 79, 79,
	79, 79, 80, 81, 43, 82, 83, 83, 84, 84,
	85, 85, 85, 85, 86, 87, 44, 88, 89, 89,
	90, 90, 91, 91, 91, 35, 93, 92, 94, 94,
	95, 95, 95, 36, 96, 97, 98, 98, 72, 72,
	73, 99, 99, 99, 99, 99, 99, 99, 99, 99,
	100, 40, 101, 101, 38, 102, 103, 104, 104, 105,
	105, 105, 105, 105, 105, 105, 105, 6, 64, 2,
	2, 5, 63, 39, 106, 59, 59, 107, 107, 1,
	16, 19, 20,
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
	1, 2, 2, 1, 3, 3, 1, 3, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 2,
	2, 4, 0, 1, 1, 2, 1, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 4, 1,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	3, 3, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 3, 3, 2, 2, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 2, 3, 2, 1, 2,
	1, 1, 1, 4, 2, 1, 1, 2, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 2, 2, 2, 4, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 1,
	1, 3, 1, 4, 2, 1, 2, 3, 5, 3,
	3, 3, 3,
}
var yyChk = [...]int{

	-1000, -7, -8, 28, -9, -18, 14, -12, -19, -20,
	-15, -16, -21, -22, 18, -11, 52, 53, 15, 37,
	-23, -26, 16, 43, 44, 4, -10, -18, -29, -30,
	-31, -76, -34, -35, -36, -37, -38, -39, -40, -41,
	-42, -43, -44, 31, -51, -92, -96, -60, -102, -106,
	-101, -65, -45, -82, -88, 33, 34, 30, 29, 35,
	38, 46, 47, 24, 41, 45, 32, 5, 5, 12,
	10, 5, 5, 5, 5, 10, 12, 10, 13, 4,
	4, 10, 11, -29, 10, 4, 10, -93, 10, 10,
	10, 10, 10, 12, 12, 10, 10, 10, 10, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 12, 12, -13, -14, -15, -16, 12, 12, -17,
	12, 10, 12, -24, -25, 18, 16, -15, -16, -27,
	-28, 16, -15, -16, -77, -78, -79, -15, -16, -80,
	-81, 26, 27, -52, -53, -54, -15, -16, -55, -56,
	-4, 17, 19, -94, -95, -15, -16, -31, -97, -98,
	-99, -15, -16, -72, -73, -63, -64, -100, 25, -31,
	39, 40, -5, -6, 22, 23, 36, -61, -50, -62,
	-15, -16, -63, -64, -31, -103, -104, -105, -54, -15,
	-16, -63, -72, -73, -64, -55, -103, -66, -67, -68,
	-15, -16, -69, -70, 54, -46, -15, -16, -47, -48,
	-49, 42, -83, -84, -85, -15, -16, -86, -87, 26,
	27, -89, -90, -91, -15, -16, -31, 11, -14, 9,
	-25, 11, 5, 13, -28, 11, 13, 11, -79, -32,
	-33, -31, -32, 10, 10, 11, -53, -57, 12, 10,
	4, -3, 5, 8, 11, -95, 11, -99, 5, 7,
	55, 7, 5, -2, 50, 51, -2, 11, -62, 11,
	-105, 11, 11, -68, 12, 10, 6, 11, -48, 10,
	4, 11, -85, -32, -32, 10, 10, 11, -91, -3,
	12, 12, 12, 11, -31, 11, -58, 20, -59, 48,
	-107, 21, 12, 12, 12, 12, 12, 12, 12, 12,
	-74, -75, -71, -15, -16, -4, -63, -64, -72, -73,
	-50, 11, 11, -17, 11, 5, -107, 5, 4, 11,
	-71, 11, 11, 12, 12, 12, 10, -1, 49, 11,
	4, 12,
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
	32, 2, 1, 43, 124, 123, 0, 155, 0, 0,
	90, 0, 0, 181, 100, 102, 60, 136, 148, 70,
	157, 164, 89, 185, 204, 182, 183, 99, 64, 135,
	147, 13, 21, 0, 6, 8, 9, 211, 212, 10,
	23, 0, 210, 0, 25, 0, 0, 29, 30, 0,
	33, 0, 36, 37, 0, 125, 126, 128, 129, 44,
	44, 0, 0, 0, 71, 73, 74, 75, 76, 0,
	80, 0, 0, 0, 158, 160, 161, 162, 0, 165,
	166, 171, 172, 173, 174, 175, 176, 177, 0, 179,
	0, 0, 202, 198, 0, 0, 0, 0, 91, 92,
	94, 95, 96, 97, 98, 0, 186, 187, 189, 190,
	191, 192, 193, 194, 195, 196, 0, 0, 103, 104,
	106, 107, 108, 0, 0, 0, 61, 62, 63, 65,
	0, 0, 0, 137, 138, 140, 141, 44, 44, 0,
	0, 0, 149, 150, 152, 153, 154, 5, 7, 0,
	26, 31, 0, 0, 34, 39, 0, 122, 127, 0,
	45, 57, 0, 132, 133, 69, 72, 81, 83, 0,
	82, 0, 77, 78, 156, 159, 163, 167, 0, 0,
	0, 0, 0, 0, 199, 200, 0, 88, 93, 184,
	188, 203, 101, 105, 117, 0, 109, 59, 66, 0,
	68, 134, 139, 0, 0, 144, 145, 146, 151, 0,
	27, 28, 35, 130, 58, 131, 0, 0, 86, 0,
	205, 0, 79, 178, 168, 169, 170, 180, 201, 197,
	0, 119, 120, 110, 111, 112, 113, 114, 115, 116,
	0, 142, 143, 0, 84, 0, 206, 0, 0, 118,
	121, 67, 24, 85, 87, 207, 0, 0, 0, 208,
	0, 209,
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
	52, 53, 54, 55,
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
		//line parser.y:150
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:156
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:164
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:167
		{
			yyVAL.stack.Pop()
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:179
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:188
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:199
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:204
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
		//line parser.y:228
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:237
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:243
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
		//line parser.y:268
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:273
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:311
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:324
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:334
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:341
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:349
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:356
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:371
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:372
		{
			yyVAL.token = yyDollar[1].token
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:375
		{
			yyVAL.token = yyDollar[2].token
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:379
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[1].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:392
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:402
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
		//line parser.y:412
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:422
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:429
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:449
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:454
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:459
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:484
		{
			yyVAL.stack.Push(&meta.Refine{Ident: tokenPath(yyDollar[2].token)})
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:496
		{
			s := yyDollar[1].token
			yyVAL.stack.Peek().(*meta.Refine).DefaultPtr = &s
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:507
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:512
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:526
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:533
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:547
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:552
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:559
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:564
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 134:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:572
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:579
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:593
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:598
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:605
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:610
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:618
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:626
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:646
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:658
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:674
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:681
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:694
		{
			n, err := strconv.ParseInt(yyDollar[2].token, 10, 32)
			if err != nil || n < 1 {
				yylex.Error(fmt.Sprintf("not a valid number for max elements %s", yyDollar[2].token))
				goto ret1
			}
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetMaxElements(int(n))
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:707
		{
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetUnbounded(true)
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:717
		{
			n, err := strconv.ParseInt(yyDollar[2].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[2].token))
				goto ret1
			}
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetMinElements(int(n))
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:742
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:752
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:759
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:762
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:770
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:777
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:800
		{
			yyVAL.boolean = yyDollar[2].boolean
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:804
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[1].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:814
		{
			yyVAL.boolean = true
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:815
		{
			yyVAL.boolean = false
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:817
		{
			yyVAL.boolean = yyDollar[2].boolean
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:821
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[1].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:835
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:842
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:851
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 208:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:855
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:866
		{
			yyVAL.token = yyDollar[2].token
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:871
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:877
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:883
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
