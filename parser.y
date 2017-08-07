%{

package main

%}

%union {
        num int
}

%token NUM

%%

pattern:
	NUM
        {
                $$ = $1
        }


%%
