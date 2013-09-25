package main

import (
	"fmt"
	"bufio"
	"os"
	"lisp"
)

func main() {
	fmt.Printf("Welcome to the Lisp REPL\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, _ := reader.ReadString('\n')
		if response, err := lisp.EvalString(line); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			fmt.Printf("=> %v\n", response)
		}
	}
}
