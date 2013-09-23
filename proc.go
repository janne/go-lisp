package lisp

import "fmt"

type Proc struct {
	params Sexp
	body   Sexp
}

func (p Proc) String() string {
	return "<Procedure>"
}

func (p Proc) Call(params Sexp) (val interface{}, err error) {
	if len(p.params) == len(params) {
		for i, name := range p.params {
			Env[name.(string)] = params[i]
		}
		val, err = Eval(p.body)
	} else {
		err = fmt.Errorf("%v has been called with %v arguments; it requires exactly %v arguments", p, len(params), len(p.params))
	}
	return
}
