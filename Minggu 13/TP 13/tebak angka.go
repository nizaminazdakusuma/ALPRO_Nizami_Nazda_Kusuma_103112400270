package main

import (
	"fmt"
)

func main() {
	const secretNumber = 7
	var guess int

	for {
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}