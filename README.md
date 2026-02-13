<img src="asset/Gor.png" width="200" height="200">

# Gor Language

Gor is an interpreted programming language made with Golang. It has similar syntax to that of JavaScript. Zero external dependencies — pure Go standard library.

## Prerequisites

- [Go](https://go.dev/dl/) 1.21 or later

## Installation

```bash
git clone https://github.com/IWhitebird/Gor.git
cd Gor
make build
```

This builds the binary to `bin/gor`.

## Usage

```
gor <file.gor>           Run a Gor source file
gor --repl               Start the interactive REPL
gor --ast <file.gor>     Print the AST of a source file as JSON
gor --version            Print version
gor --help               Show help
```

### Run a file

```bash
./bin/gor examples/fibonacci.gor
```

### Start the REPL

```bash
./bin/gor --repl
```

### Print the AST

```bash
./bin/gor --ast examples/fibonacci.gor
```

## Development

```bash
make build      # Build binary to bin/gor
make run        # Build and start REPL
make test       # Run test suite
make test-v     # Run tests with verbose output
make fmt        # Format all Go source files
make clean      # Remove built binary
```

Or without Make:

```bash
go build -o bin/gor ./cmd/gor/
go test ./test/ -count=1
go fmt ./...
```

## Data Types

Gor has 6 data types: **number** (integers), **string**, **boolean**, **null**, **object**, and **array**.

```
let num = 42
let str = "hello"
let flag = true
let empty = null
let obj = { x: 1, y: 2 }
let arr = [10, 20, 30]
```

## Variables

```
let a = 10        # mutable
const b = 20      # immutable
a = 30            # ok
```

## Operators

Arithmetic:
```
+  -  *  /  %
```

Comparison:
```
==  !=  >  <  >=  <=
```

Logical:
```
&&  ||
```

Bitwise:
```
&  |
```

String concatenation:
```
print("hello" + " " + "world")
```

## Print

```
print(42)
print("hello")
print(true)
```

## Comments

```
# This is a comment
let x = 10 # inline comment
```

## If / Else

```
if (x > 10) {
    print("big")
} else if (x > 5) {
    print("medium")
} else {
    print("small")
}
```

Conditions support truthy/falsy values — non-zero numbers, non-empty strings, and non-null values are truthy.

## Loops

Currently the only loop Gor supports is the `for` loop:

```
for (let i = 0; i < 10; i = i + 1) {
    print(i)
}
```

## Functions

Functions are declared with `fn` and support closures:

```
fn add(a, b) {
    return a + b
}
print(add(10, 20))
```

### Closures

Functions capture their enclosing scope:

```
fn makeCounter() {
    let count = 0
    fn increment() {
        count = count + 1
        return count
    }
    return increment
}
let c = makeCounter()
print(c())   # 1
print(c())   # 2
```

## Arrays

```
let arr = [10, 20, 30]
print(arr[0])          # 10
print(arr[1 + 1])      # 30
arr[1] = 99
```

## Objects

```
let obj = {
    name: "gor",
    version: 1
}
print(obj.name)
obj.version = 2

# Nested access
let data = { inner: { value: 42 } }
print(data.inner.value)
```

## Examples

See the `examples/` folder for data structure implementations:

- **`stack.gor`** — Stack (LIFO) with push/pop/peek
- **`queue.gor`** — Queue (FIFO) with enqueue/dequeue
- **`linked_list.gor`** — Linked list using nested objects
- **`binary_search.gor`** — Recursive binary search
- **`bubble_sort.gor`** — Bubble sort algorithm
- **`selection_sort.gor`** — Selection sort algorithm
- **`hash_map.gor`** — Key-value store using objects
- **`fibonacci.gor`** — Recursive and iterative Fibonacci
- **`closure_patterns.gor`** — Counter, adder, accumulator, getter/setter

## Project Structure

```
Gor/
├── cmd/gor/              # CLI entry point
├── core/
│   ├── ast/              # AST node definitions
│   ├── lexer/            # Tokenizer
│   ├── parser/           # Recursive descent parser
│   ├── interpreter/      # Tree-walking evaluator
│   └── program/          # REPL and input handling
├── gor.go                # Public API (Repl, RunFromFile, PrintAST)
├── examples/             # Example Gor programs
├── test/                 # Test suite (238 tests)
├── benchmark/            # Performance benchmarks
├── asset/                # Logo and images
└── Makefile
```

## Future Updates

- While loop
- Unary operators (`--`, `++`)
- Multithreading
