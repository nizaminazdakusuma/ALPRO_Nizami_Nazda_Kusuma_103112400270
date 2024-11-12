package main
import (
	"fmt"
	"math"
)
func main() {
	var jariJari float64

	// Memasukan input yang diberikan oleh pengguna
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	// Menghitung luas lingkaran
	luas := math.Pi * math.Pow(jariJari, 2)

	// Menampilkan hasil 
	fmt.Printf("Luas lingkaran: %2f\n ", luas)
}