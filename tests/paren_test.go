package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestParen(t *testing.T) {
	input := "(){}"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.RPAREN,
			TokenValue: ")",
		},
		{
			TokenType:  token.LBRACE,
			TokenValue: "{",
		},
		{
			TokenType:  token.RBRACE,
			TokenValue: "}",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)
}

func TestParenWithSpaces(t *testing.T) {
	input := "(     )"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.RPAREN,
			TokenValue: ")",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)
}
