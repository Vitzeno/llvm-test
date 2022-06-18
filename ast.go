package main

import (
	"fmt"
	"strconv"
)

type expr interface{}

type astRoot struct {
	expr expr
}

type binaryExpr struct {
	Op  byte
	lhs expr
	rhs expr
}

type unaryExpr struct {
	expr expr
}

type parenExpr struct {
	expr expr
}

type variable struct {
	name string
}

type number struct {
	value string
}

type assignment struct {
	variable string
	expr     expr
}

func (l *Lexer) eval(e expr) float64 {
	//spew.Dump(e)
	switch e := e.(type) {
	case *binaryExpr:
		switch e.Op {
		case '+':
			return l.eval(e.lhs) + l.eval(e.rhs)
		case '-':
			return l.eval(e.lhs) - l.eval(e.rhs)
		case '*':
			return l.eval(e.lhs) * l.eval(e.rhs)
		case '/':
			if l.eval(e.rhs) == 0 {
				l.Error("Division by zero")
				return 0
			}
			return l.eval(e.lhs) / l.eval(e.rhs)
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

	default:
		panic("unknown node type")
	}
}

func (l *Lexer) printAstNode(e expr) {
	switch e := e.(type) {
	case *binaryExpr:
		l.printAstNode(e.lhs)
		fmt.Printf(" %c ", e.Op)
		l.printAstNode(e.rhs)
	case *unaryExpr:
		fmt.Printf("(%c", '-')
		l.printAstNode(e.expr)
		fmt.Printf(")")
	case *astRoot:
		l.printAstNode(e.expr)
	case *parenExpr:
		fmt.Printf("(")
		l.printAstNode(e.expr)
		fmt.Printf(")")
	case *variable:
		fmt.Printf("%s", e.name)
	case *number:
		fmt.Printf("%s", e.value)
	case *assignment:
		fmt.Printf("%s = ", e.variable)
		l.printAstNode(e.expr)
	default:
		panic("unknown node type")
	}
}
