package main

import (
	"fmt"
)

func f() {
	fmt.Println("test")
	panic(1)
	fmt.Println("test2")
}

// there is no try...except in go.
func main() {
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()
	f()
}
