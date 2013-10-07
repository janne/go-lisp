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
		{"(define a 42)", []*Token{{openToken, "("}, {symbolToken, "define"}, {symbolToken, "a"}, {numberToken, "42"}, {closeToken, ")"}}},
		{"\t(quote\n\t\t(a b c))  ", []*Token{{openToken, "("}, {symbolToken, "quote"}, {openToken, "("}, {symbolToken, "a"}, {symbolToken, "b"}, {symbolToken, "c"}, {closeToken, ")"}, {closeToken, ")"}}},
		{"hello ; dude\n\tworld", []*Token{{symbolToken, "hello"}, {symbolToken, "world"}}},
		{"test \"a string\"", []*Token{{symbolToken, "test"}, {stringToken, "\"a string\""}}},
		{"\"only string\"", []*Token{{stringToken, "\"only string\""}}},
		{"\"string\\nwith\\\"escape\\tcharacters\"", []*Token{{stringToken, "\"string\\nwith\\\"escape\\tcharacters\""}}},
		{"\"hej\\\"hello\"", []*Token{{stringToken, "\"hej\\\"hello\""}}},
	}

	for _, test := range tests {
		x := Tokenize(test.in)
		if !equalSlices(x, test.out) {
			t.Errorf("Tokenize \"%v\" gives \"%v\", expected \"%v\"", test.in, x, test.out)
		}
	}
}
