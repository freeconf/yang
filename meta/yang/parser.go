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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:949

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 784

var yyAct = [...]int{

	204, 317, 202, 12, 218, 217, 12, 379, 366, 141,
	275, 45, 308, 201, 220, 280, 207, 272, 241, 256,
	44, 43, 249, 212, 206, 42, 224, 321, 41, 164,
	40, 184, 296, 152, 39, 178, 38, 211, 169, 136,
	159, 37, 19, 134, 15, 146, 143, 267, 142, 404,
	203, 403, 309, 11, 376, 163, 11, 408, 209, 19,
	79, 238, 176, 3, 20, 361, 368, 322, 323, 368,
	83, 84, 85, 86, 19, 176, 90, 406, 411, 405,
	19, 20, 143, 138, 142, 166, 167, 277, 150, 165,
	155, 370, 4, 365, 134, 134, 20, 134, 319, 161,
	134, 172, 20, 180, 186, 277, 214, 214, 189, 188,
	226, 233, 243, 251, 258, 173, 274, 293, 282, 230,
	191, 165, 219, 219, 170, 269, 205, 165, 190, 130,
	216, 216, 147, 137, 268, 266, 138, 134, 149, 265,
	154, 246, 264, 290, 263, 318, 150, 359, 262, 160,
	261, 171, 155, 179, 185, 260, 213, 213, 236, 161,
	225, 232, 242, 250, 257, 221, 273, 245, 281, 172,
	334, 156, 333, 19, 91, 286, 144, 140, 180, 391,
	31, 19, 153, 173, 186, 291, 137, 375, 189, 188,
	147, 288, 170, 134, 374, 20, 149, 299, 320, 295,
	191, 139, 154, 20, 327, 133, 373, 304, 190, 160,
	97, 214, 307, 313, 306, 315, 372, 134, 300, 171,
	165, 371, 134, 134, 226, 162, 324, 219, 179, 117,
	187, 116, 215, 215, 185, 216, 227, 235, 244, 252,
	259, 243, 276, 115, 283, 114, 316, 134, 329, 251,
	332, 134, 325, 109, 369, 108, 258, 208, 89, 341,
	88, 213, 82, 33, 81, 355, 19, 269, 354, 107,
	246, 345, 274, 409, 225, 347, 268, 266, 342, 343,
	282, 265, 337, 335, 264, 162, 263, 33, 20, 349,
	262, 242, 261, 238, 19, 352, 245, 260, 402, 250,
	353, 19, 148, 72, 23, 399, 257, 393, 67, 66,
	187, 75, 390, 165, 68, 389, 20, 69, 363, 358,
	73, 238, 273, 20, 74, 70, 71, 356, 303, 285,
	281, 367, 19, 19, 176, 381, 175, 215, 289, 387,
	386, 165, 19, 148, 19, 23, 176, 351, 175, 383,
	227, 385, 388, 348, 20, 20, 346, 344, 340, 384,
	181, 194, 336, 392, 20, 331, 20, 244, 292, 330,
	253, 395, 19, 153, 328, 252, 326, 314, 294, 381,
	338, 302, 259, 387, 386, 380, 400, 301, 396, 113,
	112, 327, 111, 383, 20, 385, 110, 106, 276, 105,
	104, 103, 102, 384, 319, 100, 283, 98, 95, 19,
	410, 176, 94, 175, 394, 87, 287, 197, 397, 80,
	310, 298, 298, 80, 284, 398, 6, 350, 339, 380,
	198, 20, 30, 195, 196, 181, 55, 311, 132, 131,
	129, 194, 128, 127, 126, 19, 125, 124, 123, 122,
	121, 120, 199, 197, 72, 193, 165, 119, 118, 67,
	66, 382, 75, 64, 65, 68, 198, 20, 69, 195,
	196, 73, 99, 93, 92, 74, 70, 71, 28, 27,
	210, 54, 56, 192, 183, 182, 76, 52, 177, 77,
	101, 78, 165, 51, 248, 247, 60, 7, 19, 24,
	240, 23, 239, 298, 298, 382, 253, 72, 59, 158,
	157, 34, 67, 66, 49, 75, 64, 65, 68, 378,
	20, 69, 377, 229, 73, 228, 25, 26, 74, 70,
	71, 223, 222, 57, 255, 17, 18, 254, 61, 76,
	200, 401, 77, 53, 78, 19, 364, 362, 19, 360,
	305, 174, 175, 197, 72, 357, 197, 168, 50, 67,
	66, 237, 75, 64, 65, 68, 198, 20, 69, 198,
	20, 73, 195, 196, 234, 74, 70, 71, 231, 19,
	58, 271, 270, 62, 279, 278, 76, 197, 72, 77,
	63, 78, 165, 67, 66, 165, 75, 64, 65, 68,
	198, 20, 69, 48, 47, 73, 46, 36, 35, 74,
	70, 71, 297, 19, 32, 151, 22, 7, 19, 24,
	76, 23, 72, 77, 145, 78, 165, 67, 66, 21,
	75, 64, 65, 68, 135, 20, 69, 16, 14, 73,
	20, 13, 10, 74, 70, 71, 25, 26, 312, 9,
	8, 29, 19, 5, 76, 17, 18, 77, 2, 78,
	165, 72, 1, 407, 0, 0, 67, 66, 0, 75,
	64, 65, 68, 0, 20, 69, 0, 0, 73, 0,
	0, 0, 74, 70, 71, 19, 0, 0, 0, 0,
	0, 0, 0, 76, 72, 0, 77, 0, 78, 67,
	66, 0, 75, 64, 65, 68, 0, 20, 69, 0,
	0, 73, 0, 0, 96, 74, 70, 71, 0, 0,
	0, 0, 0, 0, 0, 0, 76, 72, 0, 77,
	0, 78, 67, 66, 49, 75, 64, 65, 68, 0,
	0, 69, 0, 0, 73, 0, 72, 0, 74, 70,
	71, 67, 66, 0, 75, 64, 65, 68, 0, 76,
	69, 0, 77, 73, 78, 0, 0, 74, 70, 71,
	0, 0, 0, 0, 0, 0, 0, 0, 76, 0,
	0, 77, 0, 78,
}
var yyPact = [...]int{

	37, -1000, 605, 475, 474, 485, -1000, 418, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 254, 418, 418, 418,
	418, 407, 250, 418, 163, 470, 469, 404, 400, 705,
	-1000, -1000, -1000, -1000, 399, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 468,
	397, 394, 393, 392, 391, 389, 259, 245, 388, 384,
	382, 381, 235, 221, 454, 453, 447, 446, 445, 444,
	443, 442, 440, 439, 438, 436, 418, 435, 434, 195,
	-1000, -1000, 46, 191, 167, 38, 166, 288, -1000, 168,
	161, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 61, -1000,
	331, -1000, 672, 432, 566, 396, 396, -1000, -1000, 67,
	253, 61, 600, 281, -1000, 29, -1000, 160, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-13, -1000, -1000, -1000, 419, 320, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 409, -1000, 329, -1000, -1000, 132, -1000,
	-1000, 359, -1000, 106, -1000, -1000, -1000, 369, 61, -1000,
	-1000, -1000, -1000, 724, 724, 418, 379, 373, 319, -1000,
	-1000, -1000, -1000, -1000, 204, 414, 433, 639, -1000, -1000,
	-1000, -1000, 368, 432, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 418, -1000, 92, 398, 19, 19, 418,
	367, 566, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 365,
	396, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 360, 356, 67, -1000, -1000, -1000, -1000, -1000, 162,
	418, 353, -1000, -1000, 21, -1000, -1000, 372, 424, 349,
	61, -1000, -1000, -1000, -1000, 724, 724, 348, 600, -1000,
	-1000, -1000, -1000, -1000, 347, 281, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	344, 29, -1000, -1000, -1000, -1000, -1000, 423, 338, 160,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 414, -1000, -1000,
	258, -1000, -1000, 255, -1000, -1000, 318, 724, -1000, 310,
	137, -1000, -1000, -1000, -1000, -1000, -1000, 47, 244, -13,
	-1000, -1000, -1000, -1000, -1000, -1000, 81, 211, 206, -1000,
	196, 184, -1000, -1000, 177, 44, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 535, -13, -1000, -1000, 566, -1000,
	-1000, -1000, 306, 303, -1000, -1000, -1000, -1000, -1000, -1000,
	169, -1000, -1000, 74, -1000, -1000, -1000, -1000, -1000, -1000,
	298, 418, 50, -1000, 60, 418, -1000, -1000, 421, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 296, 535, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 532, -1000,
	-1000, -1000, 289, -1000, 41, -1000, -1000, 39, 69, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 10, 264, 398, -1000,
	68, -1000,
}
var yyPgo = [...]int{

	0, 663, 27, 1, 12, 52, 662, 658, 653, 651,
	426, 650, 649, 642, 50, 0, 641, 638, 44, 637,
	634, 39, 629, 624, 45, 616, 615, 33, 180, 614,
	257, 32, 612, 608, 607, 41, 36, 34, 30, 28,
	25, 21, 20, 11, 606, 604, 603, 590, 585, 584,
	15, 126, 583, 582, 581, 17, 10, 580, 578, 574,
	47, 561, 13, 558, 557, 38, 23, 14, 551, 550,
	549, 547, 546, 543, 540, 2, 24, 16, 538, 537,
	534, 19, 533, 532, 531, 26, 525, 523, 7, 5,
	4, 522, 519, 511, 510, 509, 40, 55, 29, 508,
	502, 500, 18, 496, 495, 494, 22, 493, 490, 488,
	35, 487, 485, 484, 31, 483, 482, 481, 58, 480,
	37, 436, 8, 9,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 32,
	32, 46, 46, 47, 48, 48, 49, 49, 50, 50,
	50, 51, 45, 45, 52, 53, 53, 54, 54, 55,
	55, 55, 55, 56, 41, 58, 58, 58, 58, 58,
	57, 59, 59, 60, 61, 33, 63, 64, 64, 65,
	65, 65, 65, 4, 4, 67, 66, 68, 69, 69,
	70, 70, 70, 70, 70, 72, 72, 36, 73, 74,
	74, 62, 62, 75, 75, 75, 75, 75, 75, 78,
	44, 79, 79, 80, 80, 81, 81, 81, 81, 81,
	81, 81, 81, 81, 81, 81, 81, 81, 82, 40,
	40, 83, 83, 84, 84, 85, 85, 85, 85, 87,
	88, 88, 88, 88, 88, 88, 88, 88, 86, 86,
	91, 92, 92, 29, 93, 94, 94, 95, 95, 96,
	96, 96, 96, 96, 97, 98, 42, 99, 100, 100,
	101, 101, 102, 102, 102, 102, 102, 43, 103, 104,
	104, 105, 105, 106, 106, 106, 106, 34, 108, 107,
	109, 109, 110, 110, 110, 35, 111, 112, 113, 113,
	89, 89, 90, 114, 114, 114, 114, 114, 114, 114,
	114, 114, 114, 115, 39, 116, 116, 37, 117, 118,
	119, 119, 120, 120, 120, 120, 120, 120, 120, 120,
	120, 77, 5, 5, 3, 2, 2, 76, 38, 121,
	71, 71, 122, 122, 1, 14, 15, 12, 13, 123,
	123,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 3, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 3, 4, 0, 1, 1, 1, 1,
	2, 1, 2, 4, 2, 4, 2, 1, 2, 1,
	1, 1, 1, 1, 1, 3, 2, 2, 1, 3,
	3, 1, 1, 1, 3, 1, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 2,
	4, 0, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	4, 0, 1, 1, 2, 1, 1, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	1, 1, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 3, 3, 2, 2, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 3, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 2, 3, 2,
	1, 2, 1, 1, 1, 4, 2, 1, 1, 2,
	3, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 2, 2, 2, 4, 2, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 1, 1, 1, 3, 4, 2,
	1, 2, 3, 5, 3, 3, 3, 3, 3, 1,
	5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -93, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, -46, 29,
	-63, -107, -111, -73, -117, -121, -116, -82, -57, -99,
	-103, -78, -52, -47, 31, 32, 28, 27, 33, 36,
	44, 45, 22, 39, 43, 30, 54, 57, 59, -5,
	5, 10, 8, -5, -5, -5, -5, 8, 10, 8,
	-5, 11, 4, 4, 8, 8, 9, -28, 8, 4,
	8, -108, 8, 8, 8, 8, 8, 10, 10, 8,
	8, 8, 8, 8, 10, 8, 10, 8, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	-5, 4, 4, 10, 56, -20, -21, -14, -15, 10,
	10, -123, 10, 8, 10, -23, -24, -18, 14, -14,
	-15, -26, -27, 14, -14, -15, 10, -94, -95, -96,
	-14, -15, -51, -97, -98, 60, 24, 25, -64, -65,
	-66, -14, -15, -67, -68, 17, 15, -109, -110, -14,
	-15, -30, -112, -113, -114, -14, -15, -51, -89, -90,
	-76, -77, -115, 23, -30, 37, 38, 21, 34, 20,
	-74, -62, -75, -14, -15, -51, -76, -77, -30, -118,
	-119, -120, -66, -14, -15, -51, -76, -89, -90, -77,
	-67, -118, -83, -84, -85, -14, -15, -51, -86, -87,
	52, -58, -14, -15, -59, -51, -60, -61, 40, -100,
	-101, -102, -14, -15, -51, -97, -98, -104, -105, -106,
	-14, -15, -51, -30, -79, -80, -81, -14, -15, -51,
	-35, -36, -37, -38, -39, -40, -41, -60, -42, -43,
	-53, -54, -55, -14, -15, -56, -51, 58, -48, -49,
	-50, -14, -15, -51, 5, 9, -21, 7, -24, 9,
	11, -27, 9, 11, 9, -96, -31, -32, -30, -31,
	-5, 8, 8, 9, -65, -69, 10, 8, -4, -5,
	6, 4, 9, -110, 9, -114, -5, -3, 53, 6,
	-3, -2, 48, 49, -2, -5, 9, -75, 9, -120,
	9, 9, -85, 10, 8, -5, 9, -60, 8, 4,
	9, -102, -31, -31, 9, -106, 9, -81, 9, -55,
	4, 9, -50, -4, 10, 10, 9, -30, 9, 10,
	-70, 18, -71, -56, -72, 46, -122, -66, 19, 10,
	10, 10, 10, 10, 10, 10, 10, -91, -92, -88,
	-14, -15, -51, -67, -76, -77, -89, -90, -62, 9,
	9, 10, -123, 9, -5, -122, -66, -5, 4, 9,
	-88, 9, 9, 10, 10, 10, 8, -1, 47, 9,
	-3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	252, 16, 0, 0, 0, 0, 0, 0, 37, 0,
	0, 15, 22, 31, 2, 3, 1, 42, 175, 174,
	0, 207, 0, 0, 119, 0, 0, 234, 149, 151,
	85, 188, 199, 131, 72, 75, 61, 64, 96, 209,
	216, 118, 238, 259, 235, 236, 148, 90, 187, 198,
	129, 74, 63, 6, 0, 0, 18, 20, 21, 267,
	268, 265, 269, 0, 266, 0, 23, 26, 0, 28,
	29, 0, 32, 0, 35, 36, 25, 0, 176, 177,
	179, 180, 181, 43, 43, 0, 0, 0, 0, 97,
	99, 100, 101, 102, 0, 0, 0, 0, 210, 212,
	213, 214, 0, 217, 218, 223, 224, 225, 226, 227,
	228, 229, 230, 0, 232, 0, 0, 0, 0, 0,
	0, 120, 121, 123, 124, 125, 126, 127, 128, 0,
	239, 240, 242, 243, 244, 245, 246, 247, 248, 249,
	250, 0, 0, 152, 153, 155, 156, 157, 158, 0,
	0, 0, 86, 87, 88, 89, 91, 0, 0, 0,
	189, 190, 192, 193, 194, 43, 43, 0, 200, 201,
	203, 204, 205, 206, 0, 132, 133, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	0, 76, 77, 79, 80, 81, 82, 0, 0, 65,
	66, 68, 69, 70, 253, 17, 19, 0, 24, 30,
	0, 33, 38, 0, 173, 178, 0, 44, 59, 0,
	0, 184, 185, 95, 98, 106, 108, 0, 0, 103,
	104, 107, 208, 211, 215, 219, 0, 0, 0, 254,
	0, 0, 255, 256, 0, 0, 117, 122, 237, 241,
	258, 150, 154, 168, 0, 159, 84, 92, 0, 94,
	186, 191, 0, 0, 197, 202, 130, 134, 73, 78,
	0, 62, 67, 0, 27, 34, 182, 60, 183, 71,
	0, 0, 111, 112, 113, 0, 260, 115, 0, 105,
	231, 220, 221, 222, 257, 251, 233, 0, 170, 171,
	160, 161, 162, 163, 164, 165, 166, 167, 0, 195,
	196, 83, 0, 109, 0, 261, 116, 0, 0, 169,
	172, 93, 270, 110, 114, 262, 0, 0, 0, 263,
	0, 264,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60,
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
		//line parser.y:142
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
		//line parser.y:150
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
		//line parser.y:164
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:179
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:186
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:189
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:202
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:213
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:226
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:231
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:247
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:250
		{
			pop(yylex)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:287
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:290
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:295
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:314
		{
			if set(yylex, meta.NewIfFeature(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:321
		{
			pop(yylex)
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:324
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:329
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:349
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:359
		{
			pop(yylex)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:371
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:383
		{
			pop(yylex)
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:388
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:398
		{
			pop(yylex)
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:403
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:420
		{
			yyVAL.token = yyDollar[1].token
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:421
		{
			yyVAL.token = yyDollar[1].token
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:424
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:434
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:441
		{
			pop(yylex)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:444
		{
			pop(yylex)
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:449
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:457
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:471
		{
			pop(yylex)
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:476
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:499
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:506
		{
			pop(yylex)
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:533
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:540
		{
			pop(yylex)
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:543
		{
			pop(yylex)
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:564
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:587
		{
			pop(yylex)
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:590
		{
			pop(yylex)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:602
		{
			pop(yylex)
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:607
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:624
		{
			pop(yylex)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:627
		{
			pop(yylex)
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:632
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:639
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:649
		{
			pop(yylex)
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:654
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:671
		{
			pop(yylex)
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:674
		{
			pop(yylex)
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:682
		{
			pop(yylex)
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:687
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:709
		{
			pop(yylex)
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:719
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:737
		{
			pop(yylex)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:742
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:757
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:762
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:769
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:788
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 234:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:795
		{
			pop(yylex)
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:800
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:805
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 237:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:815
		{
			pop(yylex)
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:820
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 251:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:847
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:854
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 253:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:857
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:862
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:872
		{
			yyVAL.boolean = true
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:873
		{
			yyVAL.boolean = false
		}
	case 257:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:876
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 258:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:886
		{
			pop(yylex)
		}
	case 259:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:891
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 262:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:902
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 263:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:907
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 264:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:914
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 265:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:919
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 266:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:926
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 267:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:933
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 268:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:940
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
