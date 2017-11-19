//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/val"
	"strconv"
)

// blindly chop off quotes
func tokenString(s string) string {
	return s[1 : len(s)-1]
}

// optionally chop off quotes
func tokenPath(s string) string {
	if len(s) > 0 && s[0] == '"' {
		return tokenString(s)
	}
	return s
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

//line parser.y:63
type yySymType struct {
	yys     int
	token   string
	boolean bool
	num     int64
	num32   int
}

const token_ident = 57346
const token_string = 57347
const token_path = 57348
const token_number = 57349
const token_custom = 57350
const token_curly_open = 57351
const token_curly_close = 57352
const token_semi = 57353
const token_rev_ident = 57354
const kywd_namespace = 57355
const kywd_description = 57356
const kywd_revision = 57357
const kywd_type = 57358
const kywd_prefix = 57359
const kywd_default = 57360
const kywd_length = 57361
const kywd_enum = 57362
const kywd_key = 57363
const kywd_config = 57364
const kywd_uses = 57365
const kywd_unique = 57366
const kywd_input = 57367
const kywd_output = 57368
const kywd_module = 57369
const kywd_container = 57370
const kywd_list = 57371
const kywd_rpc = 57372
const kywd_notification = 57373
const kywd_typedef = 57374
const kywd_grouping = 57375
const kywd_leaf = 57376
const kywd_mandatory = 57377
const kywd_reference = 57378
const kywd_leaf_list = 57379
const kywd_max_elements = 57380
const kywd_min_elements = 57381
const kywd_choice = 57382
const kywd_case = 57383
const kywd_import = 57384
const kywd_include = 57385
const kywd_action = 57386
const kywd_anyxml = 57387
const kywd_anydata = 57388
const kywd_path = 57389
const kywd_value = 57390
const kywd_true = 57391
const kywd_false = 57392
const kywd_contact = 57393
const kywd_organization = 57394
const kywd_refine = 57395
const kywd_unbounded = 57396
const kywd_augment = 57397
const kywd_submodule = 57398

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_path",
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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:855

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 622

var yyAct = [...]int{

	187, 185, 277, 12, 199, 198, 12, 330, 318, 127,
	184, 189, 268, 227, 45, 233, 44, 188, 201, 257,
	43, 42, 41, 40, 39, 220, 38, 37, 193, 205,
	149, 148, 168, 162, 15, 281, 138, 122, 132, 153,
	191, 19, 279, 243, 356, 194, 3, 19, 217, 145,
	282, 283, 19, 139, 182, 180, 68, 176, 315, 319,
	19, 63, 62, 20, 71, 60, 61, 64, 181, 20,
	65, 178, 179, 69, 20, 4, 124, 70, 66, 67,
	210, 136, 20, 141, 186, 319, 317, 11, 72, 278,
	11, 19, 147, 160, 156, 159, 164, 170, 19, 196,
	196, 172, 171, 207, 213, 222, 229, 235, 174, 254,
	200, 200, 157, 20, 173, 133, 197, 197, 190, 31,
	20, 245, 124, 244, 33, 217, 251, 242, 241, 240,
	239, 238, 136, 237, 236, 224, 223, 250, 141, 154,
	202, 19, 134, 84, 23, 147, 359, 215, 33, 90,
	352, 253, 354, 156, 353, 19, 139, 351, 327, 247,
	123, 326, 164, 20, 325, 135, 133, 140, 170, 260,
	249, 157, 172, 171, 252, 324, 146, 20, 155, 174,
	163, 169, 280, 195, 195, 173, 287, 206, 212, 221,
	228, 234, 264, 196, 256, 273, 19, 246, 154, 323,
	275, 19, 322, 321, 200, 207, 123, 150, 151, 129,
	197, 128, 320, 310, 165, 177, 135, 284, 20, 309,
	222, 289, 140, 20, 230, 142, 294, 229, 293, 146,
	267, 130, 266, 235, 292, 126, 102, 155, 101, 82,
	305, 81, 125, 302, 303, 301, 163, 245, 307, 244,
	224, 223, 169, 242, 241, 240, 239, 238, 297, 237,
	236, 308, 19, 134, 120, 23, 19, 259, 259, 75,
	159, 74, 100, 357, 180, 263, 350, 195, 347, 19,
	165, 160, 342, 159, 20, 340, 177, 181, 20, 206,
	178, 179, 339, 313, 311, 332, 306, 304, 300, 337,
	336, 20, 296, 291, 221, 290, 335, 288, 286, 338,
	274, 228, 334, 333, 255, 298, 262, 234, 341, 261,
	106, 105, 104, 103, 99, 344, 98, 97, 96, 95,
	332, 93, 91, 88, 337, 336, 87, 348, 80, 248,
	287, 335, 259, 259, 269, 230, 270, 334, 333, 279,
	295, 7, 19, 24, 119, 23, 6, 345, 343, 358,
	285, 68, 30, 276, 83, 79, 63, 62, 47, 71,
	60, 61, 64, 78, 20, 65, 77, 312, 69, 331,
	25, 26, 70, 66, 67, 76, 73, 346, 53, 17,
	18, 299, 271, 72, 349, 118, 117, 116, 19, 115,
	114, 113, 112, 111, 110, 109, 180, 68, 108, 107,
	92, 86, 63, 62, 331, 71, 60, 61, 64, 181,
	20, 65, 85, 19, 69, 28, 27, 192, 70, 66,
	67, 180, 68, 52, 54, 175, 167, 63, 62, 72,
	71, 60, 61, 64, 181, 20, 65, 272, 166, 69,
	50, 19, 161, 70, 66, 67, 94, 49, 226, 225,
	68, 58, 219, 218, 72, 63, 62, 57, 71, 60,
	61, 64, 144, 20, 65, 143, 19, 69, 34, 329,
	328, 70, 66, 67, 209, 68, 208, 204, 203, 55,
	63, 62, 72, 71, 60, 61, 64, 232, 20, 65,
	89, 231, 69, 59, 183, 51, 70, 66, 67, 316,
	314, 265, 158, 68, 152, 48, 216, 72, 63, 62,
	47, 71, 60, 61, 64, 214, 211, 65, 56, 46,
	69, 36, 68, 35, 70, 66, 67, 63, 62, 258,
	71, 60, 61, 64, 32, 72, 65, 137, 22, 69,
	131, 21, 121, 70, 66, 67, 7, 19, 24, 16,
	23, 19, 14, 13, 72, 10, 9, 8, 29, 5,
	68, 2, 1, 355, 0, 63, 62, 0, 71, 20,
	0, 64, 0, 20, 65, 25, 26, 69, 217, 0,
	0, 70, 66, 67, 17, 18, 19, 0, 160, 0,
	159, 0, 0, 0, 180, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 181, 20, 0,
	178, 179,
}
var yyPact = [...]int{

	19, -1000, 543, 422, 421, 338, -1000, 381, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 260, 380, 371, 368,
	360, 329, 230, 359, 131, 418, 407, 327, 324, 490,
	-1000, -1000, -1000, -1000, 323, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 406, 322, 320,
	319, 318, 317, 315, 261, 227, 314, 313, 312, 311,
	405, 404, 401, 400, 399, 398, 397, 396, 395, 393,
	392, 391, 348, 253, -1000, 46, 231, 224, 200, 220,
	248, -1000, 38, 214, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 182, -1000, 77, -1000, 462, 33, 409, 582, 582,
	-1000, -1000, 27, 84, 182, 462, 547, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 187, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 331,
	-1000, 127, -1000, -1000, 114, -1000, -1000, 141, -1000, 97,
	-1000, -1000, -1000, 304, 182, -1000, -1000, -1000, 509, 509,
	310, 307, 265, -1000, -1000, -1000, -1000, -1000, 221, 339,
	388, 437, -1000, -1000, -1000, -1000, 300, 33, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 358, -1000, 35, 342,
	1, 1, 355, 298, 409, -1000, -1000, -1000, -1000, -1000,
	-1000, 297, 582, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 295, 293, 27, -1000, -1000, -1000, -1000, 217,
	344, 292, -1000, -1000, 7, -1000, 306, 387, 288, 182,
	-1000, -1000, -1000, 509, 509, 287, 462, -1000, -1000, -1000,
	-1000, 286, 547, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 339, -1000,
	-1000, 208, -1000, -1000, 202, -1000, -1000, 284, 509, -1000,
	283, -1000, -1000, -1000, -1000, -1000, -1000, 39, 201, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 192, 191, 188, -1000,
	164, 153, -1000, -1000, 150, 147, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 252, -1000, -1000, -1000, 409, -1000,
	-1000, -1000, 282, 275, -1000, -1000, -1000, -1000, 200, -1000,
	-1000, -1000, -1000, -1000, 272, 353, 65, 352, -1000, 383,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 268, 252,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 384, -1000,
	-1000, 266, -1000, 146, -1000, 139, 143, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -4, 263, 342, -1000, 135, -1000,
}
var yyPgo = [...]int{

	0, 573, 35, 2, 12, 572, 571, 569, 568, 356,
	567, 566, 565, 84, 0, 563, 562, 34, 559, 552,
	37, 551, 550, 38, 548, 547, 36, 119, 544, 118,
	19, 539, 533, 531, 27, 26, 24, 23, 22, 21,
	20, 16, 14, 529, 528, 526, 525, 43, 516, 10,
	515, 514, 39, 45, 18, 512, 511, 510, 509, 505,
	504, 1, 17, 11, 503, 501, 497, 15, 489, 488,
	487, 29, 486, 484, 7, 5, 4, 480, 479, 478,
	475, 472, 49, 31, 30, 467, 463, 462, 25, 461,
	459, 458, 13, 457, 456, 452, 33, 450, 448, 436,
	32, 435, 434, 433, 40, 427, 28, 388, 8, 9,
}
var yyR1 = [...]int{

	0, 5, 6, 6, 7, 7, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 18, 10, 10, 19, 19,
	20, 20, 21, 22, 22, 17, 23, 23, 23, 23,
	15, 24, 25, 25, 26, 26, 26, 16, 16, 27,
	27, 8, 8, 30, 30, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 31, 31, 40,
	45, 45, 45, 45, 44, 46, 46, 47, 48, 32,
	50, 51, 51, 52, 52, 52, 52, 4, 4, 54,
	53, 55, 56, 56, 57, 57, 57, 35, 59, 60,
	60, 49, 49, 61, 61, 61, 61, 61, 64, 43,
	65, 65, 66, 66, 67, 67, 67, 67, 67, 67,
	67, 67, 67, 67, 67, 67, 68, 39, 39, 69,
	69, 70, 70, 71, 71, 71, 73, 74, 74, 74,
	74, 74, 74, 74, 72, 72, 77, 78, 78, 28,
	79, 80, 80, 81, 81, 82, 82, 82, 82, 83,
	84, 41, 85, 86, 86, 87, 87, 88, 88, 88,
	88, 42, 89, 90, 90, 91, 91, 92, 92, 92,
	33, 94, 93, 95, 95, 96, 96, 96, 34, 97,
	98, 99, 99, 75, 75, 76, 100, 100, 100, 100,
	100, 100, 100, 100, 100, 101, 38, 102, 102, 36,
	103, 104, 105, 105, 106, 106, 106, 106, 106, 106,
	106, 106, 63, 3, 2, 2, 62, 37, 107, 58,
	58, 108, 108, 1, 13, 14, 11, 12, 109, 109,
}
var yyR2 = [...]int{

	0, 4, 3, 3, 1, 2, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 4, 1, 2,
	1, 1, 2, 1, 2, 3, 1, 3, 1, 1,
	4, 2, 1, 2, 3, 1, 1, 2, 4, 1,
	1, 1, 2, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	0, 1, 1, 1, 2, 1, 2, 4, 2, 4,
	2, 1, 2, 1, 1, 1, 1, 1, 1, 3,
	2, 2, 1, 3, 3, 1, 3, 4, 2, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 2, 4,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 4, 0,
	1, 1, 2, 1, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 2, 4,
	2, 0, 1, 1, 2, 1, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 3,
	3, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	2, 3, 2, 1, 2, 1, 1, 1, 4, 2,
	1, 1, 2, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3, 2, 2, 2, 4,
	2, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 1, 3, 4, 2, 1,
	2, 3, 5, 3, 3, 3, 3, 3, 1, 5,
}
var yyChk = [...]int{

	-1000, -5, -6, 27, 56, -7, -9, 13, -10, -11,
	-12, -13, -14, -15, -16, -17, -18, 51, 52, 14,
	36, -21, -24, 17, 15, 42, 43, 4, 4, -8,
	-9, -27, -28, -29, -79, -32, -33, -34, -35, -36,
	-37, -38, -39, -40, -41, -42, -43, 30, -50, -93,
	-97, -59, -103, -107, -102, -68, -44, -85, -89, -64,
	32, 33, 29, 28, 34, 37, 45, 46, 23, 40,
	44, 31, 55, 5, 11, 9, 5, 5, 5, 5,
	9, 11, 9, 5, 12, 4, 4, 9, 9, 10,
	-27, 9, 4, 9, -94, 9, 9, 9, 9, 9,
	11, 11, 9, 9, 9, 9, 9, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 6,
	11, -19, -20, -13, -14, 11, 11, -109, 11, 9,
	11, -22, -23, -17, 15, -13, -14, -25, -26, 15,
	-13, -14, 11, -80, -81, -82, -13, -14, -83, -84,
	25, 26, -51, -52, -53, -13, -14, -54, -55, 18,
	16, -95, -96, -13, -14, -29, -98, -99, -100, -13,
	-14, -75, -76, -62, -63, -101, 24, -29, 38, 39,
	22, 35, 21, -60, -49, -61, -13, -14, -62, -63,
	-29, -104, -105, -106, -53, -13, -14, -62, -75, -76,
	-63, -54, -104, -69, -70, -71, -13, -14, -72, -73,
	53, -45, -13, -14, -46, -47, -48, 41, -86, -87,
	-88, -13, -14, -83, -84, -90, -91, -92, -13, -14,
	-29, -65, -66, -67, -13, -14, -34, -35, -36, -37,
	-38, -39, -40, -47, -41, -42, 10, -20, 8, -23,
	10, 12, -26, 10, 12, 10, -82, -30, -31, -29,
	-30, 9, 9, 10, -52, -56, 11, 9, -4, 5,
	7, 4, 10, -96, 10, -100, 5, -3, 54, 7,
	-3, -2, 49, 50, -2, 5, 10, -61, 10, -106,
	10, 10, -71, 11, 9, 6, 10, -47, 9, 4,
	10, -88, -30, -30, 10, -92, 10, -67, -4, 11,
	11, 10, -29, 10, -57, 19, -58, 47, -108, 20,
	11, 11, 11, 11, 11, 11, 11, 11, -77, -78,
	-74, -13, -14, -54, -62, -63, -75, -76, -49, 10,
	10, -109, 10, 5, -108, 5, 4, 10, -74, 10,
	10, 11, 11, 11, 9, -1, 48, 10, -3, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 4, 0, 7, 8,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	5, 41, 39, 40, 0, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 16, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 15, 22, 31, 2, 3, 1,
	42, 141, 140, 0, 170, 0, 0, 89, 0, 0,
	196, 117, 119, 60, 153, 163, 100, 70, 172, 179,
	88, 200, 218, 197, 198, 116, 64, 152, 162, 98,
	6, 0, 18, 20, 21, 226, 227, 224, 228, 0,
	225, 0, 23, 26, 0, 28, 29, 0, 32, 0,
	35, 36, 25, 0, 142, 143, 145, 146, 43, 43,
	0, 0, 0, 71, 73, 74, 75, 76, 0, 0,
	0, 0, 173, 175, 176, 177, 0, 180, 181, 186,
	187, 188, 189, 190, 191, 192, 0, 194, 0, 0,
	0, 0, 0, 0, 90, 91, 93, 94, 95, 96,
	97, 0, 201, 202, 204, 205, 206, 207, 208, 209,
	210, 211, 0, 0, 120, 121, 123, 124, 125, 0,
	0, 0, 61, 62, 63, 65, 0, 0, 0, 154,
	155, 157, 158, 43, 43, 0, 164, 165, 167, 168,
	169, 0, 101, 102, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 114, 115, 17, 19, 0, 24,
	30, 0, 33, 38, 0, 139, 144, 0, 44, 57,
	0, 149, 150, 69, 72, 80, 82, 0, 0, 77,
	78, 81, 171, 174, 178, 182, 0, 0, 0, 213,
	0, 0, 214, 215, 0, 0, 87, 92, 199, 203,
	217, 118, 122, 134, 0, 126, 59, 66, 0, 68,
	151, 156, 0, 0, 161, 166, 99, 103, 0, 27,
	34, 147, 58, 148, 0, 0, 85, 0, 219, 0,
	79, 193, 183, 184, 185, 216, 212, 195, 0, 136,
	137, 127, 128, 129, 130, 131, 132, 133, 0, 159,
	160, 0, 83, 0, 220, 0, 0, 135, 138, 67,
	229, 84, 86, 221, 0, 0, 0, 222, 0, 223,
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
	52, 53, 54, 55, 56,
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
		//line parser.y:141
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
		//line parser.y:149
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
		//line parser.y:163
		{
			if set(yylex, meta.SetNamespace(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:178
		{
			if push(yylex, meta.NewRevision(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:185
		{
			pop(yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:188
		{
			pop(yylex)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:201
		{
			if push(yylex, meta.NewImport(yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:212
		{
			if set(yylex, meta.SetPrefix(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:225
		{
			pop(yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:230
		{
			if push(yylex, meta.NewInclude(yyDollar[2].token, yylex.(*lexer).loader)) {
				goto ret1
			}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:246
		{
			pop(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:249
		{
			pop(yylex)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:286
		{
			pop(yylex)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:297
		{
			if push(yylex, meta.NewChoice(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:309
		{
			pop(yylex)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:314
		{
			if push(yylex, meta.NewChoiceCase(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:324
		{
			pop(yylex)
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:329
		{
			if push(yylex, meta.NewTypedef(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:346
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:347
		{
			yyVAL.token = yyDollar[1].token
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:350
		{
			if set(yylex, meta.SetDefault{Value: yyDollar[2].token}) {
				goto ret1
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:360
		{
			if push(yylex, meta.NewDataType(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:367
		{
			pop(yylex)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:370
		{
			pop(yylex)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:375
		{
			if set(yylex, meta.SetEncodedLength(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:381
		{
			if set(yylex, meta.SetPath(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:391
		{
			pop(yylex)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:396
		{
			if push(yylex, meta.NewContainer(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:418
		{
			if push(yylex, meta.NewAugment(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:425
		{
			pop(yylex)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:451
		{
			if push(yylex, meta.NewUses(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:458
		{
			pop(yylex)
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:461
		{
			pop(yylex)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:484
		{
			if push(yylex, meta.NewRefine(tokenPath(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:506
		{
			pop(yylex)
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:509
		{
			pop(yylex)
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:521
		{
			pop(yylex)
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:526
		{
			if push(yylex, meta.NewRpc(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:542
		{
			pop(yylex)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:545
		{
			pop(yylex)
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:550
		{
			if push(yylex, meta.NewRpcInput()) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:557
		{
			if push(yylex, meta.NewRpcOutput()) {
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:567
		{
			pop(yylex)
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:572
		{
			if push(yylex, meta.NewRpc(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:588
		{
			pop(yylex)
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:591
		{
			pop(yylex)
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:599
		{
			pop(yylex)
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:604
		{
			if push(yylex, meta.NewNotification(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:625
		{
			pop(yylex)
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:635
		{
			if push(yylex, meta.NewGrouping(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:653
		{
			pop(yylex)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:658
		{
			if push(yylex, meta.NewList(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:673
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:678
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:685
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:703
		{
			if set(yylex, meta.SetKey(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:710
		{
			pop(yylex)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:715
		{
			if push(yylex, meta.NewAny(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:720
		{
			if push(yylex, meta.NewAny(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:730
		{
			pop(yylex)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:735
		{
			if push(yylex, meta.NewLeaf(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:761
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:768
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:778
		{
			yyVAL.boolean = true
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:779
		{
			yyVAL.boolean = false
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:782
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:792
		{
			pop(yylex)
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:797
		{
			if push(yylex, meta.NewLeafList(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:808
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 222:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:813
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:820
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:825
		{
			if set(yylex, meta.SetDescription(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:832
		{
			if set(yylex, meta.SetReference(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:839
		{
			if set(yylex, meta.SetContact(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:846
		{
			if set(yylex, meta.SetOrganization(tokenString(yyDollar[2].token))) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
