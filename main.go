package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vitzeno/llvm-test/internal"
)

func main() {
	sourceFile := flag.String("in", "", "input file")
	flag.Parse()

	if !isFlagPassed("in") {
		fmt.Println("input file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//internal.InitCodeGen()
	internal.YYDebug = 0
	file, err := os.Open(*sourceFile)
	if err != nil {
		panic(err)
	}

	lexer := internal.NewLexer(file)
	errCount := internal.YYParse(lexer)
	if errCount != 0 {
		fmt.Println("parsing failed found error(s) in source file")
		os.Exit(1)
	}

	for _, lexErr := range lexer.Errors() {
		fmt.Println(lexErr)
	}

	// _, err = os.Create(fmt.Sprintf("%s/%s.ll", path.Dir(*sourceFile), path.Base(*sourceFile)))
	// if err != nil {
	// 	panic(err)
	// }

	//internal.PrintCode()
	//internal.WriteToFile(outputFile)
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
