%{
package yang

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/freeconf/c2g/val"
    "github.com/freeconf/c2g/meta"
    "github.com/freeconf/c2g/c2"
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

%}

%union {
    token    string
    boolean  bool
    num      int64
    num32    int
}

%token <token> token_ident
%token <token> token_string
%token <token> token_number
%token <token> token_custom
%token token_curly_open
%token token_curly_close
%token token_semi
%token <token> token_rev_ident

/* KEEP LIST IN SYNC WITH lexer.go */
%token kywd_namespace
%token kywd_description
%token kywd_revision
%token kywd_type
%token kywd_prefix
%token kywd_default
%token kywd_length
%token kywd_enum
%token kywd_key
%token kywd_config
%token kywd_uses
%token kywd_unique
%token kywd_input
%token kywd_output
%token kywd_module
%token kywd_container
%token kywd_list
%token kywd_rpc
%token kywd_notification
%token kywd_typedef
%token kywd_grouping
%token kywd_leaf
%token kywd_mandatory
%token kywd_reference
%token kywd_leaf_list
%token kywd_max_elements
%token kywd_min_elements
%token kywd_choice
%token kywd_case
%token kywd_import
%token kywd_include
%token kywd_action
%token kywd_anyxml
%token kywd_anydata
%token kywd_path
%token kywd_value
%token kywd_true
%token kywd_false
%token kywd_contact
%token kywd_organization
%token kywd_refine
%token kywd_unbounded
%token kywd_augment
%token kywd_submodule
%token kywd_str_plus
%token kywd_identity
%token kywd_base
%token kywd_feature
%token kywd_if_feature

%type <num32> enum_value
%type <boolean> bool_value
%type <num32> int_value
%type <token> string_or_number
%type <token> string_value

%%

module :
    module_def
    module_stmts
    module_body_stmts
    token_curly_close
    /* don't pop, leave on stack */

module_def :
    kywd_module token_ident token_curly_open {
        l := yylex.(*lexer)
        if l.parent != nil {
            l.Error("expected submodule for include")
            goto ret1
        } 
        yylex.(*lexer).stack.Push(meta.NewModule($2))
    }
    | kywd_submodule token_ident token_curly_open {        
        l := yylex.(*lexer)
        if l.parent == nil {
            /* may want to allow this is parsing submodules on their own has value */
            l.Error("submodule is for includes")
            goto ret1
        } 
        l.stack.Push(l.parent)
    }

module_stmts :
    module_stmt
    | module_stmts module_stmt

module_stmt :
    kywd_namespace string_value token_semi {
         if set(yylex, meta.SetNamespace($2)) {
            goto ret1
         }
    }
    | revision_stmt
    | contact_stmt
    | organization_stmt
    | description
    | reference_stmt
    | import_stmt
    | include_stmt
    | prefix_stmt

revision_def :
    kywd_revision token_rev_ident {
        if push(yylex, meta.NewRevision(peek(yylex), $2)) {
            goto ret1
        }
    }

revision_stmt :
    revision_def token_semi {
        pop(yylex)
    }
    | revision_def token_curly_open revision_body_stmts token_curly_close {
        pop(yylex)
    }

revision_body_stmts :
    revision_body_stmt
    | revision_body_stmts revision_body_stmt

revision_body_stmt :
    description
    | reference_stmt

import_def : 
    kywd_import token_ident {
        if push(yylex, meta.NewImport(peek(yylex).(*meta.Module), $2, yylex.(*lexer).loader)) {
            goto ret1
        }
    }

import_body_stmts :
    import_body_stmt
    | import_body_stmts import_body_stmt

prefix_stmt: 
    kywd_prefix string_value token_semi {
        if set(yylex, meta.SetPrefix($2)) {
            goto ret1
        }
     }

import_body_stmt :
     prefix_stmt
     | kywd_revision token_rev_ident token_semi
     | description
     | reference_stmt

import_stmt : 
    import_def token_curly_open import_body_stmts token_curly_close {
        pop(yylex)
    }

include_def : 
    kywd_include token_ident {
        if push(yylex, meta.NewInclude(peek(yylex).(*meta.Module), $2, yylex.(*lexer).loader)) {
            goto ret1
        }
    }

include_body_stmts :
    include_body_stmt
    | include_body_stmts include_body_stmt

include_body_stmt :
     kywd_revision token_rev_ident token_semi
     | description
     | reference_stmt

include_stmt :
    include_def token_semi {
        pop(yylex)
    }
    | include_def token_curly_open include_body_stmts token_curly_close {
        pop(yylex)
    }

module_body_stmt :
    rpc_stmt
    | body_stmt

module_body_stmts :
    module_body_stmt
    | module_body_stmts module_body_stmt

