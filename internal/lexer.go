package internal

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Position struct {
	line int
	col  int
}

type Lexer struct {
	pos         Position
	variables   map[string]float64
	evalFailed  bool
	parseResult ast
	reader      *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:         Position{line: 1, col: 0},
		variables:   make(map[string]float64),
		parseResult: &astRoot{},
		reader:      bufio.NewReader(reader),
	}
}

func (l *Lexer) Error(e string) {
	fmt.Println(e)
	fmt.Printf("Line: %d, Col: %d\n", l.pos.line, l.pos.col)
	l.evalFailed = true
}

func (l *Lexer) Lex(lval *YYSymType) int {
	//spew.Dump(l.variables)
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
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				// backup and let lexInt rescan the beginning of the int
				l.backup()
				digit := l.lexInt()
				lval.String = digit
				fmt.Printf("digit: %v\n", digit)
				return tokenNumber
			} else if unicode.IsLetter(r) {
				// backup and let lexIdentifier rescan the beginning of the identifier
				l.backup()
				lit := l.lexIdentifier()
				fmt.Printf("letter: %v\n", lit)
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
				fmt.Printf("symbol: %v\n", symbol)
				switch symbol {
				case "+": // special case for + fallthrough since handled in parser
					return int(r)
				case "=":
					lval.String = symbol
					return tokenAssign
				case "<":
					lval.String = symbol
					return tokenLe
				case ">":
					lval.String = symbol
					return tokenGe
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

func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if unicode.IsDigit(r) {
			lit += string(r)
		} else {
			// over-scanned int, need to move back
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexSymbol() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if unicode.IsSymbol(r) {
			lit += string(r)
		} else {
			// over=scanned identifier, need to move back
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexIdentifier() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if unicode.IsLetter(r) {
			lit += string(r)
		} else {
			// over=scanned identifier, need to move back
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.col = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.col--
}
