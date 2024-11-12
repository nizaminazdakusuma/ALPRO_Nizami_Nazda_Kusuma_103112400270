package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	var tahun int

	// Menggunakan loop untuk input
	for i := 1; i <= 3; i++ {
		fmt.Printf("Tahun: ")
		fmt.Scan(&tahun)

		// Memeriksa apakah tahun tersebut adalah kabisat
		kabisat := isLeapYear(tahun)

		// Menampilkan hasil apakah tahun tersebut adalah kabisat
		fmt.Printf("Kabisat: %t\n", kabisat)
	}
}