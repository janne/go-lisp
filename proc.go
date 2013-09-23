package lisp

import "fmt"

type Proc struct {
	Params Sexp
	Body   Sexp
}

func (p Proc) String() string {
	return "<Procedure>"
}

func (p Proc) Call(params Sexp) (val interface{}, err error) {
	if len(p.Params) == len(params) {
		for i, name := range p.Params {
			Env[name.(string)] = params[i]
		}
		val, err = Eval(p.Body)
	} else {
		err = fmt.Errorf("%v has been called with %v arguments; it requires exactly %v arguments", p, len(params), len(p.Params))
	}
	return
}
