package lisp

import "regexp"

func Tokenize(program string) (result []string) {
	program = regexp.MustCompile(";.*").ReplaceAllString(program, "")
	return regexp.MustCompile(`("[^"]*"|\(|\)|[^\s()]+)`).FindAllString(program, -1)
}
