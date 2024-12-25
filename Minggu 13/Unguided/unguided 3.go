package main

import (
	"fmt"
)

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	if target <= 0 {
		fmt.Println("Silakan masukkan target donasi yang valid.")
		return
	}

	for totalDonasi < target {
		jumlahDonatur++
		fmt.Printf("Masukkan donasi dari donatur %d: ", jumlahDonatur)
		fmt.Scan(&donasi)

		if donasi < 0 {
			fmt.Println("Silakan masukkan jumlah donasi yang valid.")
			jumlahDonatur-- 
			continue
		}

		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}