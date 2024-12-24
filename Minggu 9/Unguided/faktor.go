package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y: ")
	fmt.Scan(&y)

	isXFaktorY := (y % x == 0)

	isYFaktorX := (x % y == 0)

	fmt.Println(isXFaktorY)
	fmt.Println(isYFaktorX)
}