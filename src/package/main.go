package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"lorikeet/compiler"
	"lorikeet/lexer"
	"lorikeet/object"
	"lorikeet/parser"
	"lorikeet/vm"
)

var input = `
let factorize = fn(n) {
  let f = fn(fac, i) {
    if (i < 1 ) { return fac; }
    $f(fac * i, i - 1);
  };
  $f(1, n);
}

let get_num = fn() {
  let num = ask("Enter a number:");
  let int_num = int(num);
  if (!int_num) {
    say("Sorry, "+num+" is not a number.", "Try again!");
    $get_num();
  }
  int_num;
}

let system_loop = fn() {
  say(get_num() |> factorize());
  $system_loop();
}
system_loop();

`

func main() {
	flag.Parse()

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s", err)
		return
	}

	object.Scanner = bufio.NewScanner(os.Stdin)

	machine := vm.New(comp.Bytecode())

	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
		return
	}
}
