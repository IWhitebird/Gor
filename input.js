function checkPrime(n) {
    for(let i = 2; i < n; i++){
        if(n % i === 0){
            return false;
        }
    }
    return true;
}

function printPrime(n) {
    for(let i = 2; i < n; i++){
        if(checkPrime(i)){
            console.log(i)
        }
    }
}

printPrime(10000)

