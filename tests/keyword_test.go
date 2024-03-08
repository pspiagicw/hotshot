package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestKeywords(t *testing.T) {

	input := "if true false while case fn lambda let"

	expectedTokens := []token.Token{
		{
			TokenType:  token.IF,
			TokenValue: "if",
		},
		{
			TokenType:  token.TRUE,
			TokenValue: "true",
		},
		{
			TokenType:  token.FALSE,
			TokenValue: "false",
		},
		{
			TokenType:  token.WHILE,
			TokenValue: "while",
		},
		{
			TokenType:  token.CASE,
			TokenValue: "case",
		},
		{
			TokenType:  token.FN,
			TokenValue: "fn",
		},
		{
			TokenType:  token.LAMBDA,
			TokenValue: "lambda",
		},
		{
			TokenType:  token.LET,
			TokenValue: "let",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}
