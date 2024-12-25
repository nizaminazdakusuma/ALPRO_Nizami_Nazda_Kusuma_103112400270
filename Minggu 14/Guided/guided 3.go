package main
import "fmt"

func main() {
	var bilangan int

	fmt.Print("masukan bilangan bulat: ")
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i," ")
		}
	}

	fmt.Println()
}