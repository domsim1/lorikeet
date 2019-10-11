package lexer

import (
	"monkey/token"
	"unicode"
)

// Lexer struct
type Lexer struct {
	input        []rune
	position     int
	readPosition int // TODO: Track line number
	ru           rune
}

// New *Lexer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
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
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ru)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ru)
	case '(':
		tok = newToken(token.LPAREN, l.ru)
	case ')':
		tok = newToken(token.RPAREN, l.ru)
	case ',':
		tok = newToken(token.COMMA, l.ru)
	case ':':
		tok = newToken(token.COLON, l.ru)
	case '+':
		tok = newToken(token.PLUS, l.ru)
	case '-':
		tok = newToken(token.MINUS, l.ru)
	case '{':
		tok = newToken(token.LBRACE, l.ru)
	case '}':
		tok = newToken(token.RBRACE, l.ru)
	case '[':
		tok = newToken(token.LBRACKET, l.ru)
	case ']':
		tok = newToken(token.RBRACKET, l.ru)
	case '!':
		if l.peekRune() == '=' {
			ch := l.ru
			l.readRune()
			literal := string(ch) + string(l.ru)
			tok = token.Token{Type: token.NOTEQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ru)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ru)
	case '/':
		tok = newToken(token.SLASH, l.ru)
	case '<':
		tok = newToken(token.LT, l.ru)
	case '>':
		tok = newToken(token.GT, l.ru)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ru) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ru) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ru)
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

func newToken(tokenType token.Type, ru rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ru)}
}

func isDigit(ru rune) bool {
	return '0' <= ru && ru <= '9'
}
