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
	var tests = []struct {
		in  string
		out []string
	}{
		{"(define a 42)", []string{"(", "define", "a", "42", ")"}},
		{"\t(quote\n\t\t(a b c))  ", []string{"(", "quote", "(", "a", "b", "c", ")", ")"}},
		{"hello ; dude\n\tworld", []string{"hello", "world"}},
	}

	for _, test := range tests {
		x := Tokenize(test.in)
		if !equalSlices(x, test.out) {
			t.Errorf("Tokenize \"%v\" gives \"%v\", expected \"%v\"", test.in, x, test.out)
		}
	}
}
