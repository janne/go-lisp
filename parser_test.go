package lisp

import "testing"

func TestParse(t *testing.T) {
	var tests = map[string]string{
		"42":      "42",
		"(42)":    "42",
		"((42))":  "42",
		"(42 13)": "42",
	}

	for in, out := range tests {
		tokens := Tokenize(in)
		x, _ := NewParser(tokens).Parse()
		if x != out {
			t.Errorf("Parse('%v') = '%v', want '%v'", in, x, out)
		}
	}
}

func TestParseFailures(t *testing.T) {
	var tests = []string{
		"(42",
	}

	for _, in := range tests {
		tokens := Tokenize(in)
		x, err := NewParser(tokens).Parse()
		if err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
