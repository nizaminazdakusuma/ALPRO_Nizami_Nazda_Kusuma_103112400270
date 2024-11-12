package main
import "fmt"
func main() {
	var x, fx float64

	// Memasukan nilai x
	fmt.Print("Memasukan nilai x: ")
	fmt.Scan(&x)

	// Menghitung nilai f(X)
	fx = 2 / (x+5) + 5
	fmt.Println(fx)
}