%{

package main

import "github.com/elpinal/gec/ast"

%}

%union {
        expr ast.Expr
        num int
}

%type <expr> expr term

%token <num> NUM

%%

expr:
	term
        {
                $$ = $1
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }

term:
        NUM
        {
                $$ = &ast.Int{X: $1}
        }
|	term '+' NUM
        {
                $$ = &ast.Add{X: $1, Y: &ast.Int{X: $3}}
        }


%%
