package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"

	"github.com/llir/llvm/ir"
)

// Position tracks the position of the lexer
type Position struct {
	line int
	col  int
}

type varible struct {
	location *ir.InstAlloca
	value    float64
}

// Lexer is the lexer struct
type Lexer struct {
	rootAst    Ast
	pos        Position
	variables  map[string]varible
	evalFailed bool
	reader     *bufio.Reader
	errors     []string
	buffer     bytes.Buffer // Buffer to store tokens temporarily
}

// NewLexer creates a new lexer
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:       Position{line: 1, col: 0},
		variables: make(map[string]varible),
		rootAst:   nil,
		reader:    bufio.NewReader(reader),
	}
}

// Error is the error handler for the lexer
func (l *Lexer) Error(e string) {
	l.errors = append(l.errors, fmt.Sprintf("%s Line: %d, Col: %d", e, l.pos.line, l.pos.col))
	l.evalFailed = true
}

// Errors returns the errors
func (l *Lexer) Errors() []string {
	return l.errors
}

// Lex is the main lexer function
func (l *Lexer) Lex(lval *YYSymType) int {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return EOF
			}
			l.Error(fmt.Sprintf("unexpected error: %s", err))
		}

		l.pos.col++

		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return tokenSeparator
		// case '<':
		// // potential issues cause by rune over-scanning and backup
		// 	nextRune, _, _ := l.reader.ReadRune()
		// 	if nextRune == '=' {
		// 		fmt.Println("less than or equal")
		// 		return tokenLte
		// 	}
		// 	l.backup()
		// 	fmt.Println("less than")
		// 	return tokenLt
		// case '>':
		// // potential issues cause by rune over-scanning and backup
		// 	nextRune, _, _ := l.reader.ReadRune()
		// 	if nextRune == '=' {
		// 		return tokenGte
		// 	}
		// 	l.backup()
		// 	return tokenGt
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				// backup and let lexInt rescan the beginning of the int
				l.backup()
				digit := l.lexInt()
				lval.String = digit
				return tokenNumber
			} else if unicode.IsLetter(r) {
				// backup and let lexIdentifier rescan the beginning of the identifier
				l.backup()
				lit := l.lexIdentifier()
				//fmt.Printf("letter: %v\n", lit)
				switch lit {
				case "let":
					lval.String = lit
					return tokenLet
				case "print":
					lval.String = lit
					return tokenPrint
				case "if":
					lval.String = lit
					return tokenIf
				case "then":
					lval.String = lit
					return tokenThen
				case "else":
					lval.String = lit
					return tokenElse
				case "AND":
					lval.String = lit
					return tokenAnd
				case "OR":
					lval.String = lit
					return tokenOr
				case "NE":
					lval.String = lit
					return tokenNe
				case "while":
					lval.String = lit
					return tokenWhile
				default:
					lval.String = lit
					return tokenIdentifier
				}
			} else if unicode.IsSymbol(r) {
				// backup and let lexSymbol rescan the beginning of the identifier
				l.backup()
				symbol := l.lexSymbol()
				//fmt.Printf("symbol: %v\n", symbol)
				switch symbol {
				case "+": // special case for + fallthrough since handled in parser
					return int(r)
				case "=":
					lval.String = symbol
					return tokenAssign
				case "<":
					lval.String = symbol
					return tokenLt
				case ">":
					lval.String = symbol
					return tokenGt
				case "==":
					lval.String = symbol
					return tokenEq
				}
			} else {
				return int(r)
			}
		}
	}
}

// lexInt scans the input for an integer
func (l *Lexer) lexInt() string {
	l.buffer.Reset()
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.buffer.String()
			}
			l.Error(fmt.Sprintf("unexpected error: %s", err))
			return l.buffer.String()
		}

		l.pos.col++
		if unicode.IsDigit(r) {
			l.buffer.WriteRune(r)
		} else {
			// over-scanned int, need to move back
			l.backup()
			return l.buffer.String()
		}
	}
}

// lexSymbol scans the input for a symbol
func (l *Lexer) lexSymbol() string {
	l.buffer.Reset()
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.buffer.String()
			}
			l.Error(fmt.Sprintf("unexpected error: %s", err))
			return l.buffer.String()
		}

		l.pos.col++
		if unicode.IsSymbol(r) {
			l.buffer.WriteRune(r)
		} else {
			// over-scanned identifier, need to move back
			l.backup()
			return l.buffer.String()
		}
	}
}

// lexIdentifier scans the input for an identifier
func (l *Lexer) lexIdentifier() string {
	l.buffer.Reset()
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.buffer.String()
			}
			l.Error(fmt.Sprintf("unexpected error: %s", err))
			return l.buffer.String()
		}

		l.pos.col++
		if unicode.IsLetter(r) {
			l.buffer.WriteRune(r)
		} else {
			// over-scanned identifier, need to move back
			l.backup()
			return l.buffer.String()
		}
	}
}

// resetPosition resets the column and increments the line
func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.col = 0
}

// backup moves the reader back one rune
func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.col--
}
