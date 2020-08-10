package main

import (
	"math"
	"math/rand"

	"bohao"

	"os"

	"github.com/go-echarts/go-echarts/charts"
	"gorgonia.org/gorgonia"
)

func function(x float64) (y float64) {
	//y = 3*math.Sin(x) - 0.5 + 5*rand.Float64() + 8
	y = 0.5*x*x*x + 0.3*math.Pow(x, 2) - x - 1 + 30*rand.Float64() + 8
	return
}

func create_xy() (x, y [][]float64) {
	for i := -10.; i < 10.; i += 0.1 {
		x = append(x, []float64{float64(i)})
		y = append(y, []float64{function(float64(i))})
	}
	return
}

func test() {
	x, y := create_xy()

	// init graph
	g := gorgonia.NewGraph()

	// setup network struction
	S := bohao.NetworkStruction{
		Neuron:     []int{1, 50, 1},                                     // the front one should be input shape and  the last one should be output shape
		Dropout:    []float64{0., 0.},                                   // set each dropout layer
		Act:        []bohao.ActivationFunc{gorgonia.Mish, bohao.Linear}, // can use act func directly from outside
		Bias:       true,
		Normal:     true,
		NormalSize: 1.,
	}

	// create NN
	m := bohao.NewNN(g, S)

	// init training parameter
	para := bohao.InitParameter()
	para.Solver = "Adam"
	para.Epoches = 2000
	para.Lr = 0.1
	para.Lossfunc = bohao.RMSError

	// fit training data
	m.Fit(x, y, para)

	// set test data into NN
	pred := m.Predict(x)

	h, _ := os.Create("plot.html")
	s := charts.NewScatter()
	plotdata1 := bohao.Transpose_float(append(bohao.Transpose_float(x), bohao.Transpose_float(y)...))
	plotdata2 := bohao.Transpose_float(append(bohao.Transpose_float(x), bohao.Transpose_float(pred)...))

	s.AddYAxis("y", plotdata1)
	s.AddYAxis("pred", plotdata2)

	s.Render(h)
}

func main() {
	test()

}
