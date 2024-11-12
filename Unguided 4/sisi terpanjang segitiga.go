package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	// Menentukan titik koordinat dari x, y
	fmt.Println("Masukkan koordinat titik A (x y):")
	fmt.Scan(&ax, &ay)
	fmt.Println("Masukkan koordinat titik B (x y):")
	fmt.Scan(&bx, &by)
	fmt.Println("Masukkan koordinat titik C (x y):")
	fmt.Scan(&cx, &cy)

	// Menghitung panjang sisi segitiga
	sideAB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	sideBC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	sideCA := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	// Menentukan sisi segitiga terpanjang
	longestSide := sideAB
	if sideBC > longestSide {
		longestSide = sideBC
	}
	if sideCA > longestSide {
		longestSide = sideCA
	}

	// Menampilkan hasil dari sisi terpanjang
	fmt.Printf(" sisi terpanjang dari segitiga adalah: %.f\n", longestSide)
}