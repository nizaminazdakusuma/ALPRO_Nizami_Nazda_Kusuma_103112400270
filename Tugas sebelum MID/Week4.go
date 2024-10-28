package main
import "fmt"
func main() {
	var nama string
	var gajiPokok, tunjangan, potongan float64

	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan gaji pokok: ")
	fmt.Scanln(&gajiPokok)
	fmt.Print("Masukan tunjangan: ")
	fmt.Scanln(&tunjangan)
	fmt.Print("Masukan potongan: ")
	fmt.Scanln(&potongan)

	totalGaji := gajiPokok + tunjangan - potongan
	fmt.Printf("Total gaji karyawan: %.2f\n", totalGaji)
}