package main

import (
	"bohao"
	"bohao/pivotalsampling"
	"math/rand"
	//"time"
)

func rep(p float64, N int) (r []float64) {

	for i := 0; i < N; i++ {
		r = append(r, p)
	}
	return
}

func main() {
	//rand.Seed(time.Now().Unix())
	rand.Seed(1234567)
	N := 625
	n := 90
	p := rep(float64(n)/float64(N), N)

	var X [][]float64
	for i := 0; i < N; i++ {
		X = append(X, []float64{rand.Float64(), rand.Float64()})
	}

	sample := pivotalsampling.LocalPivotalSampling(X, p)
	var sampleData [][]float64
	for i := 0; i < len(sample); i++ {
		sampleData = append(sampleData, X[sample[i]])
	}

	s_T := bohao.Transpose_float(sampleData)
	X_T := bohao.Transpose_float(X)
	var draw [][]float64

	draw = append(draw, X_T[0])
	draw = append(draw, X_T[1])
	draw = append(draw, s_T[0])
	draw = append(draw, s_T[1])

	bohao.DrawXYScatterPlot(draw, "/Users/pengbohao/Go_code/Test local pivotal sampling/plot.html")
}
