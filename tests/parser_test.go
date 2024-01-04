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
func TestAddition(t *testing.T) {
	input := `(+ 1 2)`

	expectedTree := []ast.Statement{
		&ast.FunctionalStatement{
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
		&ast.FunctionalStatement{
			Op: &token.Token{
				TokenType:  token.PLUS,
				TokenValue: "+",
			},
			Args: []ast.Statement{
				&ast.FunctionalStatement{
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
