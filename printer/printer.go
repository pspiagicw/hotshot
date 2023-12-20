package printer

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/ast"
)

func PrintAST(ast *ast.Program) string {
	var output strings.Builder
	for i, statement := range ast.Statements {
		output.WriteString(fmt.Sprintf("[%d] ", i))
		if statement != nil {
			output.WriteString(statement.StringifyStatement())
		} else {
			output.WriteString("NIL Statement")
		}
		output.WriteString("\n")
	}
	return output.String()
}
