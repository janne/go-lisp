package lisp

import "testing"

func TestParse(t *testing.T) {
	var tests = map[string]string{
		"42": "42",
	}

	for in, out := range tests {
		x := NewParser(in).Parse()
		if x != out {
			t.Errorf("Parse(%v) = %v, want %v", in, x, out)
		}
	}
}
