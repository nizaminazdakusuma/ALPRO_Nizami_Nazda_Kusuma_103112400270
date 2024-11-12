package main
import "fmt"
func main() {
 var detik, jam, menit int

 // Memasukan detik
 fmt.Scan(&detik)

 // Mengubah detik ke jam, menit, dan detik
 jam = detik / 3600
 menit = (detik % 3600) / 60
 detik = detik % 60

 // Menampilkan hasil konversi
 fmt.Println(jam, "jam", menit, "menit dan", detik,"detik")
}