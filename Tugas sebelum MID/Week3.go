package main
import "fmt"
func main() {
	var bilangan, d1, d2, d3 int // deklarasi variabel

	fmt.Print("Masukan bilangan tiap digit: ")
	fmt.Scan(&bilangan)

	d1 = bilangan / 1000      // mengambil digit pertama
	d2 = (bilangan %100) / 10 // mengambil digit kedua
	d3 = bilangan % 10        // mengambil digit ketiga

	fmt.Println(d1 <= d2 && d2 <= d3) // mengecek apakah digit dalam urutan
}