package main

type Token int

const (
	tokenEof = iota
	tokenIllegal
	tokenIdentifier     // programmer defined names
	tokenKeyword        // reserved keywords (e.g. "func", "if", "else", "return")
	tokenComment        // single line comment (e.g. #)
	tokenSeparator      // separators (e.g. "}", ")", ";")
	tokenAssignment     // assignment operator (e.g. "=", "*=", "/=")
	tokenBinaryOperator // binary operators (e.g. "+",  "==", "!=", "<", "<=", ">", ">=")
	tokenUnaryOperator  // unary operators (e.g. "!", "++", "--")
	tokenInt            // integer literal
)

var tokenNames = map[Token]string{
	tokenEof:            "EOF",
	tokenIllegal:        "illegal",
	tokenIdentifier:     "identifier",
	tokenKeyword:        "keyword",
	tokenComment:        "comment",
	tokenSeparator:      "separator",
	tokenAssignment:     "assignment",
	tokenBinaryOperator: "binary operator",
	tokenUnaryOperator:  "unary operator",
	tokenInt:            "integer",
}
