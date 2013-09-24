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
				if val, err = evalValue(t); err == nil {
					val, err = runProc(val, expr[1:])
				}
			} else if t == "quote" { // Quote
				if len(expr) == 2 {
					val = expr[1]
				} else {
					err = fmt.Errorf("Ill-formed special form: %v", expr)
				}
			} else if t == "define" { // Define
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
				err = fmt.Errorf("Ill-formed special form: %v", expr)
			} else if t == "set!" { // Set!
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
			} else if t == "if" { // If
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
			} else if t == "begin" { // Begin
				val, err = Eval(expr[1:])
			} else if t == "lambda" {
				if len(expr) > 2 {
					params := expr[1].(Sexp)
					val = Proc{params, expr[2:]}
				} else {
					err = fmt.Errorf("Ill-formed special form: %v", expr)
				}
			} else if isBuiltin(t) { // Addition
				args := make([]interface{}, 0)
				for _, i := range expr[1:] {
					if value, err := evalValue(i); err != nil {
						return nil, err
					} else {
						args = append(args, value)
					}
				}
				val, err = runBuiltin(t.(string), args)
			} else {
				if val, err = evalValue(t); err == nil {
					val, err = runProc(val, expr[1:])
				}
			}
		}
	default:
		err = fmt.Errorf("Unknown data type: %v", input)
	}
	return
}

func runProc(proc interface{}, vars []interface{}) (val interface{}, err error) {
	if p, ok := proc.(Proc); ok {
		val, err = p.Call(vars)
	} else {
		err = fmt.Errorf("The object %v is not applicable", proc)
	}
	return
}
