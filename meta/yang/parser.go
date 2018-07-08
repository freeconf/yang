//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta"
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

func push(l yyLexer, m interface{}) bool {
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
	return l.(*lexer).stack.Peek().(meta.Meta)
}

//line parser.y:58
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
const kywd_min_elements = 57379
const kywd_choice = 57380
const kywd_case = 57381
const kywd_import = 57382
const kywd_include = 57383
const kywd_action = 57384
const kywd_anyxml = 57385
const kywd_anydata = 57386
const kywd_path = 57387
const kywd_value = 57388
const kywd_true = 57389
const kywd_false = 57390
const kywd_contact = 57391
const kywd_organization = 57392
const kywd_refine = 57393
const kywd_unbounded = 57394
const kywd_augment = 57395
const kywd_submodule = 57396
const kywd_str_plus = 57397
const kywd_identity = 57398
const kywd_base = 57399
const kywd_feature = 57400
const kywd_if_feature = 57401
const kywd_when = 57402
const kywd_must = 57403
const kywd_yang_version = 57404
const kywd_range = 57405
const kywd_extension = 57406
const kywd_argument = 57407
const kywd_yin_element = 57408
const kywd_pattern = 57409
const kywd_units = 57410
const kywd_fraction_digits = 57411
const kywd_status = 57412

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
	"kywd_yang_version",
	"kywd_range",
	"kywd_extension",
	"kywd_argument",
	"kywd_yin_element",
	"kywd_pattern",
	"kywd_units",
	"kywd_fraction_digits",
	"kywd_status",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:1102

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 974

