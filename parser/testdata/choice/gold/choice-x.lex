module "module"
[ident] "choice-x"
{ "{"
container "container"
[ident] "x"
{ "{"
choice "choice"
[ident] "how"
{ "{"
config "config"
false "false"
; ";"
case "case"
[ident] "one"
{ "{"
leaf "leaf"
[ident] "y"
{ "{"
type "type"
[ident] "uint16"
; ";"
} "}"
} "}"
case "case"
[ident] "two"
{ "{"
leaf "leaf"
[ident] "z"
{ "{"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
} "}"
} "}"
container "container"
[ident] "y"
{ "{"
choice "choice"
[ident] "boogie"
{ "{"
case "case"
[ident] "one"
{ "{"
leaf "leaf"
[ident] "y"
{ "{"
type "type"
[ident] "string"
; ";"
config "config"
true "true"
; ";"
} "}"
} "}"
case "case"
[ident] "two"
{ "{"
leaf "leaf"
[ident] "z"
{ "{"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
} "}"
} "}"
} "}"
