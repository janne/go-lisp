package lisp

import "testing"

func TestAdd(t *testing.T) {
	if !isBuiltin(NewValue("+")) {
		t.Errorf("+ is not correcly setup")
	}

	if sum, err := builtin.Add(NewValue(1), NewValue(2), NewValue(3)); sum != NewValue(6) || err != nil {
		t.Errorf("1 + 2 + 3 should == 6, is %v, error: %v", sum, err)
	}
}

func TestSub(t *testing.T) {
	if sum, err := builtin.Sub(NewValue(5), NewValue(2), NewValue(1)); sum != NewValue(2) || err != nil {
		t.Errorf("5 - 2 - 1 should == 2, is %v, error: %v", sum, err)
	}
}

func TestMul(t *testing.T) {
	if sum, err := builtin.Mul(NewValue(2), NewValue(3), NewValue(4)); sum != NewValue(24) || err != nil {
		t.Errorf("2 * 3 * 4 should == 24, is %v, error: %v", sum, err)
	}
}

func TestGt(t *testing.T) {
	if result, err := builtin.Gt(NewValue(4), NewValue(3), NewValue(2)); result == NewValue(false) || err != nil {
		t.Errorf("4 > 3 > 2 should == true, is %v, error: %v", result, err)
	}

	if result, err := builtin.Gt(NewValue(4), NewValue(4), NewValue(2)); result == NewValue(true) || err != nil {
		t.Errorf("4 > 4 > 2 should == true, is %v, error: %v", result, err)
	}
}

func TestLt(t *testing.T) {
	if result, err := builtin.Lt(NewValue(2), NewValue(3), NewValue(4)); result == NewValue(false) || err != nil {
		t.Errorf("2 < 3 < 4 should == true, is %v, error: %v", result, err)
	}
}

func TestGte(t *testing.T) {
	if result, err := builtin.Gte(NewValue(4), NewValue(4), NewValue(2)); result == NewValue(false) || err != nil {
		t.Errorf("4 >= 4 >= 2 should == true, is %v, error: %v", result, err)
	}
}

func TestLte(t *testing.T) {
	if result, err := builtin.Lte(NewValue(2), NewValue(2), NewValue(4)); result == NewValue(false) || err != nil {
		t.Errorf("2 <= 2 <= 4 should == true, is %v, error: %v", result, err)
	}
}
