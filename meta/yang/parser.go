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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:1015

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 916

var yyAct = [...]int{

	217, 347, 215, 12, 234, 416, 12, 399, 405, 214,
	150, 304, 309, 301, 284, 40, 277, 37, 39, 269,
	38, 233, 179, 248, 188, 225, 216, 351, 155, 11,
	194, 161, 11, 122, 169, 145, 22, 236, 143, 36,
	35, 221, 446, 326, 34, 220, 41, 22, 33, 22,
	3, 174, 22, 296, 32, 450, 31, 226, 23, 222,
	209, 173, 18, 266, 223, 18, 266, 445, 22, 23,
	186, 23, 185, 210, 23, 22, 152, 15, 151, 4,
	352, 353, 401, 175, 212, 444, 186, 143, 147, 453,
	23, 218, 306, 159, 175, 164, 175, 23, 171, 175,
	212, 57, 448, 182, 447, 190, 196, 429, 228, 228,
	200, 240, 143, 250, 146, 260, 271, 279, 286, 158,
	303, 163, 311, 143, 170, 230, 230, 199, 242, 181,
	143, 189, 195, 298, 227, 227, 297, 239, 295, 249,
	183, 259, 270, 278, 285, 147, 302, 202, 310, 235,
	235, 201, 245, 232, 232, 159, 244, 294, 293, 254,
	180, 164, 292, 412, 191, 205, 291, 274, 264, 171,
	156, 146, 290, 237, 289, 349, 281, 273, 413, 182,
	316, 158, 403, 318, 391, 411, 257, 163, 190, 172,
	410, 152, 321, 151, 196, 170, 165, 197, 200, 229,
	229, 334, 241, 325, 251, 181, 262, 272, 280, 287,
	350, 305, 341, 312, 189, 199, 183, 358, 329, 338,
	195, 348, 123, 143, 343, 228, 366, 143, 365, 143,
	84, 143, 156, 328, 328, 202, 180, 153, 354, 201,
	409, 143, 230, 88, 89, 90, 91, 191, 250, 95,
	360, 227, 149, 205, 322, 370, 408, 22, 162, 451,
	172, 22, 438, 332, 402, 22, 235, 148, 219, 271,
	232, 364, 176, 177, 249, 387, 346, 279, 345, 23,
	113, 142, 143, 23, 286, 197, 319, 23, 373, 22,
	157, 386, 26, 377, 254, 270, 313, 143, 379, 298,
	137, 303, 297, 278, 295, 331, 166, 22, 175, 311,
	285, 23, 143, 443, 381, 369, 229, 374, 375, 440,
	274, 384, 431, 294, 293, 428, 143, 302, 292, 23,
	273, 427, 291, 328, 328, 310, 281, 390, 290, 251,
	289, 22, 157, 337, 26, 336, 256, 407, 77, 396,
	121, 385, 120, 141, 175, 212, 388, 119, 186, 118,
	272, 393, 401, 23, 383, 380, 315, 418, 280, 22,
	112, 425, 111, 406, 198, 287, 231, 231, 378, 243,
	426, 252, 376, 263, 423, 372, 288, 389, 424, 398,
	368, 23, 305, 417, 363, 400, 430, 110, 330, 109,
	312, 306, 362, 434, 420, 407, 361, 394, 422, 359,
	22, 162, 421, 439, 357, 94, 418, 93, 342, 333,
	425, 441, 22, 87, 186, 86, 185, 344, 324, 358,
	140, 406, 23, 423, 355, 356, 117, 424, 116, 115,
	114, 108, 417, 107, 23, 106, 105, 104, 102, 97,
	92, 317, 452, 420, 349, 435, 85, 422, 419, 85,
	124, 421, 198, 64, 6, 82, 323, 7, 22, 45,
	83, 26, 320, 314, 98, 96, 437, 73, 382, 367,
	371, 339, 68, 67, 48, 76, 65, 66, 69, 139,
	23, 70, 138, 231, 74, 136, 46, 47, 75, 71,
	72, 135, 134, 54, 133, 20, 21, 419, 132, 77,
	131, 22, 78, 186, 79, 185, 252, 57, 27, 209,
	64, 130, 129, 128, 7, 22, 45, 127, 26, 126,
	125, 101, 210, 23, 73, 207, 208, 100, 99, 68,
	67, 48, 76, 65, 66, 69, 81, 23, 70, 80,
	224, 74, 288, 46, 47, 75, 71, 72, 175, 212,
	57, 53, 20, 21, 238, 55, 77, 203, 193, 78,
	192, 79, 51, 64, 57, 27, 187, 103, 22, 50,
	276, 275, 60, 268, 267, 211, 209, 73, 204, 59,
	168, 167, 68, 67, 28, 76, 65, 66, 69, 210,
	23, 70, 207, 208, 74, 415, 414, 255, 75, 71,
	72, 253, 247, 246, 56, 283, 432, 433, 64, 77,
	442, 436, 78, 22, 79, 175, 212, 57, 282, 61,
	213, 209, 73, 52, 397, 395, 392, 68, 67, 335,
	76, 65, 66, 69, 210, 23, 70, 184, 178, 74,
	49, 265, 261, 75, 71, 72, 258, 58, 300, 299,
	64, 62, 404, 206, 77, 22, 308, 78, 307, 79,
	175, 212, 57, 209, 73, 63, 44, 43, 42, 68,
	67, 30, 76, 65, 66, 69, 210, 23, 70, 29,
	327, 74, 160, 25, 154, 75, 71, 72, 24, 144,
	19, 17, 64, 16, 14, 13, 77, 22, 10, 78,
	9, 79, 175, 212, 57, 8, 73, 5, 2, 1,
	449, 68, 67, 0, 76, 65, 66, 69, 0, 23,
	70, 0, 0, 74, 0, 0, 0, 75, 71, 72,
	0, 64, 0, 340, 0, 0, 22, 0, 77, 0,
	0, 78, 0, 79, 175, 73, 57, 0, 0, 0,
	68, 67, 0, 76, 65, 66, 69, 0, 23, 70,
	0, 0, 74, 0, 64, 0, 75, 71, 72, 22,
	0, 0, 0, 0, 0, 0, 0, 77, 73, 0,
	78, 0, 79, 68, 67, 57, 76, 65, 66, 69,
	0, 23, 70, 0, 0, 74, 0, 64, 0, 75,
	71, 72, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 73, 0, 78, 0, 79, 68, 67, 57, 76,
	65, 66, 69, 0, 0, 70, 0, 22, 74, 0,
	0, 0, 75, 71, 72, 0, 73, 0, 0, 0,
	0, 68, 67, 77, 76, 0, 78, 69, 79, 23,
	70, 57, 0, 74, 266, 0, 22, 75, 71, 72,
	185, 0, 0, 0, 209, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 175, 212, 0, 210, 23, 0,
	207, 208, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 175, 0, 57,
}
var yyPact = [...]int{

	25, -1000, 513, 545, 542, 456, -1000, 451, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 415,
	451, 451, 451, 451, 442, 407, 451, 470, 441, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 469, 534, 533, 527, 440,
	439, 438, 437, 435, 433, 389, 362, 451, 432, 431,
	430, 428, 349, 342, 454, 526, 525, 523, 519, 518,
	517, 506, 504, 500, 498, 497, 491, 451, 488, 485,
	422, 345, -1000, -1000, 271, -1000, -1000, 63, 257, 242,
	68, 227, 329, -1000, 398, 186, 296, 249, -1000, -1000,
	-1000, -1000, 56, -1000, 767, 566, 653, 499, 499, -1000,
	40, -1000, 295, 176, 24, 249, 695, 825, -1000, 35,
	-1000, 37, 286, -17, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -17, -1000, -1000,
	-1000, -1000, -1000, 468, 357, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 444, -1000, 277, -1000, -1000, 467, -1000, -1000,
	245, -1000, 461, -1000, -1000, -1000, -1000, 419, 249, -1000,
	-1000, -1000, -1000, 800, 800, 451, 297, 255, 410, -1000,
	-1000, -1000, -1000, -1000, 335, 454, 477, 734, -1000, -1000,
	-1000, -1000, 409, 566, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 451, -1000, 268, 169, 448, 33,
	33, 451, 451, 405, 653, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 400, 499, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 397, 393, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 385, 295, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 218, 451, -1000, 381, -1000,
	-1000, 27, -1000, -1000, -1000, 247, 476, 376, 249, -1000,
	-1000, -1000, -1000, 800, 800, 373, 695, -1000, -1000, -1000,
	-1000, -1000, 369, 825, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 356,
	35, -1000, -1000, -1000, -1000, -1000, 474, 355, 37, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 454, -1000, -1000,
	281, -1000, -1000, 265, -1000, -1000, 347, 800, -1000, 328,
	174, -1000, -1000, -1000, -1000, -1000, -1000, 344, 254, -1000,
	-1000, -1000, -1000, -1000, 172, -1000, 63, 246, 230, -1000,
	180, 175, -1000, -1000, 153, 168, -17, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 854, -17, -1000, -1000,
	653, -1000, -1000, -1000, 322, 316, -1000, -1000, -1000, -1000,
	-1000, -1000, 97, -1000, -1000, 183, -1000, -1000, -1000, -1000,
	-1000, -1000, 313, 451, 451, 64, -1000, 72, 451, -1000,
	-1000, 472, -1000, -1000, 253, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 310, 854, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 611, -1000, -1000, -1000,
	304, -1000, 75, 57, -1000, -1000, 32, 94, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 9, 250,
	448, -1000, 79, -1000,
}
var yyPgo = [...]int{

	0, 720, 27, 1, 33, 222, 719, 718, 717, 464,
	715, 710, 708, 26, 0, 705, 704, 77, 703, 701,
	59, 700, 699, 35, 698, 694, 28, 693, 692, 31,
	43, 690, 689, 681, 56, 54, 48, 44, 40, 39,
	17, 20, 18, 15, 46, 678, 677, 676, 675, 668,
	666, 12, 91, 663, 268, 662, 8, 661, 659, 658,
	13, 11, 657, 656, 652, 53, 651, 9, 650, 648,
	22, 57, 37, 647, 639, 636, 635, 634, 633, 630,
	2, 45, 41, 629, 628, 615, 14, 614, 613, 612,
	23, 611, 607, 5, 21, 4, 606, 605, 594, 591,
	590, 34, 61, 51, 589, 584, 583, 19, 582, 581,
	580, 16, 579, 577, 576, 24, 572, 570, 568, 30,
	567, 565, 564, 561, 64, 550, 25, 503, 7, 10,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 21, 10,
	10, 22, 22, 23, 23, 24, 25, 25, 17, 26,
	26, 26, 26, 15, 27, 28, 28, 29, 29, 29,
	16, 16, 30, 30, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	31, 31, 46, 46, 48, 49, 49, 50, 50, 51,
	51, 51, 40, 52, 53, 54, 54, 55, 55, 56,
	56, 45, 45, 57, 58, 58, 59, 59, 60, 60,
	60, 60, 61, 41, 63, 63, 63, 63, 63, 63,
	62, 64, 64, 65, 66, 32, 68, 69, 69, 70,
	70, 70, 70, 4, 4, 72, 71, 73, 74, 74,
	75, 75, 75, 75, 75, 75, 77, 77, 35, 78,
	79, 79, 67, 67, 80, 80, 80, 80, 80, 80,
	80, 83, 44, 84, 84, 85, 85, 86, 86, 86,
	86, 86, 86, 86, 86, 86, 86, 86, 86, 86,
	86, 87, 39, 39, 88, 88, 89, 89, 90, 90,
	90, 90, 90, 90, 92, 93, 93, 93, 93, 93,
	93, 93, 93, 93, 91, 91, 96, 97, 97, 19,
	98, 99, 99, 100, 100, 101, 101, 101, 101, 101,
	102, 103, 42, 104, 105, 105, 106, 106, 107, 107,
	107, 107, 107, 43, 108, 109, 109, 110, 110, 111,
	111, 111, 111, 33, 113, 112, 114, 114, 115, 115,
	115, 34, 116, 117, 118, 118, 94, 94, 95, 119,
	119, 119, 119, 119, 119, 119, 119, 119, 119, 119,
	120, 38, 38, 122, 122, 122, 122, 122, 122, 122,
	121, 121, 36, 123, 124, 125, 125, 126, 126, 126,
	126, 126, 126, 126, 126, 126, 126, 126, 82, 5,
	5, 3, 2, 2, 81, 37, 127, 76, 76, 128,
	128, 1, 13, 14, 11, 12, 18, 129, 129, 47,
}
var yyR2 = [...]int{

	0, 3, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	4, 1, 2, 1, 1, 2, 1, 2, 3, 1,
	3, 1, 1, 4, 2, 1, 2, 3, 1, 1,
	2, 4, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 3, 3, 2, 2, 4, 1, 2, 1,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 3, 4, 0, 1, 1, 1, 1, 1,
	2, 1, 2, 4, 2, 4, 2, 1, 2, 1,
	1, 1, 1, 1, 1, 3, 2, 2, 1, 3,
	3, 3, 1, 1, 1, 3, 1, 2, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 2, 4, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 4, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 3, 3, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 2, 3, 2, 1, 2, 1, 1,
	1, 4, 2, 1, 1, 2, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 2, 4, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 2, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 4, 2, 1, 2, 3,
	5, 3, 3, 3, 3, 3, 3, 1, 5, 3,
}
var yyChk = [...]int{

	-1000, -6, -7, 25, 54, -8, -9, 11, -10, -11,
	-12, -13, -14, -15, -16, -17, -18, -19, -20, -21,
	49, 50, 12, 34, -24, -27, 15, 62, -98, -32,
	-33, -34, -35, -36, -37, -38, -39, -40, -41, -42,
	-43, -44, -45, -46, -47, 13, 40, 41, 28, -68,
	-112, -116, -78, -123, -127, -121, -87, 61, -62, -104,
	-108, -83, -57, -48, 7, 30, 31, 27, 26, 32,
	35, 43, 44, 21, 38, 42, 29, 53, 56, 58,
	4, 4, 9, -9, -5, 5, 10, 8, -5, -5,
	-5, -5, 8, 10, 8, -5, 5, 8, 5, 4,
	4, 4, 8, -113, 8, 8, 8, 8, 8, 10,
	8, 10, 8, -5, 8, 8, 8, 8, 10, 8,
	10, 8, -4, -5, 6, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, -5, 4, 4,
	8, 8, 10, 55, -22, -23, -13, -14, 10, 10,
	-129, 10, 8, 10, -25, -26, -17, 13, -13, -14,
	-28, -29, 13, -13, -14, 10, 10, -99, -100, -101,
	-13, -14, -52, -102, -103, 59, 23, 24, -69, -70,
	-71, -13, -14, -72, -73, 16, 14, -114, -115, -13,
	-14, -20, -117, -118, -119, -13, -14, -52, -54, -94,
	-95, -81, -82, -120, 22, -20, -53, 36, 37, 20,
	33, 19, 60, -79, -67, -80, -13, -14, -52, -54,
	-81, -82, -20, -124, -125, -126, -71, -13, -14, -52,
	-40, -54, -81, -94, -95, -82, -72, -124, -122, -13,
	-14, -52, -40, -54, -81, -82, -88, -89, -90, -13,
	-14, -52, -54, -91, -44, -92, 51, 10, -63, -13,
	-14, -64, -52, -54, -65, -66, 39, -105, -106, -107,
	-13, -14, -52, -102, -103, -109, -110, -111, -13, -14,
	-52, -20, -84, -85, -86, -13, -14, -52, -54, -34,
	-35, -36, -37, -38, -39, -41, -65, -42, -43, -58,
	-59, -60, -13, -14, -61, -52, 57, -49, -50, -51,
	-13, -14, -52, 10, 5, 9, -23, 7, -26, 9,
	5, -29, 9, 5, 9, -101, -30, -31, -20, -30,
	-5, 8, 8, 9, -70, -74, 10, 8, -4, 4,
	9, -115, 9, -119, -5, 10, 8, -3, 52, 6,
	-3, -2, 47, 48, -2, -5, -5, 9, -80, 9,
	-126, 9, 9, 9, -90, 10, 8, -5, 9, -65,
	8, 4, 9, -107, -30, -30, 9, -111, 9, -86,
	9, -60, 4, 9, -51, -4, 10, 10, 9, -20,
	9, 10, -75, 17, 63, -76, -61, -77, 45, -128,
	-71, 18, 10, 10, -55, -56, -13, -14, 10, 10,
	10, 10, 10, 10, -96, -97, -93, -13, -14, -52,
	-72, -81, -82, -40, -94, -95, -67, 9, 9, 10,
	-129, 9, -5, -5, -128, -71, -5, 4, 9, -56,
	9, -93, 9, 9, 10, 10, 10, 10, 8, -1,
	46, 9, -3, 10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 5, 0, 279, 19, 0, 0, 0,
	0, 0, 0, 40, 0, 0, 0, 191, 18, 25,
	34, 190, 0, 223, 0, 0, 130, 0, 0, 251,
	0, 162, 164, 0, 94, 204, 215, 143, 81, 84,
	62, 65, 0, 113, 114, 106, 225, 232, 129, 263,
	286, 260, 261, 161, 100, 203, 214, 141, 83, 64,
	2, 3, 6, 0, 0, 21, 23, 24, 294, 295,
	292, 297, 0, 293, 0, 26, 29, 0, 31, 32,
	0, 35, 0, 38, 39, 28, 296, 0, 192, 193,
	195, 196, 197, 42, 42, 0, 0, 0, 0, 107,
	109, 110, 111, 112, 0, 0, 0, 0, 226, 228,
	229, 230, 0, 233, 234, 239, 240, 241, 242, 243,
	244, 245, 246, 247, 0, 249, 0, 0, 0, 0,
	0, 0, 0, 0, 131, 132, 134, 135, 136, 137,
	138, 139, 140, 0, 264, 265, 267, 268, 269, 270,
	271, 272, 273, 274, 275, 276, 277, 0, 0, 253,
	254, 255, 256, 257, 258, 259, 0, 165, 166, 168,
	169, 170, 171, 172, 173, 0, 0, 72, 0, 95,
	96, 97, 98, 99, 101, 0, 0, 0, 205, 206,
	208, 209, 210, 42, 42, 0, 216, 217, 219, 220,
	221, 222, 0, 144, 145, 147, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 0,
	85, 86, 88, 89, 90, 91, 0, 0, 66, 67,
	69, 70, 71, 299, 280, 20, 22, 0, 27, 33,
	0, 36, 41, 0, 189, 194, 0, 43, 60, 0,
	0, 200, 201, 105, 108, 116, 118, 0, 0, 117,
	224, 227, 231, 235, 0, 75, 0, 0, 0, 281,
	0, 0, 282, 283, 0, 0, 74, 128, 133, 262,
	266, 285, 252, 163, 167, 184, 0, 174, 93, 102,
	0, 104, 202, 207, 0, 0, 213, 218, 142, 146,
	82, 87, 0, 63, 68, 0, 30, 37, 198, 61,
	199, 73, 0, 0, 0, 122, 123, 124, 0, 287,
	126, 0, 115, 248, 0, 77, 79, 80, 236, 237,
	238, 284, 278, 250, 0, 186, 187, 175, 176, 177,
	178, 179, 180, 181, 182, 183, 0, 211, 212, 92,
	0, 119, 0, 0, 288, 127, 0, 0, 76, 78,
	185, 188, 103, 298, 120, 121, 125, 289, 0, 0,
	0, 290, 0, 291,
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
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:184
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:191
		{
			pop(yylex)
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:194
		{
			pop(yylex)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:207
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:218
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:231
		{
			pop(yylex)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:236
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:252
		{
			pop(yylex)
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:255
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:286
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:289
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:294
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:313
		{
			if set(yylex, meta.NewMust(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:320
		{
			if set(yylex, meta.NewIfFeature(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:327
		{
			if push(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:334
		{
			pop(yylex)
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:337
		{
			pop(yylex)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:349
		{
			pop(yylex)
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:352
		{
			pop(yylex)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:357
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:377
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:387
		{
			pop(yylex)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:400
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:412
		{
			pop(yylex)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:417
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:427
		{
			pop(yylex)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:432
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:449
		{
			yyVAL.token = yyDollar[1].token
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:450
		{
			yyVAL.token = yyDollar[1].token
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:463
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:470
		{
			pop(yylex)
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:473
		{
			pop(yylex)
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:478
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetLenRange(r)) {
				goto ret1
			}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:487
		{
			r, err := meta.NewRange(yyDollar[2].token)
			if chkErr(yylex, err) {
				goto ret1
			}
			if set(yylex, meta.SetValueRange(r)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:499
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:513
		{
			pop(yylex)
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:518
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:542
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:549
		{
			pop(yylex)
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:577
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:584
		{
			pop(yylex)
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:587
		{
			pop(yylex)
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:607
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:626
		{
			pop(yylex)
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:629
		{
			pop(yylex)
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:641
		{
			pop(yylex)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:646
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:663
		{
			pop(yylex)
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:666
		{
			pop(yylex)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:671
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:678
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 202:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:688
		{
			pop(yylex)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:693
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:710
		{
			pop(yylex)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:713
		{
			pop(yylex)
		}
	case 213:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:721
		{
			pop(yylex)
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:726
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:748
		{
			pop(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:758
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:776
		{
			pop(yylex)
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:781
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:796
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:801
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:808
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 250:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:828
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:835
		{
			pop(yylex)
		}
	case 252:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:838
		{
			pop(yylex)
		}
	case 260:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:853
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:858
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 262:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:868
		{
			pop(yylex)
		}
	case 263:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:873
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 278:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:901
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:908
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 280:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:911
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:916
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:926
		{
			yyVAL.boolean = true
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:927
		{
			yyVAL.boolean = false
		}
	case 284:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:930
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 285:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:940
		{
			pop(yylex)
		}
	case 286:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:945
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 289:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:956
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 290:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:961
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 291:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:968
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 292:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:973
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 293:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:980
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 294:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:987
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 295:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:994
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 296:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1001
		{
			if set(yylex, meta.SetYangVersion(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 299:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1012
		{

		}
	}
	goto yystack /* stack new state and value */
}
