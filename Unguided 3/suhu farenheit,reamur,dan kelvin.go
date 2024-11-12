package main
import (
	"fmt"
)

// Rumus untuk mengubah celcius ke fahrenheit
func celciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9/5) + 32
}

func main() {
	var celcius float64

	// Memasukan suhu celcius
	fmt.Print("Memasukan suhu celcius: ")
	fmt.Scan(&celcius)
	
	// Mengubah suhu celcius ke farenheit, reamur, dan kelvin
	fahrenheit := celciusToFahrenheit(celcius)

	// Menampilkan hasil dalam suhu fahrenheit
	fmt.Printf("%.f derajat celcius setara dengan %.f derajat Fahrenheit\n", celcius, fahrenheit)
}