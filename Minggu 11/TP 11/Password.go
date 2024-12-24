package main

import (
	"fmt"
)

func main() {
	const correctPassword = "telutizen" 
	var masukanPassword string
	attempts := 0
	const maxAttempts = 3

	fmt.Println("Masukkan password Anda.")

	for attempts < maxAttempts {
		fmt.Printf("Kesempatan %d dari %d: ", attempts+1, maxAttempts)
		fmt.Scan(&masukanPassword)

		if masukanPassword == correctPassword {
			fmt.Println("Login berhasil!")
			return 
		} else {
			fmt.Println("Password salah.")
			attempts++
		}
	}
	fmt.Println("Login ditolak.")
}