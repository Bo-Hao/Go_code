package main 

import(
	"fmt"
	"math"
)



type DeltaFunc func(x, y, w []float64) (G []float64)


func DF(x, y, w []float64) {

}

func GradientDescent(x, y, w0 []float64, eta float64, deltaf DeltaFunc) (w1 []float64) {
	w1 = make([]float64, len(w0))
	for i := 0; i < len(w0); i ++{
		w1[i] = w0[i] - eta*deltaf(x, y, w0)[i]
	}
	return w1
}

func main(){
	fmt.Println(math.Sqrt(4.))

}