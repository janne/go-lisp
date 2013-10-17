package lisp

func EvalString(line string) (Value, error) {
	parsed, err := NewTokens(line).Parse()
	if err != nil {
		return Nil, err
	}
	evaled, err := parsed.Eval()
	if err != nil {
		return Nil, err
	}
	scope.Create("_", evaled)
	return evaled, nil
}
