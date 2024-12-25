package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)

	if  n <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif yang valid.")
		return
	}

	bilanganGanjil := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 { 
			bilanganGanjil++
		}
	}

	fmt.Printf("Terdapat %d bilangan ganjil\n", bilanganGanjil)
}