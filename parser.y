%{

package main

import "github.com/elpinal/gec/ast"

%}

%union {
        expr ast.Expr
        num int
}

%type <expr> top expr term factor

%token <num> NUM

%%

top:
        expr
        {
                $$ = $1
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }

expr:
        term
        {
                $$ = $1
        }
|	expr '+' term
        {
                $$ = &ast.Add{X: $1, Y: $3}
        }
|	expr '-' term
        {
                $$ = &ast.Sub{X: $1, Y: $3}
        }

term:
        factor
        {
                $$ = $1
        }
|	term '*' factor
        {
                $$ = &ast.Mul{X: $1, Y: $3}
        }
|	term '/' factor
        {
                $$ = &ast.Div{X: $1, Y: $3}
        }

factor:
        NUM
        {
                $$ = &ast.Int{X: $1}
        }

%%
