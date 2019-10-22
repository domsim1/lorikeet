package object

import (
	"bufio"
	"fmt"
	"strconv"
)

// Scanner for get()
var Scanner *bufio.Scanner

// Builtins functions
var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"say",
		&Builtin{Fn: func(args ...Object) Object {
			for _, arg := range args {
				s, err := strconv.Unquote(`"` + arg.Inspect() + `"`)
				if err != nil {
					panic(err)
				}
				fmt.Print(s)
			}
			fmt.Print("\n")
			return nil
		},
		},
	},
	{
		"ask",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) > 1 {
				return newError("wrong number of arguments, no more than 1. got=%d",
					len(args))
			}

			if len(args) == 1 {
				if args[0].Type() != STRING {
					return newError("argument to `say` must be STRING, got %s",
						args[0].Type())
				}

				arr := args[0].(*String)
				fmt.Print(arr.Inspect())
			}

			scanned := Scanner.Scan()
			if !scanned {
				return nil
			}

			return &String{Value: Scanner.Text()}
		},
		},
	},
	{
		"head",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != ARRAY {
				return newError("argument to `head` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return nil
		},
		},
	},
	{
		"last",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != ARRAY {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return nil
		},
		},
	},
	{
		"tail",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != ARRAY {
				return newError("argument to `tail` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Array{Elements: newElements}
			}

			return nil
		},
		},
	},
	{
		"push",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != ARRAY {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			newElements := make([]Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Array{Elements: newElements}
		},
		},
	},
	{
		"int",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *Float:
				return &Integer{Value: int64(arg.Value)}
			case *String:
				val, err := strconv.ParseInt(arg.Value, 10, 64)
				if err != nil {
					flval, err := strconv.ParseFloat(arg.Value, 64)
					if err != nil {
						return nil
					}
					return &Integer{Value: int64(flval)}
				}
				return &Integer{Value: val}
			case *Integer:
				return arg
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"float",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *Integer:
				return &Float{Value: float64(arg.Value)}
			case *String:
				val, err := strconv.ParseFloat(arg.Value, 64)
				if err != nil {
					return nil
				}
				return &Float{Value: val}
			case *Float:
				return arg
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"string",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *Float:
				return &String{Value: fmt.Sprintf("%g", arg.Value)}
			case *Integer:
				return &String{Value: fmt.Sprintf("%d", arg.Value)}
			case *String:
				return arg
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
}

// GetBuiltinByName gets builtin function by name
func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
