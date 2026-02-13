// Prime Sieve - count primes under 10000

function isPrime(n) {
    if (n < 2) return false;
    for (let i = 2; i * i <= n; i++) {
        if (n % i === 0) return false;
    }
    return true;
}

let count = 0;
for (let i = 2; i < 10000; i++) {
    if (isPrime(i)) count++;
}

console.log(count);
