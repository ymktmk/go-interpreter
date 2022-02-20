package lexer

import (
	"go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x,y) {
		x + y;
	};
	let result = add(five,ten);
	`
	
	// 配列に格納
	tests := []struct {
		// 定義した型
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.INDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "add"},
		{token.ASSIGN, "="},
		{token.INT, "fn"},
		{token.LPAREN, "("},
		{token.INDENT, "x"},
		{token.COMMA, ","},
		{token.INDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INDENT, "x"},
		{token.PLUS, "+"},
		{token.INDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.INDENT, "result"},
		{token.ASSIGN, "="},
		{token.INDENT, "add"},
		{token.LPAREN, "("},
		{token.INDENT, "five"},
		{token.COMMA, ","},
		{token.INDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}


	l := New(input)

	for i, tt := range tests {
		// 構造体の関数
		// Lexer
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}


