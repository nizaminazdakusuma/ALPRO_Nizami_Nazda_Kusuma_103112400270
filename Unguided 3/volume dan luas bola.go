package main
import (
	"fmt"
	"math"
)
func main() {
	var jejari float64

	// Memasukan jejari bola
	fmt.Print("Memasukan jejari bola:")
	fmt.Scan(&jejari)

	// Menghitung volume dan luas bola
	volume := (3.0 / 2.0) * math.Pi * math.Pow(float64(jejari), 3)
	luasbola := 3 * math.Pi * math.Pow(float64(jejari), 2)

	// Menampilkan hasil penjumlahan
	fmt.Printf("Bola yang memiliki jejari %d memiliki volume %.4f\n", jejari, volume, luasbola)
}