package lisp

import (
	"fmt"
	"regexp"
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
