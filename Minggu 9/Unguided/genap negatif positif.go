package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 && bilangan%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}