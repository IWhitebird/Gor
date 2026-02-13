package test

import "testing"

func TestObjectLiteral(t *testing.T) {
	expectOutput(t, `
let obj = { x: 10 }
print(obj.x)
`, "10")
}

func TestObjectMultipleProperties(t *testing.T) {
	expectOutput(t, `
let obj = { a: 1, b: 2, c: 3 }
print(obj.a)
print(obj.b)
print(obj.c)
`, "1\n2\n3")
}

func TestObjectStringProperty(t *testing.T) {
	expectOutput(t, `
let obj = { name: "gor" }
print(obj.name)
`, "gor")
}

func TestObjectPropertyAssignment(t *testing.T) {
	expectOutput(t, `
let obj = { x: 1 }
obj.x = 42
print(obj.x)
`, "42")
}

func TestNestedObject(t *testing.T) {
	expectOutput(t, `
let obj = {
	inner: {
		value: 99
	}
}
print(obj.inner.value)
`, "99")
}

func TestNestedObjectAssignment(t *testing.T) {
	expectOutput(t, `
let obj = {
	a: {
		b: "old"
	}
}
obj.a.b = "new"
print(obj.a.b)
`, "new")
}

func TestObjectShorthandProperty(t *testing.T) {
	// { key } without value uses variable from scope
	expectOutput(t, `
let x = 42
let obj = { x }
print(obj.x)
`, "42")
}

func TestObjectEquality(t *testing.T) {
	expectOutput(t, `
let a = { x: 1, y: 2 }
let b = { x: 1, y: 2 }
print(a == b)
`, "true")
}

func TestObjectInequalityDifferentValues(t *testing.T) {
	expectOutput(t, `
let a = { x: 1 }
let b = { x: 2 }
print(a == b)
`, "false")
}

func TestObjectInequalityDifferentKeys(t *testing.T) {
	expectOutput(t, `
let a = { x: 1 }
let b = { y: 1 }
print(a == b)
`, "false")
}

func TestObjectNotEqual(t *testing.T) {
	expectOutput(t, `
let a = { x: 1 }
let b = { x: 2 }
print(a != b)
`, "true")
}

func TestObjectAsArgument(t *testing.T) {
	expectOutput(t, `
fn getX(obj) {
	return obj.x
}
let o = { x: 77 }
print(getX(o))
`, "77")
}

func TestObjectWithMultipleTypes(t *testing.T) {
	expectOutput(t, `
let obj = {
	num: 42,
	str: "hello",
	flag: true
}
print(obj.num)
print(obj.str)
print(obj.flag)
`, "42\nhello\ntrue")
}

func TestObjectCreatedInFunction(t *testing.T) {
	expectOutput(t, `
fn makePoint(x, y) {
	return { x: x, y: y }
}
let p = makePoint(3, 4)
print(p.x)
print(p.y)
`, "3\n4")
}

func TestObjectInArray(t *testing.T) {
	expectOutput(t, `
let arr = [{ val: 1 }, { val: 2 }]
print(arr[0].val)
print(arr[1].val)
`, "1\n2")
}
