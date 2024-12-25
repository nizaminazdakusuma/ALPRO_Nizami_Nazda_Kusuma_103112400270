package main

import "fmt"

func main() {
	var kata string
	var perulangan int

	fmt.Scan(&kata, &perulangan)
	counter := 0 //penghittungan dimulai dari 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= perulangan)
	}
}