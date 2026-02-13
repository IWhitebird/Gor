package test

import "testing"

func TestLogicalAndTrue(t *testing.T) {
	expectOutput(t, "print(true && true)", "true")
}

func TestLogicalAndFalse(t *testing.T) {
	expectOutput(t, "print(true && false)", "false")
}

func TestLogicalOrTrue(t *testing.T) {
	expectOutput(t, "print(false || true)", "true")
}

func TestLogicalOrFalse(t *testing.T) {
	expectOutput(t, "print(false || false)", "false")
}

func TestChainedAnd(t *testing.T) {
	expectOutput(t, "print(true && true && true)", "true")
	expectOutput(t, "print(true && false && true)", "false")
}

func TestChainedOr(t *testing.T) {
	expectOutput(t, "print(false || false || true)", "true")
	expectOutput(t, "print(false || false || false)", "false")
}

func TestAndShortCircuit(t *testing.T) {
	// && with falsy left returns the left value
	expectOutput(t, `
if (false && true) {
	print("bad")
} else {
	print("good")
}
`, "good")
}

func TestOrShortCircuit(t *testing.T) {
	// || with truthy left returns the left value
	expectOutput(t, `
if (true || false) {
	print("good")
} else {
	print("bad")
}
`, "good")
}

func TestLogicalWithNumbers(t *testing.T) {
	// Truthy: non-zero numbers
	expectOutput(t, `
if (1 && 2) {
	print("truthy")
}
`, "truthy")
}

func TestLogicalAndWithZero(t *testing.T) {
	// 0 is falsy
	expectOutput(t, `
if (0 && 1) {
	print("bad")
} else {
	print("falsy")
}
`, "falsy")
}

func TestLogicalOrWithZero(t *testing.T) {
	expectOutput(t, `
if (0 || 1) {
	print("truthy")
}
`, "truthy")
}

func TestLogicalWithStrings(t *testing.T) {
	// Non-empty string is truthy, empty is falsy
	expectOutput(t, `
if ("hello" && "world") {
	print("truthy")
}
`, "truthy")
}

func TestBitwiseAndBool(t *testing.T) {
	expectOutput(t, "print(true & true)", "true")
	expectOutput(t, "print(true & false)", "false")
}

func TestBitwiseOrBool(t *testing.T) {
	expectOutput(t, "print(false | true)", "true")
	expectOutput(t, "print(false | false)", "false")
}
