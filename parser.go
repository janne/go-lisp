package lisp

type Parser struct {
	tokens []string
}

func NewParser(program string) *Parser {
	p := &Parser{}
	p.tokens = Tokenize(program)
	return p
}

func (p *Parser) Parse() string {
	return "foo"
}
