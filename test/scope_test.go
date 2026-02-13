package test

import "testing"

func TestBlockCreatesNewScope(t *testing.T) {
	expectOutput(t, `
let x = 1
if (true) {
	let x = 2
	print(x)
}
print(x)
`, "2\n1")
}

func TestForLoopScope(t *testing.T) {
	// Each iteration gets its own scope
	expectOutput(t, `
let total = 0
for (let i = 0; i < 3; i = i + 1) {
	let x = i * 10
	total = total + x
}
print(total)
`, "30")
}

func TestFunctionHasOwnScope(t *testing.T) {
	expectOutput(t, `
let x = 100
fn change() {
	let x = 999
	return x
}
print(change())
print(x)
`, "999\n100")
}

func TestFunctionCanReadOuterScope(t *testing.T) {
	expectOutput(t, `
let x = 42
fn read() {
	return x
}
print(read())
`, "42")
}

func TestFunctionCanModifyOuterScope(t *testing.T) {
	expectOutput(t, `
let x = 1
fn modify() {
	x = 99
	return x
}
modify()
print(x)
`, "99")
}

func TestNestedScopeChain(t *testing.T) {
	expectOutput(t, `
let a = 1
if (true) {
	let b = 2
	if (true) {
		let c = 3
		print(a + b + c)
	}
}
`, "6")
}

func TestShadowedVariableDoesNotAffectOuter(t *testing.T) {
	expectOutput(t, `
let x = "outer"
fn test() {
	let x = "inner"
	return x
}
print(test())
print(x)
`, "inner\nouter")
}

func TestForLoopVarNotVisibleOutside(t *testing.T) {
	// After loop, outer variable should be unchanged
	expectOutput(t, `
let result = 0
for (let i = 0; i < 5; i = i + 1) {
	result = result + i
}
print(result)
`, "10")
}
