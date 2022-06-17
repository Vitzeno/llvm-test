package main

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
	parseResult expr
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

func (l *Lexer) Lex(lval *yySymType) int {
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
				lit := l.lexInt()
				lval.String = lit
				return tokenNumber
			} else if unicode.IsLetter(r) {
				// backup and let lexIdentifier rescan the beginning of the identifier
				l.backup()
				lit := l.lexIdentifier()
				lval.String = lit
				switch lit {
				case "let":
					return tokenLet
				default:
					return tokenIdentifier
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
