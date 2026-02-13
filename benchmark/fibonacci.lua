-- Recursive Fibonacci - benchmark function calls and recursion

function fib(n)
    if n <= 1 then
        return n
    end
    return fib(n - 1) + fib(n - 2)
end

local result = fib(30)
print(result)
