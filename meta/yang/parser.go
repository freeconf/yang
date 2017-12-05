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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:1074

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 897

var yyAct = [...]int{

	231, 366, 442, 12, 248, 229, 12, 431, 413, 228,
	370, 425, 318, 155, 127, 247, 323, 235, 315, 298,
	42, 39, 291, 41, 40, 202, 38, 240, 234, 37,
	239, 36, 35, 34, 262, 43, 250, 185, 33, 208,
	174, 166, 23, 283, 237, 192, 160, 150, 23, 340,
	148, 179, 178, 23, 3, 15, 480, 23, 280, 199,
	157, 198, 156, 476, 24, 223, 230, 475, 23, 11,
	24, 483, 11, 368, 310, 24, 427, 23, 224, 24,
	199, 221, 222, 4, 474, 371, 372, 473, 181, 182,
	24, 152, 478, 439, 477, 190, 164, 455, 169, 24,
	180, 176, 417, 188, 180, 226, 60, 148, 148, 195,
	204, 210, 148, 242, 242, 214, 254, 438, 264, 367,
	274, 285, 293, 300, 180, 317, 213, 325, 216, 148,
	249, 249, 232, 259, 244, 244, 193, 256, 148, 215,
	437, 246, 246, 312, 258, 196, 311, 309, 436, 308,
	152, 161, 307, 268, 306, 305, 304, 151, 251, 435,
	164, 303, 163, 434, 168, 467, 169, 175, 23, 187,
	428, 406, 288, 287, 176, 194, 203, 209, 329, 241,
	241, 23, 253, 23, 263, 188, 273, 284, 292, 299,
	24, 316, 195, 324, 278, 405, 429, 330, 236, 410,
	327, 19, 204, 24, 19, 24, 332, 335, 210, 157,
	280, 156, 214, 357, 339, 161, 151, 171, 481, 193,
	271, 170, 348, 213, 369, 216, 163, 360, 196, 343,
	180, 226, 168, 177, 377, 373, 215, 353, 472, 242,
	175, 148, 469, 211, 148, 243, 243, 362, 255, 158,
	265, 187, 276, 286, 294, 301, 249, 319, 194, 326,
	244, 460, 264, 23, 128, 148, 148, 246, 203, 379,
	352, 223, 87, 23, 209, 199, 333, 198, 457, 23,
	162, 454, 27, 285, 224, 24, 91, 92, 93, 94,
	453, 293, 98, 409, 148, 24, 383, 268, 300, 154,
	407, 24, 153, 147, 402, 241, 177, 399, 205, 219,
	180, 226, 60, 396, 385, 317, 384, 398, 312, 397,
	295, 311, 309, 325, 308, 118, 392, 307, 263, 306,
	305, 304, 395, 400, 288, 287, 303, 393, 394, 403,
	211, 23, 167, 391, 148, 142, 404, 148, 148, 284,
	388, 415, 23, 233, 199, 336, 198, 292, 23, 167,
	365, 389, 364, 24, 299, 356, 433, 355, 350, 422,
	349, 243, 387, 126, 24, 125, 382, 342, 342, 346,
	24, 316, 23, 381, 426, 380, 444, 23, 378, 324,
	451, 198, 376, 361, 265, 223, 124, 345, 123, 452,
	205, 450, 23, 448, 24, 359, 219, 449, 224, 24,
	347, 221, 222, 415, 447, 286, 338, 414, 456, 146,
	145, 458, 446, 294, 24, 122, 121, 320, 459, 180,
	301, 433, 432, 463, 180, 117, 60, 116, 468, 120,
	119, 270, 444, 80, 470, 344, 451, 319, 113, 180,
	226, 464, 443, 23, 162, 326, 27, 450, 377, 448,
	115, 112, 114, 449, 212, 111, 245, 245, 110, 257,
	447, 266, 102, 277, 101, 24, 302, 109, 446, 414,
	368, 108, 482, 363, 23, 100, 342, 342, 95, 295,
	374, 375, 97, 76, 96, 88, 129, 432, 71, 70,
	90, 79, 89, 331, 72, 88, 24, 73, 443, 351,
	77, 280, 337, 6, 78, 74, 75, 199, 445, 86,
	419, 427, 67, 334, 85, 328, 7, 23, 47, 103,
	27, 180, 226, 99, 466, 386, 76, 401, 390, 358,
	408, 71, 70, 50, 79, 68, 69, 72, 424, 24,
	73, 144, 143, 77, 141, 48, 49, 78, 74, 75,
	320, 212, 140, 139, 21, 22, 420, 138, 80, 137,
	136, 81, 135, 82, 445, 134, 60, 28, 133, 51,
	57, 132, 131, 67, 130, 107, 106, 7, 23, 47,
	105, 27, 245, 104, 84, 83, 238, 76, 56, 252,
	58, 217, 71, 70, 50, 79, 68, 69, 72, 207,
	24, 73, 206, 54, 77, 266, 48, 49, 78, 74,
	75, 201, 200, 53, 290, 21, 22, 289, 63, 80,
	282, 281, 81, 62, 82, 173, 172, 60, 28, 29,
	51, 67, 441, 440, 269, 267, 23, 261, 260, 59,
	297, 302, 296, 225, 223, 76, 218, 64, 227, 55,
	71, 70, 423, 79, 68, 69, 72, 224, 24, 73,
	221, 222, 77, 421, 418, 354, 78, 74, 75, 197,
	191, 52, 279, 275, 461, 462, 67, 80, 471, 465,
	81, 23, 82, 180, 226, 60, 272, 61, 314, 223,
	76, 313, 65, 430, 220, 71, 70, 322, 79, 68,
	69, 72, 224, 24, 73, 321, 66, 77, 416, 412,
	411, 78, 74, 75, 189, 186, 184, 183, 67, 30,
	46, 45, 80, 23, 44, 81, 32, 82, 180, 226,
	60, 223, 76, 31, 341, 165, 26, 71, 70, 159,
	79, 68, 69, 72, 224, 24, 73, 25, 149, 77,
	20, 18, 17, 78, 74, 75, 16, 14, 13, 10,
	67, 9, 8, 5, 80, 23, 2, 81, 1, 82,
	180, 226, 60, 479, 76, 0, 0, 0, 0, 71,
	70, 0, 79, 68, 69, 72, 0, 24, 73, 0,
	0, 77, 0, 0, 0, 78, 74, 75, 0, 67,
	0, 0, 0, 0, 23, 0, 80, 0, 0, 81,
	0, 82, 180, 76, 60, 0, 0, 0, 71, 70,
	0, 79, 68, 69, 72, 0, 24, 73, 0, 0,
	77, 0, 67, 0, 78, 74, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 80, 76, 0, 81, 0,
	82, 71, 70, 60, 79, 68, 69, 72, 0, 0,
	73, 0, 0, 77, 0, 0, 0, 78, 74, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 81, 0, 82, 0, 0, 60,
}
var yyPact = [...]int{

	29, -1000, 576, 591, 590, 515, -1000, 500, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	492, 500, 500, 500, 500, 480, 484, 500, 528, 477,
	464, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 524, 589, 586,
	582, 581, 473, 469, 460, 457, 453, 440, 452, 427,
	500, 432, 431, 418, 417, 388, 365, 490, 580, 578,
	577, 574, 571, 568, 566, 565, 563, 559, 558, 550,
	500, 548, 547, 412, 411, -1000, -1000, 293, -1000, -1000,
	56, 292, 289, 52, 239, 441, -1000, 329, 211, 207,
	65, -1000, 30, -1000, -1000, -1000, -1000, -1000, 340, 802,
	634, 721, 45, 45, -1000, 251, -1000, 390, 210, 171,
	65, 763, 472, -1000, 370, -1000, 41, 190, -5, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -5, -1000, -1000, -1000, -1000, -1000, 520, 169,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 496, -1000, 267,
	-1000, -1000, 518, -1000, -1000, 346, -1000, 507, -1000, -1000,
	-1000, -1000, 407, 65, -1000, -1000, -1000, -1000, 835, 835,
	500, 389, 371, 401, 30, -1000, -1000, -1000, -1000, 360,
	504, 261, -1000, -1000, -1000, -1000, -1000, 357, 490, 535,
	396, 802, -1000, -1000, -1000, -1000, 384, 634, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 500, -1000,
	352, 67, 474, 38, 38, 500, 500, 383, 721, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 379, 45, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 376, 374, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	367, 390, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 306,
	500, -1000, 363, -1000, -1000, 19, -1000, -1000, -1000, 353,
	534, 334, 65, -1000, -1000, -1000, -1000, 835, 835, 323,
	763, -1000, -1000, -1000, -1000, -1000, 310, 472, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 298, 370, -1000, -1000, -1000, -1000, -1000,
	533, 295, 41, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 490, -1000, -1000, 185, -1000, -1000, 161, -1000, -1000,
	291, 835, -1000, 284, 189, -1000, -1000, -1000, -1000, -1000,
	36, -1000, -1000, -1000, -1000, -1000, 503, 160, -1000, -1000,
	-1000, -1000, -1000, 186, -1000, 56, 153, 149, -1000, 138,
	130, -1000, -1000, 107, 83, -5, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 375, -5, -1000, -1000, 721,
	-1000, -1000, -1000, 281, 272, -1000, -1000, -1000, -1000, -1000,
	-1000, 87, -1000, -1000, 201, -1000, -1000, -1000, -1000, -1000,
	-1000, 269, 36, -1000, -1000, -1000, -1000, 38, 252, 500,
	500, 58, -1000, 66, 500, -1000, -1000, 530, -1000, -1000,
	156, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	233, 375, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 679, -1000, -1000, -1000, 229, -1000, -1000, 77,
	-1000, 74, 57, -1000, -1000, 53, 84, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 10, 209,
	474, -1000, 61, -1000,
}
var yyPgo = [...]int{

	0, 783, 10, 1, 14, 264, 778, 776, 773, 513,
	772, 771, 769, 66, 0, 768, 767, 55, 766, 762,
	761, 198, 760, 758, 47, 757, 749, 46, 746, 745,
	41, 49, 744, 743, 736, 38, 33, 32, 31, 29,
	26, 21, 24, 23, 20, 35, 734, 731, 730, 729,
	727, 726, 37, 725, 724, 720, 719, 8, 718, 716,
	715, 707, 16, 132, 704, 353, 703, 7, 702, 701,
	698, 18, 12, 697, 696, 683, 74, 682, 9, 681,
	680, 45, 27, 36, 679, 675, 674, 673, 662, 659,
	658, 5, 28, 17, 657, 652, 650, 19, 649, 648,
	647, 34, 645, 644, 2, 15, 4, 643, 642, 639,
	636, 635, 40, 52, 51, 633, 631, 630, 43, 628,
	627, 624, 22, 623, 622, 621, 25, 613, 612, 609,
	39, 601, 600, 599, 598, 44, 596, 30, 580, 11,
	13,
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
	80, 81, 81, 81, 81, 4, 4, 83, 82, 84,
	85, 85, 86, 86, 86, 86, 86, 86, 88, 88,
	36, 89, 90, 90, 78, 78, 91, 91, 91, 91,
	91, 91, 91, 94, 45, 95, 95, 96, 96, 97,
	97, 97, 97, 97, 97, 97, 97, 97, 97, 97,
	97, 97, 97, 98, 40, 40, 99, 99, 100, 100,
	101, 101, 101, 101, 101, 101, 103, 104, 104, 104,
	104, 104, 104, 104, 104, 104, 102, 102, 107, 108,
	108, 19, 109, 110, 110, 111, 111, 112, 112, 112,
	112, 112, 113, 114, 43, 115, 116, 116, 117, 117,
	118, 118, 118, 118, 118, 44, 119, 120, 120, 121,
	121, 122, 122, 122, 122, 34, 123, 124, 124, 125,
	125, 126, 126, 126, 35, 127, 128, 129, 129, 105,
	105, 106, 130, 130, 130, 130, 130, 130, 130, 130,
	130, 130, 130, 131, 39, 39, 133, 133, 133, 133,
	133, 133, 133, 132, 132, 37, 134, 135, 136, 136,
	137, 137, 137, 137, 137, 137, 137, 137, 137, 137,
	137, 93, 5, 5, 3, 2, 2, 92, 38, 138,
	87, 87, 139, 139, 1, 13, 14, 11, 12, 18,
	140, 140, 48,
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
	2, 1, 1, 1, 1, 1, 1, 3, 2, 2,
	1, 3, 3, 3, 1, 1, 1, 3, 1, 2,
	4, 2, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 4, 0, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 4, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 4, 1, 1,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 3, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 3, 3, 4, 2, 0, 1, 1,
	2, 1, 1, 1, 1, 4, 2, 0, 1, 1,
	2, 1, 1, 1, 4, 2, 1, 1, 2, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 2, 4, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 4, 2, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 1, 1, 1, 3, 4, 2,
	1, 2, 3, 5, 3, 3, 3, 3, 3, 3,
	1, 5, 3,
}
var yyChk = [...]int{

	-1000, -6, -7, 25, 54, -8, -9, 11, -10, -11,
	-12, -13, -14, -15, -16, -17, -18, -19, -20, -21,
	-22, 49, 50, 12, 34, -25, -28, 15, 62, -109,
	-49, -33, -34, -35, -36, -37, -38, -39, -40, -41,
	-42, -43, -44, -45, -46, -47, -48, 13, 40, 41,
	28, 64, -79, -123, -127, -89, -134, -138, -132, -98,
	61, -73, -115, -119, -94, -68, -59, 7, 30, 31,
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
	10, 10, -110, -111, -112, -13, -14, -63, -113, -114,
	59, 23, 24, -50, -51, -52, -53, -13, -14, -54,
	65, -80, -81, -82, -13, -14, -83, -84, 16, 14,
	-124, -125, -126, -13, -14, -21, -128, -129, -130, -13,
	-14, -63, -65, -105, -106, -92, -93, -131, 22, -21,
	-64, 36, 37, 20, 33, 19, 60, -90, -78, -91,
	-13, -14, -63, -65, -92, -93, -21, -135, -136, -137,
	-82, -13, -14, -63, -41, -65, -92, -105, -106, -93,
	-83, -135, -133, -13, -14, -63, -41, -65, -92, -93,
	-99, -100, -101, -13, -14, -63, -65, -102, -45, -103,
	51, 10, -74, -13, -14, -75, -63, -65, -76, -77,
	39, -116, -117, -118, -13, -14, -63, -113, -114, -120,
	-121, -122, -13, -14, -63, -21, -95, -96, -97, -13,
	-14, -63, -65, -35, -36, -37, -38, -39, -40, -42,
	-76, -43, -44, -69, -70, -71, -13, -14, -72, -63,
	57, -60, -61, -62, -13, -14, -63, 10, 5, 9,
	-24, 7, -27, 9, 5, -30, 9, 5, 9, -112,
	-31, -32, -21, -31, -5, 8, 8, 9, -52, 10,
	8, 5, 9, -81, -85, 10, 8, -4, 4, 9,
	-126, 9, -130, -5, 10, 8, -3, 52, 6, -3,
	-2, 47, 48, -2, -5, -5, 9, -91, 9, -137,
	9, 9, 9, -101, 10, 8, -5, 9, -76, 8,
	4, 9, -118, -31, -31, 9, -122, 9, -97, 9,
	-71, 4, 9, -62, -4, 10, 10, 9, -21, 9,
	10, -55, -56, -57, -13, -14, -58, 66, -86, 17,
	63, -87, -72, -88, 45, -139, -82, 18, 10, 10,
	-66, -67, -13, -14, 10, 10, 10, 10, 10, 10,
	-107, -108, -104, -13, -14, -63, -83, -92, -93, -41,
	-105, -106, -78, 9, 9, 10, -140, 9, -57, -2,
	9, -5, -5, -139, -82, -5, 4, 9, -67, 9,
	-104, 9, 9, 10, 10, 10, 10, 10, 8, -1,
	46, 9, -3, 10,
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
	0, 0, 0, 0, 0, 1, 5, 0, 302, 20,
	0, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	213, 63, 66, 19, 26, 35, 212, 65, 0, 247,
	0, 152, 0, 0, 274, 0, 184, 186, 0, 116,
	226, 237, 165, 103, 106, 84, 87, 0, 135, 136,
	128, 246, 255, 151, 286, 309, 283, 284, 183, 122,
	225, 236, 163, 105, 86, 2, 3, 6, 0, 0,
	22, 24, 25, 317, 318, 315, 320, 0, 316, 0,
	27, 30, 0, 32, 33, 0, 36, 0, 39, 40,
	29, 319, 0, 214, 215, 217, 218, 219, 43, 43,
	0, 0, 0, 0, 67, 68, 70, 71, 72, 0,
	0, 0, 129, 131, 132, 133, 134, 0, 0, 0,
	0, 248, 249, 251, 252, 253, 0, 256, 257, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 0, 272,
	0, 0, 0, 0, 0, 0, 0, 0, 153, 154,
	156, 157, 158, 159, 160, 161, 162, 0, 287, 288,
	290, 291, 292, 293, 294, 295, 296, 297, 298, 299,
	300, 0, 0, 276, 277, 278, 279, 280, 281, 282,
	0, 187, 188, 190, 191, 192, 193, 194, 195, 0,
	0, 94, 0, 117, 118, 119, 120, 121, 123, 0,
	0, 0, 227, 228, 230, 231, 232, 43, 43, 0,
	238, 239, 241, 242, 243, 244, 0, 166, 167, 169,
	170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
	180, 181, 182, 0, 107, 108, 110, 111, 112, 113,
	0, 0, 88, 89, 91, 92, 93, 322, 303, 21,
	23, 0, 28, 34, 0, 37, 42, 0, 211, 216,
	0, 44, 61, 0, 0, 222, 223, 64, 69, 73,
	76, 75, 127, 130, 138, 140, 0, 0, 139, 245,
	250, 254, 258, 0, 97, 0, 0, 0, 304, 0,
	0, 305, 306, 0, 0, 96, 150, 155, 285, 289,
	308, 275, 185, 189, 206, 0, 196, 115, 124, 0,
	126, 224, 229, 0, 0, 235, 240, 164, 168, 104,
	109, 0, 85, 90, 0, 31, 38, 220, 62, 221,
	95, 0, 77, 78, 80, 81, 82, 0, 0, 0,
	0, 144, 145, 146, 0, 310, 148, 0, 137, 271,
	0, 99, 101, 102, 259, 260, 261, 307, 301, 273,
	0, 208, 209, 197, 198, 199, 200, 201, 202, 203,
	204, 205, 0, 233, 234, 114, 0, 74, 79, 0,
	141, 0, 0, 311, 149, 0, 0, 98, 100, 207,
	210, 125, 321, 83, 142, 143, 147, 312, 0, 0,
	0, 313, 0, 314,
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
	62, 63, 64, 65, 66,
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
		//line parser.y:147
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
		//line parser.y:155
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
		//line parser.y:169
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:188
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:195
		{
			pop(yylex)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:198
		{
			pop(yylex)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:211
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:222
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:235
		{
			pop(yylex)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:240
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:256
		{
			pop(yylex)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:259
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:289
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:292
		{
			pop(yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:297
		{
			if push(yylex, meta.NewExtension(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:316
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:319
		{
			pop(yylex)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:324
		{
			if push(yylex, meta.NewExtensionArg(peek(yylex).(*meta.Extension), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:343
		{
			if set(yylex, meta.SetYinElement(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:350
		{
			pop(yylex)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:353
		{
			pop(yylex)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:358
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:377
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:384
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:391
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:398
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:401
		{
			pop(yylex)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:413
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:416
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:421
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:441
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:451
		{
			pop(yylex)
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:464
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:476
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:481
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:491
		{
			pop(yylex)
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:496
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:513
		{
			yyVAL.token = yyDollar[1].token
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:514
		{
			yyVAL.token = yyDollar[1].token
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:517
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:527
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:534
		{
			pop(yylex)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:537
		{
			pop(yylex)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:542
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:551
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
		//line parser.y:563
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:577
		{
			pop(yylex)
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:582
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:606
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:613
		{
			pop(yylex)
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:641
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:648
		{
			pop(yylex)
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:651
		{
			pop(yylex)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:671
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:690
		{
			pop(yylex)
		}
	case 207:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:693
		{
			pop(yylex)
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:705
		{
			pop(yylex)
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:710
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:727
		{
			pop(yylex)
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:730
		{
			pop(yylex)
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:735
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:742
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
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
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:774
		{
			pop(yylex)
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:777
		{
			pop(yylex)
		}
	case 235:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:785
		{
			pop(yylex)
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:790
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 245:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:811
		{
			pop(yylex)
		}
	case 246:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:816
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 254:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:835
		{
			pop(yylex)
		}
	case 255:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:840
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 259:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:855
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 260:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:860
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 261:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:867
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 273:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:887
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 274:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:894
		{
			pop(yylex)
		}
	case 275:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:897
		{
			pop(yylex)
		}
	case 283:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:912
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 284:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:917
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 285:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:927
		{
			pop(yylex)
		}
	case 286:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:932
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 301:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:960
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:967
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 303:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:970
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 304:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:975
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:985
		{
			yyVAL.boolean = true
		}
	case 306:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:986
		{
			yyVAL.boolean = false
		}
	case 307:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:989
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 308:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:999
		{
			pop(yylex)
		}
	case 309:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1004
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 312:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1015
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 313:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1020
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 314:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1027
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 315:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1032
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 316:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1039
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 317:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1046
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 318:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1053
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 319:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1060
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 322:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1071
		{

		}
	}
	goto yystack /* stack new state and value */
}
