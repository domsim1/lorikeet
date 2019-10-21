package compiler

import "fmt"

// SymbolScope represents scope name
type SymbolScope string

// Scope names
const (
	LocalScope    SymbolScope = "LOCAL"
	GlobalScope   SymbolScope = "GLOBAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FUNCTION"
)

// Symbol struct holds symbol name,
// scope and an index
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
	Mut   bool
}

// SymbolTable store for symbols
type SymbolTable struct {
	Outer *SymbolTable

	store          map[string]Symbol
	numDefinitions int

	FreeSymbols []Symbol
}

// NewSymbolTable inits symbol tables
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	free := []Symbol{}
	return &SymbolTable{store: s, FreeSymbols: free}
}

// Define symbol in symbol table
func (s *SymbolTable) Define(name string, mut bool) (Symbol, error) {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Mut: mut}
	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	obj, ok := s.Resolve(name)
	if ok {
		if symbol.Scope == obj.Scope {
			return obj, fmt.Errorf("symbol %s is already declared",
				obj.Name)
		}
	}

	s.store[name] = symbol
	s.numDefinitions++
	return symbol, nil
}

// Resolve get symbol from symbol table
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}

// NewEnclosedSymbolTable creates a symbol table
// with and Outer symbol table
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// DefineBuiltin symbols in the BuiltinScope with give name and index,
// this function ignores symbol table scope
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}

// DefineFunctionName creates a new symbol with FunctionScope
func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{Name: name, Index: 0, Scope: FunctionScope}
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)

	symbol := Symbol{Name: original.Name, Index: len(s.FreeSymbols) - 1}
	symbol.Scope = FreeScope

	s.store[original.Name] = symbol
	return symbol
}
