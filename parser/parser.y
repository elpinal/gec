%{

package parser

import (
        "github.com/elpinal/gec/ast"
        "github.com/elpinal/gec/token"
)

%}

%union {
        top   *ast.WithDecls
        decl  *ast.Decl
        decls []*ast.Decl
        expr  ast.Expr
        token token.Token
        cmpop ast.CmpOp
}

%type <top>   top program withdecls withoutdecls
%type <expr>  expr term factor factor1 atom topexpr abs cmpexpr
%type <decl>  decl
%type <decls> decls
%type <cmpop> cmpop

%token <token> ILLEGAL

%token <token> NEWLINE
%token <token> NUM
%token <token> IDENT
%token <token> SYMBOL

%token <token> BOOL
%token <token> IF
%token <token> THEN
%token <token> ELSE

%token <token> RARROW
%token <token> EQ
%token <token> NE
%token <token> LE
%token <token> GE

%%

program:
	margin top margin
        {
                $$ = $2
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }

margin: /* empty */ | margin NEWLINE

top: withoutdecls | withdecls

withoutdecls:
        topexpr
        {
                $$ = &ast.WithDecls{Expr: $1}
        }

withdecls:
        decls newlines topexpr
        {
                $$ = &ast.WithDecls{Decls: $1, Expr: $3}
        }

newlines: NEWLINE margin

decls:
        decls newlines decl
        {
                $$ = append($1, $3)
        }
        | decl
        {
                $$ = []*ast.Decl{$1}
        }

decl:
        IDENT '=' topexpr
        {
                $$ = &ast.Decl{LHS: $1, RHS: $3}
        }
        | IDENT SYMBOL IDENT '=' topexpr
        {
                f := &ast.Abs{Param: $3, Body: $5}
                g := &ast.Abs{Param: $1, Body: f}
                $$ = &ast.Decl{LHS: $2, RHS: g}
        }

topexpr:
        abs
        | cmpexpr
        | IF cmpexpr THEN cmpexpr ELSE cmpexpr
        {
                $$ = &ast.If{Cond: $2, E1: $4, E2: $6}
        }

abs:
        '\\' IDENT RARROW topexpr
        {
                $$ = &ast.Abs{Param: $2, Body: $4}
        }

cmpexpr:
        expr
        | expr cmpop expr
        {
                $$ = &ast.Cmp{Op: $2, LHS: $1, RHS: $3}
        }

cmpop:
        EQ
        {
                $$ = ast.Eq
        }
        | NE
        {
                $$ = ast.NE
        }
        | '<'
        {
                $$ = ast.LT
        }
        | '>'
        {
                $$ = ast.GT
        }
        | LE
        {
                $$ = ast.LE
        }
        | GE
        {
                $$ = ast.GE
        }

expr:
        term
        | expr '+' term
        {
                $$ = &ast.Add{X: $1, Y: $3}
        }
        | expr '-' term
        {
                $$ = &ast.Sub{X: $1, Y: $3}
        }

term:
        factor1
        | term '*' factor1
        {
                $$ = &ast.Mul{X: $1, Y: $3}
        }
        | term '/' factor1
        {
                $$ = &ast.Div{X: $1, Y: $3}
        }

factor1:
        factor
        | factor1 SYMBOL factor
        {
                a := &ast.App{Fn: $2, Arg: $1}
                $$ = &ast.App{Fn: a, Arg: $3}
        }

factor:
        atom
        | factor atom
        {
                $$ = &ast.App{Fn: $1, Arg: $2}
        }

atom:
        NUM
        {
                $$ = &ast.Int{X: $1}
        }
        | IDENT
        {
                $$ = &ast.Ident{Name: $1}
        }
        | BOOL
        {
                $$ = &ast.Bool{X: $1}
        }
        | '(' topexpr ')'
        {
                $$ = &ast.ParenExpr{X: $2}
        }
        | '[' ']'
        {
                $$ = &ast.NilList{}
        }

%%
