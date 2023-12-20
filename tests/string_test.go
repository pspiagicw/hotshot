package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestStrings(t *testing.T) {
	input := `"something"`

	expectedTokens := []token.Token{
		{
			TokenType:  token.STRING,
			TokenValue: "something",
		},
		{
			TokenType:  token.EOF,
			TokenValue: " ",
		},
	}
	checkTokens(t, expectedTokens, input)
}
