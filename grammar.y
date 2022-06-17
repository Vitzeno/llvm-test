%{
package main
%}

%union{
String string
Expr expr 
}


%token<String> NUMBER IDENTIFIER

%type <Expr> expr assignment

%left '+' '-'
%left '*' '/'

%%
start: expr {yylex.(*Lexer).parseResult = &astRoot{$1}} 
     | assignment {yylex.(*Lexer).parseResult = $1}
     ;

expr:
      NUMBER {$$ = &number{$1} }
    | IDENTIFIER { $$ = &variable{$1}}
    | expr '+' expr { $$ = &binaryExpr{Op: '+', lhs: $1, rhs: $3} }
    | expr '-' expr { $$ = &binaryExpr{Op: '-', lhs: $1, rhs: $3} }
    | expr '*' expr { $$ = &binaryExpr{Op: '*', lhs: $1, rhs: $3} }
    | expr '/' expr { $$ = &binaryExpr{Op: '/', lhs: $1, rhs: $3} }
    | '(' expr ')'  { $$ = &parenExpr{$2}}
    | '-' expr %prec '*' { $$ = &unaryExpr{$2} }
    ;

assignment:
          IDENTIFIER '=' expr {$$ = &assignment{$1, $3}};
%%
