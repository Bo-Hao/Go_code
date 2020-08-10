package main

import (
	"bohao"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"gorgonia.org/gorgonia"
)

const raw_data string = "/Users/pengbohao/go/src/gorgonia.org/gorgonia/examples/iris/iris.csv"

func read_data() ([][]float64, []string) {
	f, _ := os.Open(raw_data)
	r := csv.NewReader(f)
	record, _ := r.ReadAll()
	record = record[1:]
	record_T := bohao.Transpose_str(record)
	y := record_T[len(record_T)-1]
	x := bohao.Transpose_float(bohao.ConvSliceFromStr2Float(record_T[:len(record_T)-1]))

	return x, y
}

func PCA_reduce() {
	fmt.Println("Start!")
	x, _ := read_data()
	pca := bohao.Cal_PCA(x)
	new_x := pca.Reduce_data_dim(x, "80%")
	fmt.Println(new_x)
	fmt.Println(pca.Explain)
}

func autoEncoder() {
	x, _ := read_data()

	g := gorgonia.NewGraph()

	S := bohao.NetworkStruction{
		Neuron:  []int{4, 2, 1, 2, 4},
		Dropout: []float64{0, 0, 0, 0, 0},
		Act:     []bohao.ActivationFunc{bohao.Linear, bohao.Linear, bohao.Linear, bohao.Linear, bohao.Linear},
	}
	para := bohao.InitParameter()
	para.Solver = "Adam"

	m := bohao.NewNN(g, S)

	m.Fit(x, x, para)

	//fmt.Println(m.Predict(x))

	m1 := m.Tear_apart(g, 0, 2)
	fmt.Println(m1.Predict(x))

}

func stackAutoEncoder() {
	x, y := read_data()

	x, _, _ = bohao.Normalized(x, 1.)

	g := gorgonia.NewGraph()

	S := bohao.AE_Struction{
		InputShape:  4,
		HiddenShape: 4,
		CoreShape:   2,
		Denoising:   0.3,
		Dropout:     []float64{0, 0, 0, 0},
		Acti:        []bohao.ActivationFunc{gorgonia.Mish, bohao.Linear, bohao.Linear, gorgonia.Mish},

		Normal:     true,
		NormalSize: 1.,
		L1reg:      0.,
		L2reg:      0.,
	}
	para := bohao.InitParameter()
	para.Solver = "Adam"
	para.BatchSize = 1000
	para.Epoches = 500
	para.Lr = 0.1

	m := bohao.NewAE(g, S)

	m.Fit(x, para)
	core := m.Encode(x)
	fmt.Println(m.B1.Value(), m.B2.Value(), m.B3.Value())
	//pred := m.Decode(core)
	/* for i := 0; i < len(pred); i++ {
		fmt.Println(pred[i], x[i])
	} */

	new_y, _ := bohao.EncodeYLabel(y)

	scatter := charts.NewScatter()

	var pt1, pt2, pt0 [][]float64
	for i := 0; i < len(core); i++ {
		switch new_y[i] {
		case 0:
			pt0 = append(pt0, core[i])
		case 1:
			pt1 = append(pt1, core[i])
		case 2:
			pt2 = append(pt2, core[i])
		}
	}
	scatter.AddYAxis("pt0", pt0)
	scatter.AddYAxis("pt1", pt1)
	scatter.AddYAxis("pt2", pt2)

	h, _ := os.Create("/Users/pengbohao/Go_code/Test_dimReduce/plot.html")
	scatter.Render(h)

}
func main() {
	//autoEncoder()
	stackAutoEncoder()
}
