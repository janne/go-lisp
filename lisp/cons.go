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

func (c Cons) Map(f func (v Value) (Value, error)) ([]Value, error) {
	result := make([]Value, 0)
	if *c.car != Nil {
		if value, err := f(*c.car); err != nil {
			return nil, err
		} else {
			result = append(result, value)
		}
		if *c.cdr != Nil {
			if c.cdr.typ == consValue {
				if values, err := c.cdr.Cons().Map(f); err != nil {
					return nil, err
				} else {
					result = append(result, values...)
				}
			} else {
				if value, err := f(*c.cdr); err != nil {
					return nil, err
				} else {
					result = append(result, value)
				}
			}
		}
	}
	return result, nil
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

func (c Cons) Sexp() Sexp {
	v, _ := c.Map(func (v Value) (Value, error) {
		return v, nil
	})
	return v
}

func (c Cons) String() string {
	arr := []string{}
	for _, v := range c.Sexp() {
		arr = append(arr, v.String())
	}
	return fmt.Sprintf(`(%v)`, strings.Join(arr, " "))
}
