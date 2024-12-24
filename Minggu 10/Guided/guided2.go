package main
import (
	"fmt"
)
func main() {
	var umur int

	fmt.Print("Masukan umur penduduk: ")
	fmt.Scan(&umur)

	if umur >= 17 {
		fmt.Println("Bisa membuat KTP")
	} else if umur < 17 {
		fmt.Println("Belum bisa membuat KTP")
	}
}