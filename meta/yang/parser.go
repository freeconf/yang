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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:899

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 649

var yyAct = [...]int{

	195, 293, 193, 12, 207, 350, 12, 338, 209, 135,
	259, 256, 192, 284, 197, 206, 45, 241, 285, 273,
	44, 196, 43, 42, 41, 40, 76, 39, 201, 297,
	213, 228, 38, 153, 176, 235, 80, 81, 82, 83,
	128, 194, 87, 161, 11, 37, 157, 11, 170, 156,
	146, 19, 251, 140, 130, 199, 298, 299, 377, 3,
	70, 225, 334, 339, 202, 65, 64, 19, 73, 62,
	63, 66, 15, 20, 67, 19, 295, 71, 339, 270,
	132, 72, 68, 69, 137, 144, 136, 149, 4, 20,
	337, 267, 74, 125, 88, 75, 155, 20, 164, 380,
	172, 178, 260, 204, 204, 180, 165, 215, 221, 230,
	237, 243, 260, 258, 218, 182, 179, 208, 208, 373,
	372, 131, 181, 294, 205, 205, 143, 253, 148, 347,
	132, 252, 128, 250, 249, 248, 247, 154, 246, 163,
	144, 171, 177, 245, 203, 203, 149, 361, 214, 220,
	229, 236, 242, 155, 257, 232, 244, 141, 231, 210,
	223, 164, 162, 346, 345, 128, 128, 344, 341, 165,
	172, 131, 375, 343, 374, 128, 178, 276, 150, 342,
	180, 143, 340, 329, 263, 138, 272, 148, 134, 296,
	182, 179, 133, 265, 154, 303, 268, 181, 127, 19,
	147, 204, 163, 292, 280, 19, 328, 104, 19, 301,
	291, 171, 141, 215, 128, 208, 31, 177, 289, 300,
	198, 20, 205, 378, 128, 162, 33, 20, 230, 305,
	20, 128, 225, 371, 128, 237, 262, 311, 128, 368,
	19, 243, 203, 308, 128, 137, 94, 136, 363, 360,
	33, 318, 319, 310, 214, 309, 258, 253, 323, 317,
	359, 252, 20, 250, 249, 248, 247, 325, 246, 229,
	321, 332, 330, 245, 232, 313, 236, 231, 327, 7,
	19, 24, 242, 23, 279, 19, 244, 168, 19, 167,
	168, 283, 167, 282, 336, 324, 266, 257, 322, 320,
	19, 142, 20, 23, 112, 314, 111, 20, 25, 26,
	20, 352, 19, 316, 106, 357, 105, 17, 18, 353,
	173, 185, 20, 158, 159, 355, 356, 358, 312, 307,
	238, 269, 354, 306, 20, 19, 147, 362, 19, 304,
	168, 302, 167, 365, 19, 142, 188, 23, 290, 86,
	352, 85, 351, 364, 357, 369, 366, 20, 353, 189,
	20, 303, 186, 187, 355, 356, 20, 79, 278, 78,
	271, 354, 277, 110, 109, 108, 107, 275, 275, 379,
	103, 102, 7, 19, 24, 101, 23, 100, 99, 97,
	173, 351, 70, 95, 92, 91, 185, 65, 64, 48,
	73, 62, 63, 66, 84, 20, 67, 264, 19, 71,
	295, 25, 26, 72, 68, 69, 77, 70, 77, 286,
	17, 18, 65, 64, 74, 73, 6, 75, 66, 261,
	20, 67, 30, 367, 71, 225, 326, 315, 72, 68,
	69, 287, 126, 124, 123, 122, 121, 120, 119, 19,
	118, 117, 275, 275, 116, 238, 190, 188, 70, 184,
	115, 114, 113, 65, 64, 96, 73, 62, 63, 66,
	189, 20, 67, 186, 187, 71, 90, 89, 28, 72,
	68, 69, 27, 54, 370, 200, 53, 55, 19, 183,
	74, 175, 174, 75, 51, 331, 188, 70, 169, 98,
	50, 234, 65, 64, 233, 73, 62, 63, 66, 189,
	20, 67, 59, 227, 71, 226, 19, 58, 72, 68,
	69, 152, 151, 34, 188, 70, 349, 348, 217, 74,
	65, 64, 75, 73, 62, 63, 66, 189, 20, 67,
	288, 216, 71, 212, 19, 211, 72, 68, 69, 56,
	240, 239, 60, 70, 191, 52, 335, 74, 65, 64,
	75, 73, 62, 63, 66, 333, 20, 67, 281, 166,
	71, 93, 160, 49, 72, 68, 69, 224, 222, 219,
	57, 255, 254, 61, 70, 74, 47, 46, 75, 65,
	64, 48, 73, 62, 63, 66, 36, 35, 67, 274,
	32, 71, 145, 70, 22, 72, 68, 69, 65, 64,
	139, 73, 62, 63, 66, 21, 74, 67, 129, 75,
	71, 16, 14, 19, 72, 68, 69, 167, 13, 10,
	9, 188, 8, 29, 5, 74, 2, 1, 75, 376,
	0, 0, 0, 0, 189, 20, 0, 186, 187,
}
var yyPact = [...]int{

	33, -1000, 267, 478, 474, 370, -1000, 411, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 359, 411, 411, 411,
	411, 396, 341, 411, 83, 473, 472, 387, 386, 562,
	-1000, -1000, -1000, -1000, 385, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 461, 381,
	380, 379, 377, 373, 372, 197, 306, 368, 367, 366,
	365, 296, 458, 457, 456, 450, 447, 446, 444, 443,
	442, 441, 440, 439, 411, 438, 188, -1000, -1000, 195,
	182, 178, 76, 175, 331, -1000, 186, 168, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 299, -1000, 272, -1000, 38,
	436, 503, 325, 325, -1000, -1000, 62, 192, 299, 38,
	395, -1000, 54, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -16, -1000, -1000, 424, 227,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 400, -1000, 287,
	-1000, -1000, 80, -1000, -1000, 322, -1000, 68, -1000, -1000,
	-1000, 361, 299, -1000, -1000, -1000, 581, 581, 364, 360,
	275, -1000, -1000, -1000, -1000, -1000, 283, 413, 437, 531,
	-1000, -1000, -1000, -1000, 339, 436, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 411, -1000, 70, 404, 8, 8,
	411, 332, 503, -1000, -1000, -1000, -1000, -1000, -1000, 330,
	325, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	324, 320, 62, -1000, -1000, -1000, -1000, 245, 411, 319,
	-1000, -1000, 21, -1000, 297, 433, 304, 299, -1000, -1000,
	-1000, 581, 581, 290, 38, -1000, -1000, -1000, -1000, 289,
	395, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 286, 54, -1000, -1000, -1000, -1000,
	432, -1000, -1000, -1000, 413, -1000, -1000, 196, -1000, -1000,
	173, -1000, -1000, 263, 581, -1000, 262, -1000, -1000, -1000,
	-1000, -1000, -1000, 44, 172, -16, -1000, -1000, -1000, -1000,
	-1000, -1000, 158, 169, 163, -1000, 157, 154, -1000, -1000,
	153, 119, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	610, -16, -1000, -1000, 503, -1000, -1000, -1000, 251, 240,
	-1000, -1000, -1000, -1000, -1000, -1000, 137, 237, -1000, -1000,
	-1000, -1000, -1000, 239, 411, 59, -1000, 411, -1000, 429,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 230, 610,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 475, -1000,
	-1000, -1000, 224, -1000, 110, -1000, 109, 164, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 11, 214, 404, -1000, 89,
	-1000,
}
var yyPgo = [...]int{

	0, 639, 29, 1, 13, 18, 637, 636, 634, 633,
	426, 632, 630, 629, 41, 0, 628, 622, 72, 621,
	618, 54, 615, 610, 53, 604, 602, 50, 216, 600,
	220, 19, 599, 597, 596, 45, 32, 27, 25, 24,
	23, 22, 20, 16, 587, 586, 583, 582, 581, 11,
	10, 580, 579, 578, 52, 577, 12, 573, 572, 43,
	64, 8, 569, 568, 565, 556, 555, 554, 2, 21,
	14, 552, 551, 550, 17, 549, 545, 543, 30, 541,
	528, 5, 15, 4, 527, 526, 523, 522, 521, 33,
	49, 46, 517, 515, 513, 31, 512, 504, 501, 35,
	500, 499, 498, 48, 494, 492, 491, 34, 489, 487,
	486, 55, 485, 28, 483, 7, 9,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 32, 32,
	45, 45, 46, 47, 47, 48, 48, 49, 49, 49,
	50, 41, 52, 52, 52, 52, 51, 53, 53, 54,
	55, 33, 57, 58, 58, 59, 59, 59, 59, 4,
	4, 61, 60, 62, 63, 63, 64, 64, 64, 64,
	36, 66, 67, 67, 56, 56, 68, 68, 68, 68,
	68, 71, 44, 72, 72, 73, 73, 74, 74, 74,
	74, 74, 74, 74, 74, 74, 74, 74, 74, 75,
	40, 40, 76, 76, 77, 77, 78, 78, 78, 80,
	81, 81, 81, 81, 81, 81, 81, 79, 79, 84,
	85, 85, 29, 86, 87, 87, 88, 88, 89, 89,
	89, 89, 90, 91, 42, 92, 93, 93, 94, 94,
	95, 95, 95, 95, 43, 96, 97, 97, 98, 98,
	99, 99, 99, 34, 101, 100, 102, 102, 103, 103,
	103, 35, 104, 105, 106, 106, 82, 82, 83, 107,
	107, 107, 107, 107, 107, 107, 107, 107, 108, 39,
	109, 109, 37, 110, 111, 112, 112, 113, 113, 113,
	113, 113, 113, 113, 113, 70, 5, 5, 3, 2,
	2, 69, 38, 114, 65, 65, 115, 115, 1, 14,
	15, 12, 13, 116, 116,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 4, 0, 1, 1, 1, 2, 1, 2, 4,
	2, 4, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 3, 2, 2, 1, 3, 3, 1, 1, 3,
	4, 2, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 2, 4, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 4, 0, 1, 1, 2, 1, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 4, 1,
	1, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	3, 3, 2, 2, 4, 2, 0, 1, 1, 2,
	1, 1, 3, 3, 4, 2, 0, 1, 1, 2,
	1, 1, 1, 2, 3, 2, 1, 2, 1, 1,
	1, 4, 2, 1, 1, 2, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 3, 2,
	2, 2, 4, 2, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 4, 2, 1, 2, 3, 5, 3, 3,
	3, 3, 3, 1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -86, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, 29, -57,
	-100, -104, -66, -110, -114, -109, -75, -51, -92, -96,
	-71, -46, 31, 32, 28, 27, 33, 36, 44, 45,
	22, 39, 43, 30, 54, 57, -5, 5, 10, 8,
	-5, -5, -5, -5, 8, 10, 8, -5, 11, 4,
	4, 8, 8, 9, -28, 8, 4, 8, -101, 8,
	8, 8, 8, 8, 10, 10, 8, 8, 8, 8,
	8, 10, 8, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, -5, 4, 10, 56, -20,
	-21, -14, -15, 10, 10, -116, 10, 8, 10, -23,
	-24, -18, 14, -14, -15, -26, -27, 14, -14, -15,
	10, -87, -88, -89, -14, -15, -90, -91, 24, 25,
	-58, -59, -60, -14, -15, -61, -62, 17, 15, -102,
	-103, -14, -15, -30, -105, -106, -107, -14, -15, -82,
	-83, -69, -70, -108, 23, -30, 37, 38, 21, 34,
	20, -67, -56, -68, -14, -15, -69, -70, -30, -111,
	-112, -113, -60, -14, -15, -69, -82, -83, -70, -61,
	-111, -76, -77, -78, -14, -15, -79, -80, 52, -52,
	-14, -15, -53, -54, -55, 40, -93, -94, -95, -14,
	-15, -90, -91, -97, -98, -99, -14, -15, -30, -72,
	-73, -74, -14, -15, -35, -36, -37, -38, -39, -40,
	-41, -54, -42, -43, -47, -48, -49, -14, -15, -50,
	58, 5, 9, -21, 7, -24, 9, 11, -27, 9,
	11, 9, -89, -31, -32, -30, -31, 8, 8, 9,
	-59, -63, 10, 8, -4, -5, 6, 4, 9, -103,
	9, -107, -5, -3, 53, 6, -3, -2, 48, 49,
	-2, -5, 9, -68, 9, -113, 9, 9, -78, 10,
	8, -5, 9, -54, 8, 4, 9, -95, -31, -31,
	9, -99, 9, -74, 9, -49, 4, -4, 10, 10,
	9, -30, 9, -64, 18, -65, -50, 46, -115, 19,
	10, 10, 10, 10, 10, 10, 10, 10, -84, -85,
	-81, -14, -15, -61, -69, -70, -82, -83, -56, 9,
	9, 10, -116, 9, -5, -115, -5, 4, 9, -81,
	9, 9, 10, 10, 10, 8, -1, 47, 9, -3,
	10,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 226, 16, 0,
	0, 0, 0, 0, 0, 37, 0, 0, 15, 22,
	31, 2, 3, 1, 42, 154, 153, 0, 183, 0,
	0, 102, 0, 0, 209, 130, 132, 72, 166, 176,
	113, 60, 63, 82, 185, 192, 101, 213, 233, 210,
	211, 129, 76, 165, 175, 111, 62, 6, 0, 0,
	18, 20, 21, 241, 242, 239, 243, 0, 240, 0,
	23, 26, 0, 28, 29, 0, 32, 0, 35, 36,
	25, 0, 155, 156, 158, 159, 43, 43, 0, 0,
	0, 83, 85, 86, 87, 88, 0, 0, 0, 0,
	186, 188, 189, 190, 0, 193, 194, 199, 200, 201,
	202, 203, 204, 205, 0, 207, 0, 0, 0, 0,
	0, 0, 103, 104, 106, 107, 108, 109, 110, 0,
	214, 215, 217, 218, 219, 220, 221, 222, 223, 224,
	0, 0, 133, 134, 136, 137, 138, 0, 0, 0,
	73, 74, 75, 77, 0, 0, 0, 167, 168, 170,
	171, 43, 43, 0, 177, 178, 180, 181, 182, 0,
	114, 115, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 0, 64, 65, 67, 68, 69,
	0, 227, 17, 19, 0, 24, 30, 0, 33, 38,
	0, 152, 157, 0, 44, 58, 0, 162, 163, 81,
	84, 92, 94, 0, 0, 89, 90, 93, 184, 187,
	191, 195, 0, 0, 0, 228, 0, 0, 229, 230,
	0, 0, 100, 105, 212, 216, 232, 131, 135, 147,
	0, 139, 71, 78, 0, 80, 164, 169, 0, 0,
	174, 179, 112, 116, 61, 66, 0, 0, 27, 34,
	160, 59, 161, 0, 0, 97, 98, 0, 234, 0,
	91, 206, 196, 197, 198, 231, 225, 208, 0, 149,
	150, 140, 141, 142, 143, 144, 145, 146, 0, 172,
	173, 70, 0, 95, 0, 235, 0, 0, 148, 151,
	79, 244, 96, 99, 236, 0, 0, 0, 237, 0,
	238,
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
	52, 53, 54, 55, 56, 57, 58,
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
		//line parser.y:140
		{
			l := yylex.(*lexer)
			if l.parent != nil {
				l.Error("expected submodule for include")
				goto ret1
			}
			yylex.(*lexer).stack.Push(meta.NewModule(yyDollar[2].token))
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:148
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
		//line parser.y:162
		{
			if set(yylex, meta.SetNamespace(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:177
		{
			if push(yylex, meta.NewRevision(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:184
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:187
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:200
		{
			if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:211
		{
			if set(yylex, meta.SetPrefix(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:224
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:229
		{
			if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:245
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:248
		{
			pop(yylex)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
			pop(yylex)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:287
		{
			pop(yylex)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:292
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:311
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:321
		{
			pop(yylex)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:332
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:344
		{
			pop(yylex)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:349
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:359
		{
			pop(yylex)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:364
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:381
		{
			yyVAL.token = yyDollar[1].token
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:382
		{
			yyVAL.token = yyDollar[1].token
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:395
		{
			if push(yylex, meta.NewDataType(peek(yylex).(meta.HasDataType), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:402
		{
			pop(yylex)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:405
		{
			pop(yylex)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:410
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:417
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:427
		{
			pop(yylex)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:432
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:454
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:461
		{
			pop(yylex)
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:487
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			pop(yylex)
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:497
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:520
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:542
		{
			pop(yylex)
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:545
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:557
		{
			pop(yylex)
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:562
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:578
		{
			pop(yylex)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:581
		{
			pop(yylex)
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:586
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:593
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:603
		{
			pop(yylex)
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:608
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:624
		{
			pop(yylex)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:627
		{
			pop(yylex)
		}
	case 174:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:635
		{
			pop(yylex)
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:640
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:661
		{
			pop(yylex)
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:671
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:689
		{
			pop(yylex)
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:694
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:709
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:714
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:721
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:739
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:746
		{
			pop(yylex)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:751
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:756
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:766
		{
			pop(yylex)
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:771
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:797
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:804
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:807
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:812
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:822
		{
			yyVAL.boolean = true
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:823
		{
			yyVAL.boolean = false
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:826
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 232:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:836
		{
			pop(yylex)
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:841
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:852
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 237:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:857
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:864
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:869
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:876
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:883
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:890
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
