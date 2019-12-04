package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input = `=+(){},;`

	var test = []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	var l = NewLexer(input)

	for i, tt := range tests {
		var tok = l.NextToken()

		if tok.Type != tt.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.ExpectedType, tok.Type)
		}

		if tok.Literal != tt.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
