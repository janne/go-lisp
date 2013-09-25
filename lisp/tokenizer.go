package lisp

import (
	"regexp"
	"strings"
)

func Tokenize(program string) []string {
	program = strings.Replace(program, "(", " ( ", -1)
	program = strings.Replace(program, ")", " ) ", -1)
	program = regexp.MustCompile(";.*").ReplaceAllString(program, "")
	program = strings.TrimSpace(program)
	t := regexp.MustCompile("\\s+").Split(program, -1)
	return t
}
