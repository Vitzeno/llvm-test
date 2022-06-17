%{
package main

import (
        "fmt"
)
%}

%union{
String string
Number float64
}


%token<String> NUMBER IDENTIFIER

%type <Number> expr

// operator precedence
%left '+' '-'
%left '*' '/'

%%
start: expr {{
            fmt.Println($1)
     }}
     | assignment;

expr:
      NUMBER {}
    | IDENTIFIER {}
    | expr '+' expr { $$ = $1 + $3 }
    | expr '-' expr { $$ = $1 - $3 }
    | expr '*' expr { $$ = $1 * $3 }
    | expr '/' expr { $$ = $1 / $3 }
    | '(' expr ')'  { $$ = $2 }
    | '-' expr %prec '*' { $$ = -$2 }
    ;

assignment:
          IDENTIFIER '=' expr {};
%%
