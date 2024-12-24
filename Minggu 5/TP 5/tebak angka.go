package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 5

	fmt.Println("Tebak angka antara 1 hingga 100. Anda memiliki 5 kesempatan.")

	for attempts > 0 {
		fmt.Print("Masukkan tebakan Anda: ")
		var guess int
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println("Input tidak valid. Harap masukkan bilangan bulat.")
			continue
		}

		attempts--
		if guess == target {
			fmt.Printf("Selamat! Anda menebak angka %d dengan benar.\n", target)
			return
		} else if guess < target {
			fmt.Println("Tebakan Anda terlalu kecil.")
		} else {
			fmt.Println("Tebakan Anda terlalu besar.")
		}

		if attempts == 0 {
			fmt.Printf("Maaf, Anda telah kehabisan kesempatan. Angka yang benar adalah %d.\n", target)
		}
	}
}