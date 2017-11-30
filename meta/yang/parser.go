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

//line parser.y:974

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 847

var yyAct = [...]int{

	207, 333, 397, 12, 205, 383, 12, 225, 142, 204,
	290, 323, 287, 295, 263, 270, 45, 223, 215, 222,
	44, 165, 311, 255, 43, 216, 236, 337, 185, 179,
	170, 137, 160, 282, 15, 147, 153, 206, 19, 135,
	11, 426, 335, 11, 19, 252, 42, 41, 213, 40,
	39, 38, 200, 324, 164, 177, 37, 19, 378, 385,
	20, 79, 144, 31, 143, 201, 20, 385, 3, 211,
	422, 83, 84, 85, 86, 338, 339, 90, 421, 20,
	19, 177, 177, 139, 176, 166, 382, 210, 151, 334,
	156, 166, 197, 97, 308, 305, 394, 4, 292, 162,
	91, 173, 20, 181, 187, 429, 218, 218, 174, 229,
	135, 238, 246, 257, 265, 272, 135, 289, 409, 297,
	138, 191, 148, 190, 135, 150, 171, 155, 208, 393,
	131, 284, 392, 391, 260, 283, 161, 139, 172, 281,
	180, 186, 135, 217, 217, 250, 228, 151, 237, 245,
	256, 264, 271, 156, 288, 226, 296, 212, 19, 154,
	162, 280, 279, 33, 278, 277, 276, 259, 301, 390,
	173, 275, 389, 193, 138, 224, 224, 174, 233, 181,
	20, 148, 303, 19, 150, 187, 386, 33, 314, 306,
	155, 192, 310, 221, 221, 171, 232, 161, 372, 19,
	319, 336, 191, 427, 190, 20, 19, 172, 328, 343,
	252, 371, 388, 330, 420, 218, 180, 167, 168, 387,
	315, 20, 186, 376, 424, 355, 423, 163, 20, 340,
	166, 197, 188, 345, 219, 219, 238, 230, 417, 239,
	248, 258, 266, 273, 292, 291, 166, 298, 144, 331,
	143, 332, 217, 166, 193, 257, 341, 19, 135, 411,
	182, 196, 349, 265, 157, 135, 209, 408, 407, 135,
	272, 267, 192, 237, 375, 373, 260, 362, 358, 20,
	368, 354, 359, 360, 224, 364, 284, 289, 163, 145,
	283, 141, 256, 140, 281, 297, 243, 352, 134, 366,
	264, 351, 221, 350, 166, 197, 365, 271, 369, 259,
	135, 363, 361, 188, 370, 357, 280, 279, 353, 278,
	277, 276, 313, 313, 288, 307, 275, 300, 348, 19,
	154, 19, 296, 380, 347, 135, 182, 135, 322, 135,
	321, 318, 196, 219, 135, 19, 346, 177, 384, 176,
	304, 20, 399, 20, 19, 149, 118, 23, 117, 401,
	19, 149, 344, 23, 239, 406, 116, 20, 115, 405,
	189, 404, 220, 220, 342, 231, 20, 240, 249, 410,
	317, 274, 20, 258, 19, 413, 329, 309, 176, 398,
	316, 266, 200, 110, 108, 109, 107, 399, 273, 418,
	89, 82, 88, 81, 401, 201, 20, 414, 198, 199,
	114, 343, 113, 112, 405, 291, 404, 313, 313, 111,
	267, 403, 106, 298, 105, 104, 19, 103, 428, 102,
	100, 166, 412, 98, 398, 72, 415, 95, 94, 402,
	67, 66, 87, 75, 80, 325, 68, 302, 20, 69,
	335, 189, 73, 252, 6, 80, 74, 70, 71, 299,
	30, 416, 367, 356, 326, 133, 403, 132, 19, 130,
	374, 129, 128, 166, 197, 202, 200, 72, 195, 127,
	400, 220, 67, 66, 402, 75, 64, 65, 68, 201,
	20, 69, 198, 199, 73, 126, 125, 124, 74, 70,
	71, 123, 240, 122, 121, 120, 119, 99, 93, 76,
	92, 28, 77, 27, 78, 166, 197, 55, 214, 54,
	227, 56, 7, 19, 24, 400, 23, 7, 19, 24,
	194, 23, 72, 184, 183, 52, 274, 67, 66, 49,
	75, 64, 65, 68, 178, 20, 69, 101, 51, 73,
	20, 25, 26, 74, 70, 71, 25, 26, 419, 262,
	17, 18, 19, 261, 76, 17, 18, 77, 60, 78,
	200, 72, 254, 253, 59, 159, 67, 66, 158, 75,
	64, 65, 68, 201, 20, 69, 34, 396, 73, 395,
	242, 241, 74, 70, 71, 235, 234, 57, 269, 268,
	61, 203, 53, 76, 19, 381, 77, 379, 78, 166,
	197, 377, 200, 72, 320, 175, 169, 50, 67, 66,
	251, 75, 64, 65, 68, 201, 20, 69, 247, 244,
	73, 58, 286, 285, 74, 70, 71, 62, 294, 293,
	19, 63, 48, 47, 46, 76, 36, 35, 77, 72,
	78, 166, 197, 312, 67, 66, 32, 75, 64, 65,
	68, 152, 20, 69, 22, 146, 73, 21, 136, 16,
	74, 70, 71, 14, 13, 327, 10, 9, 8, 19,
	29, 76, 5, 2, 77, 1, 78, 166, 72, 425,
	0, 0, 0, 67, 66, 0, 75, 64, 65, 68,
	0, 20, 69, 0, 0, 73, 0, 0, 0, 74,
	70, 71, 19, 0, 0, 0, 0, 0, 0, 0,
	76, 72, 0, 77, 0, 78, 67, 66, 0, 75,
	64, 65, 68, 0, 20, 69, 0, 0, 73, 0,
	0, 96, 74, 70, 71, 0, 0, 0, 0, 0,
	0, 0, 0, 76, 72, 0, 77, 0, 78, 67,
	66, 49, 75, 64, 65, 68, 0, 0, 69, 0,
	0, 73, 0, 72, 0, 74, 70, 71, 67, 66,
	0, 75, 64, 65, 68, 0, 76, 69, 0, 77,
	73, 78, 0, 0, 74, 70, 71, 0, 19, 0,
	177, 0, 176, 0, 0, 76, 200, 0, 77, 0,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 201,
	20, 0, 198, 199, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 166, 197,
}
var yyPact = [...]int{

	42, -1000, 515, 509, 507, 510, -1000, 450, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 393, 450, 450, 450,
	450, 434, 392, 450, 89, 506, 504, 430, 429, 732,
	-1000, -1000, -1000, -1000, 425, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 503,
	422, 421, 419, 417, 416, 414, 386, 385, 411, 405,
	404, 402, 358, 348, 502, 501, 500, 499, 497, 493,
	492, 491, 475, 468, 467, 465, 450, 463, 461, 288,
	-1000, -1000, 44, 283, 281, 54, 279, 347, -1000, 145,
	254, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 193, -1000,
	67, -1000, 699, 455, 591, 785, 785, -1000, 31, -1000,
	244, 170, 193, 627, 413, -1000, 186, -1000, 25, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -17, -1000, -1000, -1000, 454, 318, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 440, -1000, 341, -1000, -1000, 84,
	-1000, -1000, 316, -1000, 83, -1000, -1000, -1000, 378, 193,
	-1000, -1000, -1000, -1000, 751, 751, 450, 382, 372, 332,
	-1000, -1000, -1000, -1000, -1000, 330, 439, 460, 666, -1000,
	-1000, -1000, -1000, 377, 455, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 450, -1000, 450, 36, 444,
	27, 27, 450, 365, 591, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 353, 785, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 337, 325, -1000, -1000,
	-1000, -1000, -1000, -1000, 319, 244, -1000, -1000, -1000, -1000,
	-1000, -1000, 293, 450, 309, -1000, -1000, 5, -1000, -1000,
	-1000, 217, 459, 306, 193, -1000, -1000, -1000, -1000, 751,
	751, 303, 627, -1000, -1000, -1000, -1000, -1000, 302, 413,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 297, 186, -1000, -1000, -1000,
	-1000, -1000, 458, 271, 25, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 439, -1000, -1000, 201, -1000, -1000, 188, -1000,
	-1000, 266, 751, -1000, 265, 213, -1000, -1000, -1000, -1000,
	-1000, -1000, 40, 176, -17, -1000, -1000, -1000, -1000, -1000,
	-1000, 209, 202, 162, 159, -1000, 123, 122, -1000, -1000,
	119, 86, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 371, -17, -1000, -1000, 591, -1000, -1000, -1000, 259,
	258, -1000, -1000, -1000, -1000, -1000, -1000, 108, -1000, -1000,
	240, -1000, -1000, -1000, -1000, -1000, -1000, 250, 450, 48,
	-1000, 66, 450, -1000, -1000, 457, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 229, 371, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 549, -1000, -1000, -1000,
	205, -1000, 68, -1000, -1000, 60, 216, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -6, 194, 444, -1000, 95, -1000,
}
var yyPgo = [...]int{

	0, 689, 27, 1, 11, 53, 685, 683, 682, 680,
	454, 678, 677, 676, 37, 0, 674, 673, 34, 669,
	668, 31, 667, 665, 35, 664, 661, 36, 63, 656,
	157, 22, 653, 647, 646, 56, 51, 50, 49, 47,
	46, 24, 20, 16, 644, 643, 642, 641, 639, 638,
	13, 128, 266, 637, 633, 632, 12, 10, 631, 629,
	628, 33, 620, 9, 617, 616, 30, 25, 7, 615,
	614, 611, 607, 605, 602, 601, 4, 87, 69, 600,
	599, 598, 15, 597, 596, 595, 26, 591, 590, 2,
	19, 17, 589, 587, 586, 578, 575, 32, 54, 21,
	574, 573, 572, 23, 568, 563, 559, 14, 548, 547,
	544, 29, 535, 534, 533, 28, 530, 521, 520, 519,
	48, 518, 18, 517, 5, 8,
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
	86, 86, 86, 86, 88, 89, 89, 89, 89, 89,
	89, 89, 89, 87, 87, 92, 93, 93, 29, 94,
	95, 95, 96, 96, 97, 97, 97, 97, 97, 98,
	99, 42, 100, 101, 101, 102, 102, 103, 103, 103,
	103, 103, 43, 104, 105, 105, 106, 106, 107, 107,
	107, 107, 34, 109, 108, 110, 110, 111, 111, 111,
	35, 112, 113, 114, 114, 90, 90, 91, 115, 115,
	115, 115, 115, 115, 115, 115, 115, 115, 115, 116,
	39, 39, 118, 118, 118, 118, 118, 118, 117, 117,
	37, 119, 120, 121, 121, 122, 122, 122, 122, 122,
	122, 122, 122, 122, 122, 78, 5, 5, 3, 2,
	2, 77, 38, 123, 72, 72, 124, 124, 1, 14,
	15, 12, 13, 125, 125,
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
	1, 1, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 4, 1, 1, 2, 4, 2,
	0, 1, 1, 2, 1, 1, 1, 3, 3, 2,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 1,
	3, 3, 4, 2, 0, 1, 1, 2, 1, 1,
	1, 1, 2, 3, 2, 1, 2, 1, 1, 1,
	4, 2, 1, 1, 2, 3, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	2, 4, 1, 1, 1, 1, 1, 1, 2, 2,
	4, 2, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 4, 2, 1, 2, 3, 5, 3, 3,
	3, 3, 3, 1, 5,
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
	-52, -87, -88, 52, -59, -14, -15, -60, -51, -52,
	-61, -62, 40, -101, -102, -103, -14, -15, -51, -98,
	-99, -105, -106, -107, -14, -15, -51, -30, -80, -81,
	-82, -14, -15, -51, -52, -35, -36, -37, -38, -39,
	-40, -41, -61, -42, -43, -54, -55, -56, -14, -15,
	-57, -51, 58, -48, -49, -50, -14, -15, -51, 5,
	9, -21, 7, -24, 9, 11, -27, 9, 11, 9,
	-97, -31, -32, -30, -31, -5, 8, 8, 9, -66,
	-70, 10, 8, -4, -5, 6, 4, 9, -111, 9,
	-115, -5, -5, -3, 53, 6, -3, -2, 48, 49,
	-2, -5, 9, -76, 9, -122, 9, 9, 9, -86,
	10, 8, -5, 9, -61, 8, 4, 9, -103, -31,
	-31, 9, -107, 9, -82, 9, -56, 4, 9, -50,
	-4, 10, 10, 9, -30, 9, 10, -71, 18, -72,
	-57, -73, 46, -124, -67, 19, 10, 10, 10, 10,
	10, 10, 10, 10, 10, -92, -93, -89, -14, -15,
	-51, -68, -77, -78, -90, -91, -63, 9, 9, 10,
	-125, 9, -5, -124, -67, -5, 4, 9, -89, 9,
	9, 10, 10, 10, 8, -1, 47, 9, -3, 10,
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
	266, 16, 0, 0, 0, 0, 0, 0, 37, 0,
	0, 15, 22, 31, 2, 3, 1, 42, 180, 179,
	0, 212, 0, 0, 121, 0, 0, 240, 0, 153,
	155, 86, 193, 204, 134, 73, 76, 61, 64, 98,
	214, 221, 120, 251, 273, 248, 249, 152, 92, 192,
	203, 132, 75, 63, 6, 0, 0, 18, 20, 21,
	281, 282, 279, 283, 0, 280, 0, 23, 26, 0,
	28, 29, 0, 32, 0, 35, 36, 25, 0, 181,
	182, 184, 185, 186, 43, 43, 0, 0, 0, 0,
	99, 101, 102, 103, 104, 0, 0, 0, 0, 215,
	217, 218, 219, 0, 222, 223, 228, 229, 230, 231,
	232, 233, 234, 235, 236, 0, 238, 0, 0, 0,
	0, 0, 0, 0, 122, 123, 125, 126, 127, 128,
	129, 130, 131, 0, 252, 253, 255, 256, 257, 258,
	259, 260, 261, 262, 263, 264, 0, 0, 242, 243,
	244, 245, 246, 247, 0, 156, 157, 159, 160, 161,
	162, 163, 0, 0, 0, 87, 88, 89, 90, 91,
	93, 0, 0, 0, 194, 195, 197, 198, 199, 43,
	43, 0, 205, 206, 208, 209, 210, 211, 0, 135,
	136, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 0, 77, 78, 80, 81,
	82, 83, 0, 0, 65, 66, 68, 69, 70, 267,
	17, 19, 0, 24, 30, 0, 33, 38, 0, 178,
	183, 0, 44, 59, 0, 0, 189, 190, 97, 100,
	108, 110, 0, 0, 105, 106, 109, 213, 216, 220,
	224, 0, 0, 0, 0, 268, 0, 0, 269, 270,
	0, 0, 119, 124, 250, 254, 272, 241, 154, 158,
	173, 0, 164, 85, 94, 0, 96, 191, 196, 0,
	0, 202, 207, 133, 137, 74, 79, 0, 62, 67,
	0, 27, 34, 187, 60, 188, 71, 0, 0, 113,
	114, 115, 0, 274, 117, 0, 107, 237, 72, 225,
	226, 227, 271, 265, 239, 0, 175, 176, 165, 166,
	167, 168, 169, 170, 171, 172, 0, 200, 201, 84,
	0, 111, 0, 275, 118, 0, 0, 174, 177, 95,
	284, 112, 116, 276, 0, 0, 0, 277, 0, 278,
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
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:576
		{
			if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:599
		{
			pop(yylex)
		}
	case 174:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:602
		{
			pop(yylex)
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:614
		{
			pop(yylex)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:619
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:636
		{
			pop(yylex)
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:639
		{
			pop(yylex)
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:644
		{
			if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:651
		{
			if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
				goto ret1
			}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:661
		{
			pop(yylex)
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:666
		{
			if push(yylex, meta.NewRpc(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:683
		{
			pop(yylex)
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:686
		{
			pop(yylex)
		}
	case 202:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:694
		{
			pop(yylex)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:699
		{
			if push(yylex, meta.NewNotification(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:721
		{
			pop(yylex)
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:731
		{
			if push(yylex, meta.NewGrouping(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 220:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:749
		{
			pop(yylex)
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:754
		{
			if push(yylex, meta.NewList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:769
		{
			if set(yylex, meta.SetMaxElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:774
		{
			if set(yylex, meta.SetUnbounded(true)) {
				goto ret1
			}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:781
		{
			if set(yylex, meta.SetMinElements(yyDollar[2].num32)) {
				goto ret1
			}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:801
		{
			if set(yylex, meta.SetKey(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 240:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:808
		{
			pop(yylex)
		}
	case 241:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:811
		{
			pop(yylex)
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:825
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:830
		{
			if push(yylex, meta.NewAny(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 250:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:840
		{
			pop(yylex)
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:845
		{
			if push(yylex, meta.NewLeaf(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 265:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:872
		{
			if set(yylex, meta.SetMandatory(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:879
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 267:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:882
		{
			yyVAL.token = yyDollar[1].token + tokenString(yyDollar[3].token)
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:887
		{
			n, err := strconv.ParseInt(yyDollar[1].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[1].token))
				goto ret1
			}
			yyVAL.num32 = int(n)
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:897
		{
			yyVAL.boolean = true
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:898
		{
			yyVAL.boolean = false
		}
	case 271:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:901
		{
			if set(yylex, meta.SetConfig(yyDollar[2].boolean)) {
				goto ret1
			}
		}
	case 272:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:911
		{
			pop(yylex)
		}
	case 273:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:916
		{
			if push(yylex, meta.NewLeafList(peek(yylex), yyDollar[2].token)) {
				goto ret1
			}
		}
	case 276:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:927
		{
			if set(yylex, meta.SetEnumLabel(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 277:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:932
		{
			if set(yylex, val.Enum{Label: yyDollar[2].token, Id: yyDollar[4].num32}) {
				goto ret1
			}
		}
	case 278:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:939
		{
			yyVAL.num32 = yyDollar[2].num32
		}
	case 279:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:944
		{
			if set(yylex, meta.SetDescription(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 280:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:951
		{
			if set(yylex, meta.SetReference(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 281:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:958
		{
			if set(yylex, meta.SetContact(yyDollar[2].token)) {
				goto ret1
			}
		}
	case 282:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:965
		{
			if set(yylex, meta.SetOrganization(yyDollar[2].token)) {
				goto ret1
			}
		}
	}
	goto yystack /* stack new state and value */
}
