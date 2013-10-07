package lisp

import (
	"fmt"
	"strings"
)

type Sexp []Value

func (s Sexp) String() string {
	if len(s) == 1 {
		return s[0].String()
	} else {
		var arr []string
		for _, v := range s {
			arr = append(arr, v.String())
		}
		return fmt.Sprintf(`(%v)`, strings.Join(arr, " "))
	}
}

func (s Sexp) Inspect() string {
	var arr []string
	for _, v := range s {
		arr = append(arr, v.Inspect())
	}
	return fmt.Sprintf(`(%v)`, strings.Join(arr, " "))
}
