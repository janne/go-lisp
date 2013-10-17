package lisp

import "fmt"

type Proc struct {
	params Vector
	body   Cons
	scope  *Scope
}

func (p Proc) String() string {
	return "<Procedure>"
}

func (p Proc) Call(params Vector) (val Value, err error) {
	if len(p.params) == len(params) {
		old_scope := scope
		scope = p.scope
		scope.AddEnv()
		for i, name := range p.params {
			scope.Create(name.String(), params[i])
		}
		val, err = p.body.Eval()
		scope.DropEnv()
		scope = old_scope
	} else {
		err = fmt.Errorf("%v has been called with %v arguments; it requires exactly %v arguments", p, len(params), len(p.params))
	}
	return
}
