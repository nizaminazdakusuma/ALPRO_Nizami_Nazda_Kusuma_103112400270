package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string

	// Memasukan sebuah huruf
	fmt.Print("Masukkan huruf: ")
	fmt.Scanln(&huruf)

	// Mengubah huruf menjadi huruf kapital 
	huruf = strings.ToUpper(huruf)

	// Memeriksa apakah huruf adalah huruf vokal
	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
	} else if len(huruf) == 1 && (huruf >= "B" && huruf <= "Z") { // Memeriksa apakah huruf adalah huruf konsonan
		fmt.Println("Huruf Konsonan")
	} else {
		fmt.Println("Input tidak valid, silakan masukkan satu huruf.")
	}
}