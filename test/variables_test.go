package test

import "testing"

func TestLetDeclaration(t *testing.T) {
	expectOutput(t, "let x = 10\nprint(x)", "10")
}

func TestConstDeclaration(t *testing.T) {
	expectOutput(t, "const x = 42\nprint(x)", "42")
}

func TestLetReassignment(t *testing.T) {
	expectOutput(t, `
let x = 1
x = 2
print(x)
`, "2")
}

func TestLetWithoutValue(t *testing.T) {
	// let without value should not crash
	expectOutput(t, `
let x
print(null)
`, "null")
}

func TestMultipleVariables(t *testing.T) {
	expectOutput(t, `
let a = 1
let b = 2
let c = 3
print(a + b + c)
`, "6")
}

func TestVariableUsedInExpression(t *testing.T) {
	expectOutput(t, `
let x = 10
let y = x * 2
print(y)
`, "20")
}

func TestVariableShadowingInBlock(t *testing.T) {
	expectOutput(t, `
let x = 1
if (true) {
	let x = 99
	print(x)
}
print(x)
`, "99\n1")
}

func TestVariableFromOuterScope(t *testing.T) {
	expectOutput(t, `
let x = 42
if (true) {
	print(x)
}
`, "42")
}

func TestReassignInNestedScope(t *testing.T) {
	expectOutput(t, `
let x = 1
if (true) {
	x = 100
}
print(x)
`, "100")
}

func TestStringVariable(t *testing.T) {
	expectOutput(t, `let s = "hello"
print(s)`, "hello")
}

func TestBoolVariable(t *testing.T) {
	expectOutput(t, "let b = true\nprint(b)", "true")
}

func TestNullVariable(t *testing.T) {
	expectOutput(t, "print(null)", "null")
}
