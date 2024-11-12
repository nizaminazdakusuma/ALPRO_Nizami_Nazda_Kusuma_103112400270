package main
import (
	"fmt"
	"math"
)
func main() {
	// Input dari pengguna
	var radius float64
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	// Menghitung luas dan keliling lingkaran
	luas := math.Pi * math.Pow(radius, 2)
	keliling := 2 * math.Pi * radius
	
	// Menampilkan hasil
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}