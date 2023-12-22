package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestKeywords(t *testing.T) {

	input := "if true false while set echo"

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
			TokenType:  token.SET,
			TokenValue: "set",
		},
		{
			TokenType:  token.ECHO,
			TokenValue: "echo",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}
