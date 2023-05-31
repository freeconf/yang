%{
package parser

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/freeconf/yang/meta"
)

func tokenString(s string) string {
	s = strings.Trim(s, " \t\n\r")
	lastChar := len(s) -1
	if s[0] == char_doublequote && s[lastChar] == char_doublequote {
		return s[1:lastChar]
	}
	if s[0] == char_singlequote && s[lastChar] == char_singlequote {
		return s[1:lastChar]
	}
	return s
}

// Lex implements goyacc interface
func (l *lexer) Lex(lval *yySymType) int {
    t, _ := l.nextToken()
    if t.typ == parseEof {
        return 0
    }
    lval.token = t.val
    return int(t.typ)
}

// Error implements goyacc interface
func (l *lexer) Error(e string) {
    line, col := l.Position()
    l.lastError = fmt.Errorf("%s - line %d, col %d", e, line, col)
}

func chkErr(l yyLexer, e error) bool {
    if e == nil {
        return false
    }
    l.Error(e.Error())
    return true
}

func chkErr2(l *lexer, keyword string, extension *meta.Extension) bool {
    if extension != nil {
        l.builder.AddExtension(l.stack.peek(), keyword, extension)
    }
    if l.builder.LastErr != nil {
        l.Error(l.builder.LastErr.Error())
        return true
    }

    return false
}

func trimQuotes(s string) string {
    if s[0] == '"' {
        return s[1:len(s)-1]
    }
    return s
}

%}

%union {
    token    string
    boolean  bool
    num      int64
    num32    int
    args     []string
    ext      *meta.Extension
}

%token <token> token_ident
%token <token> token_string
%token <token> token_number
%token <token> token_unknown
%token token_curly_open
%token token_curly_close
%token token_semi

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
%token kywd_when
%token kywd_must
%token kywd_yang_version
%token kywd_range
%token kywd_extension
%token kywd_argument
%token kywd_yin_element
%token kywd_pattern
%token kywd_modifier
%token kywd_invert_match
%token kywd_units
%token kywd_fraction_digits
%token kywd_status
%token kywd_current
%token kywd_obsolete
%token kywd_deprecated
%token kywd_presence
%token kywd_deviation
%token kywd_deviate
%token kywd_not_supported
%token kywd_add
%token kywd_replace
%token kywd_delete
%token kywd_ordered_by
%token kywd_system
%token kywd_user
%token kywd_require_instance
%token kywd_error_app_tag
%token kywd_error_message
%token kywd_bit
%token kywd_position
%token kywd_revision_date
%token kywd_belongs_to

%type <boolean> bool_value
%type <num32> int_value
%type <token> string_or_number
%type <token> string_value
%type <args> optional_unknown_args
%type <args> unknown_args
%type <ext> unknown_stmt
%type <ext> statement_end

%%

module :
    module_def
    module_stmts
    token_curly_close
    /* don't pop, leave on stack */

module_def :
    kywd_module token_ident token_curly_open {
        l := yylex.(*lexer)
        if l.parent != nil {
            l.Error("expected submodule for include")
            goto ret1
        }        
        yylex.(*lexer).stack.push(l.builder.Module($2, l.featureSet))
    }
    | kywd_submodule token_ident token_curly_open {
        l := yylex.(*lexer)
        if l.parent == nil {
            // may want to allow this is parsing submodules on their own has value
            l.Error("submodule is for includes")
            goto ret1
        } 
        // sub modules really just re-add parent module back onto stack and let all 
        // children be added to that.
        l.stack.push(l.builder.Submodule(l.parent, $2, l.featureSet))
    }

module_stmts :
    module_stmt
    | module_stmts module_stmt

module_stmt :
    kywd_namespace string_value statement_end {
        l := yylex.(*lexer)
        l.builder.Namespace(l.stack.peek(), $2)
        if chkErr2(l, "namespace", $3) {
            goto ret1
        }
    }
    | revision_stmt
    | contact_stmt
    | organization_stmt
    | description
    | status_stmt
    | reference_stmt
    | import_stmt
    | include_stmt
    | prefix_stmt
    | yang_ver_stmt
    | rpc_stmt
    | extension_def_stmt
    | deviation_stmt
    | belongs_to_stmt
    | body_stmt


