package main

import (
	"fmt"
)

func main() {
	var beratParsel int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	kg := beratParsel / 1000
	gr := beratParsel % 1000

	biayaDasar := kg * 10000
	var biayaTambahan int

	if kg > 10 {
		biayaTambahan = 0
	} else if gr >= 500 {
		biayaTambahan = gr * 5
	} else {
		biayaTambahan = gr * 15
	}

	totalBiaya := biayaDasar + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}