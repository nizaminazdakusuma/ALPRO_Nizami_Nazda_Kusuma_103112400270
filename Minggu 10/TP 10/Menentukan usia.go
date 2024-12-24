package main

import (
	"fmt"
)

func main() {

	// Memasukkan usia
	var usia int
	
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	// Menggunakan switch case untuk menentukan kategori usia
	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori Usia: Anak-anak (0 - 12 tahun)")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori Usia: Remaja (13 - 17 tahun)")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori Usia: Dewasa (18 - 64 tahun)")
	case usia >= 65:
		fmt.Println("Kategori Usia: Lansia (65 tahun ke atas)")
	default:
		fmt.Println("Usia tidak valid")
	}
}