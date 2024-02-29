package printer

import (
	"strings"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/shivamMg/ppds/tree"
)

func PrintAST(ast *ast.Program) string {
	var output strings.Builder
	output.WriteString(tree.SprintHrn(ast))
	return output.String()
}
