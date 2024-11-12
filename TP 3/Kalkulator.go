package main

import (
	"fmt"
)

func main() {
	var bilangan1, bilangan2 float64
	var operasi string

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&bilangan1)

	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&bilangan2)

	fmt.Print("Masukkan operasi (+, -, *, /): ")
	fmt.Scan(&operasi)

	var hasil float64
	var valid bool

	// Melakukan operasi dengan kalkulator sesuai tanda
	switch operasi {
	case "+":
		hasil = bilangan1 + bilangan2
		valid = true
	case "-":
		hasil = bilangan1 - bilangan2
		valid = true
	case "*":
		hasil = bilangan1 * bilangan2
		valid = true
	case "/":
		if bilangan2 != 0 {
			hasil = bilangan1 / bilangan2
			valid = true
		} else {
			fmt.Println("Error: Pembagian tidak boleh dengan nol.")
			valid = false
		}
	default:
		fmt.Println("Error: Operasi tidak valid.")
		valid = false
	}

	// Menampilkan hasil penghitungan
	if valid {
		fmt.Printf("Hasil: %.f %s %.f = %.f\n", bilangan1, operasi, bilangan2, hasil)
	}
}