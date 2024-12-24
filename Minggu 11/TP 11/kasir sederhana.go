package main

import (
	"fmt"
)

func main() {
	var total float64
	var pilihBarang string
	var namaBarang string
	var hargaBarang float64

	fmt.Println("Selamat datang di aplikasi kasir sederhana!")
	fmt.Println("Tekan 'p' jika ingin menambah barang, 'n' untuk menyelesaikan transaksi.")

	for {
		fmt.Print("Apakah Anda ingin menambahkan barang? (p/n): ")
		fmt.Scan(&pilihBarang)

		if pilihBarang == "p" {
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&namaBarang)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scan(&hargaBarang)

			total += hargaBarang
			fmt.Printf("Total belanja: Rp%.2f\n", total)
		} else if pilihBarang == "n" {
			fmt.Printf("Total belanja: Rp%.2f\n", total)
			fmt.Println("Transaksi berhasil. selamat datang kembali")
			break
		} else {
			fmt.Println("Input tidak valid. Silakan coba lagi.")
		}
	}
}