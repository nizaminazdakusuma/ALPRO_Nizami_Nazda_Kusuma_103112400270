package main

import (
	"fmt"
)

func isPerfect(number int) bool {
	if number < 1 {
		return false
	}

	sum := 0
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	return sum == number
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)

	if isPerfect(N) {
		fmt.Printf("%d adalah bilangan sempurna.\n", N)
	} else {
		fmt.Printf("%d bukan bilangan sempurna.\n", N)
	}
}