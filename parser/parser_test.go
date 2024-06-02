package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

func TestSetStatement(t *testing.T) {
	input := `(set ["name"]something 1)`

	expectedTree := []ast.Statement{
		&ast.SetStatement{
			Target: &ast.IndexStatement{
				Key: &ast.StringStatement{
					Value: "name",
				},
				Target: &ast.IdentStatement{
					Value: &token.Token{TokenType: token.IDENT, TokenValue: "something"},
				},
			},
			Value: &ast.IntStatement{
				Value: 1,
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestSliceStatement(t *testing.T) {
	input := `[0]something`

	expectedTree := []ast.Statement{
		&ast.IndexStatement{
			Key: &ast.IntStatement{
				Value: 0,
			},
			Target: &ast.IdentStatement{
				Value: &token.Token{TokenType: token.IDENT, TokenValue: "something"},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestQuoteStatement(t *testing.T) {
	input := `'something`

	expectedTree := []ast.Statement{
		&ast.QuoteStatement{
			Body: &token.Token{
				TokenType:  token.IDENT,
				TokenValue: "something",
			},
		},
	}

	checkTree(t, input, expectedTree)

}

func TestEmptyStatement(t *testing.T) {
	input := "()"

	expectedTree := []ast.Statement{
		&ast.EmptyStatement{},
	}

	checkTree(t, input, expectedTree)

}

func TestSimpleParser(t *testing.T) {

	input := `
    1
    2
    `

	expectedTree := []ast.Statement{
		&ast.IntStatement{
			Value: 1,
		},
		&ast.IntStatement{
			Value: 2,
		},
	}

	checkTree(t, input, expectedTree)
}

func TestBoolStatements(t *testing.T) {
	input := `true
    false`

	expectedTree := []ast.Statement{
		&ast.BoolStatement{
			Value: true,
		},
		&ast.BoolStatement{
			Value: false,
		},
	}

	checkTree(t, input, expectedTree)
}

func TestSimpleString(t *testing.T) {
	input := `"hello"`

	expectedTree := []ast.Statement{
		&ast.StringStatement{
			Value: "hello",
		},
	}
	checkTree(t, input, expectedTree)
}
func TestTableStatement(t *testing.T) {
	input := `{1 2 3}`

	expectedTree := []ast.Statement{
		&ast.TableStatement{
			Elements: []ast.Statement{
				&ast.IntStatement{
					Value: 1,
				},
				&ast.IntStatement{
					Value: 2,
				},
				&ast.IntStatement{
					Value: 3,
				},
			},
		},
	}
	checkTree(t, input, expectedTree)
}

func TestComments(t *testing.T) {
	input := `; some comments about you ;
    69
    `
	expectedTree := []ast.Statement{
		&ast.IntStatement{
			Value: 69,
		},
	}

	checkTree(t, input, expectedTree)
}
func TestAddition(t *testing.T) {
	input := `(+ 1 2)`

	expectedTree := []ast.Statement{
		&ast.CallStatement{
			Op: &token.Token{
				TokenType:  token.PLUS,
				TokenValue: "+",
			},
			Args: []ast.Statement{
				&ast.IntStatement{
					Value: 1,
				},
				&ast.IntStatement{
					Value: 2,
				},
			},
		},
	}
	checkTree(t, input, expectedTree)

}
func TestNestedStatement(t *testing.T) {
	input := `(+ (+ 1 2) 3)`

	expectedTree := []ast.Statement{
		&ast.CallStatement{
			Op: &token.Token{
				TokenType:  token.PLUS,
				TokenValue: "+",
			},
			Args: []ast.Statement{
				&ast.CallStatement{
					Op: &token.Token{
						TokenType:  token.PLUS,
						TokenValue: "+",
					},
					Args: []ast.Statement{
						&ast.IntStatement{
							Value: 1,
						},
						&ast.IntStatement{
							Value: 2,
						},
					},
				},
				&ast.IntStatement{
					Value: 3,
				},
			},
		},
	}
	checkTree(t, input, expectedTree)

}
func TestValidOp(t *testing.T) {

	tt := map[string]bool{
		"(+ 1 2)":  true,
		"(- 1 2)":  true,
		"(/ 1 2)":  true,
		"(* 1 2)":  true,
		"(^ 1 2)":  true,
		"(if 1 2)": true,
		"(= 1 2)":  true,
		"(< 1 2)":  true,
		"(> 1 2)":  true,
		// Will fail in execution, but pass in parser.
		// Can be fixed with AST passes.

		"(let 1 2)": false,
		"(% 1 2)":   false,
		"(# 1 2)":   false,
		"(; 1 2)":   false,
		"(@ 1 2)":   false,
		"(, 1 2)":   false,
		"(! 1 2)":   false,
		"(? 1 2)":   false,
	}

	for input, expectedResult := range tt {
		t.Run(input, func(t *testing.T) {
			if validStatement(t, input) != expectedResult {
				t.Errorf("Test '%s' failed to match result: %t!", input, expectedResult)
			}
		})
	}
}
func TestValidStatement(t *testing.T) {
	tt := map[string]bool{
		"+": false,
		"-": false,
		"*": false,
		"/": false,
		"%": false,
		"|": false,
		";": true,

		"@": false,
		"$": false,
		"!": false,
		"?": false,
		"#": false,

		"if":    false,
		"while": false,
		"case":  false,

		"=": false,
		">": false,
		"<": false,

		"somevar":      true,
		"1":            true,
		`"somestring"`: true,

		`(+ 1 2)`:         true,
		`(+ "foo" "bar")`: true,
		// Should parse properly, execution is not a worry now! This would fail in execution, not here!
		`(/ "foo" "bar")`:              true,
		`(if (= 1 2) (echo g))`:        true,
		`(echo "Hello, World!")`:       true,
		"; this should be a comment ;": true,
		"; this should be a comment":   true,
		"(let someVar 3)":              true,
		"(somefunc somearg 1)":         true,
		"(= 1 1)":                      true,
		"(> 1 1)":                      true,
		"(< 1 1)":                      true,
		`(= "some" "some")`:            true,

		`(fn hello () (echo "Hello, World"))`: true,
		`(fn add (x y) (+ x y))`:              true,

		`(lambda () (echo "Hello, World"))`: true,

		`{ 1 2 3}`: true,
		`(cond ((= 1 1) "1 is equal")
    ((< 2 1) "2 is smaller than 1")
    (true "Always true"))`: true,
		`(import "somepackage")`: true,
		`'something`:             true,
		`(echo 'something)`:      true,
	}

	for input, expectedResult := range tt {
		t.Run(input, func(t *testing.T) {
			if validStatement(t, input) != expectedResult {
				t.Errorf("Test '%s' failed to match result: %t!", input, expectedResult)
			}
		})
	}
}
func checkTree(t *testing.T, input string, expectedTree []ast.Statement) {

	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer, false)

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
func matchIdentStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.IdentStatement)
	if !ok {
		t.Fatalf("Expected ident statement, got others")
	}
	a, ok := actualStatement.(*ast.IdentStatement)
	if !ok {
		t.Fatalf("Expected ident statement, got: %v", actualStatement)
	}
	return a.Value.TokenValue == e.Value.TokenValue && a.Value.TokenType == e.Value.TokenType
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
func matchQuoteStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.QuoteStatement)
	if !ok {
		t.Fatalf("Expected quote statement, got others")
	}
	a, ok := actualStatement.(*ast.QuoteStatement)
	if !ok {
		t.Fatalf("Expected quote statement, got others")
	}
	return a.Body.TokenValue == e.Body.TokenValue && a.Body.TokenType == e.Body.TokenType
}
func matchSetStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.SetStatement)
	if !ok {
		t.Fatalf("Expected set statement, got others")
	}
	a, ok := actualStatement.(*ast.SetStatement)
	if !ok {
		t.Fatalf("Expected set statement, got others")
	}
	return matchSliceStatement(t, e.Target, a.Target) && matchStatement(t, e.Value, a.Value)
}

func matchSliceStatement(t *testing.T, expectedStatement ast.Statement, actualStatement ast.Statement) bool {
	e, ok := expectedStatement.(*ast.IndexStatement)
	if !ok {
		t.Fatalf("Expected slice statement, got others")
	}
	a, ok := actualStatement.(*ast.IndexStatement)
	if !ok {
		t.Fatalf("Expected slice statement, got others")
	}

	return matchStatement(t, e.Key, a.Key) && matchStatement(t, e.Target, a.Target)
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
		case *ast.QuoteStatement:
			return matchQuoteStatement(t, expectedStatement, actualStatement)
		case *ast.IndexStatement:
			return matchSliceStatement(t, expectedStatement, actualStatement)
		case *ast.IdentStatement:
			return matchIdentStatement(t, expectedStatement, actualStatement)
		case *ast.SetStatement:
			return matchSetStatement(t, expectedStatement, actualStatement)
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
	parser := NewParser(lexer, false)

	program := parser.Parse()
	if len(parser.Errors()) != 0 {
		return false
	}

	snaps.MatchSnapshot(t, program)

	return true
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
