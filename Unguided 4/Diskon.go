package main

import (
	"fmt"
)

func main() {
	var totalHargaAwal, diskonPersen int

	// Memasukan total harga sebelum diskon
	fmt.Print("Masukkan harga awal: ")
	fmt.Scan(&totalHargaAwal)

	// Memasukan diskon
	fmt.Print("Masukkan diskon (dalam persen): ")
	fmt.Scan(&diskonPersen)

	// Menghitung total harga setelah diskon
	diskon := float64(diskonPersen) / 100 * float64(totalHargaAwal)
	totalHargaAkhir := float64(totalHargaAwal) - diskon

	// Menampilkan harga setelah diskon
	fmt.Printf("Total harga setelah diskon: %.f\n", totalHargaAkhir)
}