package lisp

import (
	"regexp"
)

type Pattern struct {
	typ    tokenType
	regexp *regexp.Regexp
}

type tokenType int

type Token struct {
	typ tokenType
	val string
}

const (
	whitespaceType tokenType = iota
	commentType
	stringType
	numberType
	openType
	closeType
	symbolType
)

func patterns() []Pattern {
	return []Pattern{
		{whitespaceType, regexp.MustCompile(`^\s+`)},
		{commentType, regexp.MustCompile(`^;.*`)},
		{stringType, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{numberType, regexp.MustCompile(`^((([0-9]+)?\.)?[0-9]+)`)},
		{openType, regexp.MustCompile(`^(\()`)},
		{closeType, regexp.MustCompile(`^(\))`)},
		{symbolType, regexp.MustCompile(`^([^\s()]+)`)},
	}
}

func Tokenize(program string) (result []*Token) {
	for pos := 0; pos < len(program); {
		for _, pattern := range patterns() {
			if matches := pattern.regexp.FindStringSubmatch(program[pos:]); matches != nil {
				if len(matches) > 1 {
					result = append(result, &Token{pattern.typ, matches[1]})
				}
				pos = pos + len(matches[0])
				break
			}
		}
	}
	return result
}
