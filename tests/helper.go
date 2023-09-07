package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

func checkTokens(t *testing.T, expected []token.Token, input string) {
	t.Helper()

	lexer := lexer.NewLexer(input)

	for i, expectedToken := range expected {
		actualToken := lexer.Next()
		if actualToken.TokenType != expectedToken.TokenType {
			t.Errorf("Test [%d], Expected TokenType: '%v', Actual TokenType: '%v'", i, expectedToken.TokenType, actualToken.TokenType)

		}
		if actualToken.TokenValue != expectedToken.TokenValue {
			t.Errorf("Test [%d], Expected TokenValue: '%v', Actual TokenValue: '%v'", i, expectedToken.TokenValue, actualToken.TokenValue)
		}
	}

}
