package printer

import (
	"strings"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/shivamMg/ppds/tree"
)

func PrintAST(ast *ast.Program) string {
	var output strings.Builder
	for _, statement := range ast.Statements {
		if statement != nil {
			output.WriteString(tree.SprintHrn(statement))
		} else {
			output.WriteString("NIL\n")
		}
	}
	return output.String()
}
