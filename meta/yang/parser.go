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
const kywd_current = 57413
const kywd_obsolete = 57414
const kywd_deprecated = 57415

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
	"kywd_current",
	"kywd_obsolete",
	"kywd_deprecated",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:1133

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 1110

var yyAct = [...]int{

	250, 520, 396, 13, 269, 247, 13, 480, 468, 348,
	443, 451, 246, 161, 400, 254, 132, 353, 344, 44,
	326, 43, 41, 259, 42, 318, 40, 258, 268, 39,
	38, 309, 284, 218, 37, 225, 271, 185, 16, 191,
	205, 176, 371, 155, 169, 249, 36, 197, 12, 215,
	35, 12, 452, 463, 24, 527, 215, 190, 214, 339,
	518, 153, 372, 256, 304, 45, 24, 20, 401, 402,
	20, 306, 266, 96, 97, 98, 26, 193, 194, 505,
	454, 504, 163, 163, 162, 162, 531, 415, 26, 414,
	517, 526, 350, 158, 395, 515, 394, 516, 453, 494,
	153, 174, 462, 180, 461, 153, 188, 385, 201, 384,
	213, 477, 25, 192, 210, 221, 228, 253, 262, 262,
	232, 276, 476, 287, 25, 298, 312, 321, 329, 153,
	347, 234, 356, 270, 270, 153, 281, 206, 157, 170,
	264, 264, 153, 278, 231, 24, 173, 341, 179, 340,
	211, 187, 338, 200, 337, 158, 153, 336, 335, 209,
	220, 227, 334, 261, 261, 315, 275, 26, 286, 174,
	297, 311, 320, 328, 333, 346, 180, 355, 332, 222,
	237, 255, 272, 314, 303, 188, 207, 300, 291, 475,
	350, 323, 192, 474, 466, 3, 473, 201, 361, 514,
	157, 24, 171, 25, 29, 210, 472, 170, 398, 248,
	440, 465, 11, 363, 173, 11, 464, 366, 221, 437,
	294, 179, 370, 26, 4, 228, 436, 181, 206, 232,
	187, 387, 358, 233, 373, 267, 267, 182, 280, 153,
	234, 211, 200, 399, 378, 383, 167, 24, 160, 24,
	209, 390, 407, 231, 397, 153, 24, 403, 262, 25,
	392, 382, 511, 220, 24, 153, 215, 166, 214, 26,
	227, 26, 153, 270, 159, 165, 152, 207, 26, 164,
	264, 525, 222, 499, 287, 409, 26, 496, 293, 237,
	82, 153, 493, 153, 252, 492, 192, 244, 380, 131,
	379, 130, 156, 261, 129, 25, 128, 25, 439, 312,
	172, 255, 178, 438, 25, 186, 413, 199, 321, 153,
	213, 153, 25, 208, 219, 226, 329, 260, 260, 286,
	274, 122, 285, 121, 296, 310, 319, 327, 433, 345,
	423, 354, 233, 427, 347, 341, 429, 340, 315, 291,
	338, 430, 337, 356, 311, 336, 335, 424, 425, 418,
	334, 428, 431, 320, 156, 419, 314, 120, 420, 119,
	434, 328, 333, 251, 528, 267, 332, 24, 172, 435,
	426, 446, 323, 422, 107, 178, 106, 417, 102, 346,
	101, 92, 376, 91, 186, 456, 471, 412, 355, 26,
	411, 24, 410, 408, 367, 406, 199, 24, 177, 458,
	230, 525, 265, 265, 208, 279, 483, 289, 391, 302,
	490, 509, 331, 26, 24, 375, 445, 219, 24, 26,
	389, 487, 377, 491, 226, 25, 369, 419, 488, 151,
	150, 470, 127, 446, 489, 126, 26, 125, 133, 495,
	26, 124, 485, 497, 118, 448, 89, 117, 116, 25,
	456, 482, 500, 498, 506, 25, 115, 260, 471, 114,
	113, 93, 94, 95, 458, 99, 510, 105, 103, 189,
	483, 203, 25, 100, 490, 255, 25, 512, 445, 229,
	362, 263, 263, 285, 277, 487, 288, 407, 301, 313,
	322, 330, 488, 349, 398, 357, 523, 508, 489, 90,
	134, 123, 90, 470, 24, 386, 485, 6, 310, 230,
	523, 529, 381, 88, 368, 482, 365, 319, 530, 359,
	432, 147, 108, 486, 364, 327, 26, 24, 171, 360,
	29, 104, 24, 421, 388, 24, 177, 149, 148, 146,
	145, 522, 265, 345, 24, 144, 255, 143, 189, 26,
	142, 192, 354, 78, 26, 522, 141, 26, 73, 72,
	140, 81, 25, 139, 74, 138, 26, 75, 289, 137,
	79, 306, 136, 135, 80, 76, 77, 112, 111, 110,
	444, 109, 86, 85, 524, 25, 519, 486, 229, 69,
	25, 192, 244, 25, 24, 469, 460, 24, 59, 257,
	58, 273, 25, 78, 60, 241, 235, 224, 73, 72,
	331, 81, 70, 71, 74, 481, 26, 75, 242, 26,
	79, 263, 223, 56, 80, 76, 77, 217, 216, 24,
	55, 374, 317, 214, 316, 82, 65, 241, 83, 308,
	84, 192, 444, 62, 192, 244, 62, 288, 307, 64,
	242, 26, 25, 239, 240, 25, 184, 183, 31, 479,
	478, 292, 290, 283, 282, 61, 325, 469, 324, 66,
	245, 57, 313, 459, 457, 393, 192, 455, 62, 481,
	450, 322, 404, 405, 449, 212, 204, 25, 54, 330,
	305, 299, 295, 63, 343, 342, 67, 467, 238, 352,
	351, 68, 447, 442, 441, 521, 69, 349, 87, 202,
	7, 24, 49, 198, 29, 196, 357, 195, 32, 521,
	78, 48, 47, 46, 34, 73, 72, 52, 81, 70,
	71, 74, 416, 26, 75, 33, 175, 79, 28, 50,
	51, 80, 76, 77, 168, 27, 154, 21, 22, 23,
	19, 18, 82, 17, 15, 83, 14, 84, 10, 9,
	62, 30, 69, 53, 8, 5, 7, 24, 49, 25,
	29, 2, 1, 0, 0, 0, 78, 0, 0, 484,
	0, 73, 72, 52, 81, 70, 71, 74, 0, 26,
	75, 0, 0, 79, 0, 50, 51, 80, 76, 77,
	0, 0, 0, 0, 22, 23, 0, 0, 82, 69,
	0, 83, 0, 84, 24, 0, 62, 30, 0, 53,
	0, 243, 241, 78, 236, 25, 0, 0, 73, 72,
	0, 81, 70, 71, 74, 242, 26, 75, 239, 240,
	79, 0, 0, 484, 80, 76, 77, 0, 0, 0,
	0, 69, 0, 513, 0, 82, 24, 0, 83, 0,
	84, 192, 244, 62, 241, 78, 0, 0, 0, 0,
	73, 72, 25, 81, 70, 71, 74, 242, 26, 75,
	0, 0, 79, 0, 0, 0, 80, 76, 77, 0,
	0, 501, 502, 503, 0, 0, 0, 82, 0, 0,
	83, 507, 84, 192, 244, 62, 69, 0, 0, 0,
	0, 24, 0, 0, 25, 0, 0, 0, 0, 241,
	78, 0, 0, 0, 0, 73, 72, 0, 81, 70,
	71, 74, 242, 26, 75, 0, 0, 79, 0, 0,
	0, 80, 76, 77, 0, 0, 0, 0, 69, 0,
	0, 0, 82, 24, 0, 83, 0, 84, 192, 244,
	62, 0, 78, 0, 0, 0, 0, 73, 72, 25,
	81, 70, 71, 74, 0, 26, 75, 0, 0, 79,
	306, 0, 0, 80, 76, 77, 0, 0, 0, 0,
	69, 0, 0, 0, 82, 24, 0, 83, 0, 84,
	192, 244, 62, 0, 78, 0, 0, 0, 0, 73,
	72, 25, 81, 70, 71, 74, 0, 26, 75, 0,
	0, 79, 0, 0, 0, 80, 76, 77, 0, 0,
	24, 0, 215, 0, 214, 0, 82, 0, 241, 83,
	0, 84, 0, 0, 62, 69, 0, 0, 0, 0,
	0, 242, 26, 25, 239, 240, 0, 0, 0, 78,
	0, 0, 0, 0, 73, 72, 0, 81, 70, 71,
	74, 0, 0, 75, 0, 0, 79, 192, 244, 62,
	80, 76, 77, 0, 0, 0, 213, 0, 25, 0,
	0, 82, 0, 0, 83, 0, 84, 0, 0, 62,
}
var yyPact = [...]int{

	170, -1000, 765, 589, 588, 709, -1000, 507, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 383, 507, 507, 507, 2, 507, 475, 380, 507,
	536, 469, 376, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 527,
	587, 585, 584, 583, 462, 461, 458, 450, 449, 446,
	359, 323, 507, 443, 439, 437, 434, 296, 291, 504,
	579, 578, 575, 571, 569, 566, 562, 556, 553, 551,
	546, 545, 507, 544, 543, 432, 431, -1000, -1000, 266,
	-1000, -1000, 244, 264, 238, 74, 269, 265, 257, 236,
	189, -1000, 533, 217, 227, 54, -1000, 416, -1000, -1000,
	-1000, -1000, -1000, 42, 993, 812, 909, 1028, 1028, -1000,
	595, -1000, 237, 210, 951, 54, 592, 542, -1000, 133,
	-1000, 502, 222, 6, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 6, -1000, -1000,
	-1000, -1000, -1000, 524, 530, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 483, -1000, -1000, -1000, -1000, 525, -1000,
	-1000, 521, -1000, -1000, -1000, 395, -1000, 519, -1000, -1000,
	-1000, -1000, -1000, 427, 54, -1000, -1000, -1000, -1000, -1000,
	1048, 1048, 507, 417, 384, 423, 416, -1000, -1000, -1000,
	-1000, -1000, 290, 517, 252, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 99, 510, 504, 540, 421, 993, -1000, -1000,
	-1000, -1000, -1000, 409, 812, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 507, -1000, 86, 202,
	498, 21, 21, 507, 507, 396, 909, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 394, 1028, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 393, 391, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 388, 237, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 79, 507, -1000, 378, -1000, -1000, -1000, 32,
	1048, -1000, -1000, -1000, -1000, 360, 539, 374, 54, -1000,
	-1000, -1000, -1000, -1000, 1048, 1048, 371, 592, -1000, -1000,
	-1000, -1000, -1000, -1000, 352, 542, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 342, 133, -1000, -1000, -1000, -1000, -1000, -1000,
	526, 329, 502, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 504, -1000, -1000, 216, -1000, -1000, 209, -1000,
	-1000, 304, 1048, 299, 200, -1000, -1000, -1000, -1000, -1000,
	389, -1000, -1000, -1000, -1000, 35, 206, 201, -1000, -1000,
	-1000, -1000, -1000, 184, -1000, 244, 196, 186, -1000, 183,
	179, -1000, -1000, 112, 101, 6, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 627, 6, -1000, -1000, -1000,
	909, -1000, -1000, -1000, 286, 283, -1000, -1000, -1000, -1000,
	-1000, -1000, 89, -1000, -1000, 75, -1000, -1000, -1000, -1000,
	-1000, 278, 389, -1000, -1000, -1000, -1000, -1000, 21, 274,
	35, -1000, 507, 507, 507, -1000, -1000, -1000, -1000, -1000,
	71, 498, 507, 503, -1000, -1000, -1000, 412, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 253, 627,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 854, -1000, -1000, -1000, 190, -1000, -1000, 85, -1000,
	-1000, 87, 80, 50, -1000, 235, 81, 45, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 365,
	-1000, -1000, -1000, -1000, -1000, 498, -1000, -1000, -1000, -1000,
	76, -1000,
}
var yyPgo = [...]int{

	0, 14, 2, 16, 448, 782, 781, 775, 517, 774,
	769, 768, 209, 45, 0, 766, 764, 38, 763, 761,
	760, 64, 757, 756, 43, 755, 754, 44, 748, 746,
	41, 42, 62, 745, 734, 50, 46, 34, 30, 29,
	26, 22, 24, 21, 19, 65, 733, 732, 731, 728,
	727, 725, 47, 723, 719, 714, 713, 10, 712, 711,
	710, 709, 17, 373, 708, 294, 707, 8, 706, 705,
	704, 18, 9, 703, 702, 701, 59, 700, 12, 698,
	696, 40, 23, 72, 36, 695, 694, 690, 11, 687,
	684, 683, 681, 680, 5, 117, 15, 679, 678, 676,
	20, 675, 674, 673, 32, 672, 671, 7, 28, 4,
	670, 669, 668, 667, 666, 37, 57, 39, 659, 658,
	649, 31, 646, 644, 642, 25, 640, 638, 637, 33,
	633, 632, 617, 35, 616, 614, 611, 610, 63, 609,
	27, 608, 606, 596, 1, 594, 13,
}
var yyR1 = [...]int{

	0, 5, 6, 6, 7, 7, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	22, 9, 9, 23, 23, 24, 24, 24, 25, 26,
	26, 17, 27, 27, 27, 27, 27, 15, 28, 29,
	29, 30, 30, 30, 30, 16, 16, 31, 31, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 32, 32, 20, 20, 49,
	50, 50, 51, 51, 52, 52, 52, 52, 53, 53,
	54, 55, 55, 56, 56, 57, 57, 57, 57, 58,
	47, 47, 59, 60, 60, 61, 61, 62, 62, 62,
	62, 41, 63, 64, 65, 65, 66, 66, 67, 67,
	67, 46, 46, 68, 69, 69, 70, 70, 71, 71,
	71, 71, 71, 72, 42, 74, 74, 74, 74, 74,
	74, 74, 74, 73, 75, 75, 76, 77, 33, 79,
	80, 80, 81, 81, 81, 81, 81, 81, 3, 3,
	84, 82, 82, 85, 86, 86, 87, 87, 88, 88,
	88, 88, 88, 88, 88, 88, 13, 13, 13, 90,
	91, 36, 92, 93, 93, 78, 78, 94, 94, 94,
	94, 94, 94, 94, 94, 97, 45, 98, 98, 99,
	99, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	100, 100, 100, 100, 100, 100, 101, 40, 40, 102,
	102, 103, 103, 104, 104, 104, 104, 104, 104, 104,
	106, 107, 107, 107, 107, 107, 107, 107, 107, 107,
	107, 105, 105, 110, 111, 111, 19, 112, 113, 113,
	114, 114, 115, 115, 115, 115, 115, 115, 116, 117,
	43, 118, 119, 119, 120, 120, 121, 121, 121, 121,
	121, 121, 44, 122, 123, 123, 124, 124, 125, 125,
	125, 125, 125, 34, 126, 127, 127, 128, 128, 129,
	129, 129, 129, 35, 130, 131, 132, 132, 108, 108,
	109, 133, 133, 133, 133, 133, 133, 133, 133, 133,
	133, 133, 133, 134, 39, 39, 136, 136, 136, 136,
	136, 136, 136, 136, 135, 135, 37, 137, 138, 139,
	139, 140, 140, 140, 140, 140, 140, 140, 140, 140,
	140, 140, 140, 140, 96, 4, 4, 2, 1, 1,
	95, 38, 141, 89, 89, 142, 143, 143, 144, 144,
	144, 144, 145, 12, 14, 10, 11, 18, 83, 146,
	146, 48,
}
var yyR2 = [...]int{

	0, 3, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 1, 2, 1, 1, 1, 2, 1,
	2, 3, 1, 3, 1, 1, 1, 4, 2, 1,
	2, 3, 1, 1, 1, 2, 4, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 1, 3,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	1, 3, 3, 2, 2, 4, 1, 2, 1, 1,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 3, 4, 0, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 2, 4, 2, 4, 2,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 2, 4, 2, 0, 1, 1, 2, 3, 3,
	3, 1, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 0, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 4, 1, 1, 2, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 1, 3, 3, 2, 2,
	4, 2, 0, 1, 1, 2, 1, 1, 1, 1,
	3, 3, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 4, 2, 1, 1, 2, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 2, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 4, 2, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 3, 1, 1, 1,
	3, 4, 2, 2, 4, 2, 1, 2, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 1,
	5, 3,
}
var yyChk = [...]int{

	-1000, -5, -6, 25, 54, -7, -8, 11, -9, -10,
	-11, -12, -13, -14, -15, -16, -17, -18, -19, -20,
	-21, -22, 49, 50, 12, 70, 34, -25, -28, 15,
	62, -112, -49, -33, -34, -35, -36, -37, -38, -39,
	-40, -41, -42, -43, -44, -45, -46, -47, -48, 13,
	40, 41, 28, 64, -79, -126, -130, -92, -137, -141,
	-135, -101, 61, -73, -118, -122, -97, -68, -59, 7,
	30, 31, 27, 26, 32, 35, 43, 44, 21, 38,
	42, 29, 53, 56, 58, 4, 4, 9, -8, -4,
	5, 10, 8, -4, -4, -4, 71, 72, 73, -4,
	8, 10, 8, -4, 5, 8, 10, 8, 5, 4,
	4, 4, 4, 8, 8, 8, 8, 8, 8, 10,
	8, 10, 8, -4, 8, 8, 8, 8, 10, 8,
	10, 8, -3, -4, 6, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, -4, 4, 4,
	8, 8, 10, 55, -23, -24, -12, -13, -14, 10,
	10, -146, 10, 8, 10, 10, 10, 10, -26, -27,
	-17, 13, -12, -13, -14, -29, -30, 13, -12, -13,
	-14, 10, 10, -113, -114, -115, -12, -13, -14, -63,
	-116, -117, 59, 23, 24, -50, -51, -52, -53, -12,
	-13, -14, -54, 65, -80, -81, -82, -83, -12, -13,
	-14, -84, -85, 68, 16, 14, -127, -128, -129, -12,
	-13, -14, -21, -131, -132, -133, -12, -13, -14, -63,
	-65, -108, -109, -95, -96, -134, 22, -21, -64, 36,
	37, 20, 33, 19, 60, -93, -78, -94, -12, -13,
	-14, -63, -65, -95, -96, -21, -138, -139, -140, -82,
	-12, -13, -14, -63, -41, -65, -83, -95, -108, -109,
	-96, -84, -138, -136, -12, -13, -14, -63, -41, -65,
	-95, -96, -102, -103, -104, -12, -13, -14, -63, -65,
	-105, -45, -106, 51, 10, -74, -12, -13, -14, -75,
	-32, -63, -65, -76, -21, -77, 39, -119, -120, -121,
	-12, -13, -14, -63, -116, -117, -123, -124, -125, -12,
	-13, -14, -63, -21, -98, -99, -100, -12, -13, -14,
	-63, -65, -35, -36, -37, -38, -39, -40, -42, -76,
	-43, -44, -69, -70, -71, -12, -13, -14, -72, -63,
	57, -60, -61, -62, -12, -13, -14, -63, 10, 5,
	9, -24, 7, -27, 9, 5, -30, 9, 5, 9,
	-115, -31, -32, -31, -4, 8, 8, 9, -52, 10,
	8, 5, 9, -81, 10, 8, 5, -3, 4, 9,
	-129, 9, -133, -4, 10, 8, -2, 52, 6, -2,
	-1, 47, 48, -1, -4, -4, 9, -94, 9, -140,
	9, 9, 9, -104, 10, 8, -4, 9, -76, -21,
	8, 4, 9, -121, -31, -31, 9, -125, 9, -100,
	9, -71, 4, 9, -62, -3, 10, 10, 9, 9,
	10, -55, -56, -57, -12, -13, -14, -58, 66, -86,
	-87, -88, 17, 63, 45, -89, -72, -90, -82, -91,
	-142, 69, 67, 18, 10, 10, 10, -66, -67, -12,
	-13, -14, 10, 10, 10, 10, 10, 10, -110, -111,
	-107, -12, -13, -14, -63, -84, -95, -96, -41, -108,
	-109, -78, 9, 9, 10, -146, 9, -57, -1, 9,
	-88, -4, -4, -4, 10, 8, -2, -4, 4, 9,
	-67, 9, -107, 9, 9, 10, 10, 10, 10, -143,
	-144, -12, -13, -14, -145, 46, 10, 10, 9, -144,
	-2, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 5, 0,
	335, 21, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 45, 0, 0, 0, 238, 67, 70, 20, 28,
	38, 237, 69, 0, 275, 0, 173, 0, 0, 304,
	0, 207, 209, 0, 125, 252, 264, 187, 111, 114,
	90, 93, 0, 148, 149, 139, 274, 284, 172, 317,
	342, 314, 315, 206, 133, 251, 263, 185, 113, 92,
	2, 3, 6, 0, 0, 23, 25, 26, 27, 355,
	356, 353, 359, 0, 166, 167, 168, 354, 0, 29,
	32, 0, 34, 35, 36, 0, 39, 0, 42, 43,
	44, 31, 357, 0, 239, 240, 242, 243, 244, 245,
	47, 47, 0, 0, 0, 0, 71, 72, 74, 75,
	76, 77, 0, 0, 0, 140, 142, 143, 144, 145,
	146, 147, 0, 0, 0, 0, 0, 276, 277, 279,
	280, 281, 282, 0, 285, 286, 291, 292, 293, 294,
	295, 296, 297, 298, 299, 300, 0, 302, 0, 0,
	0, 0, 0, 0, 0, 0, 174, 175, 177, 178,
	179, 180, 181, 182, 183, 184, 0, 318, 319, 321,
	322, 323, 324, 325, 326, 327, 328, 329, 330, 331,
	332, 333, 0, 0, 306, 307, 308, 309, 310, 311,
	312, 313, 0, 210, 211, 213, 214, 215, 216, 217,
	218, 219, 0, 0, 101, 0, 126, 127, 128, 129,
	130, 131, 132, 134, 65, 0, 0, 0, 253, 254,
	256, 257, 258, 259, 47, 47, 0, 265, 266, 268,
	269, 270, 271, 272, 0, 188, 189, 191, 192, 193,
	194, 195, 196, 197, 198, 199, 200, 201, 202, 203,
	204, 205, 0, 115, 116, 118, 119, 120, 121, 122,
	0, 0, 94, 95, 97, 98, 99, 100, 361, 336,
	22, 24, 0, 30, 37, 0, 40, 46, 0, 236,
	241, 0, 48, 0, 0, 248, 249, 68, 73, 78,
	81, 80, 138, 141, 151, 154, 0, 0, 153, 273,
	278, 283, 287, 0, 104, 0, 0, 0, 337, 0,
	0, 338, 339, 0, 0, 103, 171, 176, 316, 320,
	341, 305, 208, 212, 231, 0, 220, 124, 135, 66,
	0, 137, 250, 255, 0, 0, 262, 267, 186, 190,
	112, 117, 0, 91, 96, 0, 33, 41, 246, 247,
	102, 0, 82, 83, 85, 86, 87, 88, 0, 0,
	155, 156, 0, 0, 0, 161, 162, 163, 164, 165,
	0, 0, 0, 0, 358, 150, 301, 0, 106, 108,
	109, 110, 288, 289, 290, 340, 334, 303, 0, 233,
	234, 221, 222, 223, 224, 225, 226, 227, 228, 229,
	230, 0, 260, 261, 123, 0, 79, 84, 0, 152,
	157, 0, 0, 0, 343, 0, 0, 0, 345, 105,
	107, 232, 235, 136, 360, 89, 158, 159, 160, 0,
	346, 348, 349, 350, 351, 0, 169, 170, 344, 347,
	0, 352,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
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
		//line parser.y:152
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
		//line parser.y:160
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
		//line parser.y:174
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:194
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:201
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:204
		{
			pop(yylex)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:218
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:229
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:243
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:248
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:265
		{
			pop(yylex)
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:268
		{
			pop(yylex)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:298
		{
			pop(yylex)
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:301
		{
			pop(yylex)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:306
		{
			if push(yylex, meta.NewExtension(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:326
		{
			pop(yylex)
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:329
		{
			pop(yylex)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:334
		{
			if push(yylex, meta.NewExtensionArg(peek(yylex).(*meta.Extension), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:354
		{
			if set(yylex, meta.SetYinElement(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:361
		{
			pop(yylex)
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:364
		{
			pop(yylex)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:369
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:389
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:396
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:403
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:410
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:413
		{
			pop(yylex)
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:426
		{
			pop(yylex)
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:429
		{
			pop(yylex)
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:434
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:455
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:465
		{
			pop(yylex)
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:480
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:492
		{
			pop(yylex)
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:497
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:504
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:509
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:527
		{
			yyVAL.token = yyDollar[1].token
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:528
		{
			yyVAL.token = yyDollar[1].token
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:531
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:538
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:541
		{
			pop(yylex)
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:546
		{
			if push(yylex, meta.NewType(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:560
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:569
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetValueRange(r)) {
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:578
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:595
		{
			if set(yylex, meta.SetFractionDigits(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:602
		{
			if set(yylex, meta.SetPattern(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:609
		{
			pop(yylex)
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:614
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:640
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:647
		{
			pop(yylex)
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:676
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:683
		{
			pop(yylex)
		}
	case 208:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:686
		{
			pop(yylex)
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:707
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:727
		{
			pop(yylex)
		}
	case 232:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:730
		{
			pop(yylex)
		}
	case 236:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:742
		{
			pop(yylex)
		}
	case 237:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:747
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 246:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:765
		{
			pop(yylex)
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:768
		{
			pop(yylex)
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:773
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:780
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 250:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:790
		{
			pop(yylex)
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:795
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 260:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:813
		{
			pop(yylex)
		}
	case 261:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:816
		{
			pop(yylex)
		}
	case 262:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:824
		{
			pop(yylex)
		}
	case 263:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:829
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 273:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:851
		{
			pop(yylex)
		}
	case 274:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:856
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 283:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:876
		{
			pop(yylex)
		}
	case 284:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:881
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 288:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:896
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 289:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:901
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 290:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:908
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 303:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:929
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 304:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:936
		{
			pop(yylex)
		}
	case 305:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:939
		{
			pop(yylex)
		}
	case 314:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:955
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 315:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:960
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 316:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:967
		{
			pop(yylex)
		}
	case 317:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:972
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 334:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1002
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 335:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1009
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 336:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1012
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 337:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1017
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 338:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1027
		{
			yyVAL.boolean = true
		}
	case 339:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1028
		{
			yyVAL.boolean = false
		}
	case 340:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1031
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 341:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1041
		{
			pop(yylex)
		}
	case 342:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1046
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 343:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1053
		{
			pop(yylex)
		}
	case 344:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1056
		{
			pop(yylex)
		}
	case 345:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1061
		{
			if push(yylex, meta.NewEnum(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 352:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1077
		{
			if set(yylex, meta.SetEnumValue(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 353:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1084
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 354:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1091
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 355:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1098
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 356:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1105
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 357:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1112
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 358:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1119
		{
			if set(yylex, meta.SetUnits(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 361:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1130
		{

		}
	}
	goto yystack /* stack new state and value */
}
