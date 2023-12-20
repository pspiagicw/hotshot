package tests

import (
	"reflect"
	"testing"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/parser"
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
func checkTree(t *testing.T, input string, expectedTree []ast.Statement) {

	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	actualTree := parser.Parse()
	if len(parser.Errors()) != 0 {
		t.Fatalf("Errors while parsing")
	}

	for i, expectedStatement := range expectedTree {
		actualStatement := actualTree.Statements[i]

		if !matchStatement(t, expectedStatement, actualStatement) {
			t.Errorf("Statement [%d], not matching, actual: %s, expected: %s", i+1, actualStatement.StringifyStatement(), expectedStatement.StringifyStatement())
		}
	}
}
func matchStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	t.Helper()

	return reflect.DeepEqual(expectedStatement, actualStatement)
}
