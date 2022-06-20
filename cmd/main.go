package main

import (
	"fmt"
	"os"
	"path"

	"github.com/vitzeno/llvm-test/internal"
)

const sourceFile = "testdata/mock"

func main() {
	internal.YYDebug = 0
	file, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}

	lexer := internal.NewLexer(file)
	internal.YYParse(lexer)

	outputFile, err := os.Create(fmt.Sprintf("%s/%s.ll", path.Dir(sourceFile), path.Base(sourceFile)))
	if err != nil {
		panic(err)
	}

	internal.PrintCode()
	internal.WriteToFile(outputFile)
}
