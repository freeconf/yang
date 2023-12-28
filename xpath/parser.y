%{
package xpath

import (
    "fmt"
)

func (l *lexer) Lex(lval *yySymType) int {
    t, _ := l.nextToken()
    if t.typ == ParseEnd {
        return 0
    }
    lval.token = t.val
    lval.stack = l.stack
    return int(t.typ)
}

func (l *lexer) Error(e string) {
    l.lastError = fmt.Errorf("%s - col %d", e, l.pos)
}

%}

%union {
 token string
 path *Path
 stack *stack
}

%token <token> token_name
%token <token> token_literal
%token <token> token_number
%token <token> token_operator

%token kywd_slash
%token kywd_colon

%type <path> qname 

%%

segments :
    segment
    | segments segment

segment :
    stmt
    | stmt kywd_slash 

qname :
    token_name {
        $$ = &Path{Ident:$1}
    }    
    | token_name kywd_colon token_name {
        l := yylex.(*lexer)
        m, err := l.lookup($1)
        if err != nil {
            l.lastError = err
            goto ret1
        }
        $$ = &Path{Module:m.Ident(), Ident: $3}
    }    

stmt : 
    qname {
        yyVAL.stack.push($1)
    }
    | qname token_operator token_number {
        n, err := num($3)
        if err != nil {
            yylex.(*lexer).lastError = err
            goto ret1
        }
        $1.Expr = &Operator{Oper:$2, Lhs:n}
        yyVAL.stack.push($1)
    }
    | qname token_operator token_literal {
        $1.Expr = &Operator{Oper:$2, Lhs:literal($3)}
        yyVAL.stack.push($1)
    }
