//line parser.y:2
package yang

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"github.com/c2stack/c2g/meta"
	"strconv"
	"strings"
)

type yangError struct {
	s string
}

func (err *yangError) Error() string {
	return err.s
}

func tokenString(s string) string {
	return s[1 : len(s)-1]
}

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
	lval.stack = l.stack
	lval.loader = l.loader
	return int(t.typ)
}

func (l *lexer) Error(e string) {
	line, col := l.Position()
	msg := fmt.Sprintf("%s - line %d, col %d", e, line, col)
	l.lastError = &yangError{msg}
}

func HasError(l yyLexer, e error) bool {
	if e == nil {
		return false
	}
	l.Error(e.Error())
	return true
}

func popAndAddMeta(yylval *yySymType) error {
	i := yylval.stack.Pop()
	if def, ok := i.(meta.Meta); ok {
		parent := yylval.stack.Peek()
		if parentList, ok := parent.(meta.MetaList); ok {
			return parentList.AddMeta(def)
		} else {
			return &yangError{fmt.Sprintf("Cannot add \"%s\" to \"%s\"; not collection type.", i.GetIdent(), parent.GetIdent())}
		}
	} else {
		return &yangError{fmt.Sprintf("\"%s\" cannot be stored in a collection type.", i.GetIdent())}
	}
}

//line parser.y:71
type yySymType struct {
	yys     int
	token   string
	boolean bool
	num     int64
	stack   *yangMetaStack
	loader  ModuleLoader
}

const token_ident = 57346
const token_string = 57347
const token_path = 57348
const token_int = 57349
const token_number = 57350
const token_custom = 57351
const token_curly_open = 57352
const token_curly_close = 57353
const token_semi = 57354
const token_rev_ident = 57355
const kywd_namespace = 57356
const kywd_description = 57357
const kywd_revision = 57358
const kywd_type = 57359
const kywd_prefix = 57360
const kywd_default = 57361
const kywd_length = 57362
const kywd_enum = 57363
const kywd_key = 57364
const kywd_config = 57365
const kywd_uses = 57366
const kywd_unique = 57367
const kywd_input = 57368
const kywd_output = 57369
const kywd_module = 57370
const kywd_container = 57371
const kywd_list = 57372
const kywd_rpc = 57373
const kywd_notification = 57374
const kywd_typedef = 57375
const kywd_grouping = 57376
const kywd_leaf = 57377
const kywd_mandatory = 57378
const kywd_reference = 57379
const kywd_leaf_list = 57380
const kywd_max_elements = 57381
const kywd_min_elements = 57382
const kywd_choice = 57383
const kywd_case = 57384
const kywd_import = 57385
const kywd_include = 57386
const kywd_action = 57387
const kywd_anyxml = 57388
const kywd_anydata = 57389
const kywd_path = 57390
const kywd_value = 57391
const kywd_true = 57392
const kywd_false = 57393
const kywd_contact = 57394
const kywd_organization = 57395
const kywd_refine = 57396
const kywd_unbounded = 57397
const kywd_augment = 57398

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"token_ident",
	"token_string",
	"token_path",
	"token_int",
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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:923

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 599

