package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()

	var a, b, c *gorgonia.Node
	var err error

	// define the expression
	a = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("a"))
	b = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("b"))
	c, err = gorgonia.Add(a, b)
	if err != nil {
		panic(err)
	}

	//create a VM to run the program on
	machine := gorgonia.NewTapeMachine(g)

	// set initial values then run
	gorgonia.Let(a, 1.0)
	gorgonia.Let(b, 2.0)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Value())
	// Output:3.0
}
