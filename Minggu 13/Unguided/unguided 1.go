package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	if bilangan <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif.")
		return
	}

	banyakDigit := hitungBanyakDigit(bilangan)

	fmt.Printf("Banyaknya digit dari bilangan %d adalah: %d\n", bilangan, banyakDigit)
}
func hitungBanyakDigit(bilangan int) int {
	return len(strconv.Itoa(bilangan))
}