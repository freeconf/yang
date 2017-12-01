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
const kywd_feature = 57401
const kywd_if_feature = 57402
const kywd_when = 57403
const kywd_must = 57404

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
	"kywd_feature",
	"kywd_if_feature",
	"kywd_when",
	"kywd_must",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:979

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 880

var yyAct = [...]int{

	210, 340, 208, 12, 227, 404, 12, 390, 226, 207,
	145, 297, 330, 302, 277, 43, 270, 294, 219, 214,
	318, 46, 262, 45, 44, 344, 213, 168, 229, 42,
	218, 47, 188, 167, 182, 163, 41, 289, 241, 173,
	215, 40, 39, 38, 156, 150, 33, 37, 138, 19,
	140, 19, 19, 216, 147, 3, 146, 434, 430, 331,
	342, 15, 259, 209, 345, 346, 11, 81, 307, 11,
	33, 20, 19, 20, 20, 392, 259, 85, 86, 87,
	88, 180, 31, 92, 4, 142, 432, 429, 431, 437,
	154, 249, 159, 78, 20, 315, 169, 200, 169, 169,
	200, 165, 138, 176, 138, 184, 190, 341, 221, 221,
	194, 233, 99, 243, 193, 253, 264, 272, 279, 113,
	296, 174, 304, 223, 223, 196, 235, 228, 228, 401,
	238, 177, 195, 138, 225, 225, 312, 237, 134, 291,
	142, 290, 288, 267, 247, 185, 199, 287, 141, 266,
	154, 151, 257, 153, 286, 158, 159, 274, 211, 285,
	284, 283, 230, 165, 164, 282, 175, 93, 183, 189,
	417, 220, 220, 176, 232, 138, 242, 395, 252, 263,
	271, 278, 184, 295, 400, 303, 394, 399, 190, 321,
	308, 174, 194, 19, 398, 310, 193, 397, 317, 19,
	313, 177, 396, 141, 343, 393, 383, 196, 320, 320,
	350, 151, 326, 153, 195, 20, 335, 250, 221, 158,
	337, 20, 185, 138, 379, 147, 164, 146, 199, 322,
	347, 378, 138, 223, 435, 19, 175, 228, 299, 160,
	169, 243, 428, 212, 225, 183, 170, 171, 352, 362,
	148, 189, 138, 144, 358, 324, 357, 20, 338, 166,
	339, 425, 264, 138, 191, 348, 222, 222, 143, 234,
	272, 244, 247, 255, 265, 273, 280, 279, 298, 356,
	305, 220, 169, 137, 365, 138, 369, 366, 367, 267,
	419, 371, 361, 416, 296, 266, 138, 415, 291, 138,
	290, 288, 304, 329, 242, 328, 287, 320, 320, 359,
	274, 373, 314, 286, 138, 376, 19, 157, 285, 284,
	283, 166, 377, 382, 282, 263, 380, 311, 121, 138,
	120, 19, 152, 271, 23, 119, 323, 118, 20, 180,
	278, 387, 385, 392, 375, 112, 191, 111, 391, 192,
	372, 224, 224, 20, 236, 370, 245, 295, 256, 406,
	381, 281, 19, 413, 180, 303, 179, 412, 368, 325,
	389, 364, 414, 19, 411, 180, 222, 179, 410, 360,
	19, 152, 299, 23, 20, 409, 355, 408, 418, 19,
	157, 110, 19, 109, 421, 20, 91, 117, 90, 244,
	203, 84, 20, 83, 406, 354, 353, 422, 413, 426,
	351, 20, 412, 204, 20, 349, 336, 350, 316, 411,
	265, 116, 405, 410, 115, 114, 108, 107, 273, 106,
	409, 192, 408, 105, 104, 280, 436, 102, 100, 169,
	200, 59, 97, 96, 89, 420, 82, 332, 6, 423,
	309, 342, 298, 82, 30, 19, 306, 424, 374, 363,
	305, 224, 205, 203, 74, 198, 333, 405, 136, 69,
	68, 135, 77, 66, 67, 70, 204, 20, 71, 201,
	202, 75, 133, 132, 245, 76, 72, 73, 131, 130,
	129, 128, 127, 126, 125, 124, 78, 123, 56, 79,
	122, 80, 169, 200, 59, 101, 7, 19, 24, 95,
	23, 7, 19, 24, 94, 23, 74, 407, 28, 27,
	281, 69, 68, 50, 77, 66, 67, 70, 217, 20,
	71, 55, 231, 75, 20, 25, 26, 76, 72, 73,
	25, 26, 427, 57, 17, 18, 19, 197, 78, 17,
	18, 79, 187, 80, 203, 74, 59, 186, 53, 181,
	69, 68, 407, 77, 66, 67, 70, 204, 20, 71,
	103, 52, 75, 269, 268, 62, 76, 72, 73, 261,
	260, 61, 162, 161, 34, 403, 402, 78, 19, 248,
	79, 246, 80, 169, 200, 59, 203, 74, 240, 239,
	58, 276, 69, 68, 275, 77, 66, 67, 70, 204,
	20, 71, 63, 206, 75, 54, 388, 386, 76, 72,
	73, 384, 327, 178, 19, 172, 51, 258, 254, 78,
	251, 60, 79, 74, 80, 169, 200, 59, 69, 68,
	293, 77, 66, 67, 70, 292, 20, 71, 64, 301,
	75, 300, 65, 49, 76, 72, 73, 48, 36, 334,
	35, 319, 32, 19, 155, 78, 22, 149, 79, 21,
	80, 169, 74, 59, 139, 16, 14, 69, 68, 13,
	77, 66, 67, 70, 10, 20, 71, 9, 8, 75,
	29, 5, 2, 76, 72, 73, 19, 1, 433, 0,
	0, 0, 0, 0, 78, 74, 0, 79, 0, 80,
	69, 68, 59, 77, 66, 67, 70, 0, 20, 71,
	0, 0, 75, 0, 0, 0, 76, 72, 73, 0,
	98, 0, 0, 0, 0, 0, 0, 78, 0, 0,
	79, 0, 80, 74, 0, 59, 0, 0, 69, 68,
	50, 77, 66, 67, 70, 0, 0, 71, 0, 0,
	75, 0, 0, 0, 76, 72, 73, 0, 0, 0,
	0, 0, 0, 0, 0, 78, 74, 0, 79, 0,
	80, 69, 68, 59, 77, 66, 67, 70, 0, 0,
	71, 0, 19, 75, 0, 0, 0, 76, 72, 73,
	0, 74, 0, 0, 0, 0, 69, 68, 78, 77,
	0, 79, 70, 80, 20, 71, 59, 0, 75, 259,
	0, 0, 76, 72, 73, 19, 0, 180, 0, 179,
	19, 0, 0, 203, 179, 0, 0, 0, 203, 169,
	200, 0, 0, 0, 0, 0, 204, 20, 0, 201,
	202, 204, 20, 0, 201, 202, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 169, 200, 59, 0, 0, 169, 0, 59,
}
var yyPact = [...]int{

	29, -1000, 499, 515, 514, 494, -1000, 448, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 393, 448, 448, 448,
	448, 436, 388, 448, 156, 510, 505, 435, 434, 721,
	-1000, -1000, -1000, -1000, 430, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	501, 429, 426, 425, 421, 419, 418, 383, 337, 448,
	417, 416, 413, 389, 327, 320, 496, 493, 491, 490,
	489, 488, 487, 486, 485, 484, 479, 478, 448, 467,
	464, 273, -1000, -1000, 186, 258, 243, 46, 240, 367,
	-1000, 376, 229, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	222, -1000, 349, -1000, 683, 442, 575, 812, 812, -1000,
	379, -1000, 39, 207, 36, 222, 611, 779, -1000, 180,
	-1000, 38, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -8, -1000, -1000, -1000, 451, 59,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 443, -1000, 318,
	-1000, -1000, 125, -1000, -1000, 303, -1000, 84, -1000, -1000,
	-1000, 409, 222, -1000, -1000, -1000, -1000, 754, 754, 448,
	328, 247, 360, -1000, -1000, -1000, -1000, -1000, 295, 441,
	462, 650, -1000, -1000, -1000, -1000, 407, 442, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 448, -1000,
	448, 54, 445, 16, 16, 448, 406, 575, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 401, 812, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	397, 396, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 377,
	39, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 246, 448,
	-1000, 370, -1000, -1000, 22, -1000, -1000, -1000, 241, 455,
	362, 222, -1000, -1000, -1000, -1000, 754, 754, 359, 611,
	-1000, -1000, -1000, -1000, -1000, 346, 779, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 341, 180, -1000, -1000, -1000, -1000, -1000, 454,
	335, 38, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 441,
	-1000, -1000, 221, -1000, -1000, 214, -1000, -1000, 317, 754,
	-1000, 314, 196, -1000, -1000, -1000, -1000, -1000, -1000, 324,
	195, -8, -1000, -1000, -1000, -1000, -1000, -1000, 176, 167,
	192, 187, -1000, 184, 177, -1000, -1000, 174, 119, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 817, -8,
	-1000, -1000, 575, -1000, -1000, -1000, 288, 284, -1000, -1000,
	-1000, -1000, -1000, -1000, 160, -1000, -1000, 217, -1000, -1000,
	-1000, -1000, -1000, -1000, 281, 448, 56, -1000, 66, 448,
	-1000, -1000, 453, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 252, 817, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 533, -1000, -1000, -1000, 233, -1000,
	77, -1000, -1000, 48, 78, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 10, 225, 445, -1000, 79, -1000,
}
var yyPgo = [...]int{

	0, 698, 25, 1, 12, 59, 697, 692, 691, 690,
	448, 688, 687, 684, 63, 0, 679, 676, 61, 675,
	674, 50, 669, 667, 45, 666, 664, 44, 82, 662,
	40, 20, 661, 660, 658, 47, 43, 42, 41, 36,
	29, 15, 24, 23, 21, 31, 657, 653, 652, 651,
	649, 13, 158, 243, 648, 645, 640, 17, 11, 631,
	630, 628, 37, 627, 9, 626, 625, 39, 18, 28,
	623, 622, 621, 617, 616, 615, 613, 2, 26, 19,
	612, 604, 601, 14, 600, 599, 598, 38, 591, 589,
	5, 8, 4, 586, 585, 584, 583, 582, 35, 33,
	27, 581, 580, 579, 22, 575, 574, 573, 16, 571,
	570, 559, 34, 558, 557, 552, 32, 547, 543, 532,
	531, 53, 528, 30, 498, 7, 10,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	32, 32, 47, 47, 48, 49, 49, 50, 50, 51,
	51, 51, 41, 52, 53, 46, 46, 54, 55, 55,
	56, 56, 57, 57, 57, 57, 58, 42, 60, 60,
	60, 60, 60, 60, 59, 61, 61, 62, 63, 33,
	65, 66, 66, 67, 67, 67, 67, 4, 4, 69,
	68, 70, 71, 71, 72, 72, 72, 72, 72, 74,
	74, 36, 75, 76, 76, 64, 64, 77, 77, 77,
	77, 77, 77, 77, 80, 45, 81, 81, 82, 82,
	83, 83, 83, 83, 83, 83, 83, 83, 83, 83,
	83, 83, 83, 83, 84, 40, 40, 85, 85, 86,
	86, 87, 87, 87, 87, 87, 87, 89, 90, 90,
	90, 90, 90, 90, 90, 90, 90, 88, 88, 93,
	94, 94, 29, 95, 96, 96, 97, 97, 98, 98,
	98, 98, 98, 99, 100, 43, 101, 102, 102, 103,
	103, 104, 104, 104, 104, 104, 44, 105, 106, 106,
	107, 107, 108, 108, 108, 108, 34, 110, 109, 111,
	111, 112, 112, 112, 35, 113, 114, 115, 115, 91,
	91, 92, 116, 116, 116, 116, 116, 116, 116, 116,
	116, 116, 116, 117, 39, 39, 119, 119, 119, 119,
	119, 119, 119, 118, 118, 37, 120, 121, 122, 122,
	123, 123, 123, 123, 123, 123, 123, 123, 123, 123,
	123, 79, 5, 5, 3, 2, 2, 78, 38, 124,
	73, 73, 125, 125, 1, 14, 15, 12, 13, 126,
	126,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 3, 3, 3, 2, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 1, 3, 4, 0, 1,
	1, 1, 1, 1, 2, 1, 2, 4, 2, 4,
	2, 1, 2, 1, 1, 1, 1, 1, 1, 3,
	2, 2, 1, 3, 3, 1, 1, 1, 3, 1,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 4, 0, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 4, 1,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 3, 3, 2, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 1, 3, 3, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 1, 2, 3, 2, 1,
	2, 1, 1, 1, 4, 2, 1, 1, 2, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 2, 4, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 4, 2, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 1, 1, 1, 3, 4, 2,
	1, 2, 3, 5, 3, 3, 3, 3, 3, 1,
	5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -95, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, -46, -47,
	29, -65, -109, -113, -75, -120, -124, -118, -84, 62,
	-59, -101, -105, -80, -54, -48, 31, 32, 28, 27,
	33, 36, 44, 45, 22, 39, 43, 30, 54, 57,
	59, -5, 5, 10, 8, -5, -5, -5, -5, 8,
	10, 8, -5, 11, 4, 4, 8, 8, 9, -28,
	8, 4, 8, -110, 8, 8, 8, 8, 8, 10,
	8, 10, 8, -5, 8, 8, 8, 8, 10, 8,
	10, 8, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, -5, 4, 4, 10, 56, -20,
	-21, -14, -15, 10, 10, -126, 10, 8, 10, -23,
	-24, -18, 14, -14, -15, -26, -27, 14, -14, -15,
	10, -96, -97, -98, -14, -15, -52, -99, -100, 60,
	24, 25, -66, -67, -68, -14, -15, -69, -70, 17,
	15, -111, -112, -14, -15, -30, -114, -115, -116, -14,
	-15, -52, -53, -91, -92, -78, -79, -117, 23, -30,
	61, 37, 38, 21, 34, 20, -76, -64, -77, -14,
	-15, -52, -53, -78, -79, -30, -121, -122, -123, -68,
	-14, -15, -52, -41, -53, -78, -91, -92, -79, -69,
	-121, -119, -14, -15, -52, -41, -53, -78, -79, -85,
	-86, -87, -14, -15, -52, -53, -88, -45, -89, 52,
	10, -60, -14, -15, -61, -52, -53, -62, -63, 40,
	-102, -103, -104, -14, -15, -52, -99, -100, -106, -107,
	-108, -14, -15, -52, -30, -81, -82, -83, -14, -15,
	-52, -53, -35, -36, -37, -38, -39, -40, -42, -62,
	-43, -44, -55, -56, -57, -14, -15, -58, -52, 58,
	-49, -50, -51, -14, -15, -52, 5, 9, -21, 7,
	-24, 9, 11, -27, 9, 11, 9, -98, -31, -32,
	-30, -31, -5, 8, 8, 9, -67, -71, 10, 8,
	-4, -5, 6, 4, 9, -112, 9, -116, -5, -5,
	-3, 53, 6, -3, -2, 48, 49, -2, -5, 9,
	-77, 9, -123, 9, 9, 9, -87, 10, 8, -5,
	9, -62, 8, 4, 9, -104, -31, -31, 9, -108,
	9, -83, 9, -57, 4, 9, -51, -4, 10, 10,
	9, -30, 9, 10, -72, 18, -73, -58, -74, 46,
	-125, -68, 19, 10, 10, 10, 10, 10, 10, 10,
	10, 10, -93, -94, -90, -14, -15, -52, -69, -78,
	-79, -41, -91, -92, -64, 9, 9, 10, -126, 9,
	-5, -125, -68, -5, 4, 9, -90, 9, 9, 10,
	10, 10, 8, -1, 47, 9, -3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 272, 16, 0, 0, 0, 0, 0, 0,
	37, 0, 0, 15, 22, 31, 2, 3, 1, 42,
	184, 183, 0, 216, 0, 0, 123, 0, 0, 244,
	0, 155, 157, 0, 88, 197, 208, 136, 75, 78,
	62, 65, 100, 218, 225, 122, 256, 279, 253, 254,
	154, 94, 196, 207, 134, 77, 64, 6, 0, 0,
	18, 20, 21, 287, 288, 285, 289, 0, 286, 0,
	23, 26, 0, 28, 29, 0, 32, 0, 35, 36,
	25, 0, 185, 186, 188, 189, 190, 43, 43, 0,
	0, 0, 0, 101, 103, 104, 105, 106, 0, 0,
	0, 0, 219, 221, 222, 223, 0, 226, 227, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 0, 242,
	0, 0, 0, 0, 0, 0, 0, 124, 125, 127,
	128, 129, 130, 131, 132, 133, 0, 257, 258, 260,
	261, 262, 263, 264, 265, 266, 267, 268, 269, 270,
	0, 0, 246, 247, 248, 249, 250, 251, 252, 0,
	158, 159, 161, 162, 163, 164, 165, 166, 0, 0,
	72, 0, 89, 90, 91, 92, 93, 95, 0, 0,
	0, 198, 199, 201, 202, 203, 43, 43, 0, 209,
	210, 212, 213, 214, 215, 0, 137, 138, 140, 141,
	142, 143, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 153, 0, 79, 80, 82, 83, 84, 85, 0,
	0, 66, 67, 69, 70, 71, 273, 17, 19, 0,
	24, 30, 0, 33, 38, 0, 182, 187, 0, 44,
	60, 0, 0, 193, 194, 99, 102, 110, 112, 0,
	0, 107, 108, 111, 217, 220, 224, 228, 0, 0,
	0, 0, 274, 0, 0, 275, 276, 0, 0, 121,
	126, 255, 259, 278, 245, 156, 160, 177, 0, 167,
	87, 96, 0, 98, 195, 200, 0, 0, 206, 211,
	135, 139, 76, 81, 0, 63, 68, 0, 27, 34,
	191, 61, 192, 73, 0, 0, 115, 116, 117, 0,
	280, 119, 0, 109, 241, 74, 229, 230, 231, 277,
	271, 243, 0, 179, 180, 168, 169, 170, 171, 172,
	173, 174, 175, 176, 0, 204, 205, 86, 0, 113,
	0, 281, 120, 0, 0, 178, 181, 97, 290, 114,
	118, 282, 0, 0, 0, 283, 0, 284,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62,
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
			l := yylex.(*lexer)
			if l.parent != nil {
				l.Error("expected submodule for include")
				goto ret1
			}
			yylex.(*lexer).stack.Push(meta.NewModule(yyDollar[2].token, l.featureSet))
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:152
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
		//line parser.y:166
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:181
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:188
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:191
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:204
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:215
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:228
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:233
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:249
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:252
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:290
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:293
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:298
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:317
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:324
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:331
		{
			if set(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:338
		{
			pop(yylex)
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:341
		{
			pop(yylex)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:346
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:366
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:376
		{
			pop(yylex)
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:389
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:401
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:406
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:416
		{
			pop(yylex)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:421
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:438
		{
			yyVAL.token = yyDollar[1].token
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:439
		{
			yyVAL.token = yyDollar[1].token
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:442
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:452
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:459
		{
			pop(yylex)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:462
		{
			pop(yylex)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:467
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:475
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:489
		{
			pop(yylex)
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:518
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:525
		{
			pop(yylex)
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:553
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:560
		{
			pop(yylex)
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:563
		{
			pop(yylex)
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:583
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:602
		{
			pop(yylex)
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:605
		{
			pop(yylex)
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:617
		{
			pop(yylex)
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:622
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:639
		{
			pop(yylex)
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:642
		{
			pop(yylex)
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:647
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:654
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:664
		{
			pop(yylex)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:669
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:686
		{
			pop(yylex)
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:689
		{
			pop(yylex)
		}
	case 206:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:697
		{
			pop(yylex)
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:702
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:724
		{
			pop(yylex)
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:734
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 224:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:752
		{
			pop(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:757
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:772
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:777
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:784
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 243:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:804
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:811
		{
			pop(yylex)
		}
	case 245:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:814
		{
			pop(yylex)
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:829
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 254:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:834
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 255:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:844
		{
			pop(yylex)
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:849
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 271:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:877
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:884
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 273:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:887
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:892
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:902
		{
			yyVAL.boolean = true
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:903
		{
			yyVAL.boolean = false
		}
	case 277:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:906
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 278:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:916
		{
			pop(yylex)
		}
	case 279:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:921
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 282:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:932
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 283:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:937
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 284:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:944
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 285:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:949
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 286:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:956
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 287:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:963
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 288:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:970
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
