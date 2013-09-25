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

func isBuiltin(c Value) bool {
	if s, ok := c.(string); ok {
		if _, ok := builtin_commands[s]; ok {
			return true
		}
	}
	return false
}

func runBuiltin(expr []Value) (val Value, err error) {
	cmd := builtin_commands[expr[0].(string)]
	values := []reflect.Value{}
	for _, i := range expr[1:] {
		if value, err := evalValue(i); err != nil {
			return nil, err
		} else {
			values = append(values, reflect.ValueOf(value))
		}
	}
	result := reflect.ValueOf(&builtin).MethodByName(cmd).Call(values)
	val = result[0].Interface()
	err, _ = result[1].Interface().(error)
	return
}

func (Builtin) Display(vars ...Value) (Value, error) {
	var interfaces []interface{}
	for _, v := range vars {
		interfaces = append(interfaces, v)
	}
	fmt.Println(interfaces...)
	return nil, nil
}

func (Builtin) Add(vars ...Value) (Value, error) {
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

func (Builtin) Sub(vars ...Value) (Value, error) {
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

func (Builtin) Mul(vars ...Value) (Value, error) {
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

func (Builtin) Gt(vars ...Value) (Value, error) {
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

func (Builtin) Lt(vars ...Value) (Value, error) {
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

func (Builtin) Gte(vars ...Value) (Value, error) {
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

func (Builtin) Lte(vars ...Value) (Value, error) {
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
