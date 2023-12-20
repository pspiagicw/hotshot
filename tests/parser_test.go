package tests

import (
	"testing"

	"github.com/pspiagicw/hotshot/ast"
)

func TestStatement(t *testing.T) {

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
