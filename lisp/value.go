package lisp

import (
	"fmt"
	"strconv"
)

type Value struct {
	Kind Kind
	raw  interface{}
}

var Nil = Value{NilKind, nil}
var False = Value{SymbolKind, "false"}
var True = Value{SymbolKind, "true"}

type Kind uint

const (
	InvalidKind Kind = iota
	NilKind
	SymbolKind
	NumberKind
	StringKind
	SexpKind
	ProcKind
)

func (v Value) IsA(k Kind) bool {
	return k == v.Kind
}

func (v Value) String() string {
	switch v.Kind {
	case NumberKind:
		return strconv.FormatFloat(v.raw.(float64), 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v.raw)
	}
}

func (v Value) Sexp() Sexp {
	return v.raw.(Sexp)
}

func (v Value) Proc() Proc {
	return v.raw.(Proc)
}

func (v Value) Number() float64 {
	return v.raw.(float64)
}
