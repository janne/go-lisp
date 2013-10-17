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

func (v Value) Eval() (val Value, err error) {
	switch v.typ {
	case consValue:
		cons := v.Cons()
		if !cons.List() {
			return Nil, fmt.Errorf("Combination must be a proper list: %v", cons)
		}
		switch cons.car.String() {
		case "quote":
			return cons.quoteForm()
		case "if":
			return cons.ifForm()
		case "set!":
			return cons.setForm()
		case "define":
			return cons.defineForm()
		case "lambda":
			return cons.lambdaForm()
		case "begin":
			return cons.beginForm()
		default:
			if cons.isBuiltin() {
				return cons.runBuiltin()
			} else {
				return cons.procForm()
			}
		}
	case numberValue, stringValue, vectorValue, nilValue:
		val = v
	case symbolValue:
		sym := v.String()
		if v, ok := scope.Get(sym); ok {
			val = v
		} else if sym == "true" || sym == "false" {
			val = Value{symbolValue, sym}
		} else {
			return Nil, fmt.Errorf("Unbound variable: %v", sym)
		}
	default:
		return Nil, fmt.Errorf("Unknown data type: %v", v)
	}
	return
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
