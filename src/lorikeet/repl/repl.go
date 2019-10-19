package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"lorikeet/compiler"
	"lorikeet/lexer"
	"lorikeet/object"
	"lorikeet/parser"
	"lorikeet/vm"
)

// PROMPT characters
const PROMPT = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(scanLinesEscapable)
	object.Scanner = scanner

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Darn! Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Darn! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")

		// Evaluator
		// evaluator.DefineMacros(program, macroEnv)
		// expanded := evaluator.ExpandMacros(program, macroEnv)

		// evaluated := evaluator.Eval(expanded, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func scanLinesEscapable(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.LastIndexByte(data, '\n'); i >= 0 {
		if wasNewLineEscaped(data[0:i]) {
			// new line was escaped, request more data
			return 0, nil, nil
		}
		// We have a full newline-terminated line.
		return i + 1, dropCRBS(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCRBS(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// drop all /r and //
func dropCRBS(data []byte) []byte {
	if len(data) > 0 {
		return bytes.ReplaceAll(
			bytes.ReplaceAll(data, []byte("\r"), []byte("")),
			[]byte("\\\n"),
			[]byte("\n"))
	}
	return data
}

func wasNewLineEscaped(data []byte) bool {
	if len(data) > 0 {
		if data[len(data)-1] == '\r' {
			return wasNewLineEscaped(dropCR(data))
		}
		if data[len(data)-1] == '\\' {
			return true
		}
	}
	return false
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
