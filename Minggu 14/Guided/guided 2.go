package main
import "fmt"

func main() {
	var bilangan1, bilangan2, bilangan3, max, min int

	fmt.Scan(&bilangan1, &bilangan2, &bilangan3)

	if bilangan1 > bilangan2 {
		max = bilangan1
		min = bilangan2
	}else {
		max = bilangan2
		min = bilangan1
	}
	if max < bilangan3 {
		min = bilangan3
	}
	if min > bilangan3 {
		min = bilangan3
	}
	fmt.Println("Terbesar", max)
	fmt.Println("Terkecil", min)

}