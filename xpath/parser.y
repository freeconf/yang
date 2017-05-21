%{
package xpath

import (
    "fmt"
    "github.com/c2stack/c2g/c2"
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
    msg := fmt.Sprintf("%s - col %d", e, l.pos)
    l.lastError = c2.NewErr(msg)
}

%}

%union {
 token string
 stack *stack
}

%token <token> token_name
%token <token> token_literal
%token <token> token_number
%token <token> token_operator

%token kywd_slash

%%

path :
    relative_path
    | absolute_path

absolute_path :
    kywd_slash relative_path {
        abs := &AbsolutePath{}
        abs.Append(yyVAL.stack.pop())
        yyVAL.stack.push(abs)
    }

relative_path :
    step
    | relative_path step { 
         p := yyVAL.stack.pop()
         yyVAL.stack.peek().Append(p)
    }

step :
    stmt kywd_slash 
    | stmt

stmt : 
    token_name {
        yyVAL.stack.push(&Segment{Ident:$1})
    }
    | token_name token_operator token_number {
        n, err := num($3)
        if err != nil {
            yylex.(*lexer).lastError = err
            goto ret1
        }
        yyVAL.stack.push(&Segment{Ident:$1, Expr: &Operator{Oper:$2, Lhs:n}})
    }
    | token_name token_operator token_literal {
        yyVAL.stack.push(&Segment{Ident:$1, Expr: &Operator{Oper:$2, Lhs:literal($3)}})
    }
