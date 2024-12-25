package main

import (
	"fmt"
	"strings"
)

func main() {
	var bungaPita string
	var count int

	for {
		var bunga string
		count++

		fmt.Printf("Bunga %d: ", count)
		fmt.Scanln(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			count--
			break
		}

		if bungaPita == "" {
			bungaPita = bunga
		} else {
			bungaPita += " – " + bunga
		}
	}

	fmt.Println("Pita:", bungaPita)
	fmt.Println("Bunga:", count)
}