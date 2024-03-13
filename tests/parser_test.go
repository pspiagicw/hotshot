package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/token"
)

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
		"(let 1 2)": true,

		"(% 1 2)": false,
		"(# 1 2)": false,
		"(; 1 2)": false,
		"(@ 1 2)": false,
		"(, 1 2)": false,
		"(! 1 2)": false,
		"(? 1 2)": false,
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
	}

	for input, expectedResult := range tt {
		t.Run(input, func(t *testing.T) {
			if validStatement(t, input) != expectedResult {
				t.Errorf("Test '%s' failed to match result: %t!", input, expectedResult)
			}
		})
	}
}
