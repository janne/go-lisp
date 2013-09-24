package lisp

import "testing"

func TestAdd(t *testing.T) {
	if builtins["+"] != "add" {
		t.Errorf("+ is not correcly setup")
	}

	if sum, err := add(1, 2); sum != 3 || err != nil {
		t.Errorf("1 + 2 should == 3, is %v, error: %v", sum, err)
	}
}
