package main
import (
	"fmt"
)
func main() {
	var farenheit int

	// Memasukan nilai suhu
	fmt.Print("Masukan suhu dalam derajat farenheit: ")
	fmt.Scanln(&farenheit)

	// Menghitung suhu dalam derajat celcius
	Celcius := float64(farenheit-32)*5/9 + 32

	// Menampilkan hasil dalam derajat celcius
	fmt.Printf("suhu dalam celcius: %.0f\n", Celcius)
}