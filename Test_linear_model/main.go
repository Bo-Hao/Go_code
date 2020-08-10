package main

import (
	"bohao"
	"fmt"
)

func main() {
	x := [][]float64{{1}, {2}, {3}}
	y := []float64{1, 2, 3}
	x_ := [][]float64{{4}, {5}, {6}, {7}}
	R := bohao.Multi_Regression(x, y)
	y_ := R.Predict(x_)
	fmt.Println(y_)
}