var yyAct = [...]int{

	186, 198, 184, 11, 197, 11, 322, 155, 334, 124,
	183, 271, 42, 188, 228, 234, 259, 41, 187, 219,
	204, 40, 192, 159, 39, 38, 283, 200, 149, 189,
	37, 165, 135, 119, 30, 36, 18, 279, 360, 35,
	193, 129, 274, 141, 185, 34, 18, 10, 216, 10,
	3, 244, 284, 285, 190, 65, 30, 323, 19, 28,
	60, 59, 18, 68, 57, 58, 61, 256, 19, 62,
	319, 323, 66, 253, 121, 209, 67, 63, 64, 133,
	358, 138, 357, 363, 19, 280, 86, 69, 143, 216,
	152, 81, 161, 167, 169, 195, 195, 168, 321, 206,
	212, 221, 230, 236, 18, 131, 171, 130, 199, 199,
	126, 170, 125, 196, 196, 246, 356, 153, 120, 121,
	245, 162, 174, 132, 243, 137, 19, 242, 241, 133,
	150, 231, 142, 240, 151, 138, 160, 166, 239, 194,
	194, 143, 238, 205, 211, 220, 229, 235, 237, 152,
	201, 214, 248, 355, 247, 295, 331, 294, 18, 161,
	330, 329, 262, 120, 328, 167, 169, 254, 18, 168,
	250, 327, 326, 132, 261, 261, 153, 266, 171, 137,
	19, 325, 275, 170, 258, 142, 288, 324, 162, 150,
	19, 18, 195, 151, 174, 157, 277, 314, 18, 180,
	156, 313, 157, 160, 206, 199, 312, 127, 286, 166,
	196, 123, 181, 19, 290, 175, 176, 265, 361, 221,
	19, 18, 122, 156, 293, 157, 117, 116, 230, 269,
	96, 268, 18, 136, 236, 98, 194, 97, 302, 303,
	304, 354, 308, 19, 351, 80, 246, 79, 205, 310,
	346, 245, 261, 261, 19, 243, 344, 231, 242, 241,
	73, 311, 72, 220, 240, 298, 343, 317, 315, 239,
	309, 307, 229, 238, 301, 297, 249, 292, 235, 237,
	18, 291, 156, 289, 157, 287, 276, 251, 180, 257,
	316, 18, 131, 306, 130, 18, 336, 341, 305, 299,
	340, 181, 19, 337, 175, 176, 224, 225, 18, 339,
	342, 255, 264, 19, 338, 18, 136, 19, 263, 146,
	147, 345, 102, 101, 100, 99, 95, 348, 94, 93,
	19, 92, 91, 89, 336, 341, 87, 19, 340, 84,
	335, 337, 352, 78, 281, 288, 272, 339, 296, 273,
	115, 349, 338, 6, 18, 22, 347, 14, 282, 278,
	252, 77, 76, 65, 75, 74, 71, 70, 60, 59,
	44, 68, 57, 58, 61, 362, 19, 62, 335, 350,
	66, 300, 23, 24, 67, 63, 64, 270, 5, 18,
	114, 16, 17, 27, 113, 69, 179, 180, 65, 173,
	112, 111, 110, 60, 59, 109, 68, 57, 58, 61,
	181, 19, 62, 175, 176, 66, 108, 107, 106, 67,
	63, 64, 353, 105, 104, 103, 18, 88, 83, 82,
	69, 25, 50, 191, 180, 65, 49, 51, 172, 164,
	60, 59, 163, 68, 57, 58, 61, 181, 19, 62,
	47, 18, 66, 158, 90, 46, 67, 63, 64, 180,
	65, 227, 226, 55, 223, 60, 59, 69, 68, 57,
	58, 61, 181, 19, 62, 222, 18, 66, 218, 217,
	54, 67, 63, 64, 145, 65, 144, 140, 139, 31,
	60, 59, 69, 68, 57, 58, 61, 333, 19, 62,
	85, 332, 66, 208, 207, 203, 67, 63, 64, 202,
	52, 233, 232, 65, 56, 182, 48, 69, 60, 59,
	44, 68, 57, 58, 61, 320, 318, 62, 267, 154,
	66, 148, 65, 45, 67, 63, 64, 60, 59, 215,
	68, 57, 58, 61, 213, 69, 62, 210, 53, 66,
	43, 33, 32, 67, 63, 64, 6, 18, 22, 260,
	14, 18, 29, 134, 69, 21, 128, 20, 13, 12,
	65, 9, 8, 118, 7, 60, 59, 15, 68, 19,
	26, 61, 4, 19, 62, 23, 24, 66, 216, 2,
	1, 67, 63, 64, 16, 17, 178, 177, 359,
}
var yyPact = [...]int{

	22, -1000, 542, 427, 339, -1000, 362, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 361, 250, 360, 359, 357, 356,
	333, 235, 78, 425, 424, 329, 489, -1000, -1000, -1000,
	-1000, 326, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 423, 323, 322, 321, 319, 318,
	316, 218, 225, 315, 314, 313, 312, 421, 420, 419,
	414, 413, 412, 401, 398, 397, 396, 390, 386, 344,
	215, 214, -1000, 153, 210, 199, 100, 195, 89, -1000,
	217, -1000, -1000, -1000, -1000, -1000, -1000, 293, -1000, 183,
	-1000, 461, 374, 436, 265, 265, -1000, -1000, 21, 47,
	280, 461, 546, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 143, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 267, -1000, 276, -1000,
	355, 60, -1000, -1000, 300, -1000, 54, -1000, -1000, 278,
	293, -1000, -1000, -1000, 508, 508, 308, 302, 206, -1000,
	-1000, -1000, -1000, -1000, 219, -1000, 383, 341, 31, -1000,
	-1000, -1000, -1000, 275, 374, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 354, -1000, 30, 337, -1000, -1000, 353,
	2, 2, 274, 436, -1000, -1000, -1000, -1000, -1000, -1000,
	272, 265, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 270, 266, 21, -1000, -1000, -1000, -1000, 145, 342,
	264, -1000, -1000, 6, -1000, 289, 377, 263, 280, -1000,
	-1000, -1000, 508, 508, 288, 283, 260, 461, -1000, -1000,
	-1000, -1000, 259, 546, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 341,
	-1000, -1000, 194, 189, -1000, -1000, 185, -1000, -1000, 257,
	508, -1000, 256, -1000, -1000, -1000, -1000, -1000, -1000, 50,
	-1000, 175, -1000, -1000, -1000, -1000, -1000, -1000, 169, 160,
	159, 152, 149, 148, -1000, -1000, 144, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 176, -1000, -1000, -1000, 436,
	-1000, -1000, -1000, 255, 245, -1000, -1000, -1000, -1000, -1000,
	-1000, 100, -1000, -1000, -1000, -1000, -1000, -1000, 239, 351,
	36, 346, -1000, 375, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 233, 176, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 411, -1000, -1000, 230, -1000, 141, -1000, 104,
	70, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -11, 207,
	371, -1000, 71, -1000,
}
var yyPgo = [...]int{

	0, 598, 26, 11, 7, 597, 596, 590, 589, 582,
	580, 577, 574, 573, 33, 44, 0, 9, 388, 572,
	571, 569, 568, 567, 566, 41, 565, 563, 32, 59,
	562, 29, 16, 559, 552, 551, 45, 39, 35, 30,
	25, 24, 21, 17, 12, 550, 548, 547, 544, 51,
	539, 10, 533, 531, 28, 40, 27, 529, 528, 526,
	525, 516, 515, 2, 18, 13, 514, 512, 511, 15,
	510, 509, 505, 20, 504, 503, 8, 4, 1, 501,
	497, 489, 488, 487, 43, 486, 484, 480, 479, 478,
	19, 475, 464, 463, 462, 461, 14, 455, 454, 453,
	23, 450, 442, 439, 31, 438, 437, 436, 54, 433,
	22, 432, 6,
}
var yyR1 = [...]int{

	0, 7, 8, 11, 12, 12, 13, 13, 14, 14,
	15, 9, 9, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 23, 17, 17, 24, 24, 25, 25, 25,
	25, 21, 26, 27, 27, 28, 28, 28, 22, 22,
	29, 29, 10, 10, 32, 32, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 31, 31, 33, 33,
	42, 47, 47, 47, 47, 46, 48, 48, 49, 50,
	34, 52, 53, 53, 54, 54, 54, 54, 3, 3,
	4, 56, 55, 57, 58, 58, 59, 59, 59, 37,
	61, 62, 62, 51, 51, 63, 63, 63, 63, 63,
	66, 45, 67, 67, 68, 68, 69, 69, 69, 69,
	69, 69, 69, 69, 69, 69, 69, 69, 70, 41,
	41, 71, 71, 72, 72, 73, 73, 73, 75, 76,
	76, 76, 76, 76, 76, 76, 74, 74, 79, 80,
	80, 30, 81, 82, 82, 83, 83, 84, 84, 84,
	84, 85, 86, 43, 87, 88, 88, 89, 89, 90,
	90, 90, 90, 91, 92, 44, 93, 94, 94, 95,
	95, 96, 96, 96, 35, 98, 97, 99, 99, 100,
	100, 100, 36, 101, 102, 103, 103, 77, 77, 78,
	104, 104, 104, 104, 104, 104, 104, 104, 104, 105,
	40, 106, 106, 38, 107, 108, 109, 109, 110, 110,
	110, 110, 110, 110, 110, 110, 6, 65, 2, 2,
	5, 64, 39, 111, 60, 60, 112, 112, 1, 16,
	19, 20,
}
var yyR2 = [...]int{

	0, 4, 3, 2, 2, 4, 1, 2, 1, 1,
	3, 1, 2, 3, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 5, 1, 2, 3, 3, 1,
	1, 4, 2, 1, 2, 3, 1, 1, 2, 4,
	1, 1, 1, 2, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	4, 0, 1, 1, 1, 2, 1, 2, 4, 2,
	4, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	3, 1, 2, 2, 1, 3, 3, 1, 3, 4,
	2, 0, 1, 1, 2, 1, 1, 1, 1, 1,
	2, 4, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	4, 0, 1, 1, 2, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 4, 1, 1,
	2, 4, 2, 0, 1, 1, 2, 1, 1, 3,
	3, 2, 2, 4, 2, 0, 1, 1, 2, 1,
	1, 3, 3, 2, 2, 4, 2, 0, 1, 1,
	2, 1, 1, 1, 2, 3, 2, 1, 2, 1,
	1, 1, 4, 2, 1, 1, 2, 3, 3, 3,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	2, 2, 2, 4, 2, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 1,
	3, 1, 4, 2, 1, 2, 3, 5, 3, 3,
	3, 3,
}
var yyChk = [...]int{

	-1000, -7, -8, 28, -9, -18, 14, -12, -19, -20,
	-15, -16, -21, -22, 18, -11, 52, 53, 15, 37,
	-23, -26, 16, 43, 44, 4, -10, -18, -29, -30,
	-31, -81, -34, -35, -36, -37, -38, -39, -40, -41,
	-42, -43, -44, -45, 31, -52, -97, -101, -61, -107,
	-111, -106, -70, -46, -87, -93, -66, 33, 34, 30,
	29, 35, 38, 46, 47, 24, 41, 45, 32, 56,
	5, 5, 12, 10, 5, 5, 5, 5, 10, 12,
	10, 13, 4, 4, 10, 11, -29, 10, 4, 10,
	-98, 10, 10, 10, 10, 10, 12, 12, 10, 10,
	10, 10, 10, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 6, 12, 12, -13, -14,
	-15, -16, 12, 12, -17, 12, 10, 12, -24, -25,
	18, 16, -15, -16, -27, -28, 16, -15, -16, -82,
	-83, -84, -15, -16, -85, -86, 26, 27, -53, -54,
	-55, -15, -16, -56, -57, -4, 17, 19, -99, -100,
	-15, -16, -31, -102, -103, -104, -15, -16, -77, -78,
	-64, -65, -105, 25, -31, 39, 40, -5, -6, 22,
	23, 36, -62, -51, -63, -15, -16, -64, -65, -31,
	-108, -109, -110, -55, -15, -16, -64, -77, -78, -65,
	-56, -108, -71, -72, -73, -15, -16, -74, -75, 54,
	-47, -15, -16, -48, -49, -50, 42, -88, -89, -90,
	-15, -16, -91, -92, 26, 27, -94, -95, -96, -15,
	-16, -31, -67, -68, -69, -15, -16, -36, -37, -38,
	-39, -40, -41, -42, -49, -43, -44, 11, -14, 9,
	-25, 11, 5, 13, -28, 11, 13, 11, -84, -32,
	-33, -31, -32, 10, 10, 11, -54, -58, 12, 10,
	4, -3, 5, 8, 11, -100, 11, -104, 5, 7,
	55, 7, 5, -2, 50, 51, -2, 11, -63, 11,
	-110, 11, 11, -73, 12, 10, 6, 11, -49, 10,
	4, 11, -90, -32, -32, 10, 10, 11, -96, 11,
	-69, -3, 12, 12, 12, 11, -31, 11, -59, 20,
	-60, 48, -112, 21, 12, 12, 12, 12, 12, 12,
	12, 12, -79, -80, -76, -15, -16, -4, -64, -65,
	-77, -78, -51, 11, 11, -17, 11, 5, -112, 5,
	4, 11, -76, 11, 11, 12, 12, 12, 10, -1,
	49, 11, 4, 12,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 11, 0, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 42, 40,
	41, 0, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 4, 0, 0, 0, 0, 0, 0, 38,
	0, 3, 22, 32, 2, 1, 43, 143, 142, 0,
	174, 0, 0, 91, 0, 0, 200, 119, 121, 61,
	155, 167, 102, 71, 176, 183, 90, 204, 223, 201,
	202, 118, 65, 154, 166, 100, 13, 21, 0, 6,
	8, 9, 230, 231, 10, 23, 0, 229, 0, 25,
	0, 0, 29, 30, 0, 33, 0, 36, 37, 0,
	144, 145, 147, 148, 44, 44, 0, 0, 0, 72,
	74, 75, 76, 77, 0, 81, 0, 0, 0, 177,
	179, 180, 181, 0, 184, 185, 190, 191, 192, 193,
	194, 195, 196, 0, 198, 0, 0, 221, 217, 0,
	0, 0, 0, 92, 93, 95, 96, 97, 98, 99,
	0, 205, 206, 208, 209, 210, 211, 212, 213, 214,
	215, 0, 0, 122, 123, 125, 126, 127, 0, 0,
	0, 62, 63, 64, 66, 0, 0, 0, 156, 157,
	159, 160, 44, 44, 0, 0, 0, 168, 169, 171,
	172, 173, 0, 103, 104, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 5, 7, 0,
	26, 31, 0, 0, 34, 39, 0, 141, 146, 0,
	45, 58, 0, 151, 152, 70, 73, 82, 84, 0,
	83, 0, 78, 79, 175, 178, 182, 186, 0, 0,
	0, 0, 0, 0, 218, 219, 0, 89, 94, 203,
	207, 222, 120, 124, 136, 0, 128, 60, 67, 0,
	69, 153, 158, 0, 0, 163, 164, 165, 170, 101,
	105, 0, 27, 28, 35, 149, 59, 150, 0, 0,
	87, 0, 224, 0, 80, 197, 187, 188, 189, 199,
	220, 216, 0, 138, 139, 129, 130, 131, 132, 133,
	134, 135, 0, 161, 162, 0, 85, 0, 225, 0,
	0, 137, 140, 68, 24, 86, 88, 226, 0, 0,
	0, 227, 0, 228,
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
		//line parser.y:151
		{
			m := &meta.Module{Ident: yyDollar[2].token}
			yyVAL.stack.Push(m)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:157
		{
			d := yyVAL.stack.Peek()
			r := &meta.Revision{Ident: yyDollar[2].token}
			d.(*meta.Module).Revision = r
			yyVAL.stack.Push(r)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:165
		{
			yyVAL.stack.Pop()
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:168
		{
			yyVAL.stack.Pop()
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:180
		{
			yyVAL.stack.Peek().(meta.Describable).SetDescription(tokenString(yyDollar[2].token))
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:189
		{
			d := yyVAL.stack.Peek()
			d.(*meta.Module).Namespace = tokenString(yyDollar[2].token)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:200
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Prefix = tokenString(yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:205
		{
			if yyVAL.loader == nil {
				yylex.Error("No loader defined")
				goto ret1
			} else {
				if sub, err := yyVAL.loader(yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				} else {
					i := &meta.Import{Module: sub}
					yyVAL.stack.Push(i)
				}
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:229
		{
			i := yyVAL.stack.Peek().(*meta.Import)
			i.Prefix = tokenString(yyDollar[2].token)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:238
		{
			i := yyVAL.stack.Pop().(*meta.Import)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddImport(i)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:244
		{
			if yyVAL.loader == nil {
				yylex.Error("No loader defined")
				goto ret1
			} else {
				if sub, err := yyVAL.loader(yyDollar[2].token); err != nil {
					yylex.Error(err.Error())
					goto ret1
				} else {
					i := &meta.Include{Module: sub}
					yyVAL.stack.Push(i)
				}
			}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:269
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:274
		{
			i := yyVAL.stack.Pop().(*meta.Include)
			m := yyVAL.stack.Peek().(*meta.Module)
			m.AddInclude(i)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:313
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:326
		{
			yyVAL.stack.Push(&meta.Choice{Ident: yyDollar[2].token})
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:336
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:343
		{
			yyVAL.stack.Push(&meta.ChoiceCase{Ident: yyDollar[2].token})
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:351
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:358
		{
			yyVAL.stack.Push(&meta.Typedef{Ident: yyDollar[2].token})
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:373
		{
			yyVAL.token = tokenString(yyDollar[1].token)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:374
		{
			yyVAL.token = yyDollar[1].token
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:377
		{
			yyVAL.token = yyDollar[2].token
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:381
		{
			if hasType, valid := yyVAL.stack.Peek().(meta.HasDataType); valid {
				hasType.GetDataType().SetDefault(yyDollar[1].token)
			} else {
				yylex.Error("expected default statement on meta supporting details")
				goto ret1
			}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:394
		{
			y := yyVAL.stack.Peek().(meta.HasDataType)
			y.SetDataType(meta.NewDataType(y, yyDollar[2].token))
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:404
		{
			var err error
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			if err = dataType.DecodeLength(tokenString(yyDollar[2].token)); err != nil {
				yylex.Error(err.Error())
				goto ret1
			}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:414
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			dataType := hasType.GetDataType()
			dataType.SetPath(tokenString(yyDollar[2].token))
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:424
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:431
		{
			yyVAL.stack.Push(&meta.Container{Ident: yyDollar[2].token})
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:452
		{
			yyVAL.stack.Push(&meta.Augment{Ident: yyDollar[2].token})
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:457
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:485
		{
			yyVAL.stack.Push(&meta.Uses{Ident: yyDollar[2].token})
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:490
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:495
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:520
		{
			yyVAL.stack.Push(&meta.Refine{Ident: tokenPath(yyDollar[2].token)})
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:532
		{
			s := yyDollar[1].token
			yyVAL.stack.Peek().(*meta.Refine).DefaultPtr = &s
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:543
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:548
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:562
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:569
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:583
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:588
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:595
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:600
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:608
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:615
		{
			yyVAL.stack.Push(&meta.Rpc{Ident: yyDollar[2].token})
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:629
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:634
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:641
		{
			yyVAL.stack.Push(&meta.RpcInput{})
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:646
		{
			yyVAL.stack.Push(&meta.RpcOutput{})
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:654
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:662
		{
			yyVAL.stack.Push(&meta.Notification{Ident: yyDollar[2].token})
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:682
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:694
		{
			yyVAL.stack.Push(&meta.Grouping{Ident: yyDollar[2].token})
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:710
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:717
		{
			yyVAL.stack.Push(&meta.List{Ident: yyDollar[2].token})
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:730
		{
			n, err := strconv.ParseInt(yyDollar[2].token, 10, 32)
			if err != nil || n < 1 {
				yylex.Error(fmt.Sprintf("not a valid number for max elements %s", yyDollar[2].token))
				goto ret1
			}
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetMaxElements(int(n))
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:743
		{
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetUnbounded(true)
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:753
		{
			n, err := strconv.ParseInt(yyDollar[2].token, 10, 32)
			if err != nil || n < 0 {
				yylex.Error(fmt.Sprintf("not a valid number for min elements %s", yyDollar[2].token))
				goto ret1
			}
			hasDetails, valid := yyVAL.stack.Peek().(meta.HasListDetails)
			if !valid {
				yylex.Error("expected a meta that allowed list length management")
				goto ret1
			}
			hasDetails.ListDetails().SetMinElements(int(n))
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:778
		{
			if list, valid := yyVAL.stack.Peek().(*meta.List); valid {
				list.Key = strings.Split(tokenString(yyDollar[2].token), " ")
			} else {
				yylex.Error("expected a list for key statement")
				goto ret1
			}
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:788
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:795
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:798
		{
			yyVAL.stack.Push(meta.NewAny(yyDollar[2].token))
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:806
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:813
		{
			yyVAL.stack.Push(&meta.Leaf{Ident: yyDollar[2].token})
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:836
		{
			yyVAL.boolean = yyDollar[2].boolean
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:840
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetMandatory(yyDollar[1].boolean)
			} else {
				yylex.Error("expected mandatory statement on meta supporting details")
				goto ret1
			}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:850
		{
			yyVAL.boolean = true
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:851
		{
			yyVAL.boolean = false
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:853
		{
			yyVAL.boolean = yyDollar[2].boolean
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:857
		{
			if hasDetails, valid := yyVAL.stack.Peek().(meta.HasDetails); valid {
				hasDetails.Details().SetConfig(yyDollar[1].boolean)
			} else {
				yylex.Error("expected config statement on meta supporting details")
				goto ret1
			}
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:871
		{
			if HasError(yylex, popAndAddMeta(&yyVAL)) {
				goto ret1
			}
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:878
		{
			yyVAL.stack.Push(&meta.LeafList{Ident: yyDollar[2].token})
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:887
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			hasType.GetDataType().AddEnumeration(yyDollar[2].token)
		}
	case 227:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:891
		{
			hasType := yyVAL.stack.Peek().(meta.HasDataType)
			v, nan := strconv.ParseInt(yyDollar[4].token, 10, 32)
			if nan != nil {
				yylex.Error("enum value illegal : " + nan.Error())
				goto ret1
			}
			hasType.GetDataType().AddEnumerationWithValue(yyDollar[2].token, int(v))
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:902
		{
			yyVAL.token = yyDollar[2].token
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:907
		{
			m := yyVAL.stack.Peek().(meta.Describable)
			m.SetReference(tokenString(yyDollar[2].token))
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:913
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Contact = tokenString(yyDollar[2].token)
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:919
		{
			m := yyVAL.stack.Peek().(*meta.Module)
			m.Organization = tokenString(yyDollar[2].token)
		}
	}
	goto yystack /* stack new state and value */
}
