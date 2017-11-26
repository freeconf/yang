//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
	"github.com/freeconf/c2g/val"
	"strconv"
	"strings"
)

func tokenString(s string) string {
	return strings.Trim(s, " \t\n\r\"'")
}

func (l *lexer) Lex(lval *yySymType) int {
	t, _ := l.nextToken()
	if t.typ == ParseEof {
		return 0
	}
	lval.token = t.val
	return int(t.typ)
}

func (l *lexer) Error(e string) {
	line, col := l.Position()
	l.lastError = c2.NewErr(fmt.Sprintf("%s - line %d, col %d", e, line, col))
}

func chkErr(l yyLexer, e error) bool {
	if e == nil {
		return false
	}
	l.Error(e.Error())
	return true
}

func push(l yyLexer, m meta.Meta) bool {
	x := l.(*lexer)
	return chkErr(l, meta.Set(x.stack.Peek(), x.stack.Push(m)))
}

func set(l yyLexer, o interface{}) bool {
	x := l.(*lexer)
	return chkErr(l, meta.Set(x.stack.Peek(), o))
}

func pop(l yyLexer) {
	l.(*lexer).stack.Pop()
}

func peek(l yyLexer) meta.Meta {
	return l.(*lexer).stack.Peek()
}

//line parser.y:59
type yySymType struct {
	yys     int
	token   string
	boolean bool
	num     int64
	num32   int
}

const token_ident = 57346
const token_string = 57347
const token_number = 57348
const token_custom = 57349
const token_curly_open = 57350
const token_curly_close = 57351
const token_semi = 57352
const token_rev_ident = 57353
const kywd_namespace = 57354
const kywd_description = 57355
const kywd_revision = 57356
const kywd_type = 57357
const kywd_prefix = 57358
const kywd_default = 57359
const kywd_length = 57360
const kywd_enum = 57361
const kywd_key = 57362
const kywd_config = 57363
const kywd_uses = 57364
const kywd_unique = 57365
const kywd_input = 57366
const kywd_output = 57367
const kywd_module = 57368
const kywd_container = 57369
const kywd_list = 57370
const kywd_rpc = 57371
const kywd_notification = 57372
const kywd_typedef = 57373
const kywd_grouping = 57374
const kywd_leaf = 57375
const kywd_mandatory = 57376
const kywd_reference = 57377
const kywd_leaf_list = 57378
const kywd_max_elements = 57379
const kywd_min_elements = 57380
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
const kywd_refine = 57394
const kywd_unbounded = 57395
const kywd_augment = 57396
const kywd_submodule = 57397
const kywd_str_plus = 57398

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
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
	"kywd_augment",
	"kywd_submodule",
	"kywd_str_plus",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:860

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 617

