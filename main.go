package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testdata/input.test")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == tokenEof {
			fmt.Println("EOF")
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.col, tokenNames[tok], lit)
	}
}
