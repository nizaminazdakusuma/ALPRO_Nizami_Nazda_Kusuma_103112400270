package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan panjang sisi c: ")
	fmt.Scanln(&c)

	sisi := []int{a, b, c}
	sort.Ints(sisi)

	if sisi[0]+sisi[1] <= sisi[2] {
		fmt.Println("Bukan segitiga")
	} else {
		if sisi[0] == sisi[1] && sisi[1] == sisi[2] {
			fmt.Println("Segitiga sama sisi")
		} else if sisi[0] == sisi[1] || sisi[1] == sisi[2] {
			fmt.Println("Segitiga sama kaki")
		} else if sisi[0]*sisi[0]+sisi[1]*sisi[1] == sisi[2]*sisi[2] {
			fmt.Println("Segitiga siku-siku")
		} else {
			fmt.Println("Segitiga sembarang")
		}
	}
}