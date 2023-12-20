package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestIdent(t *testing.T) {

	input := "(+ one two)"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.PLUS,
			TokenValue: "+",
		},
		{
			TokenType:  token.IDENT,
			TokenValue: "one",
		},
		{
			TokenType:  token.IDENT,
			TokenValue: "two",
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
func TestIdent2(t *testing.T) {
	input := "(one)"

	expectedTokens := []token.Token{
		{
			TokenType:  token.LPAREN,
			TokenValue: "(",
		},
		{
			TokenType:  token.IDENT,
			TokenValue: "one",
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
