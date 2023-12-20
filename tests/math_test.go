package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestMath(t *testing.T) {

	input := "+-/*%"

	expectedTokens := []token.Token{
		{
			TokenType:  token.PLUS,
			TokenValue: "+",
		},
		{
			TokenType:  token.MINUS,
			TokenValue: "-",
		},
		{
			TokenType:  token.SLASH,
			TokenValue: "/",
		},
		{
			TokenType:  token.MULTIPLY,
			TokenValue: "*",
		},
		{
			TokenType:  token.MOD,
			TokenValue: "%",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}

func TestComparator(t *testing.T) {

	input := "<>="

	expectedTokens := []token.Token{
		{
			TokenType:  token.LESSTHAN,
			TokenValue: "<",
		},
		{
			TokenType:  token.GREATERTHAN,
			TokenValue: ">",
		},
		{
			TokenType:  token.EQ,
			TokenValue: "=",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}
