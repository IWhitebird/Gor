package test

import "testing"

func TestPrintNumber(t *testing.T) {
	expectOutput(t, "print(42)", "42")
}

func TestPrintString(t *testing.T) {
	expectOutput(t, `print("hello")`, "hello")
}

func TestPrintBoolTrue(t *testing.T) {
	expectOutput(t, "print(true)", "true")
}

func TestPrintBoolFalse(t *testing.T) {
	expectOutput(t, "print(false)", "false")
}

func TestPrintNull(t *testing.T) {
	expectOutput(t, "print(null)", "null")
}

func TestPrintExpression(t *testing.T) {
	expectOutput(t, "print(1 + 2 + 3)", "6")
}

func TestPrintMultipleCalls(t *testing.T) {
	expectLines(t, `
print(1)
print(2)
print(3)
`, "1", "2", "3")
}

func TestPrintZero(t *testing.T) {
	expectOutput(t, "print(0)", "0")
}

func TestPrintNegative(t *testing.T) {
	expectOutput(t, "print(0 - 5)", "-5")
}

func TestPrintEmptyString(t *testing.T) {
	expectOutput(t, `print("")`, "")
}

func TestPrintVariable(t *testing.T) {
	expectOutput(t, `
let x = 123
print(x)
`, "123")
}

func TestPrintFunctionResult(t *testing.T) {
	expectOutput(t, `
fn double(x) {
	return x * 2
}
print(double(21))
`, "42")
}
