package main

import (
	"os"
)

func main() {
	yyDebug = 4
	file, err := os.Open("testdata/input.test")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	yyParse(lexer)
}
