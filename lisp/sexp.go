package lisp

import (
	"fmt"
	"strings"
)

type Sexp []Value

func (s Sexp) String() string {
	var arr []string
	for _, v := range s {
		arr = append(arr, v.String())
	}
	return fmt.Sprintf(`[%v]`, strings.Join(arr, " "))
}

func (s Sexp) Inspect() string {
	var arr []string
	for _, v := range s {
		arr = append(arr, v.Inspect())
	}
	return fmt.Sprintf(`[%v]`, strings.Join(arr, " "))
}
