package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input)

		fmt.Printf("Anda mengetik: %s\n", input)

		if strings.ToLower(input) == "telkom" {
			break
		}
	}

	fmt.Println("Program selesai.")
}