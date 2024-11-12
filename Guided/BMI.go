package main
import "fmt"
func main() {
 var beratBadan, tinggiBadan, bmi float64

 // Memasukan berat badan dan tinggi badan
 fmt.Scan(&beratBadan, &tinggiBadan)

 // Menghitung BMI
 bmi = beratBadan / (tinggiBadan * tinggiBadan)

 // Menampilkan hasil BMI
 fmt.Printf("%.2f", bmi)
}