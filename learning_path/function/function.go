package main

import (
	"fmt"
)

//or func add(x int, y int)
func add(x, y int) int {
	return x + y
}

//multi return
func swap(x, y string) (string, string) {
	return y, x
}

//name return
func split(sum int) (x, y int) {
	x = sum + 1
	y = sum - 1
	return
}

func main() {
	fmt.Println(add(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
