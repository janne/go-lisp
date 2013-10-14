package lisp

import (
	"fmt"
	"strings"
)

type Cons struct {
	car *Value
	cdr *Value
}

func (c Cons) List() bool {
	return c.cdr.typ == consValue || c.cdr.typ == nilValue
}

func (c Cons) Map(f func (v Value) Value) []Value {
	result := make([]Value, 0)
	if *c.car != Nil {
		if c.car.typ == consValue {
			result = append(result, Value{sexpValue, c.car.Cons().Map(f)})
		} else {
			result = append(result, *c.car)
		}
		if *c.cdr != Nil {
			if c.cdr.typ == consValue {
				result = append(result, c.cdr.Cons().Map(f)...)
			} else {
				result = append(result, *c.cdr)
			}
		}
	}
	return result
}

func (c Cons) Len() int {
	l := 0
	if *c.car != Nil {
		l++
		if *c.cdr != Nil {
			if c.cdr.typ == consValue {
				l += c.cdr.Cons().Len()
			} else {
				l++
			}
		}
	}
	return l
}

func (c Cons) Sexp() (s Sexp) {
	if c.car.typ == consValue {
		cons := c.car.val.(*Cons)
		s = append(s, Value{sexpValue, cons.Sexp()})
	} else if *c.car != Nil {
		s = append(s, *c.car)
	}
	if c.cdr.typ == consValue {
		cons := c.cdr.val.(*Cons)
		s = append(s, cons.Sexp()...)
	} else if *c.cdr != Nil {
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
