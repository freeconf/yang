module "module"
[ident] "-units"
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
[ident] "lr"
{ "{"
default "default"
[string] "7"
; ";"
units "units"
[string] "u"
; ";"
type "type"
[ident] "int32"
; ";"
} "}"
leaf "leaf"
[ident] "y"
{ "{"
type "type"
[ident] "union"
{ "{"
type "type"
[ident] "int16"
; ";"
type "type"
[ident] "lr"
; ";"
} "}"
} "}"
} "}"
