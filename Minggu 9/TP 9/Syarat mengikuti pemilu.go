package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string
	
	fmt.Print("Masukkan umur warga: ")
	fmt.Scanln(&umur)

	fmt.Print("Masukkan kewarganegaraan Warga (WNI/WNA): ")
	fmt.Scanln(&kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Warga bisa mengikuti pemilu")
	} else {
		if umur < 17 {
			fmt.Println("Warga tidak bisa mengikuti pemilu karena umur Anda kurang dari 17 tahun.")
		}
		if kewarganegaraan != "WNI" {
			fmt.Println("Warga tidak bisa mengikuti pemilu karena bukan WNI.")
		}
	}
}