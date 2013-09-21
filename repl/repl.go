package main

import (
	"fmt"
	"bufio"
	"os"
	"lisp"
)

func main() {
	fmt.Printf("Hello, world!\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, _ := reader.ReadString('\n')
		if response, err := lisp.Execute(line); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			fmt.Printf("%v\n", response)
		}
	}
}
