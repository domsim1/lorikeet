package lexer

import (
	"lorikeet/token"
	"unicode"
)

// Lexer struct
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	linePosition int
	ru           rune
}

// New *Lexer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input), linePosition: 1}
	l.readRune()
	return l
}

// NextToken *Lexer
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ru {
	case '=':
		if l.peekRune() == '=' {
			ch := l.ru
			l.readRune()
			literal := string(ch) + string(l.ru)
			tok = token.Token{Type: token.EQ, Literal: literal, Line: l.linePosition}
		} else {
			tok = newToken(token.ASSIGN, l.ru, l.linePosition)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ru, l.linePosition)
	case '(':
		tok = newToken(token.LPAREN, l.ru, l.linePosition)
	case ')':
		tok = newToken(token.RPAREN, l.ru, l.linePosition)
	case ',':
		tok = newToken(token.COMMA, l.ru, l.linePosition)
	case ':':
		tok = newToken(token.COLON, l.ru, l.linePosition)
	case '+':
		tok = newToken(token.PLUS, l.ru, l.linePosition)
	case '-':
		tok = newToken(token.MINUS, l.ru, l.linePosition)
	case '{':
		tok = newToken(token.LBRACE, l.ru, l.linePosition)
	case '}':
		tok = newToken(token.RBRACE, l.ru, l.linePosition)
	case '[':
		tok = newToken(token.LBRACKET, l.ru, l.linePosition)
	case ']':
		tok = newToken(token.RBRACKET, l.ru, l.linePosition)
	case '|':
		if l.peekRune() == '>' {
			ch := l.ru
			l.readRune()
			literal := string(ch) + string(l.ru)
			tok = token.Token{Type: token.PIPE, Literal: literal, Line: l.linePosition}
		} else {
			tok = newToken(token.ILLEGAL, l.ru, l.linePosition)
		}
	case '!':
		if l.peekRune() == '=' {
			ch := l.ru
			l.readRune()
			literal := string(ch) + string(l.ru)
			tok = token.Token{Type: token.NOTEQ, Literal: literal, Line: l.linePosition}
		} else {
			tok = newToken(token.BANG, l.ru, l.linePosition)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ru, l.linePosition)
	case '/':
		tok = newToken(token.SLASH, l.ru, l.linePosition)
	case '<':
		tok = newToken(token.LT, l.ru, l.linePosition)
	case '>':
		tok = newToken(token.GT, l.ru, l.linePosition)
	case '$':
		tok = newToken(token.MONEY, l.ru, l.linePosition)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Line = l.linePosition
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.linePosition
	default:
		if isLetter(l.ru) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.linePosition
			return tok
		} else if isDigit(l.ru) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Line = l.linePosition
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ru, l.linePosition)
	}

	l.readRune()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ru) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ru = 0
	} else {
		l.ru = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhitespace() {
	for l.ru == ' ' || l.ru == '\t' || l.ru == '\n' || l.ru == '\r' {
		switch l.ru {
		case '\r':
			if l.peekRune() == '\n' {
				l.readRune()
			}
			l.linePosition++
		case '\n':
			l.linePosition++
		}
		l.readRune()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ru) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readRune()
		if l.ru == '"' || l.ru == 0 {
			break
		}
	}
	return string(l.input[position:l.position])
}

// Allowed identifier chars
func isLetter(ru rune) bool {
	return ru == '_' || unicode.IsLetter(ru)
}

func newToken(tokenType token.Type, ru rune, line int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ru), Line: line}
}

func isDigit(ru rune) bool {
	return '0' <= ru && ru <= '9'
}
