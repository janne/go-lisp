package lisp

import (
	"fmt"
	"strconv"
)

type Value struct {
	typ valueType
	val interface{}
}

var Nil = Value{nilValue, nil}
var False = Value{symbolValue, "false"}
var True = Value{symbolValue, "true"}

type valueType uint8

const (
	nilValue valueType = iota
	symbolValue
	numberValue
	stringValue
	vectorValue
	procValue
	consValue
)

func (v Value) Eval() (Value, error) {
	switch v.typ {
	case consValue:
		return v.Cons().Execute()
	case symbolValue:
		sym := v.String()
		if v, ok := scope.Get(sym); ok {
			return v, nil
		} else if sym == "true" || sym == "false" {
			return Value{symbolValue, sym}, nil
		} else {
			return Nil, fmt.Errorf("Unbound variable: %v", sym)
		}
	default:
		return v, nil
	}
}

func (v Value) String() string {
	switch v.typ {
	case numberValue:
		return strconv.FormatFloat(v.val.(float64), 'f', -1, 64)
	case nilValue:
		return "()"
	default:
		return fmt.Sprintf("%v", v.val)
	}
}

func (v Value) Inspect() string {
	switch v.typ {
	case stringValue:
		return fmt.Sprintf(`"%v"`, v.val)
	case vectorValue:
		return v.val.(Vector).Inspect()
	default:
		return v.String()
	}
}

func (v Value) Cons() Cons {
	if v.typ == consValue {
		return *v.val.(*Cons)
	} else {
		return Cons{&v, &Nil}
	}
}

func (v Value) Number() float64 {
	return v.val.(float64)
}
