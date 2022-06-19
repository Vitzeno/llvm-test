package internal

import (
	"fmt"
	"strconv"
)

type ast interface{}

type astRoot struct {
	expr ast
}

type binaryExpr struct {
	Op  string
	lhs ast
	rhs ast
}

type unaryExpr struct {
	expr ast
}

type parenExpr struct {
	expr ast
}

type variable struct {
	name string
}

type number struct {
	value string
}

type assignment struct {
	variable string
	expr     ast
}

type ifStatement struct {
	cond     ast
	thenStmt ast
	elseStmt ast
}

type whileStatement struct {
	cond ast
	body ast
}

func (l *Lexer) eval(a ast) float64 {
	//spew.Dump(e)
	switch e := a.(type) {
	case *binaryExpr:
		switch e.Op {
		case "+":
			return l.eval(e.lhs) + l.eval(e.rhs)
		case "-":
			return l.eval(e.lhs) - l.eval(e.rhs)
		case "*":
			return l.eval(e.lhs) * l.eval(e.rhs)
		case "/":
			if l.eval(e.rhs) == 0 {
				l.Error("Division by zero")
				return 0
			}
			return l.eval(e.lhs) / l.eval(e.rhs)
		case ">":
			if l.eval(e.lhs) > l.eval(e.rhs) {
				return 1
			}
			return 0
		case "<":
			if l.eval(e.lhs) < l.eval(e.rhs) {
				return 1
			}
			return 0
		case "==":
			if l.eval(e.lhs) == l.eval(e.rhs) {
				return 1
			}
			return 0
		case "NE":
			if l.eval(e.lhs) != l.eval(e.rhs) {
				return 1
			}
			return 0
		case "OR":
			if l.eval(e.lhs) != 0 || l.eval(e.rhs) != 0 {
				return 1
			}
			return 0
		case "AND":
			if l.eval(e.lhs) != 0 && l.eval(e.rhs) != 0 {
				return 1
			}
			return 0
		default:
			panic("unknown operator")
		}

	case *unaryExpr:
		return -l.eval(e.expr)

	case *astRoot:
		result := l.eval(e.expr)
		if !l.evalFailed {
			fmt.Println(result)
		}
		return result

	case *parenExpr:
		return l.eval(e.expr)

	case *variable:
		val, ok := l.variables[e.name]
		if !ok {
			l.Error(fmt.Sprintf("undefined variable: %s", e.name))
		}
		return val

	case *number:
		var err error
		val, err := strconv.ParseFloat(e.value, 64)
		if err != nil {
			l.Error("invalid number")
		}
		return val

	case *assignment:
		result := l.eval(e.expr)
		if !l.evalFailed {
			l.variables[e.variable] = result
		}
		return result

	case *ifStatement:
		if l.eval(e.cond) != 0 {
			return l.eval(e.thenStmt)
		}
		if e.elseStmt != nil {
			return l.eval(e.elseStmt)
		}
		return 0

	case *whileStatement:
		for l.eval(e.cond) != 0 {
			l.eval(e.body)
		}
		return 0

	default:
		panic("unknown node type")
	}
}

// func (l *Lexer) printAstNode(a ast) {
// 	switch e := a.(type) {
// 	case *binaryExpr:
// 		l.printAstNode(e.lhs)
// 		fmt.Printf(" %s ", e.Op)
// 		l.printAstNode(e.rhs)
// 	case *unaryExpr:
// 		fmt.Printf("(%c", '-')
// 		l.printAstNode(e.expr)
// 		fmt.Printf(")")
// 	case *astRoot:
// 		l.printAstNode(e.expr)
// 	case *parenExpr:
// 		fmt.Printf("(")
// 		l.printAstNode(e.expr)
// 		fmt.Printf(")")
// 	case *variable:
// 		fmt.Printf("%s", e.name)
// 	case *number:
// 		fmt.Printf("%s", e.value)
// 	case *assignment:
// 		fmt.Printf("%s = ", e.variable)
// 		l.printAstNode(e.expr)
// 	default:
// 		panic("unknown node type")
// 	}
// }
