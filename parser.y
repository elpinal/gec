%{

package main

%}

%union {
        num int
}

%type <num> pattern

%token <num> NUM

%%

pattern:
	NUM
        {
                $$ = $1
                if l, ok := yylex.(*exprLexer); ok {
                        l.expr = $$
                }
        }


%%
