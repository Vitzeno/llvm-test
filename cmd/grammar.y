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


%token<String> NUMBER IDENTIFIER SEPARATOR ASSIGN LET IF THEN LE GE EQ NE OR AND ELSE

%type <Ast> statements statement expression assignment control_flow

%nonassoc NO_ELSE
%nonassoc ELSE
%right ASSIGN
%left LE GE EQ NE OR AND
%left '+' '-'
%left '*' '/'
%right UMINUS
%left '(' ')'

%start program

%%
program :  /* empty */
     | program statement { fmt.Println(yylex.(*Lexer).eval($2)) } 
     ;

statement:
     expression SEPARATOR
     | assignment SEPARATOR
     | control_flow
     ;

statements:
     statement statements
     | statement
     ;

expression:
     NUMBER { $$ = &number{$1} }
    | IDENTIFIER { $$ = &variable{$1} }
    | expression '+' expression { $$ = &binaryExpr{Op: "+", lhs: $1, rhs: $3} }
    | expression '-' expression { $$ = &binaryExpr{Op: "-", lhs: $1, rhs: $3} }
    | expression '*' expression { $$ = &binaryExpr{Op: "*", lhs: $1, rhs: $3} }
    | expression '/' expression { $$ = &binaryExpr{Op: "/", lhs: $1, rhs: $3} }
    | expression LE expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | expression GE expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | expression EQ expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | expression NE expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | expression OR expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | expression AND expression { $$ = &binaryExpr{Op: $2, lhs: $1, rhs: $3} }
    | '(' expression ')'  { $$ = &parenExpr{$2} }
    | '-' expression %prec UMINUS { $$ = &unaryExpr{$2} }
    ;
    
assignment:
     LET IDENTIFIER ASSIGN expression { $$ = &assignment{variable: $2, expr: $4} }
     ;

control_flow:
     IF '(' expression ')' '{' statements '}'  %prec NO_ELSE { $$ = &ifStatement{cond: $3, thenStmt: $6, elseStmt: nil} }
     | IF '(' expression ')' '{' statements '}' ELSE '{' statements '}' { $$ = &ifStatement{cond: $3, thenStmt: $6, elseStmt: $10} }
     ;

%%
