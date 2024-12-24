package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Hasil kuadrat dari bilangan 1 hingga %d:\n", N)
	for i := 1; i <= N; i++ {
		kuadrat := i * i
		fmt.Printf("%d ", kuadrat)
	}
	fmt.Println()
}