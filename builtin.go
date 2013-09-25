package lisp

import "fmt"
import "reflect"

type Builtin struct{}

var builtin = Builtin{}

var builtin_commands = map[string]string{
	"+":       "Add",
	"-":       "Sub",
	"*":       "Mul",
	">":       "Gt",
	"<":       "Lt",
	">=":      "Gte",
	"<=":      "Lte",
	"display": "Display",
}

func isBuiltin(c interface{}) bool {
	if s, ok := c.(string); ok {
		if _, ok := builtin_commands[s]; ok {
			return true
		}
	}
	return false
}

func runBuiltin(c string, args []interface{}) (val interface{}, err error) {
	cmd := builtin_commands[c]
	values := []reflect.Value{}
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}
	result := reflect.ValueOf(&builtin).MethodByName(cmd).Call(values)
	val = result[0].Interface()
	err, _ = result[1].Interface().(error)
	return
}

func (Builtin) Display(vars ...interface{}) (interface{}, error) {
	fmt.Println(vars...)
	return nil, nil
}

func (Builtin) Add(vars ...interface{}) (interface{}, error) {
	var sum int
	for _, v := range vars {
		if i, ok := v.(int); ok {
			sum += i
		} else {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return sum, nil
}

func (Builtin) Sub(vars ...interface{}) (interface{}, error) {
	sum, ok := vars[0].(int)
	if !ok {
		return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	for _, v := range vars[1:] {
		if i, ok := v.(int); ok {
			sum -= i
		} else {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return sum, nil
}

func (Builtin) Mul(vars ...interface{}) (interface{}, error) {
	sum, ok := vars[0].(int)
	if !ok {
		return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	for _, v := range vars[1:] {
		if i, ok := v.(int); ok {
			sum *= i
		} else {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return sum, nil
}

func (Builtin) Gt(vars ...interface{}) (interface{}, error) {
	for i := 1; i < len(vars); i++ {
		v1, ok1 := vars[i-1].(int)
		v2, ok2 := vars[i].(int)
		if !ok1 && !ok2 {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1 > v2) {
			return "false", nil
		}
	}
	return "true", nil
}

func (Builtin) Lt(vars ...interface{}) (interface{}, error) {
	for i := 1; i < len(vars); i++ {
		v1, ok1 := vars[i-1].(int)
		v2, ok2 := vars[i].(int)
		if !ok1 && !ok2 {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1 < v2) {
			return "false", nil
		}
	}
	return "true", nil
}

func (Builtin) Gte(vars ...interface{}) (interface{}, error) {
	for i := 1; i < len(vars); i++ {
		v1, ok1 := vars[i-1].(int)
		v2, ok2 := vars[i].(int)
		if !ok1 && !ok2 {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1 >= v2) {
			return "false", nil
		}
	}
	return "true", nil
}

func (Builtin) Lte(vars ...interface{}) (interface{}, error) {
	for i := 1; i < len(vars); i++ {
		v1, ok1 := vars[i-1].(int)
		v2, ok2 := vars[i].(int)
		if !ok1 && !ok2 {
			return nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1 <= v2) {
			return "false", nil
		}
	}
	return "true", nil
}
