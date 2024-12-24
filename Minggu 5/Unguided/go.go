package main
import "fmt"
func main() {
	var x float64

	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)

	y := (3*x-5) * (2*x+1)

	fmt.Printf("Nilain y dari x = %.2f adalah %.2f\n", x, y)
}