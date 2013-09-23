package lisp

import "fmt"

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
		if v, ok := Env[sym]; ok {
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
				val, err = evalValue(t)
				if p, ok := val.(Proc); ok {
					val, err = p.Call(expr[1:])
				} else {
					err = fmt.Errorf("The object %v is not applicable", val)
				}
			} else if t == "quote" { // Quote
				if len(expr) == 2 {
					val = expr[1]
				} else {
					err = fmt.Errorf("Ill-formed special form: %v", expr)
				}
			} else if t == "define" { // Define
				if len(expr) < 2 || len(expr) > 3 {
					err = fmt.Errorf("Ill-formed special form: %v", expr)
				} else if len(expr) == 3 {
					val, err = evalValue(expr[2])
				}
				if err == nil {
					Env[expr[1].(string)] = val
				}
			} else if t == "set!" { // Set!
				if len(expr) == 3 {
					key := expr[1].(string)
					if _, ok := Env[key]; ok {
						val, err = evalValue(expr[2])
						if err == nil {
							Env[key] = val
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
			} else if t == "+" { // Addition
				var sum int
				for _, i := range expr[1:] {
					j, err := evalValue(i)
					if err == nil {
						v, ok := j.(int)
						if ok {
							sum += int(v)
						} else {
							err = fmt.Errorf("Cannot only add numbers: %v", i)
							break
						}
					}
				}
				val = sum
			} else {
				val, err = evalValue(t)
				if p, ok := val.(Proc); ok {
					val, err = p.Call(expr[1:])
				} else {
					err = fmt.Errorf("The object %v is not applicable", val)
				}
			}
		}
	default:
		err = fmt.Errorf("Unknown data type: %v", input)
	}
	return
}
