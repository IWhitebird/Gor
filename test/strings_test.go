package test

import "testing"

func TestStringLiteral(t *testing.T) {
	expectOutput(t, `print("hello")`, "hello")
}

func TestEmptyString(t *testing.T) {
	expectOutput(t, `print("")`, "")
}

func TestStringConcatenation(t *testing.T) {
	expectOutput(t, `print("hello" + " " + "world")`, "hello world")
}

func TestStringConcatWithVariables(t *testing.T) {
	expectOutput(t, `
let a = "foo"
let b = "bar"
print(a + b)
`, "foobar")
}

func TestStringEquality(t *testing.T) {
	expectOutput(t, `print("abc" == "abc")`, "true")
}

func TestStringInequality(t *testing.T) {
	expectOutput(t, `print("abc" != "def")`, "true")
}

func TestStringEqualityFalse(t *testing.T) {
	expectOutput(t, `print("abc" == "def")`, "false")
}

func TestStringComparison(t *testing.T) {
	expectOutput(t, `print("a" < "b")`, "true")
	expectOutput(t, `print("b" > "a")`, "true")
}

func TestStringWithSpaces(t *testing.T) {
	expectOutput(t, `print("hello world")`, "hello world")
}

func TestStringWithNumbers(t *testing.T) {
	expectOutput(t, `print("test123")`, "test123")
}

func TestStringReassignment(t *testing.T) {
	expectOutput(t, `
let s = "first"
s = "second"
print(s)
`, "second")
}

func TestStringLengthComparison(t *testing.T) {
	expectOutput(t, `print("ab" >= "a")`, "true")
	expectOutput(t, `print("a" <= "ab")`, "true")
}

func TestMultipleStringConcat(t *testing.T) {
	expectOutput(t, `
let a = "a"
let b = a + "b"
let c = b + "c"
print(c)
`, "abc")
}
