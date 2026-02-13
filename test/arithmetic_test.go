package test

import "testing"

func TestBasicAddition(t *testing.T) {
	expectOutput(t, "print(1 + 2)", "3")
}

func TestBasicSubtraction(t *testing.T) {
	expectOutput(t, "print(10 - 3)", "7")
}

func TestBasicMultiplication(t *testing.T) {
	expectOutput(t, "print(4 * 5)", "20")
}

func TestBasicDivision(t *testing.T) {
	expectOutput(t, "print(20 / 4)", "5")
}

func TestModulo(t *testing.T) {
	expectOutput(t, "print(10 % 3)", "1")
}

func TestIntegerDivisionTruncates(t *testing.T) {
	expectOutput(t, "print(7 / 2)", "3")
}

func TestOperatorPrecedence(t *testing.T) {
	expectOutput(t, "print(2 + 3 * 4)", "14")
}

func TestOperatorPrecedenceSubMul(t *testing.T) {
	expectOutput(t, "print(10 - 2 * 3)", "4")
}

func TestParenthesesOverridePrecedence(t *testing.T) {
	expectOutput(t, "print((2 + 3) * 4)", "20")
}

func TestNestedParentheses(t *testing.T) {
	expectOutput(t, "print(((1 + 2) * (3 + 4)))", "21")
}

func TestChainedAddition(t *testing.T) {
	expectOutput(t, "print(1 + 2 + 3 + 4 + 5)", "15")
}

func TestChainedMultiplication(t *testing.T) {
	expectOutput(t, "print(2 * 3 * 4)", "24")
}

func TestMixedArithmetic(t *testing.T) {
	expectOutput(t, "print(100 / 10 + 5 * 2 - 3)", "17")
}

func TestNegativeResult(t *testing.T) {
	expectOutput(t, "print(3 - 10)", "-7")
}

func TestZeroOperations(t *testing.T) {
	expectOutput(t, "print(0 + 0)", "0")
	expectOutput(t, "print(0 * 100)", "0")
	expectOutput(t, "print(0 / 5)", "0")
}

func TestBitwiseAnd(t *testing.T) {
	expectOutput(t, "print(6 & 3)", "2")
}

func TestBitwiseOr(t *testing.T) {
	expectOutput(t, "print(6 | 3)", "7")
}

func TestModuloEdgeCases(t *testing.T) {
	expectOutput(t, "print(10 % 10)", "0")
	expectOutput(t, "print(1 % 5)", "1")
}

func TestLargeNumbers(t *testing.T) {
	expectOutput(t, "print(1000000 * 1000)", "1000000000")
}
