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
	"github.com/llir/llvm/ir/value"
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
	block = entry.NewBlock("entry")

	printf = module.NewFunc("printf", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
}

func (l *Lexer) compileConstant(a ast) value.Value {
	return constant.NewFloat(types.Double, l.eval(a))
}

func (l *Lexer) codeGen(a ast) value.Value {
	switch e := a.(type) {
	case *binaryExpr:
		lhs := l.compileConstant(e.lhs)
		rhs := l.compileConstant(e.lhs)
		switch e.Op {
		case "+":
			return block.NewFAdd(lhs, rhs)
		case "-":
			return block.NewFSub(lhs, rhs)
		case "*":
			return block.NewFMul(lhs, rhs)
		case "/":
			// this causes a block missing terminator error, might be better to handle this separately outside of the switch
			if l.eval(e.rhs) == 0 {
				l.Error("Division by zero")
				return nil
			}
			return block.NewFDiv(lhs, rhs)
		case ">":
			return block.NewFCmp(enum.FPredOGE, lhs, rhs)
		case "<":
			return block.NewFCmp(enum.FPredOLE, lhs, rhs)
		case "==":
			return block.NewFCmp(enum.FPredOEQ, lhs, rhs)
		case "NE":
			return block.NewFCmp(enum.FPredONE, lhs, rhs)
		case "OR":
			return block.NewOr(lhs, rhs)
		case "AND":
			return block.NewAnd(lhs, rhs)
		default:
			panic(fmt.Sprintf("code gen: unknown operator: %s", e.Op))
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
			l.Error(fmt.Sprintf("variable [%s] already defined", e.variable))
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
		thenB := entry.NewBlock("if.then")
		var elseB *ir.Block
		if e.elseStmt != nil {
			elseB = entry.NewBlock("if.else")
		} else {
			elseB = nil
		}

		cond := l.codeGen(e.cond)
		if !l.evalFailed {
			block.NewCondBr(cond, thenB, elseB)
		}
		block = thenB
		l.codeGen(e.thenStmt)
		if !l.evalFailed {
			block.NewBr(block)
		}

		if e.elseStmt != nil {
			block = elseB
			l.codeGen(e.elseStmt)
			if !l.evalFailed {
				block.NewBr(block)
			}
		}

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

	return nil
}

func WriteToFile(w io.Writer) {
	block.NewRet(constant.NewInt(types.I32, 0))
	module.WriteTo(w)
}

func PrintCode() {
	block.NewRet(constant.NewInt(types.I32, 0))
	fmt.Printf("%v\n", module.String())
}
