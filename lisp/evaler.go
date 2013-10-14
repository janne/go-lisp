package lisp

import "fmt"

func EvalString(line string) (Value, error) {
	parsed, err := NewTokens(line).Parse()
	if err != nil {
		return Nil, err
	}
	evaled, err := Eval(parsed)
	if err != nil {
		return Nil, err
	}
	scope.Create("_", evaled)
	return evaled, nil
}

func Eval(cons Cons) (val Value, err error) {
	for _, t := range cons.Sexp() {
		val, err = evalValue(t)
		if err != nil {
			break
		}
	}
	return
}

func evalValue(input Value) (val Value, err error) {
	switch input.typ {
	case consValue:
		cons := input.Cons()
		if !cons.List() {
			return Nil, fmt.Errorf("Combination must be a proper list: %v", cons)
		}
		switch cons.car.String() {
		case "quote":
			return quoteForm(cons)
		case "if":
			return ifForm(cons)
		case "set!":
			return setForm(cons)
		case "define":
			return defineForm(cons)
		case "lambda":
			return lambdaForm(cons)
		case "begin":
			return beginForm(cons)
		default:
			if isBuiltin(cons) {
				return runBuiltin(cons)
			} else {
				return procForm(cons)
			}
		}
	case numberValue, stringValue, sexpValue, nilValue:
		val = input
	case symbolValue:
		sym := input.String()
		if v, ok := scope.Get(sym); ok {
			val = v
		} else if sym == "true" || sym == "false" {
			val = Value{symbolValue, sym}
		} else {
			return Nil, fmt.Errorf("Unbound variable: %v", sym)
		}
	default:
		return Nil, fmt.Errorf("Unknown data type: %v", input)
	}
	return
}

func procForm(cons Cons) (val Value, err error) {
	if val, err = evalValue(*cons.car); err == nil {
		if val.typ == procValue {
			var args []Value
			for _, v := range cons.cdr.Sexp() {
				if e, err := evalValue(v); err != nil {
					return Nil, err
				} else {
					args = append(args, e)
				}
			}
			val, err = val.Proc().Call(args)
		} else {
			err = fmt.Errorf("The object %v is not applicable", val)
		}
	}
	return
}

func beginForm(cons Cons) (val Value, err error) {
	return Eval(cons.cdr.val.(Cons))
}

func setForm(cons Cons) (val Value, err error) {
	expr := cons.Sexp()
	if len(expr) == 3 {
		key := expr[1].String()
		if _, ok := scope.Get(key); ok {
			val, err = evalValue(expr[2])
			if err == nil {
				scope.Set(key, val)
			}
		} else {
			err = fmt.Errorf("Unbound variable: %v", key)
		}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	}
	return
}

func ifForm(cons Cons) (val Value, err error) {
	expr := cons.Sexp()
	if len(expr) < 3 || len(expr) > 4 {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	} else {
		r, err := evalValue(expr[1])
		if err == nil {
			if !(r.typ == symbolValue && r.String() == "false") && r != Nil && len(expr) > 2 {
				val, err = evalValue(expr[2])
			} else if len(expr) == 4 {
				val, err = evalValue(expr[3])
			}
		}
	}
	return
}

func lambdaForm(cons Cons) (val Value, err error) {
	if cons.cdr.typ == consValue {
		lambda := cons.cdr.val.(Cons)
		params := lambda.car.Sexp()
		val = Value{procValue, Proc{params, lambda.cdr.val.(Cons), scope.Dup()}}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", cons)
	}
	return
}

func quoteForm(cons Cons) (val Value, err error) {
	if cons.cdr != nil {
		val = *cons.cdr
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", cons)
	}
	return
}

func defineForm(cons Cons) (val Value, err error) {
	expr := cons.Sexp()
	if len(expr) >= 2 && len(expr) <= 3 {
		if expr[1].typ == symbolValue {
			key := expr[1].String()
			if len(expr) == 3 {
				var i Value
				if i, err = evalValue(expr[2]); err == nil {
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