optional_body_stmts :
    /*empty*/
    | body_stmts

body_stmt :
    typedef_stmt
    | grouping_stmt
    | list_stmt
    | container_stmt
    | leaf_stmt
    | leaf_list_stmt
    | anyxml_stmt
    | uses_stmt
    | choice_stmt
    | action_stmt
    | notification_stmt
    | augment_stmt
    | identity_stmt
    | feature_stmt

body_stmts :
    body_stmt | body_stmts body_stmt


feature_stmt : 
    feature_def token_semi {
        pop(yylex)        
    }
    | feature_def token_curly_open optional_feature_body_stmts token_curly_close {
        pop(yylex)
    }

feature_def :
    kywd_feature token_ident {
        if push(yylex, meta.NewFeature(peek(yylex).(*meta.Module), $2)) {
            goto ret1
        }        
    }

optional_feature_body_stmts :
    /* empty */
    | feature_body_stmts

feature_body_stmts :    
    feature_body_stmt | feature_body_stmts feature_body_stmt

feature_body_stmt :    
    description
    | reference_stmt

identity_stmt : 
    identity_def token_semi {
        pop(yylex)        
    }
    | identity_def token_curly_open optional_identity_body_stmts token_curly_close {
        pop(yylex)
    }

identity_def :
    kywd_identity token_ident {
        if push(yylex, meta.NewIdentity(peek(yylex).(*meta.Module), $2)) {
            goto ret1
        }        
    }

optional_identity_body_stmts :
    /* empty */
    | identity_body_stmts

identity_body_stmts :    
    identity_body_stmt | identity_body_stmts identity_body_stmt

identity_body_stmt :    
    description
    | reference_stmt    
    | base_stmt
    
base_stmt :    
    kywd_base token_ident token_semi {
        if set(yylex, meta.SetBase($2)) {
            goto ret1
        }
    }

choice_stmt :
    choice_def
    token_curly_open
    choice_stmt_body
    token_curly_close {
        pop(yylex)
    }

choice_stmt_body :
    /* empty */
    | description
    | reference_stmt    
    | case_stmts

choice_def :
    kywd_choice token_ident {
        if push(yylex, meta.NewChoice(peek(yylex), $2)) {
            goto ret1
        }
    }

case_stmts :
    case_stmt | case_stmts case_stmt

case_stmt :
    case_def token_curly_open
    container_body_stmts
    token_curly_close {
        pop(yylex)
    }

case_def :
    kywd_case token_ident {
        if push(yylex, meta.NewChoiceCase(peek(yylex), $2)) {
            goto ret1
        }
    }

typedef_stmt :
    typedef_def
    token_curly_open
    typedef_stmt_body
    token_curly_close {
        pop(yylex)
    }

typedef_def :
    kywd_typedef token_ident {
        if push(yylex, meta.NewTypedef(peek(yylex), $2)) {
            goto ret1
        }
    }

typedef_stmt_body :
    typedef_stmt_body_stmt
    | typedef_stmt_body typedef_stmt_body_stmt

typedef_stmt_body_stmt:
    type_stmt
    | description
    | reference_stmt
    | default_stmt

string_or_number : 
    string_value { $$ = $1 }
    | token_number { $$ = $1 }

default_stmt :
    kywd_default string_or_number token_semi {
        if set(yylex, meta.SetDefault{Value:$2}) {
            goto ret1            
        }
    }

type_stmt :
    type_stmt_def type_stmt_body

type_stmt_def :
    kywd_type token_ident {
        if push(yylex, meta.NewDataType(peek(yylex), $2)) {
            goto ret1
        }
    }

type_stmt_body :
    token_semi {
        pop(yylex)
    }
    | token_curly_open type_stmt_types token_curly_close {
        pop(yylex)
    }

type_stmt_types :
    kywd_length string_value token_semi {
        if set(yylex, meta.SetEncodedLength($2)) {  
            goto ret1            
        }
    }
    | enum_stmts
    | base_stmt
    | type_stmts
    | kywd_path string_value token_semi {        
        if set(yylex, meta.SetPath($2)) {  
            goto ret1            
        }
    }

type_stmts :
    type_stmt | type_stmts type_stmt
    

container_stmt :
    container_def
    token_curly_open
    optional_container_body_stmts
    token_curly_close {
        pop(yylex)
    }

