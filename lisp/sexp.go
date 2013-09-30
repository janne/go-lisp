package lisp

import (
	"fmt"
)

type Sexp []Value

func (s Sexp) String() string {
	if len(s) == 1 {
		return fmt.Sprintf("%v", s[0])
	} else {
		return fmt.Sprintf("%v", []Value(s))
	}
}
