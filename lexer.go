package main

import (
	"bufio"
	"io"
	"unicode"
)

type Position struct {
	line int
	col  int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, col: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, tokenEof, ""
			}

			panic(err)
		}

		l.pos.col++

		switch r {
		case '\n':
			l.resetPosition()
		case '(':
			return l.pos, tokenLeftParen, string(r)
		case ')':
			return l.pos, tokenRightParen, string(r)
		case '{':
			return l.pos, tokenLeftBrace, string(r)
		case '}':
			return l.pos, tokenRightBrace, string(r)
		case '#':
			startPos := l.pos
			// skip until end of line since comment
			_, _, err := l.reader.ReadLine()
			if err != nil {
				panic(err)
			}
			l.resetPosition()
			return startPos, tokenComment, string(r)
		case '=':
			return l.pos, tokenAssignment, string(r)
		case '+':
			return l.pos, tokenAddition, string(r)
		case '-':
			return l.pos, tokenSubtraction, string(r)
		case '*':
			return l.pos, tokenMultiplication, string(r)
		case '/':
			return l.pos, tokenDivision, string(r)
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				// backup and let lexInt rescan the beginning of the int
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, tokenInt, lit
			} else if unicode.IsLetter(r) {
				// backup and let lexIdentifier rescan the beginning of the identifier
				startPos := l.pos
				l.backup()
				lit := l.lexIdentifier()
				switch lit {
				case "func", "if", "else", "return":
					return startPos, tokenKeyword, lit
				default:
					return startPos, tokenIdentifier, lit
				}
			} else {
				return l.pos, tokenIllegal, string(r)
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
			// overscanned int, need to move back
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
			// overscanned indentifier, need to move back
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
