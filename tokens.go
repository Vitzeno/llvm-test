package main

type Token int

const (
	tokenEof = iota
	tokenIllegal
	tokenIdentifier     // programmer defined names
	tokenKeyword        // reserved keywords (e.g. "func", "if", "else", "return")
	tokenComment        // single line comment (e.g. #)
	tokenLeftParen      // (
	tokenRightParen     // )
	tokenLeftBrace      // {
	tokenRightBrace     // }
	tokenAssignment     // assignment operator (e.g. "=")
	tokenAddition       // addition operator (e.g. "+")
	tokenSubtraction    // subtraction operator (e.g. "-")
	tokenMultiplication // multiplication operator (e.g. "*")
	tokenDivision       // division operator (e.g. "/")
	tokenInt            // integer literal
)

var tokenNames = map[Token]string{
	tokenEof:            "EOF",
	tokenIllegal:        "ILLEGAL",
	tokenIdentifier:     "IDENTIFIER",
	tokenKeyword:        "KEYWORD",
	tokenComment:        "COMMENT",
	tokenLeftParen:      "LEFT_PAREN",
	tokenRightParen:     "RIGHT_PAREN",
	tokenLeftBrace:      "LEFT_BRACE",
	tokenRightBrace:     "RIGHT_BRACE",
	tokenAssignment:     "ASSIGNMENT",
	tokenAddition:       "ADDITION",
	tokenSubtraction:    "SUBTRACTION",
	tokenMultiplication: "MULTIPLICATION",
	tokenDivision:       "DIVISION",
	tokenInt:            "INT",
}
