package lisp

import "testing"

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	var tests = map[string][]string{
		"(define a 42)":            {"(", "define", "a", "42", ")"},
		"\t(quote\n\t\t(a b c))  ": {"(", "quote", "(", "a", "b", "c", ")", ")"},
		"hello ; dude\n\tworld":    {"hello", "world"},
	}

	for in, out := range tests {
		x := Tokenize(in)
		if !equalSlices(x, out) {
			t.Errorf("Tokenize \"%v\" gives \"%v\", expected \"%v\"", in, x, out)
		}
	}
}
