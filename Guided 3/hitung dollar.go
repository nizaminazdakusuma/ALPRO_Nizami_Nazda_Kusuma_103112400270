package main

import "fmt"

func main() {
	var rupiah, dolar int

	// Memasukan nilai rupiah
	fmt.Println("Memasukan nilai rupiah: ")
	fmt.Scan(&rupiah)

	// Mengganti nilai rupiah ke dolar 
	dolar = rupiah / 15000
	fmt.Println(dolar)
}