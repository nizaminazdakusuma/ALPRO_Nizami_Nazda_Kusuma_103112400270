package main

import (
	"fmt"
)

func main() {
	var jumlahBarang int

	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	totalPoin := 0
	for i := 1; i <= jumlahBarang; i++ {
		if i <= 5 {
			totalPoin += 10
		} else {
			totalPoin += 15
		}
	}

	fmt.Printf("Poin yang didapatkan: %d poin\n", totalPoin)
}