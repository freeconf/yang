//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/freeconf/c2g/c2"
	"github.com/freeconf/c2g/meta"
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

//line parser.y:1083

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 912

var yyAct = [...]int{

	231, 365, 482, 12, 444, 229, 12, 433, 250, 419,
	369, 127, 412, 248, 155, 247, 318, 235, 228, 323,
	240, 39, 315, 42, 262, 291, 239, 310, 234, 340,
	43, 179, 41, 208, 298, 40, 283, 192, 38, 174,
	37, 166, 202, 185, 15, 160, 23, 150, 23, 36,
	35, 23, 23, 148, 3, 34, 237, 280, 33, 199,
	487, 480, 420, 429, 370, 371, 230, 23, 24, 11,
	24, 491, 11, 24, 24, 178, 468, 157, 467, 156,
	477, 457, 488, 4, 384, 23, 383, 440, 439, 24,
	422, 152, 364, 320, 363, 180, 164, 367, 169, 180,
	416, 176, 320, 188, 190, 148, 148, 24, 421, 195,
	204, 210, 428, 242, 242, 438, 254, 196, 264, 486,
	274, 285, 293, 300, 214, 317, 213, 325, 216, 193,
	249, 249, 232, 259, 244, 244, 157, 256, 156, 215,
	161, 246, 246, 366, 258, 479, 312, 278, 268, 478,
	152, 437, 288, 441, 431, 311, 409, 151, 309, 436,
	164, 308, 163, 307, 168, 236, 169, 175, 19, 187,
	251, 19, 306, 305, 176, 194, 203, 209, 304, 241,
	241, 303, 253, 148, 263, 188, 273, 284, 292, 299,
	148, 316, 195, 324, 148, 430, 287, 330, 148, 148,
	196, 148, 204, 405, 161, 332, 404, 335, 210, 343,
	356, 233, 193, 339, 327, 355, 151, 354, 271, 170,
	350, 214, 349, 213, 368, 216, 163, 158, 348, 353,
	171, 476, 168, 177, 376, 372, 215, 154, 473, 242,
	175, 361, 462, 211, 359, 243, 243, 459, 255, 23,
	265, 187, 276, 286, 294, 301, 249, 319, 194, 326,
	244, 456, 264, 148, 148, 378, 153, 246, 203, 455,
	471, 24, 148, 23, 209, 205, 219, 329, 23, 408,
	23, 147, 148, 285, 406, 128, 382, 295, 270, 401,
	80, 293, 268, 87, 398, 24, 180, 226, 300, 126,
	24, 125, 24, 387, 396, 241, 177, 91, 92, 93,
	94, 148, 486, 98, 288, 317, 395, 392, 393, 391,
	394, 312, 212, 325, 245, 245, 148, 257, 263, 266,
	311, 277, 397, 309, 302, 390, 308, 399, 307, 124,
	211, 123, 402, 403, 342, 342, 118, 306, 305, 284,
	117, 414, 116, 304, 386, 381, 303, 292, 287, 336,
	380, 23, 23, 167, 299, 435, 142, 205, 115, 223,
	114, 243, 424, 219, 333, 23, 425, 23, 162, 379,
	27, 316, 224, 24, 24, 446, 181, 182, 102, 324,
	101, 377, 375, 448, 265, 360, 358, 24, 453, 24,
	452, 97, 450, 96, 23, 167, 451, 454, 180, 226,
	60, 347, 414, 449, 90, 286, 89, 413, 458, 212,
	338, 23, 180, 294, 460, 198, 24, 461, 463, 223,
	301, 434, 388, 435, 346, 424, 345, 146, 145, 425,
	472, 122, 224, 24, 446, 221, 222, 319, 474, 23,
	245, 445, 448, 342, 342, 326, 295, 453, 121, 452,
	376, 450, 120, 119, 113, 451, 344, 112, 180, 484,
	60, 24, 449, 266, 111, 110, 280, 331, 413, 109,
	108, 23, 484, 199, 489, 198, 100, 367, 490, 223,
	95, 23, 6, 199, 88, 198, 180, 226, 86, 434,
	88, 129, 224, 24, 362, 221, 222, 407, 351, 302,
	445, 373, 374, 24, 337, 334, 470, 447, 67, 328,
	85, 103, 7, 23, 47, 99, 27, 400, 180, 226,
	60, 389, 76, 357, 144, 483, 143, 71, 70, 50,
	79, 68, 69, 72, 141, 24, 73, 140, 483, 77,
	139, 48, 49, 78, 74, 75, 385, 138, 137, 136,
	21, 22, 23, 162, 80, 27, 135, 81, 134, 82,
	133, 485, 60, 28, 67, 51, 447, 132, 7, 23,
	47, 131, 27, 130, 24, 107, 106, 105, 76, 104,
	84, 83, 481, 71, 70, 50, 79, 68, 69, 72,
	427, 24, 73, 57, 238, 77, 56, 48, 49, 78,
	74, 75, 252, 58, 217, 207, 21, 22, 206, 54,
	80, 67, 201, 81, 200, 82, 23, 53, 60, 28,
	290, 51, 289, 225, 223, 76, 218, 63, 282, 281,
	71, 70, 62, 79, 68, 69, 72, 224, 24, 73,
	221, 222, 77, 173, 172, 29, 78, 74, 75, 443,
	442, 269, 267, 261, 67, 260, 475, 80, 59, 23,
	81, 297, 82, 180, 226, 60, 296, 223, 76, 64,
	227, 55, 426, 71, 70, 423, 79, 68, 69, 72,
	224, 24, 73, 418, 417, 77, 197, 191, 52, 78,
	74, 75, 279, 275, 272, 61, 464, 465, 466, 314,
	80, 313, 65, 81, 469, 82, 180, 226, 60, 67,
	432, 220, 322, 352, 23, 321, 23, 66, 199, 415,
	198, 411, 223, 76, 410, 189, 186, 184, 71, 70,
	183, 79, 68, 69, 72, 224, 24, 73, 24, 30,
	77, 46, 45, 44, 78, 74, 75, 32, 31, 341,
	165, 67, 26, 159, 25, 80, 23, 149, 81, 20,
	82, 180, 226, 60, 18, 76, 17, 16, 14, 13,
	71, 70, 10, 79, 68, 69, 72, 9, 24, 73,
	8, 5, 77, 2, 1, 0, 78, 74, 75, 0,
	67, 0, 0, 0, 0, 23, 0, 80, 0, 0,
	81, 0, 82, 180, 76, 60, 0, 0, 0, 71,
	70, 0, 79, 68, 69, 72, 0, 24, 73, 0,
	0, 77, 0, 67, 0, 78, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 80, 76, 0, 81,
	0, 82, 71, 70, 60, 79, 68, 69, 72, 0,
	0, 73, 0, 23, 77, 0, 0, 0, 78, 74,
	75, 0, 76, 0, 0, 0, 0, 71, 70, 80,
	79, 0, 81, 72, 82, 24, 73, 60, 0, 77,
	280, 0, 0, 78, 74, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	180, 226,
}
var yyPact = [...]int{

	29, -1000, 567, 587, 586, 511, -1000, 489, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	406, 489, 489, 489, 489, 482, 393, 489, 520, 478,
	380, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 516, 585, 583,
	582, 581, 472, 471, 467, 466, 459, 456, 360, 342,
	489, 455, 454, 450, 433, 331, 291, 495, 579, 577,
	573, 566, 564, 562, 555, 554, 553, 546, 543, 540,
	489, 532, 530, 430, 429, -1000, -1000, 271, -1000, -1000,
	55, 256, 227, 128, 217, 550, -1000, 392, 209, 220,
	363, -1000, 39, -1000, -1000, -1000, -1000, -1000, 479, 793,
	614, 712, 469, 469, -1000, 349, -1000, 237, 208, 437,
	363, 754, 851, -1000, 36, -1000, 40, 204, -2, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -2, -1000, -1000, -1000, -1000, -1000, 514, 268,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 470, -1000, 365,
	-1000, -1000, 510, -1000, -1000, 350, -1000, 509, -1000, -1000,
	-1000, -1000, 411, 363, -1000, -1000, -1000, -1000, 826, 826,
	489, 428, 426, 402, 39, -1000, -1000, -1000, -1000, 212,
	503, 714, -1000, -1000, -1000, -1000, -1000, 207, 495, 529,
	387, 793, -1000, -1000, -1000, -1000, 386, 614, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 489, -1000,
	84, 91, 481, 17, 17, 489, 489, 383, 712, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 382, 469, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 370, 351, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	346, 237, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 76,
	489, -1000, 345, -1000, -1000, 18, -1000, -1000, -1000, 424,
	527, 326, 363, -1000, -1000, -1000, -1000, 826, 826, 311,
	754, -1000, -1000, -1000, -1000, -1000, 295, 851, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 285, 36, -1000, -1000, -1000, -1000, -1000,
	523, 280, 40, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 495, -1000, -1000, 196, -1000, -1000, 193, -1000, -1000,
	275, 826, -1000, 270, 146, -1000, -1000, -1000, -1000, -1000,
	34, -1000, -1000, -1000, -1000, 45, 185, -1000, -1000, -1000,
	-1000, -1000, 144, -1000, 55, 149, 141, -1000, 105, 78,
	-1000, -1000, 77, 143, -2, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 409, -2, -1000, -1000, 712, -1000,
	-1000, -1000, 260, 252, -1000, -1000, -1000, -1000, -1000, -1000,
	71, -1000, -1000, 69, -1000, -1000, -1000, -1000, -1000, -1000,
	238, 34, -1000, -1000, -1000, -1000, 17, 233, 45, -1000,
	489, 489, 489, -1000, -1000, -1000, -1000, 68, 489, 512,
	-1000, -1000, 261, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 229, 409, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 657, -1000, -1000, -1000, 222, -1000,
	-1000, 70, -1000, -1000, 139, 135, 51, -1000, 266, 50,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 73, -1000, -1000, -1000, -1000, 481, -1000, -1000, -1000,
	61, -1000,
}
var yyPgo = [...]int{

	0, 10, 1, 11, 285, 794, 793, 791, 492, 790,
	787, 782, 66, 0, 779, 778, 44, 777, 776, 774,
	165, 769, 767, 47, 764, 763, 45, 762, 760, 41,
	29, 759, 758, 757, 58, 55, 50, 49, 40, 38,
	21, 35, 32, 23, 30, 753, 752, 751, 749, 740,
	737, 43, 736, 735, 734, 731, 12, 729, 727, 725,
	722, 19, 132, 721, 211, 720, 7, 712, 711, 709,
	22, 16, 705, 704, 703, 27, 702, 18, 698, 697,
	37, 20, 8, 696, 694, 693, 9, 685, 682, 681,
	680, 5, 28, 17, 679, 676, 671, 34, 668, 665,
	663, 24, 662, 661, 4, 15, 13, 660, 659, 655,
	654, 653, 39, 75, 31, 642, 639, 638, 36, 637,
	632, 630, 25, 627, 624, 622, 42, 619, 618, 615,
	33, 614, 613, 612, 606, 56, 604, 26, 603, 600,
	592, 2, 571, 14,
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
	73, 73, 72, 74, 74, 75, 76, 32, 78, 79,
	79, 80, 80, 80, 80, 3, 3, 82, 81, 81,
	83, 84, 84, 85, 85, 86, 86, 86, 86, 86,
	86, 86, 88, 35, 89, 90, 90, 77, 77, 91,
	91, 91, 91, 91, 91, 91, 94, 44, 95, 95,
	96, 96, 97, 97, 97, 97, 97, 97, 97, 97,
	97, 97, 97, 97, 97, 97, 98, 39, 39, 99,
	99, 100, 100, 101, 101, 101, 101, 101, 101, 103,
	104, 104, 104, 104, 104, 104, 104, 104, 104, 102,
	102, 107, 108, 108, 18, 109, 110, 110, 111, 111,
	112, 112, 112, 112, 112, 113, 114, 42, 115, 116,
	116, 117, 117, 118, 118, 118, 118, 118, 43, 119,
	120, 120, 121, 121, 122, 122, 122, 122, 33, 123,
	124, 124, 125, 125, 126, 126, 126, 34, 127, 128,
	129, 129, 105, 105, 106, 130, 130, 130, 130, 130,
	130, 130, 130, 130, 130, 130, 131, 38, 38, 133,
	133, 133, 133, 133, 133, 133, 132, 132, 36, 134,
	135, 136, 136, 137, 137, 137, 137, 137, 137, 137,
	137, 137, 137, 137, 93, 4, 4, 2, 1, 1,
	92, 37, 138, 87, 87, 139, 140, 140, 141, 141,
	141, 142, 12, 13, 10, 11, 17, 143, 143, 47,
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
	3, 4, 2, 2, 4, 2, 1, 2, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 1, 5, 3,
}
var yyChk = [...]int{

	-1000, -5, -6, 25, 54, -7, -8, 11, -9, -10,
	-11, -12, -13, -14, -15, -16, -17, -18, -19, -20,
	-21, 49, 50, 12, 34, -24, -27, 15, 62, -109,
	-48, -32, -33, -34, -35, -36, -37, -38, -39, -40,
	-41, -42, -43, -44, -45, -46, -47, 13, 40, 41,
	28, 64, -78, -123, -127, -89, -134, -138, -132, -98,
	61, -72, -115, -119, -94, -67, -58, 7, 30, 31,
	27, 26, 32, 35, 43, 44, 21, 38, 42, 29,
	53, 56, 58, 4, 4, 9, -8, -4, 5, 10,
	8, -4, -4, -4, -4, 8, 10, 8, -4, 5,
	8, 10, 8, 5, 4, 4, 4, 4, 8, 8,
	8, 8, 8, 8, 10, 8, 10, 8, -4, 8,
	8, 8, 8, 10, 8, 10, 8, -3, -4, 6,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, -4, 4, 4, 8, 8, 10, 55, -22,
	-23, -12, -13, 10, 10, -143, 10, 8, 10, -25,
	-26, -16, 13, -12, -13, -28, -29, 13, -12, -13,
	10, 10, -110, -111, -112, -12, -13, -62, -113, -114,
	59, 23, 24, -49, -50, -51, -52, -12, -13, -53,
	65, -79, -80, -81, -12, -13, -82, -83, 16, 14,
	-124, -125, -126, -12, -13, -20, -128, -129, -130, -12,
	-13, -62, -64, -105, -106, -92, -93, -131, 22, -20,
	-63, 36, 37, 20, 33, 19, 60, -90, -77, -91,
	-12, -13, -62, -64, -92, -93, -20, -135, -136, -137,
	-81, -12, -13, -62, -40, -64, -92, -105, -106, -93,
	-82, -135, -133, -12, -13, -62, -40, -64, -92, -93,
	-99, -100, -101, -12, -13, -62, -64, -102, -44, -103,
	51, 10, -73, -12, -13, -74, -62, -64, -75, -76,
	39, -116, -117, -118, -12, -13, -62, -113, -114, -120,
	-121, -122, -12, -13, -62, -20, -95, -96, -97, -12,
	-13, -62, -64, -34, -35, -36, -37, -38, -39, -41,
	-75, -42, -43, -68, -69, -70, -12, -13, -71, -62,
	57, -59, -60, -61, -12, -13, -62, 10, 5, 9,
	-23, 7, -26, 9, 5, -29, 9, 5, 9, -112,
	-30, -31, -20, -30, -4, 8, 8, 9, -51, 10,
	8, 5, 9, -80, 10, 8, -3, 4, 9, -126,
	9, -130, -4, 10, 8, -2, 52, 6, -2, -1,
	47, 48, -1, -4, -4, 9, -91, 9, -137, 9,
	9, 9, -101, 10, 8, -4, 9, -75, 8, 4,
	9, -118, -30, -30, 9, -122, 9, -97, 9, -70,
	4, 9, -61, -3, 10, 10, 9, -20, 9, 10,
	-54, -55, -56, -12, -13, -57, 66, -84, -85, -86,
	17, 63, 45, -87, -71, -81, -88, -139, 67, 18,
	10, 10, -65, -66, -12, -13, 10, 10, 10, 10,
	10, 10, -107, -108, -104, -12, -13, -62, -82, -92,
	-93, -40, -105, -106, -77, 9, 9, 10, -143, 9,
	-56, -1, 9, -86, -4, -4, -4, 10, 8, -4,
	4, 9, -66, 9, -104, 9, 9, 10, 10, 10,
	10, -140, -141, -12, -13, -142, 46, 10, 9, -141,
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
	0, 0, 0, 0, 0, 1, 5, 0, 305, 20,
	0, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	216, 63, 66, 19, 26, 35, 215, 65, 0, 250,
	0, 155, 0, 0, 277, 0, 187, 189, 0, 116,
	229, 240, 168, 103, 106, 84, 87, 0, 135, 136,
	128, 249, 258, 154, 289, 312, 286, 287, 186, 122,
	228, 239, 166, 105, 86, 2, 3, 6, 0, 0,
	22, 24, 25, 324, 325, 322, 327, 0, 323, 0,
	27, 30, 0, 32, 33, 0, 36, 0, 39, 40,
	29, 326, 0, 217, 218, 220, 221, 222, 43, 43,
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
	0, 0, 88, 89, 91, 92, 93, 329, 306, 21,
	23, 0, 28, 34, 0, 37, 42, 0, 214, 219,
	0, 44, 61, 0, 0, 225, 226, 64, 69, 73,
	76, 75, 127, 130, 138, 141, 0, 140, 248, 253,
	257, 261, 0, 97, 0, 0, 0, 307, 0, 0,
	308, 309, 0, 0, 96, 153, 158, 288, 292, 311,
	278, 188, 192, 209, 0, 199, 115, 124, 0, 126,
	227, 232, 0, 0, 238, 243, 167, 171, 104, 109,
	0, 85, 90, 0, 31, 38, 223, 62, 224, 95,
	0, 77, 78, 80, 81, 82, 0, 0, 142, 143,
	0, 0, 0, 148, 149, 150, 151, 0, 0, 0,
	137, 274, 0, 99, 101, 102, 262, 263, 264, 310,
	304, 276, 0, 211, 212, 200, 201, 202, 203, 204,
	205, 206, 207, 208, 0, 236, 237, 114, 0, 74,
	79, 0, 139, 144, 0, 0, 0, 313, 0, 0,
	315, 98, 100, 210, 213, 125, 328, 83, 145, 146,
	147, 0, 316, 318, 319, 320, 0, 152, 314, 317,
	0, 321,
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
		//line parser.y:146
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
		//line parser.y:154
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
		//line parser.y:168
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:187
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:194
		{
			pop(yylex)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:197
		{
			pop(yylex)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:210
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:221
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:234
		{
			pop(yylex)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:239
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:255
		{
			pop(yylex)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:258
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:288
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:291
		{
			pop(yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:296
		{
			if push(yylex, meta.NewExtension(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:315
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:318
		{
			pop(yylex)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:323
		{
			if push(yylex, meta.NewExtensionArg(peek(yylex).(*meta.Extension), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:342
		{
			if set(yylex, meta.SetYinElement(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:349
		{
			pop(yylex)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:352
		{
			pop(yylex)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:357
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:376
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:383
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:390
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:397
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:400
		{
			pop(yylex)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:412
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:415
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:420
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:440
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:450
		{
			pop(yylex)
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:463
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:475
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:480
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:487
		{
			pop(yylex)
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:492
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:508
		{
			yyVAL.token = yyDollar[1].token
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:509
		{
			yyVAL.token = yyDollar[1].token
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:512
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:519
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:522
		{
			pop(yylex)
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:527
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:541
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
		//line parser.y:550
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
		//line parser.y:559
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:570
		{
			if set(yylex, meta.SetPattern(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:577
		{
			pop(yylex)
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:582
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:606
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:613
		{
			pop(yylex)
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:641
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:648
		{
			pop(yylex)
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:651
		{
			pop(yylex)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:671
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:690
		{
			pop(yylex)
		}
	case 210:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:693
		{
			pop(yylex)
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:705
		{
			pop(yylex)
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:710
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:727
		{
			pop(yylex)
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:730
		{
			pop(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:735
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:742
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:752
		{
			pop(yylex)
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:757
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:774
		{
			pop(yylex)
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:777
		{
			pop(yylex)
		}
	case 238:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:785
		{
			pop(yylex)
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:790
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 248:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:811
		{
			pop(yylex)
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:816
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 257:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:835
		{
			pop(yylex)
		}
	case 258:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:840
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 262:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:855
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 263:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:860
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 264:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:867
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 276:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:887
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 277:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:894
		{
			pop(yylex)
		}
	case 278:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:897
		{
			pop(yylex)
		}
	case 286:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:912
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 287:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:917
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 288:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:927
		{
			pop(yylex)
		}
	case 289:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:932
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 304:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:960
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:967
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 306:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:970
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 307:
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
	case 308:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:985
		{
			yyVAL.boolean = true
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:986
		{
			yyVAL.boolean = false
		}
	case 310:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:989
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 311:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:999
		{
			pop(yylex)
		}
	case 312:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1004
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 313:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1011
		{
			pop(yylex)
		}
	case 314:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1014
		{
			pop(yylex)
		}
	case 315:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1019
		{
			if push(yylex, meta.NewEnum(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 321:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1034
		{
			if set(yylex, meta.SetEnumValue(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 322:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1041
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 323:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1048
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 324:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1055
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 325:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1062
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 326:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1069
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 329:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1080
		{

		}
	}
	goto yystack /* stack new state and value */
}
