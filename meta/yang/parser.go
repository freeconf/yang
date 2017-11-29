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

//line parser.y:933

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 688

var yyAct = [...]int{

	201, 304, 199, 12, 213, 212, 12, 365, 296, 352,
	141, 295, 198, 265, 215, 203, 79, 269, 262, 45,
	247, 44, 202, 43, 208, 42, 83, 84, 85, 86,
	41, 234, 90, 241, 219, 40, 308, 207, 176, 182,
	152, 163, 15, 167, 200, 284, 134, 11, 146, 39,
	11, 19, 389, 159, 38, 3, 37, 136, 19, 162,
	393, 204, 205, 309, 310, 231, 257, 33, 354, 174,
	31, 19, 153, 20, 174, 281, 306, 347, 354, 391,
	20, 390, 396, 138, 4, 130, 278, 143, 150, 142,
	155, 33, 143, 20, 142, 388, 266, 224, 134, 161,
	97, 170, 19, 178, 184, 351, 210, 210, 186, 185,
	221, 227, 236, 243, 249, 171, 264, 266, 271, 188,
	362, 214, 214, 305, 20, 168, 187, 137, 211, 211,
	147, 91, 149, 259, 154, 258, 138, 256, 356, 255,
	134, 134, 376, 160, 254, 169, 150, 177, 183, 253,
	209, 209, 155, 238, 220, 226, 235, 242, 248, 161,
	263, 361, 270, 252, 179, 191, 134, 170, 251, 216,
	250, 237, 156, 321, 244, 320, 178, 229, 360, 359,
	137, 171, 184, 358, 134, 144, 186, 185, 147, 140,
	149, 168, 279, 274, 276, 307, 154, 188, 357, 303,
	139, 314, 133, 160, 187, 312, 294, 210, 293, 287,
	291, 169, 283, 117, 300, 116, 19, 355, 134, 221,
	177, 302, 214, 19, 286, 286, 183, 164, 165, 211,
	342, 134, 311, 322, 236, 134, 394, 179, 20, 19,
	148, 243, 23, 191, 316, 20, 134, 249, 134, 341,
	231, 209, 107, 319, 19, 280, 174, 387, 173, 19,
	153, 20, 264, 220, 384, 328, 259, 334, 258, 271,
	256, 115, 255, 114, 332, 238, 20, 254, 235, 378,
	336, 20, 253, 329, 330, 242, 339, 340, 375, 374,
	273, 248, 345, 237, 19, 324, 252, 343, 338, 286,
	286, 251, 244, 250, 335, 109, 263, 108, 349, 333,
	7, 19, 24, 270, 23, 89, 20, 88, 82, 353,
	81, 19, 367, 174, 290, 173, 372, 371, 19, 194,
	174, 331, 173, 20, 327, 323, 368, 370, 373, 25,
	26, 318, 195, 20, 369, 192, 193, 344, 17, 18,
	20, 377, 317, 315, 313, 301, 379, 282, 380, 325,
	382, 289, 288, 113, 112, 367, 366, 111, 110, 372,
	371, 106, 385, 105, 104, 381, 314, 306, 103, 368,
	370, 7, 19, 24, 102, 23, 100, 369, 98, 95,
	94, 72, 87, 80, 297, 395, 67, 66, 49, 75,
	64, 65, 68, 275, 20, 69, 80, 272, 73, 366,
	25, 26, 74, 70, 71, 6, 383, 337, 326, 17,
	18, 30, 19, 76, 298, 132, 77, 131, 78, 196,
	194, 72, 190, 129, 128, 127, 67, 66, 126, 75,
	64, 65, 68, 195, 20, 69, 192, 193, 73, 125,
	124, 277, 74, 70, 71, 19, 148, 386, 23, 123,
	122, 19, 121, 76, 120, 119, 77, 118, 78, 194,
	72, 99, 93, 92, 28, 67, 66, 20, 75, 64,
	65, 68, 195, 20, 69, 27, 55, 73, 206, 54,
	56, 74, 70, 71, 189, 19, 181, 180, 52, 175,
	101, 51, 76, 194, 72, 77, 240, 78, 239, 67,
	66, 60, 75, 64, 65, 68, 195, 20, 69, 233,
	232, 73, 59, 158, 299, 74, 70, 71, 19, 157,
	34, 364, 363, 223, 222, 218, 76, 72, 217, 77,
	57, 78, 67, 66, 246, 75, 64, 65, 68, 245,
	20, 69, 61, 197, 73, 53, 350, 348, 74, 70,
	71, 19, 346, 292, 172, 166, 50, 230, 228, 76,
	72, 225, 77, 58, 78, 67, 66, 261, 75, 64,
	65, 68, 260, 20, 69, 62, 268, 73, 267, 63,
	96, 74, 70, 71, 48, 47, 46, 36, 35, 285,
	32, 151, 76, 72, 22, 77, 145, 78, 67, 66,
	49, 75, 64, 65, 68, 21, 135, 69, 16, 14,
	73, 13, 72, 10, 74, 70, 71, 67, 66, 9,
	75, 64, 65, 68, 8, 76, 69, 29, 77, 73,
	78, 19, 5, 74, 70, 71, 2, 1, 392, 0,
	72, 0, 0, 0, 76, 67, 66, 77, 75, 78,
	0, 68, 19, 20, 69, 0, 173, 73, 231, 0,
	194, 74, 70, 71, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 195, 20, 0, 192, 193,
}
var yyPact = [...]int{

	29, -1000, 298, 481, 470, 369, -1000, 401, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 310, 401, 401, 401,
	401, 384, 307, 401, 120, 469, 468, 382, 381, 581,
	-1000, -1000, -1000, -1000, 380, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 467,
	378, 376, 370, 366, 365, 363, 242, 297, 360, 359,
	356, 355, 263, 205, 463, 461, 460, 458, 456, 455,
	446, 445, 434, 431, 430, 429, 401, 423, 421, 192,
	-1000, -1000, 89, 190, 179, 84, 175, 226, -1000, 58,
	162, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 203, -1000,
	241, -1000, 548, 409, 482, 308, 308, -1000, -1000, 45,
	210, 203, 548, 628, -1000, 38, -1000, 89, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-10, -1000, -1000, -1000, 402, 281, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 396, -1000, 442, -1000, -1000, 75, -1000,
	-1000, 246, -1000, 64, -1000, -1000, -1000, 348, 203, -1000,
	-1000, -1000, 600, 600, 354, 353, 315, -1000, -1000, -1000,
	-1000, -1000, 198, 388, 420, 515, -1000, -1000, -1000, -1000,
	346, 409, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	401, -1000, 70, 371, 15, 15, 401, 345, 482, -1000,
	-1000, -1000, -1000, -1000, -1000, 344, 308, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 343, 332, 45, -1000,
	-1000, -1000, -1000, 165, 401, 326, -1000, -1000, 25, -1000,
	351, 414, 325, 203, -1000, -1000, -1000, 600, 600, 322,
	548, -1000, -1000, -1000, -1000, 300, 628, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	295, 38, -1000, -1000, -1000, -1000, 413, 289, 89, -1000,
	-1000, -1000, -1000, -1000, -1000, 388, -1000, -1000, 239, -1000,
	-1000, 220, -1000, -1000, 288, 600, -1000, 283, -1000, -1000,
	-1000, -1000, -1000, -1000, 59, 207, -10, -1000, -1000, -1000,
	-1000, -1000, -1000, 128, 188, 173, -1000, 169, 168, -1000,
	-1000, 151, 110, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 649, -10, -1000, -1000, 482, -1000, -1000, -1000, 280,
	279, -1000, -1000, -1000, -1000, -1000, -1000, 132, -1000, -1000,
	79, -1000, -1000, -1000, -1000, -1000, 270, 401, 49, -1000,
	54, 401, -1000, -1000, 412, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 255, 649, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 448, -1000, -1000, -1000, 248, -1000, 85,
	-1000, -1000, 42, 71, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 13, 227, 371, -1000, 72, -1000,
}
var yyPgo = [...]int{

	0, 648, 36, 1, 11, 8, 647, 646, 642, 637,
	415, 634, 629, 623, 44, 0, 621, 619, 42, 618,
	616, 57, 615, 606, 48, 604, 601, 40, 70, 600,
	61, 45, 599, 598, 597, 56, 54, 49, 35, 30,
	25, 23, 21, 19, 596, 595, 594, 589, 588, 586,
	17, 585, 582, 577, 18, 13, 573, 571, 568, 66,
	567, 12, 566, 565, 43, 24, 14, 564, 563, 562,
	557, 556, 555, 553, 2, 22, 15, 552, 549, 544,
	20, 540, 538, 535, 34, 534, 533, 7, 5, 4,
	532, 531, 530, 529, 523, 53, 59, 41, 522, 520,
	519, 31, 511, 508, 506, 33, 501, 500, 499, 38,
	498, 497, 496, 39, 494, 490, 489, 62, 488, 37,
	486, 9, 10,
}
var yyR1 = [...]int{

	0, 6, 7, 7, 8, 8, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 19, 11, 11, 20, 20,
	21, 21, 22, 23, 23, 18, 24, 24, 24, 24,
	16, 25, 26, 26, 27, 27, 27, 17, 17, 28,
	28, 9, 9, 31, 31, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 32,
	32, 46, 46, 47, 48, 48, 49, 49, 50, 50,
	45, 45, 51, 52, 52, 53, 53, 54, 54, 54,
	55, 41, 57, 57, 57, 57, 56, 58, 58, 59,
	60, 33, 62, 63, 63, 64, 64, 64, 64, 4,
	4, 66, 65, 67, 68, 68, 69, 69, 69, 69,
	69, 71, 71, 36, 72, 73, 73, 61, 61, 74,
	74, 74, 74, 74, 77, 44, 78, 78, 79, 79,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 81, 40, 40, 82, 82, 83, 83, 84,
	84, 84, 86, 87, 87, 87, 87, 87, 87, 87,
	85, 85, 90, 91, 91, 29, 92, 93, 93, 94,
	94, 95, 95, 95, 95, 96, 97, 42, 98, 99,
	99, 100, 100, 101, 101, 101, 101, 43, 102, 103,
	103, 104, 104, 105, 105, 105, 34, 107, 106, 108,
	108, 109, 109, 109, 35, 110, 111, 112, 112, 88,
	88, 89, 113, 113, 113, 113, 113, 113, 113, 113,
	113, 114, 39, 115, 115, 37, 116, 117, 118, 118,
	119, 119, 119, 119, 119, 119, 119, 119, 76, 5,
	5, 3, 2, 2, 75, 38, 120, 70, 70, 121,
	121, 1, 14, 15, 12, 13, 122, 122,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 4, 2, 0, 1, 1, 2, 1, 1,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 4, 0, 1, 1, 1, 2, 1, 2, 4,
	2, 4, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 3, 2, 2, 1, 3, 3, 1, 1, 1,
	3, 1, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 1, 1, 1, 2, 4, 0, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 4, 0, 1, 1, 2, 1,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 4, 1, 1, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 3, 3, 2, 2, 4, 2, 0,
	1, 1, 2, 1, 1, 3, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 2, 3, 2, 1,
	2, 1, 1, 1, 4, 2, 1, 1, 2, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 3, 2, 2, 2, 4, 2, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 4, 2, 1, 2, 3,
	5, 3, 3, 3, 3, 3, 1, 5,
}
var yyChk = [...]int{

	-1000, -6, -7, 26, 55, -8, -10, 12, -11, -12,
	-13, -14, -15, -16, -17, -18, -19, 50, 51, 13,
	35, -22, -25, 16, 14, 41, 42, 4, 4, -9,
	-10, -28, -29, -30, -92, -33, -34, -35, -36, -37,
	-38, -39, -40, -41, -42, -43, -44, -45, -46, 29,
	-62, -106, -110, -72, -116, -120, -115, -81, -56, -98,
	-102, -77, -51, -47, 31, 32, 28, 27, 33, 36,
	44, 45, 22, 39, 43, 30, 54, 57, 59, -5,
	5, 10, 8, -5, -5, -5, -5, 8, 10, 8,
	-5, 11, 4, 4, 8, 8, 9, -28, 8, 4,
	8, -107, 8, 8, 8, 8, 8, 10, 10, 8,
	8, 8, 8, 8, 10, 8, 10, 8, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	-5, 4, 4, 10, 56, -20, -21, -14, -15, 10,
	10, -122, 10, 8, 10, -23, -24, -18, 14, -14,
	-15, -26, -27, 14, -14, -15, 10, -93, -94, -95,
	-14, -15, -96, -97, 24, 25, -63, -64, -65, -14,
	-15, -66, -67, 17, 15, -108, -109, -14, -15, -30,
	-111, -112, -113, -14, -15, -88, -89, -75, -76, -114,
	23, -30, 37, 38, 21, 34, 20, -73, -61, -74,
	-14, -15, -75, -76, -30, -117, -118, -119, -65, -14,
	-15, -75, -88, -89, -76, -66, -117, -82, -83, -84,
	-14, -15, -85, -86, 52, -57, -14, -15, -58, -59,
	-60, 40, -99, -100, -101, -14, -15, -96, -97, -103,
	-104, -105, -14, -15, -30, -78, -79, -80, -14, -15,
	-35, -36, -37, -38, -39, -40, -41, -59, -42, -43,
	-52, -53, -54, -14, -15, -55, 58, -48, -49, -50,
	-14, -15, 5, 9, -21, 7, -24, 9, 11, -27,
	9, 11, 9, -95, -31, -32, -30, -31, 8, 8,
	9, -64, -68, 10, 8, -4, -5, 6, 4, 9,
	-109, 9, -113, -5, -3, 53, 6, -3, -2, 48,
	49, -2, -5, 9, -74, 9, -119, 9, 9, -84,
	10, 8, -5, 9, -59, 8, 4, 9, -101, -31,
	-31, 9, -105, 9, -80, 9, -54, 4, 9, -50,
	-4, 10, 10, 9, -30, 9, -69, 18, -70, -55,
	-71, 46, -121, -65, 19, 10, 10, 10, 10, 10,
	10, 10, 10, -90, -91, -87, -14, -15, -66, -75,
	-76, -88, -89, -61, 9, 9, 10, -122, 9, -5,
	-121, -65, -5, 4, 9, -87, 9, 9, 10, 10,
	10, 8, -1, 47, 9, -3, 10,
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
	239, 16, 0, 0, 0, 0, 0, 0, 37, 0,
	0, 15, 22, 31, 2, 3, 1, 42, 167, 166,
	0, 196, 0, 0, 115, 0, 0, 222, 143, 145,
	82, 179, 189, 126, 70, 73, 61, 64, 92, 198,
	205, 114, 226, 246, 223, 224, 142, 86, 178, 188,
	124, 72, 63, 6, 0, 0, 18, 20, 21, 254,
	255, 252, 256, 0, 253, 0, 23, 26, 0, 28,
	29, 0, 32, 0, 35, 36, 25, 0, 168, 169,
	171, 172, 43, 43, 0, 0, 0, 93, 95, 96,
	97, 98, 0, 0, 0, 0, 199, 201, 202, 203,
	0, 206, 207, 212, 213, 214, 215, 216, 217, 218,
	0, 220, 0, 0, 0, 0, 0, 0, 116, 117,
	119, 120, 121, 122, 123, 0, 227, 228, 230, 231,
	232, 233, 234, 235, 236, 237, 0, 0, 146, 147,
	149, 150, 151, 0, 0, 0, 83, 84, 85, 87,
	0, 0, 0, 180, 181, 183, 184, 43, 43, 0,
	190, 191, 193, 194, 195, 0, 127, 128, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	0, 74, 75, 77, 78, 79, 0, 0, 65, 66,
	68, 69, 240, 17, 19, 0, 24, 30, 0, 33,
	38, 0, 165, 170, 0, 44, 59, 0, 175, 176,
	91, 94, 102, 104, 0, 0, 99, 100, 103, 197,
	200, 204, 208, 0, 0, 0, 241, 0, 0, 242,
	243, 0, 0, 113, 118, 225, 229, 245, 144, 148,
	160, 0, 152, 81, 88, 0, 90, 177, 182, 0,
	0, 187, 192, 125, 129, 71, 76, 0, 62, 67,
	0, 27, 34, 173, 60, 174, 0, 0, 107, 108,
	109, 0, 247, 111, 0, 101, 219, 209, 210, 211,
	244, 238, 221, 0, 162, 163, 153, 154, 155, 156,
	157, 158, 159, 0, 185, 186, 80, 0, 105, 0,
	248, 112, 0, 0, 161, 164, 89, 257, 106, 110,
	249, 0, 0, 0, 250, 0, 251,
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
			yylex.(*lexer).stack.Push(meta.NewModule(yyDollar[2].token))
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
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:313
		{
			pop(yylex)
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:316
		{
			pop(yylex)
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:321
		{
			if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:340
		{
			if set(yylex, meta.SetBase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:350
		{
			pop(yylex)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:361
		{
			if push(yylex, meta.NewChoice(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:373
		{
			pop(yylex)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:378
		{
			if push(yylex, meta.NewChoiceCase(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:388
		{
			pop(yylex)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:393
		{
			if push(yylex, meta.NewTypedef(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:410
		{
			yyVAL.token = yyDollar[1].token
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:411
		{
			yyVAL.token = yyDollar[1].token
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:414
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:424
		{
			if push(yylex, meta.NewDataType(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:431
		{
			pop(yylex)
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:434
		{
			pop(yylex)
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:439
		{
			if set(yylex, meta.SetEncodedLength(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:447
		{
			if set(yylex, meta.SetPath(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:461
		{
			pop(yylex)
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:466
		{
			if push(yylex, meta.NewContainer(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:488
		{
			if push(yylex, meta.NewAugment(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:495
		{
			pop(yylex)
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:521
		{
			if push(yylex, meta.NewUses(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:528
		{
			pop(yylex)
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:531
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:554
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:576
		{
			pop(yylex)
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:579
		{
			pop(yylex)
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:591
		{
			pop(yylex)
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:596
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:612
		{
			pop(yylex)
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:615
		{
			pop(yylex)
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:620
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:627
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 177:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:637
		{
			pop(yylex)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:642
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:658
		{
			pop(yylex)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:661
		{
			pop(yylex)
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:669
		{
			pop(yylex)
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:674
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:695
		{
			pop(yylex)
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:705
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 204:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:723
		{
			pop(yylex)
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:728
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:743
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:748
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:755
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:773
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:780
		{
			pop(yylex)
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:785
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:790
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:800
		{
			pop(yylex)
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:805
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:831
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:838
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:841
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:846
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:856
		{
			yyVAL.boolean = true
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:857
		{
			yyVAL.boolean = false
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:860
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 245:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:870
		{
			pop(yylex)
		}
	case 246:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:875
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 249:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:886
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 250:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:891
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 251:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:898
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 252:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:903
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 253:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:910
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 254:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:917
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 255:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:924
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
