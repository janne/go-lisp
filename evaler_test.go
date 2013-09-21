package lisp

import "testing"
import "fmt"

func TestEval(t *testing.T) {
	var tests = map[string]interface{}{
		"42":                                      "42",
		"(42)":                                    "42",
		"((42))":                                  "42",
		"(+ 42 13)":                               "55",
		"(+ (+ 1 2 3) 4)":                         "10",
		"(quote 1 2 3)":                           "[1 2 3]",
		"(quote 1 (+ 1 2) 3)":                     "[1 [+ 1 2] 3]",
		"(if true (+ 1 1) 3)":                     "2",
		"(if false 42 1)":                         "1",
		"(if false 42)":                           "<nil>",
		"(begin (define x) (if x 1 2))":           "2",
		"(define r 3)":                            "3",
		"(begin 5 (+ 3 4))":                       "7",
		"(begin (define p 3) (+ 39 p))":           "42",
		"(begin (define p 3) (set! p 4) (+ 1 p))": "5",
		"(begin (define p 3) (set! p (+ 1 1)) p)": "2",
		"(begin (define pi (+ 3 14)) pi)":         "17",
	}

	for in, out := range tests {
		tokens := Tokenize(in)
		parsed, err := Parse(tokens)
		if err != nil {
			t.Error(err)
		}
		x, err := Eval(parsed)
		if err != nil {
			t.Error(err)
		} else if fmt.Sprintf("%v", x) != out {
			t.Errorf("Parsing \"%v\" gives \"%v\", want \"%v\"", in, x, out)
		}
	}
}

func TestEvalFailures(t *testing.T) {
	var tests = []string{
		"hello",
		"(set! undefined 42)",
	}

	for _, in := range tests {
		tokens := Tokenize(in)
		parsed, err := Parse(tokens)
		if err != nil {
			t.Error(err)
		}
		x, err := Eval(parsed)
		if err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
