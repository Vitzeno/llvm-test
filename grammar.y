%{
package main
%}

%union{
String string
Expr expr 
}


%token<String> NUMBER IDENTIFIER SEPARATOR LET

%type <Expr> expr assignment

%left '+' '-'
%left '*' '/'

%%
start : program
     | start program
     ;

program: 
     expr SEPARATOR { yylex.(*Lexer).parseResult = &astRoot{$1} } 
     | assignment SEPARATOR { yylex.(*Lexer).parseResult = &astRoot{$1} }
     ;

expr:
      NUMBER { $$ = &number{$1} }
    | IDENTIFIER { $$ = &variable{$1} }
    | expr '+' expr { $$ = &binaryExpr{Op: '+', lhs: $1, rhs: $3} }
    | expr '-' expr { $$ = &binaryExpr{Op: '-', lhs: $1, rhs: $3} }
    | expr '*' expr { $$ = &binaryExpr{Op: '*', lhs: $1, rhs: $3} }
    | expr '/' expr { $$ = &binaryExpr{Op: '/', lhs: $1, rhs: $3} }
    | '(' expr ')'  { $$ = &parenExpr{$2} }
    | '-' expr %prec '*' { $$ = &unaryExpr{$2} }
    ;

assignment:
          LET IDENTIFIER '=' expr { $$ = &assignment{variable: $2, expr: $4} }
     ;
%%
