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
const kywd_identity = 57399
const kywd_base = 57400

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
	"kywd_identity",
	"kywd_base",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:904

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 661

var yyAct = [...]int{

	195, 293, 193, 12, 207, 206, 12, 352, 339, 135,
	259, 285, 192, 284, 197, 256, 241, 45, 196, 76,
	44, 273, 209, 228, 235, 43, 42, 41, 40, 80,
	81, 82, 83, 39, 213, 87, 201, 202, 38, 297,
	176, 170, 161, 19, 153, 146, 140, 194, 37, 137,
	11, 136, 130, 11, 251, 128, 157, 156, 298, 299,
	3, 380, 225, 198, 199, 20, 341, 15, 168, 33,
	376, 168, 19, 270, 334, 341, 267, 375, 349, 343,
	132, 88, 378, 150, 377, 144, 125, 149, 260, 4,
	19, 147, 383, 33, 20, 138, 155, 128, 164, 225,
	172, 178, 338, 204, 204, 180, 179, 215, 221, 230,
	237, 243, 20, 258, 260, 182, 128, 208, 208, 181,
	165, 205, 205, 128, 128, 128, 134, 131, 253, 128,
	132, 252, 143, 363, 148, 162, 250, 249, 248, 247,
	144, 128, 133, 154, 246, 163, 149, 171, 177, 245,
	203, 203, 141, 155, 214, 220, 229, 236, 242, 244,
	257, 164, 223, 173, 185, 232, 231, 348, 210, 31,
	172, 295, 128, 238, 347, 381, 178, 131, 127, 276,
	180, 179, 263, 165, 346, 345, 265, 143, 128, 296,
	182, 268, 19, 148, 181, 303, 292, 272, 162, 94,
	154, 204, 301, 280, 19, 344, 262, 141, 163, 342,
	19, 289, 329, 215, 20, 208, 291, 171, 294, 205,
	275, 275, 328, 177, 128, 137, 20, 136, 230, 300,
	311, 218, 20, 173, 104, 237, 279, 305, 374, 185,
	19, 243, 168, 310, 167, 309, 371, 308, 203, 365,
	283, 317, 282, 318, 319, 362, 258, 323, 253, 321,
	214, 252, 20, 112, 314, 111, 250, 249, 248, 247,
	106, 325, 105, 361, 246, 229, 269, 313, 327, 245,
	19, 147, 236, 19, 232, 231, 332, 167, 242, 244,
	86, 188, 85, 330, 336, 275, 275, 79, 238, 78,
	324, 322, 20, 257, 189, 20, 278, 186, 187, 320,
	19, 354, 168, 316, 167, 359, 358, 312, 188, 266,
	307, 340, 306, 19, 142, 357, 23, 360, 304, 356,
	302, 189, 20, 355, 186, 187, 290, 364, 331, 19,
	271, 168, 19, 167, 367, 20, 366, 277, 110, 109,
	369, 108, 354, 158, 159, 107, 359, 358, 353, 372,
	103, 20, 102, 303, 20, 101, 357, 100, 19, 142,
	356, 23, 99, 264, 355, 368, 97, 7, 19, 24,
	95, 23, 382, 92, 91, 84, 295, 70, 77, 286,
	20, 77, 65, 64, 48, 73, 62, 63, 66, 353,
	20, 67, 261, 370, 71, 6, 25, 26, 72, 68,
	69, 30, 326, 315, 287, 17, 18, 19, 126, 74,
	124, 123, 75, 122, 190, 188, 70, 184, 121, 120,
	119, 65, 64, 118, 73, 62, 63, 66, 189, 20,
	67, 186, 187, 71, 117, 54, 116, 72, 68, 69,
	373, 115, 114, 113, 19, 96, 90, 89, 74, 28,
	27, 75, 188, 70, 200, 53, 55, 183, 65, 64,
	175, 73, 62, 63, 66, 189, 20, 67, 174, 51,
	71, 169, 19, 98, 72, 68, 69, 50, 234, 233,
	188, 70, 59, 227, 226, 74, 65, 64, 75, 73,
	62, 63, 66, 189, 20, 67, 288, 58, 71, 152,
	19, 151, 72, 68, 69, 34, 351, 350, 217, 70,
	216, 212, 211, 74, 65, 64, 75, 73, 62, 63,
	66, 56, 20, 67, 240, 239, 71, 60, 19, 191,
	72, 68, 69, 52, 337, 335, 333, 70, 281, 166,
	160, 74, 65, 64, 75, 73, 62, 63, 66, 49,
	20, 67, 224, 222, 71, 93, 219, 57, 72, 68,
	69, 255, 254, 61, 47, 46, 36, 35, 70, 74,
	274, 32, 75, 65, 64, 48, 73, 62, 63, 66,
	145, 22, 67, 139, 21, 71, 129, 70, 16, 72,
	68, 69, 65, 64, 14, 73, 62, 63, 66, 13,
	74, 67, 10, 75, 71, 9, 8, 29, 72, 68,
	69, 7, 19, 24, 5, 23, 19, 2, 1, 74,
	379, 0, 75, 0, 0, 70, 0, 0, 0, 0,
	65, 64, 0, 73, 20, 0, 66, 0, 20, 67,
	25, 26, 71, 225, 0, 0, 72, 68, 69, 17,
	18,
}
var yyPact = [...]int{

	34, -1000, 609, 456, 455, 365, -1000, 386, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 289, 386, 386, 386,
	386, 377, 282, 386, 70, 453, 452, 376, 375, 556,
	-1000, -1000, -1000, -1000, 372, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 451, 368,
	364, 359, 357, 354, 352, 224, 262, 347, 343, 341,
	340, 255, 449, 448, 447, 442, 440, 429, 426, 425,
	424, 419, 417, 416, 386, 414, 168, -1000, -1000, 191,
	132, 116, 41, 85, 355, -1000, 77, 73, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 329, -1000, 326, -1000, 525,
	404, 469, 297, 297, -1000, -1000, 179, 59, 329, 525,
	613, -1000, 30, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1, -1000, -1000, 397, 197,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 366, -1000, 310,
	-1000, -1000, 65, -1000, -1000, 267, -1000, 62, -1000, -1000,
	-1000, 331, 329, -1000, -1000, -1000, 575, 575, 339, 298,
	227, -1000, -1000, -1000, -1000, -1000, 242, 383, 410, 497,
	-1000, -1000, -1000, -1000, 327, 404, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 386, -1000, 165, 380, 10, 10,
	386, 321, 469, -1000, -1000, -1000, -1000, -1000, -1000, 319,
	297, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	313, 311, 179, -1000, -1000, -1000, -1000, 235, 386, 308,
	-1000, -1000, 22, -1000, 256, 409, 304, 329, -1000, -1000,
	-1000, 575, 575, 300, 525, -1000, -1000, -1000, -1000, 292,
	613, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 291, 30, -1000, -1000, -1000, -1000,
	408, -1000, -1000, -1000, 383, -1000, -1000, 212, -1000, -1000,
	202, -1000, -1000, 284, 575, -1000, 277, -1000, -1000, -1000,
	-1000, -1000, -1000, 56, 199, -1, -1000, -1000, -1000, -1000,
	-1000, -1000, 69, 195, 175, -1000, 174, 164, -1000, -1000,
	157, 68, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	270, -1, -1000, -1000, 469, -1000, -1000, -1000, 264, 246,
	-1000, -1000, -1000, -1000, -1000, -1000, 123, 217, -1000, -1000,
	-1000, -1000, -1000, 240, 386, 47, -1000, 53, 386, -1000,
	-1000, 399, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	237, 270, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	441, -1000, -1000, -1000, 229, -1000, 67, -1000, -1000, 60,
	74, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 14, 166,
	380, -1000, 82, -1000,
}
var yyPgo = [...]int{

	0, 630, 39, 1, 13, 11, 628, 627, 624, 617,
	405, 616, 615, 612, 47, 0, 609, 604, 67, 598,
	596, 52, 594, 593, 46, 591, 590, 45, 169, 581,
	63, 21, 580, 577, 576, 48, 38, 33, 28, 27,
	26, 25, 20, 17, 575, 574, 573, 572, 571, 15,
	10, 567, 566, 563, 54, 562, 12, 559, 550, 42,
	37, 22, 549, 548, 546, 545, 544, 543, 539, 2,
	18, 14, 537, 535, 534, 16, 531, 522, 521, 34,
	520, 518, 7, 5, 4, 517, 516, 515, 511, 509,
	44, 57, 56, 507, 494, 493, 23, 492, 489, 488,
	24, 487, 483, 481, 41, 479, 478, 470, 40, 467,
	466, 465, 64, 464, 36, 445, 8, 9,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 32, 32,
	45, 45, 46, 47, 47, 48, 48, 49, 49, 49,
	50, 41, 52, 52, 52, 52, 51, 53, 53, 54,
	55, 33, 57, 58, 58, 59, 59, 59, 59, 4,
	4, 61, 60, 62, 63, 63, 64, 64, 64, 64,
	64, 66, 66, 36, 67, 68, 68, 56, 56, 69,
	69, 69, 69, 69, 72, 44, 73, 73, 74, 74,
	75, 75, 75, 75, 75, 75, 75, 75, 75, 75,
	75, 75, 76, 40, 40, 77, 77, 78, 78, 79,
	79, 79, 81, 82, 82, 82, 82, 82, 82, 82,
	80, 80, 85, 86, 86, 29, 87, 88, 88, 89,
	89, 90, 90, 90, 90, 91, 92, 42, 93, 94,
	94, 95, 95, 96, 96, 96, 96, 43, 97, 98,
	98, 99, 99, 100, 100, 100, 34, 102, 101, 103,
	103, 104, 104, 104, 35, 105, 106, 107, 107, 83,
	83, 84, 108, 108, 108, 108, 108, 108, 108, 108,
	108, 109, 39, 110, 110, 37, 111, 112, 113, 113,
	114, 114, 114, 114, 114, 114, 114, 114, 71, 5,
	5, 3, 2, 2, 70, 38, 115, 65, 65, 116,
	116, 1, 14, 15, 12, 13, 117, 117,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 4, 0, 1, 1, 1, 2, 1, 2, 4,
	2, 4, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 3, 2, 2, 1, 3, 3, 1, 1, 1,
	3, 1, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 1, 2, 4, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 4, 0, 1, 1, 2, 1,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 4, 1, 1, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 3, 3, 2, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 3, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 2, 3, 2, 1,
	2, 1, 1, 1, 4, 2, 1, 1, 2, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 3, 2, 2, 2, 4, 2, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 4, 2, 1, 2, 3,
	5, 3, 3, 3, 3, 3, 1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -87, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, 29, -57,
	-101, -105, -67, -111, -115, -110, -76, -51, -93, -97,
	-72, -46, 31, 32, 28, 27, 33, 36, 44, 45,
	22, 39, 43, 30, 54, 57, -5, 5, 10, 8,
	-5, -5, -5, -5, 8, 10, 8, -5, 11, 4,
	4, 8, 8, 9, -28, 8, 4, 8, -102, 8,
	8, 8, 8, 8, 10, 10, 8, 8, 8, 8,
	8, 10, 8, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, -5, 4, 10, 56, -20,
	-21, -14, -15, 10, 10, -117, 10, 8, 10, -23,
	-24, -18, 14, -14, -15, -26, -27, 14, -14, -15,
	10, -88, -89, -90, -14, -15, -91, -92, 24, 25,
	-58, -59, -60, -14, -15, -61, -62, 17, 15, -103,
	-104, -14, -15, -30, -106, -107, -108, -14, -15, -83,
	-84, -70, -71, -109, 23, -30, 37, 38, 21, 34,
	20, -68, -56, -69, -14, -15, -70, -71, -30, -112,
	-113, -114, -60, -14, -15, -70, -83, -84, -71, -61,
	-112, -77, -78, -79, -14, -15, -80, -81, 52, -52,
	-14, -15, -53, -54, -55, 40, -94, -95, -96, -14,
	-15, -91, -92, -98, -99, -100, -14, -15, -30, -73,
	-74, -75, -14, -15, -35, -36, -37, -38, -39, -40,
	-41, -54, -42, -43, -47, -48, -49, -14, -15, -50,
	58, 5, 9, -21, 7, -24, 9, 11, -27, 9,
	11, 9, -90, -31, -32, -30, -31, 8, 8, 9,
	-59, -63, 10, 8, -4, -5, 6, 4, 9, -104,
	9, -108, -5, -3, 53, 6, -3, -2, 48, 49,
	-2, -5, 9, -69, 9, -114, 9, 9, -79, 10,
	8, -5, 9, -54, 8, 4, 9, -96, -31, -31,
	9, -100, 9, -75, 9, -49, 4, -4, 10, 10,
	9, -30, 9, -64, 18, -65, -50, -66, 46, -116,
	-60, 19, 10, 10, 10, 10, 10, 10, 10, 10,
	-85, -86, -82, -14, -15, -61, -70, -71, -83, -84,
	-56, 9, 9, 10, -117, 9, -5, -116, -60, -5,
	4, 9, -82, 9, 9, 10, 10, 10, 8, -1,
	47, 9, -3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 229, 16, 0,
	0, 0, 0, 0, 0, 37, 0, 0, 15, 22,
	31, 2, 3, 1, 42, 157, 156, 0, 186, 0,
	0, 105, 0, 0, 212, 133, 135, 72, 169, 179,
	116, 60, 63, 82, 188, 195, 104, 216, 236, 213,
	214, 132, 76, 168, 178, 114, 62, 6, 0, 0,
	18, 20, 21, 244, 245, 242, 246, 0, 243, 0,
	23, 26, 0, 28, 29, 0, 32, 0, 35, 36,
	25, 0, 158, 159, 161, 162, 43, 43, 0, 0,
	0, 83, 85, 86, 87, 88, 0, 0, 0, 0,
	189, 191, 192, 193, 0, 196, 197, 202, 203, 204,
	205, 206, 207, 208, 0, 210, 0, 0, 0, 0,
	0, 0, 106, 107, 109, 110, 111, 112, 113, 0,
	217, 218, 220, 221, 222, 223, 224, 225, 226, 227,
	0, 0, 136, 137, 139, 140, 141, 0, 0, 0,
	73, 74, 75, 77, 0, 0, 0, 170, 171, 173,
	174, 43, 43, 0, 180, 181, 183, 184, 185, 0,
	117, 118, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 0, 64, 65, 67, 68, 69,
	0, 230, 17, 19, 0, 24, 30, 0, 33, 38,
	0, 155, 160, 0, 44, 58, 0, 165, 166, 81,
	84, 92, 94, 0, 0, 89, 90, 93, 187, 190,
	194, 198, 0, 0, 0, 231, 0, 0, 232, 233,
	0, 0, 103, 108, 215, 219, 235, 134, 138, 150,
	0, 142, 71, 78, 0, 80, 167, 172, 0, 0,
	177, 182, 115, 119, 61, 66, 0, 0, 27, 34,
	163, 59, 164, 0, 0, 97, 98, 99, 0, 237,
	101, 0, 91, 209, 199, 200, 201, 234, 228, 211,
	0, 152, 153, 143, 144, 145, 146, 147, 148, 149,
	0, 175, 176, 70, 0, 95, 0, 238, 102, 0,
	0, 151, 154, 79, 247, 96, 100, 239, 0, 0,
	0, 240, 0, 241,
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
	52, 53, 54, 55, 56, 57, 58,
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
		//line parser.y:140
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
		//line parser.y:148
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
		//line parser.y:162
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:177
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:184
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:187
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:200
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:211
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:224
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:229
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:245
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:248
		{
			pop(yylex)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
			pop(yylex)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:287
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:292
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:311
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			pop(yylex)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:332
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:344
		{
			pop(yylex)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:349
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:359
		{
			pop(yylex)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:364
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:381
		{
			yyVAL.token = yyDollar[1].token
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:382
		{
			yyVAL.token = yyDollar[1].token
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:395
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:402
		{
			pop(yylex)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:405
		{
			pop(yylex)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:410
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:418
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:432
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:437
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:459
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:466
		{
			pop(yylex)
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:492
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:499
		{
			pop(yylex)
		}
	case 134:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:502
		{
			pop(yylex)
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:525
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:547
		{
			pop(yylex)
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:550
		{
			pop(yylex)
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:562
		{
			pop(yylex)
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:567
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:583
		{
			pop(yylex)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:586
		{
			pop(yylex)
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:591
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:598
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:608
		{
			pop(yylex)
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:613
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:629
		{
			pop(yylex)
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:632
		{
			pop(yylex)
		}
	case 177:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:640
		{
			pop(yylex)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:645
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:666
		{
			pop(yylex)
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:676
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:694
		{
			pop(yylex)
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:699
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:714
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:719
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:726
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:744
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:751
		{
			pop(yylex)
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:756
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:761
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:771
		{
			pop(yylex)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:776
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:802
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:809
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:812
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:817
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:827
		{
			yyVAL.boolean = true
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:828
		{
			yyVAL.boolean = false
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:831
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 235:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:841
		{
			pop(yylex)
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:846
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:857
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 240:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:862
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:869
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:874
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 243:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:881
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:888
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 245:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:895
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
