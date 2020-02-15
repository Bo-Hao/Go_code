package main

import (
	"fmt"
	"math/rand"

	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type nn struct {
	g      *ExprGraph
	w0, w1 *Node

	pred    *Node
	predVal Value
}

var err error

func (m *nn) fwd(x *Node) (err error) {
	var l0, l1 *Node

	// Set first layer to be copy of input
	l0 = x

	// Dot product of l0 and w0, use as input for Sigmoid
	l0dot := Must(Mul(l0, m.w0))

	// Build hidden layer out of result
	l1 = Must(Sigmoid(l0dot))

	m.pred = l1
	Read(m.pred, &m.predVal)
	return
}

func newNN(g *ExprGraph) *nn {
	//create node for w/weights
	wB := tensor.Random(tensor.Float64, 3) //[]float64{-0.1, 0.4, -1.0}
	wT := tensor.New(tensor.WithBacking(wB), tensor.WithShape(3, 1))
	w0 := NewMatrix(g,
		tensor.Float64,
		WithName("w"),
		WithShape(3, 1),
		WithValue(wT),
	)
	return &nn{
		g:  g,
		w0: w0,
	}
}

func (m *nn) learnables() Nodes {
	return Nodes{m.w0}
}

func main() {
	rand.Seed(31337)

	//create graph and network
	g := NewGraph()
	m := newNN(g)

	//set input x to network
	xB := []float64{0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1}
	xT := tensor.New(tensor.WithBacking(xB), tensor.WithShape(4, 3))
	x := NewMatrix(
		g,
		tensor.Float64,
		WithName("X"),
		WithShape(4, 3),
		WithValue(xT),
	)

	//define validation dataset
	yB := []float64{0, 0, 1, 1}
	yT := tensor.New(tensor.WithBacking(yB), tensor.WithShape(4, 1))
	y := NewMatrix(
		g,
		tensor.Float64,
		WithName("y"),
		WithShape(4, 1),
		WithValue(yT),
	)

	//run forward pass
	if err = m.fwd(x); err != nil {
		panic(err)
	}

	//calculate cost w/MSE
	losses := Must(Sub(y, m.pred))
	square := Must(Square(losses))
	cost := Must(Mean(square))

	//do the gradient updates
	if _, err = Grad(cost, m.learnables()...); err != nil {
		panic(err)
	}

	//Instantiate VM and Solver
	vm := NewTapeMachine(g, BindDualValues(m.learnables()...))
	solver := NewVanillaSolver(WithLearnRate(1.0))
	//solver := NewRMSPropSolver(WithLearnRate(0.01))

	for i := 0; i < 1000; i++ {
		vm.Reset()
		vm.RunAll()
		solver.Step(NodesToValueGrads(m.learnables()))
		fmt.Println("\n State at iter ", i)
		fmt.Println("Cost: \n", cost.Value())
		fmt.Println("Weights: \n", m.w0.Value())
		vm.Reset()
	}
	fmt.Println("Output after Training: \n", m.predVal)

}