belongs_to_def :
    kywd_belongs_to token_string {
        l := yylex.(*lexer)
        l.stack.push(l.builder.BelongsTo(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

belongs_to_stmt :
    belongs_to_def token_curly_open prefix_stmt token_curly_close {
        yylex.(*lexer).stack.pop()
    }

revision_def :
    kywd_revision token_string {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Revision(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

revision_stmt :
    revision_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | revision_def token_curly_open revision_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

revision_body_stmts :
    revision_body_stmt
    | revision_body_stmts revision_body_stmt

revision_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | unknown_stmt

import_def : 
    kywd_import token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Import(l.stack.peek(), $2, l.loader))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

import_body_stmts :
    import_body_stmt
    | import_body_stmts import_body_stmt

prefix_stmt: 
    kywd_prefix string_value token_semi {
        l := yylex.(*lexer)
        l.builder.Prefix(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
     }

import_body_stmt :
     prefix_stmt
     | kywd_revision token_string token_semi
     | description
     | status_stmt
     | reference_stmt
     | revision_date_stmt
     | unknown_stmt

import_stmt : 
    import_def token_curly_open import_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

include_def : 
    kywd_include token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Include(l.stack.peek(), $2, yylex.(*lexer).loader))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

include_body_stmts :
    include_body_stmt
    | include_body_stmts include_body_stmt

include_body_stmt :
     kywd_revision token_string token_semi
     | description
     | status_stmt
     | reference_stmt
     | revision_date_stmt
     | unknown_stmt

include_stmt :
    include_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | include_def token_curly_open include_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

revision_date_stmt :
    kywd_revision_date token_string token_semi {
        l := yylex.(*lexer)
        l.builder.SetRevisionDate(l.stack.peek(), $2)
    }

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
    | must_stmt
    | choice_stmt
    | action_stmt
    | notification_stmt
    | augment_stmt
    | identity_stmt
    | feature_stmt
    | unknown_stmt

body_stmts :
    body_stmt | body_stmts body_stmt

extension_def_stmt :
    extension_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | extension_def token_curly_open optional_extension_def_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

extension_def : 
    kywd_extension token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ExtensionDef(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

optional_extension_def_body_stmts :
    /* empty */
    | extension_def_body_stmts

extension_def_body_stmts :
    extension_def_body_stmt | extension_def_body_stmts extension_def_body_stmt

extension_def_body_stmt :
    extension_def_argument_stmt
    | description
    | status_stmt
    | reference_stmt

extension_def_argument_stmt :
    argument_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | argument_def token_curly_open optional_argument_def_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

argument_def :
    kywd_argument token_string {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ExtensionDefArg(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

optional_argument_def_body_stmts :
    /* empty */
    | argument_def_body_stmts

argument_def_body_stmts :
    argument_def_body_stmt | argument_def_body_stmts argument_def_body_stmt

argument_def_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | yin_element_stmt

yin_element_stmt : 
    kywd_yin_element bool_value token_semi {
        l := yylex.(*lexer)
        l.builder.YinElement(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

deviation_stmt :
    deviation_def token_curly_open deviation_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

deviation_def :
    kywd_deviation string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Deviation(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

deviation_body_stmts :
        deviation_body_stmt | deviation_body_stmts deviation_body_stmt

deviation_body_stmt :
    description
    | reference_stmt
    /* 
      could get fancy here and limit
      either
         one not-supported
      or
         one or more others
    */
    | deviate_not_supported
    | deviate_replace_def deviate_stmt
    | deviate_delete_def deviate_stmt
    | deviate_add_def deviate_stmt
    | unknown_stmt

deviate_not_supported:
    kywd_deviate kywd_not_supported statement_end {
        l := yylex.(*lexer)
        l.builder.NotSupported(l.stack.peek())
        if chkErr2(l, "not-supported", $3) {
            goto ret1
        }
    }

deviate_replace_def :
    kywd_deviate kywd_replace {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ReplaceDeviate(l.stack.peek()))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

deviate_delete_def :
    kywd_deviate kywd_delete {
        l := yylex.(*lexer)
        l.stack.push(l.builder.DeleteDeviate(l.stack.peek()))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

deviate_add_def :
    kywd_deviate kywd_add {
        l := yylex.(*lexer)
        l.stack.push(l.builder.AddDeviate(l.stack.peek()))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

deviate_stmt :
    token_curly_open deviate_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

deviate_body_stmts :
    deviate_body_stmt | deviate_body_stmts deviate_body_stmt

/* 
  superset of all deviates but builder will validate
  applicability
*/
deviate_body_stmt :
    units_stmt
    | must_stmt
    | unique_stmt
    | default_stmt
    | config_stmt
    | mandatory_stmt
    | max_elements
    | min_elements
    | type_stmt /* replace only */

feature_stmt : 
    feature_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | feature_def token_curly_open optional_feature_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }


feature_def :
    kywd_feature token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Feature(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | unknown_stmt

must_stmt :
    must_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | must_def token_curly_open error_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    } 

must_def :
    kywd_must string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Must(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

error_body_stmts :
    error_body_stmt | error_body_stmts error_body_stmt

error_body_stmt :
    description
    | reference_stmt
    | error_message_stmt
    | error_app_tag_stmt
    | unknown_stmt

error_message_stmt :
    kywd_error_message string_value statement_end {
        l := yylex.(*lexer)
        l.builder.ErrorMessage(l.stack.peek(), $2)
        if chkErr2(l, "error-message", $3) {
            goto ret1
        }        
    }

error_app_tag_stmt :
    kywd_error_app_tag string_value statement_end {
        l := yylex.(*lexer)
        l.builder.ErrorAppTag(l.stack.peek(), $2)
        if chkErr2(l, "error-app-tag", $3) {
            goto ret1
        }        
    }


if_feature_stmt :
    kywd_if_feature string_value statement_end {
        l := yylex.(*lexer)
        l.builder.IfFeature(l.stack.peek(), $2)
        if chkErr2(l, "if-feature", $3) {
            goto ret1
        }
    }

when_def : 
    kywd_when string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.When(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

when_stmt :
    when_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | when_def token_curly_open when_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

when_body_stmts :
    when_body_stmt | when_body_stmts when_body_stmt

when_body_stmt :
    description
    | status_stmt
    | reference_stmt    
    | unknown_stmt

identity_stmt : 
    identity_def token_semi {
        yylex.(*lexer).stack.pop()        
    }
    | identity_def token_curly_open optional_identity_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

identity_def :
    kywd_identity token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Identity(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt    
    | base_stmt
    | if_feature_stmt
    | unknown_stmt
    
base_stmt :    
    kywd_base token_ident statement_end {
        l := yylex.(*lexer)        
        l.builder.Base(l.stack.peek(), $2)
        if chkErr2(l, "base", $3) {
            goto ret1
        }
    }

choice_stmt :
    choice_def token_curly_open choice_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

choice_body_stmts :
    choice_body_stmt | choice_body_stmts choice_body_stmt


choice_body_stmt :
    description
    | status_stmt
    | reference_stmt    
    | case_stmt
    | body_stmt
    | if_feature_stmt
    | when_stmt
    | mandatory_stmt
    | default_stmt

choice_def :
    kywd_choice token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Choice(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

case_stmt :
    case_def token_curly_open container_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

case_def :
    kywd_case token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Case(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

typedef_stmt :
    typedef_def token_curly_open typedef_stmt_body token_curly_close {
        yylex.(*lexer).stack.pop()
    }

typedef_def :
    kywd_typedef token_ident {
        l := yylex.(*lexer)        
        l.stack.push(l.builder.Typedef(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

typedef_stmt_body :
    typedef_stmt_body_stmt | typedef_stmt_body typedef_stmt_body_stmt

typedef_stmt_body_stmt:
    type_stmt
    | units_stmt
    | description
    | status_stmt
    | reference_stmt
    | default_stmt
    | unknown_stmt

string_or_number : 
    string_value { $$ = $1 }
    | token_number { $$ = $1 }

default_stmt :
    kywd_default string_value statement_end {
        l := yylex.(*lexer)        
        l.builder.Default(l.stack.peek(), $2)
        if chkErr2(l, "default", $3) {
            goto ret1
        }
    }

type_stmt :
    type_stmt_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | type_stmt_def token_curly_open optional_type_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

type_stmt_def :
    kywd_type token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Type(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

optional_type_body_stmts :
    /* empty */
    | type_body_stmts

type_body_stmts :
    type_body_stmt | type_body_stmts type_body_stmt

type_body_stmt :
    type_detail_stmt
    | kywd_path string_value statement_end {    
        l := yylex.(*lexer)
        l.builder.Path(l.stack.peek(), $2)
        if chkErr2(l, "path", $3) {
            goto ret1
        }
    }    
    | enum_stmt
    | bit_stmt
    | base_stmt
    | fraction_digits_stmt
    | type_stmt
    | require_instance_stmt
    | unknown_stmt

type_detail_stmt :
    type_detail_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | type_detail_def token_curly_open type_detail_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

type_detail_def :     
    kywd_range string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ValueRange(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }
    | kywd_length string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.LengthRange(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }
    | kywd_pattern string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Pattern(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

type_detail_body_stmts :
    type_detail_body_stmt | type_detail_body_stmts type_detail_body_stmt

type_detail_body_stmt :
    description
    | reference_stmt
    | error_message_stmt
    | error_app_tag_stmt
    | modifier_stmt
    | unknown_stmt

modifier_stmt:
    kywd_modifier kywd_invert_match token_semi {
        l := yylex.(*lexer)
        l.builder.SetInverted(l.stack.peek())
    }

require_instance_stmt :
    kywd_require_instance bool_value statement_end {
        l := yylex.(*lexer)
        l.builder.RequireInstance(l.stack.peek(), $2)
        if chkErr2(l, "require-instance", $3) {
            goto ret1
        }        
    }

status_stmt : 
    kywd_status kywd_current token_semi
    | kywd_status kywd_obsolete token_semi
    | kywd_status kywd_deprecated token_semi

fraction_digits_stmt :
    kywd_fraction_digits int_value token_semi {
        l := yylex.(*lexer)
        l.builder.FractionDigits(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

container_stmt :
    container_def token_curly_open optional_container_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

container_def :
    kywd_container token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Container(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | when_stmt
    | config_stmt
    | mandatory_stmt
    | presence_stmt
    | body_stmt

presence_stmt :
    kywd_presence string_value statement_end {
        l := yylex.(*lexer)
        l.builder.Presence(l.stack.peek(), $2)     
        if chkErr2(l, "presence", $3) {
            goto ret1
        }
    }

augment_def :
    kywd_augment string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Augment(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

augment_stmt :
    augment_def token_curly_open optional_augment_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

optional_augment_body_stmts :
    /* empty */
    | augment_body_stmts    

augment_body_stmts :
    augment_body_stmt | augment_body_stmts augment_body_stmt

augment_body_stmt : 
    description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | when_stmt
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
    | unknown_stmt

uses_def :
    kywd_uses token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Uses(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

uses_stmt :
    uses_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | uses_def token_curly_open optional_uses_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

optional_uses_body_stmts :
    /* empty */
    | uses_body_stmts

uses_body_stmts :
    uses_body_stmt | uses_body_stmts uses_body_stmt

uses_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | when_stmt
    | refine_stmt
    | augment_stmt
    | unknown_stmt

refine_def : 
    kywd_refine string_value {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Refine(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

refine_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | default_stmt
    | config_stmt 
    | mandatory_stmt
    | must_stmt
    | max_elements
    | min_elements
    | unknown_stmt

refine_stmt : 
    /* I question the point of this. declaring a refinement w/no details */
    refine_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | refine_def token_curly_open optional_refine_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

optional_refine_body_stmts :
    /* empty */
    refine_body_stmts

refine_body_stmts  :
    refine_body_stmt | refine_body_stmts refine_body_stmt

rpc_stmt :
    rpc_def token_curly_open optional_rpc_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

rpc_def :
    kywd_rpc token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Action(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | rpc_input optional_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }
    | rpc_output optional_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }
    | unknown_stmt

rpc_input :
    kywd_input token_curly_open {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ActionInput(l.stack.peek()))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

rpc_output :
    kywd_output token_curly_open {
        l := yylex.(*lexer)
        l.stack.push(l.builder.ActionOutput(l.stack.peek()))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

action_stmt :
    action_def
    token_curly_open
    optional_action_body_stmts
    token_curly_close {
        yylex.(*lexer).stack.pop()
    }

action_def :
    kywd_action token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Action(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | rpc_input optional_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }
    | rpc_output optional_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }
    | unknown_stmt

notification_stmt :
    notification_def
    token_curly_open
    optional_notification_body_stmts
    token_curly_close {
        yylex.(*lexer).stack.pop()
    }

notification_def :
    kywd_notification token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Notification(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | body_stmt

grouping_stmt :
    grouping_def token_curly_open optional_grouping_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }    

grouping_def :
    kywd_grouping token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Grouping(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

optional_grouping_body_stmts : 
    /* empty */
    | grouping_body_stmts

grouping_body_stmts :
    grouping_body_stmt | grouping_body_stmts grouping_body_stmt

grouping_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | body_stmt

list_stmt :
    list_def token_curly_open optional_list_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
     }

list_def :
    kywd_list token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.List(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
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
        l := yylex.(*lexer)        
        l.builder.MaxElements(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }
    | kywd_max_elements kywd_unbounded token_semi {
        l := yylex.(*lexer)
        l.builder.UnBounded(l.stack.peek(), true)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

min_elements : 
    kywd_min_elements int_value token_semi {
        l := yylex.(*lexer)
        l.builder.MinElements(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

list_body_stmt :
    description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | when_stmt
    | max_elements
    | min_elements
    | config_stmt
    | mandatory_stmt
    | key_stmt
    | unique_stmt
    | ordered_by_stmt
    | body_stmt


ordered_by_stmt :  
    kywd_ordered_by kywd_system statement_end {
        l := yylex.(*lexer)
        l.builder.OrderedBy(l.stack.peek(), meta.OrderedBySystem)
        if chkErr2(l, "ordered-by", $3) {
            goto ret1
        }
    }
    | kywd_ordered_by kywd_user statement_end {
        l := yylex.(*lexer)
        l.builder.OrderedBy(l.stack.peek(), meta.OrderedByUser)
        if chkErr2(l, "ordered-by", $3) {
            goto ret1
        }
    }

key_stmt: 
    kywd_key string_value statement_end {
        l := yylex.(*lexer)
        l.builder.Key(l.stack.peek(), $2)
        if chkErr2(l, "key", $3) {
            goto ret1
        }
    }

unique_stmt:    
    kywd_unique string_value token_semi

anyxml_stmt:
    anyxml_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | anyxml_def token_curly_open anyxml_body token_curly_close {
        yylex.(*lexer).stack.pop()
    }

anyxml_body :
    /* empty */
    description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | must_stmt
    | when_stmt
    | config_stmt
    | mandatory_stmt
    | unknown_stmt

anyxml_def :
    kywd_anyxml token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Any(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }
    | kywd_anydata token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Any(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

leaf_stmt:
    leaf_def token_curly_open optional_leaf_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
     }

leaf_def :
    kywd_leaf token_ident {
        l := yylex.(*lexer)        
        l.stack.push(l.builder.Leaf(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

optional_leaf_body_stmts:
	/* empty */
	leaf_body_stmts

leaf_body_stmts :
    leaf_body_stmt
    | leaf_body_stmts leaf_body_stmt


/* some are leaf-list only but builder will surface issues */
leaf_body_stmt :
    type_stmt
    | description
    | status_stmt
    | reference_stmt
    | if_feature_stmt
    | must_stmt
    | when_stmt
    | units_stmt
    | config_stmt
    | ordered_by_stmt
    | max_elements
    | min_elements
    | mandatory_stmt
    | default_stmt
    | unknown_stmt

mandatory_stmt : 
    kywd_mandatory bool_value token_semi {
        l := yylex.(*lexer)        
        l.builder.Mandatory(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
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
    | token_string {
        s := trimQuotes($1)
        n, err := strconv.ParseInt(s, 10, 32)
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
        l := yylex.(*lexer)
        l.builder.Config(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

leaf_list_stmt :
    leaf_list_def
    token_curly_open
    optional_leaf_body_stmts
    token_curly_close {
        yylex.(*lexer).stack.pop()
    }

leaf_list_def :
    kywd_leaf_list token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.LeafList(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

bit_stmt :
    bit_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | bit_def token_curly_open bit_body_stmts token_curly_close {
        yylex.(*lexer).stack.pop()
    }

bit_def :
    kywd_bit token_ident {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Bit(l.stack.peek(), $2))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }        
    }

bit_body_stmts :
    bit_body_stmt | bit_body_stmts bit_body_stmt

bit_body_stmt :
    description 
    | status_stmt
    | reference_stmt
    | position
    | unknown_stmt    

position :
    kywd_position int_value statement_end {
        l := yylex.(*lexer)
        l.builder.Position(l.stack.peek(), $2)
        if chkErr2(l, "position", $3) {
            goto ret1
        }
    }

enum_stmt :
    enum_def token_semi {
        yylex.(*lexer).stack.pop()
    }
    | enum_def token_curly_open enum_body_stmts token_curly_close {        
        yylex.(*lexer).stack.pop()
    }

enum_def : 
    kywd_enum token_string {
        l := yylex.(*lexer)
        l.stack.push(l.builder.Enum(l.stack.peek(), trimQuotes($2)))
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

enum_body_stmts :
    enum_body_stmt | enum_body_stmts enum_body_stmt

enum_body_stmt :
    description 
    | status_stmt
    | reference_stmt
    | enum_value
    | unknown_stmt

enum_value :
    kywd_value int_value statement_end {
        l := yylex.(*lexer)
        l.builder.EnumValue(l.stack.peek(), $2)
        if chkErr2(l, "value", $3) {
            goto ret1
        }
    }

description : 
    kywd_description string_value statement_end {
        l := yylex.(*lexer)
        l.builder.Description(l.stack.peek(), $2)     
        if chkErr2(l, "description", $3) {
            goto ret1
        }
    }

reference_stmt :
    kywd_reference string_value token_semi {
        l := yylex.(*lexer)
        l.builder.Reference(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

contact_stmt :
    kywd_contact string_value token_semi {
        l := yylex.(*lexer)        
        l.builder.Contact(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

organization_stmt :
    kywd_organization string_value token_semi {
        l := yylex.(*lexer)
        l.builder.Organization(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

yang_ver_stmt : 
    kywd_yang_version token_string token_semi {
        l := yylex.(*lexer)
        l.builder.YangVersion(l.stack.peek(), $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
    }

units_stmt :
    kywd_units token_string statement_end {        
        l := yylex.(*lexer)        
        l.builder.Units(l.stack.peek(), $2)
        if chkErr2(l, "units", $3) {
            goto ret1
        }
    }

statement_end :
    token_semi {
        $$ = nil
    }
    | token_curly_open unknown_stmt token_curly_close {
        $$ = $2
    }

unknown_stmt :
    token_unknown optional_unknown_args statement_end {              
        l := yylex.(*lexer)
        $$ = l.builder.Extension($1, $2)
        if chkErr(yylex, l.builder.LastErr) {
            goto ret1
        }
        // ironcically keyword extensions have have primary extensions
        if $3 != nil {
            l.builder.AddExtension($$, "", $3)
        }
        l.builder.AddExtension(l.stack.peek(), "", $$)
    }

optional_unknown_args:
	/* empty */ {
        $$ = []string{}
    }
    | unknown_args

unknown_args :
    string_or_number {
        $$ = []string{$1}
    }
    | unknown_args string_or_number {
        $$ = append($1, $2)
    }
%%

