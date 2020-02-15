package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	var x, y, z *gorgonia.Node
	var err error

	g := gorgonia.NewGraph()
	x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))

	if z, err = gorgonia.Add(x, y); err != nil {
		panic(err)
	}

	vm := gorgonia.NewTapeMachine(g)
	defer vm.Close()
	var Val float64
	Val = 2.0
	gorgonia.Let(x, Val)
	gorgonia.Let(y, 3.0)
	vm.RunAll()
	fmt.Println(z.Value())

	Val = 3.0
	gorgonia.Let(x, Val)
	gorgonia.Let(y, 3.0)
	vm.RunAll()
	fmt.Println(z.Value())

}
