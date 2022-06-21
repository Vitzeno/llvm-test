package internal

import (
	"fmt"
	"io"
	"runtime"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

var (
	module *ir.Module
	printf *ir.Func
	entry  *ir.Func
	block  *ir.Block
)

func InitCodeGen() {
	module = ir.NewModule()
	os := runtime.GOOS
	switch os {
	case "windows":
		module.TargetTriple = "x86_64-pc-windows-msvc"
	case "darwin":
		module.TargetTriple = "x86_64-apple-darwin"
	case "linux":
		module.TargetTriple = "x86_64-pc-linux-gnu"
	default:
		panic("unknown OS target")
	}

	entry = module.NewFunc("main", types.I32)
	block = entry.NewBlock("main")
	block.NewRet(constant.NewInt(types.I32, 0))

	printf = module.NewFunc("printf", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
}

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
		case ">":
			block.NewFCmp(enum.FPredOGE, lhs, rhs)
		case "<":
			block.NewFCmp(enum.FPredOLE, lhs, rhs)
		case "==":
			block.NewFCmp(enum.FPredOEQ, lhs, rhs)
		case "NE":
			block.NewFCmp(enum.FPredONE, lhs, rhs)
		case "OR":
			block.NewOr(lhs, rhs)
		case "AND":
			block.NewAnd(lhs, rhs)
		default:
			panic("unknown operator")
		}

	case *unaryExpr:
		// fallrthrough to eval
		l.eval(e.expr)

	case *astRoot:
		l.codeGen(e.expr)

	case *parenExpr:
		// fallrthrough to eval
		l.eval(e.expr)

	case *variable:
		_, ok := l.variables[e.name]
		if !ok {
			l.Error(fmt.Sprintf("undefined variable: %s", e.name))
		}

	case *number:
		var err error
		_, err = strconv.ParseFloat(e.value, 64)
		if err != nil {
			l.Error("invalid number")
		}

	case *assignment:
		_, ok := l.variables[e.variable]
		if ok {
			l.Error(fmt.Sprintf("variable already defined: %s", e.variable))
		}

		result := l.eval(e.expr)
		if !l.evalFailed {
			variable := block.NewAlloca(types.Double)
			variable.SetName(e.variable)
			block.NewStore(constant.NewFloat(types.Double, result), variable)
			l.variables[e.variable] = varible{location: variable, value: result}
		}

	case *reassignment:
		_, ok := l.variables[e.variable]
		if !ok {
			l.Error(fmt.Sprintf("undefined variable: %s", e.variable))
		}

		result := l.eval(e.expr)
		if !l.evalFailed {
			block.NewStore(constant.NewFloat(types.Double, result), l.variables[e.variable].location)
			l.variables[e.variable] = varible{location: l.variables[e.variable].location, value: result}
		}

	case *stdPrint:
		result := l.eval(e.expr)
		if !l.evalFailed {
			variable := block.NewAlloca(types.I8)
			variable.SetName("msg")
			block.NewCall(printf, variable, constant.NewInt(types.I32, int64(result)))
		}

	case *ifStatement:
		//TODO: proper implementation
		condStmt := constant.NewInt(types.I1, int64(l.eval(e.cond)))

		iftrue := entry.NewBlock("")
		iffalse := entry.NewBlock("")
		ifend := entry.NewBlock("")

		block.NewCondBr(condStmt, iftrue, iffalse)

		iftrue.NewBr(ifend)
		// add instructions to iftrue

		iffalse.NewBr(ifend)
		// add instructions to iffalse

		ifend.NewRet(constant.NewInt(types.I32, 0))

	case *whileStatement:
		//TODO: proper implementation
		initBlock := block
		whileCond := entry.NewBlock("while.cond")
		block = entry.NewBlock("while.body")
		if l.eval(e.cond) != 0 {
			l.codeGen(e.body)
			block.NewBr(whileCond)
		}
		block = initBlock

	default:
		panic("unknown node type")
	}
}

func WriteToFile(w io.Writer) {
	module.WriteTo(w)
}

func PrintCode() {
	fmt.Printf("%v\n", module.String())
}
