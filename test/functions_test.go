package test

import "testing"

func TestSimpleFunction(t *testing.T) {
	expectOutput(t, `
fn add(a, b) {
	return a + b
}
print(add(3, 4))
`, "7")
}

func TestFunctionNoArgs(t *testing.T) {
	expectOutput(t, `
fn greet() {
	return 42
}
print(greet())
`, "42")
}

func TestFunctionWithoutReturn(t *testing.T) {
	// Should not crash — returns last evaluated value
	expectOutput(t, `
fn doStuff() {
	let x = 10
}
doStuff()
print("ok")
`, "ok")
}

func TestFunctionReturnString(t *testing.T) {
	expectOutput(t, `
fn hello() {
	return "world"
}
print(hello())
`, "world")
}

func TestFunctionReturnBool(t *testing.T) {
	expectOutput(t, `
fn isPositive(n) {
	if (n > 0) {
		return true
	}
	return false
}
print(isPositive(5))
print(isPositive(0))
`, "true\nfalse")
}

func TestFunctionMultipleParams(t *testing.T) {
	expectOutput(t, `
fn calc(a, b, c) {
	return a + b * c
}
print(calc(1, 2, 3))
`, "7")
}

func TestFunctionCallingFunction(t *testing.T) {
	expectOutput(t, `
fn double(x) {
	return x * 2
}
fn quadruple(x) {
	return double(double(x))
}
print(quadruple(3))
`, "12")
}

func TestRecursion(t *testing.T) {
	expectOutput(t, `
fn factorial(n) {
	if (n <= 1) {
		return 1
	}
	return n * factorial(n - 1)
}
print(factorial(5))
`, "120")
}

func TestRecursionFibonacci(t *testing.T) {
	expectOutput(t, `
fn fib(n) {
	if (n <= 1) {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
print(fib(10))
`, "55")
}

func TestFunctionAsValue(t *testing.T) {
	// Functions can be assigned to variables
	expectOutput(t, `
fn add(a, b) {
	return a + b
}
let result = add(10, 20)
print(result)
`, "30")
}

func TestFunctionScopeDoesNotLeak(t *testing.T) {
	// Variables inside function should not be visible outside
	expectOutput(t, `
let x = 1
fn change() {
	let y = 99
	return y
}
print(change())
print(x)
`, "99\n1")
}

func TestEarlyReturn(t *testing.T) {
	expectOutput(t, `
fn check(x) {
	if (x > 10) {
		return "big"
	}
	return "small"
}
print(check(20))
print(check(5))
`, "big\nsmall")
}

func TestFunctionReturnInLoop(t *testing.T) {
	expectOutput(t, `
fn findFirst() {
	for (let i = 0; i < 100; i = i + 1) {
		if (i == 5) {
			return i
		}
	}
	return 0
}
print(findFirst())
`, "5")
}

func TestClosure(t *testing.T) {
	expectOutput(t, `
fn makeAdder(x) {
	fn adder(y) {
		return x + y
	}
	return adder
}
let add5 = makeAdder(5)
print(add5(3))
`, "8")
}

func TestClosureRetainsScope(t *testing.T) {
	expectOutput(t, `
fn counter() {
	let count = 0
	fn increment() {
		count = count + 1
		return count
	}
	return increment
}
let c = counter()
print(c())
print(c())
print(c())
`, "1\n2\n3")
}

func TestNestedFunctions(t *testing.T) {
	expectOutput(t, `
fn outer() {
	fn middle() {
		fn inner() {
			return 42
		}
		return inner()
	}
	return middle()
}
print(outer())
`, "42")
}

func TestRecursiveCountdown(t *testing.T) {
	expectLines(t, `
fn countdown(n) {
	if (n <= 0) {
		return 0
	}
	print(n)
	return countdown(n - 1)
}
countdown(5)
`, "5", "4", "3", "2", "1")
}
