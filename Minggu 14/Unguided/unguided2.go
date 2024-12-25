package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif yang valid.")
		return
	}

	if isPrima(n) {
		fmt.Println(n, "prima")
	} else {
		fmt.Println(n, "bukan prima")
	}
}

func isPrima(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}