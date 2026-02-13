// Recursive Fibonacci - benchmark function calls and recursion

function fib(n) {
    if (n <= 1) return n;
    return fib(n - 1) + fib(n - 2);
}

const result = fib(30);
console.log(result);
