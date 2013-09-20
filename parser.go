package lisp

import "regexp"
import "fmt"
import "strconv"

type Parser struct {
	tokens []string
	values []string
}

func NewParser(tokens []string) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (string, error) {
	var i int
	for i < len(p.tokens) {
		t := p.tokens[i]
		// BEGIN PARENTHESIS
		if t == "(" {
			start := i + 1
			end, err := p.findEnd(start)
			if err != nil {
				return "", err
			}
			value, err := NewParser(p.tokens[start:end]).Parse()
			if err != nil {
				return "", err
			}
			p.values = append(p.values, value)
			i = end + 1
			// END PARENTHESIS
		} else if t == ")" {
			return "", fmt.Errorf("%v: List was closed but not opened", i)
		} else {
			i++
			p.values = append(p.values, t)
		}
	}
	return p.Eval()
}

func (p *Parser) Eval() (string, error) {
	t := p.values[0]
	if m, _ := regexp.MatchString("^\\d+$", t); m {
		return t, nil
	} else if t == "+" {
		var sum int
		for _, val := range p.values[1:] {
			i, err := strconv.Atoi(val)
			if err != nil {
				fmt.Errorf("Can only add numbers: %v", val)
			}
			sum += i
		}
		return strconv.Itoa(sum), nil
	} else {
		return "", fmt.Errorf("Unknown symbol: %v", t)
	}
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
	return 0, fmt.Errorf("List was opened but not closed")
}
