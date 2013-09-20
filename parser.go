package lisp

import "regexp"
import "fmt"
import "strconv"

type Parser struct {
	tokens []string
	values []interface{}
}

var Env map[string]interface{}

func init() {
	Env = make(map[string]interface{})
}

func NewParser(tokens []string) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (interface{}, error) {
	var pos int
	for pos < len(p.tokens) {
		t := p.tokens[pos]
		if m, _ := regexp.MatchString("^\\d+$", t); m { // Number
			i, err := strconv.Atoi(t)
			if err != nil {
				fmt.Errorf("Failed to convert number: %v", t)
			}
			p.values = append(p.values, i)
			pos++
		} else if t == "(" { // Open parenthesis
			start := pos + 1
			end, err := p.findEnd(start)
			if err != nil {
				return "", err
			}
			value, err := NewParser(p.tokens[start:end]).Parse()
			if err != nil {
				return "", err
			}
			p.values = append(p.values, value)
			pos = end + 1
		} else if t == ")" { // Close parenthesis
			return "", fmt.Errorf("List was closed but not opened")
			// Env
		} else if val, ok := Env[t]; ok {
			p.values = append(p.values, val)
			pos++
		} else { // Symbol
			p.values = append(p.values, t)
			pos++
		}
	}
	return p.Eval()
}

func (p *Parser) Eval() (interface{}, error) {
	t := p.values[0]
	if t == "quote" { // (quote exp)
		return p.values[1:], nil
	} else if t == "define" { // Define
		if len(p.values) == 3 {
			Env[p.values[1].(string)] = p.values[2]
			return p.values[2], nil
		} else {
			return nil, fmt.Errorf("Define require two parameters")
		}
	} else if _, ok := t.(int); ok { // Int
		return t, nil
	} else if _, ok := t.([]interface{}); ok { // Array
		return t, nil
	} else if t == "if" { // If
		if p.values[1] == "true" && len(p.values) > 2 {
			return p.values[2], nil
		} else if len(p.values) > 3 {
			return p.values[3], nil
		}
		return "nil", nil
	} else if t == "begin" { // Begin
		return p.values[len(p.values)-1], nil
	} else if t == "+" { // Addition
		var sum int
		for _, i := range p.values[1:] {
			v, ok := i.(int)
			if ok {
				sum += int(v)
			} else {
				return nil, fmt.Errorf("Cannot only add numbers: %v", i)
			}
		}
		return sum, nil
	}
	// Unknown
	return "", fmt.Errorf("Unknown symbol: %v", t)
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
