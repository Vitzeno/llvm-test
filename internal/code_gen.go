package internal

import (
	"fmt"
	"io"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

var module = ir.NewModule()
var entry = module.NewFunc("main", types.I32)
var block = entry.NewBlock("")

func (l *Lexer) codeGen(a ast) {
	switch e := a.(type) {
	case *binaryExpr:
		lhs := constant.NewFloat(types.Double, l.eval(e.lhs))
		rhs := constant.NewFloat(types.Double, l.eval(e.rhs))
		switch e.Op {
		case "+":
			block.NewFAdd(lhs, rhs)
		case "-":
			block.NewFSub(lhs, rhs)
		case "*":
			block.NewFMul(lhs, rhs)
		case "/":
			if l.eval(e.rhs) == 0 {
				l.Error("Division by zero")
				return
			}
			block.NewFDiv(lhs, rhs)
		default:
			panic("unknown operator")
		}
	default:
		panic("unknown node type")
	}
}

func WriteToFile(w io.Writer) {
	block.NewRet(constant.NewInt(types.I32, 0))
	module.WriteTo(w)
}

func PrintCode() {
	block.NewRet(constant.NewFloat(types.Double, 0))
	fmt.Printf("module.String(): %v\n", module.String())
}
