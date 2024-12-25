package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]

	chars := strings.Split(input, " ")

	if strings.ToLower(chars[0]) == strings.ToLower(chars[1]) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}