package main

import (
	"fmt"
)

func main() {
	var b int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	fmt.Print("Faktor: ")
	isPrima := true
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			if i != 1 && i != b {
				isPrima = false
			}
		}
	}
	fmt.Println()

	fmt.Println("Prima:", isPrima)
}