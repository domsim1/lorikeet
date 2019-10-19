package object

// NewEnclosedEnvironment creates *Enviroments enclosed to code blocks
// this provides a scope for identifiers created in functions
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates *Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment struct
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get value stored in enviroment by identifier
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set value in enviroment by identifier
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
