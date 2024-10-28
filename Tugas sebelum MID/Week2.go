package main
import (
	"fmt"
)
func main() {
	panjang := 10
	lebar := 12

	keliling := 2 * panjang + 2 * lebar
	luas := panjang * lebar

	fmt.Printf("keliling persegi panjang: %d meter\n", keliling)
	fmt.Printf("Luas persegi panjang: %d meter persegi\n", luas)
}