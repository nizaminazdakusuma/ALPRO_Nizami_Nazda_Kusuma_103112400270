package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	fmt.Printf("Bilangan prima dari 1 hingga %d adalah:\n", N)

	for i := 1; i <= N; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}