var yyAct = [...]int{

	189, 280, 187, 12, 201, 200, 12, 333, 321, 129,
	271, 186, 235, 272, 203, 45, 191, 44, 43, 229,
	42, 73, 222, 207, 41, 195, 284, 170, 140, 190,
	260, 77, 78, 79, 80, 155, 164, 84, 188, 15,
	40, 11, 134, 151, 11, 90, 39, 150, 38, 37,
	355, 122, 359, 192, 245, 124, 147, 196, 68, 33,
	193, 285, 286, 63, 62, 47, 71, 60, 61, 64,
	3, 131, 65, 130, 219, 69, 322, 126, 257, 70,
	66, 67, 138, 33, 143, 249, 120, 362, 19, 19,
	72, 354, 357, 149, 356, 158, 122, 166, 172, 4,
	198, 198, 174, 173, 209, 215, 224, 231, 237, 159,
	20, 20, 330, 329, 176, 125, 202, 202, 282, 122,
	137, 135, 142, 247, 126, 246, 244, 175, 243, 199,
	199, 148, 242, 157, 138, 165, 171, 122, 197, 197,
	143, 324, 208, 214, 223, 230, 236, 149, 241, 226,
	167, 179, 156, 225, 240, 158, 239, 238, 122, 217,
	232, 204, 125, 144, 166, 281, 31, 132, 255, 159,
	172, 254, 137, 135, 174, 173, 252, 85, 142, 250,
	328, 128, 263, 283, 327, 148, 176, 122, 127, 290,
	267, 121, 279, 157, 326, 198, 91, 278, 288, 175,
	276, 325, 165, 259, 262, 262, 360, 209, 171, 122,
	287, 202, 156, 122, 318, 322, 19, 167, 323, 19,
	292, 313, 224, 179, 199, 312, 298, 122, 101, 231,
	295, 353, 350, 197, 122, 237, 345, 122, 20, 19,
	141, 20, 320, 219, 304, 208, 343, 310, 308, 131,
	247, 130, 246, 244, 342, 243, 305, 306, 212, 242,
	223, 20, 311, 316, 266, 226, 314, 230, 19, 225,
	162, 300, 161, 236, 297, 241, 296, 309, 307, 262,
	262, 240, 232, 239, 238, 303, 299, 7, 19, 24,
	20, 23, 253, 270, 301, 269, 19, 136, 335, 23,
	294, 293, 340, 339, 103, 265, 102, 19, 19, 162,
	20, 161, 336, 341, 338, 315, 25, 26, 20, 152,
	153, 344, 83, 264, 82, 17, 18, 337, 347, 20,
	20, 19, 346, 335, 348, 161, 334, 340, 339, 182,
	351, 76, 107, 75, 290, 291, 256, 336, 289, 338,
	19, 141, 183, 20, 277, 180, 181, 258, 106, 282,
	105, 361, 337, 7, 19, 24, 104, 23, 100, 99,
	98, 334, 20, 68, 97, 96, 94, 92, 63, 62,
	47, 71, 60, 61, 64, 89, 20, 65, 19, 136,
	69, 23, 25, 26, 70, 66, 67, 88, 81, 19,
	251, 17, 18, 74, 273, 72, 184, 182, 68, 178,
	20, 74, 248, 63, 62, 349, 71, 60, 61, 64,
	183, 20, 65, 180, 181, 69, 302, 6, 274, 70,
	66, 67, 352, 30, 119, 118, 19, 117, 116, 115,
	72, 114, 113, 112, 182, 68, 111, 110, 109, 108,
	63, 62, 53, 71, 60, 61, 64, 183, 20, 65,
	93, 19, 69, 87, 86, 28, 70, 66, 67, 182,
	68, 27, 194, 52, 54, 63, 62, 72, 71, 60,
	61, 64, 183, 20, 65, 275, 177, 69, 169, 19,
	168, 70, 66, 67, 50, 163, 95, 49, 68, 228,
	227, 58, 72, 63, 62, 221, 71, 60, 61, 64,
	220, 20, 65, 57, 19, 69, 146, 145, 34, 70,
	66, 67, 332, 68, 331, 211, 210, 206, 63, 62,
	72, 71, 60, 61, 64, 205, 20, 65, 55, 234,
	69, 233, 68, 59, 70, 66, 67, 63, 62, 185,
	71, 60, 61, 64, 51, 72, 65, 319, 19, 69,
	317, 268, 160, 70, 66, 67, 154, 68, 48, 218,
	216, 213, 63, 62, 72, 71, 56, 46, 64, 36,
	20, 65, 35, 261, 69, 219, 32, 139, 70, 66,
	67, 19, 22, 162, 133, 161, 21, 123, 16, 182,
	14, 13, 10, 9, 8, 29, 5, 2, 1, 358,
	0, 0, 183, 20, 0, 180, 181,
}
var yyPact = [...]int{

	44, -1000, 275, 467, 461, 351, -1000, 406, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 333, 406, 406, 406,
	406, 390, 314, 406, 166, 460, 459, 389, 377, 36,
	-1000, -1000, -1000, -1000, 369, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 456, 368, 367,
	366, 362, 361, 360, 218, 296, 358, 352, 350, 334,
	445, 444, 443, 442, 439, 438, 437, 435, 434, 433,
	431, 430, 406, 181, -1000, -1000, 75, 178, 171, 63,
	157, 375, -1000, 226, 153, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 295, -1000, 294, -1000, 501, 386, 448, 578,
	578, -1000, -1000, 206, 203, 295, 501, 545, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-5, -1000, 407, 76, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 393, -1000, 283, -1000, -1000, 160, -1000, -1000, 337,
	-1000, 67, -1000, -1000, -1000, 348, 295, -1000, -1000, -1000,
	520, 520, 315, 297, 255, -1000, -1000, -1000, -1000, -1000,
	285, 398, 424, 476, -1000, -1000, -1000, -1000, 345, 386,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 406, -1000,
	112, 353, 13, 13, 406, 339, 448, -1000, -1000, -1000,
	-1000, -1000, -1000, 336, 578, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 292, 291, 206, -1000, -1000, -1000,
	-1000, 266, 406, 277, -1000, -1000, 34, -1000, 286, 422,
	276, 295, -1000, -1000, -1000, 520, 520, 269, 501, -1000,
	-1000, -1000, -1000, 268, 545, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 398, -1000, -1000, 215, -1000, -1000, 211, -1000, -1000,
	257, 520, -1000, 254, -1000, -1000, -1000, -1000, -1000, -1000,
	196, 208, -5, -1000, -1000, -1000, -1000, -1000, -1000, 131,
	191, 184, -1000, 174, 170, -1000, -1000, 103, 102, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 318, -5, -1000,
	-1000, 448, -1000, -1000, -1000, 245, 237, -1000, -1000, -1000,
	-1000, 241, -1000, -1000, -1000, -1000, -1000, 227, 406, 57,
	406, -1000, 411, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 223, 318, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 423, -1000, -1000, 222, -1000, 81, -1000, 40, 84,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 5, 197, 353,
	-1000, 77, -1000,
}
var yyPgo = [...]int{

	0, 609, 26, 1, 10, 13, 608, 607, 606, 605,
	427, 604, 603, 602, 38, 0, 601, 600, 39, 598,
	597, 55, 596, 594, 42, 592, 587, 28, 166, 586,
	53, 30, 583, 582, 579, 49, 48, 46, 40, 24,
	20, 18, 17, 15, 577, 576, 571, 570, 54, 569,
	11, 568, 566, 35, 57, 14, 562, 561, 560, 557,
	554, 549, 2, 29, 16, 543, 541, 539, 12, 538,
	535, 527, 23, 526, 525, 7, 5, 4, 524, 522,
	518, 517, 516, 56, 47, 43, 513, 510, 505, 22,
	501, 500, 499, 19, 497, 496, 495, 36, 494, 490,
	488, 27, 486, 474, 473, 60, 472, 25, 452, 8,
	9,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 32, 32, 41,
	46, 46, 46, 46, 45, 47, 47, 48, 49, 33,
	51, 52, 52, 53, 53, 53, 53, 4, 4, 55,
	54, 56, 57, 57, 58, 58, 58, 36, 60, 61,
	61, 50, 50, 62, 62, 62, 62, 62, 65, 44,
	66, 66, 67, 67, 68, 68, 68, 68, 68, 68,
	68, 68, 68, 68, 68, 68, 69, 40, 40, 70,
	70, 71, 71, 72, 72, 72, 74, 75, 75, 75,
	75, 75, 75, 75, 73, 73, 78, 79, 79, 29,
	80, 81, 81, 82, 82, 83, 83, 83, 83, 84,
	85, 42, 86, 87, 87, 88, 88, 89, 89, 89,
	89, 43, 90, 91, 91, 92, 92, 93, 93, 93,
	34, 95, 94, 96, 96, 97, 97, 97, 35, 98,
	99, 100, 100, 76, 76, 77, 101, 101, 101, 101,
	101, 101, 101, 101, 101, 102, 39, 103, 103, 37,
	104, 105, 106, 106, 107, 107, 107, 107, 107, 107,
	107, 107, 64, 5, 5, 3, 2, 2, 63, 38,
	108, 59, 59, 109, 109, 1, 14, 15, 12, 13,
	110, 110,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	0, 1, 1, 1, 2, 1, 2, 4, 2, 4,
	2, 1, 2, 1, 1, 1, 1, 1, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 2, 4,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 0,
	1, 1, 2, 1, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 3,
	3, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	2, 3, 2, 1, 2, 1, 1, 1, 4, 2,
	1, 1, 2, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3, 2, 2, 2, 4,
	2, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 1, 1, 3, 4,
	2, 1, 2, 3, 5, 3, 3, 3, 3, 3,
	1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -80, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, 29, -51, -94,
	-98, -60, -104, -108, -103, -69, -45, -86, -90, -65,
	31, 32, 28, 27, 33, 36, 44, 45, 22, 39,
	43, 30, 54, -5, 5, 10, 8, -5, -5, -5,
	-5, 8, 10, 8, -5, 11, 4, 4, 8, 8,
	9, -28, 8, 4, 8, -95, 8, 8, 8, 8,
	8, 10, 10, 8, 8, 8, 8, 8, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	-5, 10, 56, -20, -21, -14, -15, 10, 10, -110,
	10, 8, 10, -23, -24, -18, 14, -14, -15, -26,
	-27, 14, -14, -15, 10, -81, -82, -83, -14, -15,
	-84, -85, 24, 25, -52, -53, -54, -14, -15, -55,
	-56, 17, 15, -96, -97, -14, -15, -30, -99, -100,
	-101, -14, -15, -76, -77, -63, -64, -102, 23, -30,
	37, 38, 21, 34, 20, -61, -50, -62, -14, -15,
	-63, -64, -30, -105, -106, -107, -54, -14, -15, -63,
	-76, -77, -64, -55, -105, -70, -71, -72, -14, -15,
	-73, -74, 52, -46, -14, -15, -47, -48, -49, 40,
	-87, -88, -89, -14, -15, -84, -85, -91, -92, -93,
	-14, -15, -30, -66, -67, -68, -14, -15, -35, -36,
	-37, -38, -39, -40, -41, -48, -42, -43, 5, 9,
	-21, 7, -24, 9, 11, -27, 9, 11, 9, -83,
	-31, -32, -30, -31, 8, 8, 9, -53, -57, 10,
	8, -4, -5, 6, 4, 9, -97, 9, -101, -5,
	-3, 53, 6, -3, -2, 48, 49, -2, -5, 9,
	-62, 9, -107, 9, 9, -72, 10, 8, -5, 9,
	-48, 8, 4, 9, -89, -31, -31, 9, -93, 9,
	-68, -4, 10, 10, 9, -30, 9, -58, 18, -59,
	46, -109, 19, 10, 10, 10, 10, 10, 10, 10,
	10, -78, -79, -75, -14, -15, -55, -63, -64, -76,
	-77, -50, 9, 9, -110, 9, -5, -109, -5, 4,
	9, -75, 9, 9, 10, 10, 10, 8, -1, 47,
	9, -3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 213, 16, 0, 0, 0, 0,
	0, 0, 37, 0, 0, 15, 22, 31, 2, 3,
	1, 42, 141, 140, 0, 170, 0, 0, 89, 0,
	0, 196, 117, 119, 60, 153, 163, 100, 70, 172,
	179, 88, 200, 220, 197, 198, 116, 64, 152, 162,
	98, 6, 0, 0, 18, 20, 21, 228, 229, 226,
	230, 0, 227, 0, 23, 26, 0, 28, 29, 0,
	32, 0, 35, 36, 25, 0, 142, 143, 145, 146,
	43, 43, 0, 0, 0, 71, 73, 74, 75, 76,
	0, 0, 0, 0, 173, 175, 176, 177, 0, 180,
	181, 186, 187, 188, 189, 190, 191, 192, 0, 194,
	0, 0, 0, 0, 0, 0, 90, 91, 93, 94,
	95, 96, 97, 0, 201, 202, 204, 205, 206, 207,
	208, 209, 210, 211, 0, 0, 120, 121, 123, 124,
	125, 0, 0, 0, 61, 62, 63, 65, 0, 0,
	0, 154, 155, 157, 158, 43, 43, 0, 164, 165,
	167, 168, 169, 0, 101, 102, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 214, 17,
	19, 0, 24, 30, 0, 33, 38, 0, 139, 144,
	0, 44, 57, 0, 149, 150, 69, 72, 80, 82,
	0, 0, 77, 78, 81, 171, 174, 178, 182, 0,
	0, 0, 215, 0, 0, 216, 217, 0, 0, 87,
	92, 199, 203, 219, 118, 122, 134, 0, 126, 59,
	66, 0, 68, 151, 156, 0, 0, 161, 166, 99,
	103, 0, 27, 34, 147, 58, 148, 0, 0, 85,
	0, 221, 0, 79, 193, 183, 184, 185, 218, 212,
	195, 0, 136, 137, 127, 128, 129, 130, 131, 132,
	133, 0, 159, 160, 0, 83, 0, 222, 0, 0,
	135, 138, 67, 231, 84, 86, 223, 0, 0, 0,
	224, 0, 225,
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
	52, 53, 54, 55, 56,
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
		//line parser.y:138
		{
			l := yylex.(*lexer)
			if l.parent != nil {
				l.Error("expected submodule for include")
				goto ret1
			}
			yylex.(*lexer).stack.Push(meta.NewModule(yyDollar[2].token))
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:146
		{
			l := yylex.(*lexer)
			if l.parent == nil {
				/* may want to allow this is parsing submodules on their own has value */
				l.Error("submodule is for includes")
				goto ret1
			}
			l.stack.Push(l.parent)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:160
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:175
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:182
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:185
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:198
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:209
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:222
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:227
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:243
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:246
		{
			pop(yylex)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:283
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:294
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:306
		{
			pop(yylex)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:311
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			pop(yylex)
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:326
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:343
		{
			yyVAL.token = yyDollar[1].token
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:344
		{
			yyVAL.token = yyDollar[1].token
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:347
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:357
		{
			if push(yylex, meta.NewDataType(peek(yylex).(meta.HasDataType), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:364
		{
			pop(yylex)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:367
		{
			pop(yylex)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:372
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:378
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:388
		{
			pop(yylex)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:393
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:415
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:422
		{
			pop(yylex)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:448
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:455
		{
			pop(yylex)
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:458
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:481
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:503
		{
			pop(yylex)
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:506
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:518
		{
			pop(yylex)
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:523
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:539
		{
			pop(yylex)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:542
		{
			pop(yylex)
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:547
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:554
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:564
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:585
		{
			pop(yylex)
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:588
		{
			pop(yylex)
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:596
		{
			pop(yylex)
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:601
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:622
		{
			pop(yylex)
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:632
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:650
		{
			pop(yylex)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:655
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:670
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:675
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:682
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:700
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:707
		{
			pop(yylex)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:712
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:717
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:727
		{
			pop(yylex)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:732
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:758
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:765
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:768
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:773
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:783
		{
			yyVAL.boolean = true
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:784
		{
			yyVAL.boolean = false
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:787
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 219:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:797
		{
			pop(yylex)
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:802
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:813
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 224:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:818
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:825
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:830
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:837
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:844
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:851
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
