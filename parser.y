%{

package main

%}

%union {
        num int
}

%type <num> expr

%token <num> NUM

%%

expr:
	NUM
        {
                $$ = $1
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }


%%
