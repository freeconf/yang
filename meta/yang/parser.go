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

//line parser.y:1101

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 950

var yyAct = [...]int{

	233, 490, 369, 12, 231, 451, 12, 440, 423, 416,
	253, 321, 230, 251, 155, 326, 127, 373, 250, 242,
	237, 318, 42, 301, 265, 343, 236, 294, 41, 40,
	39, 38, 37, 36, 286, 35, 34, 241, 204, 210,
	185, 174, 33, 43, 150, 192, 166, 160, 313, 179,
	496, 23, 178, 15, 148, 23, 23, 239, 201, 23,
	200, 201, 283, 200, 23, 23, 232, 225, 500, 11,
	3, 495, 11, 24, 248, 374, 375, 24, 24, 479,
	226, 24, 23, 223, 224, 488, 24, 24, 157, 485,
	156, 152, 475, 487, 474, 148, 164, 464, 169, 4,
	323, 176, 180, 188, 24, 420, 180, 228, 60, 196,
	206, 212, 199, 244, 244, 199, 257, 190, 267, 197,
	277, 288, 296, 303, 216, 320, 371, 328, 193, 215,
	148, 218, 234, 252, 252, 148, 262, 217, 148, 249,
	249, 447, 261, 246, 246, 315, 259, 446, 445, 161,
	152, 314, 312, 486, 311, 310, 309, 151, 308, 307,
	164, 271, 163, 448, 168, 306, 169, 175, 281, 187,
	291, 254, 370, 290, 176, 195, 205, 211, 332, 243,
	243, 23, 256, 194, 266, 188, 276, 287, 295, 302,
	444, 319, 196, 327, 333, 443, 23, 162, 148, 27,
	23, 437, 197, 24, 206, 346, 436, 335, 148, 409,
	212, 193, 338, 161, 408, 342, 151, 360, 24, 438,
	330, 171, 24, 216, 413, 351, 163, 372, 215, 274,
	218, 392, 168, 177, 23, 380, 217, 356, 170, 158,
	175, 244, 363, 213, 376, 245, 245, 180, 258, 365,
	268, 187, 279, 289, 297, 304, 24, 322, 195, 329,
	154, 252, 484, 235, 148, 267, 194, 249, 153, 148,
	205, 246, 481, 273, 148, 80, 211, 497, 382, 469,
	23, 180, 228, 148, 148, 466, 288, 128, 157, 386,
	156, 388, 349, 387, 296, 87, 368, 358, 367, 357,
	147, 303, 24, 463, 462, 148, 177, 243, 271, 91,
	92, 93, 94, 148, 494, 98, 396, 397, 320, 238,
	395, 399, 19, 315, 401, 19, 328, 391, 23, 314,
	312, 266, 311, 310, 309, 291, 308, 307, 290, 403,
	412, 406, 213, 306, 23, 148, 410, 353, 118, 352,
	24, 407, 287, 23, 418, 181, 182, 126, 348, 125,
	295, 405, 494, 23, 167, 402, 24, 302, 142, 442,
	428, 400, 398, 245, 214, 24, 247, 247, 430, 260,
	283, 269, 394, 280, 319, 24, 305, 390, 124, 453,
	123, 180, 327, 117, 115, 116, 114, 268, 385, 455,
	180, 228, 460, 355, 384, 461, 23, 459, 201, 457,
	200, 102, 97, 101, 96, 456, 418, 383, 289, 458,
	417, 90, 465, 89, 381, 467, 297, 379, 24, 207,
	221, 470, 364, 304, 428, 441, 476, 362, 468, 350,
	442, 298, 430, 341, 334, 339, 146, 480, 23, 167,
	322, 453, 145, 122, 121, 452, 482, 120, 329, 119,
	113, 455, 199, 112, 460, 111, 380, 110, 347, 459,
	24, 457, 109, 214, 108, 100, 492, 456, 95, 88,
	129, 458, 417, 201, 371, 88, 424, 435, 23, 6,
	492, 498, 200, 359, 354, 86, 225, 499, 345, 345,
	340, 337, 331, 23, 247, 103, 441, 99, 366, 226,
	24, 225, 223, 224, 426, 377, 378, 452, 478, 404,
	393, 454, 361, 207, 226, 24, 323, 144, 269, 221,
	143, 141, 425, 140, 139, 180, 434, 60, 433, 138,
	336, 137, 491, 23, 162, 136, 27, 135, 134, 133,
	180, 228, 60, 132, 131, 130, 491, 107, 106, 105,
	104, 389, 84, 493, 305, 24, 67, 83, 85, 489,
	7, 23, 47, 432, 27, 57, 240, 56, 255, 58,
	76, 219, 209, 454, 208, 71, 70, 50, 79, 68,
	69, 72, 54, 24, 73, 203, 202, 77, 53, 48,
	49, 78, 74, 75, 293, 292, 63, 285, 21, 22,
	345, 345, 80, 298, 284, 81, 62, 82, 173, 172,
	60, 28, 29, 51, 67, 450, 449, 272, 7, 23,
	47, 270, 27, 264, 263, 59, 300, 299, 76, 64,
	229, 55, 431, 71, 70, 50, 79, 68, 69, 72,
	429, 24, 73, 427, 422, 77, 421, 48, 49, 78,
	74, 75, 198, 191, 411, 52, 21, 22, 282, 278,
	80, 275, 67, 81, 61, 82, 317, 23, 60, 28,
	316, 51, 65, 439, 227, 225, 76, 220, 222, 325,
	324, 71, 70, 66, 79, 68, 69, 72, 226, 24,
	73, 223, 224, 77, 419, 415, 414, 78, 74, 75,
	189, 186, 471, 472, 473, 67, 184, 483, 80, 183,
	23, 81, 477, 82, 180, 228, 60, 30, 225, 76,
	46, 45, 44, 32, 71, 70, 31, 79, 68, 69,
	72, 226, 24, 73, 344, 165, 77, 26, 159, 25,
	78, 74, 75, 149, 20, 18, 17, 67, 16, 14,
	13, 80, 23, 10, 81, 9, 82, 180, 228, 60,
	225, 76, 8, 5, 2, 1, 71, 70, 0, 79,
	68, 69, 72, 226, 24, 73, 0, 0, 77, 0,
	0, 0, 78, 74, 75, 0, 0, 0, 0, 67,
	0, 0, 0, 80, 23, 0, 81, 0, 82, 180,
	228, 60, 0, 76, 0, 0, 0, 0, 71, 70,
	0, 79, 68, 69, 72, 0, 24, 73, 0, 0,
	77, 0, 0, 0, 78, 74, 75, 0, 67, 0,
	0, 0, 0, 23, 0, 80, 0, 0, 81, 0,
	82, 180, 76, 60, 0, 0, 0, 71, 70, 0,
	79, 68, 69, 72, 0, 24, 73, 0, 0, 77,
	0, 67, 0, 78, 74, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 80, 76, 0, 81, 0, 82,
	71, 70, 60, 79, 68, 69, 72, 0, 0, 73,
	0, 23, 77, 0, 0, 0, 78, 74, 75, 0,
	76, 0, 0, 0, 0, 71, 70, 80, 79, 0,
	81, 72, 82, 24, 73, 60, 0, 77, 283, 0,
	0, 78, 74, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 180, 228,
}
var yyPact = [...]int{

	45, -1000, 617, 563, 558, 559, -1000, 480, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	413, 480, 480, 480, 480, 470, 404, 480, 502, 467,
	403, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 500, 556, 555,
	554, 553, 466, 464, 459, 457, 455, 452, 386, 385,
	480, 451, 449, 446, 445, 380, 349, 474, 551, 550,
	549, 545, 544, 543, 541, 537, 535, 530, 529, 527,
	480, 526, 523, 444, 438, -1000, -1000, 290, -1000, -1000,
	53, 258, 250, 80, 229, 184, -1000, 351, 228, 211,
	332, -1000, 52, -1000, -1000, -1000, -1000, -1000, 44, 831,
	665, 750, 47, 47, -1000, 491, -1000, 222, 219, 341,
	332, 792, 889, -1000, 43, -1000, 188, 210, -1, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1, -1000, -1000, -1000, -1000, -1000, 497, 169,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 437, -1000, 531,
	-1000, -1000, 496, -1000, -1000, 436, -1000, 495, -1000, -1000,
	-1000, -1000, 434, 332, -1000, -1000, -1000, -1000, 864, 864,
	480, 350, 284, 430, 52, -1000, -1000, -1000, -1000, 339,
	489, 394, -1000, -1000, -1000, -1000, -1000, -1000, 289, 488,
	474, 518, 428, 831, -1000, -1000, -1000, -1000, 423, 665,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	480, -1000, 288, 120, 478, 28, 28, 480, 480, 418,
	750, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 415,
	47, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 408, 395, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 389, 222, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 283, 480, -1000, 378, -1000, -1000, 23, -1000,
	-1000, -1000, 223, 516, 373, 332, -1000, -1000, -1000, -1000,
	864, 864, 363, 792, -1000, -1000, -1000, -1000, -1000, 362,
	889, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 356, 43, -1000, -1000,
	-1000, -1000, -1000, 515, 352, 188, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 474, -1000, -1000, 204, -1000, -1000,
	199, -1000, -1000, 337, 864, -1000, 331, 214, -1000, -1000,
	-1000, -1000, -1000, 39, -1000, -1000, -1000, -1000, 469, 196,
	191, -1000, -1000, -1000, -1000, -1000, 209, -1000, 53, 185,
	180, -1000, 138, 137, -1000, -1000, 131, 153, -1, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 476, -1,
	-1000, -1000, 750, -1000, -1000, -1000, 295, 294, -1000, -1000,
	-1000, -1000, -1000, -1000, 87, -1000, -1000, 280, -1000, -1000,
	-1000, -1000, -1000, -1000, 276, 39, -1000, -1000, -1000, -1000,
	28, 270, 469, -1000, 480, 480, 480, -1000, -1000, -1000,
	-1000, -1000, 84, 478, 480, 514, -1000, -1000, -1000, 70,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 263,
	476, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 708, -1000, -1000, -1000, 253, -1000, -1000, 79, -1000,
	-1000, 143, 83, 75, -1000, 316, 61, 40, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 268,
	-1000, -1000, -1000, -1000, 478, -1000, -1000, -1000, -1000, 58,
	-1000,
}
var yyPgo = [...]int{

	0, 17, 2, 16, 287, 775, 774, 773, 489, 772,
	765, 763, 66, 0, 760, 759, 53, 758, 756, 755,
	319, 754, 753, 44, 749, 748, 47, 747, 745, 46,
	25, 744, 736, 733, 42, 36, 35, 33, 32, 31,
	30, 29, 28, 22, 43, 732, 731, 730, 727, 719,
	716, 40, 711, 710, 706, 705, 9, 704, 693, 690,
	689, 15, 132, 688, 263, 683, 7, 682, 680, 676,
	21, 11, 674, 671, 669, 48, 668, 12, 665, 663,
	45, 19, 74, 10, 662, 656, 654, 8, 653, 650,
	642, 641, 640, 4, 26, 20, 639, 637, 636, 23,
	635, 634, 633, 24, 631, 627, 5, 18, 13, 626,
	625, 622, 619, 618, 41, 52, 49, 616, 614, 607,
	34, 606, 605, 604, 27, 598, 596, 595, 38, 592,
	584, 582, 39, 581, 579, 578, 577, 57, 576, 37,
	575, 573, 569, 1, 563, 14,
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
	79, 80, 80, 80, 80, 80, 3, 3, 83, 81,
	81, 84, 85, 85, 86, 86, 87, 87, 87, 87,
	87, 87, 87, 87, 89, 90, 35, 91, 92, 92,
	77, 77, 93, 93, 93, 93, 93, 93, 93, 96,
	44, 97, 97, 98, 98, 99, 99, 99, 99, 99,
	99, 99, 99, 99, 99, 99, 99, 99, 99, 100,
	39, 39, 101, 101, 102, 102, 103, 103, 103, 103,
	103, 103, 105, 106, 106, 106, 106, 106, 106, 106,
	106, 106, 104, 104, 109, 110, 110, 18, 111, 112,
	112, 113, 113, 114, 114, 114, 114, 114, 115, 116,
	42, 117, 118, 118, 119, 119, 120, 120, 120, 120,
	120, 43, 121, 122, 122, 123, 123, 124, 124, 124,
	124, 33, 125, 126, 126, 127, 127, 128, 128, 128,
	34, 129, 130, 131, 131, 107, 107, 108, 132, 132,
	132, 132, 132, 132, 132, 132, 132, 132, 132, 133,
	38, 38, 135, 135, 135, 135, 135, 135, 135, 134,
	134, 36, 136, 137, 138, 138, 139, 139, 139, 139,
	139, 139, 139, 139, 139, 139, 139, 139, 95, 4,
	4, 2, 1, 1, 94, 37, 140, 88, 88, 141,
	142, 142, 143, 143, 143, 144, 12, 13, 10, 11,
	17, 82, 145, 145, 47,
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
	2, 1, 1, 1, 1, 1, 1, 1, 3, 2,
	4, 2, 0, 1, 1, 2, 3, 3, 3, 1,
	1, 1, 1, 1, 3, 3, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 2,
	4, 0, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 4, 1, 1, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 1, 1, 1, 3,
	3, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	1, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	4, 2, 1, 1, 2, 3, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	2, 4, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 4, 2, 2, 4, 2,
	1, 2, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 1, 5, 3,
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
	-104, -44, -105, 51, 10, -73, -12, -13, -74, -62,
	-64, -75, -76, 39, -118, -119, -120, -12, -13, -62,
	-115, -116, -122, -123, -124, -12, -13, -62, -20, -97,
	-98, -99, -12, -13, -62, -64, -34, -35, -36, -37,
	-38, -39, -41, -75, -42, -43, -68, -69, -70, -12,
	-13, -71, -62, 57, -59, -60, -61, -12, -13, -62,
	10, 5, 9, -23, 7, -26, 9, 5, -29, 9,
	5, 9, -114, -30, -31, -20, -30, -4, 8, 8,
	9, -51, 10, 8, 5, 9, -80, 10, 8, 5,
	-3, 4, 9, -128, 9, -132, -4, 10, 8, -2,
	52, 6, -2, -1, 47, 48, -1, -4, -4, 9,
	-93, 9, -139, 9, 9, 9, -103, 10, 8, -4,
	9, -75, 8, 4, 9, -120, -30, -30, 9, -124,
	9, -99, 9, -70, 4, 9, -61, -3, 10, 10,
	9, -20, 9, 10, -54, -55, -56, -12, -13, -57,
	66, -85, -86, -87, 17, 63, 45, -88, -71, -89,
	-81, -90, -141, 69, 67, 18, 10, 10, 10, -65,
	-66, -12, -13, 10, 10, 10, 10, 10, 10, -109,
	-110, -106, -12, -13, -62, -83, -94, -95, -40, -107,
	-108, -77, 9, 9, 10, -145, 9, -56, -1, 9,
	-87, -4, -4, -4, 10, 8, -2, -4, 4, 9,
	-66, 9, -106, 9, 9, 10, 10, 10, 10, -142,
	-143, -12, -13, -144, 46, 10, 10, 9, -143, -2,
	10,
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
	0, 0, 0, 0, 0, 1, 5, 0, 309, 20,
	0, 0, 0, 0, 0, 0, 41, 0, 0, 0,
	219, 63, 66, 19, 26, 35, 218, 65, 0, 253,
	0, 158, 0, 0, 280, 0, 190, 192, 0, 116,
	232, 243, 171, 103, 106, 84, 87, 0, 136, 137,
	128, 252, 261, 157, 292, 316, 289, 290, 189, 122,
	231, 242, 169, 105, 86, 2, 3, 6, 0, 0,
	22, 24, 25, 328, 329, 326, 332, 0, 327, 0,
	27, 30, 0, 32, 33, 0, 36, 0, 39, 40,
	29, 330, 0, 220, 221, 223, 224, 225, 43, 43,
	0, 0, 0, 0, 67, 68, 70, 71, 72, 0,
	0, 0, 129, 131, 132, 133, 134, 135, 0, 0,
	0, 0, 0, 254, 255, 257, 258, 259, 0, 262,
	263, 268, 269, 270, 271, 272, 273, 274, 275, 276,
	0, 278, 0, 0, 0, 0, 0, 0, 0, 0,
	159, 160, 162, 163, 164, 165, 166, 167, 168, 0,
	293, 294, 296, 297, 298, 299, 300, 301, 302, 303,
	304, 305, 306, 307, 0, 0, 282, 283, 284, 285,
	286, 287, 288, 0, 193, 194, 196, 197, 198, 199,
	200, 201, 0, 0, 94, 0, 117, 118, 119, 120,
	121, 123, 0, 0, 0, 233, 234, 236, 237, 238,
	43, 43, 0, 244, 245, 247, 248, 249, 250, 0,
	172, 173, 175, 176, 177, 178, 179, 180, 181, 182,
	183, 184, 185, 186, 187, 188, 0, 107, 108, 110,
	111, 112, 113, 0, 0, 88, 89, 91, 92, 93,
	334, 310, 21, 23, 0, 28, 34, 0, 37, 42,
	0, 217, 222, 0, 44, 61, 0, 0, 228, 229,
	64, 69, 73, 76, 75, 127, 130, 139, 142, 0,
	0, 141, 251, 256, 260, 264, 0, 97, 0, 0,
	0, 311, 0, 0, 312, 313, 0, 0, 96, 156,
	161, 291, 295, 315, 281, 191, 195, 212, 0, 202,
	115, 124, 0, 126, 230, 235, 0, 0, 241, 246,
	170, 174, 104, 109, 0, 85, 90, 0, 31, 38,
	226, 62, 227, 95, 0, 77, 78, 80, 81, 82,
	0, 0, 143, 144, 0, 0, 0, 149, 150, 151,
	152, 153, 0, 0, 0, 0, 331, 138, 277, 0,
	99, 101, 102, 265, 266, 267, 314, 308, 279, 0,
	214, 215, 203, 204, 205, 206, 207, 208, 209, 210,
	211, 0, 239, 240, 114, 0, 74, 79, 0, 140,
	145, 0, 0, 0, 317, 0, 0, 0, 319, 98,
	100, 213, 216, 125, 333, 83, 146, 147, 148, 0,
	320, 322, 323, 324, 0, 154, 155, 318, 321, 0,
	325,
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
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:466
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:478
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:483
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:490
		{
			pop(yylex)
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:495
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:512
		{
			yyVAL.token = yyDollar[1].token
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:513
		{
			yyVAL.token = yyDollar[1].token
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:516
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:523
		{
			pop(yylex)
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:526
		{
			pop(yylex)
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:531
		{
			if push(yylex, meta.NewType(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:545
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:554
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetValueRange(r)) {
				goto ret1
			}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:563
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:576
		{
			if set(yylex, meta.SetFractionDigits(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:583
		{
			if set(yylex, meta.SetPattern(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:590
		{
			pop(yylex)
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:595
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:619
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:626
		{
			pop(yylex)
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:654
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:661
		{
			pop(yylex)
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:664
		{
			pop(yylex)
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:684
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:703
		{
			pop(yylex)
		}
	case 213:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:706
		{
			pop(yylex)
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:718
		{
			pop(yylex)
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:723
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:740
		{
			pop(yylex)
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:743
		{
			pop(yylex)
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:748
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:755
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 230:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:765
		{
			pop(yylex)
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:770
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:787
		{
			pop(yylex)
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:790
		{
			pop(yylex)
		}
	case 241:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:798
		{
			pop(yylex)
		}
	case 242:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:803
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 251:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:824
		{
			pop(yylex)
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:829
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 260:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:848
		{
			pop(yylex)
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:853
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 265:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:868
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 266:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:873
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 267:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:880
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 279:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:900
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 280:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:907
		{
			pop(yylex)
		}
	case 281:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:910
		{
			pop(yylex)
		}
	case 289:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:925
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 290:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:930
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 291:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:937
		{
			pop(yylex)
		}
	case 292:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:942
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 308:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:971
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:978
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 310:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:981
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 311:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:986
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 312:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:996
		{
			yyVAL.boolean = true
		}
	case 313:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:997
		{
			yyVAL.boolean = false
		}
	case 314:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1000
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 315:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1010
		{
			pop(yylex)
		}
	case 316:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1015
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 317:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1022
		{
			pop(yylex)
		}
	case 318:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1025
		{
			pop(yylex)
		}
	case 319:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1030
		{
			if push(yylex, meta.NewEnum(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 325:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1045
		{
			if set(yylex, meta.SetEnumValue(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 326:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1052
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 327:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1059
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 328:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1066
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 329:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1073
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 330:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1080
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 331:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1087
		{
			if set(yylex, meta.SetUnits(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 334:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1098
		{

		}
	}
	goto yystack /* stack new state and value */
}
