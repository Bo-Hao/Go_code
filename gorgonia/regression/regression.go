package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var err error

type nn struct {
	g  *gorgonia.ExprGraph
	x  *gorgonia.Node
	w0 *gorgonia.Node

	pred    *gorgonia.Node
	predVal gorgonia.Value
}

func newNN(g *gorgonia.ExprGraph) *nn {
	w0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	//w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))

	return &nn{
		g:  g,
		w0: w0,
		//w1: w1,
	}
}

func newNN_load(g *gorgonia.ExprGraph) *nn {
	//w0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	//w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))

	weight, err := readmodel()
	if err != nil {
		log.Fatal("can't load")
	}
	w0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 1), gorgonia.WithName("w0"), gorgonia.WithValue(weight))
	return &nn{
		g:  g,
		w0: w0,
		//w1: w1,
	}

}

func (m *nn) forward(x *gorgonia.Node) (err error) {
	var l0 *gorgonia.Node
	var l0dot *gorgonia.Node

	l0 = x
	l0dot = gorgonia.Must(gorgonia.Mul(l0, m.w0))

	m.pred = l0dot

	gorgonia.Read(m.pred, &m.predVal)
	return
}

func (m *nn) learnables() gorgonia.Nodes {
	return gorgonia.Nodes{m.w0}
}

func generate_data() (xset []float64, yset []float64) {

	for i := 0.0; i < 10; i = i + 1.0 {
		xset = append(xset, i)
		yset = append(yset, i /* + rand.Float64() */)
	}
	return
}

func save(nodes []*gorgonia.Node) error {
	f, err := os.Create("example_gorgonia")
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	for _, node := range nodes {
		err := enc.Encode(node.Value())
		if err != nil {
			return err
		}
	}
	return nil
}

func readmodel() (tensor.Tensor, error) {
	f, err := os.Open("example_gorgonia")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	var w0 *tensor.Dense
	log.Println("decoding xT")
	err = dec.Decode(&w0)
	if err != nil {
		return nil, err
	}

	return w0, nil
}

func main1() {
	rand.Seed(1377)
	var epoches int = 100

	g := gorgonia.NewGraph()
	m := newNN(g)

	xset, yset := generate_data()
	xlen := len(xset)

	xT := tensor.New(tensor.WithBacking(xset), tensor.WithShape(xlen, 1))
	xVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("X"), gorgonia.WithValue(xT))

	yT := tensor.New(tensor.WithBacking(yset), tensor.WithShape(xlen, 1))
	yVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("y"), gorgonia.WithValue(yT))

	//define input output
	x := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(xlen, 1), gorgonia.WithName("X"))
	y := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(xlen, 1), gorgonia.WithName("y"))

	//forward pass
	if err = m.forward(x); err != nil {
		log.Fatal(err)
	}

	//define loss function
	losses := gorgonia.Must(gorgonia.Neg(gorgonia.Must(gorgonia.Square(gorgonia.Must(gorgonia.Sub(m.pred, y))))))
	cost := gorgonia.Must(gorgonia.Mean(losses))

	//record cost
	var costVal gorgonia.Value
	gorgonia.Read(cost, &costVal)

	//upgrade gradient
	if _, err = gorgonia.Grad(cost, m.learnables()...); err != nil {
		log.Fatal("Unable to upgrade gradient")
	}
	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...))
	solver := gorgonia.NewAdamSolver(gorgonia.WithLearnRate(0.5))
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
	err = save([]*gorgonia.Node{m.w0})

	fmt.Println(m.w0.Value())
}

func main2() {
	g := gorgonia.NewGraph()
	m := newNN_load(g)

	var xtest []float64
	for i := 10.5; i < 20.5; i = i + 1.0 {
		xtest = append(xtest, i)
	}
	xlen := len(xtest)

	xT := tensor.New(tensor.WithBacking(xtest), tensor.WithShape(xlen, 1))
	xVal := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithName("X"), gorgonia.WithValue(xT))

	//define input output
	x := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(xlen, 1), gorgonia.WithName("X"))

	//forward pass
	if err = m.forward(x); err != nil {
		log.Fatal(err)
	}

	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...))
	fmt.Println(m.w0.Value())
	gorgonia.Let(x, xVal)
	vm.RunAll()
	log.Printf("done")
	fmt.Println(m.pred.Value().Data())

}

func main() {
	main1()
	main2()
}
