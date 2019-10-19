package lexer

import (
	"testing"

	"lorikeet/token"
)

func TestNextToken(t *testing.T) {
	input :=
		`let five = 5;
		 let ten = 10;

		 let add = fn(x, y) {
			 x + y;
		 };

		 let result = add(five, ten);
		 !-/*5;
		 5 < 10 > 5;

		 if (5 < 10) {
			 return true;
		 } else {
			 return false;
		 }

		 10 == 10;
		 10 != 9;
		 "foobar"
		 "foo bar"
		 [1, 2];
		 {"foo": "bar"}
		 let 汉字 = "かんじ";
		 "🐵 🙈 🙉 🙊 🐒"
		 macro(x, y) { x + y; };
		 $do();
		 getName()|>greet("hello");
		 1.5;
		`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
		expectedLine    int
	}{
		{token.LET, "let", 1},
		{token.IDENT, "five", 1},
		{token.ASSIGN, "=", 1},
		{token.INT, "5", 1},
		{token.SEMICOLON, ";", 1},
		{token.LET, "let", 2},
		{token.IDENT, "ten", 2},
		{token.ASSIGN, "=", 2},
		{token.INT, "10", 2},
		{token.SEMICOLON, ";", 2},
		{token.LET, "let", 4},
		{token.IDENT, "add", 4},
		{token.ASSIGN, "=", 4},
		{token.FUNCTION, "fn", 4},
		{token.LPAREN, "(", 4},
		{token.IDENT, "x", 4},
		{token.COMMA, ",", 4},
		{token.IDENT, "y", 4},
		{token.RPAREN, ")", 4},
		{token.LBRACE, "{", 4},
		{token.IDENT, "x", 5},
		{token.PLUS, "+", 5},
		{token.IDENT, "y", 5},
		{token.SEMICOLON, ";", 5},
		{token.RBRACE, "}", 6},
		{token.SEMICOLON, ";", 6},
		{token.LET, "let", 8},
		{token.IDENT, "result", 8},
		{token.ASSIGN, "=", 8},
		{token.IDENT, "add", 8},
		{token.LPAREN, "(", 8},
		{token.IDENT, "five", 8},
		{token.COMMA, ",", 8},
		{token.IDENT, "ten", 8},
		{token.RPAREN, ")", 8},
		{token.SEMICOLON, ";", 8},
		{token.BANG, "!", 9},
		{token.MINUS, "-", 9},
		{token.SLASH, "/", 9},
		{token.ASTERISK, "*", 9},
		{token.INT, "5", 9},
		{token.SEMICOLON, ";", 9},
		{token.INT, "5", 10},
		{token.LT, "<", 10},
		{token.INT, "10", 10},
		{token.GT, ">", 10},
		{token.INT, "5", 10},
		{token.SEMICOLON, ";", 10},
		{token.IF, "if", 12},
		{token.LPAREN, "(", 12},
		{token.INT, "5", 12},
		{token.LT, "<", 12},
		{token.INT, "10", 12},
		{token.RPAREN, ")", 12},
		{token.LBRACE, "{", 12},
		{token.RETURN, "return", 13},
		{token.TRUE, "true", 13},
		{token.SEMICOLON, ";", 13},
		{token.RBRACE, "}", 14},
		{token.ELSE, "else", 14},
		{token.LBRACE, "{", 14},
		{token.RETURN, "return", 15},
		{token.FALSE, "false", 15},
		{token.SEMICOLON, ";", 15},
		{token.RBRACE, "}", 16},
		{token.INT, "10", 18},
		{token.EQ, "==", 18},
		{token.INT, "10", 18},
		{token.SEMICOLON, ";", 18},
		{token.INT, "10", 19},
		{token.NOTEQ, "!=", 19},
		{token.INT, "9", 19},
		{token.SEMICOLON, ";", 19},
		{token.STRING, "foobar", 20},
		{token.STRING, "foo bar", 21},
		{token.LBRACKET, "[", 22},
		{token.INT, "1", 22},
		{token.COMMA, ",", 22},
		{token.INT, "2", 22},
		{token.RBRACKET, "]", 22},
		{token.SEMICOLON, ";", 22},
		{token.LBRACE, "{", 23},
		{token.STRING, "foo", 23},
		{token.COLON, ":", 23},
		{token.STRING, "bar", 23},
		{token.RBRACE, "}", 23},
		{token.LET, "let", 24},
		{token.IDENT, "汉字", 24},
		{token.ASSIGN, "=", 24},
		{token.STRING, "かんじ", 24},
		{token.SEMICOLON, ";", 24},
		{token.STRING, "🐵 🙈 🙉 🙊 🐒", 25},
		{token.MACRO, "macro", 26},
		{token.LPAREN, "(", 26},
		{token.IDENT, "x", 26},
		{token.COMMA, ",", 26},
		{token.IDENT, "y", 26},
		{token.RPAREN, ")", 26},
		{token.LBRACE, "{", 26},
		{token.IDENT, "x", 26},
		{token.PLUS, "+", 26},
		{token.IDENT, "y", 26},
		{token.SEMICOLON, ";", 26},
		{token.RBRACE, "}", 26},
		{token.SEMICOLON, ";", 26},
		{token.MONEY, "$", 27},
		{token.IDENT, "do", 27},
		{token.LPAREN, "(", 27},
		{token.RPAREN, ")", 27},
		{token.SEMICOLON, ";", 27},
		{token.IDENT, "getName", 28},
		{token.LPAREN, "(", 28},
		{token.RPAREN, ")", 28},
		{token.PIPE, "|>", 28},
		{token.IDENT, "greet", 28},
		{token.LPAREN, "(", 28},
		{token.STRING, "hello", 28},
		{token.RPAREN, ")", 28},
		{token.SEMICOLON, ";", 28},
		{token.FLOAT, "1.5", 29},
		{token.SEMICOLON, ";", 29},
		{token.EOF, "", 30},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}
	}
}
