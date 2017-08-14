%{

package main

import "github.com/elpinal/gec/ast"

%}

%union {
        top *ast.WithDecls
        decl ast.Decl
        decls []ast.Decl
        expr ast.Expr
        exprs []ast.Expr
        num int
        ident string
        args []string
}

%type <top> top
%type <expr> expr term factor atom
%type <exprs> atoms
%type <decl> decl
%type <decls> decls
%type <args> args

%token ILLEGAL
%token <num> NUM
%token <ident> IDENT

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
                $$ = []ast.Decl{$1}
        }

decl:
        IDENT '=' expr
        {
                $$ = &ast.Assign{LHS: $1, RHS: $3}
        }
|	IDENT args '=' expr
        {
                $$ = &ast.DeclFunc{Name: $1, Args: $2, RHS: $4}
        }

args:
        IDENT
        {
                $$ = []string{$1}
        }
|	args IDENT
        {
                $$ = append($1, $2)
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
