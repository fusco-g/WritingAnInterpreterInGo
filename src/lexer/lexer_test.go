package lexer

import (
	"interpeter/token"
	"testing"
)

// test source code

func TestNextToken(t *testing.T) {
	input := `let five = 5;
 let ten = 10;
 let add = fn(x, y) {
 x + y;
 };
 let result = add(five, ten);
 `

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.O_ROUND_BRACKET, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.C_ROUND_BRACKET, ")"},
		{token.O_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.C_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.O_ROUND_BRACKET, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.C_ROUND_BRACKET, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for index := range test {
		token := lexer.NextToken()

		if token.Type != test[index].expectedType {
			t.Errorf("Got = %v - Expected = %v. AT INDEX = %d", token.Type, test[index].expectedType, index)
			t.Fatal("Token Type Wrong!")
		}

		if token.Literal != test[index].expectedLiteral {
			t.Errorf("Got = %v - Expected = %v. AT INDEX = %d", token.Literal, test[index].expectedLiteral, index)
			t.Fatal("Wrong Token Literal!")
		}
	}
}
