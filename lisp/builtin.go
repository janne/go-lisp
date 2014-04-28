package lisp

import "fmt"

type Builtin struct{}

var builtin = Builtin{}

var builtin_commands = map[string]string{
	"+":             "Add",
	"-":             "Sub",
	"*":             "Mul",
	">":             "Gt",
	"<":             "Lt",
	">=":            "Gte",
	"<=":            "Lte",
	"display":       "Display",
	"cons":          "Cons",
	"car":           "Car",
	"cdr":           "Cdr",
	"string?":       "StringHuh",
	"string=?":      "StringEqualHuh",
	"string-length": "StringLength",
	"string-append": "StringAppend",
}

func (Builtin) Display(vars ...Value) (Value, error) {
	if len(vars) == 1 {
		fmt.Println(vars[0])
	} else {
		return badlyFormattedArguments(vars)
	}
	return Nil, nil
}

func (Builtin) Cons(vars ...Value) (Value, error) {
	if len(vars) == 2 {
		cons := Cons{&vars[0], &vars[1]}
		return Value{consValue, &cons}, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Car(vars ...Value) (Value, error) {
	if len(vars) == 1 && vars[0].typ == consValue {
		cons := vars[0].Cons()
		return *cons.car, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Cdr(vars ...Value) (Value, error) {
	if len(vars) == 1 && vars[0].typ == consValue {
		cons := vars[0].Cons()
		return *cons.cdr, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Add(vars ...Value) (Value, error) {
	sum := 0.0
	for _, v := range vars {
		if v.typ == numberValue {
			sum += v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Sub(vars ...Value) (Value, error) {
	if len(vars) == 0 || vars[0].typ != numberValue {
		return badlyFormattedArguments(vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum -= v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Mul(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return Value{numberValue, 1.0}, nil
	}
	if vars[0].typ != numberValue {
		return badlyFormattedArguments(vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum *= v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Gt(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() > v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lt(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() < v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Gte(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() >= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lte(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() <= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) StringHuh(vars ...Value) (Value, error) {
	if len(vars) != 1 {
		return badlyFormattedArguments(vars)
	}
	if vars[0].typ == stringValue {
		return True, nil
	}
	return False, nil
}

func (Builtin) StringEqualHuh(vars ...Value) (Value, error) {
	if len(vars) < 2 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != stringValue || v2.typ != stringValue {
			return badlyFormattedArguments(vars)
		} else if v1.String() != v2.String() {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) StringLength(vars ...Value) (Value, error) {
	if len(vars) != 1 || vars[0].typ != stringValue {
		return badlyFormattedArguments(vars)
	}
	return Value{numberValue, float64(len(vars[0].String()))}, nil
}

func (Builtin) StringAppend(vars ...Value) (Value, error) {
	result := ""
	for i := 0; i < len(vars); i++ {
		v := vars[i]
		if v.typ != stringValue {
			return badlyFormattedArguments(vars)
		} else {
			result = result + v.String()
		}
	}
	return Value{stringValue, result}, nil
}

func badlyFormattedArguments(vars []Value) (Value, error) {
	return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
}
