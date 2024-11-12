package main

import "fmt"

func main() {
	var ( 
		apel, pisang, jeruk string
	    temp string
	)

	// Memasukan dan menghitung string satu, dua, dan tiga
	fmt.Print("Apel: ")
	fmt.Scanln(&apel)
	fmt.Print("Pisang: ")
	fmt.Scanln(&pisang)
	fmt.Print("Jeruk: ")
	fmt.Scanln(&jeruk)
	fmt.Println("Output awal = " + apel + " " + pisang + " " + jeruk)
	temp = apel
	apel = pisang
	pisang = jeruk
	jeruk = temp

	// Memasukan hasil dari string diatas
	fmt.Println("Output akhir = " + apel + " " + pisang + " " + jeruk)
}