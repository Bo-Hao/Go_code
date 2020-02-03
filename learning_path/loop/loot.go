package main

import (
	"fmt"
	"math"
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//short if
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if and else
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

//switch
func choose(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2, 3, 4:
		fmt.Println("2, 3 or 4")
	}
}

func anotherchoose(i int){
	switch{
	case i > 5:
		fmt.Println(i, " > 5")
	case i < 5:
		fmt.Println(i, " < 5")
	case i == 5:
		fmt.Print(i, " = 5")
	}
}

func main() {
	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//infinite loop
	/* for {
	} */

	//if
	fmt.Println(sqrt(2), sqrt(-4))

	//short if
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//if and else
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20))

	//switch
	choose(1)
    choose(3)
    choose(4)

    anotherchoose(6)

}
