package tests

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
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
		matchToken(t, i, expectedToken, actualToken)
	}

}
func matchToken(t *testing.T, i int, expected token.Token, actual *token.Token) {
	t.Helper()
	if actual.TokenType != expected.TokenType {
		t.Errorf("Test [%d], Expected TokenType: '%v', Actual TokenType: '%v'", i, expected.TokenType, actual.TokenType)

	}
	if actual.TokenValue != expected.TokenValue {
		t.Errorf("Test [%d], Expected TokenValue: '%v', Actual TokenValue: '%v'", i, expected.TokenValue, actual.TokenValue)
	}

}
func checkTree(t *testing.T, input string, expectedTree []ast.Statement) {

	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	actualTree := parser.Parse()
	if len(parser.Errors()) != 0 {
		for _, err := range parser.Errors() {
			t.Log(err)
		}
		t.Fatalf("Errors while parsing")
	}

	for i, expectedStatement := range expectedTree {
		actualStatement := actualTree.Statements[i]

		matchStatement(t, expectedStatement, actualStatement)
	}
}
func matchIntStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.IntStatement)
	if !ok {
		t.Fatalf("Expected int statement, got others")
	}
	a, ok := actualStatement.(*ast.IntStatement)
	if !ok {
		t.Fatalf("Expected int statement, got: %v", actualStatement)
	}
	return e.Value == a.Value
}
func matchStringStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.StringStatement)
	if !ok {
		t.Fatalf("Expected string statement, got others")
	}
	a, ok := actualStatement.(*ast.StringStatement)
	if !ok {
		t.Fatalf("Expected string statement, got: %v", actualStatement)
	}
	return e.Value == a.Value
}
func matchBoolStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.BoolStatement)
	if !ok {
		t.Fatalf("Expected bool statement, got others")
	}
	a, ok := actualStatement.(*ast.BoolStatement)
	if !ok {
		t.Fatalf("Expected bool statement, got: %v", actualStatement)
	}
	return e.Value == a.Value
}
func matchTableStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.TableStatement)
	if !ok {
		t.Fatalf("Expected table statement, got others")
	}
	a, ok := actualStatement.(*ast.TableStatement)
	if !ok {
		t.Fatalf("Expected table statement, got others")
	}
	result := true
	for index, _ := range e.Elements {
		result = result && matchStatement(t, e.Elements[index], a.Elements[index])
	}
	return result
}
func matchEmptyStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	_, ok := expectedStatement.(*ast.EmptyStatement)
	if !ok {
		t.Fatalf("Expected string statement, got others")
	}
	_, ok = actualStatement.(*ast.EmptyStatement)
	if !ok {
		t.Fatalf("Expected string statement, got: %v", actualStatement)
	}
	return true
}
func matchStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	t.Helper()
	f, ok := expectedStatement.(*ast.CallStatement)
	if !ok {
		switch s := expectedStatement.(type) {
		case *ast.IntStatement:
			return matchIntStatement(t, expectedStatement, actualStatement)
		case *ast.EmptyStatement:
			return matchEmptyStatement(t, expectedStatement, actualStatement)
		case *ast.StringStatement:
			return matchStringStatement(t, expectedStatement, actualStatement)
		case *ast.BoolStatement:
			return matchBoolStatement(t, expectedStatement, actualStatement)
		case *ast.TableStatement:
			return matchTableStatement(t, expectedStatement, actualStatement)
		default:
			t.Fatalf("Some bloody type found!: %v", s)
		}
	}
	g, ok := actualStatement.(*ast.CallStatement)
	if !ok {
		t.Fatalf("Error expected statement was Functional, actual isn't!\n")
	}
	matchToken(t, 0, *f.Op, g.Op)
	for i, fArgs := range f.Args {
		exArgs := g.Args[i]
		if !matchStatement(t, fArgs, exArgs) {
			return false
		}
	}
	return true
}
func validStatement(t *testing.T, input string) bool {
	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	program := parser.Parse()
	if len(parser.Errors()) != 0 {
		return false
	}

	snaps.MatchSnapshot(t, program)

	return true
}
