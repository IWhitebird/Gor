package test

import "testing"

func TestIfTrue(t *testing.T) {
	expectOutput(t, `
if (true) {
	print("yes")
}
`, "yes")
}

func TestIfFalse(t *testing.T) {
	expectOutput(t, `
if (false) {
	print("bad")
}
print("done")
`, "done")
}

func TestIfElse(t *testing.T) {
	expectOutput(t, `
if (false) {
	print("bad")
} else {
	print("good")
}
`, "good")
}

func TestIfElseIf(t *testing.T) {
	expectOutput(t, `
let x = 2
if (x == 1) {
	print("one")
} else if (x == 2) {
	print("two")
} else {
	print("other")
}
`, "two")
}

func TestIfElseIfElse(t *testing.T) {
	expectOutput(t, `
let x = 5
if (x == 1) {
	print("one")
} else if (x == 2) {
	print("two")
} else {
	print("other")
}
`, "other")
}

func TestIfWithExpression(t *testing.T) {
	expectOutput(t, `
let x = 10
if (x > 5) {
	print("big")
} else {
	print("small")
}
`, "big")
}

func TestNestedIf(t *testing.T) {
	expectOutput(t, `
let x = 10
if (x > 5) {
	if (x > 8) {
		print("very big")
	} else {
		print("medium")
	}
}
`, "very big")
}

func TestIfTruthyNumber(t *testing.T) {
	expectOutput(t, `
if (1) {
	print("truthy")
}
`, "truthy")
}

func TestIfFalsyZero(t *testing.T) {
	expectOutput(t, `
if (0) {
	print("bad")
} else {
	print("falsy")
}
`, "falsy")
}

func TestIfTruthyString(t *testing.T) {
	expectOutput(t, `
if ("hello") {
	print("truthy")
}
`, "truthy")
}

func TestIfFalsyEmptyString(t *testing.T) {
	expectOutput(t, `
if ("") {
	print("bad")
} else {
	print("falsy")
}
`, "falsy")
}

func TestIfFalsyNull(t *testing.T) {
	expectOutput(t, `
if (null) {
	print("bad")
} else {
	print("null is falsy")
}
`, "null is falsy")
}

func TestIfWithLogicalAnd(t *testing.T) {
	expectOutput(t, `
if (true && true) {
	print("both")
}
`, "both")
}

func TestIfWithLogicalOr(t *testing.T) {
	expectOutput(t, `
if (false || true) {
	print("one")
}
`, "one")
}

func TestIfWithComparison(t *testing.T) {
	expectOutput(t, `
let a = 5
let b = 10
if (a < b && b > 0) {
	print("yes")
}
`, "yes")
}

func TestMultipleElseIf(t *testing.T) {
	expectOutput(t, `
let x = 3
if (x == 1) {
	print("one")
} else if (x == 2) {
	print("two")
} else if (x == 3) {
	print("three")
} else if (x == 4) {
	print("four")
} else {
	print("other")
}
`, "three")
}
