package lisp

import "fmt"

type Proc struct {
	Params Sexp
	Body   Sexp
}

func (p Proc) Call(params Sexp) (val interface{}, err error) {
	if len(p.Params) == len(params) {
		for i, name := range p.Params {
			Env[name.(string)] = params[i]
		}
		val, err = Eval(p.Body)
	} else {
		err = fmt.Errorf("Number of parameters mismatch, %v for %v", len(params), len(p.Params))
	}
	return
}
