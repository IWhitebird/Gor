package main

import "fmt"

func Testo() {
	printPrime(10000)
}

func checkPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrime(n int) {
	for i := 2; i < n; i++ {
		if checkPrime(i) {
			fmt.Println(i)
		}
	}
}
