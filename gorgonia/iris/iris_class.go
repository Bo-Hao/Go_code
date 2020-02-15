package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

const raw_data string = "/Users/pengbohao/go/src/gorgonia.org/gorgonia/examples/iris/iris.csv"

var err error

func read_iris() [][]string {
	f, err := os.Open(raw_data)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	sample_size := len(record) - 1
	result := make([][]string, sample_size)
	var spp string = ""

	for i := 1; i < sample_size+1; i++ {
		switch record[i][4] {
		case "setosa":
			spp = "0"
		case "versicolor":
			spp = "1"
		case "virginica":
			spp = "2"
		}
		result[i-1] = append(record[i][:4], spp)
	}
	return result
}
func strtofloat(s string) float64 {
	var result float64
	if result, err = strconv.ParseFloat(s, 64); err != nil {
		log.Fatal(err)
	}
	return result
}

func train_test_split(data [][]string, p float64) ([]float64, []float64, []float64, []float64) {
	train_size := int(math.Round(float64(len(data)) * (1.0 - p)))
	test_size := len(data) - train_size
	var train_data, test_data []float64
	var train_label, test_label []float64

	idx := rand.Perm(train_size)[:test_size+1]
	sort.Ints(idx)
	j := 0
	var tmp1 []float64
	var tmp2 float64
	for i := 0; i < len(data); i++ {
		tmp1 = []float64{strtofloat(data[i][0]), strtofloat(data[i][1]), strtofloat(data[i][2]), strtofloat(data[i][3])}
		tmp2 = strtofloat(data[i][4])
		if i == int(idx[j]) && j < test_size {
			for k := 0; k < len(tmp1); k++ {
				test_data = append(test_data, tmp1[k])
			}
			test_label = append(test_label, tmp2)
			j++
		} else {
			for k := 0; k < len(tmp1); k++ {
				train_data = append(train_data, tmp1[k])
			}
			train_label = append(train_label, tmp2)
		}
	}
	return train_data, train_label, test_data, test_label
}

type nn struct {
	g              *gorgonia.ExprGraph
	w0, w1, w2, w3 *gorgonia.Node

	pred    *gorgonia.Node
	predVal gorgonia.Value
}

func newNN(g *gorgonia.ExprGraph) *nn {
	w0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(4, 3), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(3, 3), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	w2 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(3, 3), gorgonia.WithName("w2"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	w3 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(3, 1), gorgonia.WithName("w3"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))

	return &nn{
		g:  g,
		w0: w0,
		w1: w1,
		w2: w2,
		w3: w3,
	}

}

func (m *nn) forward(x *gorgonia.Node) (err error) {
	var l0, l1, l2, l3 *gorgonia.Node
	var l0dot, l1dot, l2dot, l3dot *gorgonia.Node

	l0 = x
	l0dot = gorgonia.Must(gorgonia.Mul(l0, m.w0))

	l1 = gorgonia.Must(gorgonia.Rectify(l0dot))
	l1dot = gorgonia.Must(gorgonia.Mul(l1, m.w1))

	l2 = gorgonia.Must(gorgonia.Rectify(l1dot))
	l2dot = gorgonia.Must(gorgonia.Mul(l2, m.w2))

	l3 = gorgonia.Must(gorgonia.Rectify(l2dot))
	l3dot = gorgonia.Must(gorgonia.Mul(l3, m.w3))

	m.pred = l3dot

	gorgonia.Read(m.pred, &m.predVal)
	return
}

func (m *nn) learnables() gorgonia.Nodes {
	return gorgonia.Nodes{m.w0, m.w1, m.w2, m.w3}
}

func main() {
	rand.Seed(1377)
	var epoches int = 10

	g := gorgonia.NewGraph()
	m := newNN(g)

	csvdata := read_iris()
	train_data, train_label, test_data, test_label := train_test_split(csvdata, 0.2)

	xT := tensor.New(tensor.WithBacking(train_data), tensor.WithShape(120, 4))
	xVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("x"), gorgonia.WithValue(xT))

	yT := tensor.New(tensor.WithBacking(train_label), tensor.WithShape(120, 1))
	yVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("y"), gorgonia.WithValue(yT))

	//define input output
	x := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(120, 4), gorgonia.WithName("x"))
	y := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(120, 1), gorgonia.WithName("y"))

	//forward pass
	if err = m.forward(x); err != nil {
		log.Fatal(err)
	}

	//define loss function
	losses := gorgonia.Must(gorgonia.Square(gorgonia.Must(gorgonia.Sub(m.pred, y))))
	cost := gorgonia.Must(gorgonia.Mean(losses))

	//record cost
	var costVal gorgonia.Value
	gorgonia.Read(cost, &costVal)

	//upgrade gradient
	if _, err = gorgonia.Grad(cost, m.learnables()...); err != nil {
		log.Fatal("Unable to upgrade gradient")
	}
	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...))
	solver := gorgonia.NewAdamSolver(gorgonia.WithLearnRate(0.01))
	defer vm.Close()

	for epoch := 0; epoch < epoches; epoch++ {
		gorgonia.UnsafeLet(x, xVal)
		gorgonia.UnsafeLet(y, yVal)

		vm.RunAll()
		solver.Step(gorgonia.NodesToValueGrads(m.learnables()))
		vm.Reset()

		log.Printf("Done!")
	}
	log.Printf("training finished!")
	log.Printf("testint start!")
	//run test

	testxT := tensor.New(tensor.WithBacking(test_data), tensor.WithShape(30, 4))
	xVal = gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("test x"), gorgonia.WithValue(testxT))
	//x = gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(30, 4), gorgonia.WithName("x"))

	/* testyT := tensor.New(tensor.WithBacking(test_label), tensor.WithShape(30, 1))
	testyVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("test y"), gorgonia.WithValue(testyT)) */

	fmt.Println(m.pred.Value().Data())
	gorgonia.Let(x, xVal)

	vm.RunAll()

	fmt.Println(m.pred.Value().Data())
	vm.Reset()
	fmt.Println(len(test_label))

}
