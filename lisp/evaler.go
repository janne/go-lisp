package lisp

import "fmt"

var scope *Scope

func init() {
	scope = NewScope()
	scope.AddEnv()
}

func EvalString(line string) (string, error) {
	tokenized := Tokenize(line)
	parsed, err := Parse(tokenized)
	if err != nil {
		return "", err
	}
	evaled, err := Eval(parsed)
	if err != nil {
		return "", err
	}
	scope.Create("_", evaled)
	return fmt.Sprintf("%v", evaled), nil
}

func Eval(expr Sexp) (val interface{}, err error) {
	for _, t := range expr {
		val, err = evalValue(t)
		if err != nil {
			break
		}
	}
	return
}

func evalList(expr []interface{}) (args []interface{}, err error) {
	for _, i := range expr {
		if value, err := evalValue(i); err != nil {
			return nil, err
		} else {
			args = append(args, value)
		}
	}
	return
}

func evalValue(input interface{}) (val interface{}, err error) {
	switch input.(type) {
	case int: // Int
		val = input
	case string: // Symbol
		sym := input.(string)
		if v, ok := scope.Get(sym); ok {
			val = v
		} else if sym == "true" || sym == "false" {
			val = sym
		} else {
			err = fmt.Errorf("Unbound variable: %v", sym)
		}
	case Sexp:
		expr := input.(Sexp)
		if len(expr) > 0 {
			t := expr[0]
			if _, ok := t.(Sexp); ok {
				val, err = procForm(expr)
			} else if t == "quote" {
				val, err = quoteForm(expr)
			} else if t == "if" {
				val, err = ifForm(expr)
			} else if t == "set!" {
				val, err = setForm(expr)
			} else if t == "define" {
				val, err = defineForm(expr)
			} else if t == "lambda" {
				val, err = lambdaForm(expr)
			} else if t == "begin" {
				val, err = beginForm(expr)
			} else if isBuiltin(t) {
				var args []interface{}
				args, err = evalList(expr[1:])
				if err == nil {
					val, err = runBuiltin(t.(string), args)
				}
			} else {
				val, err = procForm(expr)
			}
		}
	default:
		err = fmt.Errorf("Unknown data type: %v", input)
	}
	return
}

func procForm(expr []interface{}) (val interface{}, err error) {
	if val, err = evalValue(expr[0]); err == nil {
		val, err = runProc(val, expr[1:])
	}
	return
}

func beginForm(expr []interface{}) (val interface{}, err error) {
	return Eval(expr[1:])
}

func setForm(expr []interface{}) (val interface{}, err error) {
	if len(expr) == 3 {
		key := expr[1].(string)
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

func ifForm(expr []interface{}) (val interface{}, err error) {
	if len(expr) < 3 || len(expr) > 4 {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	} else {
		r, err := evalValue(expr[1])
		if err == nil {
			if r != "false" && r != nil && len(expr) > 2 {
				val, err = evalValue(expr[2])
			} else if len(expr) == 4 {
				val, err = evalValue(expr[3])
			}
		}
	}
	return
}

func lambdaForm(expr []interface{}) (val interface{}, err error) {
	if len(expr) > 2 {
		params := expr[1].(Sexp)
		val = Proc{params, expr[2:], scope.Dup()}
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	}
	return
}

func quoteForm(expr []interface{}) (val interface{}, err error) {
	if len(expr) == 2 {
		val = expr[1]
	} else {
		err = fmt.Errorf("Ill-formed special form: %v", expr)
	}
	return
}

func defineForm(expr []interface{}) (val interface{}, err error) {
	if len(expr) >= 2 && len(expr) <= 3 {
		if key, ok := expr[1].(string); ok {
			if len(expr) == 3 {
				var i interface{}
				if i, err = evalValue(expr[2]); err == nil {
					scope.Create(key, i)
				}
			} else {
				scope.Create(key, nil)
			}
			return key, err
		}
	}
	return nil, fmt.Errorf("Ill-formed special form: %v", expr)
}

func runProc(proc interface{}, vars []interface{}) (val interface{}, err error) {
	if p, ok := proc.(Proc); ok {
		var args []interface{}
		for _, v := range vars {
			if e, err := evalValue(v); err != nil {
				return nil, err
			} else {
				args = append(args, e)
			}
		}
		val, err = p.Call(args)
	} else {
		err = fmt.Errorf("The object %v is not applicable", proc)
	}
	return
}
