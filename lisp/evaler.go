package lisp

import "fmt"

func EvalString(line string) (Value, error) {
	parsed, err := NewTokens(line).Parse()
	if err != nil {
		return Nil, err
	}
	evaled, err := parsed.Eval()
	if err != nil {
		return Nil, err
	}
	scope.Create("_", evaled)
	return evaled, nil
}

func procForm(cons Cons) (val Value, err error) {
	if val, err = cons.car.Eval(); err == nil {
		if val.typ == procValue {
			var args Vector
			if args, err = cons.cdr.Cons().Map(func(v Value) (Value, error) {
				return v.Eval()
			}); err != nil {
				return
			} else {
				val, err = val.val.(Proc).Call(args)
			}
		} else {
			err = fmt.Errorf("The object %v is not applicable", val)
		}
	}
	return
}

func beginForm(cons Cons) (val Value, err error) {
	return cons.cdr.Cons().Eval()
}

func setForm(cons Cons) (val Value, err error) {
	expr := cons.Vector()
	if len(expr) == 3 {
		key := expr[1].String()
		if _, ok := scope.Get(key); ok {
			val, err = expr[2].Eval()
			if err == nil {
				scope.Set(key, val)
			}
		} else {
			err = fmt.Errorf("Unbound variable: %v", key)
		}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", cons)
	}
	return
}

func ifForm(cons Cons) (val Value, err error) {
	expr := cons.Vector()
	val = Nil
	if len(expr) < 3 || len(expr) > 4 {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	} else {
		r, err := expr[1].Eval()
		if err == nil {
			if !(r.typ == symbolValue && r.String() == "false") && r != Nil && len(expr) > 2 {
				val, err = expr[2].Eval()
			} else if len(expr) == 4 {
				val, err = expr[3].Eval()
			}
		}
	}
	return
}

func lambdaForm(cons Cons) (val Value, err error) {
	if cons.cdr.typ == consValue {
		lambda := cons.cdr.Cons()
		if (lambda.car.typ == consValue || lambda.car.typ == nilValue) && lambda.cdr.typ == consValue {
			params := lambda.car.Cons().Vector()
			val = Value{procValue, Proc{params, lambda.cdr.Cons(), scope.Dup()}}
		} else {
			err = fmt.Errorf("Ill-formed special form: %v", cons)
		}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", cons)
	}
	return
}

func quoteForm(cons Cons) (val Value, err error) {
	if cons.cdr != nil {
		if *cons.cdr.Cons().cdr == Nil {
			val = *cons.cdr.Cons().car
		} else {
			val = Value{consValue, cons}
		}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", cons)
	}
	return
}

func defineForm(cons Cons) (val Value, err error) {
	expr := cons.Vector()
	if len(expr) >= 2 && len(expr) <= 3 {
		if expr[1].typ == symbolValue {
			key := expr[1].String()
			if len(expr) == 3 {
				var i Value
				if i, err = expr[2].Eval(); err == nil {
					scope.Create(key, i)
				}
			} else {
				scope.Create(key, Nil)
			}
			return expr[1], err
		}
	}
	return Nil, fmt.Errorf("Ill-formed special form: %v", expr)
}
