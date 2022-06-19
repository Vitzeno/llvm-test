package internal

type Token int

const EOF = 0

const (
	tokenNumber     = NUMBER     // integer literal
	tokenIdentifier = IDENTIFIER // programmer defined names
	tokenSeparator  = SEPARATOR  // separator e.g. semicolon
	tokenAssign     = ASSIGN     // assignment operator
	tokenLet        = LET        // let keyword
	tokenIf         = IF         // if keyword
	tokenThen       = THEN       // then keyword
	tokenElse       = ELSE       // else keyword
	tokenLe         = LE         // <
	tokenGe         = GE         // >
	tokenEq         = EQ         // ==
	tokenNe         = NE         // !=
	tokenOr         = OR         // ||
	tokenAnd        = AND        // &&
	tokenWhile      = WHILE      // while keyword
)
