package main

import (
	"fmt"
)

func main() {
	var base, exponent int

	fmt.Print("Masukkan bilangan bulat positif (basis dan eksponen): ")
	fmt.Scan(&base, &exponent)

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	fmt.Printf("%d dipangkatkan dengan %d adalah: %d\n", base, exponent, result)
}