package compiler

import (
	"testing"

	"github.com/pspiagicw/hotshot/code"
)

func TestScopes(t *testing.T) {
	compiler := NewCompiler()

	if compiler.scopeIndex != 0 {
		t.Errorf("scopeIndex wrong. got=%d", compiler.scopeIndex)
	}
	globalSymbolTable := compiler.symbols

	compiler.emit(code.ADD, -1)
	compiler.enterScope()
	if compiler.scopeIndex != 1 {
		t.Errorf("scopeIndex wrong. got=%d", compiler.scopeIndex)
	}

	compiler.emit(code.SUB, -1)
	if len(compiler.scopes[compiler.scopeIndex].instructions) != 1 {
		t.Errorf("instructions length wrong. got=%d", len(compiler.scopes[compiler.scopeIndex].instructions))
	}

	if compiler.symbols.Outer != globalSymbolTable {
		t.Errorf("compiler did not enclose symbol table properly")
	}

	compiler.leaveScope()
	if compiler.scopeIndex != 0 {
		t.Errorf("scopeIndex wrong. got=%d", compiler.scopeIndex)
	}

	if compiler.symbols != globalSymbolTable {
		t.Errorf("compiler did not restore global symbol table")
	}
	if compiler.symbols.Outer != nil {
		t.Errorf("compiler modified global symbol table incorrectly")
	}

	compiler.emit(code.MUL, -1)

	if len(compiler.scopes[compiler.scopeIndex].instructions) != 2 {
		t.Errorf("instructions length wrong. got=%d", len(compiler.scopes[compiler.scopeIndex].instructions))
	}
}
