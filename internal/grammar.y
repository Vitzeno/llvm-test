%{
package internal

%}

%union{
String string
Ast ast 
}

%token<String> NUMBER IDENTIFIER SEPARATOR ASSIGN LET IF THEN LE GE EQ NE OR AND ELSE WHILE PRINT

%type <Ast> statements statement expression assignment reassignment print control_flow while_statement

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
     | program statement { YYlex.(*Lexer).codeGen($2) } 
     ;

statement:
     expression SEPARATOR
     | assignment SEPARATOR
     | reassignment SEPARATOR
     | print SEPARATOR
     | control_flow
     | while_statement
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

reassignment:
     IDENTIFIER ASSIGN expression { $$ = &reassignment{variable: $1, expr: $3} }
     ;

print:
     PRINT '(' expression ')' { $$ = &stdPrint{expr: $3} }
     ;

control_flow:
     IF '(' expression ')' '{' statements '}'  %prec NO_ELSE { $$ = &ifStatement{cond: $3, thenStmt: $6, elseStmt: nil} }
     | IF '(' expression ')' '{' statements '}' ELSE '{' statements '}' { $$ = &ifStatement{cond: $3, thenStmt: $6, elseStmt: $10} }
     ;

while_statement:
     WHILE '(' expression ')' '{' statements '}' { $$ = &whileStatement{cond: $3, body: $6} }

%%
