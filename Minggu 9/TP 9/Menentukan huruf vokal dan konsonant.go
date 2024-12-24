package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scanln(&huruf)

	huruf = strings.ToUpper(huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
	} else if len(huruf) == 1 && (huruf >= "B" && huruf <= "Z") {
		fmt.Println("Huruf Konsonan")
	} else {
		fmt.Println("Input tidak valid, silakan masukkan satu huruf.")
	}
}