-- Prime Sieve - count primes under 10000

function isPrime(n)
    if n < 2 then
        return false
    end
    local i = 2
    while i * i <= n do
        if n % i == 0 then
            return false
        end
        i = i + 1
    end
    return true
end

local count = 0
for i = 2, 9999 do
    if isPrime(i) then
        count = count + 1
    end
end

print(count)
