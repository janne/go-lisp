package lisp

import "regexp"
import "fmt"
import "strconv"

type Value interface{}

type Sexp []Value

func (s Sexp) String() string {
	if len(s) == 1 {
		return fmt.Sprintf("%v", s[0])
	} else {
		return fmt.Sprintf("%v", []Value(s))
	}
}

func Parse(tokens []string) (Sexp, error) {
	var pos int
	values := make(Sexp, 0)
	for pos < len(tokens) {
		t := tokens[pos]
		if m, _ := regexp.MatchString("^\\d+$", t); m { // Number
			i, err := strconv.Atoi(t)
			if err != nil {
				fmt.Errorf("Failed to convert number: %v", t)
			}
			values = append(values, i)
			pos++
		} else if t == "(" { // Open parenthesis
			start := pos + 1
			end, err := findEnd(tokens, start)
			if err != nil {
				return nil, err
			}
			x, err := Parse(tokens[start:end])
			if err != nil {
				return nil, err
			}
			values = append(values, x)
			pos = end + 1
		} else if t == ")" { // Close parenthesis
			return nil, fmt.Errorf("List was closed but not opened")
		} else { // Symbol
			values = append(values, t)
			pos++
		}
	}
	return values, nil
}

func findEnd(tokens []string, start int) (int, error) {
	depth := 1

	for i := start; i < len(tokens); i++ {
		t := tokens[i]
		if t == "(" {
			depth++
		} else if t == ")" {
			depth--
		}
		if depth == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("List was opened but not closed")
}
