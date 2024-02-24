package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vitzeno/llvm-test/parser"
)

func main() {
	sourceFile := flag.String("in", "", "input file")
	flag.Parse()

	if !isFlagPassed("in") {
		fmt.Println("input file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	parser.YYDebug = 0
	file, err := os.Open(*sourceFile)
	if err != nil {
		panic(err)
	}

	lexer := parser.NewLexer(file)
	errCount := parser.YYParse(lexer)
	if errCount != 0 {
		fmt.Println("parsing failed found error(s) in source file")
		os.Exit(1)
	}

	for _, lexErr := range lexer.Errors() {
		fmt.Println(lexErr)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
