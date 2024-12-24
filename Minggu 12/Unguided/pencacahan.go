package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	if number <= 0 {
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}

	var digits []int

	for number > 0 {
		digit := number % 10
		digits = append(digits, digit)
		number = number / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println(digits[i])
	}
}