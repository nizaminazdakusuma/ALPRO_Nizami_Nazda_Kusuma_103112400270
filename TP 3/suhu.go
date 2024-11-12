package main
import (
	"fmt"
)
func main() {
	var farenheit float64

	fmt.Print("Masukan suhu dalam derajat farenheit: ")
	fmt.Scan(&farenheit)

	kelvin := (farenheit-32)*5/9 + 273.15

	fmt.Printf("suhu dalam derajat kelvin: %.f k\n", kelvin)
}