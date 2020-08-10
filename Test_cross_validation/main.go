package main

import (
	"bohao"
	"math/rand"

	"gorgonia.org/gorgonia"
)

const raw_data string = "/Users/pengbohao/go/src/gorgonia.org/gorgonia/examples/iris/iris.csv"

func read_data() (x, y [][]float64) {

	for i := 0; i < 100; i++ {
		x = append(x, []float64{float64(i) * 0.5, float64(i) - 6, 5 * float64(i) * rand.Float64()})
		y = append(y, []float64{float64(i)*0.7 + 2. + rand.Float64()*10.})
	}

	return x, y
}

func main() {
	x, y := read_data()
	// construct a three layers network.

	P := bohao.Parameter{
		Lr:        0.01,
		Epoches:   200,
		BatchSize: 625,
		Solver:    "Adam",
		Lossfunc:  bohao.RMSError,
	}

	N := bohao.NetworkStruction{
		Neuron:     []int{3, 20, 1},
		Dropout:    []float64{0.2, 0.2},
		Act:        []bohao.ActivationFunc{gorgonia.Mish, bohao.Linear},
		Bias:       true,
		Normal:     true,
		NormalSize: 1.,
		L1reg:      0.,
		L2reg:      0.,
	}

	g := gorgonia.NewGraph()
	m := bohao.NewNN(g, N)
	m.Fit(x, y, P)

}
