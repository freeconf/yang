module "module"
[ident] "example-ba"...
{ "{"
namespace "namespace"
[string] "\"http://ex"...
; ";"
prefix "prefix"
[string] "\"barmod\""
; ";"
import "import"
[ident] "example-fo"...
{ "{"
prefix "prefix"
[string] "\"foomod\""
; ";"
} "}"
uses "uses"
[ident] "foomod:x"
; ";"
augment "augment"
[string] "/top"
{ "{"
leaf "leaf"
[ident] "bar"
{ "{"
type "type"
[ident] "boolean"
; ";"
} "}"
} "}"
} "}"
