package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

//go generate goyacc -o parser.go -v parser.output grammar.y

func main() {
	yyDebug = 0
	file, err := os.Open("testdata/input.test")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	yyParse(lexer)

	spew.Dump(lexer.parseResult)
	if !lexer.evalFailed {
		lexer.eval(lexer.parseResult)
	}

}
