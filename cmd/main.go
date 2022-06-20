package main

import (
	"os"

	"github.com/vitzeno/llvm-test/internal"
)

func main() {
	internal.YYDebug = 0
	file, err := os.Open("testdata/mock.test")
	if err != nil {
		panic(err)
	}

	lexer := internal.NewLexer(file)
	internal.YYParse(lexer)
}
