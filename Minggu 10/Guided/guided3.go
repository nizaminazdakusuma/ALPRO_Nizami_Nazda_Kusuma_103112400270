package main
import (
	"fmt"
	"strings"
)
func main() {
	var huruf string

	fmt.Print("Masukan huruf: ")
	fmt.Scanln(&huruf)

	huruf = strings . ToUpper(huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf vokal")
	} else if len(huruf) == 1 && (huruf >= "B" && huruf <= "Z") {
		fmt.Println("Huruf konsonan")
	} else {
		fmt.Println("Bukan huruf, Masukan huruf ")
	}
}