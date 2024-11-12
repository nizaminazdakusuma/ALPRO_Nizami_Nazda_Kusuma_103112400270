package main
import "fmt"
func main() {
	var alas, tinggi, luas float64

	// Memasukan alas dan tinggi
	fmt.Print("Memasukan alas dan tinggi: ")
	fmt.Scan(&alas, &tinggi)

	// Menghitung luas lingkaran
	luas = 0.5 * alas * tinggi
	fmt.Println(luas)
}