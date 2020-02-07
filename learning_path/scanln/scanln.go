package main 

import(
	"fmt"
)

func main() {
	var x, y  int
	fmt.Println("Please enter two integer with a blank between them:")
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)

}	