%{

package main

import (
        "github.com/elpinal/gec/ast"
        "github.com/elpinal/gec/token"
)

%}

%union {
        top *ast.WithDecls
        decl *ast.Decl
        decls []*ast.Decl
        expr ast.Expr
        token token.Token
}

%type <top> top
%type <expr> expr term factor atom absexpr abs
%type <decl> decl
%type <decls> decls

%token <token> ILLEGAL NUM IDENT RARROW BOOL IF THEN ELSE

%%

top:
        absexpr
        {
                $$ = &ast.WithDecls{Expr: $1}
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }
|	decls ';' absexpr
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
                $$ = []*ast.Decl{$1}
        }

decl:
        IDENT '=' absexpr
        {
                $$ = &ast.Decl{LHS: $1, RHS: $3}
        }

absexpr:
        abs
        {
                $$ = $1
        }
|	expr
        {
                $$ = $1
        }
|	IF expr THEN expr ELSE expr
        {
                $$ = &ast.If{Cond: $2, E1: $4, E2: $6}
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
        atom
        {
                $$ = $1
        }
|	factor atom
        {
                $$ = &ast.App{Fn: $1, Arg: $2}
        }

atom:
        NUM
        {
                $$ = &ast.Int{X: $1}
        }
|	IDENT
        {
                $$ = &ast.Ident{Name: $1}
        }
|	BOOL
        {
                $$ = &ast.Bool{X: $1}
        }

abs:
        '\\' IDENT RARROW expr
        {
                $$ = &ast.Abs{Param: $2, Body: $4}
        }

%%
