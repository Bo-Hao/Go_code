package main

import (
	"fmt"

	"gorgonia.org/gorgonia/examples/mnist"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var inputs, targets tensor.Tensor
var err error
var dt tensor.Dtype = gorgonia.Float64

type nn struct {
	g *gorgonia.ExprGraph
	w0, w1, w2 *gorgonia.Node 
	out *gorgonia.Node 
	preVal gorgonia.Value
}

func newNN(g *gorgonia.ExprGraph) *nn {
	//create node for w/weight
	w0 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(784, 300), gorgonia.WithName("w0"), gorgonia.WithInit((gorgonia.GlorotN(1.0))))
	w1 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(300, 100), gorgonia.WithName("w1"), gorgonia.WithInit((gorgonia.GlorotN(1.0))))
	w2 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(100, 10), gorgonia.WithName("w2"), gorgonia.WithInit((gorgonia.GlorotN(1.0))))

	return &nn{
		g : g, 
		w0 : w0, 
		w1 : w1, 
		w2 : w2, 
	}
}

func (m *nn) fwd(x *gorgonia.Node) {
	var l0, l1, l2 *gorgonia.Node
	var l0dot, l1dot *gorgonia.Node 

	//set first layer to be copy of input 
	l0 = x

	//dot product of l0 and w0, use as input for ReLU
	if l0dot, err = gorgonia.Mul(l0, m.w0); err != nil{
		panic(err)
	}

	//build hidden layer out of result 
	l1 = gorgonia.Must(gorgonia.Rectify(l0dot))

	//more layers
	if l1dot, err = gorgonia.Mul(l1, m.w1); err != nil{
		panic(err)
	}
	l2 = gorgonia.Must(gorgonia.Rectify(l1dot))

	var out *gorgonia.Node 
	if out, err = gorgonia.Mul(l2, m.w2); err != nil {
		panic(err)
	}

	m.out, err = gorgonia.SoftMax(out)
	gorgonia.Read(m.out, &m.preVal)
	return
}

func main() {
	inputs, targets, err = mnist.Load("train", "./mnist", gorgonia.Float64)
	
	g := gorgonia.NewGraph()
	m := newNN(g)

	x := gorgonia.NewMatrix(
		g, 
		dt, 
		gorgonia.WithName("X"),
		gorgonia.WithShape(784, )
	)
	m.fwd()

	losses, err := gorgonia.HadamardProd(m.out, y)
	if err != nil{
		panic(err)
	}

	cost := gorgonia.Must(gorgonia.Mean(losses))
	cost = gorgonia.Must(gorgonia.Neg(cost))

	//we wanna track costs 
	var costVal gorgonia.Value 
	gorgonia.Read(cost, &costVal)




}
