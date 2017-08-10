%{

package main

import "github.com/elpinal/gec/ast"

%}

%union {
        decl *ast.Assign
        decls []*ast.Assign
        expr ast.Expr
        num int
        ident string
}

%type <expr> top expr term factor
%type <decl> decl
%type <decls> decls

%token ILLEGAL
%token <num> NUM
%token <ident> IDENT

%%

top:
        expr
        {
                $$ = $1
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }
|	decls ';' expr
        {
                $$ = &ast.WithDecls{Decls: $1, Expr: $3}
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }

decls:
        decls ';' decl
        {
                $$ = append($1, $3)
        }
|       decl
        {
                $$ = []*ast.Assign{$1}
        }

decl:
        IDENT '=' expr
        {
                $$ = &ast.Assign{LHS: $1, RHS: $3}
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
