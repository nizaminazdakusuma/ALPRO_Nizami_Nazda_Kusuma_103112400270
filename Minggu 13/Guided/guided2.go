package main

import "fmt"

func main() {
	var bilangan int
	var continueLoop bool
	for continueLoop = true; continueLoop; {
		fmt.Scan(&bilangan)
		continueLoop = bilangan <= 0
	}

	fmt.Printf("%d adalah bilangan bulat positif\n", bilangan)
}
