package test

import "testing"

func TestCommentIgnored(t *testing.T) {
	expectOutput(t, `
# this is a comment
print(1)
# another comment
`, "1")
}

func TestCommentAtEndOfLine(t *testing.T) {
	expectOutput(t, `
let x = 10 # inline comment
print(x)
`, "10")
}

func TestMultipleStatements(t *testing.T) {
	expectLines(t, `
let a = 1
let b = 2
let c = 3
print(a)
print(b)
print(c)
`, "1", "2", "3")
}

func TestComplexProgram(t *testing.T) {
	// Combines many features
	expectOutput(t, `
fn isPrime(n) {
	if (n <= 1) {
		return false
	}
	for (let i = 2; i * i <= n; i = i + 1) {
		if (n % i == 0) {
			return false
		}
	}
	return true
}

let count = 0
for (let i = 2; i < 20; i = i + 1) {
	if (isPrime(i)) {
		count = count + 1
	}
}
print(count)
`, "8")
}

func TestFunctionReturningObject(t *testing.T) {
	expectOutput(t, `
fn pair(a, b) {
	return { first: a, second: b }
}
let p = pair(10, 20)
print(p.first + p.second)
`, "30")
}

func TestArrayOfObjects(t *testing.T) {
	expectOutput(t, `
let items = [
	{ name: "a", val: 1 },
	{ name: "b", val: 2 },
	{ name: "c", val: 3 }
]
let total = 0
for (let i = 0; i < 3; i = i + 1) {
	total = total + items[i].val
}
print(total)
`, "6")
}

func TestRecursiveGCD(t *testing.T) {
	expectOutput(t, `
fn gcd(a, b) {
	if (b == 0) {
		return a
	}
	return gcd(b, a % b)
}
print(gcd(48, 18))
`, "6")
}

func TestRecursivePower(t *testing.T) {
	expectOutput(t, `
fn power(base, exp) {
	if (exp == 0) {
		return 1
	}
	return base * power(base, exp - 1)
}
print(power(2, 10))
`, "1024")
}

func TestMutuallyDependentData(t *testing.T) {
	expectOutput(t, `
let a = [1, 2, 3]
let b = [4, 5, 6]
let sum = 0
for (let i = 0; i < 3; i = i + 1) {
	sum = sum + a[i] + b[i]
}
print(sum)
`, "21")
}

func TestStringInIfCondition(t *testing.T) {
	expectOutput(t, `
let name = "gor"
if (name == "gor") {
	print("match")
} else {
	print("no match")
}
`, "match")
}

func TestBooleanArithmetic(t *testing.T) {
	// true/false in comparisons
	expectOutput(t, "print(true == true)", "true")
	expectOutput(t, "print(true != false)", "true")
}

func TestObjectWithArrayProperty(t *testing.T) {
	expectOutput(t, `
let obj = {
	items: [10, 20, 30]
}
print(obj.items[1])
`, "20")
}

func TestMultipleFunctions(t *testing.T) {
	expectOutput(t, `
fn add(a, b) { return a + b }
fn sub(a, b) { return a - b }
fn mul(a, b) { return a * b }
print(add(1, mul(2, sub(10, 7))))
`, "7")
}

func TestDeepNesting(t *testing.T) {
	expectOutput(t, `
fn a() {
	fn b() {
		fn c() {
			return 42
		}
		return c()
	}
	return b()
}
print(a())
`, "42")
}

func TestVariableOverwriteMultipleTimes(t *testing.T) {
	expectOutput(t, `
let x = 1
x = 2
x = 3
x = 4
x = 5
print(x)
`, "5")
}

func TestComplexExpression(t *testing.T) {
	expectOutput(t, "print((1 + 2) * (3 + 4) - (5 * 2))", "11")
}

func TestEmptyIfBody(t *testing.T) {
	// If body with only variable declaration, no print
	expectOutput(t, `
if (true) {
	let x = 1
}
print("ok")
`, "ok")
}

func TestBubbleSort(t *testing.T) {
	expectLines(t, `
let arr = [5, 3, 1, 4, 2]
for (let i = 0; i < 5; i = i + 1) {
	for (let j = 0; j < 4; j = j + 1) {
		if (arr[j] > arr[j + 1]) {
			let temp = arr[j]
			arr[j] = arr[j + 1]
			arr[j + 1] = temp
		}
	}
}
for (let i = 0; i < 5; i = i + 1) {
	print(arr[i])
}
`, "1", "2", "3", "4", "5")
}

func TestFibIterative(t *testing.T) {
	expectOutput(t, `
fn fibIter(n) {
	let a = 0
	let b = 1
	for (let i = 0; i < n; i = i + 1) {
		let temp = b
		b = a + b
		a = temp
	}
	return a
}
print(fibIter(10))
`, "55")
}

func TestObjectCreatedInsideLoop(t *testing.T) {
	expectLines(t, `
for (let i = 0; i < 3; i = i + 1) {
	let obj = { val: i * 10 }
	print(obj.val)
}
`, "0", "10", "20")
}
