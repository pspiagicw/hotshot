package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestNumbers(t *testing.T) {

	input := "1"

	expectedTokens := []token.Token{
		{
			TokenType:  token.NUM,
			TokenValue: "1",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}

	checkTokens(t, expectedTokens, input)

}
func TestNum2(t *testing.T) {
	input := "(1)"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.NUM,
			TokenValue: "1",
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

func TestNumNegative(t *testing.T) {

	input := "(-1)+1"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.NUM,
			TokenValue: "-1",
		},
		{
			TokenType:  token.RPAREN,
			TokenValue: ")",
		},
		{
			TokenType:  token.NUM,
			TokenValue: "1",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}

	checkTokens(t, expectedTokens, input)
}
