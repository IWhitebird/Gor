package test

import "testing"

// Operator precedence (low to high):
//   Assignment → Object → Logical(&&/||) → Comparison(==/!=/>/</>=/<=)
//   → Additive(+/-/&/|) → Multiplicative(*/ / /%) → Call/Member → Primary

func TestPrecedenceMultOverAdd(t *testing.T) {
	// * binds tighter than +
	expectOutput(t, "print(2 + 3 * 4)", "14")
	expectOutput(t, "print(3 * 4 + 2)", "14")
}

func TestPrecedenceDivOverSub(t *testing.T) {
	expectOutput(t, "print(10 - 6 / 2)", "7")
}

func TestPrecedenceComparisonOverLogical(t *testing.T) {
	// && binds looser than ==
	// true && true should be true
	expectOutput(t, "print(1 == 1 && 2 == 2)", "true")
	expectOutput(t, "print(1 == 2 || 3 == 3)", "true")
}

func TestPrecedenceAddOverComparison(t *testing.T) {
	// + binds tighter than <
	// j < n - 1 means j < (n - 1), NOT (j < n) - 1
	expectOutput(t, `
let j = 3
let n = 5
if (j < n - 1) {
    print("yes")
} else {
    print("no")
}
`, "yes")
}

func TestPrecedenceSubtractInComparison(t *testing.T) {
	// 5 > 10 - 7 means 5 > (10 - 7) = 5 > 3 = true
	expectOutput(t, "print(5 > 10 - 7)", "true")
}

func TestPrecedenceAddInEquality(t *testing.T) {
	// 2 + 3 == 5 means (2 + 3) == 5 = true
	expectOutput(t, "print(2 + 3 == 5)", "true")
}

func TestPrecedenceMultInComparison(t *testing.T) {
	// 2 * 3 > 5 means (2 * 3) > 5 = 6 > 5 = true
	expectOutput(t, "print(2 * 3 > 5)", "true")
}

func TestPrecedenceForLoopCondition(t *testing.T) {
	// for (let i = 0; i < n - 1; ...) should parse as i < (n - 1)
	expectOutput(t, `
let count = 0
let n = 5
for (let i = 0; i < n - 1; i = i + 1) {
    count = count + 1
}
print(count)
`, "4")
}

func TestPrecedenceNestedForLoopCondition(t *testing.T) {
	// for (let j = 0; j < n - 1 - i; ...) should parse as j < ((n - 1) - i)
	expectOutput(t, `
let count = 0
let n = 5
for (let i = 0; i < n; i = i + 1) {
    for (let j = 0; j < n - 1 - i; j = j + 1) {
        count = count + 1
    }
}
print(count)
`, "10")
}

func TestPrecedenceMultiplyInLoopBound(t *testing.T) {
	// i * i <= n means (i * i) <= n
	expectOutput(t, `
fn isPrime(n) {
    if (n <= 1) { return false }
    for (let i = 2; i * i <= n; i = i + 1) {
        if (n % i == 0) { return false }
    }
    return true
}
print(isPrime(7))
print(isPrime(9))
print(isPrime(11))
`, "true\nfalse\ntrue")
}

func TestPrecedenceChainedComparisons(t *testing.T) {
	// a < b == true means (a < b) == true
	expectOutput(t, "print((3 < 5) == true)", "true")
}

func TestPrecedenceParensOverride(t *testing.T) {
	// Explicit parens should always win
	expectOutput(t, "print((2 + 3) * 4)", "20")
	expectOutput(t, "print(2 * (3 + 4))", "14")
}

func TestPrecedenceLogicalOrVsAnd(t *testing.T) {
	// && and || at same level, left-to-right
	// false && true || true → (false && true) || true → false || true → true
	expectOutput(t, `
if (false && true || true) {
    print("yes")
} else {
    print("no")
}
`, "yes")
}

func TestPrecedenceComparisonNotGroupedWithAdd(t *testing.T) {
	// 1 + 2 < 4 + 5 means (1 + 2) < (4 + 5) = 3 < 9 = true
	expectOutput(t, "print(1 + 2 < 4 + 5)", "true")
}

func TestPrecedenceModuloOverComparison(t *testing.T) {
	// n % 2 == 0 means (n % 2) == 0
	expectOutput(t, `
let n = 10
print(n % 2 == 0)
`, "true")
}

func TestPrecedenceBubbleSortPattern(t *testing.T) {
	// The exact pattern from bubble sort that used to crash
	expectOutput(t, `
fn sort(arr, n) {
    for (let i = 0; i < n - 1; i = i + 1) {
        for (let j = 0; j < n - 1 - i; j = j + 1) {
            if (arr[j] > arr[j + 1]) {
                let temp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp
            }
        }
    }
    return arr
}
let a = [3, 1, 2]
sort(a, 3)
print(a[0])
print(a[1])
print(a[2])
`, "1\n2\n3")
}
