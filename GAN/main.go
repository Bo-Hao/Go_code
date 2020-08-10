package main

import (
	//"bohao"
	"bohao"
	"fmt"
	"log"

	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func test_bias() {
	g := NewGraph()
	a := NewMatrix(g, tensor.Float64, WithShape(1, 1), WithName("b"), WithInit(Uniform(-1, 1)))
	//a := NewScalar(g, tensor.Float64, WithName("a"), WithValue(float64(10)))
	b := NewMatrix(g, tensor.Float64, WithShape(3, 1), WithName("b"), WithInit(Uniform(-1, 1)))
	fmt.Println(a.Shape(), b.Shape())
	fmt.Printf("a = %v\nb =\n%v\n", a.Value(), b.Value())

	ba, err := BroadcastAdd(b, a, nil, []byte{0})

	if err != nil {
		panic(err)
	}

	// Now, let's run the program
	machine := NewTapeMachine(g)
	defer machine.Close()
	if err = machine.RunAll(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("b +⃗ a =\n%v", ba.Value())
}
func main() {
	/* //bohao.Example_auto()
	//bohao.Example_NN()
	//bohao.Do_VAE()
	*/
	

	//bohao.PCA([][]float64{{1, 2, 3, 4, 5}, {1, 2, 4, 4, 5}, {3, 3, 4, 5, 6}})
	bohao.PCA([][]float64{{1, 2}, {1, 2}})
}
