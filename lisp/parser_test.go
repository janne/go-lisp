package lisp

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	var tests = map[string]string{
		"42":            "42",
		"(+ (+ 1 2) 3)": "[+ [+ 1 2] 3]",
	}
	for in, out := range tests {
		tokens := Tokenize(in)
		parsed, _ := Parse(tokens)
		result := fmt.Sprintf("%v", parsed)
		if result != out {
			t.Errorf("Parse \"%v\" gives \"%v\", expected \"%v\"", in, result, out)
		}
	}
}

func TestParseFailures(t *testing.T) {
	var tests = []string{
		"(42",
	}

	for _, in := range tests {
		tokens := Tokenize(in)
		x, err := Parse(tokens)
		if err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
