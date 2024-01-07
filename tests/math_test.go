package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestMath(t *testing.T) {

	input := "+-/*%^"

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
			TokenType:  token.POWER,
			TokenValue: "^",
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
func TestOthers(t *testing.T) {
	input := "|;$@!,?#"

	expectedTokens := []token.Token{
		{
			TokenType:  token.PIPE,
			TokenValue: "|",
		},
		{
			TokenType:  token.SEMICOLON,
			TokenValue: ";",
		},
		{
			TokenType:  token.DOLLAR,
			TokenValue: "$",
		},
		{
			TokenType:  token.AT,
			TokenValue: "@",
		},
		{
			TokenType:  token.BANG,
			TokenValue: "!",
		},
		{
			TokenType:  token.COMMA,
			TokenValue: ",",
		},
		{
			TokenType:  token.QUESTION,
			TokenValue: "?",
		},
		{
			TokenType:  token.HASH,
			TokenValue: "#",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)

}
