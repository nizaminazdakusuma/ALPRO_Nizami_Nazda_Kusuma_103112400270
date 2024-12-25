package main

import (
	"fmt"
	"strings"
)

func main() {
	masukan := []string{
		"Inigolang.",
		"Gilinggiling.",
		"Inigolanggolanggo.",
	}

	for i, input := range masukan {

		count := countGo(input)

		fmt.Printf("%d. %s -> %d\n", i+1, input, count)
	}
}

func countGo(str string) int {
	str = strings.TrimSuffix(str, ".")

	count := strings.Count(str, "go")

	return count
}