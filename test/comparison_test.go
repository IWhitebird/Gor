package test

import "testing"

func TestEqualNumbers(t *testing.T) {
	expectOutput(t, "print(1 == 1)", "true")
}

func TestNotEqualNumbers(t *testing.T) {
	expectOutput(t, "print(1 != 2)", "true")
}

func TestEqualNumbersFalse(t *testing.T) {
	expectOutput(t, "print(1 == 2)", "false")
}

func TestGreaterThan(t *testing.T) {
	expectOutput(t, "print(5 > 3)", "true")
	expectOutput(t, "print(3 > 5)", "false")
}

func TestLessThan(t *testing.T) {
	expectOutput(t, "print(3 < 5)", "true")
	expectOutput(t, "print(5 < 3)", "false")
}

func TestGreaterOrEqual(t *testing.T) {
	expectOutput(t, "print(5 >= 5)", "true")
	expectOutput(t, "print(6 >= 5)", "true")
	expectOutput(t, "print(4 >= 5)", "false")
}

func TestLessOrEqual(t *testing.T) {
	expectOutput(t, "print(5 <= 5)", "true")
	expectOutput(t, "print(4 <= 5)", "true")
	expectOutput(t, "print(6 <= 5)", "false")
}

func TestEqualBooleans(t *testing.T) {
	expectOutput(t, "print(true == true)", "true")
	expectOutput(t, "print(false == false)", "true")
	expectOutput(t, "print(true == false)", "false")
}

func TestNotEqualBooleans(t *testing.T) {
	expectOutput(t, "print(true != false)", "true")
	expectOutput(t, "print(true != true)", "false")
}

func TestEqualStrings(t *testing.T) {
	expectOutput(t, `print("abc" == "abc")`, "true")
	expectOutput(t, `print("abc" == "xyz")`, "false")
}

func TestCrossTypeEquality(t *testing.T) {
	// Different types should not be equal
	expectOutput(t, `print(1 == true)`, "false")
	expectOutput(t, `print(0 == null)`, "false")
	expectOutput(t, `print("1" == 1)`, "false")
}

func TestNullEquality(t *testing.T) {
	expectOutput(t, "print(null == null)", "true")
	expectOutput(t, "print(null != null)", "false")
}

func TestComparisonInVariable(t *testing.T) {
	expectOutput(t, `
let result = 10 > 5
print(result)
`, "true")
}

func TestChainedComparisons(t *testing.T) {
	// a == b == true is parsed as (a == b) == true
	expectOutput(t, "print((5 == 5) == true)", "true")
}
