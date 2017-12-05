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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:1079

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 942

var yyAct = [...]int{

	365, 231, 248, 229, 12, 432, 443, 12, 318, 419,
	412, 155, 369, 228, 323, 127, 247, 315, 235, 42,
	298, 250, 291, 39, 41, 40, 340, 234, 179, 38,
	37, 239, 36, 35, 310, 262, 240, 34, 178, 43,
	208, 33, 283, 202, 192, 185, 174, 160, 15, 166,
	150, 148, 237, 23, 23, 3, 23, 482, 23, 280,
	199, 483, 198, 23, 181, 182, 223, 230, 370, 371,
	11, 23, 485, 11, 367, 24, 24, 474, 24, 224,
	24, 456, 221, 222, 4, 24, 439, 479, 157, 478,
	156, 438, 152, 24, 157, 437, 156, 164, 436, 169,
	180, 473, 176, 23, 188, 180, 226, 60, 416, 190,
	195, 204, 210, 214, 242, 242, 320, 254, 180, 264,
	366, 274, 285, 293, 300, 24, 317, 213, 325, 216,
	196, 249, 249, 232, 259, 148, 244, 244, 215, 256,
	246, 246, 312, 258, 161, 193, 480, 311, 309, 288,
	180, 152, 308, 307, 278, 306, 305, 268, 151, 287,
	304, 164, 435, 163, 303, 168, 251, 169, 175, 429,
	187, 477, 388, 476, 23, 176, 194, 203, 209, 468,
	241, 241, 23, 253, 475, 263, 188, 273, 284, 292,
	299, 148, 316, 195, 324, 405, 24, 440, 430, 404,
	330, 280, 327, 204, 24, 409, 343, 332, 161, 210,
	214, 171, 470, 196, 356, 335, 148, 151, 148, 271,
	339, 180, 226, 368, 213, 461, 216, 163, 193, 148,
	348, 170, 376, 168, 177, 215, 353, 372, 458, 455,
	242, 175, 148, 148, 211, 359, 243, 243, 361, 255,
	148, 265, 187, 276, 286, 294, 301, 249, 319, 194,
	326, 23, 244, 264, 148, 128, 246, 454, 408, 203,
	378, 406, 401, 87, 236, 209, 148, 19, 398, 329,
	19, 396, 23, 24, 285, 394, 390, 91, 92, 93,
	94, 158, 293, 98, 384, 346, 383, 382, 386, 300,
	270, 268, 80, 154, 24, 381, 241, 177, 180, 226,
	387, 288, 23, 395, 392, 393, 317, 312, 397, 380,
	223, 287, 311, 309, 325, 391, 118, 308, 307, 263,
	306, 305, 399, 224, 24, 304, 148, 402, 379, 303,
	199, 211, 377, 420, 427, 375, 142, 403, 148, 233,
	284, 360, 414, 153, 147, 364, 345, 363, 292, 180,
	226, 60, 358, 347, 424, 299, 434, 338, 146, 23,
	167, 422, 243, 145, 331, 333, 122, 121, 23, 162,
	120, 27, 316, 320, 205, 219, 445, 452, 119, 421,
	324, 24, 425, 428, 113, 265, 295, 112, 148, 148,
	24, 451, 453, 449, 23, 162, 447, 27, 450, 355,
	111, 354, 448, 414, 352, 457, 286, 23, 413, 199,
	110, 198, 459, 350, 294, 349, 24, 424, 462, 460,
	109, 301, 433, 434, 108, 100, 23, 469, 199, 24,
	198, 126, 95, 125, 445, 452, 344, 367, 319, 471,
	88, 129, 444, 342, 342, 425, 326, 376, 24, 451,
	212, 449, 245, 245, 447, 257, 450, 266, 88, 277,
	448, 336, 302, 351, 23, 167, 205, 337, 124, 413,
	123, 334, 219, 484, 362, 117, 115, 116, 114, 67,
	466, 373, 374, 102, 23, 101, 24, 328, 97, 433,
	96, 6, 223, 76, 90, 103, 89, 86, 71, 70,
	444, 79, 68, 69, 72, 224, 24, 73, 446, 99,
	77, 400, 389, 357, 78, 74, 75, 144, 143, 141,
	140, 139, 138, 137, 136, 80, 385, 135, 81, 134,
	82, 180, 226, 60, 57, 133, 132, 67, 131, 85,
	130, 7, 23, 47, 107, 27, 106, 212, 105, 104,
	84, 76, 342, 342, 83, 295, 71, 70, 50, 79,
	68, 69, 72, 238, 24, 73, 446, 56, 77, 252,
	48, 49, 78, 74, 75, 58, 217, 207, 245, 21,
	22, 206, 54, 80, 201, 200, 81, 53, 82, 290,
	289, 60, 28, 63, 51, 282, 281, 62, 173, 67,
	172, 266, 29, 7, 23, 47, 407, 27, 442, 441,
	269, 267, 261, 76, 260, 59, 297, 296, 71, 70,
	50, 79, 68, 69, 72, 64, 24, 73, 227, 55,
	77, 426, 48, 49, 78, 74, 75, 302, 423, 418,
	417, 21, 22, 197, 191, 80, 52, 279, 81, 275,
	82, 272, 61, 60, 28, 314, 51, 313, 65, 431,
	220, 322, 321, 66, 415, 411, 410, 67, 189, 186,
	184, 183, 23, 30, 46, 45, 463, 464, 465, 225,
	223, 76, 218, 44, 467, 32, 71, 70, 31, 79,
	68, 69, 72, 224, 24, 73, 221, 222, 77, 341,
	165, 26, 78, 74, 75, 159, 25, 149, 20, 67,
	18, 472, 17, 80, 23, 16, 81, 14, 82, 180,
	226, 60, 223, 76, 13, 10, 9, 8, 71, 70,
	5, 79, 68, 69, 72, 224, 24, 73, 2, 1,
	77, 481, 0, 0, 78, 74, 75, 0, 0, 0,
	0, 67, 0, 0, 0, 80, 23, 0, 81, 0,
	82, 180, 226, 60, 0, 76, 0, 0, 0, 0,
	71, 70, 0, 79, 68, 69, 72, 0, 24, 73,
	0, 0, 77, 0, 0, 0, 78, 74, 75, 0,
	67, 0, 0, 0, 0, 23, 0, 80, 0, 0,
	81, 0, 82, 180, 76, 60, 0, 0, 0, 71,
	70, 0, 79, 68, 69, 72, 0, 24, 73, 0,
	0, 77, 0, 67, 0, 78, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 80, 76, 0, 81,
	0, 82, 71, 70, 60, 79, 68, 69, 72, 0,
	0, 73, 0, 23, 77, 0, 0, 0, 78, 74,
	75, 0, 76, 0, 0, 0, 0, 71, 70, 80,
	79, 0, 81, 72, 82, 24, 73, 60, 0, 77,
	280, 0, 23, 78, 74, 75, 198, 0, 0, 0,
	223, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	180, 226, 0, 224, 24, 0, 221, 222, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 180,
	0, 60,
}
var yyPact = [...]int{

	30, -1000, 602, 560, 556, 540, -1000, 463, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	496, 463, 463, 463, 463, 434, 490, 463, 514, 427,
	485, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 500, 555, 554,
	552, 550, 426, 422, 412, 402, 389, 386, 478, 477,
	463, 380, 372, 369, 368, 470, 433, 445, 546, 544,
	542, 541, 535, 533, 530, 529, 528, 527, 526, 525,
	463, 524, 523, 365, 360, -1000, -1000, 344, -1000, -1000,
	51, 343, 293, 80, 281, 392, -1000, 357, 221, 201,
	41, -1000, 44, -1000, -1000, -1000, -1000, -1000, 424, 793,
	670, 482, 46, 46, -1000, 300, -1000, 249, 209, 162,
	41, 754, 851, -1000, 59, -1000, 91, 192, -4, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -4, -1000, -1000, -1000, -1000, -1000, 492, 270,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 367, -1000, 366,
	-1000, -1000, 476, -1000, -1000, 462, -1000, 472, -1000, -1000,
	-1000, -1000, 358, 41, -1000, -1000, -1000, -1000, 826, 826,
	463, 348, 287, 354, 44, -1000, -1000, -1000, -1000, 415,
	468, 405, -1000, -1000, -1000, -1000, -1000, 401, 445, 519,
	353, 793, -1000, -1000, -1000, -1000, 342, 670, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 463, -1000,
	347, 68, 441, 21, 21, 463, 463, 336, 482, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 333, 46, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 329, 310, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	296, 249, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 286,
	463, -1000, 289, -1000, -1000, 20, -1000, -1000, -1000, 164,
	518, 277, 41, -1000, -1000, -1000, -1000, 826, 826, 276,
	754, -1000, -1000, -1000, -1000, -1000, 272, 851, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 269, 59, -1000, -1000, -1000, -1000, -1000,
	517, 263, 91, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 445, -1000, -1000, 189, -1000, -1000, 185, -1000, -1000,
	262, 826, -1000, 259, 195, -1000, -1000, -1000, -1000, -1000,
	42, -1000, -1000, -1000, -1000, 326, 159, -1000, -1000, -1000,
	-1000, -1000, 188, -1000, 51, 152, 88, -1000, 85, 81,
	-1000, -1000, 76, 187, -4, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 880, -4, -1000, -1000, 482, -1000,
	-1000, -1000, 258, 230, -1000, -1000, -1000, -1000, -1000, -1000,
	71, -1000, -1000, 86, -1000, -1000, -1000, -1000, -1000, -1000,
	229, 42, -1000, -1000, -1000, -1000, 21, 216, 326, -1000,
	463, 463, 463, -1000, -1000, -1000, -1000, 486, 463, -1000,
	-1000, 170, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 203, 880, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 712, -1000, -1000, -1000, 92, -1000, -1000,
	67, -1000, -1000, 174, 163, 161, 79, 136, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 11,
	-1000, 52, 441, -1000, 62, -1000,
}
var yyPgo = [...]int{

	0, 751, 12, 0, 15, 265, 749, 748, 740, 501,
	737, 736, 735, 67, 1, 734, 727, 48, 725, 722,
	720, 274, 718, 717, 50, 716, 715, 47, 711, 710,
	49, 26, 709, 698, 695, 41, 37, 33, 32, 30,
	29, 23, 25, 24, 19, 39, 693, 685, 684, 683,
	681, 680, 45, 679, 678, 676, 675, 10, 674, 673,
	672, 671, 14, 133, 670, 349, 669, 5, 668, 667,
	665, 17, 8, 662, 661, 659, 34, 657, 13, 656,
	654, 44, 36, 21, 653, 650, 649, 9, 648, 641,
	639, 638, 3, 27, 18, 635, 627, 626, 20, 625,
	624, 622, 35, 621, 620, 6, 16, 2, 619, 618,
	612, 610, 608, 46, 38, 28, 607, 606, 605, 42,
	603, 600, 599, 22, 597, 595, 594, 43, 592, 591,
	587, 40, 586, 585, 579, 577, 52, 573, 31, 544,
	11,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 22,
	10, 10, 23, 23, 24, 24, 25, 26, 26, 17,
	27, 27, 27, 27, 15, 28, 29, 29, 30, 30,
	30, 16, 16, 31, 31, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 32, 32, 20, 20, 49, 50, 50, 51, 51,
	52, 52, 52, 53, 53, 54, 55, 55, 56, 56,
	57, 57, 57, 58, 47, 47, 59, 60, 60, 61,
	61, 62, 62, 62, 41, 63, 64, 65, 65, 66,
	66, 67, 67, 46, 46, 68, 69, 69, 70, 70,
	71, 71, 71, 71, 72, 42, 74, 74, 74, 74,
	74, 74, 73, 75, 75, 76, 77, 33, 79, 80,
	80, 81, 81, 81, 81, 4, 4, 83, 82, 82,
	84, 85, 85, 86, 86, 87, 87, 87, 87, 87,
	87, 87, 89, 36, 90, 91, 91, 78, 78, 92,
	92, 92, 92, 92, 92, 92, 95, 45, 96, 96,
	97, 97, 98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 99, 40, 40, 100,
	100, 101, 101, 102, 102, 102, 102, 102, 102, 104,
	105, 105, 105, 105, 105, 105, 105, 105, 105, 103,
	103, 108, 109, 109, 19, 110, 111, 111, 112, 112,
	113, 113, 113, 113, 113, 114, 115, 43, 116, 117,
	117, 118, 118, 119, 119, 119, 119, 119, 44, 120,
	121, 121, 122, 122, 123, 123, 123, 123, 34, 124,
	125, 125, 126, 126, 127, 127, 127, 35, 128, 129,
	130, 130, 106, 106, 107, 131, 131, 131, 131, 131,
	131, 131, 131, 131, 131, 131, 132, 39, 39, 134,
	134, 134, 134, 134, 134, 134, 133, 133, 37, 135,
	136, 137, 137, 138, 138, 138, 138, 138, 138, 138,
	138, 138, 138, 138, 94, 5, 5, 3, 2, 2,
	93, 38, 139, 88, 88, 1, 13, 14, 11, 12,
	18, 140, 140, 48,
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
	1, 1, 2, 1, 2, 4, 2, 4, 2, 1,
	2, 1, 1, 1, 1, 1, 1, 3, 2, 4,
	2, 0, 1, 1, 2, 3, 3, 3, 1, 1,
	1, 1, 3, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 4, 0, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	4, 1, 1, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 3, 3, 2, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 3, 3, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 1, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 4, 2, 1,
	1, 2, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 3, 2, 4, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 2,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 3, 1, 1, 1,
	3, 4, 2, 3, 5, 3, 3, 3, 3, 3,
	3, 1, 5, 3,
}
var yyChk = [...]int{

	-1000, -6, -7, 25, 54, -8, -9, 11, -10, -11,
	-12, -13, -14, -15, -16, -17, -18, -19, -20, -21,
	-22, 49, 50, 12, 34, -25, -28, 15, 62, -110,
	-49, -33, -34, -35, -36, -37, -38, -39, -40, -41,
	-42, -43, -44, -45, -46, -47, -48, 13, 40, 41,
	28, 64, -79, -124, -128, -90, -135, -139, -133, -99,
	61, -73, -116, -120, -95, -68, -59, 7, 30, 31,
	27, 26, 32, 35, 43, 44, 21, 38, 42, 29,
	53, 56, 58, 4, 4, 9, -9, -5, 5, 10,
	8, -5, -5, -5, -5, 8, 10, 8, -5, 5,
	8, 10, 8, 5, 4, 4, 4, 4, 8, 8,
	8, 8, 8, 8, 10, 8, 10, 8, -5, 8,
	8, 8, 8, 10, 8, 10, 8, -4, -5, 6,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, -5, 4, 4, 8, 8, 10, 55, -23,
	-24, -13, -14, 10, 10, -140, 10, 8, 10, -26,
	-27, -17, 13, -13, -14, -29, -30, 13, -13, -14,
	10, 10, -111, -112, -113, -13, -14, -63, -114, -115,
	59, 23, 24, -50, -51, -52, -53, -13, -14, -54,
	65, -80, -81, -82, -13, -14, -83, -84, 16, 14,
	-125, -126, -127, -13, -14, -21, -129, -130, -131, -13,
	-14, -63, -65, -106, -107, -93, -94, -132, 22, -21,
	-64, 36, 37, 20, 33, 19, 60, -91, -78, -92,
	-13, -14, -63, -65, -93, -94, -21, -136, -137, -138,
	-82, -13, -14, -63, -41, -65, -93, -106, -107, -94,
	-83, -136, -134, -13, -14, -63, -41, -65, -93, -94,
	-100, -101, -102, -13, -14, -63, -65, -103, -45, -104,
	51, 10, -74, -13, -14, -75, -63, -65, -76, -77,
	39, -117, -118, -119, -13, -14, -63, -114, -115, -121,
	-122, -123, -13, -14, -63, -21, -96, -97, -98, -13,
	-14, -63, -65, -35, -36, -37, -38, -39, -40, -42,
	-76, -43, -44, -69, -70, -71, -13, -14, -72, -63,
	57, -60, -61, -62, -13, -14, -63, 10, 5, 9,
	-24, 7, -27, 9, 5, -30, 9, 5, 9, -113,
	-31, -32, -21, -31, -5, 8, 8, 9, -52, 10,
	8, 5, 9, -81, 10, 8, -4, 4, 9, -127,
	9, -131, -5, 10, 8, -3, 52, 6, -3, -2,
	47, 48, -2, -5, -5, 9, -92, 9, -138, 9,
	9, 9, -102, 10, 8, -5, 9, -76, 8, 4,
	9, -119, -31, -31, 9, -123, 9, -98, 9, -71,
	4, 9, -62, -4, 10, 10, 9, -21, 9, 10,
	-55, -56, -57, -13, -14, -58, 66, -85, -86, -87,
	17, 63, 45, -88, -72, -82, -89, 18, 67, 10,
	10, -66, -67, -13, -14, 10, 10, 10, 10, 10,
	10, -108, -109, -105, -13, -14, -63, -83, -93, -94,
	-41, -106, -107, -78, 9, 9, 10, -140, 9, -57,
	-2, 9, -87, -5, -5, -5, 4, -5, 9, -67,
	9, -105, 9, 9, 10, 10, 10, 10, 10, 8,
	10, -1, 46, 9, -3, 10,
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
	0, 0, 0, 0, 0, 1, 5, 0, 305, 20,
	0, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	216, 63, 66, 19, 26, 35, 215, 65, 0, 250,
	0, 155, 0, 0, 277, 0, 187, 189, 0, 116,
	229, 240, 168, 103, 106, 84, 87, 0, 135, 136,
	128, 249, 258, 154, 289, 312, 286, 287, 186, 122,
	228, 239, 166, 105, 86, 2, 3, 6, 0, 0,
	22, 24, 25, 318, 319, 316, 321, 0, 317, 0,
	27, 30, 0, 32, 33, 0, 36, 0, 39, 40,
	29, 320, 0, 217, 218, 220, 221, 222, 43, 43,
	0, 0, 0, 0, 67, 68, 70, 71, 72, 0,
	0, 0, 129, 131, 132, 133, 134, 0, 0, 0,
	0, 251, 252, 254, 255, 256, 0, 259, 260, 265,
	266, 267, 268, 269, 270, 271, 272, 273, 0, 275,
	0, 0, 0, 0, 0, 0, 0, 0, 156, 157,
	159, 160, 161, 162, 163, 164, 165, 0, 290, 291,
	293, 294, 295, 296, 297, 298, 299, 300, 301, 302,
	303, 0, 0, 279, 280, 281, 282, 283, 284, 285,
	0, 190, 191, 193, 194, 195, 196, 197, 198, 0,
	0, 94, 0, 117, 118, 119, 120, 121, 123, 0,
	0, 0, 230, 231, 233, 234, 235, 43, 43, 0,
	241, 242, 244, 245, 246, 247, 0, 169, 170, 172,
	173, 174, 175, 176, 177, 178, 179, 180, 181, 182,
	183, 184, 185, 0, 107, 108, 110, 111, 112, 113,
	0, 0, 88, 89, 91, 92, 93, 323, 306, 21,
	23, 0, 28, 34, 0, 37, 42, 0, 214, 219,
	0, 44, 61, 0, 0, 225, 226, 64, 69, 73,
	76, 75, 127, 130, 138, 141, 0, 140, 248, 253,
	257, 261, 0, 97, 0, 0, 0, 307, 0, 0,
	308, 309, 0, 0, 96, 153, 158, 288, 292, 311,
	278, 188, 192, 209, 0, 199, 115, 124, 0, 126,
	227, 232, 0, 0, 238, 243, 167, 171, 104, 109,
	0, 85, 90, 0, 31, 38, 223, 62, 224, 95,
	0, 77, 78, 80, 81, 82, 0, 0, 142, 143,
	0, 0, 0, 148, 149, 150, 151, 0, 0, 137,
	274, 0, 99, 101, 102, 262, 263, 264, 310, 304,
	276, 0, 211, 212, 200, 201, 202, 203, 204, 205,
	206, 207, 208, 0, 236, 237, 114, 0, 74, 79,
	0, 139, 144, 0, 0, 0, 0, 0, 98, 100,
	210, 213, 125, 322, 83, 145, 146, 147, 313, 0,
	152, 0, 0, 314, 0, 315,
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
	62, 63, 64, 65, 66, 67,
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
		//line parser.y:148
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
		//line parser.y:156
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
		//line parser.y:170
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:189
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:196
		{
			pop(yylex)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:199
		{
			pop(yylex)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:212
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:223
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:236
		{
			pop(yylex)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:257
		{
			pop(yylex)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:260
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:290
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:293
		{
			pop(yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:298
		{
			if push(yylex, meta.NewExtension(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:317
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:320
		{
			pop(yylex)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:325
		{
			if push(yylex, meta.NewExtensionArg(peek(yylex).(*meta.Extension), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:344
		{
			if set(yylex, meta.SetYinElement(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:351
		{
			pop(yylex)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:354
		{
			pop(yylex)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:378
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:392
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:399
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:402
		{
			pop(yylex)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:414
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:417
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:422
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:442
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:452
		{
			pop(yylex)
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:465
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:477
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:482
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:489
		{
			pop(yylex)
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:510
		{
			yyVAL.token = yyDollar[1].token
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:511
		{
			yyVAL.token = yyDollar[1].token
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:514
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:521
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:524
		{
			pop(yylex)
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:529
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:543
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:552
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetValueRange(r)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:561
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:572
		{
			if set(yylex, meta.SetPattern(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:579
		{
			pop(yylex)
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:584
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:608
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:615
		{
			pop(yylex)
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:643
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:650
		{
			pop(yylex)
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:653
		{
			pop(yylex)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:673
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:692
		{
			pop(yylex)
		}
	case 210:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:695
		{
			pop(yylex)
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:707
		{
			pop(yylex)
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:712
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:729
		{
			pop(yylex)
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:732
		{
			pop(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:737
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:744
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:754
		{
			pop(yylex)
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:759
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:776
		{
			pop(yylex)
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:779
		{
			pop(yylex)
		}
	case 238:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:787
		{
			pop(yylex)
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:792
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 248:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:813
		{
			pop(yylex)
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:818
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 257:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:837
		{
			pop(yylex)
		}
	case 258:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:842
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 262:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:857
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 263:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:862
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 264:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:869
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 276:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:889
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 277:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:896
		{
			pop(yylex)
		}
	case 278:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:899
		{
			pop(yylex)
		}
	case 286:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:914
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 287:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:919
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 288:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:929
		{
			pop(yylex)
		}
	case 289:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:934
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 304:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:962
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:969
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 306:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:972
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 307:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:977
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 308:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:987
		{
			yyVAL.boolean = true
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:988
		{
			yyVAL.boolean = false
		}
	case 310:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:991
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 311:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1001
		{
			pop(yylex)
		}
	case 312:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1006
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 313:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1013
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 314:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1018
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 315:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1032
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 316:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1037
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 317:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1044
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 318:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1051
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 319:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1058
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 320:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1065
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 323:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1076
		{

		}
	}
	goto yystack /* stack new state and value */
}
