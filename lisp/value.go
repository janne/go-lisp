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
	invalidValue valueType = iota
	nilValue
	symbolValue
	numberValue
	stringValue
	sexpValue
	procValue
	consValue
)

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
	case sexpValue:
		return v.val.(Sexp).Inspect()
	default:
		return v.String()
	}
}

func (v Value) Sexp() Sexp {
	return v.val.(Sexp)
}

func (v Value) Proc() Proc {
	return v.val.(Proc)
}

func (v Value) Number() float64 {
	return v.val.(float64)
}
