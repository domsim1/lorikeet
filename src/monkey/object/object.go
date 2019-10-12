package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

// Type of object
type Type string

// Types
const (
	INTEGER  = "INTEGER"
	BOOLEAN  = "BOOLEAN"
	NULL     = "NULL"
	RETURN   = "RETURN_VALUE"
	ERROR    = "ERROR"
	FUNCTION = "FUNCTION"
	STRING   = "STRING"
	BUILTIN  = "BUILTIN"
	ARRAY    = "ARRAY"
	HASH     = "HASH"
	QUOTE    = "QUOTE"
)

// Object methods
type Object interface {
	Type() Type
	Inspect() string
}

// Function struct
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type will return the function type "FUNCTION"
func (f *Function) Type() Type { return FUNCTION }

// Inspect will return the function
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}\n")

	return out.String()
}

// Integer object
type Integer struct {
	Value int64
}

// Type will return the integer type "INTEGER"
func (i *Integer) Type() Type { return INTEGER }

// Inspect will return the Integer value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean object
type Boolean struct {
	Value bool
}

// Type will return the boolean type "BOOLEAN"
func (b *Boolean) Type() Type { return BOOLEAN }

// Inspect will return the boolean value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// String object
type String struct {
	Value string
}

// Type will return the string type "STRING"
func (s *String) Type() Type { return STRING }

// Inspect will return the string value
func (s *String) Inspect() string { return s.Value }

//Null struct
type Null struct{}

// Type will return the null type "NULL"
func (n *Null) Type() Type { return NULL }

// Inspect will return the null value
func (n *Null) Inspect() string { return "null" }

//ReturnValue object
type ReturnValue struct {
	Value Object
}

// Type will return the return value type "RETURN_VALUE"
func (rv *ReturnValue) Type() Type { return RETURN }

// Inspect will return the return value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error object
type Error struct {
	Message string
}

// Type will return the Error type "ERROR"
func (e *Error) Type() Type { return ERROR }

// Inspect will return the error message
func (e *Error) Inspect() string { return "Error: " + e.Message }

// BuiltinFunction type
type BuiltinFunction func(args ...Object) Object

// Builtin object
type Builtin struct {
	Fn BuiltinFunction
}

// Type will return the inbuilt function type "BUILTIN"
func (b *Builtin) Type() Type { return BUILTIN }

// Inspect will return "builtin function" string
func (b *Builtin) Inspect() string { return "builtin function" }

// Array object
type Array struct {
	Elements []Object
}

// Type will return the array type "ARRAY"
func (ao *Array) Type() Type { return ARRAY }

// Inspect will return array value
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey object
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey boolean
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey Int
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey String
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair object
type HashPair struct {
	Key   Object
	Value Object
}

// Hash object
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type will return the hash type "HASH"
func (h *Hash) Type() Type { return HASH }

// Inspect will return hash value
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable provides HashKey function to HashKey object
type Hashable interface {
	HashKey() HashKey
}

// Quote object
type Quote struct {
	Node ast.Node
}

// Type will return qoute type "QUOTE"
func (q *Quote) Type() Type { return QUOTE }

// Inspect will return quote value
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
