export const EXAMPLES: Record<string, string> = {
  'Hello World': `# Hello World in Gor
print("Hello, World!")

let name = "Gor"
print("Welcome to " + name)`,

  'Fibonacci': `# Fibonacci — recursive and iterative

fn fibRecursive(n) {
    if (n <= 1) {
        return n
    }
    return fibRecursive(n - 1) + fibRecursive(n - 2)
}

fn fibIterative(n) {
    let a = 0
    let b = 1
    for (let i = 0; i < n; i = i + 1) {
        let temp = b
        b = a + b
        a = temp
    }
    return a
}

print("Fibonacci (recursive):")
for (let i = 0; i < 10; i = i + 1) {
    print(fibRecursive(i))
}

print("Fibonacci (iterative):")
for (let i = 0; i < 20; i = i + 1) {
    print(fibIterative(i))
}`,

  'Bubble Sort': `# Bubble Sort

fn bubbleSort(arr, n) {
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

let data = [64, 34, 25, 12, 22, 11, 90]

print("Before sort:")
for (let i = 0; i < 7; i = i + 1) {
    print(data[i])
}

bubbleSort(data, 7)

print("After sort:")
for (let i = 0; i < 7; i = i + 1) {
    print(data[i])
}`,

  'Binary Search': `# Binary Search on a sorted array

fn binarySearch(arr, target, low, high) {
    if (low > high) {
        return 0 - 1
    }
    let mid = (low + high) / 2
    if (arr[mid] == target) {
        return mid
    }
    if (arr[mid] < target) {
        return binarySearch(arr, target, mid + 1, high)
    }
    return binarySearch(arr, target, low, mid - 1)
}

let sorted = [2, 5, 8, 12, 16, 23, 38, 45, 56, 72, 91]

print("Searching for 23:")
print(binarySearch(sorted, 23, 0, 10))

print("Searching for 56:")
print(binarySearch(sorted, 56, 0, 10))

print("Searching for 99 (not found):")
print(binarySearch(sorted, 99, 0, 10))`,

  'Closures': `# Closure Patterns

fn makeCounter(start) {
    let count = start
    fn increment() {
        count = count + 1
        return count
    }
    return increment
}

let counter = makeCounter(0)
print("Counter:")
print(counter())
print(counter())
print(counter())

fn makeAdder(x) {
    fn add(y) {
        return x + y
    }
    return add
}

let add10 = makeAdder(10)
let add50 = makeAdder(50)
print("add10(5):")
print(add10(5))
print("add50(5):")
print(add50(5))

fn makeAccumulator() {
    let total = 0
    fn accumulate(n) {
        total = total + n
        return total
    }
    return accumulate
}

let acc = makeAccumulator()
print("Accumulator:")
print(acc(10))
print(acc(20))
print(acc(30))`,

  'Stack': `# Stack — Last In, First Out (LIFO)

let stack = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
let top = 0

fn push(val) {
    stack[top] = val
    top = top + 1
    return val
}

fn pop() {
    if (top == 0) {
        print("Stack underflow!")
        return 0
    }
    top = top - 1
    return stack[top]
}

fn peek() {
    if (top == 0) { return 0 }
    return stack[top - 1]
}

fn size() { return top }

push(10)
push(20)
push(30)

print("Size after 3 pushes:")
print(size())
print("Peek:")
print(peek())
print("Pop:")
print(pop())
print(pop())
print("Size after 2 pops:")
print(size())`,

  'Objects': `# Hash Map — using objects as key-value stores

let map = {
    alice: 95,
    bob: 87,
    charlie: 92,
    diana: 88
}

print("Alice's score:")
print(map.alice)

map.bob = 91
print("Bob's updated score:")
print(map.bob)

let students = {
    alice: { grade: 95, passed: true },
    bob: { grade: 91, passed: true },
    charlie: { grade: 45, passed: false }
}

print("Charlie passed:")
print(students.charlie.passed)
print("Charlie's grade:")
print(students.charlie.grade)`,
}
