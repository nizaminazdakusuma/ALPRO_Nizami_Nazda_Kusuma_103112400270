package main
import "fmt"
func main() {
 var bilangan, d1, d2, d3 int

 // Memasukan bilangan
 fmt.Scan(&bilangan)

 // Menentukan apakah bilangan semakin besar atau tidak
 d1 = bilangan / 100
 d2 = bilangan % 100 / 10
 d3 = bilangan % 100 % 10

 // Menampilkan apakah hasil semakin besar atau tidak
 fmt.Println(d1 <= d2 && d2 <= d3)
}