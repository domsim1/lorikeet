package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"lorikeet/compiler"
	"lorikeet/lexer"
	"lorikeet/parser"
	"lorikeet/repl"
	"lorikeet/vm"
	"os"
)

var file string

func main() {
	flag.StringVar(&file, "file", "", "Usage")
	flag.Parse()
	if file == "" {
		fmt.Printf("The Lorikeet programming language!\n")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("could not read file!\n")
	}

	l := lexer.New(string(data))
	p := parser.New(l)
	program := p.ParseProgram()

	comp := compiler.New()
	compErr := comp.Compile(program)
	if compErr != nil {
		fmt.Printf("compiler error: %s", err)
		return
	}

	machine := vm.New(comp.Bytecode())

	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
		return
	}

}
