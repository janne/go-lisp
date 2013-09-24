package lisp

import "testing"

func TestScope(t *testing.T) {
	scope := NewScope()
	if scope.Env() != nil {
		t.Errorf("Env should be nil initially")
	}

	env := scope.AddEnv()
	if env != scope.Env() {
		t.Errorf("AddEnv() returns %v, should be same as scope.Env(): %v", env, scope.Env())
	}

	env2 := scope.AddEnv()
	if env2 != scope.Env() {
		t.Errorf("AddEnv() returns %v, should be same as scope.Env(): %v", env2, scope.Env())
	}

	env3 := scope.DropEnv()
	if env3 != scope.Env() {
		t.Errorf("DropEnv() returns %v, should be same as scope.Env(): %v", env3, scope.Env())
	}

	if env != env3 {
		t.Errorf("Original env: %v should be same as dropped env from DropEnv(): %v", env, env3)
	}

	env4 := scope.DropEnv()
	if env4 != nil {
		t.Errorf("DropEnv should be back to nil but is %v", env4)
	}
}
