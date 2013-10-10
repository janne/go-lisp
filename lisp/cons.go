package lisp

import (
	"fmt"
	"strings"
)

type Cons struct {
	car *Value
	cdr *Value
}

func (c Cons) Sexp() (s Sexp) {
	s = append(s, *c.car)
	if c.cdr.typ == consValue {
		s = append(s, c.cdr.val.(Cons).Sexp()...)
	} else if c.cdr.typ != nilValue {
		s = append(s, *c.cdr)
	}
	return
}

func (c Cons) String() string {
	arr := []string{}
	arr = append(arr, c.car.String())
	if c.cdr.typ != sexpValue {
		arr = append(arr, ".")
	}
	arr = append(arr, c.cdr.String())
	return fmt.Sprintf(`(%v)`, strings.Join(arr, " "))
}
