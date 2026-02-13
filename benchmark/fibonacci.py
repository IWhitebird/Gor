"""Recursive Fibonacci - benchmark function calls and recursion"""

def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)

result = fib(30)
print(result)
