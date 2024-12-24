package main
import (
	"fmt"
)
func main() {
	var nilai float64

	fmt.Print("Masukan nilai Mahasiswa: ")
	fmt.Scanln(&nilai)

	if nilai > 90 {
		fmt.Println("indeks: A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("indeks: AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("indeks: B")
	} else {
		fmt.Println("indeks: C")
	}
}