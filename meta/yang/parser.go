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
const kywd_yang_version = 57405

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
	"kywd_yang_version",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:989

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 909

var yyAct = [...]int{

	214, 344, 212, 12, 408, 394, 12, 233, 148, 211,
	301, 223, 306, 334, 298, 281, 274, 266, 218, 48,
	47, 46, 222, 322, 231, 217, 230, 245, 44, 43,
	42, 192, 348, 141, 45, 41, 40, 20, 186, 39,
	293, 153, 159, 177, 143, 167, 49, 434, 184, 220,
	346, 389, 396, 20, 3, 184, 172, 183, 20, 21,
	213, 207, 438, 11, 263, 263, 11, 33, 20, 174,
	175, 396, 171, 15, 208, 21, 433, 205, 206, 393,
	21, 20, 184, 4, 173, 204, 215, 145, 20, 160,
	21, 303, 157, 141, 162, 319, 150, 345, 149, 102,
	173, 204, 61, 21, 169, 173, 180, 316, 188, 194,
	21, 225, 225, 181, 237, 173, 247, 178, 257, 268,
	276, 283, 141, 300, 405, 308, 303, 200, 173, 232,
	232, 96, 242, 198, 199, 197, 229, 229, 441, 241,
	295, 294, 292, 145, 141, 227, 227, 144, 239, 291,
	290, 289, 156, 157, 161, 421, 288, 287, 261, 162,
	286, 234, 251, 404, 168, 154, 179, 169, 187, 193,
	141, 224, 224, 403, 236, 271, 246, 180, 256, 267,
	275, 282, 402, 299, 181, 307, 188, 312, 178, 399,
	170, 270, 194, 398, 314, 195, 325, 226, 226, 20,
	238, 317, 248, 144, 259, 269, 277, 284, 347, 302,
	200, 309, 321, 156, 354, 401, 198, 199, 197, 161,
	330, 21, 225, 341, 339, 387, 154, 168, 219, 20,
	349, 350, 400, 254, 35, 141, 397, 179, 163, 141,
	232, 351, 335, 151, 356, 247, 187, 229, 383, 382,
	83, 21, 193, 170, 318, 439, 227, 147, 20, 160,
	35, 87, 88, 89, 90, 164, 268, 94, 253, 146,
	80, 141, 360, 140, 276, 432, 173, 204, 195, 141,
	21, 283, 224, 369, 141, 20, 155, 429, 24, 141,
	373, 251, 423, 420, 370, 371, 375, 419, 300, 365,
	295, 294, 292, 141, 116, 246, 308, 21, 226, 291,
	290, 289, 377, 386, 384, 141, 288, 287, 380, 141,
	286, 366, 271, 137, 379, 20, 267, 381, 436, 183,
	435, 248, 216, 207, 275, 315, 189, 203, 270, 20,
	155, 282, 24, 376, 391, 395, 208, 21, 278, 205,
	206, 329, 269, 374, 150, 20, 149, 184, 299, 183,
	277, 21, 362, 410, 361, 328, 307, 284, 372, 20,
	412, 184, 173, 183, 61, 368, 418, 21, 364, 311,
	333, 414, 332, 20, 302, 359, 358, 417, 413, 416,
	422, 21, 309, 124, 327, 123, 425, 415, 357, 355,
	324, 324, 353, 340, 426, 21, 320, 120, 410, 20,
	119, 122, 430, 121, 189, 412, 326, 207, 118, 117,
	203, 354, 115, 409, 114, 113, 414, 112, 111, 110,
	208, 21, 417, 413, 416, 93, 86, 92, 85, 109,
	440, 196, 415, 228, 228, 342, 240, 343, 249, 411,
	260, 108, 352, 285, 107, 313, 173, 204, 61, 7,
	20, 26, 105, 24, 103, 100, 99, 91, 409, 76,
	84, 336, 346, 84, 71, 70, 52, 79, 68, 69,
	72, 6, 21, 73, 310, 95, 77, 32, 27, 28,
	78, 74, 75, 428, 411, 378, 363, 18, 19, 324,
	324, 80, 278, 58, 81, 367, 82, 337, 139, 61,
	25, 138, 20, 136, 135, 134, 133, 132, 131, 209,
	207, 76, 202, 130, 196, 129, 71, 70, 128, 79,
	68, 69, 72, 208, 21, 73, 205, 206, 77, 127,
	126, 125, 78, 74, 75, 104, 98, 97, 30, 29,
	221, 57, 385, 80, 228, 235, 81, 59, 82, 173,
	204, 61, 201, 191, 190, 55, 431, 185, 106, 54,
	20, 273, 272, 64, 265, 264, 63, 249, 207, 76,
	166, 165, 36, 407, 71, 70, 406, 79, 68, 69,
	72, 208, 21, 73, 252, 250, 77, 244, 243, 60,
	78, 74, 75, 280, 279, 65, 210, 56, 392, 390,
	388, 80, 331, 285, 81, 182, 82, 173, 204, 61,
	20, 176, 53, 262, 258, 255, 62, 297, 207, 76,
	296, 66, 424, 305, 71, 70, 427, 79, 68, 69,
	72, 208, 21, 73, 304, 67, 77, 51, 50, 38,
	78, 74, 75, 37, 323, 34, 20, 158, 23, 152,
	22, 80, 142, 17, 81, 76, 82, 173, 204, 61,
	71, 70, 16, 79, 68, 69, 72, 14, 21, 73,
	13, 10, 77, 9, 8, 31, 78, 74, 75, 5,
	2, 338, 1, 437, 0, 20, 0, 80, 0, 0,
	81, 0, 82, 173, 76, 61, 0, 0, 0, 71,
	70, 0, 79, 68, 69, 72, 0, 21, 73, 0,
	0, 77, 0, 0, 0, 78, 74, 75, 20, 0,
	0, 0, 0, 0, 0, 0, 80, 76, 0, 81,
	0, 82, 71, 70, 61, 79, 68, 69, 72, 0,
	21, 73, 0, 0, 77, 0, 0, 0, 78, 74,
	75, 0, 101, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 81, 0, 82, 76, 0, 61, 0, 0,
	71, 70, 52, 79, 68, 69, 72, 0, 0, 73,
	0, 0, 77, 0, 0, 0, 78, 74, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 80, 76, 0,
	81, 0, 82, 71, 70, 61, 79, 68, 69, 72,
	0, 0, 73, 0, 20, 77, 0, 0, 0, 78,
	74, 75, 0, 76, 0, 0, 0, 0, 71, 70,
	80, 79, 0, 81, 72, 82, 21, 73, 61, 0,
	77, 263, 0, 0, 78, 74, 75, 7, 20, 26,
	0, 24, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 173, 204, 0, 0, 0, 0, 0, 0, 0,
	21, 0, 0, 0, 0, 0, 27, 28, 0, 0,
	0, 0, 0, 0, 0, 18, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 25,
}
var yyPact = [...]int{

	28, -1000, 845, 545, 544, 447, -1000, 468, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 428, 468, 468,
	468, 468, 459, 427, 468, 480, 120, 543, 542, 458,
	457, 753, -1000, -1000, -1000, -1000, 456, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 541, 454, 446, 443, 431, 421, 420, 417,
	414, 468, 411, 410, 402, 399, 403, 385, 537, 536,
	535, 524, 521, 519, 514, 513, 512, 511, 510, 509,
	468, 507, 504, 263, -1000, -1000, 186, 259, 247, 88,
	233, 272, -1000, 75, 228, 255, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 45, -1000, 356, -1000, 715, 499, 607,
	40, 40, -1000, 396, -1000, 216, 223, 24, 45, 643,
	811, -1000, 68, -1000, 55, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -23, -1000, -1000,
	-1000, 479, 370, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	448, -1000, 326, -1000, -1000, 96, -1000, -1000, 245, -1000,
	84, -1000, -1000, -1000, -1000, 397, 45, -1000, -1000, -1000,
	-1000, 786, 786, 468, 386, 357, 342, -1000, -1000, -1000,
	-1000, -1000, 372, 465, 503, 682, -1000, -1000, -1000, -1000,
	394, 499, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 468, -1000, 468, 44, 466, 182, 182, 468,
	393, 607, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	390, 40, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 389, 377, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 376, 216, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 354, 468, -1000, 369, -1000, -1000, 25, -1000,
	-1000, -1000, 313, 501, 366, 45, -1000, -1000, -1000, -1000,
	786, 786, 359, 643, -1000, -1000, -1000, -1000, -1000, 344,
	811, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 334, 68, -1000, -1000,
	-1000, -1000, -1000, 491, 315, 55, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 465, -1000, -1000, 239, -1000, -1000, 238,
	-1000, -1000, 305, 786, -1000, 304, 215, -1000, -1000, -1000,
	-1000, -1000, -1000, 33, 226, -23, -1000, -1000, -1000, -1000,
	-1000, -1000, 183, 179, 222, 205, -1000, 172, 163, -1000,
	-1000, 153, 114, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 312, -23, -1000, -1000, 607, -1000, -1000, -1000,
	288, 284, -1000, -1000, -1000, -1000, -1000, -1000, 145, -1000,
	-1000, 346, -1000, -1000, -1000, -1000, -1000, -1000, 283, 468,
	52, -1000, 67, 468, -1000, -1000, 489, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 278, 312, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 557, -1000,
	-1000, -1000, 266, -1000, 66, -1000, -1000, 37, 320, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 15, 246, 466, -1000,
	128, -1000,
}
var yyPgo = [...]int{

	0, 693, 32, 1, 13, 242, 692, 690, 689, 685,
	481, 684, 683, 681, 60, 0, 680, 677, 73, 672,
	663, 662, 44, 660, 659, 41, 658, 657, 42, 67,
	655, 228, 23, 654, 653, 649, 39, 36, 35, 30,
	29, 28, 34, 21, 20, 19, 46, 648, 647, 645,
	644, 633, 12, 86, 332, 631, 630, 627, 14, 10,
	626, 625, 624, 40, 623, 9, 622, 621, 43, 11,
	7, 615, 612, 610, 609, 608, 607, 606, 2, 25,
	18, 605, 604, 603, 15, 599, 598, 597, 27, 595,
	594, 4, 26, 24, 586, 583, 582, 581, 580, 45,
	72, 56, 576, 575, 574, 17, 573, 572, 571, 16,
	569, 568, 567, 38, 565, 564, 563, 31, 562, 557,
	555, 551, 49, 550, 22, 503, 5, 8,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 20, 11, 11, 21,
	21, 22, 22, 23, 24, 24, 18, 25, 25, 25,
	25, 16, 26, 27, 27, 28, 28, 28, 17, 17,
	29, 29, 9, 9, 32, 32, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 31, 31, 31, 31,
	31, 33, 33, 48, 48, 49, 50, 50, 51, 51,
	52, 52, 52, 42, 53, 54, 47, 47, 55, 56,
	56, 57, 57, 58, 58, 58, 58, 59, 43, 61,
	61, 61, 61, 61, 61, 60, 62, 62, 63, 64,
	34, 66, 67, 67, 68, 68, 68, 68, 4, 4,
	70, 69, 71, 72, 72, 73, 73, 73, 73, 73,
	75, 75, 37, 76, 77, 77, 65, 65, 78, 78,
	78, 78, 78, 78, 78, 81, 46, 82, 82, 83,
	83, 84, 84, 84, 84, 84, 84, 84, 84, 84,
	84, 84, 84, 84, 84, 85, 41, 41, 86, 86,
	87, 87, 88, 88, 88, 88, 88, 88, 90, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 89, 89,
	94, 95, 95, 30, 96, 97, 97, 98, 98, 99,
	99, 99, 99, 99, 100, 101, 44, 102, 103, 103,
	104, 104, 105, 105, 105, 105, 105, 45, 106, 107,
	107, 108, 108, 109, 109, 109, 109, 35, 111, 110,
	112, 112, 113, 113, 113, 36, 114, 115, 116, 116,
	92, 92, 93, 117, 117, 117, 117, 117, 117, 117,
	117, 117, 117, 117, 118, 40, 40, 120, 120, 120,
	120, 120, 120, 120, 119, 119, 38, 121, 122, 123,
	123, 124, 124, 124, 124, 124, 124, 124, 124, 124,
	124, 124, 80, 5, 5, 3, 2, 2, 79, 39,
	125, 74, 74, 126, 126, 1, 14, 15, 12, 13,
	19, 127, 127,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 1,
	2, 1, 1, 2, 1, 2, 3, 1, 3, 1,
	1, 4, 2, 1, 2, 3, 1, 1, 2, 4,
	1, 1, 1, 2, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 3, 3, 3, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 3, 4, 0,
	1, 1, 1, 1, 1, 2, 1, 2, 4, 2,
	4, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	3, 2, 2, 1, 3, 3, 1, 1, 1, 3,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 0, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 0, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	1, 1, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 3, 3, 2, 2, 4, 2, 0, 1,
	1, 2, 1, 1, 1, 3, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 2, 3, 2,
	1, 2, 1, 1, 1, 4, 2, 1, 1, 2,
	3, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 2, 4, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 4, 2, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 1, 1, 3, 4,
	2, 1, 2, 3, 5, 3, 3, 3, 3, 3,
	3, 1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, -20, 50, 51,
	13, 35, -23, -26, 16, 63, 14, 41, 42, 4,
	4, -9, -10, -29, -30, -31, -96, -34, -35, -36,
	-37, -38, -39, -40, -41, -42, -43, -44, -45, -46,
	-47, -48, 29, -66, -110, -114, -76, -121, -125, -119,
	-85, 62, -60, -102, -106, -81, -55, -49, 31, 32,
	28, 27, 33, 36, 44, 45, 22, 39, 43, 30,
	54, 57, 59, -5, 5, 10, 8, -5, -5, -5,
	-5, 8, 10, 8, -5, 5, 11, 4, 4, 8,
	8, 9, -29, 8, 4, 8, -111, 8, 8, 8,
	8, 8, 10, 8, 10, 8, -5, 8, 8, 8,
	8, 10, 8, 10, 8, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, -5, 4, 4,
	10, 56, -21, -22, -14, -15, 10, 10, -127, 10,
	8, 10, -24, -25, -18, 14, -14, -15, -27, -28,
	14, -14, -15, 10, 10, -97, -98, -99, -14, -15,
	-53, -100, -101, 60, 24, 25, -67, -68, -69, -14,
	-15, -70, -71, 17, 15, -112, -113, -14, -15, -31,
	-115, -116, -117, -14, -15, -53, -54, -92, -93, -79,
	-80, -118, 23, -31, 61, 37, 38, 21, 34, 20,
	-77, -65, -78, -14, -15, -53, -54, -79, -80, -31,
	-122, -123, -124, -69, -14, -15, -53, -42, -54, -79,
	-92, -93, -80, -70, -122, -120, -14, -15, -53, -42,
	-54, -79, -80, -86, -87, -88, -14, -15, -53, -54,
	-89, -46, -90, 52, 10, -61, -14, -15, -62, -53,
	-54, -63, -64, 40, -103, -104, -105, -14, -15, -53,
	-100, -101, -107, -108, -109, -14, -15, -53, -31, -82,
	-83, -84, -14, -15, -53, -54, -36, -37, -38, -39,
	-40, -41, -43, -63, -44, -45, -56, -57, -58, -14,
	-15, -59, -53, 58, -50, -51, -52, -14, -15, -53,
	5, 9, -22, 7, -25, 9, 11, -28, 9, 11,
	9, -99, -32, -33, -31, -32, -5, 8, 8, 9,
	-68, -72, 10, 8, -4, -5, 6, 4, 9, -113,
	9, -117, -5, -5, -3, 53, 6, -3, -2, 48,
	49, -2, -5, 9, -78, 9, -124, 9, 9, 9,
	-88, 10, 8, -5, 9, -63, 8, 4, 9, -105,
	-32, -32, 9, -109, 9, -84, 9, -58, 4, 9,
	-52, -4, 10, 10, 9, -31, 9, 10, -73, 18,
	-74, -59, -75, 46, -126, -69, 19, 10, 10, 10,
	10, 10, 10, 10, 10, 10, -94, -95, -91, -14,
	-15, -53, -70, -79, -80, -42, -92, -93, -65, 9,
	9, 10, -127, 9, -5, -126, -69, -5, 4, 9,
	-91, 9, 9, 10, 10, 10, 8, -1, 47, 9,
	-3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 5, 42, 40, 41, 0, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 273, 17, 0, 0, 0, 0,
	0, 0, 38, 0, 0, 0, 16, 23, 32, 2,
	3, 1, 43, 185, 184, 0, 217, 0, 0, 124,
	0, 0, 245, 0, 156, 158, 0, 89, 198, 209,
	137, 76, 79, 63, 66, 101, 219, 226, 123, 257,
	280, 254, 255, 155, 95, 197, 208, 135, 78, 65,
	6, 0, 0, 19, 21, 22, 288, 289, 286, 291,
	0, 287, 0, 24, 27, 0, 29, 30, 0, 33,
	0, 36, 37, 26, 290, 0, 186, 187, 189, 190,
	191, 44, 44, 0, 0, 0, 0, 102, 104, 105,
	106, 107, 0, 0, 0, 0, 220, 222, 223, 224,
	0, 227, 228, 233, 234, 235, 236, 237, 238, 239,
	240, 241, 0, 243, 0, 0, 0, 0, 0, 0,
	0, 125, 126, 128, 129, 130, 131, 132, 133, 134,
	0, 258, 259, 261, 262, 263, 264, 265, 266, 267,
	268, 269, 270, 271, 0, 0, 247, 248, 249, 250,
	251, 252, 253, 0, 159, 160, 162, 163, 164, 165,
	166, 167, 0, 0, 73, 0, 90, 91, 92, 93,
	94, 96, 0, 0, 0, 199, 200, 202, 203, 204,
	44, 44, 0, 210, 211, 213, 214, 215, 216, 0,
	138, 139, 141, 142, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 0, 80, 81, 83,
	84, 85, 86, 0, 0, 67, 68, 70, 71, 72,
	274, 18, 20, 0, 25, 31, 0, 34, 39, 0,
	183, 188, 0, 45, 61, 0, 0, 194, 195, 100,
	103, 111, 113, 0, 0, 108, 109, 112, 218, 221,
	225, 229, 0, 0, 0, 0, 275, 0, 0, 276,
	277, 0, 0, 122, 127, 256, 260, 279, 246, 157,
	161, 178, 0, 168, 88, 97, 0, 99, 196, 201,
	0, 0, 207, 212, 136, 140, 77, 82, 0, 64,
	69, 0, 28, 35, 192, 62, 193, 74, 0, 0,
	116, 117, 118, 0, 281, 120, 0, 110, 242, 75,
	230, 231, 232, 278, 272, 244, 0, 180, 181, 169,
	170, 171, 172, 173, 174, 175, 176, 177, 0, 205,
	206, 87, 0, 114, 0, 282, 121, 0, 0, 179,
	182, 98, 292, 115, 119, 283, 0, 0, 0, 284,
	0, 285,
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
	62, 63,
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
		//line parser.y:145
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
		//line parser.y:153
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
		//line parser.y:167
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:183
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:190
		{
			pop(yylex)
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:193
		{
			pop(yylex)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:206
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:217
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:230
		{
			pop(yylex)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:235
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:251
		{
			pop(yylex)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:254
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:292
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:295
		{
			pop(yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:300
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:319
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:326
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:333
		{
			if set(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:340
		{
			pop(yylex)
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:343
		{
			pop(yylex)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:348
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:368
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:378
		{
			pop(yylex)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:391
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:403
		{
			pop(yylex)
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:408
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:418
		{
			pop(yylex)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:423
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:440
		{
			yyVAL.token = yyDollar[1].token
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:441
		{
			yyVAL.token = yyDollar[1].token
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:444
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:454
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:461
		{
			pop(yylex)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:464
		{
			pop(yylex)
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:469
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:477
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:491
		{
			pop(yylex)
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:496
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:520
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:527
		{
			pop(yylex)
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:555
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:562
		{
			pop(yylex)
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:565
		{
			pop(yylex)
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:585
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:604
		{
			pop(yylex)
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:607
		{
			pop(yylex)
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:619
		{
			pop(yylex)
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:624
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:641
		{
			pop(yylex)
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:644
		{
			pop(yylex)
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:649
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:656
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:666
		{
			pop(yylex)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:671
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:688
		{
			pop(yylex)
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:691
		{
			pop(yylex)
		}
	case 207:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:699
		{
			pop(yylex)
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:704
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:726
		{
			pop(yylex)
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:736
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:754
		{
			pop(yylex)
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:759
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:774
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:779
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:786
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:806
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:813
		{
			pop(yylex)
		}
	case 246:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:816
		{
			pop(yylex)
		}
	case 254:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:831
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 255:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:836
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 256:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:846
		{
			pop(yylex)
		}
	case 257:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:851
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 272:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:879
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:886
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 274:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:889
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:894
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:904
		{
			yyVAL.boolean = true
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:905
		{
			yyVAL.boolean = false
		}
	case 278:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:908
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 279:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:918
		{
			pop(yylex)
		}
	case 280:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:923
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 283:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:934
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 284:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:939
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 285:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:946
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 286:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:951
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 287:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:958
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 288:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:965
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 289:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:972
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 290:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:979
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
