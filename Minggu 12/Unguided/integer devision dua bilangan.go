package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y (y <= x): ")
	fmt.Scan(&y)

	if x < 0 || y <= 0 || x < y {
		fmt.Println("Input tidak valid. Pastikan x >= y dan y > 0.")
		return
	}

	hasil := 0
	sisa := x

	for sisa >= y {
		sisa -= y
		hasil++
	}

	fmt.Printf("%d div %d = %d\n", x, y, hasil)
}