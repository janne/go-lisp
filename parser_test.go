package lisp

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	var tests = map[string]interface{}{
		"42":                            "42",
		"(42)":                          "42",
		"((42))":                        "42",
		"(42 13)":                       "42",
		"(+ 42 13)":                     "55",
		"(+ (+ 1 2 3) 4)":               "10",
		"(quote 1 2 3)":                 "[1 2 3]",
		"if true (+ 1 1) 3":             "2",
		"if false 42 1":                 "1",
		"if false 42":                   "nil",
		"(define r 3)":                  "3",
		"(begin 5 (+ 3 4))":             "7",
		"(begin (define p 3) (+ 39 p))": "42",
	}

	for in, out := range tests {
		tokens := Tokenize(in)
		x, err := NewParser(tokens).Parse()
		if err != nil {
			t.Error(err)
		} else if fmt.Sprintf("%v", x) != out {
			t.Errorf("Parsing \"%v\" gives \"%v\", want \"%v\"", in, x, out)
		}
	}
}

func TestParseFailures(t *testing.T) {
	var tests = []string{
		"(42",
		"hello",
	}

	for _, in := range tests {
		tokens := Tokenize(in)
		x, err := NewParser(tokens).Parse()
		if err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
