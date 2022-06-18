%{
package main

import (
     "fmt"
)

%}

%union{
String string
Expr expr 
}


%token<String> NUMBER IDENTIFIER SEPARATOR LET

%type <Expr> expressions expression assignment 

%left '+' '-'
%left '*' '/'

%start program

%%
program :  /* empty */
     | program expressions SEPARATOR { fmt.Println(yylex.(*Lexer).eval($2)) } 
     ;

expressions:
     expression  
     | assignment 
     ;

expression:
     NUMBER { $$ = &number{$1} }
    | IDENTIFIER { $$ = &variable{$1} }
    | expression '+' expression { $$ = &binaryExpr{Op: '+', lhs: $1, rhs: $3} }
    | expression '-' expression { $$ = &binaryExpr{Op: '-', lhs: $1, rhs: $3} }
    | expression '*' expression { $$ = &binaryExpr{Op: '*', lhs: $1, rhs: $3} }
    | expression '/' expression { $$ = &binaryExpr{Op: '/', lhs: $1, rhs: $3} }
    | '(' expression ')'  { $$ = &parenExpr{$2} }
    | '-' expression %prec '*' { $$ = &unaryExpr{$2} }
    ;

assignment:
          LET IDENTIFIER '=' expression { $$ = &assignment{variable: $2, expr: $4} }
     ;
%%
