package main

import (
	"fmt"
	"math"
	)

func main(){
	var x [5] float64
	for i := 0; i < 5; i++ {
		var f = math.Pow(float64(i), 2)
		x[i] = f
	}
	fmt.Println(x)

	for _, value := range x{
		fmt.Println(value)
	}

	// short representation 
	y := [5]float64{
		1, 
		2, 
		3, 
		4, 
		5,	// Notice: remain this comma.
	}
	fmt.Println(y)
}
