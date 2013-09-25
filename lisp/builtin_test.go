package lisp

import "testing"

func TestAdd(t *testing.T) {
	if !isBuiltin("+") {
		t.Errorf("+ is not correcly setup")
	}

	if sum, err := builtin.Add(1, 2, 3); sum != 6 || err != nil {
		t.Errorf("1 + 2 + 3 should == 6, is %v, error: %v", sum, err)
	}
}

func TestSub(t *testing.T) {
	if sum, err := builtin.Sub(5, 2, 1); sum != 2 || err != nil {
		t.Errorf("5 - 2 - 1 should == 2, is %v, error: %v", sum, err)
	}
}

func TestMul(t *testing.T) {
	if sum, err := builtin.Mul(2, 3, 4); sum != 24 || err != nil {
		t.Errorf("2 * 3 * 4 should == 24, is %v, error: %v", sum, err)
	}
}

func TestGt(t *testing.T) {
	if result, err := builtin.Gt(4, 3, 2); result == "false" || err != nil {
		t.Errorf("4 > 3 > 2 should == true, is %v, error: %v", result, err)
	}

	if result, err := builtin.Gt(4, 4, 2); result == "true" || err != nil {
		t.Errorf("4 > 4 > 2 should == true, is %v, error: %v", result, err)
	}
}

func TestLt(t *testing.T) {
	if result, err := builtin.Lt(2, 3, 4); result == "false" || err != nil {
		t.Errorf("2 < 3 < 4 should == true, is %v, error: %v", result, err)
	}
}

func TestGte(t *testing.T) {
	if result, err := builtin.Gte(4, 4, 2); result == "false" || err != nil {
		t.Errorf("4 >= 4 >= 2 should == true, is %v, error: %v", result, err)
	}
}

func TestLte(t *testing.T) {
	if result, err := builtin.Lte(2, 2, 4); result == "false" || err != nil {
		t.Errorf("2 <= 2 <= 4 should == true, is %v, error: %v", result, err)
	}
}
