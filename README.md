![gor](https://github.com/IWhitebird/Gor/assets/115157819/37a84bbc-7c71-4be9-91cf-ca3179399481)
# Gor Language

Gor is interpreted programming langauge made with Golang. It has similar syntax to that of JavaScript.

## Data types:
- Language has 5 datatypes: number, string, null, object, array.

## Variable declaration (static and non-static):
```
let a = 10
const b = ""
```
## Print statement:
```
print(a)
```

## Arithmetic operators & logical operators:
```
- `+`, `-`, `*`, `/`, `,`, `&`, `|`
```

## Comparison operators:
```
- `==`, `<=`, `>=`, `!=`, `&&`, `||`
```

## Array and object declaration:
```
let a = [10, 2]
print(a[0])

const b = {
    hi: 10
}
print(b.hi)
```

## Function declaration:
- Functions can be declared using fn keyword 
```
fn myFunc(a, b) {
    return a + b
}
myFunc(10, 20)
```

## Loops:
- Currently, the only loop Gor supports is the for loop.
```
for (let i = 0; i < 10; i = i + 1) {
    print(i)
}
```

## Future Updates:
- While loop
- Unary operators (--, ++)
- Multithreading
