package test

import "testing"

func TestBasicForLoop(t *testing.T) {
	expectLines(t, `
for (let i = 0; i < 5; i = i + 1) {
	print(i)
}
`, "0", "1", "2", "3", "4")
}

func TestForLoopSum(t *testing.T) {
	expectOutput(t, `
let sum = 0
for (let i = 1; i <= 10; i = i + 1) {
	sum = sum + i
}
print(sum)
`, "55")
}

func TestForLoopZeroIterations(t *testing.T) {
	// Condition false from start — body never runs
	expectOutput(t, `
for (let i = 10; i < 5; i = i + 1) {
	print("bad")
}
print("done")
`, "done")
}

func TestForLoopSingleIteration(t *testing.T) {
	expectOutput(t, `
for (let i = 0; i < 1; i = i + 1) {
	print("once")
}
`, "once")
}

func TestForLoopCountDown(t *testing.T) {
	expectLines(t, `
for (let i = 3; i > 0; i = i - 1) {
	print(i)
}
`, "3", "2", "1")
}

func TestNestedForLoop(t *testing.T) {
	expectOutput(t, `
let count = 0
for (let i = 0; i < 3; i = i + 1) {
	for (let j = 0; j < 3; j = j + 1) {
		count = count + 1
	}
}
print(count)
`, "9")
}

func TestForLoopWithArray(t *testing.T) {
	expectLines(t, `
let arr = [10, 20, 30]
for (let i = 0; i < 3; i = i + 1) {
	print(arr[i])
}
`, "10", "20", "30")
}

func TestForLoopModifyOuter(t *testing.T) {
	expectOutput(t, `
let x = 0
for (let i = 0; i < 5; i = i + 1) {
	x = x + 2
}
print(x)
`, "10")
}

func TestForLoopStepByTwo(t *testing.T) {
	expectLines(t, `
for (let i = 0; i < 10; i = i + 2) {
	print(i)
}
`, "0", "2", "4", "6", "8")
}

func TestForLoopScopeIsolation(t *testing.T) {
	// Loop variable should not leak
	expectOutput(t, `
let x = 99
for (let i = 0; i < 3; i = i + 1) {
	let y = i
}
print(x)
`, "99")
}

func TestForLoopMultiply(t *testing.T) {
	expectOutput(t, `
let result = 1
for (let i = 1; i <= 5; i = i + 1) {
	result = result * i
}
print(result)
`, "120")
}
