package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	if n % 2 == 0 {
		fmt.Println("Masukan harus berupa bilangan bulat ganjil.")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j || i + j == n+1 {
				fmt.Print(j, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}