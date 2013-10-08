package lisp

import (
	"fmt"
	"regexp"
	"strconv"
)

type Tokens []*Token

type tokenType uint8

type Token struct {
	typ tokenType
	val string
}

type Pattern struct {
	typ    tokenType
	regexp *regexp.Regexp
}

func (t Token) String() string {
	return fmt.Sprintf("%v", t.val)
}

const (
	whitespaceToken tokenType = iota
	commentToken
	stringToken
	numberToken
	openToken
	closeToken
	symbolToken
)

func patterns() []Pattern {
	return []Pattern{
		{whitespaceToken, regexp.MustCompile(`^\s+`)},
		{commentToken, regexp.MustCompile(`^;.*`)},
		{stringToken, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{numberToken, regexp.MustCompile(`^((([0-9]+)?\.)?[0-9]+)`)},
		{openToken, regexp.MustCompile(`^(\()`)},
		{closeToken, regexp.MustCompile(`^(\))`)},
		{symbolToken, regexp.MustCompile(`^([^\s()]+)`)},
	}
}

func NewTokens(program string) (tokens Tokens) {
	for pos := 0; pos < len(program); {
		for _, pattern := range patterns() {
			if matches := pattern.regexp.FindStringSubmatch(program[pos:]); matches != nil {
				if len(matches) > 1 {
					tokens = append(tokens, &Token{pattern.typ, matches[1]})
				}
				pos = pos + len(matches[0])
				break
			}
		}
	}
	return
}

func (tokens Tokens) Parse() (Sexp, error) {
	var pos int
	values := make(Sexp, 0)
	for pos < len(tokens) {
		t := tokens[pos]
		switch t.typ {
		case numberToken:
			if i, err := strconv.ParseFloat(t.val, 64); err != nil {
				return nil, fmt.Errorf("Failed to convert number: %v", t.val)
			} else {
				values = append(values, Value{numberValue, i})
				pos++
			}
		case stringToken:
			values = append(values, Value{stringValue, t.val[1 : len(t.val)-1]})
			pos++
		case symbolToken:
			values = append(values, Value{symbolValue, t.val})
			pos++
		case openToken:
			start := pos + 1
			end, err := tokens.findClose(start)
			if err != nil {
				return nil, err
			}
			x, err := tokens[start:end].Parse()
			if err != nil {
				return nil, err
			}
			values = append(values, Value{sexpValue, x})
			pos = end + 1
		case closeToken:
			return nil, fmt.Errorf("List was closed but not opened")
		}
	}
	return values, nil
}

func (t Tokens) findClose(start int) (int, error) {
	depth := 1
	for i := start; i < len(t); i++ {
		t := t[i]
		switch t.typ {
		case openToken:
			depth++
		case closeToken:
			depth--
		}
		if depth == 0 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("List was opened but not closed")
}
