package main

import (
	"os"
)

func main() {
	yyDebug = 0
	file, err := os.Open("testdata/input.test")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	yyParse(lexer)
	// if output != 0 {
	// 	panic("parse error")
	// }

	//spew.Dump(lexer.parseResult)
	if !lexer.evalFailed {
		lexer.eval(lexer.parseResult)
	}

}
