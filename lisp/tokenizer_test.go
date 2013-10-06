package lisp

import "testing"

func equalSlices(a, b []*Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.val != b[i].val || v.typ != b[i].typ {
			return false
		}
	}
	return true
}

func TestTokenize(t *testing.T) {
	var tests = []struct {
		in  string
		out []*Token
	}{
		{"(define a 42)", []*Token{{openType, "("}, {symbolType, "define"}, {symbolType, "a"}, {numberType, "42"}, {closeType, ")"}}},
		{"\t(quote\n\t\t(a b c))  ", []*Token{{openType, "("}, {symbolType, "quote"}, {openType, "("}, {symbolType, "a"}, {symbolType, "b"}, {symbolType, "c"}, {closeType, ")"}, {closeType, ")"}}},
		{"hello ; dude\n\tworld", []*Token{{symbolType, "hello"}, {symbolType, "world"}}},
		{"test \"a string\"", []*Token{{symbolType, "test"}, {stringType, "\"a string\""}}},
		{"\"only string\"", []*Token{{stringType, "\"only string\""}}},
		{"\"string\\nwith\\\"escape\\tcharacters\"", []*Token{{stringType, "\"string\\nwith\\\"escape\\tcharacters\""}}},
		{"\"hej\\\"hello\"", []*Token{{stringType, "\"hej\\\"hello\""}}},
	}

	for _, test := range tests {
		x := Tokenize(test.in)
		if !equalSlices(x, test.out) {
			t.Errorf("Tokenize \"%v\" gives \"%v\", expected \"%v\"", test.in, x, test.out)
		}
	}
}
