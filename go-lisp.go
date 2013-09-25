package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/janne/go-lisp/lisp"
	"io/ioutil"
	"os"
)

func main() {
	var file = flag.String("file", "", "File to interpret")
	var repl = flag.Bool("i", false, "Interactive mode")
	flag.Parse()

	if *repl {
		Repl()
	} else if *file != "" {
		if output, err := ioutil.ReadFile(*file); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			if _, err := lisp.EvalString(string(output)); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	} else {
		if output, err := ioutil.ReadAll(bufio.NewReader(os.Stdin)); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			lisp.EvalString(string(output))
		}
	}
}

func Repl() {
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
