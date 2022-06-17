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
	//spew.Dump(l.variables)
	switch e := e.(type) {
	case *binaryExpr:
		fmt.Printf("binaryExpr: %c\n", e.Op)
		switch e.Op {
		case '+':
			return l.eval(e.lhs) + l.eval(e.rhs)
		case '-':
			return l.eval(e.lhs) - l.eval(e.rhs)
		case '*':
			return l.eval(e.lhs) * l.eval(e.rhs)
		case '/':
			return l.eval(e.lhs) / l.eval(e.rhs)
		default:
			panic("unknown operator")
		}

	case *unaryExpr:
		//fmt.Printf("unary: %v\n", l.eval(e.expr))
		return -l.eval(e.expr)

	case *astRoot:
		//fmt.Printf("%f\n", l.eval(e.expr))
		result := l.eval(e.expr)
		if !l.evalFailed {
			fmt.Println(result)
		}
		return result

	case *parenExpr:
		//fmt.Printf("%T\n", e.expr)
		return l.eval(e.expr)

	case *variable:
		//fmt.Printf("variable %s\n", e.name)
		val, ok := l.variables[e.name]
		if !ok {
			l.Error(fmt.Sprintf("undefined variable: %s", e.name))
		}
		return val

	case *number:
		//fmt.Printf("number: %s\n", e.value)
		var err error
		val, err := strconv.ParseFloat(e.value, 64)
		if err != nil {
			l.Error("invalid number")
		}
		return val

	case *assignment:
		//fmt.Printf("assignment: %s = %v\n", e.variable, l.eval(e.expr))
		result := l.eval(e.expr)
		if !l.evalFailed {
			l.variables[e.variable] = result
		}
		return result

	default:
		panic("unknown node type")
	}
}
