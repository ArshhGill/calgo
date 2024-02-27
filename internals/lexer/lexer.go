package lexer

import (
	"calgo/internals/token"
	"fmt"
)

const NULLCHR rune = '\x00'

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func NewLexer(input string) *Lexer {
	uwu := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           NULLCHR,
	}

	uwu.readChar()

	return uwu
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhites()

	switch l.ch {
	case '+':
		tok = token.Token{
			Type:    token.PLUS,
			Literal: string(l.ch),
		}
	case '-':
		tok = token.Token{
			Type:    token.MINUS,
			Literal: string(l.ch),
		}
	case '*':
		tok = token.Token{
			Type:    token.MULTIPLY,
			Literal: string(l.ch),
		}
	case '/':
		tok = token.Token{
			Type:    token.DIVIDE,
			Literal: string(l.ch),
		}
	case '%':
		tok = token.Token{
			Type:    token.MODULO,
			Literal: string(l.ch),
		}
	case NULLCHR:
		tok = token.Token{
			Type:    token.EOL,
			Literal: string(l.ch),
		}
	default:
		if ifDig(l.ch) {
			fmt.Println("uwu")
			return l.readDig()
		} else {
			return token.Token{
				Type:    token.ILLEGAL,
				Literal: string(l.ch),
			}
		}

	}

	l.readChar()

	return tok
}

func (l *Lexer) readDig() token.Token {
	pos := l.position
	for ifDig(l.ch) {
		l.readChar()
	}
	fmt.Println(l.input[pos:l.position], l.position, pos)

	return token.Token{
		Type:    token.INT,
		Literal: l.input[pos:l.position],
	}
}

func ifDig(ch rune) bool {
	return '0' <= ch && '9' >= ch
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = NULLCHR
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhites() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}
