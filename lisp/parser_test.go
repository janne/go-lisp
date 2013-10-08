package lisp

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"42", "42"},
		{"(+ (+ 1 2) 3)", "(+ (+ 1 2) 3)"},
	}
	for _, test := range tests {
		tokens := NewTokens(test.in)
		parsed, _ := Parse(tokens)
		result := fmt.Sprintf("%v", parsed)
		if result != test.out {
			t.Errorf("Parse \"%v\" gives \"%v\", expected \"%v\"", test.in, result, test.out)
		}
	}
}

func TestParseFailures(t *testing.T) {
	var tests = []string{
		"(42",
	}

	for _, in := range tests {
		tokens := NewTokens(in)
		x, err := Parse(tokens)
		if err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
