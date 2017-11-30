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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:975

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 870

var yyAct = [...]int{

	207, 334, 398, 12, 205, 384, 12, 225, 142, 204,
	291, 324, 288, 296, 271, 264, 312, 211, 223, 256,
	185, 236, 170, 338, 215, 222, 210, 45, 160, 147,
	153, 135, 19, 283, 179, 137, 427, 206, 216, 19,
	11, 423, 44, 11, 46, 19, 43, 42, 41, 40,
	39, 38, 37, 200, 20, 165, 164, 325, 253, 422,
	3, 20, 386, 19, 213, 79, 201, 20, 31, 336,
	15, 244, 395, 76, 19, 83, 84, 85, 86, 166,
	197, 90, 389, 139, 293, 20, 166, 135, 151, 4,
	156, 177, 166, 197, 339, 340, 20, 309, 97, 162,
	425, 173, 424, 181, 187, 135, 218, 218, 174, 229,
	166, 238, 247, 258, 266, 273, 335, 290, 135, 298,
	138, 193, 191, 224, 224, 150, 233, 155, 135, 190,
	192, 208, 221, 221, 131, 232, 161, 139, 172, 171,
	180, 186, 285, 217, 217, 251, 228, 151, 237, 246,
	257, 265, 272, 156, 289, 242, 297, 284, 148, 306,
	162, 282, 281, 280, 279, 278, 277, 276, 261, 260,
	173, 226, 302, 19, 138, 91, 304, 174, 144, 181,
	143, 430, 315, 307, 150, 187, 410, 388, 311, 352,
	155, 351, 320, 394, 323, 20, 322, 161, 393, 392,
	253, 337, 193, 191, 118, 331, 117, 172, 171, 344,
	190, 192, 116, 329, 115, 218, 180, 148, 356, 377,
	166, 197, 186, 177, 316, 341, 379, 386, 212, 110,
	163, 109, 224, 135, 33, 188, 238, 219, 219, 346,
	230, 221, 239, 249, 259, 267, 274, 391, 292, 144,
	299, 143, 217, 332, 383, 333, 258, 350, 33, 428,
	342, 157, 19, 154, 266, 135, 293, 108, 145, 107,
	89, 273, 88, 237, 82, 359, 81, 360, 361, 363,
	242, 141, 355, 19, 20, 365, 390, 387, 290, 140,
	373, 163, 134, 257, 167, 168, 298, 135, 285, 372,
	367, 265, 353, 421, 418, 20, 412, 135, 272, 370,
	409, 261, 260, 284, 135, 371, 188, 282, 281, 280,
	279, 278, 277, 276, 308, 289, 408, 135, 19, 154,
	166, 182, 196, 297, 381, 135, 209, 319, 135, 376,
	374, 19, 268, 177, 369, 176, 219, 305, 366, 364,
	20, 19, 149, 400, 23, 19, 362, 177, 358, 176,
	402, 354, 385, 20, 349, 301, 407, 239, 348, 19,
	404, 406, 347, 20, 345, 343, 330, 20, 405, 403,
	411, 19, 149, 310, 23, 318, 414, 259, 317, 114,
	399, 20, 113, 314, 314, 267, 112, 111, 400, 106,
	419, 105, 274, 20, 104, 402, 103, 182, 102, 100,
	98, 95, 344, 196, 94, 404, 406, 87, 303, 292,
	336, 415, 6, 405, 403, 80, 326, 299, 30, 429,
	80, 300, 417, 368, 357, 399, 327, 413, 133, 132,
	189, 416, 220, 220, 130, 231, 129, 240, 250, 128,
	19, 275, 127, 126, 125, 124, 123, 202, 200, 72,
	195, 122, 121, 120, 67, 66, 119, 75, 64, 65,
	68, 201, 20, 69, 198, 199, 73, 99, 93, 92,
	74, 70, 71, 28, 401, 27, 55, 214, 54, 314,
	314, 76, 268, 227, 77, 56, 78, 166, 197, 194,
	184, 183, 52, 178, 101, 51, 263, 7, 19, 24,
	262, 23, 60, 255, 254, 59, 159, 72, 158, 34,
	397, 189, 67, 66, 49, 75, 64, 65, 68, 401,
	20, 69, 396, 243, 73, 241, 25, 26, 74, 70,
	71, 235, 375, 420, 234, 17, 18, 19, 57, 76,
	270, 220, 77, 269, 78, 200, 72, 61, 203, 53,
	382, 67, 66, 380, 75, 64, 65, 68, 201, 20,
	69, 378, 240, 73, 321, 175, 19, 74, 70, 71,
	176, 169, 50, 252, 200, 248, 245, 58, 76, 287,
	286, 77, 62, 78, 166, 197, 19, 201, 20, 295,
	198, 199, 294, 63, 200, 72, 48, 275, 47, 36,
	67, 66, 35, 75, 64, 65, 68, 201, 20, 69,
	313, 32, 73, 166, 152, 22, 74, 70, 71, 146,
	21, 136, 19, 16, 14, 13, 10, 76, 9, 8,
	77, 72, 78, 166, 197, 29, 67, 66, 5, 75,
	64, 65, 68, 2, 20, 69, 1, 426, 73, 0,
	0, 0, 74, 70, 71, 0, 0, 328, 0, 0,
	0, 19, 0, 76, 0, 0, 77, 0, 78, 166,
	72, 0, 0, 0, 0, 67, 66, 0, 75, 64,
	65, 68, 0, 20, 69, 0, 19, 73, 0, 0,
	0, 74, 70, 71, 0, 72, 0, 0, 0, 0,
	67, 66, 76, 75, 0, 77, 68, 78, 20, 69,
	0, 0, 73, 253, 0, 0, 74, 70, 71, 0,
	0, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 0, 166, 197, 67, 66, 0, 75, 64,
	65, 68, 0, 20, 69, 0, 0, 73, 0, 0,
	96, 74, 70, 71, 0, 0, 0, 0, 0, 0,
	0, 0, 76, 72, 0, 77, 0, 78, 67, 66,
	49, 75, 64, 65, 68, 0, 0, 69, 0, 0,
	73, 0, 72, 0, 74, 70, 71, 67, 66, 0,
	75, 64, 65, 68, 0, 76, 69, 0, 77, 73,
	78, 0, 0, 74, 70, 71, 0, 19, 0, 177,
	0, 176, 0, 0, 76, 200, 0, 77, 0, 78,
	7, 19, 24, 0, 23, 0, 0, 0, 201, 20,
	0, 198, 199, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 0, 0, 25,
	26, 0, 0, 0, 166, 197, 0, 0, 17, 18,
}
var yyPact = [...]int{

	34, -1000, 818, 481, 479, 495, -1000, 425, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 266, 425, 425, 425,
	425, 409, 262, 425, 164, 475, 474, 406, 403, 751,
	-1000, -1000, -1000, -1000, 402, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 473,
	401, 400, 398, 396, 393, 391, 259, 221, 389, 388,
	384, 381, 204, 196, 462, 459, 458, 457, 452, 451,
	450, 449, 448, 445, 442, 440, 425, 435, 434, 282,
	-1000, -1000, 61, 279, 271, 241, 258, 368, -1000, 249,
	251, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 270, -1000,
	342, -1000, 718, 437, 583, 804, 804, -1000, 32, -1000,
	19, 160, 270, 619, 683, -1000, 26, -1000, 50, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -25, -1000, -1000, -1000, 426, 356, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 411, -1000, 338, -1000, -1000, 148,
	-1000, -1000, 315, -1000, 86, -1000, -1000, -1000, 374, 270,
	-1000, -1000, -1000, -1000, 770, 770, 425, 380, 377, 328,
	-1000, -1000, -1000, -1000, -1000, 186, 420, 432, 658, -1000,
	-1000, -1000, -1000, 367, 437, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 425, -1000, 425, 63, 414,
	46, 46, 425, 366, 583, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 365, 804, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 363, 359, -1000, -1000,
	-1000, -1000, -1000, -1000, 355, 19, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 181, 425, 352, -1000, -1000, 18, -1000,
	-1000, -1000, 210, 430, 349, 270, -1000, -1000, -1000, -1000,
	770, 770, 347, 619, -1000, -1000, -1000, -1000, -1000, 340,
	683, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 339, 26, -1000, -1000,
	-1000, -1000, -1000, 429, 335, 50, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 420, -1000, -1000, 289, -1000, -1000, 280,
	-1000, -1000, 331, 770, -1000, 330, 209, -1000, -1000, -1000,
	-1000, -1000, -1000, 208, 277, -25, -1000, -1000, -1000, -1000,
	-1000, -1000, 177, 72, 276, 237, -1000, 189, 188, -1000,
	-1000, 183, 62, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 563, -25, -1000, -1000, 583, -1000, -1000, -1000,
	317, 301, -1000, -1000, -1000, -1000, -1000, -1000, 176, -1000,
	-1000, 170, -1000, -1000, -1000, -1000, -1000, -1000, 297, 425,
	43, -1000, 76, 425, -1000, -1000, 428, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 295, 563, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 534, -1000, -1000,
	-1000, 294, -1000, 49, -1000, -1000, 31, 92, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -11, 250, 414, -1000, 171,
	-1000,
}
var yyPgo = [...]int{

	0, 657, 23, 1, 11, 57, 656, 653, 648, 645,
	422, 639, 638, 636, 37, 0, 635, 634, 70, 633,
	631, 35, 630, 629, 29, 625, 624, 30, 68, 621,
	228, 16, 620, 612, 609, 52, 51, 50, 49, 48,
	47, 46, 42, 27, 44, 608, 606, 603, 602, 599,
	13, 131, 336, 592, 590, 589, 12, 10, 587, 586,
	585, 33, 583, 9, 582, 581, 22, 38, 7, 575,
	574, 571, 563, 560, 559, 558, 4, 26, 17, 557,
	553, 550, 14, 548, 544, 541, 21, 535, 533, 2,
	25, 18, 532, 520, 519, 518, 516, 28, 56, 55,
	515, 514, 513, 19, 512, 510, 506, 15, 505, 504,
	503, 34, 502, 501, 500, 20, 499, 495, 493, 488,
	64, 487, 24, 486, 5, 8,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 32,
	32, 46, 46, 47, 48, 48, 49, 49, 50, 50,
	50, 51, 52, 45, 45, 53, 54, 54, 55, 55,
	56, 56, 56, 56, 57, 41, 59, 59, 59, 59,
	59, 59, 58, 60, 60, 61, 62, 33, 64, 65,
	65, 66, 66, 66, 66, 4, 4, 68, 67, 69,
	70, 70, 71, 71, 71, 71, 71, 73, 73, 36,
	74, 75, 75, 63, 63, 76, 76, 76, 76, 76,
	76, 76, 79, 44, 80, 80, 81, 81, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 83, 40, 40, 84, 84, 85, 85, 86,
	86, 86, 86, 86, 86, 88, 89, 89, 89, 89,
	89, 89, 89, 89, 87, 87, 92, 93, 93, 29,
	94, 95, 95, 96, 96, 97, 97, 97, 97, 97,
	98, 99, 42, 100, 101, 101, 102, 102, 103, 103,
	103, 103, 103, 43, 104, 105, 105, 106, 106, 107,
	107, 107, 107, 34, 109, 108, 110, 110, 111, 111,
	111, 35, 112, 113, 114, 114, 90, 90, 91, 115,
	115, 115, 115, 115, 115, 115, 115, 115, 115, 115,
	116, 39, 39, 118, 118, 118, 118, 118, 118, 117,
	117, 37, 119, 120, 121, 121, 122, 122, 122, 122,
	122, 122, 122, 122, 122, 122, 78, 5, 5, 3,
	2, 2, 77, 38, 123, 72, 72, 124, 124, 1,
	14, 15, 12, 13, 125, 125,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 3, 3, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 1, 3, 4, 0, 1, 1, 1,
	1, 1, 2, 1, 2, 4, 2, 4, 2, 1,
	2, 1, 1, 1, 1, 1, 1, 3, 2, 2,
	1, 3, 3, 1, 1, 1, 3, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 4, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 4, 0, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 3, 3,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 3, 3, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 2, 3, 2, 1, 2, 1, 1,
	1, 4, 2, 1, 1, 2, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 2, 4, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 3, 1,
	1, 1, 3, 4, 2, 1, 2, 3, 5, 3,
	3, 3, 3, 3, 1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -94, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, -46, 29,
	-64, -108, -112, -74, -119, -123, -117, -83, -58, -100,
	-104, -79, -53, -47, 31, 32, 28, 27, 33, 36,
	44, 45, 22, 39, 43, 30, 54, 57, 59, -5,
	5, 10, 8, -5, -5, -5, -5, 8, 10, 8,
	-5, 11, 4, 4, 8, 8, 9, -28, 8, 4,
	8, -109, 8, 8, 8, 8, 8, 10, 8, 10,
	8, 8, 8, 8, 8, 10, 8, 10, 8, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, -5, 4, 4, 10, 56, -20, -21, -14, -15,
	10, 10, -125, 10, 8, 10, -23, -24, -18, 14,
	-14, -15, -26, -27, 14, -14, -15, 10, -95, -96,
	-97, -14, -15, -51, -98, -99, 60, 24, 25, -65,
	-66, -67, -14, -15, -68, -69, 17, 15, -110, -111,
	-14, -15, -30, -113, -114, -115, -14, -15, -51, -52,
	-90, -91, -77, -78, -116, 23, -30, 61, 37, 38,
	21, 34, 20, -75, -63, -76, -14, -15, -51, -52,
	-77, -78, -30, -120, -121, -122, -67, -14, -15, -51,
	-52, -77, -90, -91, -78, -68, -120, -118, -14, -15,
	-51, -52, -77, -78, -84, -85, -86, -14, -15, -51,
	-52, -87, -44, -88, 52, -59, -14, -15, -60, -51,
	-52, -61, -62, 40, -101, -102, -103, -14, -15, -51,
	-98, -99, -105, -106, -107, -14, -15, -51, -30, -80,
	-81, -82, -14, -15, -51, -52, -35, -36, -37, -38,
	-39, -40, -41, -61, -42, -43, -54, -55, -56, -14,
	-15, -57, -51, 58, -48, -49, -50, -14, -15, -51,
	5, 9, -21, 7, -24, 9, 11, -27, 9, 11,
	9, -97, -31, -32, -30, -31, -5, 8, 8, 9,
	-66, -70, 10, 8, -4, -5, 6, 4, 9, -111,
	9, -115, -5, -5, -3, 53, 6, -3, -2, 48,
	49, -2, -5, 9, -76, 9, -122, 9, 9, 9,
	-86, 10, 8, -5, 9, -61, 8, 4, 9, -103,
	-31, -31, 9, -107, 9, -82, 9, -56, 4, 9,
	-50, -4, 10, 10, 9, -30, 9, 10, -71, 18,
	-72, -57, -73, 46, -124, -67, 19, 10, 10, 10,
	10, 10, 10, 10, 10, 10, -92, -93, -89, -14,
	-15, -51, -68, -77, -78, -90, -91, -63, 9, 9,
	10, -125, 9, -5, -124, -67, -5, 4, 9, -89,
	9, 9, 10, 10, 10, 8, -1, 47, 9, -3,
	10,
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
	267, 16, 0, 0, 0, 0, 0, 0, 37, 0,
	0, 15, 22, 31, 2, 3, 1, 42, 181, 180,
	0, 213, 0, 0, 121, 0, 0, 241, 0, 153,
	155, 86, 194, 205, 134, 73, 76, 61, 64, 98,
	215, 222, 120, 252, 274, 249, 250, 152, 92, 193,
	204, 132, 75, 63, 6, 0, 0, 18, 20, 21,
	282, 283, 280, 284, 0, 281, 0, 23, 26, 0,
	28, 29, 0, 32, 0, 35, 36, 25, 0, 182,
	183, 185, 186, 187, 43, 43, 0, 0, 0, 0,
	99, 101, 102, 103, 104, 0, 0, 0, 0, 216,
	218, 219, 220, 0, 223, 224, 229, 230, 231, 232,
	233, 234, 235, 236, 237, 0, 239, 0, 0, 0,
	0, 0, 0, 0, 122, 123, 125, 126, 127, 128,
	129, 130, 131, 0, 253, 254, 256, 257, 258, 259,
	260, 261, 262, 263, 264, 265, 0, 0, 243, 244,
	245, 246, 247, 248, 0, 156, 157, 159, 160, 161,
	162, 163, 164, 0, 0, 0, 87, 88, 89, 90,
	91, 93, 0, 0, 0, 195, 196, 198, 199, 200,
	43, 43, 0, 206, 207, 209, 210, 211, 212, 0,
	135, 136, 138, 139, 140, 141, 142, 143, 144, 145,
	146, 147, 148, 149, 150, 151, 0, 77, 78, 80,
	81, 82, 83, 0, 0, 65, 66, 68, 69, 70,
	268, 17, 19, 0, 24, 30, 0, 33, 38, 0,
	179, 184, 0, 44, 59, 0, 0, 190, 191, 97,
	100, 108, 110, 0, 0, 105, 106, 109, 214, 217,
	221, 225, 0, 0, 0, 0, 269, 0, 0, 270,
	271, 0, 0, 119, 124, 251, 255, 273, 242, 154,
	158, 174, 0, 165, 85, 94, 0, 96, 192, 197,
	0, 0, 203, 208, 133, 137, 74, 79, 0, 62,
	67, 0, 27, 34, 188, 60, 189, 71, 0, 0,
	113, 114, 115, 0, 275, 117, 0, 107, 238, 72,
	226, 227, 228, 272, 266, 240, 0, 176, 177, 166,
	167, 168, 169, 170, 171, 172, 173, 0, 201, 202,
	84, 0, 111, 0, 276, 118, 0, 0, 175, 178,
	95, 285, 112, 116, 277, 0, 0, 0, 278, 0,
	279,
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
		//line parser.y:143
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
		//line parser.y:151
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
		//line parser.y:165
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:180
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:187
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:190
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:203
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:214
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:227
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:232
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:248
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:251
		{
			pop(yylex)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:288
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:291
		{
			pop(yylex)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:296
		{
			if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:315
		{
			if set(yylex, meta.NewIfFeature(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:322
		{
			if set(yylex, meta.NewWhen(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:329
		{
			pop(yylex)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:332
		{
			pop(yylex)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:337
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:357
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:367
		{
			pop(yylex)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:380
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:392
		{
			pop(yylex)
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:397
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:407
		{
			pop(yylex)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:412
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:429
		{
			yyVAL.token = yyDollar[1].token
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:430
		{
			yyVAL.token = yyDollar[1].token
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:433
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:443
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:450
		{
			pop(yylex)
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			pop(yylex)
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:466
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:480
		{
			pop(yylex)
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:485
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:509
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:516
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:544
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:551
		{
			pop(yylex)
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:554
		{
			pop(yylex)
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:577
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:600
		{
			pop(yylex)
		}
	case 175:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:603
		{
			pop(yylex)
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:615
		{
			pop(yylex)
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:620
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:637
		{
			pop(yylex)
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:640
		{
			pop(yylex)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:645
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:652
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:662
		{
			pop(yylex)
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:667
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:684
		{
			pop(yylex)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:687
		{
			pop(yylex)
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:695
		{
			pop(yylex)
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:700
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:722
		{
			pop(yylex)
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:732
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:750
		{
			pop(yylex)
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:755
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:770
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:775
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:782
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:802
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:809
		{
			pop(yylex)
		}
	case 242:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:812
		{
			pop(yylex)
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:826
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 250:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:831
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 251:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:841
		{
			pop(yylex)
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:846
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 266:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:873
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:880
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 268:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:883
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:888
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:898
		{
			yyVAL.boolean = true
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:899
		{
			yyVAL.boolean = false
		}
	case 272:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:902
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 273:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:912
		{
			pop(yylex)
		}
	case 274:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:917
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 277:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:928
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 278:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:933
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 279:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:940
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 280:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:945
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 281:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:952
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 282:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:959
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 283:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:966
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