container_def :
    kywd_container token_ident {
        if push(yylex, meta.NewContainer(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_container_body_stmts :
	/* empty */
	| container_body_stmts

container_body_stmts :
    container_body_stmt
    | container_body_stmts container_body_stmt

container_body_stmt :
    description
    | reference_stmt
    | config_stmt
    | mandatory_stmt
    | body_stmt

augment_def :
    kywd_augment string_value {
        if push(yylex, meta.NewAugment(peek(yylex), $2)) {
            goto ret1
        }
    }

augment_stmt :
    augment_def token_curly_open optional_augment_body_stmts token_curly_close {
        pop(yylex)
    }

optional_augment_body_stmts :
    /* empty */
    | augment_body_stmts    

augment_body_stmts :
    augment_body_stmt | augment_body_stmts augment_body_stmt

augment_body_stmt : 
    description
    | reference_stmt
    | list_stmt
    | container_stmt
    | leaf_stmt
    | leaf_list_stmt
    | anyxml_stmt
    | uses_stmt
    | choice_stmt
    | case_stmt
    | action_stmt
    | notification_stmt

uses_def :
    kywd_uses token_ident {
        if push(yylex, meta.NewUses(peek(yylex), $2)) {
            goto ret1
        }
    }

uses_stmt :
    uses_def token_semi {
        pop(yylex)
    }
    | uses_def token_curly_open optional_uses_body_stmts token_curly_close {
        pop(yylex)
    }

optional_uses_body_stmts :
    /* empty */
    | uses_body_stmts

uses_body_stmts :
    uses_body_stmt | uses_body_stmts uses_body_stmt

uses_body_stmt :
    /*
     when
     if-feature
     status
     uses-augment
    */
    description
    | reference_stmt
    | refine_stmt

refine_def : 
    kywd_refine string_value {
        if push(yylex, meta.NewRefine(peek(yylex).(*meta.Uses), $2)) {
            goto ret1
        }
    }

refine_body_stmt :
    /*
     if-feature
     when
     presence
    */
    description
    | reference_stmt
    | default_stmt
    | config_stmt 
    | mandatory_stmt
    | max_elements
    | min_elements

refine_stmt : 
    /* I question the point of this. declaring a refinement w/no details */
    refine_def token_semi {
        pop(yylex)
    }
    | refine_def token_curly_open optional_refine_body_stmts token_curly_close {
        pop(yylex)
    }

optional_refine_body_stmts :
    /* empty */
    refine_body_stmts

refine_body_stmts  :
    refine_body_stmt | refine_body_stmts refine_body_stmt

rpc_stmt :
    rpc_def token_curly_open optional_rpc_body_stmts token_curly_close {
        pop(yylex)
    }

rpc_def :
    kywd_rpc token_ident {
        if push(yylex, meta.NewRpc(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_rpc_body_stmts :
    /* empty */
    | rpc_body_stmts

rpc_body_stmts :
    rpc_body_stmt | rpc_body_stmts rpc_body_stmt

rpc_body_stmt:
    description
    | reference_stmt
    | rpc_input optional_body_stmts token_curly_close {
        pop(yylex)
    }
    | rpc_output optional_body_stmts token_curly_close {
        pop(yylex)
    }

rpc_input :
    kywd_input token_curly_open {
        if push(yylex, meta.NewRpcInput(peek(yylex).(*meta.Rpc))) {
            goto ret1
        }
    }

rpc_output :
    kywd_output token_curly_open {
        if push(yylex, meta.NewRpcOutput(peek(yylex).(*meta.Rpc))) {
            goto ret1
        }
    }

action_stmt :
    action_def
    token_curly_open
    optional_action_body_stmts
    token_curly_close {
        pop(yylex)
    }

action_def :
    kywd_action token_ident {
        if push(yylex, meta.NewRpc(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_action_body_stmts :
    /* empty */
    | action_body_stmts

action_body_stmts :
    action_body_stmt | action_body_stmts action_body_stmt

action_body_stmt:
    description
    | reference_stmt
    | rpc_input optional_body_stmts token_curly_close {
        pop(yylex)
    }
    | rpc_output optional_body_stmts token_curly_close {
        pop(yylex)
    }

notification_stmt :
    notification_def
    token_curly_open
    optional_notification_body_stmts
    token_curly_close {
        pop(yylex)
    }

notification_def :
    kywd_notification token_ident {
        if push(yylex, meta.NewNotification(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_notification_body_stmts :
	/* empty */
	| notification_body_stmts

notification_body_stmts :
    notification_body_stmt
    | notification_body_stmts notification_body_stmt

notification_body_stmt :
    description
    | reference_stmt
    | body_stmt

grouping_stmt :
    grouping_def
    grouping_body_defined {
        pop(yylex)
    }

grouping_body_defined:
    token_curly_open
    grouping_body_stmts
    token_curly_close

grouping_def :
    kywd_grouping token_ident {
        if push(yylex, meta.NewGrouping(peek(yylex), $2)) {
            goto ret1
        }
    }

grouping_body_stmts :
    grouping_body_stmt |
    grouping_body_stmts grouping_body_stmt

grouping_body_stmt :
    description
    | reference_stmt
    | body_stmt

list_stmt :
    list_def token_curly_open
    optional_list_body_stmts
    token_curly_close{
        pop(yylex)
     }

list_def :
    kywd_list token_ident {
        if push(yylex, meta.NewList(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_list_body_stmts :
	/* empty */
	list_body_stmts

list_body_stmts :
    list_body_stmt
    | list_body_stmts list_body_stmt

max_elements : 
    kywd_max_elements int_value token_semi {
        if set(yylex, meta.SetMaxElements($2)) {
            goto ret1
        }
    }
    | kywd_max_elements kywd_unbounded token_semi {
        if set(yylex, meta.SetUnbounded(true)) {
            goto ret1
        }
    }

min_elements : 
    kywd_min_elements int_value token_semi {
        if set(yylex, meta.SetMinElements($2)) {
            goto ret1
        }
    }

list_body_stmt :
    description
    | reference_stmt
    | max_elements
    | min_elements
    | config_stmt
    | mandatory_stmt
    | key_stmt
    | kywd_unique string_value token_semi
    | body_stmt

key_stmt: 
    kywd_key string_value token_semi {
        if set(yylex, meta.SetKey($2)) {
            goto ret1
        }
    }

anyxml_stmt:
    anyxml_def token_semi {
        pop(yylex)
    }

anyxml_def :
    kywd_anyxml token_ident {
        if push(yylex, meta.NewAny(peek(yylex), $2)) {
            goto ret1
        }
    }
    | kywd_anydata token_ident {
        if push(yylex, meta.NewAny(peek(yylex), $2)) {
            goto ret1
        }
    }

leaf_stmt:
    leaf_def
    token_curly_open
    optional_leaf_body_stmts
    token_curly_close {
        pop(yylex)
     }

leaf_def :
    kywd_leaf token_ident {
        if push(yylex, meta.NewLeaf(peek(yylex), $2)) {
            goto ret1
        }
    }

optional_leaf_body_stmts:
	/* empty */
	leaf_body_stmts

leaf_body_stmts :
    leaf_body_stmt
    | leaf_body_stmts leaf_body_stmt

/* TODO: when, if, units, must, status, reference */
leaf_body_stmt :
    type_stmt
    | description
    | reference_stmt
    | config_stmt
    | max_elements
    | min_elements
    | mandatory_stmt
    | default_stmt

mandatory_stmt : 
    kywd_mandatory bool_value token_semi {
        if set(yylex, meta.SetMandatory($2)) {
            goto ret1
        }
    }

string_value :
    token_string {
        $$ = tokenString($1)
    }
    | string_value kywd_str_plus token_string {
        $$ = $1 + tokenString($3)
    }

int_value : 
    token_number {
        n, err := strconv.ParseInt($1, 10, 32)
        if err != nil || n < 0 {
            yylex.Error(fmt.Sprintf("not a valid number for min elements %s", $1))
            goto ret1
        }       
        $$ = int(n)
    }

bool_value :
    kywd_true {$$ = true} 
    | kywd_false {$$ = false}

config_stmt : 
    kywd_config bool_value token_semi {
        if set(yylex, meta.SetConfig($2)) {
            goto ret1
        }
    }

leaf_list_stmt :
    leaf_list_def
    token_curly_open
    optional_leaf_body_stmts
    token_curly_close {
        pop(yylex)
    }

leaf_list_def :
    kywd_leaf_list token_ident {
        if push(yylex, meta.NewLeafList(peek(yylex), $2)) {
            goto ret1
        }
    }

enum_stmts :
    enum_stmt
    | enum_stmts enum_stmt

enum_stmt :
    kywd_enum token_ident token_semi {
        if set(yylex, meta.SetEnumLabel($2))  {
            goto ret1
        }
    }
    | kywd_enum token_ident token_curly_open enum_value token_curly_close {        
        if set(yylex, val.Enum{Label:$2, Id:$4})  {
            goto ret1
        }
    }

enum_value :
    kywd_value int_value token_semi {
        $$ = $2
    }

description : 
    kywd_description string_value statement_end {
        if set(yylex, meta.SetDescription($2)) {
            goto ret1
        }
    }

reference_stmt :
    kywd_reference string_value token_semi {        
        if set(yylex, meta.SetReference($2)) {
            goto ret1
        }
    }

contact_stmt :
    kywd_contact string_value token_semi {
        if set(yylex, meta.SetContact($2)) {
            goto ret1
        }
    }

organization_stmt :
    kywd_organization string_value token_semi {
        if set(yylex, meta.SetOrganization($2)) {
            goto ret1
        }
    }

statement_end :
    token_semi
    | token_curly_open token_custom string_or_number statement_end token_curly_close  
%%

