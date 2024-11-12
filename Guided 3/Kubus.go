package main

import "fmt"

func main() {
	var sisi int
	var volume int

	// Memasukan panjang sisi kubus
	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)

	// Menghitung Volume kubus
	volume = sisi * sisi * sisi

	// Menampilkan hasil volume
	fmt.Printf("Volume kubus = %d\n", volume)
}