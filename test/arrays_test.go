package test

import "testing"

func TestArrayLiteral(t *testing.T) {
	expectOutput(t, `
let arr = [1, 2, 3]
print(arr[0])
`, "1")
}

func TestArrayIndexAccess(t *testing.T) {
	expectOutput(t, `
let arr = [10, 20, 30]
print(arr[0])
print(arr[1])
print(arr[2])
`, "10\n20\n30")
}

func TestArrayWithStrings(t *testing.T) {
	expectOutput(t, `
let arr = ["hello", "world"]
print(arr[0])
print(arr[1])
`, "hello\nworld")
}

func TestArrayAssignment(t *testing.T) {
	expectOutput(t, `
let arr = [1, 2, 3]
arr[1] = 99
print(arr[1])
`, "99")
}

func TestArrayWithExpressions(t *testing.T) {
	expectOutput(t, `
let arr = [1 + 1, 2 * 3, 10 - 4]
print(arr[0])
print(arr[1])
print(arr[2])
`, "2\n6\n6")
}

func TestEmptyArray(t *testing.T) {
	expectOutput(t, `
let arr = []
print("ok")
`, "ok")
}

func TestArraySingleElement(t *testing.T) {
	expectOutput(t, `
let arr = [42]
print(arr[0])
`, "42")
}

func TestArrayMixedTypes(t *testing.T) {
	expectOutput(t, `
let arr = [1, "hello", true]
print(arr[0])
print(arr[1])
print(arr[2])
`, "1\nhello\ntrue")
}

func TestArrayInLoop(t *testing.T) {
	expectOutput(t, `
let arr = [10, 20, 30, 40, 50]
let sum = 0
for (let i = 0; i < 5; i = i + 1) {
	sum = sum + arr[i]
}
print(sum)
`, "150")
}

func TestArrayModifyInLoop(t *testing.T) {
	expectOutput(t, `
let arr = [0, 0, 0]
for (let i = 0; i < 3; i = i + 1) {
	arr[i] = i * 10
}
print(arr[0])
print(arr[1])
print(arr[2])
`, "0\n10\n20")
}

func TestNestedArrays(t *testing.T) {
	expectOutput(t, `
let matrix = [[1, 2], [3, 4]]
print(matrix[0][0])
print(matrix[0][1])
print(matrix[1][0])
print(matrix[1][1])
`, "1\n2\n3\n4")
}

func TestArrayAsArgument(t *testing.T) {
	expectOutput(t, `
fn first(arr) {
	return arr[0]
}
let a = [99, 88, 77]
print(first(a))
`, "99")
}

func TestArrayComputedIndex(t *testing.T) {
	expectOutput(t, `
let arr = [10, 20, 30]
let i = 1
print(arr[i])
print(arr[i + 1])
`, "20\n30")
}
