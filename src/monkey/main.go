package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Printf("This is the Monkey programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
