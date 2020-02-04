package main

import (
	"fmt"
)

//Simply, the array without declaring its length is slice. 
func main(){
	var x [] float64 
	fmt.Println(x)

	// define length and capability with 5
	x = make([]float64, 5)
	fmt.Println(x)

	// define length and capability with 5 and 10
	x = make([]float64, 5, 10)
	fmt.Println(x)

	//another way to create slice
	arr := [9]float64{1,2,3,4,5,6,7,8,9}
	y := arr[0:5]
	fmt.Println(y)

	//append 
	slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	
	//copy: it can copy a slice to another slice 
	slice3 := [] int {1, 2, 3}
	slice4 := make([]int, 2, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}