package lisp

import "fmt"

var builtins = map[string]string{
	"+": "add",
}

func isBuiltin(c interface{}) bool {
	if s, ok := c.(string); ok {
		if _, ok := builtins[s]; ok {
			return true
		}
	}
	return false
}

func runBuiltin(c string, args []interface{}) (val interface{}, err error) {
	cmd := builtins[c]
	if cmd == "add" {
		val, err = add(args...)
	}
	return
}

func add(vars ...interface{}) (interface{}, error) {
	var sum int
	for _, v := range vars {
		if i, ok := v.(int); ok {
			sum += i
		} else {
			return nil, fmt.Errorf("Can only add numbers: %v", i)
		}
	}
	return sum, nil
}
