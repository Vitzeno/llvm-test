package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type ast any

type root struct {
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

type reassignment struct {
	variable string
	expr     ast
}

type stdPrint struct {
	expr ast
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

func (l *Lexer) eval(a ast) (float64, error) {
	switch e := a.(type) {
	case *binaryExpr:
		return l.evalBinaryExpr(e)
	case *unaryExpr:
		return l.evalUnaryExpr(e)
	case *root:
		return l.evalRoot(e)
	case *parenExpr:
		return l.evalParenExpr(e)
	case *variable:
		return l.evalVariable(e)
	case *number:
		return l.evalNumber(e)
	case *assignment:
		return l.evalAssignment(e)
	case *reassignment:
		return l.evalReassignment(e)
	case *stdPrint:
		return l.evalStdPrint(e)
	case *ifStatement:
		return l.evalIfStatement(e)
	case *whileStatement:
		return l.evalWhileStatement(e)
	default:
		return 0, fmt.Errorf("unknown node type")
	}
}

func (l *Lexer) evalBinaryExpr(e *binaryExpr) (float64, error) {
	switch e.Op {
	case "+":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		return lhs + rhs, nil
	case "-":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		return lhs - rhs, nil
	case "*":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		return lhs * rhs, nil
	case "/":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if rhs == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return lhs / rhs, nil
	case ">":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs > rhs {
			return 1, nil
		}
		return 0, nil
	case "<":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs < rhs {
			return 1, nil
		}
		return 0, nil
	case "==":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs == rhs {
			return 1, nil
		}
		return 0, nil
	case "NE":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs != rhs {
			return 1, nil
		}
		return 0, nil
	case "OR":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs != 0 || rhs != 0 {
			return 1, nil
		}
		return 0, nil
	case "AND":
		lhs, err := l.eval(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := l.eval(e.rhs)
		if err != nil {
			return 0, err
		}

		if lhs != 0 && rhs != 0 {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", e.Op)
	}
}

func (l *Lexer) evalUnaryExpr(e *unaryExpr) (float64, error) {
	result, err := l.eval(e.expr)
	return -result, err
}

func (l *Lexer) evalRoot(e *root) (float64, error) {
	return l.eval(e.expr)
}

func (l *Lexer) evalParenExpr(e *parenExpr) (float64, error) {
	return l.eval(e.expr)
}

func (l *Lexer) evalVariable(e *variable) (float64, error) {
	val, ok := l.variables[e.name]
	if !ok {
		return 0, fmt.Errorf("undefined variable: %s", e.name)
	}
	return val.value, nil
}

func (l *Lexer) evalNumber(e *number) (float64, error) {
	val, err := strconv.ParseFloat(e.value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", e.value)
	}
	return val, nil
}

func (l *Lexer) evalAssignment(e *assignment) (float64, error) {
	_, ok := l.variables[e.variable]
	if ok {
		return 0, fmt.Errorf("variable already defined: %s", e.variable)
	}

	result, err := l.eval(e.expr)
	if err != nil {
		return 0, err
	}

	if !l.evalFailed {
		l.variables[e.variable] = varible{location: nil, value: result}
	}
	return result, nil
}

func (l *Lexer) evalReassignment(e *reassignment) (float64, error) {
	_, ok := l.variables[e.variable]
	if !ok {
		return 0, fmt.Errorf("undefined variable: %s", e.variable)
	}

	result, err := l.eval(e.expr)
	if err != nil {
		return 0, err
	}

	if !l.evalFailed {
		l.variables[e.variable] = varible{location: l.variables[e.variable].location, value: result}
	}
	return result, nil
}

func (l *Lexer) evalStdPrint(e *stdPrint) (float64, error) {
	result, err := l.eval(e.expr)
	if err != nil {
		return 0, err
	}
	if !l.evalFailed {
		fmt.Println(result)
	}
	return result, nil
}

func (l *Lexer) evalIfStatement(e *ifStatement) (float64, error) {
	result, err := l.eval(e.cond)
	if err != nil {
		return 0, err
	}

	if result != 0 {
		return l.eval(e.thenStmt)
	}
	if e.elseStmt != nil {
		return l.eval(e.elseStmt)
	}
	return 0, nil
}

func (l *Lexer) evalWhileStatement(e *whileStatement) (float64, error) {
	result, err := l.eval(e.cond)
	if err != nil {
		return 0, err
	}

	for result != 0 {
		_, err := l.eval(e.body)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}

// PrintAST prints the AST starting from the given node
func PrintAST(a ast, indentLevel int) {
	switch e := a.(type) {
	case *binaryExpr:
		fmt.Printf("%sBinaryExpr(%s)\n", strings.Repeat("\t", indentLevel), e.Op)
		PrintAST(e.lhs, indentLevel+1)
		PrintAST(e.rhs, indentLevel+1)
	case *unaryExpr:
		fmt.Printf("%sUnaryExpr\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.expr, indentLevel+1)
	case *root:
		fmt.Printf("%sAstRoot\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.expr, indentLevel+1)
	case *parenExpr:
		fmt.Printf("%sParenExpr\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.expr, indentLevel+1)
	case *variable:
		fmt.Printf("%sVariable(%s)\n", strings.Repeat("\t", indentLevel), e.name)
	case *number:
		fmt.Printf("%sNumber(%s)\n", strings.Repeat("\t", indentLevel), e.value)
	case *assignment:
		fmt.Printf("%sAssignment(%s)\n", strings.Repeat("\t", indentLevel), e.variable)
		PrintAST(e.expr, indentLevel+1)
	case *reassignment:
		fmt.Printf("%sReassignment(%s)\n", strings.Repeat("\t", indentLevel), e.variable)
		PrintAST(e.expr, indentLevel+1)
	case *stdPrint:
		fmt.Printf("%sStdPrint\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.expr, indentLevel+1)
	case *ifStatement:
		fmt.Printf("%sIfStatement\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.cond, indentLevel+1)
		PrintAST(e.thenStmt, indentLevel+1)
		if e.elseStmt != nil {
			PrintAST(e.elseStmt, indentLevel+1)
		}
	case *whileStatement:
		fmt.Printf("%sWhileStatement\n", strings.Repeat("\t", indentLevel))
		PrintAST(e.cond, indentLevel+1)
		PrintAST(e.body, indentLevel+1)
	default:
		fmt.Println("Unknown node type")
	}
}
