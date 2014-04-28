package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/janne/go-lisp/lisp"
	"io/ioutil"
	"os"
	"strings"
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
	expr := ""
	for {
		if expr == "" {
			fmt.Printf("\n> ")
		}
		line, _ := reader.ReadString('\n')
		expr = fmt.Sprintf("%v%v", expr, line)
		openCount := strings.Count(expr, "(")
		closeCount := strings.Count(expr, ")")
		if openCount < closeCount {
			fmt.Printf("ERROR: Malformed expression: %v", line)
			expr = ""
		} else if openCount == closeCount {
			if strings.TrimSpace(expr) != "" {
				if response, err := lisp.EvalString(expr); err != nil {
					fmt.Printf("ERROR: %v\n", err)
				} else {
					if response == lisp.Nil {
						fmt.Println(";Unspecified return value")
					} else {
						fmt.Printf(";Value: %v\n", response.Inspect())
					}
				}
			}
			expr = ""
		}
	}
}
