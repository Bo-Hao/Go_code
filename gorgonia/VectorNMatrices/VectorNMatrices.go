// The goal here is to create a graph that will compute z = Wx, note that W is an nxn matrix.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	g := G.NewGraph()
	//matrix
	matB := []float64{0.9, 0.7, 0.4, 0.2}                                //step1
	matT := tensor.New(tensor.WithBacking(matB), tensor.WithShape(2, 2)) //step2
	mat := G.NewMatrix(
		g,
		tensor.Float64,
		G.WithName("W"),
		G.WithShape(2, 2),
		G.WithValue(matT),
	) //step3

	b := G.NewScalar(
		g,
		tensor.Float64,
		G.WithName("b"),
		G.WithValue(3.0),
	)

	//vector
	vecB := []float64{5, 7}
	vecT := tensor.New(tensor.WithBacking(vecB), tensor.WithShape(2))
	vec := G.NewVector(
		g,
		tensor.Float64,
		G.WithName("x"),
		G.WithShape(2),
		G.WithValue(vecT),
	)

	z, err := G.Add(G.Must(G.Mul(mat, vec)), b)
	if err != nil {
		panic(err)
	}
	machine := G.NewTapeMachine(g)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}

	fmt.Println(z.Value().Data()) // Output: [9.4 3.4]

	ioutil.WriteFile("simple_graph.dot", []byte(g.ToDot()), 0644)
}
