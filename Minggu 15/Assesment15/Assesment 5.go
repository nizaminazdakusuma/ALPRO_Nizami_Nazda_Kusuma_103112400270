package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan ganjil: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j || i+j == n+1 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}