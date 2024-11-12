package main
import (
	"fmt"
)
func main() {
	var bmi, tinggiBadan float64

	// Masukan tinggi badan dan bmi
	fmt.Print("Masukan nilai BMI: ")
	_, err := fmt.Scan(&bmi)
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan memasukan angka.")
		return
	}

	fmt.Print("Memasukan tinggi badan dalam meter: ")
	_, err = fmt.Scan(&tinggiBadan)
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan memasukan angka.")
		return
	}

	// Menghitung berat badan
	beratBadan := bmi * (tinggiBadan * tinggiBadan)

	// Menampilkan hasil dari penjumlahan di atas
	fmt.Printf(" Berat badan seorang dengan BMI %.f dan tinggi badan %.f meter adalah %f kg.\n", bmi, tinggiBadan, beratBadan)

}