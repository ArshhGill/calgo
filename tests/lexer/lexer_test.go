package lexer_test

import (
	"calgo/internals/lexer"
	"calgo/internals/token"
	"testing"
)

func Test_lexer(t *testing.T) {
	input := "1+2"

	tokens := []token.Token{
		{
			Type:    token.INT,
			Literal: "1",
		},
		{
			Type:    token.PLUS,
			Literal: "+",
		},
		{
			Type:    token.INT,
			Literal: "2",
		},

		// {
		// 	Type:    token.INT,
		// 	Literal: "1",
		// },
		// {
		// 	Type:    token.MINUS,
		// 	Literal: "-",
		// },
		// {
		// 	Type:    token.INT,
		// 	Literal: "2",
		// },
		//
		// {
		// 	Type:    token.EOL,
		// 	Literal: "",
		// },
	}

	lex := lexer.NewLexer(input)

	for _, tt := range tokens {
		returnedToken := lex.NextToken()
		if returnedToken != tt {
			t.Fatalf("Expected %v, got %v, you suck lmfaooo", tt, returnedToken)
		}
	}
}
