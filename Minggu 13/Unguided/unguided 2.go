package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	if false {
		fmt.Println("Silakan masukkan bilangan desimal yang valid.")
		return
	}

	roundedUp := math.Ceil(input)

	for i := input; i <= roundedUp; i += 0.1 {
		fmt.Printf("%.1f\n", i)
	}
}