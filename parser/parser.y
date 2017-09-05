%{

package parser

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

%type <top> top program
%type <expr> expr term factor atom absexpr abs cmpexpr
%type <decl> decl
%type <decls> decls

%token <token> ILLEGAL NEWLINE NUM IDENT RARROW BOOL IF THEN ELSE EQ NE LE GE SYMBOL

%%

program:
	margin top margin
        {
                $$ = $2
        }

top:
        absexpr
        {
                $$ = &ast.WithDecls{Expr: $1}
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }
|	decls newlines absexpr
        {
                $$ = &ast.WithDecls{Decls: $1, Expr: $3}
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }

margin:
|	margin NEWLINE

newlines:
        NEWLINE margin

decls:
        decls newlines decl
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
|	IDENT SYMBOL IDENT '=' absexpr
        {
                f := &ast.Abs{Param: $3, Body: $5}
                g := &ast.Abs{Param: $1, Body: f}
                $$ = &ast.Decl{LHS: $2, RHS: g}
        }

absexpr:
        abs
        {
                $$ = $1
        }
|	cmpexpr
        {
                $$ = $1
        }
|	IF cmpexpr THEN cmpexpr ELSE cmpexpr
        {
                $$ = &ast.If{Cond: $2, E1: $4, E2: $6}
        }

cmpexpr:
        expr
|	expr EQ expr
        {
                $$ = &ast.Cmp{Op: ast.Eq, LHS: $1, RHS: $3}
        }
|	expr NE expr
        {
                $$ = &ast.Cmp{Op: ast.NE, LHS: $1, RHS: $3}
        }
|	expr '<' expr
        {
                $$ = &ast.Cmp{Op: ast.LT, LHS: $1, RHS: $3}
        }
|	expr '>' expr
        {
                $$ = &ast.Cmp{Op: ast.GT, LHS: $1, RHS: $3}
        }
|	expr LE expr
        {
                $$ = &ast.Cmp{Op: ast.LE, LHS: $1, RHS: $3}
        }
|	expr GE expr
        {
                $$ = &ast.Cmp{Op: ast.GE, LHS: $1, RHS: $3}
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
|	'(' absexpr ')'
        {
                $$ = &ast.ParenExpr{X: $2}
        }
|	'[' ']'
        {
                $$ = &ast.NilList{}
        }

abs:
        '\\' IDENT RARROW absexpr
        {
                $$ = &ast.Abs{Param: $2, Body: $4}
        }

%%