var yyAct = [...]int{

	233, 491, 370, 12, 231, 452, 12, 441, 251, 424,
	374, 323, 417, 250, 155, 42, 230, 253, 283, 127,
	237, 19, 328, 242, 19, 320, 41, 128, 236, 296,
	40, 179, 39, 38, 315, 87, 288, 303, 178, 265,
	43, 241, 210, 204, 192, 174, 185, 166, 37, 91,
	92, 93, 94, 150, 248, 98, 239, 160, 36, 345,
	35, 23, 497, 148, 3, 15, 232, 375, 376, 11,
	346, 34, 11, 33, 285, 476, 157, 475, 156, 501,
	485, 338, 23, 24, 23, 162, 496, 27, 118, 486,
	23, 152, 356, 4, 465, 23, 164, 201, 169, 200,
	157, 176, 156, 188, 24, 448, 24, 148, 142, 196,
	206, 212, 24, 244, 244, 421, 257, 24, 267, 216,
	277, 290, 298, 305, 215, 322, 197, 330, 207, 221,
	238, 218, 193, 252, 252, 190, 262, 180, 317, 217,
	300, 249, 249, 447, 261, 246, 246, 148, 259, 316,
	152, 199, 293, 314, 282, 489, 313, 151, 271, 292,
	164, 161, 163, 194, 168, 372, 169, 175, 446, 187,
	254, 312, 445, 23, 176, 195, 205, 211, 444, 243,
	243, 311, 256, 310, 266, 188, 276, 289, 297, 304,
	279, 321, 196, 329, 309, 24, 308, 488, 487, 480,
	148, 438, 23, 335, 206, 449, 437, 495, 348, 197,
	212, 371, 439, 340, 411, 193, 151, 337, 216, 344,
	361, 234, 207, 215, 24, 161, 163, 373, 221, 414,
	218, 352, 168, 274, 170, 381, 357, 377, 217, 347,
	175, 244, 148, 148, 23, 158, 194, 364, 367, 238,
	148, 187, 366, 23, 167, 378, 379, 148, 195, 23,
	389, 252, 388, 200, 235, 267, 24, 225, 201, 249,
	205, 425, 436, 246, 148, 24, 211, 410, 148, 148,
	226, 24, 383, 223, 224, 154, 332, 171, 290, 325,
	148, 180, 482, 23, 23, 201, 298, 200, 393, 427,
	470, 390, 153, 305, 387, 271, 180, 243, 60, 467,
	464, 325, 147, 392, 300, 24, 24, 426, 317, 293,
	322, 435, 177, 434, 397, 401, 292, 463, 330, 316,
	148, 266, 213, 314, 245, 245, 313, 258, 413, 268,
	403, 280, 291, 299, 306, 405, 324, 148, 331, 199,
	408, 312, 398, 399, 289, 419, 409, 148, 369, 394,
	368, 311, 297, 310, 359, 393, 358, 23, 412, 304,
	443, 429, 407, 404, 309, 214, 308, 247, 247, 354,
	260, 353, 269, 431, 281, 402, 321, 307, 400, 24,
	454, 126, 350, 125, 329, 177, 396, 124, 461, 123,
	391, 498, 334, 460, 23, 23, 273, 456, 80, 386,
	458, 462, 385, 238, 180, 228, 384, 419, 457, 382,
	380, 418, 459, 117, 466, 116, 24, 24, 23, 468,
	365, 213, 469, 471, 363, 429, 442, 477, 495, 181,
	182, 443, 23, 162, 115, 27, 114, 431, 481, 102,
	24, 101, 454, 472, 473, 474, 453, 483, 351, 97,
	461, 96, 245, 478, 24, 460, 90, 381, 89, 456,
	343, 349, 458, 146, 214, 180, 145, 493, 122, 23,
	457, 238, 121, 418, 459, 341, 268, 225, 23, 167,
	120, 493, 499, 119, 113, 112, 111, 110, 500, 109,
	226, 24, 108, 100, 95, 247, 336, 442, 372, 291,
	24, 88, 23, 6, 201, 360, 200, 299, 453, 86,
	225, 88, 129, 355, 306, 342, 180, 228, 60, 269,
	339, 333, 103, 226, 24, 99, 223, 224, 479, 406,
	395, 324, 362, 492, 144, 143, 141, 494, 140, 331,
	67, 139, 85, 138, 7, 23, 47, 492, 27, 180,
	228, 60, 137, 136, 76, 135, 134, 307, 199, 71,
	70, 50, 79, 68, 69, 72, 133, 24, 73, 132,
	131, 77, 130, 48, 49, 78, 74, 75, 107, 106,
	105, 104, 21, 22, 84, 83, 80, 490, 433, 81,
	57, 82, 240, 56, 60, 28, 255, 51, 67, 58,
	219, 455, 7, 23, 47, 209, 27, 208, 54, 203,
	202, 53, 76, 295, 294, 63, 287, 71, 70, 50,
	79, 68, 69, 72, 286, 24, 73, 62, 173, 77,
	172, 48, 49, 78, 74, 75, 29, 451, 450, 272,
	21, 22, 270, 264, 80, 67, 263, 81, 59, 82,
	23, 302, 60, 28, 301, 51, 64, 227, 225, 76,
	220, 229, 55, 455, 71, 70, 432, 79, 68, 69,
	72, 226, 24, 73, 223, 224, 77, 430, 428, 423,
	78, 74, 75, 422, 198, 191, 52, 67, 284, 484,
	278, 80, 23, 275, 81, 61, 82, 180, 228, 60,
	225, 76, 319, 318, 65, 440, 71, 70, 222, 79,
	68, 69, 72, 226, 24, 73, 327, 326, 77, 66,
	420, 416, 78, 74, 75, 415, 189, 186, 184, 67,
	183, 30, 46, 80, 23, 45, 81, 44, 82, 180,
	228, 60, 225, 76, 32, 31, 165, 26, 71, 70,
	159, 79, 68, 69, 72, 226, 24, 73, 25, 149,
	77, 20, 18, 17, 78, 74, 75, 16, 14, 13,
	10, 67, 9, 8, 5, 80, 23, 2, 81, 1,
	82, 180, 228, 60, 0, 76, 0, 0, 0, 0,
	71, 70, 0, 79, 68, 69, 72, 0, 24, 73,
	0, 0, 77, 285, 0, 0, 78, 74, 75, 0,
	0, 0, 0, 67, 0, 0, 0, 80, 23, 0,
	81, 0, 82, 180, 228, 60, 0, 76, 0, 0,
	0, 0, 71, 70, 0, 79, 68, 69, 72, 0,
	24, 73, 0, 0, 77, 0, 0, 0, 78, 74,
	75, 0, 67, 0, 0, 0, 0, 23, 0, 80,
	0, 0, 81, 0, 82, 180, 76, 60, 0, 0,
	0, 71, 70, 0, 79, 68, 69, 72, 0, 24,
	73, 0, 0, 77, 0, 67, 0, 78, 74, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 80, 76,
	0, 81, 0, 82, 71, 70, 60, 79, 68, 69,
	72, 0, 0, 73, 0, 23, 77, 0, 0, 0,
	78, 74, 75, 0, 76, 0, 0, 0, 0, 71,
	70, 80, 79, 0, 81, 72, 82, 24, 73, 60,
	0, 77, 285, 0, 0, 78, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 180, 228,
}
var yyPact = [...]int{

	39, -1000, 601, 591, 590, 543, -1000, 506, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	458, 506, 506, 506, 506, 496, 451, 506, 530, 495,
	441, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 527, 587, 586,
	585, 584, 494, 491, 489, 488, 487, 486, 436, 415,
	506, 485, 482, 474, 470, 389, 383, 516, 578, 576,
	575, 572, 562, 561, 559, 558, 549, 547, 544, 542,
	506, 541, 540, 468, 465, -1000, -1000, 302, -1000, -1000,
	282, 292, 275, 92, 235, 430, -1000, 241, 224, 277,
	416, -1000, 70, -1000, -1000, -1000, -1000, -1000, 281, 855,
	648, 732, 500, 500, -1000, 467, -1000, 355, 223, 774,
	416, 816, 913, -1000, 232, -1000, 78, 276, 8, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 8, -1000, -1000, -1000, -1000, -1000, 526, 393,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 499, -1000, 72,
	-1000, -1000, 525, -1000, -1000, 476, -1000, 520, -1000, -1000,
	-1000, -1000, 461, 416, -1000, -1000, -1000, -1000, 888, 888,
	506, 463, 384, 449, 70, -1000, -1000, -1000, -1000, 371,
	518, 83, -1000, -1000, -1000, -1000, -1000, -1000, 356, 510,
	516, 538, 425, 855, -1000, -1000, -1000, -1000, 421, 648,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	506, -1000, 350, 159, 502, 20, 20, 506, 506, 411,
	732, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 410,
	500, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 407, 403, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 400, 355, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 252, 506, -1000, 391, -1000, -1000, 35, 888,
	-1000, -1000, -1000, -1000, 351, 536, 387, 416, -1000, -1000,
	-1000, -1000, 888, 888, 379, 816, -1000, -1000, -1000, -1000,
	-1000, 376, 913, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 364, 232,
	-1000, -1000, -1000, -1000, -1000, 535, 363, 78, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 516, -1000, -1000, 267,
	-1000, -1000, 204, -1000, -1000, 359, 888, 329, 219, -1000,
	-1000, -1000, -1000, -1000, 49, -1000, -1000, -1000, -1000, 254,
	196, 191, -1000, -1000, -1000, -1000, -1000, 202, -1000, 282,
	168, 162, -1000, 158, 133, -1000, -1000, 95, 195, 8,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 247,
	8, -1000, -1000, -1000, 732, -1000, -1000, -1000, 318, 301,
	-1000, -1000, -1000, -1000, -1000, -1000, 84, -1000, -1000, 68,
	-1000, -1000, -1000, -1000, -1000, 300, 49, -1000, -1000, -1000,
	-1000, 20, 291, 254, -1000, 506, 506, 506, -1000, -1000,
	-1000, -1000, -1000, 67, 502, 506, 534, -1000, -1000, -1000,
	190, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	283, 247, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 690, -1000, -1000, -1000, 71, -1000, -1000, 79,
	-1000, -1000, 188, 187, 145, -1000, 161, 76, 52, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	392, -1000, -1000, -1000, -1000, 502, -1000, -1000, -1000, -1000,
	69, -1000,
}
var yyPgo = [...]int{

	0, 10, 2, 19, 27, 789, 787, 784, 513, 783,
	782, 780, 66, 0, 779, 778, 65, 777, 773, 772,
	18, 771, 769, 53, 768, 760, 57, 757, 756, 47,
	59, 70, 755, 754, 73, 71, 60, 58, 48, 33,
	32, 30, 26, 15, 40, 747, 745, 742, 741, 740,
	738, 46, 737, 736, 735, 731, 12, 730, 729, 727,
	726, 22, 221, 718, 264, 715, 7, 714, 713, 712,
	25, 11, 705, 703, 700, 34, 698, 16, 696, 695,
	44, 23, 54, 17, 694, 693, 689, 9, 688, 687,
	676, 672, 671, 4, 28, 20, 666, 664, 661, 37,
	658, 656, 653, 39, 652, 649, 5, 13, 8, 648,
	647, 646, 640, 638, 45, 38, 31, 637, 634, 626,
	36, 625, 624, 623, 29, 621, 620, 619, 43, 618,
	617, 615, 42, 610, 609, 606, 603, 56, 602, 41,
	600, 598, 597, 1, 547, 14,
}
var yyR1 = [...]int{

	0, 5, 6, 6, 7, 7, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 21,
	9, 9, 22, 22, 23, 23, 24, 25, 25, 16,
	26, 26, 26, 26, 14, 27, 28, 28, 29, 29,
	29, 15, 15, 30, 30, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 31, 31, 19, 19, 48, 49, 49, 50, 50,
	51, 51, 51, 52, 52, 53, 54, 54, 55, 55,
	56, 56, 56, 57, 46, 46, 58, 59, 59, 60,
	60, 61, 61, 61, 40, 62, 63, 64, 64, 65,
	65, 66, 66, 45, 45, 67, 68, 68, 69, 69,
	70, 70, 70, 70, 71, 41, 73, 73, 73, 73,
	73, 73, 73, 72, 74, 74, 75, 76, 32, 78,
	79, 79, 80, 80, 80, 80, 80, 3, 3, 83,
	81, 81, 84, 85, 85, 86, 86, 87, 87, 87,
	87, 87, 87, 87, 87, 89, 90, 35, 91, 92,
	92, 77, 77, 93, 93, 93, 93, 93, 93, 93,
	96, 44, 97, 97, 98, 98, 99, 99, 99, 99,
	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
	100, 39, 39, 101, 101, 102, 102, 103, 103, 103,
	103, 103, 103, 105, 106, 106, 106, 106, 106, 106,
	106, 106, 106, 104, 104, 109, 110, 110, 18, 111,
	112, 112, 113, 113, 114, 114, 114, 114, 114, 115,
	116, 42, 117, 118, 118, 119, 119, 120, 120, 120,
	120, 120, 43, 121, 122, 122, 123, 123, 124, 124,
	124, 124, 33, 125, 126, 126, 127, 127, 128, 128,
	128, 34, 129, 130, 131, 131, 107, 107, 108, 132,
	132, 132, 132, 132, 132, 132, 132, 132, 132, 132,
	133, 38, 38, 135, 135, 135, 135, 135, 135, 135,
	134, 134, 36, 136, 137, 138, 138, 139, 139, 139,
	139, 139, 139, 139, 139, 139, 139, 139, 139, 95,
	4, 4, 2, 1, 1, 94, 37, 140, 88, 88,
	141, 142, 142, 143, 143, 143, 144, 12, 13, 10,
	11, 17, 82, 145, 145, 47,
}
var yyR2 = [...]int{

	0, 3, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 1, 2, 1, 1, 2, 1, 2, 3,
	1, 3, 1, 1, 4, 2, 1, 2, 3, 1,
	1, 2, 4, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 3, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 1, 3, 3, 2, 2, 4, 1,
	2, 1, 1, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 1, 3, 4, 0, 1, 1, 1,
	1, 1, 1, 2, 1, 2, 4, 2, 4, 2,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 4, 2, 0, 1, 1, 2, 3, 3, 3,
	1, 1, 1, 1, 1, 3, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 4, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 4, 1, 1, 2, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 3, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 4, 2, 1, 1, 2, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 2, 4, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 2, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 3, 1, 1, 1, 3, 4, 2, 2, 4,
	2, 1, 2, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 1, 5, 3,
}
var yyChk = [...]int{

	-1000, -5, -6, 25, 54, -7, -8, 11, -9, -10,
	-11, -12, -13, -14, -15, -16, -17, -18, -19, -20,
	-21, 49, 50, 12, 34, -24, -27, 15, 62, -111,
	-48, -32, -33, -34, -35, -36, -37, -38, -39, -40,
	-41, -42, -43, -44, -45, -46, -47, 13, 40, 41,
	28, 64, -78, -125, -129, -91, -136, -140, -134, -100,
	61, -72, -117, -121, -96, -67, -58, 7, 30, 31,
	27, 26, 32, 35, 43, 44, 21, 38, 42, 29,
	53, 56, 58, 4, 4, 9, -8, -4, 5, 10,
	8, -4, -4, -4, -4, 8, 10, 8, -4, 5,
	8, 10, 8, 5, 4, 4, 4, 4, 8, 8,
	8, 8, 8, 8, 10, 8, 10, 8, -4, 8,
	8, 8, 8, 10, 8, 10, 8, -3, -4, 6,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, -4, 4, 4, 8, 8, 10, 55, -22,
	-23, -12, -13, 10, 10, -145, 10, 8, 10, -25,
	-26, -16, 13, -12, -13, -28, -29, 13, -12, -13,
	10, 10, -112, -113, -114, -12, -13, -62, -115, -116,
	59, 23, 24, -49, -50, -51, -52, -12, -13, -53,
	65, -79, -80, -81, -82, -12, -13, -83, -84, 68,
	16, 14, -126, -127, -128, -12, -13, -20, -130, -131,
	-132, -12, -13, -62, -64, -107, -108, -94, -95, -133,
	22, -20, -63, 36, 37, 20, 33, 19, 60, -92,
	-77, -93, -12, -13, -62, -64, -94, -95, -20, -137,
	-138, -139, -81, -12, -13, -62, -40, -64, -82, -94,
	-107, -108, -95, -83, -137, -135, -12, -13, -62, -40,
	-64, -94, -95, -101, -102, -103, -12, -13, -62, -64,
	-104, -44, -105, 51, 10, -73, -12, -13, -74, -31,
	-62, -64, -75, -20, -76, 39, -118, -119, -120, -12,
	-13, -62, -115, -116, -122, -123, -124, -12, -13, -62,
	-20, -97, -98, -99, -12, -13, -62, -64, -34, -35,
	-36, -37, -38, -39, -41, -75, -42, -43, -68, -69,
	-70, -12, -13, -71, -62, 57, -59, -60, -61, -12,
	-13, -62, 10, 5, 9, -23, 7, -26, 9, 5,
	-29, 9, 5, 9, -114, -30, -31, -30, -4, 8,
	8, 9, -51, 10, 8, 5, 9, -80, 10, 8,
	5, -3, 4, 9, -128, 9, -132, -4, 10, 8,
	-2, 52, 6, -2, -1, 47, 48, -1, -4, -4,
	9, -93, 9, -139, 9, 9, 9, -103, 10, 8,
	-4, 9, -75, -20, 8, 4, 9, -120, -30, -30,
	9, -124, 9, -99, 9, -70, 4, 9, -61, -3,
	10, 10, 9, 9, 10, -54, -55, -56, -12, -13,
	-57, 66, -85, -86, -87, 17, 63, 45, -88, -71,
	-89, -81, -90, -141, 69, 67, 18, 10, 10, 10,
	-65, -66, -12, -13, 10, 10, 10, 10, 10, 10,
	-109, -110, -106, -12, -13, -62, -83, -94, -95, -40,
	-107, -108, -77, 9, 9, 10, -145, 9, -56, -1,
	9, -87, -4, -4, -4, 10, 8, -2, -4, 4,
	9, -66, 9, -106, 9, 9, 10, 10, 10, 10,
	-142, -143, -12, -13, -144, 46, 10, 10, 9, -143,
	-2, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 5, 0, 310, 20,
	0, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	220, 63, 66, 19, 26, 35, 219, 65, 0, 254,
	0, 159, 0, 0, 281, 0, 191, 193, 0, 116,
	233, 244, 172, 103, 106, 84, 87, 0, 137, 138,
	129, 253, 262, 158, 293, 317, 290, 291, 190, 123,
	232, 243, 170, 105, 86, 2, 3, 6, 0, 0,
	22, 24, 25, 329, 330, 327, 333, 0, 328, 0,
	27, 30, 0, 32, 33, 0, 36, 0, 39, 40,
	29, 331, 0, 221, 222, 224, 225, 226, 43, 43,
	0, 0, 0, 0, 67, 68, 70, 71, 72, 0,
	0, 0, 130, 132, 133, 134, 135, 136, 0, 0,
	0, 0, 0, 255, 256, 258, 259, 260, 0, 263,
	264, 269, 270, 271, 272, 273, 274, 275, 276, 277,
	0, 279, 0, 0, 0, 0, 0, 0, 0, 0,
	160, 161, 163, 164, 165, 166, 167, 168, 169, 0,
	294, 295, 297, 298, 299, 300, 301, 302, 303, 304,
	305, 306, 307, 308, 0, 0, 283, 284, 285, 286,
	287, 288, 289, 0, 194, 195, 197, 198, 199, 200,
	201, 202, 0, 0, 94, 0, 117, 118, 119, 120,
	121, 122, 124, 61, 0, 0, 0, 234, 235, 237,
	238, 239, 43, 43, 0, 245, 246, 248, 249, 250,
	251, 0, 173, 174, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 0, 107,
	108, 110, 111, 112, 113, 0, 0, 88, 89, 91,
	92, 93, 335, 311, 21, 23, 0, 28, 34, 0,
	37, 42, 0, 218, 223, 0, 44, 0, 0, 229,
	230, 64, 69, 73, 76, 75, 128, 131, 140, 143,
	0, 0, 142, 252, 257, 261, 265, 0, 97, 0,
	0, 0, 312, 0, 0, 313, 314, 0, 0, 96,
	157, 162, 292, 296, 316, 282, 192, 196, 213, 0,
	203, 115, 125, 62, 0, 127, 231, 236, 0, 0,
	242, 247, 171, 175, 104, 109, 0, 85, 90, 0,
	31, 38, 227, 228, 95, 0, 77, 78, 80, 81,
	82, 0, 0, 144, 145, 0, 0, 0, 150, 151,
	152, 153, 154, 0, 0, 0, 0, 332, 139, 278,
	0, 99, 101, 102, 266, 267, 268, 315, 309, 280,
	0, 215, 216, 204, 205, 206, 207, 208, 209, 210,
	211, 212, 0, 240, 241, 114, 0, 74, 79, 0,
	141, 146, 0, 0, 0, 318, 0, 0, 0, 320,
	98, 100, 214, 217, 126, 334, 83, 147, 148, 149,
	0, 321, 323, 324, 325, 0, 155, 156, 319, 322,
	0, 326,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70,
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
		//line parser.y:149
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
		//line parser.y:157
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
		//line parser.y:171
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:190
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:197
		{
			pop(yylex)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:200
		{
			pop(yylex)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:213
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:224
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:237
		{
			pop(yylex)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:242
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:258
		{
			pop(yylex)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:261
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:291
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:294
		{
			pop(yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:299
		{
			if push(yylex, meta.NewExtension(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:318
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			pop(yylex)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:326
		{
			if push(yylex, meta.NewExtensionArg(peek(yylex).(*meta.Extension), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:345
		{
			if set(yylex, meta.SetYinElement(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:352
		{
			pop(yylex)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:355
		{
			pop(yylex)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:360
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:379
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:386
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:393
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:400
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:403
		{
			pop(yylex)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:415
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:418
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:423
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:443
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:453
		{
			pop(yylex)
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:467
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:479
		{
			pop(yylex)
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:484
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:491
		{
			pop(yylex)
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:496
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:513
		{
			yyVAL.token = yyDollar[1].token
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:514
		{
			yyVAL.token = yyDollar[1].token
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:517
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:524
		{
			pop(yylex)
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:527
		{
			pop(yylex)
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:532
		{
			if push(yylex, meta.NewType(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:546
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:555
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetValueRange(r)) {
				goto ret1
			}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:564
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:577
		{
			if set(yylex, meta.SetFractionDigits(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:584
		{
			if set(yylex, meta.SetPattern(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:591
		{
			pop(yylex)
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:596
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:620
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:627
		{
			pop(yylex)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:655
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:662
		{
			pop(yylex)
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:665
		{
			pop(yylex)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:685
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:704
		{
			pop(yylex)
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:707
		{
			pop(yylex)
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:719
		{
			pop(yylex)
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:724
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:741
		{
			pop(yylex)
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:744
		{
			pop(yylex)
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:749
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:756
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:766
		{
			pop(yylex)
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:771
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:788
		{
			pop(yylex)
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:791
		{
			pop(yylex)
		}
	case 242:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:799
		{
			pop(yylex)
		}
	case 243:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:804
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 252:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:825
		{
			pop(yylex)
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:830
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 261:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:849
		{
			pop(yylex)
		}
	case 262:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:854
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 266:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:869
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 267:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:874
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 268:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:881
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 280:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:901
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 281:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:908
		{
			pop(yylex)
		}
	case 282:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:911
		{
			pop(yylex)
		}
	case 290:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:926
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 291:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:931
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 292:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:938
		{
			pop(yylex)
		}
	case 293:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:943
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 309:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:972
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 310:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:979
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 311:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:982
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 312:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:987
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 313:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:997
		{
			yyVAL.boolean = true
		}
	case 314:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:998
		{
			yyVAL.boolean = false
		}
	case 315:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1001
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 316:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1011
		{
			pop(yylex)
		}
	case 317:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1016
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 318:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1023
		{
			pop(yylex)
		}
	case 319:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1026
		{
			pop(yylex)
		}
	case 320:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1031
		{
			if push(yylex, meta.NewEnum(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 326:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1046
		{
			if set(yylex, meta.SetEnumValue(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 327:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1053
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 328:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1060
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 329:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1067
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 330:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1074
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 331:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1081
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 332:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1088
		{
			if set(yylex, meta.SetUnits(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 335:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1099
		{

		}
	}
	goto yystack /* stack new state and value */
}
