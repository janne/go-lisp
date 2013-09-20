package lisp

import "strings"
import "regexp"

func Tokenize(program string) []string {
	r := strings.NewReplacer("(", " ( ", ")", " ) ")
	program = r.Replace(program)
	program = strings.TrimSpace(program)
	t := regexp.MustCompile("\\s+").Split(program, -1)
	return t
}
