package main
import (
	"fmt"
)
func main() {
	var bilanganBulat int

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 1 {
		fmt.Println("Bukan positif")
	} else {
		fmt.Println("Positif")
	}
}