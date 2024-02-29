package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestKeywords(t *testing.T) {

	input := "if true false for case fn"

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
			TokenType:  token.FOR,
			TokenValue: "for",
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
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}
