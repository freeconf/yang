module "module"
[ident] "union"
{ "{"
namespace "namespace"
[string] "\"\""
; ";"
prefix "prefix"
[string] "\"\""
; ";"
revision "revision"
[string] "0"
; ";"
typedef "typedef"
[ident] "t"
{ "{"
type "type"
[ident] "union"
{ "{"
type "type"
[ident] "enumeratio"...
{ "{"
enum "enum"
[string] "a"
; ";"
enum "enum"
[string] "b"
; ";"
} "}"
type "type"
[ident] "int8"
; ";"
} "}"
} "}"
leaf "leaf"
[ident] "x"
{ "{"
type "type"
[ident] "union"
{ "{"
type "type"
[ident] "int32"
; ";"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
leaf-list "leaf-list"
[ident] "y"
{ "{"
type "type"
[ident] "union"
{ "{"
type "type"
[ident] "int64"
; ";"
type "type"
[ident] "decimal64"
; ";"
} "}"
} "}"
leaf "leaf"
[ident] "q"
{ "{"
type "type"
[ident] "t"
; ";"
} "}"
leaf-list "leaf-list"
[ident] "p"
{ "{"
type "type"
[ident] "t"
; ";"
} "}"
} "}"
