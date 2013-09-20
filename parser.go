package lisp

import "regexp"
import "errors"

type Parser struct {
	tokens []string
}

func NewParser(tokens []string) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) findEnd(start int) (int, error) {
	depth := 1

	for i := start; i < len(p.tokens); i++ {
		t := p.tokens[i]
		if t == "(" {
			depth++
		} else if t == ")" {
			depth--
		}
		if depth == 0 {
			return i, nil
		}
	}
	return 0, errors.New("List was opened but not closed")
}

func (p *Parser) Parse() (string, error) {
	for i := 0; i < len(p.tokens); i++ {
		t := p.tokens[i]
		if m, _ := regexp.MatchString("^\\d+$", t); m {
			return t, nil
		} else if t == "(" {
			start := i + 1
			end, err := p.findEnd(start)
			if err != nil {
				return "", err
			}
			return NewParser(p.tokens[start:end]).Parse()
		} else if t == ")" {
			return "", errors.New("List was closed but not opened")
		}
	}
	return "", errors.New("Invalid input")
}
