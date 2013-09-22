package lisp

var Env map[string]interface{}

func init() {
	Env = make(map[string]interface{})
}
