package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah kerucut (n): ")

	for i := 0; i < n; i++ {
		var r, h float64

		fmt.Printf("Masukkan jari-jari alas dan tinggi kerucut ke-%d: ", i+1)

		volume := (1.0 / 3.0) * math.Pi * r * r * h

		fmt.Printf("Volume kerucut ke-%d: %f\n", i+1, volume)
	}
}