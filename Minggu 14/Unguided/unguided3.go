package main

import (
	"fmt"
	"strings"
)

func main() {
	expectedColors := []string{"merah", "kuning", "hijau", "ungu"}
	var berhasil bool = true

	for i := 1; i <= 5; i++ {
		var input string
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&input)

		colors := strings.Fields(input)

		if len(colors) != 4 {
			fmt.Println("Input tidak valid, harus ada 4 warna.")
			berhasil = false
			continue
		}

		if !isEqual(colors, expectedColors) {
			berhasil = false
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}