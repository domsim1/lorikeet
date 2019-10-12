package main

import (
	"fmt"
	"lorikeet/repl"
	"os"
)

func main() {
	fmt.Printf("The Lorikeet programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
