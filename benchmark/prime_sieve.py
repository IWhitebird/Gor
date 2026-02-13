"""Prime Sieve - count primes under 10000"""

def is_prime(n):
    if n < 2:
        return False
    i = 2
    while i * i <= n:
        if n % i == 0:
            return False
        i += 1
    return True

count = 0
for i in range(2, 10000):
    if is_prime(i):
        count += 1

print(count)
