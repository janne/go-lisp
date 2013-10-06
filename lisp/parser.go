package lisp

import (
	"fmt"
	"strconv"
)

func Parse(tokens []*Token) (Sexp, error) {
	var pos int
	values := make(Sexp, 0)
	for pos < len(tokens) {
		t := tokens[pos]
		switch t.typ {
		case numberType:
			if i, err := strconv.ParseFloat(t.val, 64); err != nil {
				return nil, fmt.Errorf("Failed to convert number: %v", t.val)
			} else {
				values = append(values, NewValue(i))
				pos++
			}
		case stringType, symbolType:
			values = append(values, NewValue(t.val))
			pos++
		case openType:
			start := pos + 1
			end, err := findEnd(tokens, start)
			if err != nil {
				return nil, err
			}
			x, err := Parse(tokens[start:end])
			if err != nil {
				return nil, err
			}
			values = append(values, NewValue(x))
			pos = end + 1
		case closeType:
			return nil, fmt.Errorf("List was closed but not opened")
		}
	}
	return values, nil
}

func findEnd(tokens []*Token, start int) (int, error) {
	depth := 1

	for i := start; i < len(tokens); i++ {
		t := tokens[i]
		switch t.typ {
		case openType:
			depth++
		case closeType:
			depth--
		}
		if depth == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("List was opened but not closed")
}
