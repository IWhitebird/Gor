package test

import "testing"

func TestBasicClosure(t *testing.T) {
	expectOutput(t, `
fn makeGreeter(greeting) {
	fn greet(name) {
		return greeting + " " + name
	}
	return greet
}
let hi = makeGreeter("hi")
print(hi("world"))
`, "hi world")
}

func TestClosureCounter(t *testing.T) {
	expectLines(t, `
fn counter() {
	let n = 0
	fn inc() {
		n = n + 1
		return n
	}
	return inc
}
let c = counter()
print(c())
print(c())
print(c())
`, "1", "2", "3")
}

func TestTwoIndependentClosures(t *testing.T) {
	// Two closures from same factory should have independent state
	expectLines(t, `
fn counter() {
	let n = 0
	fn inc() {
		n = n + 1
		return n
	}
	return inc
}
let a = counter()
let b = counter()
print(a())
print(a())
print(b())
print(a())
print(b())
`, "1", "2", "1", "3", "2")
}

func TestClosureOverMultipleVars(t *testing.T) {
	expectOutput(t, `
fn make(x, y) {
	fn sum() {
		return x + y
	}
	return sum
}
let f = make(10, 20)
print(f())
`, "30")
}

func TestClosureModifiesOuter(t *testing.T) {
	expectOutput(t, `
fn make() {
	let val = 0
	fn set(v) {
		val = v
		return val
	}
	fn get() {
		return val
	}
	return [set, get]
}
let pair = make()
let setter = pair[0]
let getter = pair[1]
setter(42)
print(getter())
`, "42")
}

func TestClosureOverLoop(t *testing.T) {
	expectOutput(t, `
fn makeAdder(n) {
	fn add(x) {
		return n + x
	}
	return add
}
let adders = [makeAdder(1), makeAdder(2), makeAdder(3)]
let sum = 0
for (let i = 0; i < 3; i = i + 1) {
	sum = sum + adders[i](10)
}
print(sum)
`, "36")
}

func TestNestedClosures(t *testing.T) {
	expectOutput(t, `
fn outer(x) {
	fn middle(y) {
		fn inner(z) {
			return x + y + z
		}
		return inner
	}
	return middle
}
let f = outer(1)
let g = f(2)
print(g(3))
`, "6")
}
