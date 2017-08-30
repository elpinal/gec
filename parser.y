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
        exprs []ast.Expr
        token token.Token
}

%type <top> top
%type <expr> expr term factor atom
%type <exprs> atoms
%type <decl> decl
%type <decls> decls

%token <token> ILLEGAL NUM IDENT

%%

top:
        expr
        {
                $$ = &ast.WithDecls{Expr: $1}
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
                $$ = []*ast.Decl{$1}
        }

decl:
        IDENT '=' expr
        {
                $$ = &ast.Decl{LHS: $1, RHS: $3}
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
|	IDENT atoms
        {
                $$ = &ast.App{FnName: $1, Args: $2}
        }

atoms:
        atoms atom
        {
                $$ = append($1, $2)
        }
|	atom
        {
                $$ = []ast.Expr{$1}
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

%%
