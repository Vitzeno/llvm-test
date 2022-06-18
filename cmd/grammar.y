%{
package main

import (
     "fmt"
)

%}

%union{
String string
Ast ast 
}


%token<String> NUMBER IDENTIFIER SEPARATOR LET IF THEN LE GE EQ NE OR AND ELSE

%type <Ast> expressions expression assignment 

%right '='
%left '<' '>' LE GE EQ NE
%left '+' '-'
%left '*' '/'
%right UMINUS
%left '(' ')'

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
    | '-' expression %prec UMINUS { $$ = &unaryExpr{$2} }
    ;

assignment:
          LET IDENTIFIER '=' expression { $$ = &assignment{variable: $2, expr: $4} }
     ;
%%
