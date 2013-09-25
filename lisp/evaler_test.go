package lisp

import "testing"
import "fmt"

func TestEval(t *testing.T) {
	var tests = map[string]interface{}{
		"()":                                                                                             "<nil>",
		"42":                                                                                             "42",
		"1 2 3":                                                                                          "3",
		"(+ 42 13)":                                                                                      "55",
		"(+ (+ 1 2 3) 4)":                                                                                "10",
		"(quote (1 2 3))":                                                                                "[1 2 3]",
		"(quote (1 (+ 1 2) 3))":                                                                          "[1 [+ 1 2] 3]",
		"(quote hej)":                                                                                    "hej",
		"(if true (+ 1 1) 3)":                                                                            "2",
		"(if false 42 1)":                                                                                "1",
		"(if false 42)":                                                                                  "<nil>",
		"(begin (define x) (if x 1 2))":                                                                  "2",
		"(define r 3)":                                                                                   "r",
		"(begin 5 (+ 3 4))":                                                                              "7",
		"(begin (define p 3) (+ 39 p))":                                                                  "42",
		"(begin (define p 3) (set! p 4) (+ 1 p))":                                                        "5",
		"(begin (define p 3) (set! p (+ 1 1)) p)":                                                        "2",
		"(begin (define pi (+ 3 14)) pi)":                                                                "17",
		"((lambda (a) (+ a 1)) 42)":                                                                      "43",
		"(begin (define p 10) p)":                                                                        "10",
		"(begin (define inc (lambda (a) (+ a 1))) (inc 42))":                                             "43",
		"(define a 10) ((lambda () (define a 20))) a":                                                    "10",
		"(define a 0) ((lambda () (set! a 10))) a":                                                       "10",
		"((lambda (i) i) (+ 5 5))":                                                                       "10",
		"(define inc ((lambda () (begin (define a 0) (lambda () (set! a (+ a 1))))))) (inc) (inc) (inc)": "3",
		"(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1)))))) (fact 20)":                      "2432902008176640000",
	}

	for in, out := range tests {
		x, err := EvalString(in)
		if err != nil {
			t.Error(err)
		} else if fmt.Sprintf("%v", x) != out {
			t.Errorf("Eval \"%v\" gives \"%v\", want \"%v\"", in, x, out)
		}
	}
}

func TestEvalFailures(t *testing.T) {
	var tests = []string{
		"hello",
		"(set! undefined 42)",
		"(lambda (a))",
		"(1 2 3)",
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
