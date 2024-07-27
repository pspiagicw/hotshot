package compiler

import "github.com/pspiagicw/hotshot/object"

type SymbolScope string

const (
	Global SymbolScope = "GLOBAL"
	Local  SymbolScope = "LOCAL"
	Built  SymbolScope = "BUILTIN"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	Outer *SymbolTable

	store          map[string]Symbol
	numDefinitions int
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	table := &SymbolTable{store: s}

	for i, builtin := range object.BuiltinList() {
		table.DefineBuiltin(i, builtin.Name)
	}

	return table
}

func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Scope: Built, Index: index}
	s.store[name] = symbol
	return symbol
}
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: Global}
	if s.Outer != nil {
		symbol.Scope = Local
	}

	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
	}
	return obj, ok
}
