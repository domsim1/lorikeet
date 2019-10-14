package compiler

// SymbolScope represents scope name
type SymbolScope string

// Scope names
const (
	GlobalScope SymbolScope = "GLOBAL"
)

// Symbol struct holds symbol name,
// scope and an index
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable store for symbols
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable inits symbol tables
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// Define symbol in symbol table
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: GlobalScope}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve get symbol from symbol table
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
