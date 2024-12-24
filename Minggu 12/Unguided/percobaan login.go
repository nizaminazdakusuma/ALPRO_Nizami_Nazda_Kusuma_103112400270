package main

import (
	"fmt"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	var username, password string
	failedAttempts := 0

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == correctUsername && password == correctPassword {
			break
		} else {
			failedAttempts++
			fmt.Println("Username atau password salah. Silakan coba lagi.")
		}
	}

	fmt.Printf("%d percobaan gagal login\n", failedAttempts)
}